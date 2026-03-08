/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkClassCapabilities defines the capabilities supported by a NetworkClass
type NetworkClassCapabilities struct {
	// SupportsIPv4 indicates whether this network class supports IPv4
	// +kubebuilder:validation:Optional
	SupportsIPv4 bool `json:"supportsIPv4,omitempty"`

	// SupportsIPv6 indicates whether this network class supports IPv6
	// +kubebuilder:validation:Optional
	SupportsIPv6 bool `json:"supportsIPv6,omitempty"`

	// SupportsDualStack indicates whether this network class supports dual-stack (IPv4+IPv6)
	// +kubebuilder:validation:Optional
	SupportsDualStack bool `json:"supportsDualStack,omitempty"`
}

// NetworkClassSpec defines the desired state of NetworkClass
type NetworkClassSpec struct {
	// ImplementationStrategy identifies the network implementation (e.g., "udn-net", "physical-net")
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	ImplementationStrategy string `json:"implementationStrategy"`

	// Capabilities describes the networking capabilities of this class
	// +kubebuilder:validation:Optional
	Capabilities NetworkClassCapabilities `json:"capabilities,omitempty"`

	// Constraints holds provider-specific constraints as key-value pairs
	// +kubebuilder:validation:Optional
	Constraints map[string]string `json:"constraints,omitempty"`
}

// NetworkClassState represents the state of a NetworkClass
// +kubebuilder:validation:Enum=Pending;Ready;Failed
type NetworkClassState string

const (
	// NetworkClassStatePending indicates the network class is being validated
	NetworkClassStatePending NetworkClassState = "Pending"

	// NetworkClassStateReady indicates the network class is ready for use
	NetworkClassStateReady NetworkClassState = "Ready"

	// NetworkClassStateFailed indicates the network class validation failed
	NetworkClassStateFailed NetworkClassState = "Failed"
)

// NetworkClassStatus defines the observed state of NetworkClass
type NetworkClassStatus struct {
	// State provides the current state of the NetworkClass
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Pending;Ready;Failed
	State NetworkClassState `json:"state,omitempty"`

	// Message provides human-readable status or error information
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Type=string
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=netclass
// +kubebuilder:printcolumn:name="Strategy",type=string,JSONPath=`.spec.implementationStrategy`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// NetworkClass is the Schema for the networkclasses API
type NetworkClass struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of NetworkClass
	// +required
	Spec NetworkClassSpec `json:"spec"`

	// status defines the observed state of NetworkClass
	// +optional
	Status NetworkClassStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// NetworkClassList contains a list of NetworkClass
type NetworkClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkClass `json:"items"`
}

// GetName returns the name of the NetworkClass resource
func (n *NetworkClass) GetName() string {
	return n.ObjectMeta.Name
}

func init() {
	SchemeBuilder.Register(&NetworkClass{}, &NetworkClassList{})
}
