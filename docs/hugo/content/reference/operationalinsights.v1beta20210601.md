---
---
<h2 id="operationalinsights.azure.com/v1beta20210601">operationalinsights.azure.com/v1beta20210601</h2>
<div>
<p>Package v1beta20210601 contains API Schema definitions for the operationalinsights v1beta20210601 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="operationalinsights.azure.com/v1beta20210601.PrivateLinkScopedResource_Status">PrivateLinkScopedResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: The full resource Id of the private link scope resource.</p>
</td>
</tr>
<tr>
<td>
<code>scopeId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScopeId: The private link scope unique Identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.PrivateLinkScopedResource_StatusARM">PrivateLinkScopedResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: The full resource Id of the private link scope resource.</p>
</td>
</tr>
<tr>
<td>
<code>scopeId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScopeId: The private link scope unique Identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.PublicNetworkAccessType_Status">PublicNetworkAccessType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM</a>, <a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.Workspace">Workspace
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/resourceDefinitions/workspaces">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/resourceDefinitions/workspaces</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">
Workspaces_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: The ETag of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures">
WorkspaceFeatures
</a>
</em>
</td>
<td>
<p>Features: Workspace features.</p>
</td>
</tr>
<tr>
<td>
<code>forceCmkForQuery</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesProvisioningState">
WorkspacePropertiesProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForIngestion">
WorkspacePropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForQuery">
WorkspacePropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.</p>
</td>
</tr>
<tr>
<td>
<code>retentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
documentation for details.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku">
WorkspaceSku
</a>
</em>
</td>
<td>
<p>Sku: The SKU (tier) of a workspace.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>workspaceCapping</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCapping">
WorkspaceCapping
</a>
</em>
</td>
<td>
<p>WorkspaceCapping: The daily volume cap for ingestion.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">
Workspace_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceCapping">WorkspaceCapping
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceCapping">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceCapping</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dailyQuotaGb</code><br/>
<em>
float64
</em>
</td>
<td>
<p>DailyQuotaGb: The workspace daily quota for ingestion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceCappingARM">WorkspaceCappingARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceCapping">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceCapping</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dailyQuotaGb</code><br/>
<em>
float64
</em>
</td>
<td>
<p>DailyQuotaGb: The workspace daily quota for ingestion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceCappingStatusDataIngestionStatus">WorkspaceCappingStatusDataIngestionStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCapping_Status">WorkspaceCapping_Status</a>, <a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCapping_StatusARM">WorkspaceCapping_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ApproachingQuota&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ForceOff&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ForceOn&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OverQuota&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RespectQuota&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SubscriptionSuspended&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceCapping_Status">WorkspaceCapping_Status
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dailyQuotaGb</code><br/>
<em>
float64
</em>
</td>
<td>
<p>DailyQuotaGb: The workspace daily quota for ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>dataIngestionStatus</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCappingStatusDataIngestionStatus">
WorkspaceCappingStatusDataIngestionStatus
</a>
</em>
</td>
<td>
<p>DataIngestionStatus: The status of data ingestion for this workspace.</p>
</td>
</tr>
<tr>
<td>
<code>quotaNextResetTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>QuotaNextResetTime: The time when the quota will be rest.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceCapping_StatusARM">WorkspaceCapping_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dailyQuotaGb</code><br/>
<em>
float64
</em>
</td>
<td>
<p>DailyQuotaGb: The workspace daily quota for ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>dataIngestionStatus</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCappingStatusDataIngestionStatus">
WorkspaceCappingStatusDataIngestionStatus
</a>
</em>
</td>
<td>
<p>DataIngestionStatus: The status of data ingestion for this workspace.</p>
</td>
</tr>
<tr>
<td>
<code>quotaNextResetTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>QuotaNextResetTime: The time when the quota will be rest.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures">WorkspaceFeatures
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceFeatures">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceFeatures</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalProperties</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>AdditionalProperties: Unmatched properties from the message are deserialized this collection</p>
</td>
</tr>
<tr>
<td>
<code>clusterResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ClusterResourceReference: Dedicated LA cluster resourceId that is linked to the workspaces.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
</td>
</tr>
<tr>
<td>
<code>enableDataExport</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDataExport: Flag that indicate if data should be exported.</p>
</td>
</tr>
<tr>
<td>
<code>enableLogAccessUsingOnlyResourcePermissions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.</p>
</td>
</tr>
<tr>
<td>
<code>immediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceFeaturesARM">WorkspaceFeaturesARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceFeatures">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceFeatures</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalProperties</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>AdditionalProperties: Unmatched properties from the message are deserialized this collection</p>
</td>
</tr>
<tr>
<td>
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
</td>
</tr>
<tr>
<td>
<code>enableDataExport</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDataExport: Flag that indicate if data should be exported.</p>
</td>
</tr>
<tr>
<td>
<code>enableLogAccessUsingOnlyResourcePermissions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.</p>
</td>
</tr>
<tr>
<td>
<code>immediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures_Status">WorkspaceFeatures_Status
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterResourceId: Dedicated LA cluster resourceId that is linked to the workspaces.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
</td>
</tr>
<tr>
<td>
<code>enableDataExport</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDataExport: Flag that indicate if data should be exported.</p>
</td>
</tr>
<tr>
<td>
<code>enableLogAccessUsingOnlyResourcePermissions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.</p>
</td>
</tr>
<tr>
<td>
<code>immediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures_StatusARM">WorkspaceFeatures_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterResourceId: Dedicated LA cluster resourceId that is linked to the workspaces.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
</td>
</tr>
<tr>
<td>
<code>enableDataExport</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableDataExport: Flag that indicate if data should be exported.</p>
</td>
</tr>
<tr>
<td>
<code>enableLogAccessUsingOnlyResourcePermissions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.</p>
</td>
</tr>
<tr>
<td>
<code>immediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_SpecARM">Workspaces_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceProperties">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceFeaturesARM">
WorkspaceFeaturesARM
</a>
</em>
</td>
<td>
<p>Features: Workspace features.</p>
</td>
</tr>
<tr>
<td>
<code>forceCmkForQuery</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesProvisioningState">
WorkspacePropertiesProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForIngestion">
WorkspacePropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForQuery">
WorkspacePropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.</p>
</td>
</tr>
<tr>
<td>
<code>retentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
documentation for details.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuARM">
WorkspaceSkuARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU (tier) of a workspace.</p>
</td>
</tr>
<tr>
<td>
<code>workspaceCapping</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCappingARM">
WorkspaceCappingARM
</a>
</em>
</td>
<td>
<p>WorkspaceCapping: The daily volume cap for ingestion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesProvisioningState">WorkspacePropertiesProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM</a>, <a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Canceled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ProvisioningAccount&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForIngestion">WorkspacePropertiesPublicNetworkAccessForIngestion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM</a>, <a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForQuery">WorkspacePropertiesPublicNetworkAccessForQuery
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM</a>, <a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesStatusProvisioningState">WorkspacePropertiesStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM</a>, <a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Canceled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ProvisioningAccount&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace_StatusARM">Workspace_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>createdDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedDate: Workspace creation date.</p>
</td>
</tr>
<tr>
<td>
<code>customerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomerId: This is a read-only property. Represents the ID associated with the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures_StatusARM">
WorkspaceFeatures_StatusARM
</a>
</em>
</td>
<td>
<p>Features: Workspace features.</p>
</td>
</tr>
<tr>
<td>
<code>forceCmkForQuery</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.</p>
</td>
</tr>
<tr>
<td>
<code>modifiedDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>ModifiedDate: Workspace modification date.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkScopedResources</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.PrivateLinkScopedResource_StatusARM">
[]PrivateLinkScopedResource_StatusARM
</a>
</em>
</td>
<td>
<p>PrivateLinkScopedResources: List of linked private link scope resources.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesStatusProvisioningState">
WorkspacePropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.</p>
</td>
</tr>
<tr>
<td>
<code>retentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
documentation for details.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku_StatusARM">
WorkspaceSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>workspaceCapping</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCapping_StatusARM">
WorkspaceCapping_StatusARM
</a>
</em>
</td>
<td>
<p>WorkspaceCapping: The daily volume cap for ingestion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSku">WorkspaceSku
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceSku">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceSku</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacityReservationLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
selected.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuName">
WorkspaceSkuName
</a>
</em>
</td>
<td>
<p>Name: The name of the SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSkuARM">WorkspaceSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">WorkspacePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceSku">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceSku</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacityReservationLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
selected.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuName">
WorkspaceSkuName
</a>
</em>
</td>
<td>
<p>Name: The name of the SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSkuName">WorkspaceSkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku">WorkspaceSku</a>, <a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuARM">WorkspaceSkuARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CapacityReservation&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Free&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LACluster&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PerGB2018&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PerNode&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standalone&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSkuStatusCapacityReservationLevel">WorkspaceSkuStatusCapacityReservationLevel
(<code>int</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku_Status">WorkspaceSku_Status</a>, <a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku_StatusARM">WorkspaceSku_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>100</p></td>
<td></td>
</tr><tr><td><p>1000</p></td>
<td></td>
</tr><tr><td><p>200</p></td>
<td></td>
</tr><tr><td><p>2000</p></td>
<td></td>
</tr><tr><td><p>300</p></td>
<td></td>
</tr><tr><td><p>400</p></td>
<td></td>
</tr><tr><td><p>500</p></td>
<td></td>
</tr><tr><td><p>5000</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSkuStatusName">WorkspaceSkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku_Status">WorkspaceSku_Status</a>, <a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku_StatusARM">WorkspaceSku_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CapacityReservation&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Free&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LACluster&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PerGB2018&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PerNode&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standalone&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSku_Status">WorkspaceSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacityReservationLevel</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuStatusCapacityReservationLevel">
WorkspaceSkuStatusCapacityReservationLevel
</a>
</em>
</td>
<td>
<p>CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
selected.</p>
</td>
</tr>
<tr>
<td>
<code>lastSkuUpdate</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastSkuUpdate: The last time when the sku was updated.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuStatusName">
WorkspaceSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The name of the SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspaceSku_StatusARM">WorkspaceSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">WorkspaceProperties_StatusARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacityReservationLevel</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuStatusCapacityReservationLevel">
WorkspaceSkuStatusCapacityReservationLevel
</a>
</em>
</td>
<td>
<p>CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
selected.</p>
</td>
</tr>
<tr>
<td>
<code>lastSkuUpdate</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastSkuUpdate: The last time when the sku was updated.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSkuStatusName">
WorkspaceSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The name of the SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.Workspace_Status">Workspace_Status
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace">Workspace</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>createdDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedDate: Workspace creation date.</p>
</td>
</tr>
<tr>
<td>
<code>customerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomerId: This is a read-only property. Represents the ID associated with the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: The ETag of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures_Status">
WorkspaceFeatures_Status
</a>
</em>
</td>
<td>
<p>Features: Workspace features.</p>
</td>
</tr>
<tr>
<td>
<code>forceCmkForQuery</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>modifiedDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>ModifiedDate: Workspace modification date.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkScopedResources</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.PrivateLinkScopedResource_Status">
[]PrivateLinkScopedResource_Status
</a>
</em>
</td>
<td>
<p>PrivateLinkScopedResources: List of linked private link scope resources.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesStatusProvisioningState">
WorkspacePropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.</p>
</td>
</tr>
<tr>
<td>
<code>retentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
documentation for details.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku_Status">
WorkspaceSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The SKU of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>workspaceCapping</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCapping_Status">
WorkspaceCapping_Status
</a>
</em>
</td>
<td>
<p>WorkspaceCapping: The daily volume cap for ingestion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.Workspace_StatusARM">Workspace_StatusARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: The ETag of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceProperties_StatusARM">
WorkspaceProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Workspace properties.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.WorkspacesSpecAPIVersion">WorkspacesSpecAPIVersion
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;2021-06-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.Workspaces_Spec">Workspaces_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#operationalinsights.azure.com/v1beta20210601.Workspace">Workspace</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: The ETag of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceFeatures">
WorkspaceFeatures
</a>
</em>
</td>
<td>
<p>Features: Workspace features.</p>
</td>
</tr>
<tr>
<td>
<code>forceCmkForQuery</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesProvisioningState">
WorkspacePropertiesProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForIngestion">
WorkspacePropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesPublicNetworkAccessForQuery">
WorkspacePropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.</p>
</td>
</tr>
<tr>
<td>
<code>retentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
documentation for details.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceSku">
WorkspaceSku
</a>
</em>
</td>
<td>
<p>Sku: The SKU (tier) of a workspace.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>workspaceCapping</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspaceCapping">
WorkspaceCapping
</a>
</em>
</td>
<td>
<p>WorkspaceCapping: The daily volume cap for ingestion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="operationalinsights.azure.com/v1beta20210601.Workspaces_SpecARM">Workspaces_SpecARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: The ETag of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the workspace.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#operationalinsights.azure.com/v1beta20210601.WorkspacePropertiesARM">
WorkspacePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Workspace properties.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</tbody>
</table>
<hr/>
