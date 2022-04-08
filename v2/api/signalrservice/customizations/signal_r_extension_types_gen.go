// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20211001 "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1alpha1api20211001"
	alpha20211001s "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1alpha1api20211001storage"
	v20211001 "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1beta20211001"
	v20211001s "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1beta20211001storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SignalRExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *SignalRExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20211001.SignalR{},
		&alpha20211001s.SignalR{},
		&v20211001.SignalR{},
		&v20211001s.SignalR{}}
}
