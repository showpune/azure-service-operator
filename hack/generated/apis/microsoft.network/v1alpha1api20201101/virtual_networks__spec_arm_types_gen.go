// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"

type VirtualNetworks_SpecARM struct {
	//APIVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	APIVersion VirtualNetworksSpecAPIVersion `json:"apiVersion"`

	//ExtendedLocation: The extended location of the virtual network.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Location: Location to deploy resource to
	Location string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name"`

	//Properties: Properties of the virtual network.
	Properties VirtualNetworks_Spec_PropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type.
	Type VirtualNetworksSpecType `json:"type"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworks_SpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (virtualNetworksSpecARM VirtualNetworks_SpecARM) GetAPIVersion() string {
	return string(virtualNetworksSpecARM.APIVersion)
}

// GetName returns the Name of the resource
func (virtualNetworksSpecARM VirtualNetworks_SpecARM) GetName() string {
	return virtualNetworksSpecARM.Name
}

// GetType returns the Type of the resource
func (virtualNetworksSpecARM VirtualNetworks_SpecARM) GetType() string {
	return string(virtualNetworksSpecARM.Type)
}

// +kubebuilder:validation:Enum={"2020-11-01"}
type VirtualNetworksSpecAPIVersion string

const VirtualNetworksSpecAPIVersion20201101 = VirtualNetworksSpecAPIVersion("2020-11-01")

// +kubebuilder:validation:Enum={"Microsoft.Network/virtualNetworks"}
type VirtualNetworksSpecType string

const VirtualNetworksSpecTypeMicrosoftNetworkVirtualNetworks = VirtualNetworksSpecType("Microsoft.Network/virtualNetworks")

type VirtualNetworks_Spec_PropertiesARM struct {
	//AddressSpace: The AddressSpace that contains an array of IP address ranges that
	//can be used by subnets.
	AddressSpace AddressSpaceARM `json:"addressSpace"`

	//BgpCommunities: Bgp Communities sent over ExpressRoute with each route
	//corresponding to a prefix in this VNET.
	BgpCommunities *VirtualNetworkBgpCommunitiesARM `json:"bgpCommunities,omitempty"`

	//DdosProtectionPlan: The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResourceARM `json:"ddosProtectionPlan,omitempty"`

	//DhcpOptions: The dhcpOptions that contains an array of DNS servers available to
	//VMs deployed in the virtual network.
	DhcpOptions *DhcpOptionsARM `json:"dhcpOptions,omitempty"`

	//EnableDdosProtection: Indicates if DDoS protection is enabled for all the
	//protected resources in the virtual network. It requires a DDoS protection plan
	//associated with the resource.
	EnableDdosProtection *bool `json:"enableDdosProtection,omitempty"`

	//EnableVmProtection: Indicates if VM protection is enabled for all the subnets in
	//the virtual network.
	EnableVmProtection *bool `json:"enableVmProtection,omitempty"`

	//IpAllocations: Array of IpAllocation which reference this VNET.
	IpAllocations []SubResourceARM `json:"ipAllocations,omitempty"`

	//Subnets: A list of subnets in a Virtual Network.
	Subnets []VirtualNetworks_Spec_Properties_SubnetsARM `json:"subnets,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions
type DhcpOptionsARM struct {
	//DnsServers: The list of DNS servers IP addresses.
	DnsServers []string `json:"dnsServers"`
}

type VirtualNetworks_Spec_Properties_SubnetsARM struct {
	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name string `json:"name"`

	//Properties: Properties of the subnet.
	Properties *VirtualNetworks_Spec_Properties_Subnets_PropertiesARM `json:"properties,omitempty"`
}

type VirtualNetworks_Spec_Properties_Subnets_PropertiesARM struct {
	//AddressPrefix: The address prefix for the subnet.
	AddressPrefix string `json:"addressPrefix"`

	//AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	//Delegations: An array of references to the delegations on the subnet.
	Delegations []VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM `json:"delegations,omitempty"`

	//IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResourceARM `json:"ipAllocations,omitempty"`

	//NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResourceARM `json:"natGateway,omitempty"`

	//NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *SubResourceARM `json:"networkSecurityGroup,omitempty"`

	//PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on
	//private end point in the subnet.
	PrivateEndpointNetworkPolicies *string `json:"privateEndpointNetworkPolicies,omitempty"`

	//PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on
	//private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *string `json:"privateLinkServiceNetworkPolicies,omitempty"`

	//RouteTable: The reference to the RouteTable resource.
	RouteTable *SubResourceARM `json:"routeTable,omitempty"`

	//ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []SubResourceARM `json:"serviceEndpointPolicies,omitempty"`

	//ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormatARM `json:"serviceEndpoints,omitempty"`
}

type VirtualNetworks_Spec_Properties_Subnets_Properties_DelegationsARM struct {
	//Name: The name of the resource that is unique within a subnet. This name can be
	//used to access the resource.
	Name string `json:"name"`

	//Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormatARM `json:"properties,omitempty"`
}
