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

type AccessControlTranslationObservation struct {
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type AccessControlTranslationParameters struct {
}

type ApplyServerSideEncryptionByDefaultObservation struct {

	// The AWS KMS master key ID used for the SSE-KMS encryption. This can only be used when you set the value of sse_algorithm as aws:kms. The default aws/s3 AWS KMS master key is used if this element is absent while the sse_algorithm is aws:kms.
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// The server-side encryption algorithm to use. Valid values are AES256 and aws:kms
	SseAlgorithm *string `json:"sseAlgorithm,omitempty" tf:"sse_algorithm,omitempty"`
}

type ApplyServerSideEncryptionByDefaultParameters struct {
}

type BucketObservation struct {

	// The canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, and log-delivery-write. Defaults to private.  Conflicts with grant. Use the resource aws_s3_bucket_acl instead.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Sets the accelerate configuration of an existing bucket. Can be Enabled or Suspended. Cannot be used in cn-north-1 or us-gov-west-1.
	// Use the resource aws_s3_bucket_accelerate_configuration instead.
	AccelerationStatus *string `json:"accelerationStatus,omitempty" tf:"acceleration_status,omitempty"`

	// The ARN of the bucket. Will be of format arn:aws:s3:::bucketname.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The bucket domain name. Will be of format bucketname.s3.amazonaws.com.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// The bucket region-specific domain name. The bucket domain name including the region name, please refer here for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent redirect issues from CloudFront to S3 Origin URL.
	BucketRegionalDomainName *string `json:"bucketRegionalDomainName,omitempty" tf:"bucket_regional_domain_name,omitempty"`

	// A rule of Cross-Origin Resource Sharing. See CORS rule below for details. Use the resource aws_s3_bucket_cors_configuration instead.
	CorsRule []CorsRuleObservation `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// An ACL policy grant. See Grant below for details. Conflicts with acl. Use the resource aws_s3_bucket_acl instead.
	Grant []GrantObservation `json:"grant,omitempty" tf:"grant,omitempty"`

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A configuration of object lifecycle management. See Lifecycle Rule below for details.
	// Use the resource aws_s3_bucket_lifecycle_configuration instead.
	LifecycleRule []LifecycleRuleObservation `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A configuration of S3 bucket logging parameters. See Logging below for details.
	// Use the resource aws_s3_bucket_logging instead.
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// A configuration of S3 object locking. See Object Lock Configuration below for details.
	// Use the object_lock_enabled parameter and the resource aws_s3_bucket_object_lock_configuration instead.
	ObjectLockConfiguration []ObjectLockConfigurationObservation `json:"objectLockConfiguration,omitempty" tf:"object_lock_configuration,omitempty"`

	// A valid bucket policy JSON document. In this case, please make sure you use the verbose/specific version of the policy.
	// Use the resource aws_s3_bucket_policy instead.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// A configuration of replication configuration. See Replication Configuration below for details.
	// Use the resource aws_s3_bucket_replication_configuration instead.
	ReplicationConfiguration []ReplicationConfigurationObservation `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`

	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either BucketOwner or Requester. By default, the owner of the S3 bucket would incur the costs of any data transfer.
	// See Requester Pays Buckets developer guide for more information.
	// Use the resource aws_s3_bucket_request_payment_configuration instead.
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// A configuration of server-side encryption configuration. See Server Side Encryption Configuration below for details.
	// Use the resource aws_s3_bucket_server_side_encryption_configuration instead.
	ServerSideEncryptionConfiguration []ServerSideEncryptionConfigurationObservation `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A configuration of the S3 bucket versioning state. See Versioning below for details. Use the resource aws_s3_bucket_versioning instead.
	Versioning []VersioningObservation `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A configuration of the S3 bucket website. See Website below for details.
	// Use the resource aws_s3_bucket_website_configuration instead.
	Website []WebsiteObservation `json:"website,omitempty" tf:"website,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketParameters struct {

	// A boolean that indicates all objects (including any locked objects) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The Route 53 Hosted Zone ID for this bucket's region.
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Valid values are true or false.
	// +kubebuilder:validation:Optional
	ObjectLockEnabled *bool `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// The AWS region this bucket resides in.
	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the bucket. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CorsRuleObservation struct {

	// List of headers allowed.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// One or more HTTP methods that you allow the origin to execute. Can be GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleParameters struct {
}

type DefaultRetentionObservation struct {

	// The number of days that you want to specify for the default retention period.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The default Object Lock retention mode you want to apply to new objects placed in this bucket. Valid values are GOVERNANCE and COMPLIANCE.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The number of years that you want to specify for the default retention period.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type DefaultRetentionParameters struct {
}

type DestinationObservation struct {

	// Specifies the overrides to use for object owners on replication. Must be used in conjunction with account_id owner override configuration.
	AccessControlTranslation []AccessControlTranslationObservation `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// The Account ID to use for overriding the object owner on replication. Must be used in conjunction with access_control_translation override configuration.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Enables replication metrics  (documented below).
	Metrics []MetricsObservation `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Destination KMS encryption key ARN for SSE-KMS replication. Must be used in conjunction with
	// sse_kms_encrypted_objects source selection criteria.
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id,omitempty"`

	// Enables S3 Replication Time Control (S3 RTC) (documented below).
	ReplicationTime []ReplicationTimeObservation `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// Specifies the Amazon S3 storage class to which you want the object to transition.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type DestinationParameters struct {
}

type ExpirationObservation struct {

	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationParameters struct {
}

type FilterObservation struct {

	// Object keyname prefix that identifies subset of objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A map of tags that identifies subset of objects to which the rule applies.
	// The rule applies only to objects having all the tags in its tagset.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilterParameters struct {
}

type GrantObservation struct {

	// Canonical user id to grant for. Used only when type is CanonicalUser.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of permissions to apply for grantee. Valid values are READ, WRITE, READ_ACP, WRITE_ACP, FULL_CONTROL.
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Type of grantee to apply for. Valid values are CanonicalUser and Group. AmazonCustomerByEmail is not supported.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Uri address to grant for. Used only when type is Group.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GrantParameters struct {
}

type LifecycleRuleObservation struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Specifies lifecycle rule status.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire. See Expiration below for details.
	Expiration []ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when noncurrent object versions expire. See Noncurrent Version Expiration below for details.
	NoncurrentVersionExpiration []NoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Specifies when noncurrent object versions transitions. See Noncurrent Version Transition below for details.
	NoncurrentVersionTransition []NoncurrentVersionTransitionObservation `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies object tags key and value.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a period in the object's transitions. See Transition below for details.
	Transition []TransitionObservation `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LifecycleRuleParameters struct {
}

type LoggingObservation struct {

	// The name of the bucket that will receive the log objects.
	TargetBucket *string `json:"targetBucket,omitempty" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type LoggingParameters struct {
}

type MetricsObservation struct {

	// Threshold within which objects are to be replicated. The only valid value is 15.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// The status of the rule. Either Enabled or Disabled. The rule is ignored if status is not Enabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type MetricsParameters struct {
}

type NoncurrentVersionExpirationObservation struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationParameters struct {
}

type NoncurrentVersionTransitionObservation struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Amazon S3 storage class to which you want the object to transition.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type NoncurrentVersionTransitionParameters struct {
}

type ObjectLockConfigurationObservation struct {

	// Indicates whether this bucket has an Object Lock configuration enabled. Valid value is Enabled. Use the top-level argument object_lock_enabled instead.
	ObjectLockEnabled *string `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// The Object Lock rule in place for this bucket (documented below).
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ObjectLockConfigurationParameters struct {
}

type ReplicationConfigurationObservation struct {

	// The ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Specifies the rules managing the replication (documented below).
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`
}

type ReplicationConfigurationParameters struct {
}

type ReplicationTimeObservation struct {

	// Threshold within which objects are to be replicated. The only valid value is 15.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// The status of the rule. Either Enabled or Disabled. The rule is ignored if status is not Enabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReplicationTimeParameters struct {
}

type RuleObservation struct {

	// The default retention period that you want to apply to new objects placed in this bucket (documented below).
	DefaultRetention []DefaultRetentionObservation `json:"defaultRetention,omitempty" tf:"default_retention,omitempty"`
}

type RuleParameters struct {
}

type RulesObservation struct {

	// Whether delete markers are replicated. The only valid value is Enabled. To disable, omit this argument. This argument is only valid with V2 replication configurations (i.e., when filter is used).
	DeleteMarkerReplicationStatus *string `json:"deleteMarkerReplicationStatus,omitempty" tf:"delete_marker_replication_status,omitempty"`

	// Specifies the destination for the rule (documented below).
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies (documented below).
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object keyname prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The priority associated with the rule. Priority should only be set if filter is configured. If not provided, defaults to 0. Priority must be unique between multiple rules.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies special object selection criteria (documented below).
	SourceSelectionCriteria []SourceSelectionCriteriaObservation `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// The status of the rule. Either Enabled or Disabled. The rule is ignored if status is not Enabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RulesParameters struct {
}

type ServerSideEncryptionConfigurationObservation struct {

	// A single object for server-side encryption by default configuration. (documented below)
	Rule []ServerSideEncryptionConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ServerSideEncryptionConfigurationParameters struct {
}

type ServerSideEncryptionConfigurationRuleObservation struct {

	// A single object for setting server-side encryption by default. (documented below)
	ApplyServerSideEncryptionByDefault []ApplyServerSideEncryptionByDefaultObservation `json:"applyServerSideEncryptionByDefault,omitempty" tf:"apply_server_side_encryption_by_default,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`
}

type ServerSideEncryptionConfigurationRuleParameters struct {
}

type SourceSelectionCriteriaObservation struct {

	// Match SSE-KMS encrypted objects (documented below). If specified, replica_kms_key_id
	// in destination must be specified as well.
	SseKMSEncryptedObjects []SseKMSEncryptedObjectsObservation `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type SourceSelectionCriteriaParameters struct {
}

type SseKMSEncryptedObjectsObservation struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SseKMSEncryptedObjectsParameters struct {
}

type TransitionObservation struct {

	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Amazon S3 storage class to which you want the object to transition.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type TransitionParameters struct {
}

type VersioningObservation struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable MFA delete for either Change the versioning state of your bucket or Permanently delete an object version. Default is false. This cannot be used to toggle this setting but is available to allow managed buckets to reflect the state in AWS
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type VersioningParameters struct {
}

type WebsiteObservation struct {

	// An absolute path to the document to return in case of a 4XX error.
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied.
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type WebsiteParameters struct {
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. Provides a S3 bucket resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}