package provisioner

import (
	"context"
	"fmt"
	"os/exec"
)

// Provisioner handles K3s agent provisioning
type Provisioner struct{}

// NewProvisioner creates a new provisioner
func NewProvisioner() *Provisioner {
	return &Provisioner{}
}

// InstallK3sAgent installs K3s agent on the target host
func (p *Provisioner) InstallK3sAgent(ctx context.Context, hostname string, serverURL, token string) error {
	installCmd := fmt.Sprintf("curl -sfL https://get.k3s.io | K3S_URL='%s' K3S_TOKEN='%s' sh -", serverURL, token)

	cmd := exec.CommandContext(ctx, "tailscale", "ssh", hostname, installCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("installing k3s: %w, output: %s", err, string(output))
	}

	return nil
}
