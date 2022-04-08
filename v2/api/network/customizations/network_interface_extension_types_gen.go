// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	alpha20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	v20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NetworkInterfaceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NetworkInterfaceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20201101.NetworkInterface{},
		&alpha20201101s.NetworkInterface{},
		&v20201101.NetworkInterface{},
		&v20201101s.NetworkInterface{}}
}
