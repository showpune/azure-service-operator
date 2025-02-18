// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of FlexibleServersDatabases_Spec. Use v1beta20210601.FlexibleServersDatabases_Spec instead
type FlexibleServersDatabases_SpecARM struct {
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *DatabasePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersDatabases_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (databases FlexibleServersDatabases_SpecARM) GetAPIVersion() string {
	return "2021-06-01"
}

// GetName returns the Name of the resource
func (databases FlexibleServersDatabases_SpecARM) GetName() string {
	return databases.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (databases FlexibleServersDatabases_SpecARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

//Deprecated version of DatabaseProperties. Use v1beta20210601.DatabaseProperties instead
type DatabasePropertiesARM struct {
	Charset   *string `json:"charset,omitempty"`
	Collation *string `json:"collation,omitempty"`
}
