// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

//Deprecated version of EventSubscriptions_Spec. Use v1beta20200601.EventSubscriptions_Spec instead
type EventSubscriptions_SpecARM struct {
	Location   *string                         `json:"location,omitempty"`
	Name       string                          `json:"name,omitempty"`
	Properties *EventSubscriptionPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string               `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &EventSubscriptions_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscriptions EventSubscriptions_SpecARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (subscriptions EventSubscriptions_SpecARM) GetName() string {
	return subscriptions.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscriptions EventSubscriptions_SpecARM) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

//Deprecated version of EventSubscriptionProperties. Use v1beta20200601.EventSubscriptionProperties instead
type EventSubscriptionPropertiesARM struct {
	DeadLetterDestination *StorageBlobDeadLetterDestinationARM            `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestinationARM                `json:"destination,omitempty"`
	EventDeliverySchema   *EventSubscriptionPropertiesEventDeliverySchema `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                                         `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilterARM                     `json:"filter,omitempty"`
	Labels                []string                                        `json:"labels,omitempty"`
	RetryPolicy           *RetryPolicyARM                                 `json:"retryPolicy,omitempty"`
}

//Deprecated version of EventSubscriptionDestination. Use v1beta20200601.EventSubscriptionDestination instead
type EventSubscriptionDestinationARM struct {
	AzureFunction    *AzureFunctionEventSubscriptionDestinationARM    `json:"azureFunctionEventSubscriptionDestination,omitempty"`
	EventHub         *EventHubEventSubscriptionDestinationARM         `json:"eventHubEventSubscriptionDestination,omitempty"`
	HybridConnection *HybridConnectionEventSubscriptionDestinationARM `json:"hybridConnectionEventSubscriptionDestination,omitempty"`
	ServiceBusQueue  *ServiceBusQueueEventSubscriptionDestinationARM  `json:"serviceBusQueueEventSubscriptionDestination,omitempty"`
	ServiceBusTopic  *ServiceBusTopicEventSubscriptionDestinationARM  `json:"serviceBusTopicEventSubscriptionDestination,omitempty"`
	StorageQueue     *StorageQueueEventSubscriptionDestinationARM     `json:"storageQueueEventSubscriptionDestination,omitempty"`
	WebHook          *WebHookEventSubscriptionDestinationARM          `json:"webHookEventSubscriptionDestination,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because EventSubscriptionDestinationARM represents a discriminated union (JSON OneOf)
func (destination EventSubscriptionDestinationARM) MarshalJSON() ([]byte, error) {
	if destination.AzureFunction != nil {
		return json.Marshal(destination.AzureFunction)
	}
	if destination.EventHub != nil {
		return json.Marshal(destination.EventHub)
	}
	if destination.HybridConnection != nil {
		return json.Marshal(destination.HybridConnection)
	}
	if destination.ServiceBusQueue != nil {
		return json.Marshal(destination.ServiceBusQueue)
	}
	if destination.ServiceBusTopic != nil {
		return json.Marshal(destination.ServiceBusTopic)
	}
	if destination.StorageQueue != nil {
		return json.Marshal(destination.StorageQueue)
	}
	if destination.WebHook != nil {
		return json.Marshal(destination.WebHook)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the EventSubscriptionDestinationARM
func (destination *EventSubscriptionDestinationARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "AzureFunction" {
		destination.AzureFunction = &AzureFunctionEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.AzureFunction)
	}
	if discriminator == "EventHub" {
		destination.EventHub = &EventHubEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.EventHub)
	}
	if discriminator == "HybridConnection" {
		destination.HybridConnection = &HybridConnectionEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.HybridConnection)
	}
	if discriminator == "ServiceBusQueue" {
		destination.ServiceBusQueue = &ServiceBusQueueEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.ServiceBusQueue)
	}
	if discriminator == "ServiceBusTopic" {
		destination.ServiceBusTopic = &ServiceBusTopicEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.ServiceBusTopic)
	}
	if discriminator == "StorageQueue" {
		destination.StorageQueue = &StorageQueueEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.StorageQueue)
	}
	if discriminator == "WebHook" {
		destination.WebHook = &WebHookEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.WebHook)
	}

	// No error
	return nil
}

//Deprecated version of EventSubscriptionFilter. Use v1beta20200601.EventSubscriptionFilter instead
type EventSubscriptionFilterARM struct {
	AdvancedFilters        []AdvancedFilterARM `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string            `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool               `json:"isSubjectCaseSensitive,omitempty"`
	SubjectBeginsWith      *string             `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string             `json:"subjectEndsWith,omitempty"`
}

//Deprecated version of RetryPolicy. Use v1beta20200601.RetryPolicy instead
type RetryPolicyARM struct {
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int `json:"maxDeliveryAttempts,omitempty"`
}

//Deprecated version of StorageBlobDeadLetterDestination. Use v1beta20200601.StorageBlobDeadLetterDestination instead
type StorageBlobDeadLetterDestinationARM struct {
	EndpointType *StorageBlobDeadLetterDestinationEndpointType  `json:"endpointType,omitempty"`
	Properties   *StorageBlobDeadLetterDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of AdvancedFilter. Use v1beta20200601.AdvancedFilter instead
type AdvancedFilterARM struct {
	BoolEquals                *AdvancedFilter_BoolEqualsARM                `json:"boolEqualsAdvancedFilter,omitempty"`
	NumberGreaterThan         *AdvancedFilter_NumberGreaterThanARM         `json:"numberGreaterThanAdvancedFilter,omitempty"`
	NumberGreaterThanOrEquals *AdvancedFilter_NumberGreaterThanOrEqualsARM `json:"numberGreaterThanOrEqualsAdvancedFilter,omitempty"`
	NumberIn                  *AdvancedFilter_NumberInARM                  `json:"numberInAdvancedFilter,omitempty"`
	NumberLessThan            *AdvancedFilter_NumberLessThanARM            `json:"numberLessThanAdvancedFilter,omitempty"`
	NumberLessThanOrEquals    *AdvancedFilter_NumberLessThanOrEqualsARM    `json:"numberLessThanOrEqualsAdvancedFilter,omitempty"`
	NumberNotIn               *AdvancedFilter_NumberNotInARM               `json:"numberNotInAdvancedFilter,omitempty"`
	StringBeginsWith          *AdvancedFilter_StringBeginsWithARM          `json:"stringBeginsWithAdvancedFilter,omitempty"`
	StringContains            *AdvancedFilter_StringContainsARM            `json:"stringContainsAdvancedFilter,omitempty"`
	StringEndsWith            *AdvancedFilter_StringEndsWithARM            `json:"stringEndsWithAdvancedFilter,omitempty"`
	StringIn                  *AdvancedFilter_StringInARM                  `json:"stringInAdvancedFilter,omitempty"`
	StringNotIn               *AdvancedFilter_StringNotInARM               `json:"stringNotInAdvancedFilter,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AdvancedFilterARM represents a discriminated union (JSON OneOf)
func (filter AdvancedFilterARM) MarshalJSON() ([]byte, error) {
	if filter.BoolEquals != nil {
		return json.Marshal(filter.BoolEquals)
	}
	if filter.NumberGreaterThan != nil {
		return json.Marshal(filter.NumberGreaterThan)
	}
	if filter.NumberGreaterThanOrEquals != nil {
		return json.Marshal(filter.NumberGreaterThanOrEquals)
	}
	if filter.NumberIn != nil {
		return json.Marshal(filter.NumberIn)
	}
	if filter.NumberLessThan != nil {
		return json.Marshal(filter.NumberLessThan)
	}
	if filter.NumberLessThanOrEquals != nil {
		return json.Marshal(filter.NumberLessThanOrEquals)
	}
	if filter.NumberNotIn != nil {
		return json.Marshal(filter.NumberNotIn)
	}
	if filter.StringBeginsWith != nil {
		return json.Marshal(filter.StringBeginsWith)
	}
	if filter.StringContains != nil {
		return json.Marshal(filter.StringContains)
	}
	if filter.StringEndsWith != nil {
		return json.Marshal(filter.StringEndsWith)
	}
	if filter.StringIn != nil {
		return json.Marshal(filter.StringIn)
	}
	if filter.StringNotIn != nil {
		return json.Marshal(filter.StringNotIn)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the AdvancedFilterARM
func (filter *AdvancedFilterARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["operatorType"]
	if discriminator == "BoolEquals" {
		filter.BoolEquals = &AdvancedFilter_BoolEqualsARM{}
		return json.Unmarshal(data, filter.BoolEquals)
	}
	if discriminator == "NumberGreaterThan" {
		filter.NumberGreaterThan = &AdvancedFilter_NumberGreaterThanARM{}
		return json.Unmarshal(data, filter.NumberGreaterThan)
	}
	if discriminator == "NumberGreaterThanOrEquals" {
		filter.NumberGreaterThanOrEquals = &AdvancedFilter_NumberGreaterThanOrEqualsARM{}
		return json.Unmarshal(data, filter.NumberGreaterThanOrEquals)
	}
	if discriminator == "NumberIn" {
		filter.NumberIn = &AdvancedFilter_NumberInARM{}
		return json.Unmarshal(data, filter.NumberIn)
	}
	if discriminator == "NumberLessThan" {
		filter.NumberLessThan = &AdvancedFilter_NumberLessThanARM{}
		return json.Unmarshal(data, filter.NumberLessThan)
	}
	if discriminator == "NumberLessThanOrEquals" {
		filter.NumberLessThanOrEquals = &AdvancedFilter_NumberLessThanOrEqualsARM{}
		return json.Unmarshal(data, filter.NumberLessThanOrEquals)
	}
	if discriminator == "NumberNotIn" {
		filter.NumberNotIn = &AdvancedFilter_NumberNotInARM{}
		return json.Unmarshal(data, filter.NumberNotIn)
	}
	if discriminator == "StringBeginsWith" {
		filter.StringBeginsWith = &AdvancedFilter_StringBeginsWithARM{}
		return json.Unmarshal(data, filter.StringBeginsWith)
	}
	if discriminator == "StringContains" {
		filter.StringContains = &AdvancedFilter_StringContainsARM{}
		return json.Unmarshal(data, filter.StringContains)
	}
	if discriminator == "StringEndsWith" {
		filter.StringEndsWith = &AdvancedFilter_StringEndsWithARM{}
		return json.Unmarshal(data, filter.StringEndsWith)
	}
	if discriminator == "StringIn" {
		filter.StringIn = &AdvancedFilter_StringInARM{}
		return json.Unmarshal(data, filter.StringIn)
	}
	if discriminator == "StringNotIn" {
		filter.StringNotIn = &AdvancedFilter_StringNotInARM{}
		return json.Unmarshal(data, filter.StringNotIn)
	}

	// No error
	return nil
}

//Deprecated version of AzureFunctionEventSubscriptionDestination. Use v1beta20200601.AzureFunctionEventSubscriptionDestination instead
type AzureFunctionEventSubscriptionDestinationARM struct {
	EndpointType AzureFunctionEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *AzureFunctionEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of EventHubEventSubscriptionDestination. Use v1beta20200601.EventHubEventSubscriptionDestination instead
type EventHubEventSubscriptionDestinationARM struct {
	EndpointType EventHubEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *EventHubEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of HybridConnectionEventSubscriptionDestination. Use v1beta20200601.HybridConnectionEventSubscriptionDestination instead
type HybridConnectionEventSubscriptionDestinationARM struct {
	EndpointType HybridConnectionEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *HybridConnectionEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of ServiceBusQueueEventSubscriptionDestination. Use v1beta20200601.ServiceBusQueueEventSubscriptionDestination instead
type ServiceBusQueueEventSubscriptionDestinationARM struct {
	EndpointType ServiceBusQueueEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *ServiceBusQueueEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of ServiceBusTopicEventSubscriptionDestination. Use v1beta20200601.ServiceBusTopicEventSubscriptionDestination instead
type ServiceBusTopicEventSubscriptionDestinationARM struct {
	EndpointType ServiceBusTopicEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *ServiceBusTopicEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of StorageBlobDeadLetterDestinationProperties. Use v1beta20200601.StorageBlobDeadLetterDestinationProperties instead
type StorageBlobDeadLetterDestinationPropertiesARM struct {
	BlobContainerName *string `json:"blobContainerName,omitempty"`
	ResourceId        *string `json:"resourceId,omitempty"`
}

//Deprecated version of StorageQueueEventSubscriptionDestination. Use v1beta20200601.StorageQueueEventSubscriptionDestination instead
type StorageQueueEventSubscriptionDestinationARM struct {
	EndpointType StorageQueueEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *StorageQueueEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of WebHookEventSubscriptionDestination. Use v1beta20200601.WebHookEventSubscriptionDestination instead
type WebHookEventSubscriptionDestinationARM struct {
	EndpointType WebHookEventSubscriptionDestinationEndpointType   `json:"endpointType,omitempty"`
	Properties   *WebHookEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of AdvancedFilter_BoolEquals. Use v1beta20200601.AdvancedFilter_BoolEquals instead
type AdvancedFilter_BoolEqualsARM struct {
	Key          *string                              `json:"key,omitempty"`
	OperatorType AdvancedFilterBoolEqualsOperatorType `json:"operatorType,omitempty"`
	Value        *bool                                `json:"value,omitempty"`
}

//Deprecated version of AdvancedFilter_NumberGreaterThan. Use v1beta20200601.AdvancedFilter_NumberGreaterThan instead
type AdvancedFilter_NumberGreaterThanARM struct {
	Key          *string                                     `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberGreaterThanOperatorType `json:"operatorType,omitempty"`
	Value        *float64                                    `json:"value,omitempty"`
}

//Deprecated version of AdvancedFilter_NumberGreaterThanOrEquals. Use v1beta20200601.AdvancedFilter_NumberGreaterThanOrEquals instead
type AdvancedFilter_NumberGreaterThanOrEqualsARM struct {
	Key          *string                                             `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberGreaterThanOrEqualsOperatorType `json:"operatorType,omitempty"`
	Value        *float64                                            `json:"value,omitempty"`
}

//Deprecated version of AdvancedFilter_NumberIn. Use v1beta20200601.AdvancedFilter_NumberIn instead
type AdvancedFilter_NumberInARM struct {
	Key          *string                            `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberInOperatorType `json:"operatorType,omitempty"`
	Values       []float64                          `json:"values,omitempty"`
}

//Deprecated version of AdvancedFilter_NumberLessThan. Use v1beta20200601.AdvancedFilter_NumberLessThan instead
type AdvancedFilter_NumberLessThanARM struct {
	Key          *string                                  `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberLessThanOperatorType `json:"operatorType,omitempty"`
	Value        *float64                                 `json:"value,omitempty"`
}

//Deprecated version of AdvancedFilter_NumberLessThanOrEquals. Use v1beta20200601.AdvancedFilter_NumberLessThanOrEquals instead
type AdvancedFilter_NumberLessThanOrEqualsARM struct {
	Key          *string                                          `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberLessThanOrEqualsOperatorType `json:"operatorType,omitempty"`
	Value        *float64                                         `json:"value,omitempty"`
}

//Deprecated version of AdvancedFilter_NumberNotIn. Use v1beta20200601.AdvancedFilter_NumberNotIn instead
type AdvancedFilter_NumberNotInARM struct {
	Key          *string                               `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberNotInOperatorType `json:"operatorType,omitempty"`
	Values       []float64                             `json:"values,omitempty"`
}

//Deprecated version of AdvancedFilter_StringBeginsWith. Use v1beta20200601.AdvancedFilter_StringBeginsWith instead
type AdvancedFilter_StringBeginsWithARM struct {
	Key          *string                                    `json:"key,omitempty"`
	OperatorType AdvancedFilterStringBeginsWithOperatorType `json:"operatorType,omitempty"`
	Values       []string                                   `json:"values,omitempty"`
}

//Deprecated version of AdvancedFilter_StringContains. Use v1beta20200601.AdvancedFilter_StringContains instead
type AdvancedFilter_StringContainsARM struct {
	Key          *string                                  `json:"key,omitempty"`
	OperatorType AdvancedFilterStringContainsOperatorType `json:"operatorType,omitempty"`
	Values       []string                                 `json:"values,omitempty"`
}

//Deprecated version of AdvancedFilter_StringEndsWith. Use v1beta20200601.AdvancedFilter_StringEndsWith instead
type AdvancedFilter_StringEndsWithARM struct {
	Key          *string                                  `json:"key,omitempty"`
	OperatorType AdvancedFilterStringEndsWithOperatorType `json:"operatorType,omitempty"`
	Values       []string                                 `json:"values,omitempty"`
}

//Deprecated version of AdvancedFilter_StringIn. Use v1beta20200601.AdvancedFilter_StringIn instead
type AdvancedFilter_StringInARM struct {
	Key          *string                            `json:"key,omitempty"`
	OperatorType AdvancedFilterStringInOperatorType `json:"operatorType,omitempty"`
	Values       []string                           `json:"values,omitempty"`
}

//Deprecated version of AdvancedFilter_StringNotIn. Use v1beta20200601.AdvancedFilter_StringNotIn instead
type AdvancedFilter_StringNotInARM struct {
	Key          *string                               `json:"key,omitempty"`
	OperatorType AdvancedFilterStringNotInOperatorType `json:"operatorType,omitempty"`
	Values       []string                              `json:"values,omitempty"`
}

//Deprecated version of AzureFunctionEventSubscriptionDestinationProperties. Use v1beta20200601.AzureFunctionEventSubscriptionDestinationProperties instead
type AzureFunctionEventSubscriptionDestinationPropertiesARM struct {
	MaxEventsPerBatch             *int    `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes *int    `json:"preferredBatchSizeInKilobytes,omitempty"`
	ResourceId                    *string `json:"resourceId,omitempty"`
}

//Deprecated version of EventHubEventSubscriptionDestinationProperties. Use v1beta20200601.EventHubEventSubscriptionDestinationProperties instead
type EventHubEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Deprecated version of HybridConnectionEventSubscriptionDestinationProperties. Use v1beta20200601.HybridConnectionEventSubscriptionDestinationProperties instead
type HybridConnectionEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Deprecated version of ServiceBusQueueEventSubscriptionDestinationProperties. Use v1beta20200601.ServiceBusQueueEventSubscriptionDestinationProperties instead
type ServiceBusQueueEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Deprecated version of ServiceBusTopicEventSubscriptionDestinationProperties. Use v1beta20200601.ServiceBusTopicEventSubscriptionDestinationProperties instead
type ServiceBusTopicEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Deprecated version of StorageQueueEventSubscriptionDestinationProperties. Use v1beta20200601.StorageQueueEventSubscriptionDestinationProperties instead
type StorageQueueEventSubscriptionDestinationPropertiesARM struct {
	QueueName  *string `json:"queueName,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

//Deprecated version of WebHookEventSubscriptionDestinationProperties. Use v1beta20200601.WebHookEventSubscriptionDestinationProperties instead
type WebHookEventSubscriptionDestinationPropertiesARM struct {
	AzureActiveDirectoryApplicationIdOrUri *string `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`
	AzureActiveDirectoryTenantId           *string `json:"azureActiveDirectoryTenantId,omitempty"`
	EndpointUrl                            *string `json:"endpointUrl,omitempty"`
	MaxEventsPerBatch                      *int    `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes          *int    `json:"preferredBatchSizeInKilobytes,omitempty"`
}
