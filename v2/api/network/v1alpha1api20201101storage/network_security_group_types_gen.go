// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"fmt"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.NetworkSecurityGroup
//Deprecated version of NetworkSecurityGroup. Use v1beta20201101.NetworkSecurityGroup instead
type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroups_Spec                                           `json:"spec,omitempty"`
	Status            NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NetworkSecurityGroup{}

// GetConditions returns the conditions of the resource
func (group *NetworkSecurityGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *NetworkSecurityGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ conversion.Convertible = &NetworkSecurityGroup{}

// ConvertFrom populates our NetworkSecurityGroup from the provided hub NetworkSecurityGroup
func (group *NetworkSecurityGroup) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201101s.NetworkSecurityGroup)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/NetworkSecurityGroup but received %T instead", hub)
	}

	return group.AssignPropertiesFromNetworkSecurityGroup(source)
}

// ConvertTo populates the provided hub NetworkSecurityGroup from our NetworkSecurityGroup
func (group *NetworkSecurityGroup) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201101s.NetworkSecurityGroup)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/NetworkSecurityGroup but received %T instead", hub)
	}

	return group.AssignPropertiesToNetworkSecurityGroup(destination)
}

var _ genruntime.KubernetesResource = &NetworkSecurityGroup{}

// AzureName returns the Azure name of the resource
func (group *NetworkSecurityGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (group NetworkSecurityGroup) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (group *NetworkSecurityGroup) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (group *NetworkSecurityGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *NetworkSecurityGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups"
func (group *NetworkSecurityGroup) GetType() string {
	return "Microsoft.Network/networkSecurityGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *NetworkSecurityGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (group *NetworkSecurityGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return &genruntime.ResourceReference{
		Group: ownerGroup,
		Kind:  ownerKind,
		Name:  group.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (group *NetworkSecurityGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// AssignPropertiesFromNetworkSecurityGroup populates our NetworkSecurityGroup from the provided source NetworkSecurityGroup
func (group *NetworkSecurityGroup) AssignPropertiesFromNetworkSecurityGroup(source *v20201101s.NetworkSecurityGroup) error {

	// ObjectMeta
	group.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NetworkSecurityGroups_Spec
	err := spec.AssignPropertiesFromNetworkSecurityGroupsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNetworkSecurityGroupsSpec() to populate field Spec")
	}
	group.Spec = spec

	// Status
	var status NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
	err = status.AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded() to populate field Status")
	}
	group.Status = status

	// No error
	return nil
}

// AssignPropertiesToNetworkSecurityGroup populates the provided destination NetworkSecurityGroup from our NetworkSecurityGroup
func (group *NetworkSecurityGroup) AssignPropertiesToNetworkSecurityGroup(destination *v20201101s.NetworkSecurityGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *group.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201101s.NetworkSecurityGroups_Spec
	err := group.Spec.AssignPropertiesToNetworkSecurityGroupsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNetworkSecurityGroupsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
	err = group.Status.AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *NetworkSecurityGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "NetworkSecurityGroup",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.NetworkSecurityGroup
//Deprecated version of NetworkSecurityGroup. Use v1beta20201101.NetworkSecurityGroup instead
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroup `json:"items"`
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
//Deprecated version of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded instead
type NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	Conditions           []conditions.Condition                                             `json:"conditions,omitempty"`
	DefaultSecurityRules []SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded     `json:"defaultSecurityRules,omitempty"`
	Etag                 *string                                                            `json:"etag,omitempty"`
	FlowLogs             []FlowLog_Status_SubResourceEmbedded                               `json:"flowLogs,omitempty"`
	Id                   *string                                                            `json:"id,omitempty"`
	Location             *string                                                            `json:"location,omitempty"`
	Name                 *string                                                            `json:"name,omitempty"`
	NetworkInterfaces    []NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded `json:"networkInterfaces,omitempty"`
	PropertyBag          genruntime.PropertyBag                                             `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                                                            `json:"provisioningState,omitempty"`
	ResourceGuid         *string                                                            `json:"resourceGuid,omitempty"`
	SecurityRules        []SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded     `json:"securityRules,omitempty"`
	Subnets              []Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded           `json:"subnets,omitempty"`
	Tags                 map[string]string                                                  `json:"tags,omitempty"`
	Type                 *string                                                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}

// ConvertStatusFrom populates our NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded from the provided source
func (embedded *NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded)
	if ok {
		// Populate our instance from source
		return embedded.AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = embedded.AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded)
	if ok {
		// Populate destination from our instance
		return embedded.AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded{}
	err := embedded.AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded populates our NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded from the provided source NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesFromNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(source *v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	embedded.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DefaultSecurityRules
	if source.DefaultSecurityRules != nil {
		defaultSecurityRuleList := make([]SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded, len(source.DefaultSecurityRules))
		for defaultSecurityRuleIndex, defaultSecurityRuleItem := range source.DefaultSecurityRules {
			// Shadow the loop variable to avoid aliasing
			defaultSecurityRuleItem := defaultSecurityRuleItem
			var defaultSecurityRule SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := defaultSecurityRule.AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(&defaultSecurityRuleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded() to populate field DefaultSecurityRules")
			}
			defaultSecurityRuleList[defaultSecurityRuleIndex] = defaultSecurityRule
		}
		embedded.DefaultSecurityRules = defaultSecurityRuleList
	} else {
		embedded.DefaultSecurityRules = nil
	}

	// Etag
	embedded.Etag = genruntime.ClonePointerToString(source.Etag)

	// FlowLogs
	if source.FlowLogs != nil {
		flowLogList := make([]FlowLog_Status_SubResourceEmbedded, len(source.FlowLogs))
		for flowLogIndex, flowLogItem := range source.FlowLogs {
			// Shadow the loop variable to avoid aliasing
			flowLogItem := flowLogItem
			var flowLog FlowLog_Status_SubResourceEmbedded
			err := flowLog.AssignPropertiesFromFlowLogStatusSubResourceEmbedded(&flowLogItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromFlowLogStatusSubResourceEmbedded() to populate field FlowLogs")
			}
			flowLogList[flowLogIndex] = flowLog
		}
		embedded.FlowLogs = flowLogList
	} else {
		embedded.FlowLogs = nil
	}

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	embedded.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	embedded.Name = genruntime.ClonePointerToString(source.Name)

	// NetworkInterfaces
	if source.NetworkInterfaces != nil {
		networkInterfaceList := make([]NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded, len(source.NetworkInterfaces))
		for networkInterfaceIndex, networkInterfaceItem := range source.NetworkInterfaces {
			// Shadow the loop variable to avoid aliasing
			networkInterfaceItem := networkInterfaceItem
			var networkInterface NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := networkInterface.AssignPropertiesFromNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(&networkInterfaceItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded() to populate field NetworkInterfaces")
			}
			networkInterfaceList[networkInterfaceIndex] = networkInterface
		}
		embedded.NetworkInterfaces = networkInterfaceList
	} else {
		embedded.NetworkInterfaces = nil
	}

	// ProvisioningState
	embedded.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ResourceGuid
	embedded.ResourceGuid = genruntime.ClonePointerToString(source.ResourceGuid)

	// SecurityRules
	if source.SecurityRules != nil {
		securityRuleList := make([]SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded, len(source.SecurityRules))
		for securityRuleIndex, securityRuleItem := range source.SecurityRules {
			// Shadow the loop variable to avoid aliasing
			securityRuleItem := securityRuleItem
			var securityRule SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := securityRule.AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(&securityRuleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded() to populate field SecurityRules")
			}
			securityRuleList[securityRuleIndex] = securityRule
		}
		embedded.SecurityRules = securityRuleList
	} else {
		embedded.SecurityRules = nil
	}

	// Subnets
	if source.Subnets != nil {
		subnetList := make([]Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded, len(source.Subnets))
		for subnetIndex, subnetItem := range source.Subnets {
			// Shadow the loop variable to avoid aliasing
			subnetItem := subnetItem
			var subnet Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := subnet.AssignPropertiesFromSubnetStatusNetworkSecurityGroupSubResourceEmbedded(&subnetItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromSubnetStatusNetworkSecurityGroupSubResourceEmbedded() to populate field Subnets")
			}
			subnetList[subnetIndex] = subnet
		}
		embedded.Subnets = subnetList
	} else {
		embedded.Subnets = nil
	}

	// Tags
	embedded.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	embedded.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded populates the provided destination NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded from our NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesToNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbedded(destination *v20201101s.NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(embedded.Conditions)

	// DefaultSecurityRules
	if embedded.DefaultSecurityRules != nil {
		defaultSecurityRuleList := make([]v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded, len(embedded.DefaultSecurityRules))
		for defaultSecurityRuleIndex, defaultSecurityRuleItem := range embedded.DefaultSecurityRules {
			// Shadow the loop variable to avoid aliasing
			defaultSecurityRuleItem := defaultSecurityRuleItem
			var defaultSecurityRule v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := defaultSecurityRuleItem.AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(&defaultSecurityRule)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded() to populate field DefaultSecurityRules")
			}
			defaultSecurityRuleList[defaultSecurityRuleIndex] = defaultSecurityRule
		}
		destination.DefaultSecurityRules = defaultSecurityRuleList
	} else {
		destination.DefaultSecurityRules = nil
	}

	// Etag
	destination.Etag = genruntime.ClonePointerToString(embedded.Etag)

	// FlowLogs
	if embedded.FlowLogs != nil {
		flowLogList := make([]v20201101s.FlowLog_Status_SubResourceEmbedded, len(embedded.FlowLogs))
		for flowLogIndex, flowLogItem := range embedded.FlowLogs {
			// Shadow the loop variable to avoid aliasing
			flowLogItem := flowLogItem
			var flowLog v20201101s.FlowLog_Status_SubResourceEmbedded
			err := flowLogItem.AssignPropertiesToFlowLogStatusSubResourceEmbedded(&flowLog)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToFlowLogStatusSubResourceEmbedded() to populate field FlowLogs")
			}
			flowLogList[flowLogIndex] = flowLog
		}
		destination.FlowLogs = flowLogList
	} else {
		destination.FlowLogs = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(embedded.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(embedded.Name)

	// NetworkInterfaces
	if embedded.NetworkInterfaces != nil {
		networkInterfaceList := make([]v20201101s.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded, len(embedded.NetworkInterfaces))
		for networkInterfaceIndex, networkInterfaceItem := range embedded.NetworkInterfaces {
			// Shadow the loop variable to avoid aliasing
			networkInterfaceItem := networkInterfaceItem
			var networkInterface v20201101s.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := networkInterfaceItem.AssignPropertiesToNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(&networkInterface)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded() to populate field NetworkInterfaces")
			}
			networkInterfaceList[networkInterfaceIndex] = networkInterface
		}
		destination.NetworkInterfaces = networkInterfaceList
	} else {
		destination.NetworkInterfaces = nil
	}

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(embedded.ProvisioningState)

	// ResourceGuid
	destination.ResourceGuid = genruntime.ClonePointerToString(embedded.ResourceGuid)

	// SecurityRules
	if embedded.SecurityRules != nil {
		securityRuleList := make([]v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded, len(embedded.SecurityRules))
		for securityRuleIndex, securityRuleItem := range embedded.SecurityRules {
			// Shadow the loop variable to avoid aliasing
			securityRuleItem := securityRuleItem
			var securityRule v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := securityRuleItem.AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(&securityRule)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded() to populate field SecurityRules")
			}
			securityRuleList[securityRuleIndex] = securityRule
		}
		destination.SecurityRules = securityRuleList
	} else {
		destination.SecurityRules = nil
	}

	// Subnets
	if embedded.Subnets != nil {
		subnetList := make([]v20201101s.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded, len(embedded.Subnets))
		for subnetIndex, subnetItem := range embedded.Subnets {
			// Shadow the loop variable to avoid aliasing
			subnetItem := subnetItem
			var subnet v20201101s.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
			err := subnetItem.AssignPropertiesToSubnetStatusNetworkSecurityGroupSubResourceEmbedded(&subnet)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToSubnetStatusNetworkSecurityGroupSubResourceEmbedded() to populate field Subnets")
			}
			subnetList[subnetIndex] = subnet
		}
		destination.Subnets = subnetList
	} else {
		destination.Subnets = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(embedded.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(embedded.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroups_Spec
type NetworkSecurityGroups_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NetworkSecurityGroups_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroups_Spec from the provided source
func (groups *NetworkSecurityGroups_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201101s.NetworkSecurityGroups_Spec)
	if ok {
		// Populate our instance from source
		return groups.AssignPropertiesFromNetworkSecurityGroupsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.NetworkSecurityGroups_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = groups.AssignPropertiesFromNetworkSecurityGroupsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NetworkSecurityGroups_Spec
func (groups *NetworkSecurityGroups_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201101s.NetworkSecurityGroups_Spec)
	if ok {
		// Populate destination from our instance
		return groups.AssignPropertiesToNetworkSecurityGroupsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.NetworkSecurityGroups_Spec{}
	err := groups.AssignPropertiesToNetworkSecurityGroupsSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromNetworkSecurityGroupsSpec populates our NetworkSecurityGroups_Spec from the provided source NetworkSecurityGroups_Spec
func (groups *NetworkSecurityGroups_Spec) AssignPropertiesFromNetworkSecurityGroupsSpec(source *v20201101s.NetworkSecurityGroups_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	groups.AzureName = source.AzureName

	// Location
	groups.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	groups.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		groups.Owner = &owner
	} else {
		groups.Owner = nil
	}

	// Tags
	groups.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		groups.PropertyBag = propertyBag
	} else {
		groups.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNetworkSecurityGroupsSpec populates the provided destination NetworkSecurityGroups_Spec from our NetworkSecurityGroups_Spec
func (groups *NetworkSecurityGroups_Spec) AssignPropertiesToNetworkSecurityGroupsSpec(destination *v20201101s.NetworkSecurityGroups_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(groups.PropertyBag)

	// AzureName
	destination.AzureName = groups.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(groups.Location)

	// OriginalVersion
	destination.OriginalVersion = groups.OriginalVersion

	// Owner
	if groups.Owner != nil {
		owner := groups.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(groups.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20201101.FlowLog_Status_SubResourceEmbedded
//Deprecated version of FlowLog_Status_SubResourceEmbedded. Use v1beta20201101.FlowLog_Status_SubResourceEmbedded instead
type FlowLog_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromFlowLogStatusSubResourceEmbedded populates our FlowLog_Status_SubResourceEmbedded from the provided source FlowLog_Status_SubResourceEmbedded
func (embedded *FlowLog_Status_SubResourceEmbedded) AssignPropertiesFromFlowLogStatusSubResourceEmbedded(source *v20201101s.FlowLog_Status_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToFlowLogStatusSubResourceEmbedded populates the provided destination FlowLog_Status_SubResourceEmbedded from our FlowLog_Status_SubResourceEmbedded
func (embedded *FlowLog_Status_SubResourceEmbedded) AssignPropertiesToFlowLogStatusSubResourceEmbedded(destination *v20201101s.FlowLog_Status_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
//Deprecated version of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded instead
type NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status `json:"extendedLocation,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded populates our NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded from the provided source NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesFromNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(source *v20201101s.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ExtendedLocation
	if source.ExtendedLocation != nil {
		var extendedLocation ExtendedLocation_Status
		err := extendedLocation.AssignPropertiesFromExtendedLocationStatus(source.ExtendedLocation)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromExtendedLocationStatus() to populate field ExtendedLocation")
		}
		embedded.ExtendedLocation = &extendedLocation
	} else {
		embedded.ExtendedLocation = nil
	}

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded populates the provided destination NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded from our NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesToNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbedded(destination *v20201101s.NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// ExtendedLocation
	if embedded.ExtendedLocation != nil {
		var extendedLocation v20201101s.ExtendedLocation_Status
		err := embedded.ExtendedLocation.AssignPropertiesToExtendedLocationStatus(&extendedLocation)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToExtendedLocationStatus() to populate field ExtendedLocation")
		}
		destination.ExtendedLocation = &extendedLocation
	} else {
		destination.ExtendedLocation = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
//Deprecated version of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded instead
type SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded populates our SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded from the provided source SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(source *v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded populates the provided destination SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded from our SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupSubResourceEmbedded(destination *v20201101s.SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
//Deprecated version of Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded. Use v1beta20201101.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded instead
type Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSubnetStatusNetworkSecurityGroupSubResourceEmbedded populates our Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded from the provided source Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesFromSubnetStatusNetworkSecurityGroupSubResourceEmbedded(source *v20201101s.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSubnetStatusNetworkSecurityGroupSubResourceEmbedded populates the provided destination Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded from our Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded
func (embedded *Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded) AssignPropertiesToSubnetStatusNetworkSecurityGroupSubResourceEmbedded(destination *v20201101s.Subnet_Status_NetworkSecurityGroup_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroup{}, &NetworkSecurityGroupList{})
}
