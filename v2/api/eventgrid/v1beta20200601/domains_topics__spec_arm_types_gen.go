// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DomainsTopics_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the domain topic.
	Name string `json:"name,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DomainsTopics_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topics DomainsTopics_SpecARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (topics DomainsTopics_SpecARM) GetName() string {
	return topics.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topics DomainsTopics_SpecARM) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}
