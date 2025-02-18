// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

//Deprecated version of ManagementPolicy_Status. Use v1beta20210401.ManagementPolicy_Status instead
type ManagementPolicy_StatusARM struct {
	Id         *string                               `json:"id,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	Properties *ManagementPolicyProperties_StatusARM `json:"properties,omitempty"`
	Type       *string                               `json:"type,omitempty"`
}

//Deprecated version of ManagementPolicyProperties_Status. Use v1beta20210401.ManagementPolicyProperties_Status instead
type ManagementPolicyProperties_StatusARM struct {
	LastModifiedTime *string                           `json:"lastModifiedTime,omitempty"`
	Policy           *ManagementPolicySchema_StatusARM `json:"policy,omitempty"`
}

//Deprecated version of ManagementPolicySchema_Status. Use v1beta20210401.ManagementPolicySchema_Status instead
type ManagementPolicySchema_StatusARM struct {
	Rules []ManagementPolicyRule_StatusARM `json:"rules,omitempty"`
}

//Deprecated version of ManagementPolicyRule_Status. Use v1beta20210401.ManagementPolicyRule_Status instead
type ManagementPolicyRule_StatusARM struct {
	Definition *ManagementPolicyDefinition_StatusARM `json:"definition,omitempty"`
	Enabled    *bool                                 `json:"enabled,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	Type       *ManagementPolicyRuleStatusType       `json:"type,omitempty"`
}

//Deprecated version of ManagementPolicyDefinition_Status. Use v1beta20210401.ManagementPolicyDefinition_Status instead
type ManagementPolicyDefinition_StatusARM struct {
	Actions *ManagementPolicyAction_StatusARM `json:"actions,omitempty"`
	Filters *ManagementPolicyFilter_StatusARM `json:"filters,omitempty"`
}

//Deprecated version of ManagementPolicyRuleStatusType. Use v1beta20210401.ManagementPolicyRuleStatusType instead
type ManagementPolicyRuleStatusType string

const ManagementPolicyRuleStatusTypeLifecycle = ManagementPolicyRuleStatusType("Lifecycle")

//Deprecated version of ManagementPolicyAction_Status. Use v1beta20210401.ManagementPolicyAction_Status instead
type ManagementPolicyAction_StatusARM struct {
	BaseBlob *ManagementPolicyBaseBlob_StatusARM `json:"baseBlob,omitempty"`
	Snapshot *ManagementPolicySnapShot_StatusARM `json:"snapshot,omitempty"`
	Version  *ManagementPolicyVersion_StatusARM  `json:"version,omitempty"`
}

//Deprecated version of ManagementPolicyFilter_Status. Use v1beta20210401.ManagementPolicyFilter_Status instead
type ManagementPolicyFilter_StatusARM struct {
	BlobIndexMatch []TagFilter_StatusARM `json:"blobIndexMatch,omitempty"`
	BlobTypes      []string              `json:"blobTypes,omitempty"`
	PrefixMatch    []string              `json:"prefixMatch,omitempty"`
}

//Deprecated version of ManagementPolicyBaseBlob_Status. Use v1beta20210401.ManagementPolicyBaseBlob_Status instead
type ManagementPolicyBaseBlob_StatusARM struct {
	Delete                      *DateAfterModification_StatusARM `json:"delete,omitempty"`
	EnableAutoTierToHotFromCool *bool                            `json:"enableAutoTierToHotFromCool,omitempty"`
	TierToArchive               *DateAfterModification_StatusARM `json:"tierToArchive,omitempty"`
	TierToCool                  *DateAfterModification_StatusARM `json:"tierToCool,omitempty"`
}

//Deprecated version of ManagementPolicySnapShot_Status. Use v1beta20210401.ManagementPolicySnapShot_Status instead
type ManagementPolicySnapShot_StatusARM struct {
	Delete        *DateAfterCreation_StatusARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreation_StatusARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreation_StatusARM `json:"tierToCool,omitempty"`
}

//Deprecated version of ManagementPolicyVersion_Status. Use v1beta20210401.ManagementPolicyVersion_Status instead
type ManagementPolicyVersion_StatusARM struct {
	Delete        *DateAfterCreation_StatusARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreation_StatusARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreation_StatusARM `json:"tierToCool,omitempty"`
}

//Deprecated version of TagFilter_Status. Use v1beta20210401.TagFilter_Status instead
type TagFilter_StatusARM struct {
	Name  *string `json:"name,omitempty"`
	Op    *string `json:"op,omitempty"`
	Value *string `json:"value,omitempty"`
}

//Deprecated version of DateAfterCreation_Status. Use v1beta20210401.DateAfterCreation_Status instead
type DateAfterCreation_StatusARM struct {
	DaysAfterCreationGreaterThan *float64 `json:"daysAfterCreationGreaterThan,omitempty"`
}

//Deprecated version of DateAfterModification_Status. Use v1beta20210401.DateAfterModification_Status instead
type DateAfterModification_StatusARM struct {
	DaysAfterLastAccessTimeGreaterThan *float64 `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`
	DaysAfterModificationGreaterThan   *float64 `json:"daysAfterModificationGreaterThan,omitempty"`
}
