package controller

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	nodesv1alpha1 "github.com/yourusername/tailscale-k3s-operator/api/v1alpha1"
	"github.com/yourusername/tailscale-k3s-operator/internal/provisioner"
	"github.com/yourusername/tailscale-k3s-operator/internal/tailscale"
)

// TailscaleK3sAgentReconciler reconciles a TailscaleK3sAgent object
type TailscaleK3sAgentReconciler struct {
    client.Client
    Scheme          *runtime.Scheme
    TailscaleClient *tailscale.Client
    Provisioner     *provisioner.Provisioner
}

func (r *TailscaleK3sAgentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    log := log.FromContext(ctx)

    var agent nodesv1alpha1.TailscaleK3sAgent
    if err := r.Get(ctx, req.NamespacedName, &agent); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // Check if the device is still available in Tailscale
    devices, err := r.TailscaleClient.ListDevices(ctx)
    if err != nil {
        log.Error(err, "failed to list Tailscale devices")
        return ctrl.Result{RequeueAfter: time.Minute}, nil
    }

    var matchingDevice *tailscale.Device
    for _, device := range devices {
        if device.ID == agent.Spec.TailscaleID {
            matchingDevice = &device
            break
        }
    }

    if matchingDevice == nil {
        log.Info("device not found in Tailscale", "id", agent.Spec.TailscaleID)
        agent.Status.Phase = "Failed"
        agent.Status.Error = "Device not found in Tailscale"
        if err := r.Status().Update(ctx, &agent); err != nil {
            return ctrl.Result{}, err
        }
        return ctrl.Result{}, nil
    }

    // Update status based on device state
    agent.Status.LastSeen = &metav1.Time{Time: matchingDevice.LastSeen}

    // If not yet provisioned, install K3s
    if agent.Status.Phase != "Ready" {
        err := r.Provisioner.InstallK3sAgent(
            ctx,
            agent.Spec.Hostname,
            agent.Spec.K3sServerURL,
            agent.Spec.K3sToken,
        )
        if err != nil {
            log.Error(err, "failed to install K3s agent")
            agent.Status.Phase = "Failed"
            agent.Status.Error = err.Error()
            if err := r.Status().Update(ctx, &agent); err != nil {
                return ctrl.Result{}, err
            }
            return ctrl.Result{RequeueAfter: time.Minute}, nil
        }

        agent.Status.Phase = "Ready"
        agent.Status.LastProvisioned = &metav1.Time{Time: time.Now()}
        agent.Status.Error = ""
    }

    if err := r.Status().Update(ctx, &agent); err != nil {
        return ctrl.Result{}, err
    }

    return ctrl.Result{RequeueAfter: time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TailscaleK3sAgentReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&nodesv1alpha1.TailscaleK3sAgent{}).
        Complete(r)
}