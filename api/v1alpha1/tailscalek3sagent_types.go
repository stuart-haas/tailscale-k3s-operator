package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TailscaleK3sAgentSpec defines the desired state of TailscaleK3sAgent
type TailscaleK3sAgentSpec struct {
	// Tags are the Tailscale tags assigned to this node
	Tags []string `json:"tags"`

	// K3sServerURL is the URL of the K3s server to join
	K3sServerURL string `json:"k3sServerURL"`

	// K3sToken is the token used for K3s agent registration
	K3sToken string `json:"k3sToken,omitempty"`

	// ClientID is the OAuth client ID from Tailscale
	ClientID string `json:"clientId"`

	// ClientSecret is the OAuth client secret from Tailscale
	ClientSecret string `json:"clientSecret"`

	// TailscaleOrgName is your Tailscale organization name
	TailscaleOrgName string `json:"tailscaleOrgName"`
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

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
//+kubebuilder:printcolumn:name="Hostname",type=string,JSONPath=`.spec.hostname`
//+kubebuilder:printcolumn:name="IP",type=string,JSONPath=`.spec.ipAddress`
//+kubebuilder:printcolumn:name="LastSeen",type="string",JSONPath=`.status.lastSeen`
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// TailscaleK3sAgent is the Schema for the tailscalek3sagents API
type TailscaleK3sAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TailscaleK3sAgentSpec   `json:"spec,omitempty"`
	Status TailscaleK3sAgentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TailscaleK3sAgentList contains a list of TailscaleK3sAgent
type TailscaleK3sAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TailscaleK3sAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TailscaleK3sAgent{}, &TailscaleK3sAgentList{})
}
