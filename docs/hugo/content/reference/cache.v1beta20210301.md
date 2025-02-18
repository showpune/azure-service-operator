---
---
<h2 id="cache.azure.com/v1beta20210301">cache.azure.com/v1beta20210301</h2>
<div>
<p>Package v1beta20210301 contains API Schema definitions for the cache v1beta20210301 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="cache.azure.com/v1beta20210301.ClusterPropertiesARM">ClusterPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterprise_SpecARM">RedisEnterprise_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/ClusterProperties">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/ClusterProperties</a></p>
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ClusterPropertiesMinimumTlsVersion">
ClusterPropertiesMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. &lsquo;1.2&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.ClusterPropertiesMinimumTlsVersion">ClusterPropertiesMinimumTlsVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.ClusterPropertiesARM">ClusterPropertiesARM</a>, <a href="#cache.azure.com/v1beta20210301.RedisEnterprise_Spec">RedisEnterprise_Spec</a>)
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
<tbody><tr><td><p>&#34;1.0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.ClusterPropertiesStatusMinimumTlsVersion">ClusterPropertiesStatusMinimumTlsVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.ClusterProperties_StatusARM">ClusterProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Cluster_Status">Cluster_Status</a>)
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
<tbody><tr><td><p>&#34;1.0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.ClusterProperties_StatusARM">ClusterProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Cluster_StatusARM">Cluster_StatusARM</a>)
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
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: DNS name of the cluster endpoint</p>
</td>
</tr>
<tr>
<td>
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ClusterPropertiesStatusMinimumTlsVersion">
ClusterPropertiesStatusMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. &lsquo;1.2&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the specified RedisEnterprise cluster</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Current provisioning status of the cluster</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Version of redis the cluster supports, e.g. &lsquo;6&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>resourceState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ResourceState_Status">
ResourceState_Status
</a>
</em>
</td>
<td>
<p>ResourceState: Current resource status of the cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Cluster_Status">Cluster_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterprise">RedisEnterprise</a>)
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
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: DNS name of the cluster endpoint</p>
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ClusterPropertiesStatusMinimumTlsVersion">
ClusterPropertiesStatusMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. &lsquo;1.2&rsquo;</p>
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
<code>privateEndpointConnections</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PrivateEndpointConnection_Status_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the specified RedisEnterprise cluster</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Current provisioning status of the cluster</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Version of redis the cluster supports, e.g. &lsquo;6&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>resourceState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ResourceState_Status">
ResourceState_Status
</a>
</em>
</td>
<td>
<p>ResourceState: Current resource status of the cluster</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Sku_Status">
Sku_Status
</a>
</em>
</td>
<td>
<p>Sku: The SKU to create, which affects price, performance, and features.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The Availability Zones where this cluster will be deployed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Cluster_StatusARM">Cluster_StatusARM
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
<a href="#cache.azure.com/v1beta20210301.ClusterProperties_StatusARM">
ClusterProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Other properties of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Sku_StatusARM">
Sku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU to create, which affects price, performance, and features.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The Availability Zones where this cluster will be deployed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesARM">DatabasePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_SpecARM">RedisEnterpriseDatabases_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/DatabaseProperties">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/DatabaseProperties</a></p>
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
<code>clientProtocol</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesClientProtocol">
DatabasePropertiesClientProtocol
</a>
</em>
</td>
<td>
<p>ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
TLS-encrypted.</p>
</td>
</tr>
<tr>
<td>
<code>clusteringPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesClusteringPolicy">
DatabasePropertiesClusteringPolicy
</a>
</em>
</td>
<td>
<p>ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesEvictionPolicy">
DatabasePropertiesEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Redis eviction policy - default is VolatileLRU.</p>
</td>
</tr>
<tr>
<td>
<code>modules</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ModuleARM">
[]ModuleARM
</a>
</em>
</td>
<td>
<p>Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceARM">
PersistenceARM
</a>
</em>
</td>
<td>
<p>Persistence: Persistence-related configuration for the RedisEnterprise database</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesClientProtocol">DatabasePropertiesClientProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabasePropertiesARM">DatabasePropertiesARM</a>, <a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">RedisEnterpriseDatabases_Spec</a>)
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
<tbody><tr><td><p>&#34;Encrypted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Plaintext&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesClusteringPolicy">DatabasePropertiesClusteringPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabasePropertiesARM">DatabasePropertiesARM</a>, <a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">RedisEnterpriseDatabases_Spec</a>)
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
<tbody><tr><td><p>&#34;EnterpriseCluster&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OSSCluster&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesEvictionPolicy">DatabasePropertiesEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabasePropertiesARM">DatabasePropertiesARM</a>, <a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">RedisEnterpriseDatabases_Spec</a>)
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
<tbody><tr><td><p>&#34;AllKeysLFU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllKeysLRU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllKeysRandom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NoEviction&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileLFU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileLRU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileRandom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileTTL&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesStatusClientProtocol">DatabasePropertiesStatusClientProtocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
<tbody><tr><td><p>&#34;Encrypted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Plaintext&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesStatusClusteringPolicy">DatabasePropertiesStatusClusteringPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
<tbody><tr><td><p>&#34;EnterpriseCluster&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OSSCluster&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabasePropertiesStatusEvictionPolicy">DatabasePropertiesStatusEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
<tbody><tr><td><p>&#34;AllKeysLFU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllKeysLRU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllKeysRandom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NoEviction&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileLFU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileLRU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileRandom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VolatileTTL&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Database_StatusARM">Database_StatusARM</a>)
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
<code>clientProtocol</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesStatusClientProtocol">
DatabasePropertiesStatusClientProtocol
</a>
</em>
</td>
<td>
<p>ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
TLS-encrypted.</p>
</td>
</tr>
<tr>
<td>
<code>clusteringPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesStatusClusteringPolicy">
DatabasePropertiesStatusClusteringPolicy
</a>
</em>
</td>
<td>
<p>ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesStatusEvictionPolicy">
DatabasePropertiesStatusEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Redis eviction policy - default is VolatileLRU</p>
</td>
</tr>
<tr>
<td>
<code>modules</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Module_StatusARM">
[]Module_StatusARM
</a>
</em>
</td>
<td>
<p>Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Persistence_StatusARM">
Persistence_StatusARM
</a>
</em>
</td>
<td>
<p>Persistence: Persistence settings</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Current provisioning status of the database</p>
</td>
</tr>
<tr>
<td>
<code>resourceState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ResourceState_Status">
ResourceState_Status
</a>
</em>
</td>
<td>
<p>ResourceState: Current resource status of the database</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Database_Status">Database_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabase">RedisEnterpriseDatabase</a>)
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
<code>clientProtocol</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesStatusClientProtocol">
DatabasePropertiesStatusClientProtocol
</a>
</em>
</td>
<td>
<p>ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
TLS-encrypted.</p>
</td>
</tr>
<tr>
<td>
<code>clusteringPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesStatusClusteringPolicy">
DatabasePropertiesStatusClusteringPolicy
</a>
</em>
</td>
<td>
<p>ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.</p>
</td>
</tr>
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
<code>evictionPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesStatusEvictionPolicy">
DatabasePropertiesStatusEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Redis eviction policy - default is VolatileLRU</p>
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
<code>modules</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Module_Status">
[]Module_Status
</a>
</em>
</td>
<td>
<p>Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.</p>
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
<code>persistence</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Persistence_Status">
Persistence_Status
</a>
</em>
</td>
<td>
<p>Persistence: Persistence settings</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Current provisioning status of the database</p>
</td>
</tr>
<tr>
<td>
<code>resourceState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ResourceState_Status">
ResourceState_Status
</a>
</em>
</td>
<td>
<p>ResourceState: Current resource status of the database</p>
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
<h3 id="cache.azure.com/v1beta20210301.Database_StatusARM">Database_StatusARM
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
<a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">
DatabaseProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Other properties of the database.</p>
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
<h3 id="cache.azure.com/v1beta20210301.Module">Module
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">RedisEnterpriseDatabases_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Module">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Module</a></p>
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
<code>args</code><br/>
<em>
string
</em>
</td>
<td>
<p>Args: Configuration options for the module, e.g. &lsquo;ERROR_RATE 0.00 INITIAL_SIZE 400&rsquo;.</p>
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
<p>Name: The name of the module, e.g. &lsquo;RedisBloom&rsquo;, &lsquo;RediSearch&rsquo;, &lsquo;RedisTimeSeries&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.ModuleARM">ModuleARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabasePropertiesARM">DatabasePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Module">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Module</a></p>
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
<code>args</code><br/>
<em>
string
</em>
</td>
<td>
<p>Args: Configuration options for the module, e.g. &lsquo;ERROR_RATE 0.00 INITIAL_SIZE 400&rsquo;.</p>
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
<p>Name: The name of the module, e.g. &lsquo;RedisBloom&rsquo;, &lsquo;RediSearch&rsquo;, &lsquo;RedisTimeSeries&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Module_Status">Module_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
<code>args</code><br/>
<em>
string
</em>
</td>
<td>
<p>Args: Configuration options for the module, e.g. &lsquo;ERROR_RATE 0.00 INITIAL_SIZE 400&rsquo;.</p>
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
<p>Name: The name of the module, e.g. &lsquo;RedisBloom&rsquo;, &lsquo;RediSearch&rsquo;, &lsquo;RedisTimeSeries&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of the module, e.g. &lsquo;1.0&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Module_StatusARM">Module_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>)
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
<code>args</code><br/>
<em>
string
</em>
</td>
<td>
<p>Args: Configuration options for the module, e.g. &lsquo;ERROR_RATE 0.00 INITIAL_SIZE 400&rsquo;.</p>
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
<p>Name: The name of the module, e.g. &lsquo;RedisBloom&rsquo;, &lsquo;RediSearch&rsquo;, &lsquo;RedisTimeSeries&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of the module, e.g. &lsquo;1.0&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Persistence">Persistence
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">RedisEnterpriseDatabases_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Persistence">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Persistence</a></p>
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
<code>aofEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AofEnabled: Sets whether AOF is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>aofFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceAofFrequency">
PersistenceAofFrequency
</a>
</em>
</td>
<td>
<p>AofFrequency: Sets the frequency at which data is written to disk.</p>
</td>
</tr>
<tr>
<td>
<code>rdbEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RdbEnabled: Sets whether RDB is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>rdbFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceRdbFrequency">
PersistenceRdbFrequency
</a>
</em>
</td>
<td>
<p>RdbFrequency: Sets the frequency at which a snapshot of the database is created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PersistenceARM">PersistenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabasePropertiesARM">DatabasePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Persistence">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Persistence</a></p>
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
<code>aofEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AofEnabled: Sets whether AOF is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>aofFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceAofFrequency">
PersistenceAofFrequency
</a>
</em>
</td>
<td>
<p>AofFrequency: Sets the frequency at which data is written to disk.</p>
</td>
</tr>
<tr>
<td>
<code>rdbEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RdbEnabled: Sets whether RDB is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>rdbFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceRdbFrequency">
PersistenceRdbFrequency
</a>
</em>
</td>
<td>
<p>RdbFrequency: Sets the frequency at which a snapshot of the database is created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PersistenceAofFrequency">PersistenceAofFrequency
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Persistence">Persistence</a>, <a href="#cache.azure.com/v1beta20210301.PersistenceARM">PersistenceARM</a>)
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
<tbody><tr><td><p>&#34;1s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;always&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PersistenceRdbFrequency">PersistenceRdbFrequency
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Persistence">Persistence</a>, <a href="#cache.azure.com/v1beta20210301.PersistenceARM">PersistenceARM</a>)
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
<tbody><tr><td><p>&#34;12h&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1h&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;6h&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PersistenceStatusAofFrequency">PersistenceStatusAofFrequency
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Persistence_Status">Persistence_Status</a>, <a href="#cache.azure.com/v1beta20210301.Persistence_StatusARM">Persistence_StatusARM</a>)
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
<tbody><tr><td><p>&#34;1s&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;always&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PersistenceStatusRdbFrequency">PersistenceStatusRdbFrequency
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Persistence_Status">Persistence_Status</a>, <a href="#cache.azure.com/v1beta20210301.Persistence_StatusARM">Persistence_StatusARM</a>)
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
<tbody><tr><td><p>&#34;12h&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1h&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;6h&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Persistence_Status">Persistence_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
<code>aofEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AofEnabled: Sets whether AOF is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>aofFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceStatusAofFrequency">
PersistenceStatusAofFrequency
</a>
</em>
</td>
<td>
<p>AofFrequency: Sets the frequency at which data is written to disk.</p>
</td>
</tr>
<tr>
<td>
<code>rdbEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RdbEnabled: Sets whether RDB is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>rdbFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceStatusRdbFrequency">
PersistenceStatusRdbFrequency
</a>
</em>
</td>
<td>
<p>RdbFrequency: Sets the frequency at which a snapshot of the database is created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Persistence_StatusARM">Persistence_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>)
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
<code>aofEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AofEnabled: Sets whether AOF is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>aofFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceStatusAofFrequency">
PersistenceStatusAofFrequency
</a>
</em>
</td>
<td>
<p>AofFrequency: Sets the frequency at which data is written to disk.</p>
</td>
</tr>
<tr>
<td>
<code>rdbEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RdbEnabled: Sets whether RDB is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>rdbFrequency</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.PersistenceStatusRdbFrequency">
PersistenceStatusRdbFrequency
</a>
</em>
</td>
<td>
<p>RdbFrequency: Sets the frequency at which a snapshot of the database is created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Cluster_Status">Cluster_Status</a>)
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
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.ClusterProperties_StatusARM">ClusterProperties_StatusARM</a>)
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
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.ProvisioningState_Status">ProvisioningState_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.ClusterProperties_StatusARM">ClusterProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Cluster_Status">Cluster_Status</a>, <a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.RedisEnterprise">RedisEnterprise
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise</a></p>
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
<a href="#cache.azure.com/v1beta20210301.RedisEnterprise_Spec">
RedisEnterprise_Spec
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ClusterPropertiesMinimumTlsVersion">
ClusterPropertiesMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. &lsquo;1.2&rsquo;.</p>
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
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create RedisEnterprise operation.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The Availability Zones where this cluster will be deployed.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Cluster_Status">
Cluster_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.RedisEnterpriseDatabase">RedisEnterpriseDatabase
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise_databases">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise_databases</a></p>
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
<a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">
RedisEnterpriseDatabases_Spec
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
<code>clientProtocol</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesClientProtocol">
DatabasePropertiesClientProtocol
</a>
</em>
</td>
<td>
<p>ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
TLS-encrypted.</p>
</td>
</tr>
<tr>
<td>
<code>clusteringPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesClusteringPolicy">
DatabasePropertiesClusteringPolicy
</a>
</em>
</td>
<td>
<p>ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesEvictionPolicy">
DatabasePropertiesEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Redis eviction policy - default is VolatileLRU.</p>
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
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>modules</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Module">
[]Module
</a>
</em>
</td>
<td>
<p>Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.</p>
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
reference to a cache.azure.com/RedisEnterprise resource</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Persistence">
Persistence
</a>
</em>
</td>
<td>
<p>Persistence: Persistence-related configuration for the RedisEnterprise database</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.</p>
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
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Database_Status">
Database_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.RedisEnterpriseDatabasesSpecAPIVersion">RedisEnterpriseDatabasesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-03-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_Spec">RedisEnterpriseDatabases_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterpriseDatabase">RedisEnterpriseDatabase</a>)
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
<code>clientProtocol</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesClientProtocol">
DatabasePropertiesClientProtocol
</a>
</em>
</td>
<td>
<p>ClientProtocol: Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is
TLS-encrypted.</p>
</td>
</tr>
<tr>
<td>
<code>clusteringPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesClusteringPolicy">
DatabasePropertiesClusteringPolicy
</a>
</em>
</td>
<td>
<p>ClusteringPolicy: Clustering policy - default is OSSCluster. Specified at create time.</p>
</td>
</tr>
<tr>
<td>
<code>evictionPolicy</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesEvictionPolicy">
DatabasePropertiesEvictionPolicy
</a>
</em>
</td>
<td>
<p>EvictionPolicy: Redis eviction policy - default is VolatileLRU.</p>
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
<p>Location: Location to deploy resource to</p>
</td>
</tr>
<tr>
<td>
<code>modules</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Module">
[]Module
</a>
</em>
</td>
<td>
<p>Modules: Optional set of redis modules to enable in this database - modules can only be added at creation time.</p>
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
reference to a cache.azure.com/RedisEnterprise resource</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Persistence">
Persistence
</a>
</em>
</td>
<td>
<p>Persistence: Persistence-related configuration for the RedisEnterprise database</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: TCP port of the database endpoint. Specified at create time. Defaults to an available port.</p>
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
<h3 id="cache.azure.com/v1beta20210301.RedisEnterpriseDatabases_SpecARM">RedisEnterpriseDatabases_SpecARM
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Location to deploy resource to</p>
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
<p>Name: The name of the database.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.DatabasePropertiesARM">
DatabasePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of RedisEnterprise databases, as opposed to general resource properties like location, tags</p>
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
<h3 id="cache.azure.com/v1beta20210301.RedisEnterpriseSpecAPIVersion">RedisEnterpriseSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-03-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.RedisEnterprise_Spec">RedisEnterprise_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterprise">RedisEnterprise</a>)
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ClusterPropertiesMinimumTlsVersion">
ClusterPropertiesMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. &lsquo;1.2&rsquo;.</p>
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
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create RedisEnterprise operation.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The Availability Zones where this cluster will be deployed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.RedisEnterprise_SpecARM">RedisEnterprise_SpecARM
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
<p>Name: The name of the RedisEnterprise cluster.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.ClusterPropertiesARM">
ClusterPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of RedisEnterprise clusters, as opposed to general resource properties like location, tags</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create RedisEnterprise operation.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The Availability Zones where this cluster will be deployed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.ResourceState_Status">ResourceState_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.ClusterProperties_StatusARM">ClusterProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Cluster_Status">Cluster_Status</a>, <a href="#cache.azure.com/v1beta20210301.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20210301.Database_Status">Database_Status</a>)
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
<tbody><tr><td><p>&#34;CreateFailed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DeleteFailed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DisableFailed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disabling&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnableFailed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabling&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UpdateFailed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterprise_Spec">RedisEnterprise_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Sku</a></p>
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
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, &hellip;)
for Enterprise SKUs and (3, 9, 15, &hellip;) for Flash SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.RedisEnterprise_SpecARM">RedisEnterprise_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Sku</a></p>
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
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, &hellip;)
for Enterprise SKUs and (3, 9, 15, &hellip;) for Flash SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.SkuName">SkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Sku">Sku</a>, <a href="#cache.azure.com/v1beta20210301.SkuARM">SkuARM</a>)
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
<tbody><tr><td><p>&#34;Enterprise_E10&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enterprise_E100&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enterprise_E20&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enterprise_E50&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnterpriseFlash_F1500&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnterpriseFlash_F300&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnterpriseFlash_F700&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.SkuStatusName">SkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Sku_Status">Sku_Status</a>, <a href="#cache.azure.com/v1beta20210301.Sku_StatusARM">Sku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Enterprise_E10&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enterprise_E100&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enterprise_E20&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enterprise_E50&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnterpriseFlash_F1500&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnterpriseFlash_F300&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnterpriseFlash_F700&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Cluster_Status">Cluster_Status</a>)
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
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, &hellip;)
for Enterprise SKUs and (3, 9, 15, &hellip;) for Flash SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20210301.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20210301.Cluster_StatusARM">Cluster_StatusARM</a>)
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
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, &hellip;)
for Enterprise SKUs and (3, 9, 15, &hellip;) for Flash SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20210301.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.)</p>
</td>
</tr>
</tbody>
</table>
<hr/>
