/*
Copyright 2024.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IdlePodSpec defines the desired state of IdlePod
type IdlePodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of IdlePod
	// Important: Run "make" to regenerate code after modifying this file
	PodTemplate corev1.PodTemplateSpec `json:"template,omitempty"`
	Definition  string                 `json:"definition,omitempty"`
}

// IdlePodStatus defines the observed state of IdlePod
type IdlePodStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of IdlePod
	// Important: Run "make" to regenerate code after modifying this file
	Phase string `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IdlePod is the Schema for the IdlePods API
type IdlePod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdlePodSpec   `json:"spec,omitempty"`
	Status IdlePodStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IdlePodList contains a list of IdlePod
type IdlePodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdlePod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IdlePod{}, &IdlePodList{})
}
