// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type UserAssignedIdentities_SpecARM struct {
	//Location: The Azure region where the identity lives.
	Location *string `json:"location,omitempty"`

	//Name: The name of the identity resource.
	Name string `json:"name,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &UserAssignedIdentities_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identities UserAssignedIdentities_SpecARM) GetAPIVersion() string {
	return "2018-11-30"
}

// GetName returns the Name of the resource
func (identities UserAssignedIdentities_SpecARM) GetName() string {
	return identities.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identities UserAssignedIdentities_SpecARM) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}
