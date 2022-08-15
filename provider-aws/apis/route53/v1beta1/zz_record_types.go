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

type AliasObservation struct {
}

type AliasParameters struct {

	// Set to true if you want Route 53 to determine whether to respond to DNS queries using this resource record set by checking the health of the resource record set. Some resources have special requirements, see related part of documentation.
	// +kubebuilder:validation:Required
	EvaluateTargetHealth *bool `json:"evaluateTargetHealth" tf:"evaluate_target_health,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type FailoverRoutingPolicyObservation struct {
}

type FailoverRoutingPolicyParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type GeolocationRoutingPolicyObservation struct {
}

type GeolocationRoutingPolicyParameters struct {

	// A two-letter continent code. See http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetGeoLocation.html for code details. Either continent or country must be specified.
	// +kubebuilder:validation:Optional
	Continent *string `json:"continent,omitempty" tf:"continent,omitempty"`

	// A two-character country code or * to indicate a default resource record set.
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// A subdivision code for a country.
	// +kubebuilder:validation:Optional
	Subdivision *string `json:"subdivision,omitempty" tf:"subdivision,omitempty"`
}

type LatencyRoutingPolicyObservation struct {
}

type LatencyRoutingPolicyParameters struct {

	// An AWS region from which to measure latency. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type RecordObservation struct {

	// FQDN built using the zone domain and name.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RecordParameters struct {

	// An alias block. Conflicts with ttl & records.
	// Alias record documented below.
	// +kubebuilder:validation:Optional
	Alias []AliasParameters `json:"alias,omitempty" tf:"alias,omitempty"`

	// Allow creation of this record in Terraform to overwrite an existing record, if any. This does not affect the ability to update the record in Terraform and does not prevent other resources within Terraform or manual Route 53 changes outside Terraform from overwriting this record. false by default. This configuration is not recommended for most environments.
	// +kubebuilder:validation:Optional
	AllowOverwrite *bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite,omitempty"`

	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	FailoverRoutingPolicy []FailoverRoutingPolicyParameters `json:"failoverRoutingPolicy,omitempty" tf:"failover_routing_policy,omitempty"`

	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	GeolocationRoutingPolicy []GeolocationRoutingPolicyParameters `json:"geolocationRoutingPolicy,omitempty" tf:"geolocation_routing_policy,omitempty"`

	// The health check the record should be associated with.
	// +crossplane:generate:reference:type=HealthCheck
	// +kubebuilder:validation:Optional
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckIDRef *v1.Reference `json:"healthCheckIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HealthCheckIDSelector *v1.Selector `json:"healthCheckIdSelector,omitempty" tf:"-"`

	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	LatencyRoutingPolicy []LatencyRoutingPolicyParameters `json:"latencyRoutingPolicy,omitempty" tf:"latency_routing_policy,omitempty"`

	// Set to true to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	// +kubebuilder:validation:Optional
	MultivalueAnswerRoutingPolicy *bool `json:"multivalueAnswerRoutingPolicy,omitempty" tf:"multivalue_answer_routing_policy,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration string .
	// +kubebuilder:validation:Optional
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// An AWS region from which to measure latency. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency
	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Unique identifier to differentiate records with routing policies from one another. Required if using failover, geolocation, latency, multivalue_answer, or weighted routing policies documented below.
	// +kubebuilder:validation:Optional
	SetIdentifier *string `json:"setIdentifier,omitempty" tf:"set_identifier,omitempty"`

	// The TTL of the record.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	// +kubebuilder:validation:Optional
	WeightedRoutingPolicy []WeightedRoutingPolicyParameters `json:"weightedRoutingPolicy,omitempty" tf:"weighted_routing_policy,omitempty"`

	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type WeightedRoutingPolicyObservation struct {
}

type WeightedRoutingPolicyParameters struct {

	// A numeric value indicating the relative weight of the record. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted.
	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Record is the Schema for the Records API. Provides a Route53 record resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordSpec   `json:"spec"`
	Status            RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
