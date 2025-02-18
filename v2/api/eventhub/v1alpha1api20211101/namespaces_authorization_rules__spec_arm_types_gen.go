// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of NamespacesAuthorizationRules_Spec. Use v1beta20211101.NamespacesAuthorizationRules_Spec instead
type NamespacesAuthorizationRules_SpecARM struct {
	Location   *string                         `json:"location,omitempty"`
	Name       string                          `json:"name,omitempty"`
	Properties *AuthorizationRulePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string               `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesAuthorizationRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rules NamespacesAuthorizationRules_SpecARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rules NamespacesAuthorizationRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/authorizationRules"
func (rules NamespacesAuthorizationRules_SpecARM) GetType() string {
	return "Microsoft.EventHub/namespaces/authorizationRules"
}

//Deprecated version of AuthorizationRuleProperties. Use v1beta20211101.AuthorizationRuleProperties instead
type AuthorizationRulePropertiesARM struct {
	Rights []AuthorizationRulePropertiesRights `json:"rights,omitempty"`
}
