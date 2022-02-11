// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	dbformysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501"
	"github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServersFirewallRuleExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServersFirewallRuleExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&dbformysql.FlexibleServersFirewallRule{},
		&v1alpha1api20210501storage.FlexibleServersFirewallRule{}}
}
