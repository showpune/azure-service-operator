// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601storage

import (
	"fmt"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200601.DomainsTopic
//Deprecated version of DomainsTopic. Use v1beta20200601.DomainsTopic instead
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainsTopics_Spec `json:"spec,omitempty"`
	Status            DomainTopic_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (topic *DomainsTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &DomainsTopic{}

// ConvertFrom populates our DomainsTopic from the provided hub DomainsTopic
func (topic *DomainsTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20200601s.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1beta20200601storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignPropertiesFromDomainsTopic(source)
}

// ConvertTo populates the provided hub DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20200601s.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected eventgrid/v1beta20200601storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignPropertiesToDomainsTopic(destination)
}

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (topic *DomainsTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceKind returns the kind of the resource
func (topic *DomainsTopic) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (topic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DomainTopic_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (topic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DomainTopic_Status); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st DomainTopic_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// AssignPropertiesFromDomainsTopic populates our DomainsTopic from the provided source DomainsTopic
func (topic *DomainsTopic) AssignPropertiesFromDomainsTopic(source *v20200601s.DomainsTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DomainsTopics_Spec
	err := spec.AssignPropertiesFromDomainsTopicsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDomainsTopicsSpec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status DomainTopic_Status
	err = status.AssignPropertiesFromDomainTopicStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDomainTopicStatus() to populate field Status")
	}
	topic.Status = status

	// No error
	return nil
}

// AssignPropertiesToDomainsTopic populates the provided destination DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) AssignPropertiesToDomainsTopic(destination *v20200601s.DomainsTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec v20200601s.DomainsTopics_Spec
	err := topic.Spec.AssignPropertiesToDomainsTopicsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDomainsTopicsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20200601s.DomainTopic_Status
	err = topic.Status.AssignPropertiesToDomainTopicStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDomainTopicStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200601.DomainsTopic
//Deprecated version of DomainsTopic. Use v1beta20200601.DomainsTopic instead
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

//Storage version of v1alpha1api20200601.DomainTopic_Status
//Deprecated version of DomainTopic_Status. Use v1beta20200601.DomainTopic_Status instead
type DomainTopic_Status struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SystemData        *SystemData_Status     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DomainTopic_Status{}

// ConvertStatusFrom populates our DomainTopic_Status from the provided source
func (topic *DomainTopic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20200601s.DomainTopic_Status)
	if ok {
		// Populate our instance from source
		return topic.AssignPropertiesFromDomainTopicStatus(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.DomainTopic_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignPropertiesFromDomainTopicStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DomainTopic_Status
func (topic *DomainTopic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20200601s.DomainTopic_Status)
	if ok {
		// Populate destination from our instance
		return topic.AssignPropertiesToDomainTopicStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.DomainTopic_Status{}
	err := topic.AssignPropertiesToDomainTopicStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromDomainTopicStatus populates our DomainTopic_Status from the provided source DomainTopic_Status
func (topic *DomainTopic_Status) AssignPropertiesFromDomainTopicStatus(source *v20200601s.DomainTopic_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	topic.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		topic.PropertyBag = propertyBag
	} else {
		topic.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDomainTopicStatus populates the provided destination DomainTopic_Status from our DomainTopic_Status
func (topic *DomainTopic_Status) AssignPropertiesToDomainTopicStatus(destination *v20200601s.DomainTopic_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topic.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(topic.ProvisioningState)

	// SystemData
	if topic.SystemData != nil {
		var systemDatum v20200601s.SystemData_Status
		err := topic.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20200601.DomainsTopics_Spec
type DomainsTopics_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a eventgrid.azure.com/Domain resource
	Owner       *genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner,omitempty" kind:"Domain"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DomainsTopics_Spec{}

// ConvertSpecFrom populates our DomainsTopics_Spec from the provided source
func (topics *DomainsTopics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20200601s.DomainsTopics_Spec)
	if ok {
		// Populate our instance from source
		return topics.AssignPropertiesFromDomainsTopicsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20200601s.DomainsTopics_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topics.AssignPropertiesFromDomainsTopicsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DomainsTopics_Spec
func (topics *DomainsTopics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20200601s.DomainsTopics_Spec)
	if ok {
		// Populate destination from our instance
		return topics.AssignPropertiesToDomainsTopicsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20200601s.DomainsTopics_Spec{}
	err := topics.AssignPropertiesToDomainsTopicsSpec(dst)
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

// AssignPropertiesFromDomainsTopicsSpec populates our DomainsTopics_Spec from the provided source DomainsTopics_Spec
func (topics *DomainsTopics_Spec) AssignPropertiesFromDomainsTopicsSpec(source *v20200601s.DomainsTopics_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	topics.AzureName = source.AzureName

	// Location
	topics.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	topics.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		topics.Owner = &owner
	} else {
		topics.Owner = nil
	}

	// Tags
	topics.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		topics.PropertyBag = propertyBag
	} else {
		topics.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDomainsTopicsSpec populates the provided destination DomainsTopics_Spec from our DomainsTopics_Spec
func (topics *DomainsTopics_Spec) AssignPropertiesToDomainsTopicsSpec(destination *v20200601s.DomainsTopics_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(topics.PropertyBag)

	// AzureName
	destination.AzureName = topics.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(topics.Location)

	// OriginalVersion
	destination.OriginalVersion = topics.OriginalVersion

	// Owner
	if topics.Owner != nil {
		owner := topics.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(topics.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
