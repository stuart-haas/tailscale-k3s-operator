package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TailscaleK3sAgentSpec defines the desired state of TailscaleK3sAgent
type TailscaleK3sAgentSpec struct {
    // TailscaleID is the unique identifier from Tailscale
    TailscaleID string `json:"tailscaleId"`

    // Hostname is the device hostname in Tailscale
    Hostname string `json:"hostname"`

    // IPAddress is the Tailscale IP address
    IPAddress string `json:"ipAddress"`

    // Tags are the Tailscale tags assigned to this node
    Tags []string `json:"tags"`

    // K3sServerURL is the URL of the K3s server to join
    K3sServerURL string `json:"k3sServerURL"`

    // K3sToken is the token used for K3s agent registration
    K3sToken string `json:"k3sToken,omitempty"`
}

// TailscaleK3sAgentStatus defines the observed state of TailscaleK3sAgent
type TailscaleK3sAgentStatus struct {
    // Phase represents the current state of the agent
    // +kubebuilder:validation:Enum=Pending;Provisioning;Ready;Failed
    Phase string `json:"phase"`

    // LastProvisioned is the timestamp of the last provisioning attempt
    LastProvisioned *metav1.Time `json:"lastProvisioned,omitempty"`

    // LastSeen is when the agent was last seen in Tailscale
    LastSeen *metav1.Time `json:"lastSeen,omitempty"`

    // K3sVersion is the version of K3s running on the agent
    K3sVersion string `json:"k3sVersion,omitempty"`

    // Error message if provisioning failed
    Error string `json:"error,omitempty"`
}
