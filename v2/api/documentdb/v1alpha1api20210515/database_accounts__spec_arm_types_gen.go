// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

//Deprecated version of DatabaseAccounts_Spec. Use v1beta20210515.DatabaseAccounts_Spec instead
type DatabaseAccounts_SpecARM struct {
	Identity   *ManagedServiceIdentityARM                `json:"identity,omitempty"`
	Kind       *DatabaseAccountsSpecKind                 `json:"kind,omitempty"`
	Location   *string                                   `json:"location,omitempty"`
	Name       string                                    `json:"name,omitempty"`
	Properties *DatabaseAccountCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (accounts DatabaseAccounts_SpecARM) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (accounts DatabaseAccounts_SpecARM) GetName() string {
	return accounts.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts"
func (accounts DatabaseAccounts_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts"
}

//Deprecated version of DatabaseAccountCreateUpdateProperties. Use v1beta20210515.DatabaseAccountCreateUpdateProperties instead
type DatabaseAccountCreateUpdatePropertiesARM struct {
	AnalyticalStorageConfiguration     *AnalyticalStorageConfigurationARM                             `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                      *ApiPropertiesARM                                              `json:"apiProperties,omitempty"`
	BackupPolicy                       *BackupPolicyARM                                               `json:"backupPolicy,omitempty"`
	Capabilities                       []CapabilityARM                                                `json:"capabilities,omitempty"`
	ConnectorOffer                     *DatabaseAccountCreateUpdatePropertiesConnectorOffer           `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicyARM                                          `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicyARM                                                `json:"cors,omitempty"`
	DatabaseAccountOfferType           *DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                                                        `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                                                          `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                                                          `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                                                          `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                                                          `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                                                          `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                                                          `json:"enableMultipleWriteLocations,omitempty"`
	IpRules                            []IpAddressOrRangeARM                                          `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                                                          `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                                                        `json:"keyVaultKeyUri,omitempty"`
	Locations                          []LocationARM                                                  `json:"locations,omitempty"`
	NetworkAclBypass                   *DatabaseAccountCreateUpdatePropertiesNetworkAclBypass         `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                                                       `json:"networkAclBypassResourceIds,omitempty"`
	PublicNetworkAccess                *DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess      `json:"publicNetworkAccess,omitempty"`
	VirtualNetworkRules                []VirtualNetworkRuleARM                                        `json:"virtualNetworkRules,omitempty"`
}

//Deprecated version of DatabaseAccountsSpecKind. Use v1beta20210515.DatabaseAccountsSpecKind instead
// +kubebuilder:validation:Enum={"GlobalDocumentDB","MongoDB","Parse"}
type DatabaseAccountsSpecKind string

const (
	DatabaseAccountsSpecKindGlobalDocumentDB = DatabaseAccountsSpecKind("GlobalDocumentDB")
	DatabaseAccountsSpecKindMongoDB          = DatabaseAccountsSpecKind("MongoDB")
	DatabaseAccountsSpecKindParse            = DatabaseAccountsSpecKind("Parse")
)

//Deprecated version of ManagedServiceIdentity. Use v1beta20210515.ManagedServiceIdentity instead
type ManagedServiceIdentityARM struct {
	Type *ManagedServiceIdentityType `json:"type,omitempty"`
}

//Deprecated version of AnalyticalStorageConfiguration. Use v1beta20210515.AnalyticalStorageConfiguration instead
type AnalyticalStorageConfigurationARM struct {
	SchemaType *AnalyticalStorageConfigurationSchemaType `json:"schemaType,omitempty"`
}

//Deprecated version of ApiProperties. Use v1beta20210515.ApiProperties instead
type ApiPropertiesARM struct {
	ServerVersion *ApiPropertiesServerVersion `json:"serverVersion,omitempty"`
}

//Deprecated version of BackupPolicy. Use v1beta20210515.BackupPolicy instead
type BackupPolicyARM struct {
	Continuous *ContinuousModeBackupPolicyARM `json:"continuousModeBackupPolicy,omitempty"`
	Periodic   *PeriodicModeBackupPolicyARM   `json:"periodicModeBackupPolicy,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupPolicyARM represents a discriminated union (JSON OneOf)
func (policy BackupPolicyARM) MarshalJSON() ([]byte, error) {
	if policy.Continuous != nil {
		return json.Marshal(policy.Continuous)
	}
	if policy.Periodic != nil {
		return json.Marshal(policy.Periodic)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupPolicyARM
func (policy *BackupPolicyARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "Continuous" {
		policy.Continuous = &ContinuousModeBackupPolicyARM{}
		return json.Unmarshal(data, policy.Continuous)
	}
	if discriminator == "Periodic" {
		policy.Periodic = &PeriodicModeBackupPolicyARM{}
		return json.Unmarshal(data, policy.Periodic)
	}

	// No error
	return nil
}

//Deprecated version of Capability. Use v1beta20210515.Capability instead
type CapabilityARM struct {
	Name *string `json:"name,omitempty"`
}

//Deprecated version of ConsistencyPolicy. Use v1beta20210515.ConsistencyPolicy instead
type ConsistencyPolicyARM struct {
	DefaultConsistencyLevel *ConsistencyPolicyDefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                                      `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                                      `json:"maxStalenessPrefix,omitempty"`
}

//Deprecated version of CorsPolicy. Use v1beta20210515.CorsPolicy instead
type CorsPolicyARM struct {
	AllowedHeaders  *string `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int    `json:"maxAgeInSeconds,omitempty"`
}

//Deprecated version of IpAddressOrRange. Use v1beta20210515.IpAddressOrRange instead
type IpAddressOrRangeARM struct {
	IpAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

//Deprecated version of Location. Use v1beta20210515.Location instead
type LocationARM struct {
	FailoverPriority *int    `json:"failoverPriority,omitempty"`
	IsZoneRedundant  *bool   `json:"isZoneRedundant,omitempty"`
	LocationName     *string `json:"locationName,omitempty"`
}

//Deprecated version of ManagedServiceIdentityType. Use v1beta20210515.ManagedServiceIdentityType instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned             = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeSystemAssignedUserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
	ManagedServiceIdentityTypeUserAssigned               = ManagedServiceIdentityType("UserAssigned")
)

//Deprecated version of VirtualNetworkRule. Use v1beta20210515.VirtualNetworkRule instead
type VirtualNetworkRuleARM struct {
	Id                               *string `json:"id,omitempty"`
	IgnoreMissingVNetServiceEndpoint *bool   `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

//Deprecated version of ContinuousModeBackupPolicy. Use v1beta20210515.ContinuousModeBackupPolicy instead
type ContinuousModeBackupPolicyARM struct {
	Type ContinuousModeBackupPolicyType `json:"type,omitempty"`
}

//Deprecated version of PeriodicModeBackupPolicy. Use v1beta20210515.PeriodicModeBackupPolicy instead
type PeriodicModeBackupPolicyARM struct {
	PeriodicModeProperties *PeriodicModePropertiesARM   `json:"periodicModeProperties,omitempty"`
	Type                   PeriodicModeBackupPolicyType `json:"type,omitempty"`
}

//Deprecated version of PeriodicModeProperties. Use v1beta20210515.PeriodicModeProperties instead
type PeriodicModePropertiesARM struct {
	BackupIntervalInMinutes        *int `json:"backupIntervalInMinutes,omitempty"`
	BackupRetentionIntervalInHours *int `json:"backupRetentionIntervalInHours,omitempty"`
}
