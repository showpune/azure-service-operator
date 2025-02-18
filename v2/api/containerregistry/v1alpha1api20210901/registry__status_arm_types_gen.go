// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210901

//Deprecated version of Registry_Status. Use v1beta20210901.Registry_Status instead
type Registry_StatusARM struct {
	Id         *string                       `json:"id,omitempty"`
	Identity   *IdentityProperties_StatusARM `json:"identity,omitempty"`
	Location   *string                       `json:"location,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *RegistryProperties_StatusARM `json:"properties,omitempty"`
	Sku        *Sku_StatusARM                `json:"sku,omitempty"`
	SystemData *SystemData_StatusARM         `json:"systemData,omitempty"`
	Tags       map[string]string             `json:"tags,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}

//Deprecated version of IdentityProperties_Status. Use v1beta20210901.IdentityProperties_Status instead
type IdentityProperties_StatusARM struct {
	PrincipalId            *string                                     `json:"principalId,omitempty"`
	TenantId               *string                                     `json:"tenantId,omitempty"`
	Type                   *IdentityPropertiesStatusType               `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentityProperties_StatusARM `json:"userAssignedIdentities,omitempty"`
}

//Deprecated version of RegistryProperties_Status. Use v1beta20210901.RegistryProperties_Status instead
type RegistryProperties_StatusARM struct {
	AdminUserEnabled           *bool                                                     `json:"adminUserEnabled,omitempty"`
	CreationDate               *string                                                   `json:"creationDate,omitempty"`
	DataEndpointEnabled        *bool                                                     `json:"dataEndpointEnabled,omitempty"`
	DataEndpointHostNames      []string                                                  `json:"dataEndpointHostNames,omitempty"`
	Encryption                 *EncryptionProperty_StatusARM                             `json:"encryption,omitempty"`
	LoginServer                *string                                                   `json:"loginServer,omitempty"`
	NetworkRuleBypassOptions   *RegistryPropertiesStatusNetworkRuleBypassOptions         `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet             *NetworkRuleSet_StatusARM                                 `json:"networkRuleSet,omitempty"`
	Policies                   *Policies_StatusARM                                       `json:"policies,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *RegistryPropertiesStatusProvisioningState                `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *RegistryPropertiesStatusPublicNetworkAccess              `json:"publicNetworkAccess,omitempty"`
	Status                     *Status_StatusARM                                         `json:"status,omitempty"`
	ZoneRedundancy             *RegistryPropertiesStatusZoneRedundancy                   `json:"zoneRedundancy,omitempty"`
}

//Deprecated version of Sku_Status. Use v1beta20210901.Sku_Status instead
type Sku_StatusARM struct {
	Name *SkuStatusName `json:"name,omitempty"`
	Tier *SkuStatusTier `json:"tier,omitempty"`
}

//Deprecated version of SystemData_Status. Use v1beta20210901.SystemData_Status instead
type SystemData_StatusARM struct {
	CreatedAt          *string                             `json:"createdAt,omitempty"`
	CreatedBy          *string                             `json:"createdBy,omitempty"`
	CreatedByType      *SystemDataStatusCreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                             `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                             `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

//Deprecated version of EncryptionProperty_Status. Use v1beta20210901.EncryptionProperty_Status instead
type EncryptionProperty_StatusARM struct {
	KeyVaultProperties *KeyVaultProperties_StatusARM   `json:"keyVaultProperties,omitempty"`
	Status             *EncryptionPropertyStatusStatus `json:"status,omitempty"`
}

//Deprecated version of IdentityPropertiesStatusType. Use v1beta20210901.IdentityPropertiesStatusType instead
type IdentityPropertiesStatusType string

const (
	IdentityPropertiesStatusTypeNone                       = IdentityPropertiesStatusType("None")
	IdentityPropertiesStatusTypeSystemAssigned             = IdentityPropertiesStatusType("SystemAssigned")
	IdentityPropertiesStatusTypeSystemAssignedUserAssigned = IdentityPropertiesStatusType("SystemAssigned, UserAssigned")
	IdentityPropertiesStatusTypeUserAssigned               = IdentityPropertiesStatusType("UserAssigned")
)

//Deprecated version of NetworkRuleSet_Status. Use v1beta20210901.NetworkRuleSet_Status instead
type NetworkRuleSet_StatusARM struct {
	DefaultAction *NetworkRuleSetStatusDefaultAction `json:"defaultAction,omitempty"`
	IpRules       []IPRule_StatusARM                 `json:"ipRules,omitempty"`
}

//Deprecated version of Policies_Status. Use v1beta20210901.Policies_Status instead
type Policies_StatusARM struct {
	ExportPolicy     *ExportPolicy_StatusARM     `json:"exportPolicy,omitempty"`
	QuarantinePolicy *QuarantinePolicy_StatusARM `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicy_StatusARM  `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicy_StatusARM      `json:"trustPolicy,omitempty"`
}

//Deprecated version of PrivateEndpointConnection_Status_SubResourceEmbedded. Use v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbedded instead
type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

//Deprecated version of SkuStatusName. Use v1beta20210901.SkuStatusName instead
type SkuStatusName string

const (
	SkuStatusNameBasic    = SkuStatusName("Basic")
	SkuStatusNameClassic  = SkuStatusName("Classic")
	SkuStatusNamePremium  = SkuStatusName("Premium")
	SkuStatusNameStandard = SkuStatusName("Standard")
)

//Deprecated version of SkuStatusTier. Use v1beta20210901.SkuStatusTier instead
type SkuStatusTier string

const (
	SkuStatusTierBasic    = SkuStatusTier("Basic")
	SkuStatusTierClassic  = SkuStatusTier("Classic")
	SkuStatusTierPremium  = SkuStatusTier("Premium")
	SkuStatusTierStandard = SkuStatusTier("Standard")
)

//Deprecated version of Status_Status. Use v1beta20210901.Status_Status instead
type Status_StatusARM struct {
	DisplayStatus *string `json:"displayStatus,omitempty"`
	Message       *string `json:"message,omitempty"`
	Timestamp     *string `json:"timestamp,omitempty"`
}

//Deprecated version of SystemDataStatusCreatedByType. Use v1beta20210901.SystemDataStatusCreatedByType instead
type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

//Deprecated version of SystemDataStatusLastModifiedByType. Use v1beta20210901.SystemDataStatusLastModifiedByType instead
type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)

//Deprecated version of UserIdentityProperties_Status. Use v1beta20210901.UserIdentityProperties_Status instead
type UserIdentityProperties_StatusARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

//Deprecated version of ExportPolicy_Status. Use v1beta20210901.ExportPolicy_Status instead
type ExportPolicy_StatusARM struct {
	Status *ExportPolicyStatusStatus `json:"status,omitempty"`
}

//Deprecated version of IPRule_Status. Use v1beta20210901.IPRule_Status instead
type IPRule_StatusARM struct {
	Action *IPRuleStatusAction `json:"action,omitempty"`
	Value  *string             `json:"value,omitempty"`
}

//Deprecated version of KeyVaultProperties_Status. Use v1beta20210901.KeyVaultProperties_Status instead
type KeyVaultProperties_StatusARM struct {
	Identity                 *string `json:"identity,omitempty"`
	KeyIdentifier            *string `json:"keyIdentifier,omitempty"`
	KeyRotationEnabled       *bool   `json:"keyRotationEnabled,omitempty"`
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`
	VersionedKeyIdentifier   *string `json:"versionedKeyIdentifier,omitempty"`
}

//Deprecated version of QuarantinePolicy_Status. Use v1beta20210901.QuarantinePolicy_Status instead
type QuarantinePolicy_StatusARM struct {
	Status *QuarantinePolicyStatusStatus `json:"status,omitempty"`
}

//Deprecated version of RetentionPolicy_Status. Use v1beta20210901.RetentionPolicy_Status instead
type RetentionPolicy_StatusARM struct {
	Days            *int                         `json:"days,omitempty"`
	LastUpdatedTime *string                      `json:"lastUpdatedTime,omitempty"`
	Status          *RetentionPolicyStatusStatus `json:"status,omitempty"`
}

//Deprecated version of TrustPolicy_Status. Use v1beta20210901.TrustPolicy_Status instead
type TrustPolicy_StatusARM struct {
	Status *TrustPolicyStatusStatus `json:"status,omitempty"`
	Type   *TrustPolicyStatusType   `json:"type,omitempty"`
}
