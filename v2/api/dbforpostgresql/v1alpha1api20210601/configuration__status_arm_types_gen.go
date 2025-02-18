// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

//Deprecated version of Configuration_Status. Use v1beta20210601.Configuration_Status instead
type Configuration_StatusARM struct {
	Id         *string                            `json:"id,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Properties *ConfigurationProperties_StatusARM `json:"properties,omitempty"`
	SystemData *SystemData_StatusARM              `json:"systemData,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

//Deprecated version of ConfigurationProperties_Status. Use v1beta20210601.ConfigurationProperties_Status instead
type ConfigurationProperties_StatusARM struct {
	AllowedValues          *string                                `json:"allowedValues,omitempty"`
	DataType               *ConfigurationPropertiesStatusDataType `json:"dataType,omitempty"`
	DefaultValue           *string                                `json:"defaultValue,omitempty"`
	Description            *string                                `json:"description,omitempty"`
	DocumentationLink      *string                                `json:"documentationLink,omitempty"`
	IsConfigPendingRestart *bool                                  `json:"isConfigPendingRestart,omitempty"`
	IsDynamicConfig        *bool                                  `json:"isDynamicConfig,omitempty"`
	IsReadOnly             *bool                                  `json:"isReadOnly,omitempty"`
	Source                 *string                                `json:"source,omitempty"`
	Unit                   *string                                `json:"unit,omitempty"`
	Value                  *string                                `json:"value,omitempty"`
}

//Deprecated version of SystemData_Status. Use v1beta20210601.SystemData_Status instead
type SystemData_StatusARM struct {
	CreatedAt          *string                             `json:"createdAt,omitempty"`
	CreatedBy          *string                             `json:"createdBy,omitempty"`
	CreatedByType      *SystemDataStatusCreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                             `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                             `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

//Deprecated version of ConfigurationPropertiesStatusDataType. Use v1beta20210601.ConfigurationPropertiesStatusDataType
//instead
type ConfigurationPropertiesStatusDataType string

const (
	ConfigurationPropertiesStatusDataTypeBoolean     = ConfigurationPropertiesStatusDataType("Boolean")
	ConfigurationPropertiesStatusDataTypeEnumeration = ConfigurationPropertiesStatusDataType("Enumeration")
	ConfigurationPropertiesStatusDataTypeInteger     = ConfigurationPropertiesStatusDataType("Integer")
	ConfigurationPropertiesStatusDataTypeNumeric     = ConfigurationPropertiesStatusDataType("Numeric")
)

//Deprecated version of SystemDataStatusCreatedByType. Use v1beta20210601.SystemDataStatusCreatedByType instead
type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

//Deprecated version of SystemDataStatusLastModifiedByType. Use v1beta20210601.SystemDataStatusLastModifiedByType instead
type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)
