package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
type OCFDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OCFDeviceSpec   `json:"spec,omitempty"`
	Status OCFDeviceStatus `json:"status,omitempty"`
}

type OCFDeviceSpec struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Owned   bool   `json:"owned,omitempty"`
	OwnerID string `json:"ownerId,omitempty"`

	ResourceTypes []ResourceType `json:"resourceTypes,omitempty"`
}

type ResourceType struct {
	Name string `json:"name,omitempty"`
}

type OCFDeviceStatus struct {
	Options []Options `json:"options,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OCFDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []OCFDevice `json:"items,omitempty"`
}

type Options struct {
	CertIdentity     string        `long:"certIdentity"`
	DiscoveryTimeout time.Duration `long:"discoveryTimeout"`

	MfgCert       string `long:"mfgCert"`
	MfgKey        string `long:"mfgKey"`
	MfgTrustCA    string `long:"mfgTrustCA"`
	MfgTrustCAKey string `long:"mfgTrustCAKey"`

	IdentityCert              string `long:"identityCert"`
	IdentityKey               string `long:"identityKey"`
	IdentityIntermediateCA    string `long:"identityIntermediateCA"`
	IdentityIntermediateCAKey string `long:"identityIntermediateCAKey"`
	IdentityTrustCA           string `long:"identityTrustCA"`
	IdentityTrustCAKey        string `long:"identityTrustCAKey"`
}
