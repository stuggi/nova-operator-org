package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NovaComputeSpec defines the desired state of NovaCompute
// +k8s:openapi-gen=true
type NovaComputeSpec struct {
        // Label is the value of the 'daemon=' label to set on a node that should run the daemon
        Label string `json:"label"`

        // container image to run for the daemon
        NovaComputeImage string `json:"novaComputeImage"`
}

// NovaComputeStatus defines the observed state of NovaCompute
// +k8s:openapi-gen=true
type NovaComputeStatus struct {
        // Count is the number of nodes the daemon is deployed to
        Count int32 `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NovaCompute is the Schema for the novacomputes API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type NovaCompute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NovaComputeSpec   `json:"spec,omitempty"`
	Status NovaComputeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NovaComputeList contains a list of NovaCompute
type NovaComputeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NovaCompute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NovaCompute{}, &NovaComputeList{})
}
