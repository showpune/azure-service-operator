// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	dbformysqlv1alpha1api20210501 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501"
	"github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501storage"
	dbformysqlv1beta20210501 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501"
	"github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServerExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServerExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&dbformysqlv1alpha1api20210501.FlexibleServer{},
		&v1alpha1api20210501storage.FlexibleServer{},
		&dbformysqlv1beta20210501.FlexibleServer{},
		&v1beta20210501storage.FlexibleServer{}}
}
