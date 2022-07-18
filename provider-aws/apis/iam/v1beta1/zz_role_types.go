/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InlinePolicyObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type InlinePolicyParameters struct {
}

type RoleObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InlinePolicy []InlinePolicyObservation `json:"inlinePolicy,omitempty" tf:"inline_policy,omitempty"`

	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type RoleParameters struct {

	// +kubebuilder:validation:Required
	AssumeRolePolicy *string `json:"assumeRolePolicy" tf:"assume_role_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDetachPolicies *bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSessionDuration *float64 `json:"maxSessionDuration,omitempty" tf:"max_session_duration,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Role is the Schema for the Roles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleSpec   `json:"spec"`
	Status            RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}