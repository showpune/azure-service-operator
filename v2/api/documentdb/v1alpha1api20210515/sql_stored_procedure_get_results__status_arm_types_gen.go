// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

//Deprecated version of SqlStoredProcedureGetResults_Status. Use v1beta20210515.SqlStoredProcedureGetResults_Status instead
type SqlStoredProcedureGetResults_StatusARM struct {
	Id         *string                                    `json:"id,omitempty"`
	Location   *string                                    `json:"location,omitempty"`
	Name       *string                                    `json:"name,omitempty"`
	Properties *SqlStoredProcedureGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
	Type       *string                                    `json:"type,omitempty"`
}

//Deprecated version of SqlStoredProcedureGetProperties_Status. Use v1beta20210515.SqlStoredProcedureGetProperties_Status instead
type SqlStoredProcedureGetProperties_StatusARM struct {
	Resource *SqlStoredProcedureGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

//Deprecated version of SqlStoredProcedureGetProperties_Status_Resource. Use v1beta20210515.SqlStoredProcedureGetProperties_Status_Resource instead
type SqlStoredProcedureGetProperties_Status_ResourceARM struct {
	Body *string  `json:"body,omitempty"`
	Etag *string  `json:"_etag,omitempty"`
	Id   *string  `json:"id,omitempty"`
	Rid  *string  `json:"_rid,omitempty"`
	Ts   *float64 `json:"_ts,omitempty"`
}
