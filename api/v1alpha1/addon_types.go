/*
Copyright 2021.

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

// AddOnSpec defines the desired state of AddOn
type AddOnSpec struct {
	TargetNamespace          string `json:"targetNamespace"`
	CatalogSourceImage       string `json:"catalogSourceImage"`
	OperatorGroupInstallMode string `json:"operatorGroupInstallMode"`
	SubscriptionChannel      string `json:"subscriptionChannel"`
}

// AddOnStatus defines the observed state of AddOn
type AddOnStatus struct {
	NamespaceStatus     string `json:"namespaceStatus"`
	CatalogSourceStatus string `json:"catalogSourceStatus"`
	OperatorGroupStatus string `json:"operatorGroupStatus"`
	SubscriptionStatus  string `json:"subscriptionStatus"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AddOn is the Schema for the addons API
type AddOn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AddOnSpec   `json:"spec,omitempty"`
	Status AddOnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddOnList contains a list of AddOn
type AddOnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddOn `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AddOn{}, &AddOnList{})
}
