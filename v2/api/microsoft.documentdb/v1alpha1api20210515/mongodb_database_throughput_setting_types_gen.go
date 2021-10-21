// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources={mongodbdatabasethroughputsettings/status,mongodbdatabasethroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return mongodbDatabaseThroughputSetting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	mongodbDatabaseThroughputSetting.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-microsoft-documentdb-azure-com-v1alpha1api20210515-mongodbdatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.mongodbdatabasethroughputsettings.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &MongodbDatabaseThroughputSetting{}

// Default applies defaults to the MongodbDatabaseThroughputSetting resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) Default() {
	mongodbDatabaseThroughputSetting.defaultImpl()
	var temp interface{} = mongodbDatabaseThroughputSetting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the MongodbDatabaseThroughputSetting resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) defaultImpl() {}

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetResourceKind returns the kind of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &mongodbDatabaseThroughputSetting.Spec
}

// GetStatus returns the status of this resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &mongodbDatabaseThroughputSetting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(mongodbDatabaseThroughputSetting.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: mongodbDatabaseThroughputSetting.Namespace,
		Name:      mongodbDatabaseThroughputSetting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		mongodbDatabaseThroughputSetting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	mongodbDatabaseThroughputSetting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-microsoft-documentdb-azure-com-v1alpha1api20210515-mongodbdatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.mongodbdatabasethroughputsettings.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &MongodbDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) ValidateCreate() error {
	validations := mongodbDatabaseThroughputSetting.createValidations()
	var temp interface{} = mongodbDatabaseThroughputSetting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) ValidateDelete() error {
	validations := mongodbDatabaseThroughputSetting.deleteValidations()
	var temp interface{} = mongodbDatabaseThroughputSetting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := mongodbDatabaseThroughputSetting.updateValidations()
	var temp interface{} = mongodbDatabaseThroughputSetting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) createValidations() []func() error {
	return []func() error{mongodbDatabaseThroughputSetting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return mongodbDatabaseThroughputSetting.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&mongodbDatabaseThroughputSetting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromMongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) AssignPropertiesFromMongodbDatabaseThroughputSetting(source *v1alpha1api20210515storage.MongodbDatabaseThroughputSetting) error {

	// Spec
	var spec DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec()")
	}
	mongodbDatabaseThroughputSetting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_Status
	err = status.AssignPropertiesFromThroughputSettingsGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromThroughputSettingsGetResultsStatus()")
	}
	mongodbDatabaseThroughputSetting.Status = status

	// No error
	return nil
}

// AssignPropertiesToMongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) AssignPropertiesToMongodbDatabaseThroughputSetting(destination *v1alpha1api20210515storage.MongodbDatabaseThroughputSetting) error {

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
	err := mongodbDatabaseThroughputSetting.Spec.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.ThroughputSettingsGetResults_Status
	err = mongodbDatabaseThroughputSetting.Status.AssignPropertiesToThroughputSettingsGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToThroughputSettingsGetResultsStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (mongodbDatabaseThroughputSetting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: mongodbDatabaseThroughputSetting.Spec.OriginalVersion(),
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion string

const DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion20210515 = DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion("2021-05-15")

type DatabaseAccountsMongodbDatabasesThroughputSettings_Spec struct {
	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"MongodbDatabase"`

	// +kubebuilder:validation:Required
	//Resource: Cosmos DB resource throughput object. Either throughput is required or
	//autoscaleSettings is required, but not both.
	Resource ThroughputSettingsResource `json:"resource"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if databaseAccountsMongodbDatabasesThroughputSettingsSpec == nil {
		return nil, nil
	}
	var result DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM

	// Set property ‘Location’:
	if databaseAccountsMongodbDatabasesThroughputSettingsSpec.Location != nil {
		location := *databaseAccountsMongodbDatabasesThroughputSettingsSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	resourceARM, err := databaseAccountsMongodbDatabasesThroughputSettingsSpec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(ThroughputSettingsResourceARM)

	// Set property ‘Tags’:
	if databaseAccountsMongodbDatabasesThroughputSettingsSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range databaseAccountsMongodbDatabasesThroughputSettingsSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) CreateEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		databaseAccountsMongodbDatabasesThroughputSettingsSpec.Location = &location
	}

	// Set property ‘Owner’:
	databaseAccountsMongodbDatabasesThroughputSettingsSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource ThroughputSettingsResource
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	databaseAccountsMongodbDatabasesThroughputSettingsSpec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		databaseAccountsMongodbDatabasesThroughputSettingsSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			databaseAccountsMongodbDatabasesThroughputSettingsSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from the provided source
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec)
	if ok {
		// Populate our instance from source
		return databaseAccountsMongodbDatabasesThroughputSettingsSpec.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databaseAccountsMongodbDatabasesThroughputSettingsSpec.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec)
	if ok {
		// Populate destination from our instance
		return databaseAccountsMongodbDatabasesThroughputSettingsSpec.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}
	err := databaseAccountsMongodbDatabasesThroughputSettingsSpec.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec populates our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from the provided source DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(source *v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) error {

	// Location
	databaseAccountsMongodbDatabasesThroughputSettingsSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	databaseAccountsMongodbDatabasesThroughputSettingsSpec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignPropertiesFromThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesFromThroughputSettingsResource()")
		}
		databaseAccountsMongodbDatabasesThroughputSettingsSpec.Resource = resource
	} else {
		databaseAccountsMongodbDatabasesThroughputSettingsSpec.Resource = ThroughputSettingsResource{}
	}

	// Tags
	databaseAccountsMongodbDatabasesThroughputSettingsSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec populates the provided destination DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(destination *v1alpha1api20210515storage.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(databaseAccountsMongodbDatabasesThroughputSettingsSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = databaseAccountsMongodbDatabasesThroughputSettingsSpec.OriginalVersion()

	// Owner
	destination.Owner = databaseAccountsMongodbDatabasesThroughputSettingsSpec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.ThroughputSettingsResource
	err := databaseAccountsMongodbDatabasesThroughputSettingsSpec.Resource.AssignPropertiesToThroughputSettingsResource(&resource)
	if err != nil {
		return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesToThroughputSettingsResource()")
	}
	destination.Resource = &resource

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databaseAccountsMongodbDatabasesThroughputSettingsSpec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (databaseAccountsMongodbDatabasesThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
