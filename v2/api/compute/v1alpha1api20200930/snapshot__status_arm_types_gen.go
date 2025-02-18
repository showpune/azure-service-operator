// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

//Deprecated version of Snapshot_Status. Use v1beta20200930.Snapshot_Status instead
type Snapshot_StatusARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM   `json:"extendedLocation,omitempty"`
	Id               *string                       `json:"id,omitempty"`
	Location         *string                       `json:"location,omitempty"`
	ManagedBy        *string                       `json:"managedBy,omitempty"`
	Name             *string                       `json:"name,omitempty"`
	Properties       *SnapshotProperties_StatusARM `json:"properties,omitempty"`
	Sku              *SnapshotSku_StatusARM        `json:"sku,omitempty"`
	Tags             map[string]string             `json:"tags,omitempty"`
	Type             *string                       `json:"type,omitempty"`
}

//Deprecated version of SnapshotProperties_Status. Use v1beta20200930.SnapshotProperties_Status instead
type SnapshotProperties_StatusARM struct {
	CreationData                 *CreationData_StatusARM                   `json:"creationData,omitempty"`
	DiskAccessId                 *string                                   `json:"diskAccessId,omitempty"`
	DiskSizeBytes                *int                                      `json:"diskSizeBytes,omitempty"`
	DiskSizeGB                   *int                                      `json:"diskSizeGB,omitempty"`
	DiskState                    *DiskState_Status                         `json:"diskState,omitempty"`
	Encryption                   *Encryption_StatusARM                     `json:"encryption,omitempty"`
	EncryptionSettingsCollection *EncryptionSettingsCollection_StatusARM   `json:"encryptionSettingsCollection,omitempty"`
	HyperVGeneration             *SnapshotPropertiesStatusHyperVGeneration `json:"hyperVGeneration,omitempty"`
	Incremental                  *bool                                     `json:"incremental,omitempty"`
	NetworkAccessPolicy          *NetworkAccessPolicy_Status               `json:"networkAccessPolicy,omitempty"`
	OsType                       *SnapshotPropertiesStatusOsType           `json:"osType,omitempty"`
	ProvisioningState            *string                                   `json:"provisioningState,omitempty"`
	PurchasePlan                 *PurchasePlan_StatusARM                   `json:"purchasePlan,omitempty"`
	TimeCreated                  *string                                   `json:"timeCreated,omitempty"`
	UniqueId                     *string                                   `json:"uniqueId,omitempty"`
}

//Deprecated version of SnapshotSku_Status. Use v1beta20200930.SnapshotSku_Status instead
type SnapshotSku_StatusARM struct {
	Name *SnapshotSkuStatusName `json:"name,omitempty"`
	Tier *string                `json:"tier,omitempty"`
}

//Deprecated version of SnapshotSkuStatusName. Use v1beta20200930.SnapshotSkuStatusName instead
type SnapshotSkuStatusName string

const (
	SnapshotSkuStatusNamePremiumLRS  = SnapshotSkuStatusName("Premium_LRS")
	SnapshotSkuStatusNameStandardLRS = SnapshotSkuStatusName("Standard_LRS")
	SnapshotSkuStatusNameStandardZRS = SnapshotSkuStatusName("Standard_ZRS")
)
