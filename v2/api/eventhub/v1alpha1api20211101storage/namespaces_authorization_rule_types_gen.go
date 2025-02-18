// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101storage

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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
//Storage version of v1alpha1api20211101.NamespacesAuthorizationRule
//Deprecated version of NamespacesAuthorizationRule. Use v1beta20211101.NamespacesAuthorizationRule instead
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesAuthorizationRules_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_Status          `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesAuthorizationRule{}

// ConvertFrom populates our NamespacesAuthorizationRule from the provided hub NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1beta20211101storage.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignPropertiesFromNamespacesAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1beta20211101storage.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignPropertiesToNamespacesAuthorizationRule(destination)
}

var _ genruntime.KubernetesResource = &NamespacesAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceKind returns the kind of the resource
func (rule *NamespacesAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *NamespacesAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/authorizationRules"
func (rule *NamespacesAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NamespacesAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// AssignPropertiesFromNamespacesAuthorizationRule populates our NamespacesAuthorizationRule from the provided source NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignPropertiesFromNamespacesAuthorizationRule(source *v1beta20211101storage.NamespacesAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesAuthorizationRules_Spec
	err := spec.AssignPropertiesFromNamespacesAuthorizationRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesAuthorizationRulesSpec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status AuthorizationRule_Status
	err = status.AssignPropertiesFromAuthorizationRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromAuthorizationRuleStatus() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesAuthorizationRule populates the provided destination NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignPropertiesToNamespacesAuthorizationRule(destination *v1beta20211101storage.NamespacesAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1beta20211101storage.NamespacesAuthorizationRules_Spec
	err := rule.Spec.AssignPropertiesToNamespacesAuthorizationRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesAuthorizationRulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1beta20211101storage.AuthorizationRule_Status
	err = rule.Status.AssignPropertiesToAuthorizationRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToAuthorizationRuleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NamespacesAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20211101.NamespacesAuthorizationRule
//Deprecated version of NamespacesAuthorizationRule. Use v1beta20211101.NamespacesAuthorizationRule instead
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

//Storage version of v1alpha1api20211101.AuthorizationRule_Status
//Deprecated version of AuthorizationRule_Status. Use v1beta20211101.AuthorizationRule_Status instead
type AuthorizationRule_Status struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rights      []string               `json:"rights,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AuthorizationRule_Status{}

// ConvertStatusFrom populates our AuthorizationRule_Status from the provided source
func (rule *AuthorizationRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1beta20211101storage.AuthorizationRule_Status)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromAuthorizationRuleStatus(src)
	}

	// Convert to an intermediate form
	src = &v1beta20211101storage.AuthorizationRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromAuthorizationRuleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our AuthorizationRule_Status
func (rule *AuthorizationRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1beta20211101storage.AuthorizationRule_Status)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToAuthorizationRuleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20211101storage.AuthorizationRule_Status{}
	err := rule.AssignPropertiesToAuthorizationRuleStatus(dst)
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

// AssignPropertiesFromAuthorizationRuleStatus populates our AuthorizationRule_Status from the provided source AuthorizationRule_Status
func (rule *AuthorizationRule_Status) AssignPropertiesFromAuthorizationRuleStatus(source *v1beta20211101storage.AuthorizationRule_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	rule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// Rights
	rule.Rights = genruntime.CloneSliceOfString(source.Rights)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		rule.SystemData = &systemDatum
	} else {
		rule.SystemData = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		rule.PropertyBag = propertyBag
	} else {
		rule.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToAuthorizationRuleStatus populates the provided destination AuthorizationRule_Status from our AuthorizationRule_Status
func (rule *AuthorizationRule_Status) AssignPropertiesToAuthorizationRuleStatus(destination *v1beta20211101storage.AuthorizationRule_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(rule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// Rights
	destination.Rights = genruntime.CloneSliceOfString(rule.Rights)

	// SystemData
	if rule.SystemData != nil {
		var systemDatum v1beta20211101storage.SystemData_Status
		err := rule.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20211101.NamespacesAuthorizationRules_Spec
type NamespacesAuthorizationRules_Spec struct {
	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a eventhub.azure.com/Namespace resource
	Owner       *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Rights      []string                           `json:"rights,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesAuthorizationRules_Spec{}

// ConvertSpecFrom populates our NamespacesAuthorizationRules_Spec from the provided source
func (rules *NamespacesAuthorizationRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1beta20211101storage.NamespacesAuthorizationRules_Spec)
	if ok {
		// Populate our instance from source
		return rules.AssignPropertiesFromNamespacesAuthorizationRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1beta20211101storage.NamespacesAuthorizationRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rules.AssignPropertiesFromNamespacesAuthorizationRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesAuthorizationRules_Spec
func (rules *NamespacesAuthorizationRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1beta20211101storage.NamespacesAuthorizationRules_Spec)
	if ok {
		// Populate destination from our instance
		return rules.AssignPropertiesToNamespacesAuthorizationRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20211101storage.NamespacesAuthorizationRules_Spec{}
	err := rules.AssignPropertiesToNamespacesAuthorizationRulesSpec(dst)
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

// AssignPropertiesFromNamespacesAuthorizationRulesSpec populates our NamespacesAuthorizationRules_Spec from the provided source NamespacesAuthorizationRules_Spec
func (rules *NamespacesAuthorizationRules_Spec) AssignPropertiesFromNamespacesAuthorizationRulesSpec(source *v1beta20211101storage.NamespacesAuthorizationRules_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	rules.AzureName = source.AzureName

	// Location
	rules.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	rules.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rules.Owner = &owner
	} else {
		rules.Owner = nil
	}

	// Rights
	rules.Rights = genruntime.CloneSliceOfString(source.Rights)

	// Tags
	rules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		rules.PropertyBag = propertyBag
	} else {
		rules.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesAuthorizationRulesSpec populates the provided destination NamespacesAuthorizationRules_Spec from our NamespacesAuthorizationRules_Spec
func (rules *NamespacesAuthorizationRules_Spec) AssignPropertiesToNamespacesAuthorizationRulesSpec(destination *v1beta20211101storage.NamespacesAuthorizationRules_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rules.PropertyBag)

	// AzureName
	destination.AzureName = rules.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(rules.Location)

	// OriginalVersion
	destination.OriginalVersion = rules.OriginalVersion

	// Owner
	if rules.Owner != nil {
		owner := rules.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Rights
	destination.Rights = genruntime.CloneSliceOfString(rules.Rights)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(rules.Tags)

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
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
