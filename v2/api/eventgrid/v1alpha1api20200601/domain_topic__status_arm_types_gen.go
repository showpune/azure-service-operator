// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

//Deprecated version of DomainTopic_Status. Use v1beta20200601.DomainTopic_Status instead
type DomainTopic_StatusARM struct {
	Id         *string                          `json:"id,omitempty"`
	Name       *string                          `json:"name,omitempty"`
	Properties *DomainTopicProperties_StatusARM `json:"properties,omitempty"`
	SystemData *SystemData_StatusARM            `json:"systemData,omitempty"`
	Type       *string                          `json:"type,omitempty"`
}

//Deprecated version of DomainTopicProperties_Status. Use v1beta20200601.DomainTopicProperties_Status instead
type DomainTopicProperties_StatusARM struct {
	ProvisioningState *DomainTopicPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`
}

//Deprecated version of SystemData_Status. Use v1beta20200601.SystemData_Status instead
type SystemData_StatusARM struct {
	CreatedAt          *string                             `json:"createdAt,omitempty"`
	CreatedBy          *string                             `json:"createdBy,omitempty"`
	CreatedByType      *SystemDataStatusCreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                             `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                             `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

//Deprecated version of DomainTopicPropertiesStatusProvisioningState. Use
//v1beta20200601.DomainTopicPropertiesStatusProvisioningState instead
type DomainTopicPropertiesStatusProvisioningState string

const (
	DomainTopicPropertiesStatusProvisioningStateCanceled  = DomainTopicPropertiesStatusProvisioningState("Canceled")
	DomainTopicPropertiesStatusProvisioningStateCreating  = DomainTopicPropertiesStatusProvisioningState("Creating")
	DomainTopicPropertiesStatusProvisioningStateDeleting  = DomainTopicPropertiesStatusProvisioningState("Deleting")
	DomainTopicPropertiesStatusProvisioningStateFailed    = DomainTopicPropertiesStatusProvisioningState("Failed")
	DomainTopicPropertiesStatusProvisioningStateSucceeded = DomainTopicPropertiesStatusProvisioningState("Succeeded")
	DomainTopicPropertiesStatusProvisioningStateUpdating  = DomainTopicPropertiesStatusProvisioningState("Updating")
)

//Deprecated version of SystemDataStatusCreatedByType. Use v1beta20200601.SystemDataStatusCreatedByType instead
type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

//Deprecated version of SystemDataStatusLastModifiedByType. Use v1beta20200601.SystemDataStatusLastModifiedByType instead
type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)
