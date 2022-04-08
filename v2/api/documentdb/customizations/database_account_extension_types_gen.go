// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	v20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DatabaseAccountExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DatabaseAccountExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20210515.DatabaseAccount{},
		&alpha20210515s.DatabaseAccount{},
		&v20210515.DatabaseAccount{},
		&v20210515s.DatabaseAccount{}}
}
