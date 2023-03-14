/*
Copyright 2023.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeloxilbappSpec defines the desired state of Kubeloxilbapp
type KubeloxilbappSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="ghcr.io/loxilb-io/kube-loxilb:latest"
	// kube-loxilb Container Image URL
	ContainerImage string `json:"containerImage"`

	// +kubebuilder:validation:Required
	// LoxiLB's API server access URL.
	LoxiURL []string `json:"loxiURL"`

	// +kubebuilder:validation:Required
	// Keystone Container Image URL
	ExternalCIDR string `json:"externalCIDR"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// If you want to turn on LoxiLB BGP mode, set true
	SetBGP bool `json:"setBGP"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=0
	// LoxiLB LoadBalancer mode. 0=default, 1=One Arm mode, 2=Full NAT mode
	SetLBMode int `json:"setLBMode"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=IfNotPresent
	// kube-loxilb image download option
	ImagePullPolicy string `json:"imagePullPolicy"`
}

// KubeloxilbappStatus defines the observed state of Kubeloxilbapp
type KubeloxilbappStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Kubeloxilbapp is the Schema for the kubeloxilbapps API
type Kubeloxilbapp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeloxilbappSpec   `json:"spec,omitempty"`
	Status KubeloxilbappStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KubeloxilbappList contains a list of Kubeloxilbapp
type KubeloxilbappList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kubeloxilbapp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kubeloxilbapp{}, &KubeloxilbappList{})
}
