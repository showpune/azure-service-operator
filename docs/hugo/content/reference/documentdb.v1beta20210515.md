---
---
<h2 id="documentdb.azure.com/v1beta20210515">documentdb.azure.com/v1beta20210515</h2>
<div>
<p>Package v1beta20210515 contains API Schema definitions for the documentdb v1beta20210515 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration">AnalyticalStorageConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AnalyticalStorageConfiguration">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AnalyticalStorageConfiguration</a></p>
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
<code>schemaType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfigurationSchemaType">
AnalyticalStorageConfigurationSchemaType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AnalyticalStorageConfigurationARM">AnalyticalStorageConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AnalyticalStorageConfiguration">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AnalyticalStorageConfiguration</a></p>
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
<code>schemaType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfigurationSchemaType">
AnalyticalStorageConfigurationSchemaType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AnalyticalStorageConfigurationSchemaType">AnalyticalStorageConfigurationSchemaType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration">AnalyticalStorageConfiguration</a>, <a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfigurationARM">AnalyticalStorageConfigurationARM</a>)
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
<tbody><tr><td><p>&#34;FullFidelity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;WellDefined&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration_Status">AnalyticalStorageConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>schemaType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageSchemaType_Status">
AnalyticalStorageSchemaType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration_StatusARM">AnalyticalStorageConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>schemaType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageSchemaType_Status">
AnalyticalStorageSchemaType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AnalyticalStorageSchemaType_Status">AnalyticalStorageSchemaType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration_Status">AnalyticalStorageConfiguration_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration_StatusARM">AnalyticalStorageConfiguration_StatusARM</a>)
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
<tbody><tr><td><p>&#34;FullFidelity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;WellDefined&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ApiProperties">ApiProperties
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ApiProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ApiProperties</a></p>
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
<code>serverVersion</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiPropertiesServerVersion">
ApiPropertiesServerVersion
</a>
</em>
</td>
<td>
<p>ServerVersion: Describes the ServerVersion of an a MongoDB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ApiPropertiesARM">ApiPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ApiProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ApiProperties</a></p>
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
<code>serverVersion</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiPropertiesServerVersion">
ApiPropertiesServerVersion
</a>
</em>
</td>
<td>
<p>ServerVersion: Describes the ServerVersion of an a MongoDB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ApiPropertiesServerVersion">ApiPropertiesServerVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ApiProperties">ApiProperties</a>, <a href="#documentdb.azure.com/v1beta20210515.ApiPropertiesARM">ApiPropertiesARM</a>)
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
<tbody><tr><td><p>&#34;3.2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;3.6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;4.0&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ApiPropertiesStatusServerVersion">ApiPropertiesStatusServerVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ApiProperties_Status">ApiProperties_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.ApiProperties_StatusARM">ApiProperties_StatusARM</a>)
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
<tbody><tr><td><p>&#34;3.2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;3.6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;4.0&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ApiProperties_Status">ApiProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>serverVersion</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiPropertiesStatusServerVersion">
ApiPropertiesStatusServerVersion
</a>
</em>
</td>
<td>
<p>ServerVersion: Describes the ServerVersion of an a MongoDB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ApiProperties_StatusARM">ApiProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>serverVersion</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiPropertiesStatusServerVersion">
ApiPropertiesStatusServerVersion
</a>
</em>
</td>
<td>
<p>ServerVersion: Describes the ServerVersion of an a MongoDB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource">AutoUpgradePolicyResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource">AutoscaleSettingsResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoUpgradePolicyResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoUpgradePolicyResource</a></p>
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
<code>throughputPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputPolicyResource">
ThroughputPolicyResource
</a>
</em>
</td>
<td>
<p>ThroughputPolicy: Cosmos DB resource throughput policy</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResourceARM">AutoUpgradePolicyResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResourceARM">AutoscaleSettingsResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoUpgradePolicyResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoUpgradePolicyResource</a></p>
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
<code>throughputPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputPolicyResourceARM">
ThroughputPolicyResourceARM
</a>
</em>
</td>
<td>
<p>ThroughputPolicy: Cosmos DB resource throughput policy</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource_Status">AutoUpgradePolicyResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource_Status">AutoscaleSettingsResource_Status</a>)
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
<code>throughputPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputPolicyResource_Status">
ThroughputPolicyResource_Status
</a>
</em>
</td>
<td>
<p>ThroughputPolicy: Represents throughput policy which service must adhere to for auto-upgrade</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource_StatusARM">AutoUpgradePolicyResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource_StatusARM">AutoscaleSettingsResource_StatusARM</a>)
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
<code>throughputPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputPolicyResource_StatusARM">
ThroughputPolicyResource_StatusARM
</a>
</em>
</td>
<td>
<p>ThroughputPolicy: Represents throughput policy which service must adhere to for auto-upgrade</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettings">AutoscaleSettings
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">CreateUpdateOptions</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettings">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettings</a></p>
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
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput, the resource can scale up to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettingsARM">AutoscaleSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">CreateUpdateOptionsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettings">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettings</a></p>
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
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput, the resource can scale up to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource">AutoscaleSettingsResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">ThroughputSettingsResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettingsResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettingsResource</a></p>
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
<code>autoUpgradePolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource">
AutoUpgradePolicyResource
</a>
</em>
</td>
<td>
<p>AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy</p>
</td>
</tr>
<tr>
<td>
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput container can scale up to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettingsResourceARM">AutoscaleSettingsResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResourceARM">ThroughputSettingsResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettingsResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettingsResource</a></p>
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
<code>autoUpgradePolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResourceARM">
AutoUpgradePolicyResourceARM
</a>
</em>
</td>
<td>
<p>AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy</p>
</td>
</tr>
<tr>
<td>
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput container can scale up to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource_Status">AutoscaleSettingsResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_Status_Resource">ThroughputSettingsGetProperties_Status_Resource</a>)
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
<code>autoUpgradePolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource_Status">
AutoUpgradePolicyResource_Status
</a>
</em>
</td>
<td>
<p>AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy</p>
</td>
</tr>
<tr>
<td>
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput container can scale up to.</p>
</td>
</tr>
<tr>
<td>
<code>targetMaxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>TargetMaxThroughput: Represents target maximum throughput container can scale up to once offer is no longer in pending
state.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource_StatusARM">AutoscaleSettingsResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_Status_ResourceARM">ThroughputSettingsGetProperties_Status_ResourceARM</a>)
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
<code>autoUpgradePolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource_StatusARM">
AutoUpgradePolicyResource_StatusARM
</a>
</em>
</td>
<td>
<p>AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy</p>
</td>
</tr>
<tr>
<td>
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput container can scale up to.</p>
</td>
</tr>
<tr>
<td>
<code>targetMaxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>TargetMaxThroughput: Represents target maximum throughput container can scale up to once offer is no longer in pending
state.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettings_Status">AutoscaleSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.OptionsResource_Status">OptionsResource_Status</a>)
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
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput, the resource can scale up to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.AutoscaleSettings_StatusARM">AutoscaleSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.OptionsResource_StatusARM">OptionsResource_StatusARM</a>)
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
<code>maxThroughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxThroughput: Represents maximum throughput, the resource can scale up to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.BackupPolicy">BackupPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/BackupPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/BackupPolicy</a></p>
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
<code>continuousModeBackupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicy">
ContinuousModeBackupPolicy
</a>
</em>
</td>
<td>
<p>Continuous: Mutually exclusive with all other properties</p>
</td>
</tr>
<tr>
<td>
<code>periodicModeBackupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicy">
PeriodicModeBackupPolicy
</a>
</em>
</td>
<td>
<p>Periodic: Mutually exclusive with all other properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.BackupPolicyARM">BackupPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/BackupPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/BackupPolicy</a></p>
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
<code>continuousModeBackupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicyARM">
ContinuousModeBackupPolicyARM
</a>
</em>
</td>
<td>
<p>Continuous: Mutually exclusive with all other properties</p>
</td>
</tr>
<tr>
<td>
<code>periodicModeBackupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyARM">
PeriodicModeBackupPolicyARM
</a>
</em>
</td>
<td>
<p>Periodic: Mutually exclusive with all other properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.BackupPolicyType_Status">BackupPolicyType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.BackupPolicy_Status">BackupPolicy_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.BackupPolicy_StatusARM">BackupPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Continuous&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Periodic&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.BackupPolicy_Status">BackupPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicyType_Status">
BackupPolicyType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.BackupPolicy_StatusARM">BackupPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicyType_Status">
BackupPolicyType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Capability">Capability
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Capability">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Capability</a></p>
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the Cosmos DB capability. For example, &ldquo;name&rdquo;: &ldquo;EnableCassandra&rdquo;. Current values also include
&ldquo;EnableTable&rdquo; and &ldquo;EnableGremlin&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CapabilityARM">CapabilityARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Capability">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Capability</a></p>
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the Cosmos DB capability. For example, &ldquo;name&rdquo;: &ldquo;EnableCassandra&rdquo;. Current values also include
&ldquo;EnableTable&rdquo; and &ldquo;EnableGremlin&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Capability_Status">Capability_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the Cosmos DB capability. For example, &ldquo;name&rdquo;: &ldquo;EnableCassandra&rdquo;. Current values also include
&ldquo;EnableTable&rdquo; and &ldquo;EnableGremlin&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Capability_StatusARM">Capability_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the Cosmos DB capability. For example, &ldquo;name&rdquo;: &ldquo;EnableCassandra&rdquo;. Current values also include
&ldquo;EnableTable&rdquo; and &ldquo;EnableGremlin&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CompositePath">CompositePath
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy">IndexingPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath</a></p>
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
<code>order</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePathOrder">
CompositePathOrder
</a>
</em>
</td>
<td>
<p>Order: Sort order for composite paths.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CompositePathARM">CompositePathARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyARM">IndexingPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath</a></p>
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
<code>order</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePathOrder">
CompositePathOrder
</a>
</em>
</td>
<td>
<p>Order: Sort order for composite paths.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CompositePathOrder">CompositePathOrder
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.CompositePath">CompositePath</a>, <a href="#documentdb.azure.com/v1beta20210515.CompositePathARM">CompositePathARM</a>)
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
<tbody><tr><td><p>&#34;ascending&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;descending&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CompositePathStatusOrder">CompositePathStatusOrder
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.CompositePath_Status">CompositePath_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.CompositePath_StatusARM">CompositePath_StatusARM</a>)
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
<tbody><tr><td><p>&#34;ascending&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;descending&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CompositePath_Status">CompositePath_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">IndexingPolicy_Status</a>)
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
<code>order</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePathStatusOrder">
CompositePathStatusOrder
</a>
</em>
</td>
<td>
<p>Order: Sort order for composite paths.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CompositePath_StatusARM">CompositePath_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">IndexingPolicy_StatusARM</a>)
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
<code>order</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePathStatusOrder">
CompositePathStatusOrder
</a>
</em>
</td>
<td>
<p>Order: Sort order for composite paths.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy">ConflictResolutionPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResource">SqlContainerResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy</a></p>
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
<code>conflictResolutionPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionProcedure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyMode">
ConflictResolutionPolicyMode
</a>
</em>
</td>
<td>
<p>Mode: Indicates the conflict resolution mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyARM">ConflictResolutionPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResourceARM">SqlContainerResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy</a></p>
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
<code>conflictResolutionPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionProcedure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyMode">
ConflictResolutionPolicyMode
</a>
</em>
</td>
<td>
<p>Mode: Indicates the conflict resolution mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyMode">ConflictResolutionPolicyMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy">ConflictResolutionPolicy</a>, <a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyARM">ConflictResolutionPolicyARM</a>)
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
<tbody><tr><td><p>&#34;Custom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LastWriterWins&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyStatusMode">ConflictResolutionPolicyStatusMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy_Status">ConflictResolutionPolicy_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy_StatusARM">ConflictResolutionPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Custom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LastWriterWins&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy_Status">ConflictResolutionPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_Resource">SqlContainerGetProperties_Status_Resource</a>)
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
<code>conflictResolutionPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionProcedure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyStatusMode">
ConflictResolutionPolicyStatusMode
</a>
</em>
</td>
<td>
<p>Mode: Indicates the conflict resolution mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy_StatusARM">ConflictResolutionPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_ResourceARM">SqlContainerGetProperties_Status_ResourceARM</a>)
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
<code>conflictResolutionPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionProcedure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyStatusMode">
ConflictResolutionPolicyStatusMode
</a>
</em>
</td>
<td>
<p>Mode: Indicates the conflict resolution mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConnectorOffer_Status">ConnectorOffer_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<tbody><tr><td><p>&#34;Small&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConsistencyPolicy">ConsistencyPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConsistencyPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConsistencyPolicy</a></p>
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
<code>defaultConsistencyLevel</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicyDefaultConsistencyLevel">
ConsistencyPolicyDefaultConsistencyLevel
</a>
</em>
</td>
<td>
<p>DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>maxIntervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
set to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>maxStalenessPrefix</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConsistencyPolicyARM">ConsistencyPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConsistencyPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConsistencyPolicy</a></p>
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
<code>defaultConsistencyLevel</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicyDefaultConsistencyLevel">
ConsistencyPolicyDefaultConsistencyLevel
</a>
</em>
</td>
<td>
<p>DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>maxIntervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
set to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>maxStalenessPrefix</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConsistencyPolicyDefaultConsistencyLevel">ConsistencyPolicyDefaultConsistencyLevel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy">ConsistencyPolicy</a>, <a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicyARM">ConsistencyPolicyARM</a>)
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
<tbody><tr><td><p>&#34;BoundedStaleness&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ConsistentPrefix&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Eventual&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Session&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Strong&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConsistencyPolicyStatusDefaultConsistencyLevel">ConsistencyPolicyStatusDefaultConsistencyLevel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy_Status">ConsistencyPolicy_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy_StatusARM">ConsistencyPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;BoundedStaleness&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ConsistentPrefix&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Eventual&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Session&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Strong&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConsistencyPolicy_Status">ConsistencyPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>defaultConsistencyLevel</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicyStatusDefaultConsistencyLevel">
ConsistencyPolicyStatusDefaultConsistencyLevel
</a>
</em>
</td>
<td>
<p>DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>maxIntervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
set to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>maxStalenessPrefix</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ConsistencyPolicy_StatusARM">ConsistencyPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>defaultConsistencyLevel</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicyStatusDefaultConsistencyLevel">
ConsistencyPolicyStatusDefaultConsistencyLevel
</a>
</em>
</td>
<td>
<p>DefaultConsistencyLevel: The default consistency level and configuration settings of the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>maxIntervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxIntervalInSeconds: When used with the Bounded Staleness consistency level, this value represents the time amount of
staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is
set to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>maxStalenessPrefix</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxStalenessPrefix: When used with the Bounded Staleness consistency level, this value represents the number of stale
requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set
to &lsquo;BoundedStaleness&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContainerPartitionKey">ContainerPartitionKey
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResource">SqlContainerResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey</a></p>
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
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKeyKind">
ContainerPartitionKeyKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
are supported for container create.</p>
</td>
</tr>
<tr>
<td>
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths using which data within the container can be partitioned</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
int
</em>
</td>
<td>
<p>Version: Indicates the version of the partition key definition</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContainerPartitionKeyARM">ContainerPartitionKeyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResourceARM">SqlContainerResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey</a></p>
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
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKeyKind">
ContainerPartitionKeyKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
are supported for container create.</p>
</td>
</tr>
<tr>
<td>
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths using which data within the container can be partitioned</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
int
</em>
</td>
<td>
<p>Version: Indicates the version of the partition key definition</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContainerPartitionKeyKind">ContainerPartitionKeyKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKey">ContainerPartitionKey</a>, <a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKeyARM">ContainerPartitionKeyARM</a>)
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
<tbody><tr><td><p>&#34;Hash&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MultiHash&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Range&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContainerPartitionKeyStatusKind">ContainerPartitionKeyStatusKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKey_Status">ContainerPartitionKey_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKey_StatusARM">ContainerPartitionKey_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Hash&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MultiHash&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Range&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContainerPartitionKey_Status">ContainerPartitionKey_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_Resource">SqlContainerGetProperties_Status_Resource</a>)
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
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKeyStatusKind">
ContainerPartitionKeyStatusKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
are supported for container create</p>
</td>
</tr>
<tr>
<td>
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths using which data within the container can be partitioned</p>
</td>
</tr>
<tr>
<td>
<code>systemKey</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SystemKey: Indicates if the container is using a system generated partition key</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
int
</em>
</td>
<td>
<p>Version: Indicates the version of the partition key definition</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContainerPartitionKey_StatusARM">ContainerPartitionKey_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_ResourceARM">SqlContainerGetProperties_Status_ResourceARM</a>)
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
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKeyStatusKind">
ContainerPartitionKeyStatusKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
are supported for container create</p>
</td>
</tr>
<tr>
<td>
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths using which data within the container can be partitioned</p>
</td>
</tr>
<tr>
<td>
<code>systemKey</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SystemKey: Indicates if the container is using a system generated partition key</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
int
</em>
</td>
<td>
<p>Version: Indicates the version of the partition key definition</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicy">ContinuousModeBackupPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.BackupPolicy">BackupPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContinuousModeBackupPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContinuousModeBackupPolicy</a></p>
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
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicyType">
ContinuousModeBackupPolicyType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicyARM">ContinuousModeBackupPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.BackupPolicyARM">BackupPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContinuousModeBackupPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContinuousModeBackupPolicy</a></p>
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
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicyType">
ContinuousModeBackupPolicyType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicyType">ContinuousModeBackupPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicy">ContinuousModeBackupPolicy</a>, <a href="#documentdb.azure.com/v1beta20210515.ContinuousModeBackupPolicyARM">ContinuousModeBackupPolicyARM</a>)
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
<tbody><tr><td><p>&#34;Continuous&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CorsPolicy">CorsPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CorsPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CorsPolicy</a></p>
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
<code>allowedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedHeaders: The request headers that the origin domain may specify on the CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedMethods</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedOrigins</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.</p>
</td>
</tr>
<tr>
<td>
<code>exposedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
the request issuer.</p>
</td>
</tr>
<tr>
<td>
<code>maxAgeInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CorsPolicyARM">CorsPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CorsPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CorsPolicy</a></p>
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
<code>allowedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedHeaders: The request headers that the origin domain may specify on the CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedMethods</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedOrigins</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.</p>
</td>
</tr>
<tr>
<td>
<code>exposedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
the request issuer.</p>
</td>
</tr>
<tr>
<td>
<code>maxAgeInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CorsPolicy_Status">CorsPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>allowedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedHeaders: The request headers that the origin domain may specify on the CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedMethods</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedOrigins</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.</p>
</td>
</tr>
<tr>
<td>
<code>exposedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
the request issuer.</p>
</td>
</tr>
<tr>
<td>
<code>maxAgeInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CorsPolicy_StatusARM">CorsPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>allowedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedHeaders: The request headers that the origin domain may specify on the CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedMethods</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedMethods: The methods (HTTP request verbs) that the origin domain may use for a CORS request.</p>
</td>
</tr>
<tr>
<td>
<code>allowedOrigins</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedOrigins: The origin domains that are permitted to make a request against the service via CORS.</p>
</td>
</tr>
<tr>
<td>
<code>exposedHeaders</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExposedHeaders: The response headers that may be sent in the response to the CORS request and exposed by the browser to
the request issuer.</p>
</td>
</tr>
<tr>
<td>
<code>maxAgeInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxAgeInSeconds: The maximum amount time that a browser should cache the preflight OPTIONS request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CreateUpdateOptions">CreateUpdateOptions
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollections_Spec">DatabaseAccountsMongodbDatabasesCollections_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabases_Spec">DatabaseAccountsMongodbDatabases_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec">DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec">DatabaseAccountsSqlDatabasesContainersTriggers_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec">DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainers_Spec">DatabaseAccountsSqlDatabasesContainers_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabases_Spec">DatabaseAccountsSqlDatabases_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CreateUpdateOptions">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CreateUpdateOptions</a></p>
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettings">
AutoscaleSettings
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Request Units per second. For example, &ldquo;throughput&rdquo;: 10000.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">CreateUpdateOptionsARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionCreateUpdatePropertiesARM">MongoDBCollectionCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseCreateUpdatePropertiesARM">MongoDBDatabaseCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlContainerCreateUpdatePropertiesARM">SqlContainerCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseCreateUpdatePropertiesARM">SqlDatabaseCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureCreateUpdatePropertiesARM">SqlStoredProcedureCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlTriggerCreateUpdatePropertiesARM">SqlTriggerCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionCreateUpdatePropertiesARM">SqlUserDefinedFunctionCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CreateUpdateOptions">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CreateUpdateOptions</a></p>
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsARM">
AutoscaleSettingsARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Request Units per second. For example, &ldquo;throughput&rdquo;: 10000.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccount">DatabaseAccount
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">
DatabaseAccounts_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>analyticalStorageConfiguration</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration">
AnalyticalStorageConfiguration
</a>
</em>
</td>
<td>
<p>AnalyticalStorageConfiguration: Analytical storage specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiProperties">
ApiProperties
</a>
</em>
</td>
<td>
</td>
</tr>
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
<code>backupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicy">
BackupPolicy
</a>
</em>
</td>
<td>
<p>BackupPolicy: The object representing the policy for taking backups on an account.</p>
</td>
</tr>
<tr>
<td>
<code>capabilities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Capability">
[]Capability
</a>
</em>
</td>
<td>
<p>Capabilities: List of Cosmos DB capabilities for the account</p>
</td>
</tr>
<tr>
<td>
<code>connectorOffer</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesConnectorOffer">
DatabaseAccountCreateUpdatePropertiesConnectorOffer
</a>
</em>
</td>
<td>
<p>ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.</p>
</td>
</tr>
<tr>
<td>
<code>consistencyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy">
ConsistencyPolicy
</a>
</em>
</td>
<td>
<p>ConsistencyPolicy: The consistency policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>cors</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CorsPolicy">
[]CorsPolicy
</a>
</em>
</td>
<td>
<p>Cors: The CORS policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>databaseAccountOfferType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType">
DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType
</a>
</em>
</td>
<td>
<p>DatabaseAccountOfferType: The offer type for the database</p>
</td>
</tr>
<tr>
<td>
<code>defaultIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
identity needs to be explicitly set by the users. It can be &ldquo;FirstPartyIdentity&rdquo;, &ldquo;SystemAssignedIdentity&rdquo; and more.</p>
</td>
</tr>
<tr>
<td>
<code>disableKeyBasedMetadataWriteAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
via account keys</p>
</td>
</tr>
<tr>
<td>
<code>enableAnalyticalStorage</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticFailover</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
failover priorities configured for the account.</p>
</td>
</tr>
<tr>
<td>
<code>enableCassandraConnector</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account</p>
</td>
</tr>
<tr>
<td>
<code>enableFreeTier</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFreeTier: Flag to indicate whether Free Tier is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleWriteLocations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleWriteLocations: Enables the account to write in multiple locations</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity">
ManagedServiceIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IpAddressOrRange">
[]IpAddressOrRange
</a>
</em>
</td>
<td>
<p>IpRules: Array of IpAddressOrRange objects.</p>
</td>
</tr>
<tr>
<td>
<code>isVirtualNetworkFilterEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultKeyUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultKeyUri: The URI of the key vault</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSpecKind">
DatabaseAccountsSpecKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of database account. This can only be set at database account creation.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location">
[]Location
</a>
</em>
</td>
<td>
<p>Locations: An array that contains the georeplication locations enabled for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypass</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesNetworkAclBypass">
DatabaseAccountCreateUpdatePropertiesNetworkAclBypass
</a>
</em>
</td>
<td>
<p>NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypassResourceIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountOperatorSpec">
DatabaseAccountOperatorSpec
</a>
</em>
</td>
<td>
<p>OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
passed directly to Azure</p>
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
<code>publicNetworkAccess</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess">
DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether requests from Public Network are allowed.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.VirtualNetworkRule">
[]VirtualNetworkRule
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">
DatabaseAccountGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_SpecARM">DatabaseAccounts_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/DatabaseAccountCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/DatabaseAccountCreateUpdateProperties</a></p>
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
<code>analyticalStorageConfiguration</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfigurationARM">
AnalyticalStorageConfigurationARM
</a>
</em>
</td>
<td>
<p>AnalyticalStorageConfiguration: Analytical storage specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiPropertiesARM">
ApiPropertiesARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>backupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicyARM">
BackupPolicyARM
</a>
</em>
</td>
<td>
<p>BackupPolicy: The object representing the policy for taking backups on an account.</p>
</td>
</tr>
<tr>
<td>
<code>capabilities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CapabilityARM">
[]CapabilityARM
</a>
</em>
</td>
<td>
<p>Capabilities: List of Cosmos DB capabilities for the account</p>
</td>
</tr>
<tr>
<td>
<code>connectorOffer</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesConnectorOffer">
DatabaseAccountCreateUpdatePropertiesConnectorOffer
</a>
</em>
</td>
<td>
<p>ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.</p>
</td>
</tr>
<tr>
<td>
<code>consistencyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicyARM">
ConsistencyPolicyARM
</a>
</em>
</td>
<td>
<p>ConsistencyPolicy: The consistency policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>cors</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CorsPolicyARM">
[]CorsPolicyARM
</a>
</em>
</td>
<td>
<p>Cors: The CORS policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>databaseAccountOfferType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType">
DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType
</a>
</em>
</td>
<td>
<p>DatabaseAccountOfferType: The offer type for the database</p>
</td>
</tr>
<tr>
<td>
<code>defaultIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
identity needs to be explicitly set by the users. It can be &ldquo;FirstPartyIdentity&rdquo;, &ldquo;SystemAssignedIdentity&rdquo; and more.</p>
</td>
</tr>
<tr>
<td>
<code>disableKeyBasedMetadataWriteAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
via account keys</p>
</td>
</tr>
<tr>
<td>
<code>enableAnalyticalStorage</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticFailover</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
failover priorities configured for the account.</p>
</td>
</tr>
<tr>
<td>
<code>enableCassandraConnector</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account</p>
</td>
</tr>
<tr>
<td>
<code>enableFreeTier</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFreeTier: Flag to indicate whether Free Tier is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleWriteLocations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleWriteLocations: Enables the account to write in multiple locations</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IpAddressOrRangeARM">
[]IpAddressOrRangeARM
</a>
</em>
</td>
<td>
<p>IpRules: Array of IpAddressOrRange objects.</p>
</td>
</tr>
<tr>
<td>
<code>isVirtualNetworkFilterEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultKeyUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultKeyUri: The URI of the key vault</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.LocationARM">
[]LocationARM
</a>
</em>
</td>
<td>
<p>Locations: An array that contains the georeplication locations enabled for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypass</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesNetworkAclBypass">
DatabaseAccountCreateUpdatePropertiesNetworkAclBypass
</a>
</em>
</td>
<td>
<p>NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypassResourceIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess">
DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether requests from Public Network are allowed.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.VirtualNetworkRuleARM">
[]VirtualNetworkRuleARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesConnectorOffer">DatabaseAccountCreateUpdatePropertiesConnectorOffer
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
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
<tbody><tr><td><p>&#34;Small&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType">DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
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
<tbody><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesNetworkAclBypass">DatabaseAccountCreateUpdatePropertiesNetworkAclBypass
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
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
<tbody><tr><td><p>&#34;AzureServices&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess">DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
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
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_StatusARM">DatabaseAccountGetResults_StatusARM</a>)
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
<code>analyticalStorageConfiguration</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration_StatusARM">
AnalyticalStorageConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>AnalyticalStorageConfiguration: Analytical storage specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiProperties_StatusARM">
ApiProperties_StatusARM
</a>
</em>
</td>
<td>
<p>ApiProperties: API specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>backupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicy_StatusARM">
BackupPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>BackupPolicy: The object representing the policy for taking backups on an account.</p>
</td>
</tr>
<tr>
<td>
<code>capabilities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Capability_StatusARM">
[]Capability_StatusARM
</a>
</em>
</td>
<td>
<p>Capabilities: List of Cosmos DB capabilities for the account</p>
</td>
</tr>
<tr>
<td>
<code>connectorOffer</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConnectorOffer_Status">
ConnectorOffer_Status
</a>
</em>
</td>
<td>
<p>ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.</p>
</td>
</tr>
<tr>
<td>
<code>consistencyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy_StatusARM">
ConsistencyPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>ConsistencyPolicy: The consistency policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>cors</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CorsPolicy_StatusARM">
[]CorsPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>Cors: The CORS policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>databaseAccountOfferType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountOfferType_Status">
DatabaseAccountOfferType_Status
</a>
</em>
</td>
<td>
<p>DatabaseAccountOfferType: The offer type for the Cosmos DB database account. Default value: Standard.</p>
</td>
</tr>
<tr>
<td>
<code>defaultIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
identity needs to be explicitly set by the users. It can be &ldquo;FirstPartyIdentity&rdquo;, &ldquo;SystemAssignedIdentity&rdquo; and more.</p>
</td>
</tr>
<tr>
<td>
<code>disableKeyBasedMetadataWriteAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
via account keys</p>
</td>
</tr>
<tr>
<td>
<code>documentEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>DocumentEndpoint: The connection endpoint for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>enableAnalyticalStorage</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticFailover</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
failover priorities configured for the account.</p>
</td>
</tr>
<tr>
<td>
<code>enableCassandraConnector</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account</p>
</td>
</tr>
<tr>
<td>
<code>enableFreeTier</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFreeTier: Flag to indicate whether Free Tier is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleWriteLocations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleWriteLocations: Enables the account to write in multiple locations</p>
</td>
</tr>
<tr>
<td>
<code>failoverPolicies</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.FailoverPolicy_StatusARM">
[]FailoverPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>FailoverPolicies: An array that contains the regions ordered by their failover priorities.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IpAddressOrRange_StatusARM">
[]IpAddressOrRange_StatusARM
</a>
</em>
</td>
<td>
<p>IpRules: List of IpRules.</p>
</td>
</tr>
<tr>
<td>
<code>isVirtualNetworkFilterEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultKeyUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultKeyUri: The URI of the key vault</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location_StatusARM">
[]Location_StatusARM
</a>
</em>
</td>
<td>
<p>Locations: An array that contains all of the locations enabled for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypass</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.NetworkAclBypass_Status">
NetworkAclBypass_Status
</a>
</em>
</td>
<td>
<p>NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypassResourceIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of Private Endpoint Connections configured for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PublicNetworkAccess_Status">
PublicNetworkAccess_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether requests from Public Network are allowed</p>
</td>
</tr>
<tr>
<td>
<code>readLocations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location_StatusARM">
[]Location_StatusARM
</a>
</em>
</td>
<td>
<p>ReadLocations: An array that contains of the read locations enabled for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.VirtualNetworkRule_StatusARM">
[]VirtualNetworkRule_StatusARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>writeLocations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location_StatusARM">
[]Location_StatusARM
</a>
</em>
</td>
<td>
<p>WriteLocations: An array that contains the write location for the Cosmos DB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountGetResultsStatusKind">DatabaseAccountGetResultsStatusKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_StatusARM">DatabaseAccountGetResults_StatusARM</a>)
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
<tbody><tr><td><p>&#34;GlobalDocumentDB&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MongoDB&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Parse&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccount">DatabaseAccount</a>)
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
<code>analyticalStorageConfiguration</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration_Status">
AnalyticalStorageConfiguration_Status
</a>
</em>
</td>
<td>
<p>AnalyticalStorageConfiguration: Analytical storage specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiProperties_Status">
ApiProperties_Status
</a>
</em>
</td>
<td>
<p>ApiProperties: API specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>backupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicy_Status">
BackupPolicy_Status
</a>
</em>
</td>
<td>
<p>BackupPolicy: The object representing the policy for taking backups on an account.</p>
</td>
</tr>
<tr>
<td>
<code>capabilities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Capability_Status">
[]Capability_Status
</a>
</em>
</td>
<td>
<p>Capabilities: List of Cosmos DB capabilities for the account</p>
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
<code>connectorOffer</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConnectorOffer_Status">
ConnectorOffer_Status
</a>
</em>
</td>
<td>
<p>ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.</p>
</td>
</tr>
<tr>
<td>
<code>consistencyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy_Status">
ConsistencyPolicy_Status
</a>
</em>
</td>
<td>
<p>ConsistencyPolicy: The consistency policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>cors</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CorsPolicy_Status">
[]CorsPolicy_Status
</a>
</em>
</td>
<td>
<p>Cors: The CORS policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>databaseAccountOfferType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountOfferType_Status">
DatabaseAccountOfferType_Status
</a>
</em>
</td>
<td>
<p>DatabaseAccountOfferType: The offer type for the Cosmos DB database account. Default value: Standard.</p>
</td>
</tr>
<tr>
<td>
<code>defaultIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
identity needs to be explicitly set by the users. It can be &ldquo;FirstPartyIdentity&rdquo;, &ldquo;SystemAssignedIdentity&rdquo; and more.</p>
</td>
</tr>
<tr>
<td>
<code>disableKeyBasedMetadataWriteAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
via account keys</p>
</td>
</tr>
<tr>
<td>
<code>documentEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>DocumentEndpoint: The connection endpoint for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>enableAnalyticalStorage</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticFailover</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
failover priorities configured for the account.</p>
</td>
</tr>
<tr>
<td>
<code>enableCassandraConnector</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account</p>
</td>
</tr>
<tr>
<td>
<code>enableFreeTier</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFreeTier: Flag to indicate whether Free Tier is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleWriteLocations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleWriteLocations: Enables the account to write in multiple locations</p>
</td>
</tr>
<tr>
<td>
<code>failoverPolicies</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.FailoverPolicy_Status">
[]FailoverPolicy_Status
</a>
</em>
</td>
<td>
<p>FailoverPolicies: An array that contains the regions ordered by their failover priorities.</p>
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
<p>Id: The unique resource identifier of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status">
ManagedServiceIdentity_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IpAddressOrRange_Status">
[]IpAddressOrRange_Status
</a>
</em>
</td>
<td>
<p>IpRules: List of IpRules.</p>
</td>
</tr>
<tr>
<td>
<code>isVirtualNetworkFilterEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultKeyUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultKeyUri: The URI of the key vault</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResultsStatusKind">
DatabaseAccountGetResultsStatusKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of database account. This can only be set at database account creation.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location_Status">
[]Location_Status
</a>
</em>
</td>
<td>
<p>Locations: An array that contains all of the locations enabled for the Cosmos DB account.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypass</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.NetworkAclBypass_Status">
NetworkAclBypass_Status
</a>
</em>
</td>
<td>
<p>NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypassResourceIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PrivateEndpointConnection_Status_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of Private Endpoint Connections configured for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PublicNetworkAccess_Status">
PublicNetworkAccess_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether requests from Public Network are allowed</p>
</td>
</tr>
<tr>
<td>
<code>readLocations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location_Status">
[]Location_Status
</a>
</em>
</td>
<td>
<p>ReadLocations: An array that contains of the read locations enabled for the Cosmos DB account.</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.VirtualNetworkRule_Status">
[]VirtualNetworkRule_Status
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>writeLocations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location_Status">
[]Location_Status
</a>
</em>
</td>
<td>
<p>WriteLocations: An array that contains the write location for the Cosmos DB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_StatusARM">DatabaseAccountGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_StatusARM">
ManagedServiceIdentity_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResultsStatusKind">
DatabaseAccountGetResultsStatusKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of database account. This can only be set at database account creation.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">
DatabaseAccountGetProperties_StatusARM
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountOfferType_Status">DatabaseAccountOfferType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<tbody><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountOperatorSecrets">DatabaseAccountOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountOperatorSpec">DatabaseAccountOperatorSpec</a>)
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
<code>documentEndpoint</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>DocumentEndpoint: indicates where the DocumentEndpoint secret should be placed. If omitted, the secret will not be
retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>primaryMasterKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>PrimaryMasterKey: indicates where the PrimaryMasterKey secret should be placed. If omitted, the secret will not be
retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>primaryReadonlyMasterKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>PrimaryReadonlyMasterKey: indicates where the PrimaryReadonlyMasterKey secret should be placed. If omitted, the secret
will not be retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryMasterKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>SecondaryMasterKey: indicates where the SecondaryMasterKey secret should be placed. If omitted, the secret will not be
retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryReadonlyMasterKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>SecondaryReadonlyMasterKey: indicates where the SecondaryReadonlyMasterKey secret should be placed. If omitted, the
secret will not be retrieved from Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountOperatorSpec">DatabaseAccountOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure</p>
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
<code>secrets</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountOperatorSecrets">
DatabaseAccountOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsSpecAPIVersion">DatabaseAccountsMongodbDatabasesCollectionsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecAPIVersion">DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec">DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabaseCollectionThroughputSetting">MongodbDatabaseCollectionThroughputSetting</a>)
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/MongodbDatabaseCollection resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM">DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsUpdatePropertiesARM">
ThroughputSettingsUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to update Azure Cosmos DB resource throughput.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollections_Spec">DatabaseAccountsMongodbDatabasesCollections_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabaseCollection">MongodbDatabaseCollection</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/MongodbDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionResource">
MongoDBCollectionResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB MongoDB collection resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollections_SpecARM">DatabaseAccountsMongodbDatabasesCollections_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB collection name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionCreateUpdatePropertiesARM">
MongoDBCollectionCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB MongoDB collection.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesSpecAPIVersion">DatabaseAccountsMongodbDatabasesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion">DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec">DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabaseThroughputSetting">MongodbDatabaseThroughputSetting</a>)
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/MongodbDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM">DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsUpdatePropertiesARM">
ThroughputSettingsUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to update Azure Cosmos DB resource throughput.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabases_Spec">DatabaseAccountsMongodbDatabases_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabase">MongodbDatabase</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/DatabaseAccount resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseResource">
MongoDBDatabaseResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB MongoDB database resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabases_SpecARM">DatabaseAccountsMongodbDatabases_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB database name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseCreateUpdatePropertiesARM">
MongoDBDatabaseCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB MongoDB database.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSpecAPIVersion">DatabaseAccountsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSpecKind">DatabaseAccountsSpecKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_SpecARM">DatabaseAccounts_SpecARM</a>)
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
<tbody><tr><td><p>&#34;GlobalDocumentDB&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MongoDB&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Parse&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersSpecAPIVersion">DatabaseAccountsSqlDatabasesContainersSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecAPIVersion">DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec">DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerStoredProcedure">SqlDatabaseContainerStoredProcedure</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureResource">
SqlStoredProcedureResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL storedProcedure resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM">DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB storedProcedure name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureCreateUpdatePropertiesARM">
SqlStoredProcedureCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB storedProcedure.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecAPIVersion">DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec">DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerThroughputSetting">SqlDatabaseContainerThroughputSetting</a>)
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM">DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsUpdatePropertiesARM">
ThroughputSettingsUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to update Azure Cosmos DB resource throughput.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggersSpecAPIVersion">DatabaseAccountsSqlDatabasesContainersTriggersSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec">DatabaseAccountsSqlDatabasesContainersTriggers_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerTrigger">SqlDatabaseContainerTrigger</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResource">
SqlTriggerResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL trigger resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM">DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB trigger name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerCreateUpdatePropertiesARM">
SqlTriggerCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB trigger.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecAPIVersion">DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec">DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerUserDefinedFunction">SqlDatabaseContainerUserDefinedFunction</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionResource">
SqlUserDefinedFunctionResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL userDefinedFunction resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM">DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB userDefinedFunction name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionCreateUpdatePropertiesARM">
SqlUserDefinedFunctionCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB userDefinedFunction.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainers_Spec">DatabaseAccountsSqlDatabasesContainers_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainer">SqlDatabaseContainer</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerResource">
SqlContainerResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL container resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainers_SpecARM">DatabaseAccountsSqlDatabasesContainers_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB container name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerCreateUpdatePropertiesARM">
SqlContainerCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB container.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesSpecAPIVersion">DatabaseAccountsSqlDatabasesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSettingsSpecAPIVersion">DatabaseAccountsSqlDatabasesThroughputSettingsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-15&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSettings_Spec">DatabaseAccountsSqlDatabasesThroughputSettings_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseThroughputSetting">SqlDatabaseThroughputSetting</a>)
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/SqlDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM">DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsUpdatePropertiesARM">
ThroughputSettingsUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to update Azure Cosmos DB resource throughput.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabases_Spec">DatabaseAccountsSqlDatabases_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabase">SqlDatabase</a>)
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/DatabaseAccount resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseResource">
SqlDatabaseResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL database resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabases_SpecARM">DatabaseAccountsSqlDatabases_SpecARM
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB database name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseCreateUpdatePropertiesARM">
SqlDatabaseCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB SQL database.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccount">DatabaseAccount</a>)
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
<code>analyticalStorageConfiguration</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AnalyticalStorageConfiguration">
AnalyticalStorageConfiguration
</a>
</em>
</td>
<td>
<p>AnalyticalStorageConfiguration: Analytical storage specific properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ApiProperties">
ApiProperties
</a>
</em>
</td>
<td>
</td>
</tr>
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
<code>backupPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.BackupPolicy">
BackupPolicy
</a>
</em>
</td>
<td>
<p>BackupPolicy: The object representing the policy for taking backups on an account.</p>
</td>
</tr>
<tr>
<td>
<code>capabilities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Capability">
[]Capability
</a>
</em>
</td>
<td>
<p>Capabilities: List of Cosmos DB capabilities for the account</p>
</td>
</tr>
<tr>
<td>
<code>connectorOffer</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesConnectorOffer">
DatabaseAccountCreateUpdatePropertiesConnectorOffer
</a>
</em>
</td>
<td>
<p>ConnectorOffer: The cassandra connector offer type for the Cosmos DB database C* account.</p>
</td>
</tr>
<tr>
<td>
<code>consistencyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConsistencyPolicy">
ConsistencyPolicy
</a>
</em>
</td>
<td>
<p>ConsistencyPolicy: The consistency policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>cors</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CorsPolicy">
[]CorsPolicy
</a>
</em>
</td>
<td>
<p>Cors: The CORS policy for the Cosmos DB database account.</p>
</td>
</tr>
<tr>
<td>
<code>databaseAccountOfferType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType">
DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferType
</a>
</em>
</td>
<td>
<p>DatabaseAccountOfferType: The offer type for the database</p>
</td>
</tr>
<tr>
<td>
<code>defaultIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultIdentity: The default identity for accessing key vault used in features like customer managed keys. The default
identity needs to be explicitly set by the users. It can be &ldquo;FirstPartyIdentity&rdquo;, &ldquo;SystemAssignedIdentity&rdquo; and more.</p>
</td>
</tr>
<tr>
<td>
<code>disableKeyBasedMetadataWriteAccess</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableKeyBasedMetadataWriteAccess: Disable write operations on metadata resources (databases, containers, throughput)
via account keys</p>
</td>
</tr>
<tr>
<td>
<code>enableAnalyticalStorage</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAnalyticalStorage: Flag to indicate whether to enable storage analytics.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutomaticFailover</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutomaticFailover: Enables automatic failover of the write region in the rare event that the region is unavailable
due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the
failover priorities configured for the account.</p>
</td>
</tr>
<tr>
<td>
<code>enableCassandraConnector</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCassandraConnector: Enables the cassandra connector on the Cosmos DB C* account</p>
</td>
</tr>
<tr>
<td>
<code>enableFreeTier</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFreeTier: Flag to indicate whether Free Tier is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleWriteLocations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleWriteLocations: Enables the account to write in multiple locations</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity">
ManagedServiceIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IpAddressOrRange">
[]IpAddressOrRange
</a>
</em>
</td>
<td>
<p>IpRules: Array of IpAddressOrRange objects.</p>
</td>
</tr>
<tr>
<td>
<code>isVirtualNetworkFilterEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsVirtualNetworkFilterEnabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultKeyUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultKeyUri: The URI of the key vault</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSpecKind">
DatabaseAccountsSpecKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of database account. This can only be set at database account creation.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>locations</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Location">
[]Location
</a>
</em>
</td>
<td>
<p>Locations: An array that contains the georeplication locations enabled for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypass</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesNetworkAclBypass">
DatabaseAccountCreateUpdatePropertiesNetworkAclBypass
</a>
</em>
</td>
<td>
<p>NetworkAclBypass: Indicates what services are allowed to bypass firewall checks.</p>
</td>
</tr>
<tr>
<td>
<code>networkAclBypassResourceIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NetworkAclBypassResourceIds: An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountOperatorSpec">
DatabaseAccountOperatorSpec
</a>
</em>
</td>
<td>
<p>OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
passed directly to Azure</p>
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
<code>publicNetworkAccess</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess">
DatabaseAccountCreateUpdatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether requests from Public Network are allowed.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.VirtualNetworkRule">
[]VirtualNetworkRule
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: List of Virtual Network ACL rules configured for the Cosmos DB account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.DatabaseAccounts_SpecARM">DatabaseAccounts_SpecARM
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
<code>identity</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentityARM">
ManagedServiceIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Identity for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSpecKind">
DatabaseAccountsSpecKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of database account. This can only be set at database account creation.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: Cosmos DB database account name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">
DatabaseAccountCreateUpdatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties to create and update Azure Cosmos DB database accounts.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ExcludedPath">ExcludedPath
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy">IndexingPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath</a></p>
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ExcludedPathARM">ExcludedPathARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyARM">IndexingPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath</a></p>
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ExcludedPath_Status">ExcludedPath_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">IndexingPolicy_Status</a>)
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ExcludedPath_StatusARM">ExcludedPath_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">IndexingPolicy_StatusARM</a>)
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.FailoverPolicy_Status">FailoverPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>failoverPriority</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
regions in which the database account exists.</p>
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
<p>Id: The unique identifier of the region in which the database account replicates to. Example:
&amp;lt;accountName&amp;gt;-&amp;lt;locationName&amp;gt;.</p>
</td>
</tr>
<tr>
<td>
<code>locationName</code><br/>
<em>
string
</em>
</td>
<td>
<p>LocationName: The name of the region in which the database account exists.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.FailoverPolicy_StatusARM">FailoverPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>failoverPriority</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
regions in which the database account exists.</p>
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
<p>Id: The unique identifier of the region in which the database account replicates to. Example:
&amp;lt;accountName&amp;gt;-&amp;lt;locationName&amp;gt;.</p>
</td>
</tr>
<tr>
<td>
<code>locationName</code><br/>
<em>
string
</em>
</td>
<td>
<p>LocationName: The name of the region in which the database account exists.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IncludedPath">IncludedPath
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy">IndexingPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath</a></p>
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
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Indexes">
[]Indexes
</a>
</em>
</td>
<td>
<p>Indexes: List of indexes for this path</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IncludedPathARM">IncludedPathARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyARM">IndexingPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath</a></p>
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
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesARM">
[]IndexesARM
</a>
</em>
</td>
<td>
<p>Indexes: List of indexes for this path</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IncludedPath_Status">IncludedPath_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">IndexingPolicy_Status</a>)
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
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Indexes_Status">
[]Indexes_Status
</a>
</em>
</td>
<td>
<p>Indexes: List of indexes for this path</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IncludedPath_StatusARM">IncludedPath_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">IndexingPolicy_StatusARM</a>)
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
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.Indexes_StatusARM">
[]Indexes_StatusARM
</a>
</em>
</td>
<td>
<p>Indexes: List of indexes for this path</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Indexes">Indexes
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IncludedPath">IncludedPath</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes</a></p>
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
<code>dataType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesDataType">
IndexesDataType
</a>
</em>
</td>
<td>
<p>DataType: The datatype for which the indexing behavior is applied to.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesKind">
IndexesKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of index.</p>
</td>
</tr>
<tr>
<td>
<code>precision</code><br/>
<em>
int
</em>
</td>
<td>
<p>Precision: The precision of the index. -1 is maximum precision.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexesARM">IndexesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IncludedPathARM">IncludedPathARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes</a></p>
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
<code>dataType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesDataType">
IndexesDataType
</a>
</em>
</td>
<td>
<p>DataType: The datatype for which the indexing behavior is applied to.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesKind">
IndexesKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of index.</p>
</td>
</tr>
<tr>
<td>
<code>precision</code><br/>
<em>
int
</em>
</td>
<td>
<p>Precision: The precision of the index. -1 is maximum precision.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexesDataType">IndexesDataType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.Indexes">Indexes</a>, <a href="#documentdb.azure.com/v1beta20210515.IndexesARM">IndexesARM</a>)
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
<tbody><tr><td><p>&#34;LineString&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MultiPolygon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Number&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Point&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Polygon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;String&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexesKind">IndexesKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.Indexes">Indexes</a>, <a href="#documentdb.azure.com/v1beta20210515.IndexesARM">IndexesARM</a>)
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
<tbody><tr><td><p>&#34;Hash&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Range&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spatial&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexesStatusDataType">IndexesStatusDataType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.Indexes_Status">Indexes_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.Indexes_StatusARM">Indexes_StatusARM</a>)
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
<tbody><tr><td><p>&#34;LineString&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MultiPolygon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Number&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Point&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Polygon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;String&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexesStatusKind">IndexesStatusKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.Indexes_Status">Indexes_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.Indexes_StatusARM">Indexes_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Hash&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Range&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spatial&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Indexes_Status">Indexes_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IncludedPath_Status">IncludedPath_Status</a>)
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
<code>dataType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesStatusDataType">
IndexesStatusDataType
</a>
</em>
</td>
<td>
<p>DataType: The datatype for which the indexing behavior is applied to.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesStatusKind">
IndexesStatusKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of index.</p>
</td>
</tr>
<tr>
<td>
<code>precision</code><br/>
<em>
int
</em>
</td>
<td>
<p>Precision: The precision of the index. -1 is maximum precision.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Indexes_StatusARM">Indexes_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IncludedPath_StatusARM">IncludedPath_StatusARM</a>)
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
<code>dataType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesStatusDataType">
IndexesStatusDataType
</a>
</em>
</td>
<td>
<p>DataType: The datatype for which the indexing behavior is applied to.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexesStatusKind">
IndexesStatusKind
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of index.</p>
</td>
</tr>
<tr>
<td>
<code>precision</code><br/>
<em>
int
</em>
</td>
<td>
<p>Precision: The precision of the index. -1 is maximum precision.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexingPolicy">IndexingPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResource">SqlContainerResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy</a></p>
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
<code>automatic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Automatic: Indicates if the indexing policy is automatic</p>
</td>
</tr>
<tr>
<td>
<code>compositeIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePath">
[]CompositePath
</a>
</em>
</td>
<td>
<p>CompositeIndexes: List of composite path list</p>
</td>
</tr>
<tr>
<td>
<code>excludedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ExcludedPath">
[]ExcludedPath
</a>
</em>
</td>
<td>
<p>ExcludedPaths: List of paths to exclude from indexing</p>
</td>
</tr>
<tr>
<td>
<code>includedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IncludedPath">
[]IncludedPath
</a>
</em>
</td>
<td>
<p>IncludedPaths: List of paths to include in the indexing</p>
</td>
</tr>
<tr>
<td>
<code>indexingMode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyIndexingMode">
IndexingPolicyIndexingMode
</a>
</em>
</td>
<td>
<p>IndexingMode: Indicates the indexing mode.</p>
</td>
</tr>
<tr>
<td>
<code>spatialIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialSpec">
[]SpatialSpec
</a>
</em>
</td>
<td>
<p>SpatialIndexes: List of spatial specifics</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexingPolicyARM">IndexingPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResourceARM">SqlContainerResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy</a></p>
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
<code>automatic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Automatic: Indicates if the indexing policy is automatic</p>
</td>
</tr>
<tr>
<td>
<code>compositeIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePathARM">
[]CompositePathARM
</a>
</em>
</td>
<td>
<p>CompositeIndexes: List of composite path list</p>
</td>
</tr>
<tr>
<td>
<code>excludedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ExcludedPathARM">
[]ExcludedPathARM
</a>
</em>
</td>
<td>
<p>ExcludedPaths: List of paths to exclude from indexing</p>
</td>
</tr>
<tr>
<td>
<code>includedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IncludedPathARM">
[]IncludedPathARM
</a>
</em>
</td>
<td>
<p>IncludedPaths: List of paths to include in the indexing</p>
</td>
</tr>
<tr>
<td>
<code>indexingMode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyIndexingMode">
IndexingPolicyIndexingMode
</a>
</em>
</td>
<td>
<p>IndexingMode: Indicates the indexing mode.</p>
</td>
</tr>
<tr>
<td>
<code>spatialIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialSpecARM">
[]SpatialSpecARM
</a>
</em>
</td>
<td>
<p>SpatialIndexes: List of spatial specifics</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexingPolicyIndexingMode">IndexingPolicyIndexingMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy">IndexingPolicy</a>, <a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyARM">IndexingPolicyARM</a>)
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
<tbody><tr><td><p>&#34;consistent&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;lazy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexingPolicyStatusIndexingMode">IndexingPolicyStatusIndexingMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">IndexingPolicy_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">IndexingPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;consistent&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;lazy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">IndexingPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_Resource">SqlContainerGetProperties_Status_Resource</a>)
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
<code>automatic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Automatic: Indicates if the indexing policy is automatic</p>
</td>
</tr>
<tr>
<td>
<code>compositeIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePath_Status">
[]CompositePath_Status
</a>
</em>
</td>
<td>
<p>CompositeIndexes: List of composite path list</p>
</td>
</tr>
<tr>
<td>
<code>excludedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ExcludedPath_Status">
[]ExcludedPath_Status
</a>
</em>
</td>
<td>
<p>ExcludedPaths: List of paths to exclude from indexing</p>
</td>
</tr>
<tr>
<td>
<code>includedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IncludedPath_Status">
[]IncludedPath_Status
</a>
</em>
</td>
<td>
<p>IncludedPaths: List of paths to include in the indexing</p>
</td>
</tr>
<tr>
<td>
<code>indexingMode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyStatusIndexingMode">
IndexingPolicyStatusIndexingMode
</a>
</em>
</td>
<td>
<p>IndexingMode: Indicates the indexing mode.</p>
</td>
</tr>
<tr>
<td>
<code>spatialIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialSpec_Status">
[]SpatialSpec_Status
</a>
</em>
</td>
<td>
<p>SpatialIndexes: List of spatial specifics</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">IndexingPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_ResourceARM">SqlContainerGetProperties_Status_ResourceARM</a>)
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
<code>automatic</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Automatic: Indicates if the indexing policy is automatic</p>
</td>
</tr>
<tr>
<td>
<code>compositeIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CompositePath_StatusARM">
[]CompositePath_StatusARM
</a>
</em>
</td>
<td>
<p>CompositeIndexes: List of composite path list</p>
</td>
</tr>
<tr>
<td>
<code>excludedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ExcludedPath_StatusARM">
[]ExcludedPath_StatusARM
</a>
</em>
</td>
<td>
<p>ExcludedPaths: List of paths to exclude from indexing</p>
</td>
</tr>
<tr>
<td>
<code>includedPaths</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IncludedPath_StatusARM">
[]IncludedPath_StatusARM
</a>
</em>
</td>
<td>
<p>IncludedPaths: List of paths to include in the indexing</p>
</td>
</tr>
<tr>
<td>
<code>indexingMode</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyStatusIndexingMode">
IndexingPolicyStatusIndexingMode
</a>
</em>
</td>
<td>
<p>IndexingMode: Indicates the indexing mode.</p>
</td>
</tr>
<tr>
<td>
<code>spatialIndexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialSpec_StatusARM">
[]SpatialSpec_StatusARM
</a>
</em>
</td>
<td>
<p>SpatialIndexes: List of spatial specifics</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IpAddressOrRange">IpAddressOrRange
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IpAddressOrRange">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IpAddressOrRange</a></p>
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
<code>ipAddressOrRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
or “23.40.210.0/8”.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IpAddressOrRangeARM">IpAddressOrRangeARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IpAddressOrRange">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IpAddressOrRange</a></p>
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
<code>ipAddressOrRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
or “23.40.210.0/8”.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IpAddressOrRange_Status">IpAddressOrRange_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>ipAddressOrRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
or “23.40.210.0/8”.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.IpAddressOrRange_StatusARM">IpAddressOrRange_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>ipAddressOrRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpAddressOrRange: A single IPv4 address or a single IPv4 address range in CIDR format. Provided IPs must be
well-formatted and cannot be contained in one of the following ranges: 10.0.0.0/8, 100.64.0.0/10, 172.16.0.0/12,
192.168.0.0/16, since these are not enforceable by the IP address filter. Example of valid inputs: “23.40.210.245”
or “23.40.210.0/8”.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Location">Location
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Location">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Location</a></p>
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
<code>failoverPriority</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
regions in which the database account exists.</p>
</td>
</tr>
<tr>
<td>
<code>isZoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region</p>
</td>
</tr>
<tr>
<td>
<code>locationName</code><br/>
<em>
string
</em>
</td>
<td>
<p>LocationName: The name of the region.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.LocationARM">LocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Location">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Location</a></p>
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
<code>failoverPriority</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
regions in which the database account exists.</p>
</td>
</tr>
<tr>
<td>
<code>isZoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region</p>
</td>
</tr>
<tr>
<td>
<code>locationName</code><br/>
<em>
string
</em>
</td>
<td>
<p>LocationName: The name of the region.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Location_Status">Location_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>documentEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>DocumentEndpoint: The connection endpoint for the specific region. Example:
https://&amp;lt;accountName&amp;gt;-&amp;lt;locationName&amp;gt;.documents.azure.com:443/</p>
</td>
</tr>
<tr>
<td>
<code>failoverPriority</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
regions in which the database account exists.</p>
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
<p>Id: The unique identifier of the region within the database account. Example: &amp;lt;accountName&amp;gt;-&amp;lt;locationName&amp;gt;.</p>
</td>
</tr>
<tr>
<td>
<code>isZoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region</p>
</td>
</tr>
<tr>
<td>
<code>locationName</code><br/>
<em>
string
</em>
</td>
<td>
<p>LocationName: The name of the region.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.Location_StatusARM">Location_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<code>documentEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>DocumentEndpoint: The connection endpoint for the specific region. Example:
https://&amp;lt;accountName&amp;gt;-&amp;lt;locationName&amp;gt;.documents.azure.com:443/</p>
</td>
</tr>
<tr>
<td>
<code>failoverPriority</code><br/>
<em>
int
</em>
</td>
<td>
<p>FailoverPriority: The failover priority of the region. A failover priority of 0 indicates a write region. The maximum
value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the
regions in which the database account exists.</p>
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
<p>Id: The unique identifier of the region within the database account. Example: &amp;lt;accountName&amp;gt;-&amp;lt;locationName&amp;gt;.</p>
</td>
</tr>
<tr>
<td>
<code>isZoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsZoneRedundant: Flag to indicate whether or not this region is an AvailabilityZone region</p>
</td>
</tr>
<tr>
<td>
<code>locationName</code><br/>
<em>
string
</em>
</td>
<td>
<p>LocationName: The name of the region.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentity">ManagedServiceIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ManagedServiceIdentity">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ManagedServiceIdentity</a></p>
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
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentityType">
ManagedServiceIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned,UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentityARM">ManagedServiceIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_SpecARM">DatabaseAccounts_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ManagedServiceIdentity">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ManagedServiceIdentity</a></p>
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
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentityType">
ManagedServiceIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned,UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentityStatusType">ManagedServiceIdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status">ManagedServiceIdentity_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_StatusARM">ManagedServiceIdentity_StatusARM</a>)
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
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned,UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentityType">ManagedServiceIdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity">ManagedServiceIdentity</a>, <a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentityARM">ManagedServiceIdentityARM</a>)
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
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned,UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status">ManagedServiceIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of the system assigned identity. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant id of the system assigned identity. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentityStatusType">
ManagedServiceIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned,UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status_UserAssignedIdentities">
map[string]./api/documentdb/v1beta20210515.ManagedServiceIdentity_Status_UserAssignedIdentities
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with resource. The user identity dictionary key
references will be ARM resource ids in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_StatusARM">ManagedServiceIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_StatusARM">DatabaseAccountGetResults_StatusARM</a>)
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
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of the system assigned identity. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant id of the system assigned identity. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentityStatusType">
ManagedServiceIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned,UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status_UserAssignedIdentitiesARM">
map[string]./api/documentdb/v1beta20210515.ManagedServiceIdentity_Status_UserAssignedIdentitiesARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with resource. The user identity dictionary key
references will be ARM resource ids in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status_UserAssignedIdentities">ManagedServiceIdentity_Status_UserAssignedIdentities
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status">ManagedServiceIdentity_Status</a>)
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
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_Status_UserAssignedIdentitiesARM">ManagedServiceIdentity_Status_UserAssignedIdentitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ManagedServiceIdentity_StatusARM">ManagedServiceIdentity_StatusARM</a>)
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
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionCreateUpdatePropertiesARM">MongoDBCollectionCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollections_SpecARM">DatabaseAccountsMongodbDatabasesCollections_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionResourceARM">
MongoDBCollectionResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB MongoDB collection resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_StatusARM">MongoDBCollectionGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetResults_StatusARM">MongoDBCollectionGetResults_StatusARM</a>)
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_StatusARM">
OptionsResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_Status_ResourceARM">
MongoDBCollectionGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_Status_Resource">MongoDBCollectionGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetResults_Status">MongoDBCollectionGetResults_Status</a>)
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB MongoDB collection</p>
</td>
</tr>
<tr>
<td>
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndex_Status">
[]MongoIndex_Status
</a>
</em>
</td>
<td>
<p>Indexes: List of index keys</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>shardKey</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ShardKey: A key-value pair of shard keys to be applied for the request.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_Status_ResourceARM">MongoDBCollectionGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_StatusARM">MongoDBCollectionGetProperties_StatusARM</a>)
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB MongoDB collection</p>
</td>
</tr>
<tr>
<td>
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndex_StatusARM">
[]MongoIndex_StatusARM
</a>
</em>
</td>
<td>
<p>Indexes: List of index keys</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>shardKey</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ShardKey: A key-value pair of shard keys to be applied for the request.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionGetResults_Status">MongoDBCollectionGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabaseCollection">MongodbDatabaseCollection</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_Status">
OptionsResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_Status_Resource">
MongoDBCollectionGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionGetResults_StatusARM">MongoDBCollectionGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_StatusARM">
MongoDBCollectionGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB MongoDB collection</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionResource">MongoDBCollectionResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollections_Spec">DatabaseAccountsMongodbDatabasesCollections_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource</a></p>
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
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
<p>Id: Name of the Cosmos DB MongoDB collection</p>
</td>
</tr>
<tr>
<td>
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndex">
[]MongoIndex
</a>
</em>
</td>
<td>
<p>Indexes: List of index keys</p>
</td>
</tr>
<tr>
<td>
<code>shardKey</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ShardKey: The shard key and partition kind pair, only support &ldquo;Hash&rdquo; partition kind</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBCollectionResourceARM">MongoDBCollectionResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionCreateUpdatePropertiesARM">MongoDBCollectionCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource</a></p>
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
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
<p>Id: Name of the Cosmos DB MongoDB collection</p>
</td>
</tr>
<tr>
<td>
<code>indexes</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexARM">
[]MongoIndexARM
</a>
</em>
</td>
<td>
<p>Indexes: List of index keys</p>
</td>
</tr>
<tr>
<td>
<code>shardKey</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ShardKey: The shard key and partition kind pair, only support &ldquo;Hash&rdquo; partition kind</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseCreateUpdatePropertiesARM">MongoDBDatabaseCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabases_SpecARM">DatabaseAccountsMongodbDatabases_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseResourceARM">
MongoDBDatabaseResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB MongoDB database resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_StatusARM">MongoDBDatabaseGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetResults_StatusARM">MongoDBDatabaseGetResults_StatusARM</a>)
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_StatusARM">
OptionsResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_Status_ResourceARM">
MongoDBDatabaseGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_Status_Resource">MongoDBDatabaseGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetResults_Status">MongoDBDatabaseGetResults_Status</a>)
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
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB MongoDB database</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_Status_ResourceARM">MongoDBDatabaseGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_StatusARM">MongoDBDatabaseGetProperties_StatusARM</a>)
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
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB MongoDB database</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetResults_Status">MongoDBDatabaseGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabase">MongodbDatabase</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_Status">
OptionsResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_Status_Resource">
MongoDBDatabaseGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetResults_StatusARM">MongoDBDatabaseGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_StatusARM">
MongoDBDatabaseGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB MongoDB database</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseResource">MongoDBDatabaseResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabases_Spec">DatabaseAccountsMongodbDatabases_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseResource</a></p>
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
<p>Id: Name of the Cosmos DB MongoDB database</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoDBDatabaseResourceARM">MongoDBDatabaseResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseCreateUpdatePropertiesARM">MongoDBDatabaseCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBDatabaseResource</a></p>
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
<p>Id: Name of the Cosmos DB MongoDB database</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndex">MongoIndex
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionResource">MongoDBCollectionResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex</a></p>
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
<code>key</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexKeys">
MongoIndexKeys
</a>
</em>
</td>
<td>
<p>Key: Cosmos DB MongoDB collection resource object</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexOptions">
MongoIndexOptions
</a>
</em>
</td>
<td>
<p>Options: Cosmos DB MongoDB collection index options</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexARM">MongoIndexARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionResourceARM">MongoDBCollectionResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex</a></p>
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
<code>key</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexKeysARM">
MongoIndexKeysARM
</a>
</em>
</td>
<td>
<p>Key: Cosmos DB MongoDB collection resource object</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexOptionsARM">
MongoIndexOptionsARM
</a>
</em>
</td>
<td>
<p>Options: Cosmos DB MongoDB collection index options</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexKeys">MongoIndexKeys
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndex">MongoIndex</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys</a></p>
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
<code>keys</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexKeysARM">MongoIndexKeysARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndexARM">MongoIndexARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys</a></p>
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
<code>keys</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexKeys_Status">MongoIndexKeys_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndex_Status">MongoIndex_Status</a>)
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
<code>keys</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexKeys_StatusARM">MongoIndexKeys_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndex_StatusARM">MongoIndex_StatusARM</a>)
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
<code>keys</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexOptions">MongoIndexOptions
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndex">MongoIndex</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions</a></p>
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
<code>expireAfterSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExpireAfterSeconds: Expire after seconds</p>
</td>
</tr>
<tr>
<td>
<code>unique</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Unique: Is unique or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexOptionsARM">MongoIndexOptionsARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndexARM">MongoIndexARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions</a></p>
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
<code>expireAfterSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExpireAfterSeconds: Expire after seconds</p>
</td>
</tr>
<tr>
<td>
<code>unique</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Unique: Is unique or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexOptions_Status">MongoIndexOptions_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndex_Status">MongoIndex_Status</a>)
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
<code>expireAfterSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExpireAfterSeconds: Expire after seconds</p>
</td>
</tr>
<tr>
<td>
<code>unique</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Unique: Is unique or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndexOptions_StatusARM">MongoIndexOptions_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoIndex_StatusARM">MongoIndex_StatusARM</a>)
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
<code>expireAfterSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExpireAfterSeconds: Expire after seconds</p>
</td>
</tr>
<tr>
<td>
<code>unique</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Unique: Is unique or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndex_Status">MongoIndex_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_Status_Resource">MongoDBCollectionGetProperties_Status_Resource</a>)
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
<code>key</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexKeys_Status">
MongoIndexKeys_Status
</a>
</em>
</td>
<td>
<p>Key: Cosmos DB MongoDB collection index keys</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexOptions_Status">
MongoIndexOptions_Status
</a>
</em>
</td>
<td>
<p>Options: Cosmos DB MongoDB collection index key options</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongoIndex_StatusARM">MongoIndex_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_Status_ResourceARM">MongoDBCollectionGetProperties_Status_ResourceARM</a>)
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
<code>key</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexKeys_StatusARM">
MongoIndexKeys_StatusARM
</a>
</em>
</td>
<td>
<p>Key: Cosmos DB MongoDB collection index keys</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoIndexOptions_StatusARM">
MongoIndexOptions_StatusARM
</a>
</em>
</td>
<td>
<p>Options: Cosmos DB MongoDB collection index key options</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongodbDatabase">MongodbDatabase
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabases_Spec">
DatabaseAccountsMongodbDatabases_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/DatabaseAccount resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseResource">
MongoDBDatabaseResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB MongoDB database resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetResults_Status">
MongoDBDatabaseGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongodbDatabaseCollection">MongodbDatabaseCollection
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollections_Spec">
DatabaseAccountsMongodbDatabasesCollections_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/MongodbDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionResource">
MongoDBCollectionResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB MongoDB collection resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetResults_Status">
MongoDBCollectionGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongodbDatabaseCollectionThroughputSetting">MongodbDatabaseCollectionThroughputSetting
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections_throughputSettings">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections_throughputSettings</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec">
DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/MongodbDatabaseCollection resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_Status">
ThroughputSettingsGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.MongodbDatabaseThroughputSetting">MongodbDatabaseThroughputSetting
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec">
DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/MongodbDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_Status">
ThroughputSettingsGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.NetworkAclBypass_Status">NetworkAclBypass_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<tbody><tr><td><p>&#34;AzureServices&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.OptionsResource_Status">OptionsResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetResults_Status">MongoDBCollectionGetResults_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetResults_Status">MongoDBDatabaseGetResults_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetResults_Status">SqlContainerGetResults_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetResults_Status">SqlDatabaseGetResults_Status</a>)
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettings_Status">
AutoscaleSettings_Status
</a>
</em>
</td>
<td>
<p>AutoscaleSettings: Specifies the Autoscale settings.</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Value of the Cosmos DB resource throughput or autoscaleSettings. Use the ThroughputSetting resource when
retrieving offer details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.OptionsResource_StatusARM">OptionsResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongoDBCollectionGetProperties_StatusARM">MongoDBCollectionGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.MongoDBDatabaseGetProperties_StatusARM">MongoDBDatabaseGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_StatusARM">SqlContainerGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_StatusARM">SqlDatabaseGetProperties_StatusARM</a>)
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettings_StatusARM">
AutoscaleSettings_StatusARM
</a>
</em>
</td>
<td>
<p>AutoscaleSettings: Specifies the Autoscale settings.</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Value of the Cosmos DB resource throughput or autoscaleSettings. Use the ThroughputSetting resource when
retrieving offer details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicy">PeriodicModeBackupPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.BackupPolicy">BackupPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeBackupPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeBackupPolicy</a></p>
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
<code>periodicModeProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PeriodicModeProperties">
PeriodicModeProperties
</a>
</em>
</td>
<td>
<p>PeriodicModeProperties: Configuration values for periodic mode backup</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyType">
PeriodicModeBackupPolicyType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyARM">PeriodicModeBackupPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.BackupPolicyARM">BackupPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeBackupPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeBackupPolicy</a></p>
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
<code>periodicModeProperties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PeriodicModePropertiesARM">
PeriodicModePropertiesARM
</a>
</em>
</td>
<td>
<p>PeriodicModeProperties: Configuration values for periodic mode backup</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyType">
PeriodicModeBackupPolicyType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyType">PeriodicModeBackupPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicy">PeriodicModeBackupPolicy</a>, <a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyARM">PeriodicModeBackupPolicyARM</a>)
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
<tbody><tr><td><p>&#34;Periodic&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.PeriodicModeProperties">PeriodicModeProperties
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicy">PeriodicModeBackupPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeProperties</a></p>
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
<code>backupIntervalInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupIntervalInMinutes: An integer representing the interval in minutes between two backups</p>
</td>
</tr>
<tr>
<td>
<code>backupRetentionIntervalInHours</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupRetentionIntervalInHours: An integer representing the time (in hours) that each backup is retained</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.PeriodicModePropertiesARM">PeriodicModePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.PeriodicModeBackupPolicyARM">PeriodicModeBackupPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/PeriodicModeProperties</a></p>
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
<code>backupIntervalInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupIntervalInMinutes: An integer representing the interval in minutes between two backups</p>
</td>
</tr>
<tr>
<td>
<code>backupRetentionIntervalInHours</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupRetentionIntervalInHours: An integer representing the time (in hours) that each backup is retained</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<h3 id="documentdb.azure.com/v1beta20210515.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<h3 id="documentdb.azure.com/v1beta20210515.PublicNetworkAccess_Status">PublicNetworkAccess_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<h3 id="documentdb.azure.com/v1beta20210515.SpatialSpec">SpatialSpec
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy">IndexingPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec</a></p>
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
<tr>
<td>
<code>types</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialSpecTypes">
[]SpatialSpecTypes
</a>
</em>
</td>
<td>
<p>Types: List of path&rsquo;s spatial type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SpatialSpecARM">SpatialSpecARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyARM">IndexingPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec</a></p>
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
<tr>
<td>
<code>types</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialSpecTypes">
[]SpatialSpecTypes
</a>
</em>
</td>
<td>
<p>Types: List of path&rsquo;s spatial type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SpatialSpecTypes">SpatialSpecTypes
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SpatialSpec">SpatialSpec</a>, <a href="#documentdb.azure.com/v1beta20210515.SpatialSpecARM">SpatialSpecARM</a>)
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
<tbody><tr><td><p>&#34;LineString&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MultiPolygon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Point&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Polygon&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SpatialSpec_Status">SpatialSpec_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">IndexingPolicy_Status</a>)
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
<tr>
<td>
<code>types</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialType_Status">
[]SpatialType_Status
</a>
</em>
</td>
<td>
<p>Types: List of path&rsquo;s spatial type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SpatialSpec_StatusARM">SpatialSpec_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">IndexingPolicy_StatusARM</a>)
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
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
(/path/*)</p>
</td>
</tr>
<tr>
<td>
<code>types</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SpatialType_Status">
[]SpatialType_Status
</a>
</em>
</td>
<td>
<p>Types: List of path&rsquo;s spatial type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SpatialType_Status">SpatialType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SpatialSpec_Status">SpatialSpec_Status</a>, <a href="#documentdb.azure.com/v1beta20210515.SpatialSpec_StatusARM">SpatialSpec_StatusARM</a>)
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
<tbody><tr><td><p>&#34;LineString&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MultiPolygon&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Point&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Polygon&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerCreateUpdatePropertiesARM">SqlContainerCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainers_SpecARM">DatabaseAccountsSqlDatabasesContainers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerResourceARM">
SqlContainerResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL container resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_StatusARM">SqlContainerGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetResults_StatusARM">SqlContainerGetResults_StatusARM</a>)
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_StatusARM">
OptionsResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_ResourceARM">
SqlContainerGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_Resource">SqlContainerGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetResults_Status">SqlContainerGetResults_Status</a>)
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy_Status">
ConflictResolutionPolicy_Status
</a>
</em>
</td>
<td>
<p>ConflictResolutionPolicy: The conflict resolution policy for the container.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>DefaultTtl: Default time to live</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL container</p>
</td>
</tr>
<tr>
<td>
<code>indexingPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_Status">
IndexingPolicy_Status
</a>
</em>
</td>
<td>
<p>IndexingPolicy: The configuration of the indexing policy. By default, the indexing is automatic for all document paths
within the container</p>
</td>
</tr>
<tr>
<td>
<code>partitionKey</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKey_Status">
ContainerPartitionKey_Status
</a>
</em>
</td>
<td>
<p>PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>uniqueKeyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicy_Status">
UniqueKeyPolicy_Status
</a>
</em>
</td>
<td>
<p>UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
collection in the Azure Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_ResourceARM">SqlContainerGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_StatusARM">SqlContainerGetProperties_StatusARM</a>)
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy_StatusARM">
ConflictResolutionPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>ConflictResolutionPolicy: The conflict resolution policy for the container.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>DefaultTtl: Default time to live</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL container</p>
</td>
</tr>
<tr>
<td>
<code>indexingPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy_StatusARM">
IndexingPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>IndexingPolicy: The configuration of the indexing policy. By default, the indexing is automatic for all document paths
within the container</p>
</td>
</tr>
<tr>
<td>
<code>partitionKey</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKey_StatusARM">
ContainerPartitionKey_StatusARM
</a>
</em>
</td>
<td>
<p>PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>uniqueKeyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicy_StatusARM">
UniqueKeyPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
collection in the Azure Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerGetResults_Status">SqlContainerGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainer">SqlDatabaseContainer</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_Status">
OptionsResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_Resource">
SqlContainerGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerGetResults_StatusARM">SqlContainerGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_StatusARM">
SqlContainerGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB container</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerResource">SqlContainerResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainers_Spec">DatabaseAccountsSqlDatabasesContainers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource</a></p>
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicy">
ConflictResolutionPolicy
</a>
</em>
</td>
<td>
<p>ConflictResolutionPolicy: The conflict resolution policy for the container.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>DefaultTtl: Default time to live</p>
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
<p>Id: Name of the Cosmos DB SQL container</p>
</td>
</tr>
<tr>
<td>
<code>indexingPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicy">
IndexingPolicy
</a>
</em>
</td>
<td>
<p>IndexingPolicy: Cosmos DB indexing policy</p>
</td>
</tr>
<tr>
<td>
<code>partitionKey</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKey">
ContainerPartitionKey
</a>
</em>
</td>
<td>
<p>PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions</p>
</td>
</tr>
<tr>
<td>
<code>uniqueKeyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicy">
UniqueKeyPolicy
</a>
</em>
</td>
<td>
<p>UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
collection in the Azure Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlContainerResourceARM">SqlContainerResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerCreateUpdatePropertiesARM">SqlContainerCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource</a></p>
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
<code>analyticalStorageTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>AnalyticalStorageTtl: Analytical TTL.</p>
</td>
</tr>
<tr>
<td>
<code>conflictResolutionPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ConflictResolutionPolicyARM">
ConflictResolutionPolicyARM
</a>
</em>
</td>
<td>
<p>ConflictResolutionPolicy: The conflict resolution policy for the container.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTtl</code><br/>
<em>
int
</em>
</td>
<td>
<p>DefaultTtl: Default time to live</p>
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
<p>Id: Name of the Cosmos DB SQL container</p>
</td>
</tr>
<tr>
<td>
<code>indexingPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.IndexingPolicyARM">
IndexingPolicyARM
</a>
</em>
</td>
<td>
<p>IndexingPolicy: Cosmos DB indexing policy</p>
</td>
</tr>
<tr>
<td>
<code>partitionKey</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ContainerPartitionKeyARM">
ContainerPartitionKeyARM
</a>
</em>
</td>
<td>
<p>PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions</p>
</td>
</tr>
<tr>
<td>
<code>uniqueKeyPolicy</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicyARM">
UniqueKeyPolicyARM
</a>
</em>
</td>
<td>
<p>UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
collection in the Azure Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabase">SqlDatabase
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabases_Spec">
DatabaseAccountsSqlDatabases_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/DatabaseAccount resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseResource">
SqlDatabaseResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL database resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetResults_Status">
SqlDatabaseGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseContainer">SqlDatabaseContainer
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainers_Spec">
DatabaseAccountsSqlDatabasesContainers_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerResource">
SqlContainerResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL container resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetResults_Status">
SqlContainerGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseContainerStoredProcedure">SqlDatabaseContainerStoredProcedure
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_storedProcedures">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_storedProcedures</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec">
DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureResource">
SqlStoredProcedureResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL storedProcedure resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetResults_Status">
SqlStoredProcedureGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseContainerThroughputSetting">SqlDatabaseContainerThroughputSetting
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_throughputSettings">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_throughputSettings</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec">
DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_Status">
ThroughputSettingsGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseContainerTrigger">SqlDatabaseContainerTrigger
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_triggers">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_triggers</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec">
DatabaseAccountsSqlDatabasesContainersTriggers_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResource">
SqlTriggerResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL trigger resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetResults_Status">
SqlTriggerGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseContainerUserDefinedFunction">SqlDatabaseContainerUserDefinedFunction
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_userDefinedFunctions">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_userDefinedFunctions</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec">
DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
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
<p>Location: The location of the resource group to which the resource belongs.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptions">
CreateUpdateOptions
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
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
reference to a documentdb.azure.com/SqlDatabaseContainer resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionResource">
SqlUserDefinedFunctionResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL userDefinedFunction resource object</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetResults_Status">
SqlUserDefinedFunctionGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseCreateUpdatePropertiesARM">SqlDatabaseCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabases_SpecARM">DatabaseAccountsSqlDatabases_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseResourceARM">
SqlDatabaseResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL database resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_StatusARM">SqlDatabaseGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetResults_StatusARM">SqlDatabaseGetResults_StatusARM</a>)
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_StatusARM">
OptionsResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_Status_ResourceARM">
SqlDatabaseGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_Status_Resource">SqlDatabaseGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetResults_Status">SqlDatabaseGetResults_Status</a>)
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
<code>_colls</code><br/>
<em>
string
</em>
</td>
<td>
<p>Colls: A system generated property that specified the addressable path of the collections resource.</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL database</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>_users</code><br/>
<em>
string
</em>
</td>
<td>
<p>Users: A system generated property that specifies the addressable path of the users resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_Status_ResourceARM">SqlDatabaseGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_StatusARM">SqlDatabaseGetProperties_StatusARM</a>)
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
<code>_colls</code><br/>
<em>
string
</em>
</td>
<td>
<p>Colls: A system generated property that specified the addressable path of the collections resource.</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL database</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>_users</code><br/>
<em>
string
</em>
</td>
<td>
<p>Users: A system generated property that specifies the addressable path of the users resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseGetResults_Status">SqlDatabaseGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabase">SqlDatabase</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.OptionsResource_Status">
OptionsResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_Status_Resource">
SqlDatabaseGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseGetResults_StatusARM">SqlDatabaseGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseGetProperties_StatusARM">
SqlDatabaseGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB SQL database</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseResource">SqlDatabaseResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabases_Spec">DatabaseAccountsSqlDatabases_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseResource</a></p>
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
<p>Id: Name of the Cosmos DB SQL database</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseResourceARM">SqlDatabaseResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseCreateUpdatePropertiesARM">SqlDatabaseCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseResource</a></p>
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
<p>Id: Name of the Cosmos DB SQL database</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlDatabaseThroughputSetting">SqlDatabaseThroughputSetting
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings</a></p>
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
<a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSettings_Spec">
DatabaseAccountsSqlDatabasesThroughputSettings_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location of the resource group to which the resource belongs.</p>
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
reference to a documentdb.azure.com/SqlDatabase resource</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">
ThroughputSettingsResource
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
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
<p>Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
type is set with &ldquo;defaultExperience&rdquo;: &ldquo;Cassandra&rdquo;. Current &ldquo;defaultExperience&rdquo; values also include &ldquo;Table&rdquo;, &ldquo;Graph&rdquo;,
&ldquo;DocumentDB&rdquo;, and &ldquo;MongoDB&rdquo;.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_Status">
ThroughputSettingsGetResults_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureCreateUpdatePropertiesARM">SqlStoredProcedureCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM">DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureResourceARM">
SqlStoredProcedureResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL storedProcedure resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_StatusARM">SqlStoredProcedureGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetResults_StatusARM">SqlStoredProcedureGetResults_StatusARM</a>)
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
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_Status_ResourceARM">
SqlStoredProcedureGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_Status_Resource">SqlStoredProcedureGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetResults_Status">SqlStoredProcedureGetResults_Status</a>)
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Stored Procedure</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL storedProcedure</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_Status_ResourceARM">SqlStoredProcedureGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_StatusARM">SqlStoredProcedureGetProperties_StatusARM</a>)
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Stored Procedure</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL storedProcedure</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetResults_Status">SqlStoredProcedureGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerStoredProcedure">SqlDatabaseContainerStoredProcedure</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_Status_Resource">
SqlStoredProcedureGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetResults_StatusARM">SqlStoredProcedureGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureGetProperties_StatusARM">
SqlStoredProcedureGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB storedProcedure</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureResource">SqlStoredProcedureResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec">DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureResource</a></p>
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Stored Procedure</p>
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
<p>Id: Name of the Cosmos DB SQL storedProcedure</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlStoredProcedureResourceARM">SqlStoredProcedureResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlStoredProcedureCreateUpdatePropertiesARM">SqlStoredProcedureCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureResource</a></p>
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Stored Procedure</p>
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
<p>Id: Name of the Cosmos DB SQL storedProcedure</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerCreateUpdatePropertiesARM">SqlTriggerCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM">DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceARM">
SqlTriggerResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL trigger resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetPropertiesStatusResourceTriggerOperation">SqlTriggerGetPropertiesStatusResourceTriggerOperation
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_Resource">SqlTriggerGetProperties_Status_Resource</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_ResourceARM">SqlTriggerGetProperties_Status_ResourceARM</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Replace&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetPropertiesStatusResourceTriggerType">SqlTriggerGetPropertiesStatusResourceTriggerType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_Resource">SqlTriggerGetProperties_Status_Resource</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_ResourceARM">SqlTriggerGetProperties_Status_ResourceARM</a>)
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
<tbody><tr><td><p>&#34;Post&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Pre&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_StatusARM">SqlTriggerGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetResults_StatusARM">SqlTriggerGetResults_StatusARM</a>)
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
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_ResourceARM">
SqlTriggerGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_Resource">SqlTriggerGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetResults_Status">SqlTriggerGetResults_Status</a>)
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Trigger</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL trigger</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>triggerOperation</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetPropertiesStatusResourceTriggerOperation">
SqlTriggerGetPropertiesStatusResourceTriggerOperation
</a>
</em>
</td>
<td>
<p>TriggerOperation: The operation the trigger is associated with</p>
</td>
</tr>
<tr>
<td>
<code>triggerType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetPropertiesStatusResourceTriggerType">
SqlTriggerGetPropertiesStatusResourceTriggerType
</a>
</em>
</td>
<td>
<p>TriggerType: Type of the Trigger</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_ResourceARM">SqlTriggerGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_StatusARM">SqlTriggerGetProperties_StatusARM</a>)
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Trigger</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL trigger</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>triggerOperation</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetPropertiesStatusResourceTriggerOperation">
SqlTriggerGetPropertiesStatusResourceTriggerOperation
</a>
</em>
</td>
<td>
<p>TriggerOperation: The operation the trigger is associated with</p>
</td>
</tr>
<tr>
<td>
<code>triggerType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetPropertiesStatusResourceTriggerType">
SqlTriggerGetPropertiesStatusResourceTriggerType
</a>
</em>
</td>
<td>
<p>TriggerType: Type of the Trigger</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetResults_Status">SqlTriggerGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerTrigger">SqlDatabaseContainerTrigger</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_Status_Resource">
SqlTriggerGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerGetResults_StatusARM">SqlTriggerGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerGetProperties_StatusARM">
SqlTriggerGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB trigger</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerResource">SqlTriggerResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec">DatabaseAccountsSqlDatabasesContainersTriggers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerResource</a></p>
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Trigger</p>
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
<p>Id: Name of the Cosmos DB SQL trigger</p>
</td>
</tr>
<tr>
<td>
<code>triggerOperation</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceTriggerOperation">
SqlTriggerResourceTriggerOperation
</a>
</em>
</td>
<td>
<p>TriggerOperation: The operation the trigger is associated with.</p>
</td>
</tr>
<tr>
<td>
<code>triggerType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceTriggerType">
SqlTriggerResourceTriggerType
</a>
</em>
</td>
<td>
<p>TriggerType: Type of the Trigger.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerResourceARM">SqlTriggerResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerCreateUpdatePropertiesARM">SqlTriggerCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerResource</a></p>
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the Trigger</p>
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
<p>Id: Name of the Cosmos DB SQL trigger</p>
</td>
</tr>
<tr>
<td>
<code>triggerOperation</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceTriggerOperation">
SqlTriggerResourceTriggerOperation
</a>
</em>
</td>
<td>
<p>TriggerOperation: The operation the trigger is associated with.</p>
</td>
</tr>
<tr>
<td>
<code>triggerType</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceTriggerType">
SqlTriggerResourceTriggerType
</a>
</em>
</td>
<td>
<p>TriggerType: Type of the Trigger.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerResourceTriggerOperation">SqlTriggerResourceTriggerOperation
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResource">SqlTriggerResource</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceARM">SqlTriggerResourceARM</a>)
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
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Replace&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlTriggerResourceTriggerType">SqlTriggerResourceTriggerType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResource">SqlTriggerResource</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlTriggerResourceARM">SqlTriggerResourceARM</a>)
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
<tbody><tr><td><p>&#34;Post&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Pre&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionCreateUpdatePropertiesARM">SqlUserDefinedFunctionCreateUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM">DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionCreateUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionCreateUpdateProperties</a></p>
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
<code>options</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.CreateUpdateOptionsARM">
CreateUpdateOptionsARM
</a>
</em>
</td>
<td>
<p>Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are &ldquo;If-Match&rdquo;,
&ldquo;If-None-Match&rdquo;, &ldquo;Session-Token&rdquo; and &ldquo;Throughput&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionResourceARM">
SqlUserDefinedFunctionResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB SQL userDefinedFunction resource object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_StatusARM">SqlUserDefinedFunctionGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetResults_StatusARM">SqlUserDefinedFunctionGetResults_StatusARM</a>)
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
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_Status_ResourceARM">
SqlUserDefinedFunctionGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_Status_Resource">SqlUserDefinedFunctionGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetResults_Status">SqlUserDefinedFunctionGetResults_Status</a>)
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the User Defined Function</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL userDefinedFunction</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_Status_ResourceARM">SqlUserDefinedFunctionGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_StatusARM">SqlUserDefinedFunctionGetProperties_StatusARM</a>)
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the User Defined Function</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
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
<p>Id: Name of the Cosmos DB SQL userDefinedFunction</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetResults_Status">SqlUserDefinedFunctionGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerUserDefinedFunction">SqlDatabaseContainerUserDefinedFunction</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_Status_Resource">
SqlUserDefinedFunctionGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetResults_StatusARM">SqlUserDefinedFunctionGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionGetProperties_StatusARM">
SqlUserDefinedFunctionGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB userDefinedFunction</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionResource">SqlUserDefinedFunctionResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec">DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource</a></p>
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the User Defined Function</p>
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
<p>Id: Name of the Cosmos DB SQL userDefinedFunction</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionResourceARM">SqlUserDefinedFunctionResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlUserDefinedFunctionCreateUpdatePropertiesARM">SqlUserDefinedFunctionCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource</a></p>
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
<code>body</code><br/>
<em>
string
</em>
</td>
<td>
<p>Body: Body of the User Defined Function</p>
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
<p>Id: Name of the Cosmos DB SQL userDefinedFunction</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputPolicyResource">ThroughputPolicyResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource">AutoUpgradePolicyResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputPolicyResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputPolicyResource</a></p>
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
<code>incrementPercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>IncrementPercent: Represents the percentage by which throughput can increase every time throughput policy kicks in.</p>
</td>
</tr>
<tr>
<td>
<code>isEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsEnabled: Determines whether the ThroughputPolicy is active or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputPolicyResourceARM">ThroughputPolicyResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResourceARM">AutoUpgradePolicyResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputPolicyResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputPolicyResource</a></p>
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
<code>incrementPercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>IncrementPercent: Represents the percentage by which throughput can increase every time throughput policy kicks in.</p>
</td>
</tr>
<tr>
<td>
<code>isEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsEnabled: Determines whether the ThroughputPolicy is active or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputPolicyResource_Status">ThroughputPolicyResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource_Status">AutoUpgradePolicyResource_Status</a>)
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
<code>incrementPercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>IncrementPercent: Represents the percentage by which throughput can increase every time throughput policy kicks in.</p>
</td>
</tr>
<tr>
<td>
<code>isEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsEnabled: Determines whether the ThroughputPolicy is active or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputPolicyResource_StatusARM">ThroughputPolicyResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.AutoUpgradePolicyResource_StatusARM">AutoUpgradePolicyResource_StatusARM</a>)
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
<code>incrementPercent</code><br/>
<em>
int
</em>
</td>
<td>
<p>IncrementPercent: Represents the percentage by which throughput can increase every time throughput policy kicks in.</p>
</td>
</tr>
<tr>
<td>
<code>isEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsEnabled: Determines whether the ThroughputPolicy is active or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_StatusARM">ThroughputSettingsGetProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_StatusARM">ThroughputSettingsGetResults_StatusARM</a>)
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
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_Status_ResourceARM">
ThroughputSettingsGetProperties_Status_ResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_Status_Resource">ThroughputSettingsGetProperties_Status_Resource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_Status">ThroughputSettingsGetResults_Status</a>)
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource_Status">
AutoscaleSettingsResource_Status
</a>
</em>
</td>
<td>
<p>AutoscaleSettings: Cosmos DB resource for autoscale settings. Either throughput is required or autoscaleSettings is
required, but not both.</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
</td>
</tr>
<tr>
<td>
<code>minimumThroughput</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinimumThroughput: The minimum throughput of the resource</p>
</td>
</tr>
<tr>
<td>
<code>offerReplacePending</code><br/>
<em>
string
</em>
</td>
<td>
<p>OfferReplacePending: The throughput replace is pending</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Value of the Cosmos DB resource throughput. Either throughput is required or autoscaleSettings is required,
but not both.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_Status_ResourceARM">ThroughputSettingsGetProperties_Status_ResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_StatusARM">ThroughputSettingsGetProperties_StatusARM</a>)
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource_StatusARM">
AutoscaleSettingsResource_StatusARM
</a>
</em>
</td>
<td>
<p>AutoscaleSettings: Cosmos DB resource for autoscale settings. Either throughput is required or autoscaleSettings is
required, but not both.</p>
</td>
</tr>
<tr>
<td>
<code>_etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: A system generated property representing the resource etag required for optimistic concurrency control.</p>
</td>
</tr>
<tr>
<td>
<code>minimumThroughput</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinimumThroughput: The minimum throughput of the resource</p>
</td>
</tr>
<tr>
<td>
<code>offerReplacePending</code><br/>
<em>
string
</em>
</td>
<td>
<p>OfferReplacePending: The throughput replace is pending</p>
</td>
</tr>
<tr>
<td>
<code>_rid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Rid: A system generated property. A unique identifier.</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Value of the Cosmos DB resource throughput. Either throughput is required or autoscaleSettings is required,
but not both.</p>
</td>
</tr>
<tr>
<td>
<code>_ts</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Ts: A system generated property that denotes the last updated timestamp of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_Status">ThroughputSettingsGetResults_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.MongodbDatabaseCollectionThroughputSetting">MongodbDatabaseCollectionThroughputSetting</a>, <a href="#documentdb.azure.com/v1beta20210515.MongodbDatabaseThroughputSetting">MongodbDatabaseThroughputSetting</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseContainerThroughputSetting">SqlDatabaseContainerThroughputSetting</a>, <a href="#documentdb.azure.com/v1beta20210515.SqlDatabaseThroughputSetting">SqlDatabaseThroughputSetting</a>)
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_Status_Resource">
ThroughputSettingsGetProperties_Status_Resource
</a>
</em>
</td>
<td>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsGetResults_StatusARM">ThroughputSettingsGetResults_StatusARM
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
<p>Id: The unique resource identifier of the ARM resource.</p>
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
<p>Location: The location of the resource group to which the resource belongs.</p>
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
<p>Name: The name of the ARM resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsGetProperties_StatusARM">
ThroughputSettingsGetProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an Azure Cosmos DB resource throughput</p>
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
<p>Type: The type of Azure resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsResource">ThroughputSettingsResource
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec">DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec">DatabaseAccountsMongodbDatabasesThroughputSettings_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec">DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSettings_Spec">DatabaseAccountsSqlDatabasesThroughputSettings_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsResource</a></p>
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResource">
AutoscaleSettingsResource
</a>
</em>
</td>
<td>
<p>AutoscaleSettings: Cosmos DB provisioned throughput settings object</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Value of the Cosmos DB resource throughput. Either throughput is required or autoscaleSettings is required,
but not both.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsResourceARM">ThroughputSettingsResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsUpdatePropertiesARM">ThroughputSettingsUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsResource">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsResource</a></p>
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
<code>autoscaleSettings</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.AutoscaleSettingsResourceARM">
AutoscaleSettingsResourceARM
</a>
</em>
</td>
<td>
<p>AutoscaleSettings: Cosmos DB provisioned throughput settings object</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Value of the Cosmos DB resource throughput. Either throughput is required or autoscaleSettings is required,
but not both.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.ThroughputSettingsUpdatePropertiesARM">ThroughputSettingsUpdatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM">DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM">DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM">DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM</a>, <a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM">DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsUpdateProperties">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsUpdateProperties</a></p>
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
<code>resource</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.ThroughputSettingsResourceARM">
ThroughputSettingsResourceARM
</a>
</em>
</td>
<td>
<p>Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
both.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKey">UniqueKey
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicy">UniqueKeyPolicy</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey</a></p>
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
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths must be unique for each document in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKeyARM">UniqueKeyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicyARM">UniqueKeyPolicyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey</a></p>
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
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths must be unique for each document in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKeyPolicy">UniqueKeyPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResource">SqlContainerResource</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy</a></p>
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
<code>uniqueKeys</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKey">
[]UniqueKey
</a>
</em>
</td>
<td>
<p>UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKeyPolicyARM">UniqueKeyPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerResourceARM">SqlContainerResourceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy</a></p>
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
<code>uniqueKeys</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKeyARM">
[]UniqueKeyARM
</a>
</em>
</td>
<td>
<p>UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKeyPolicy_Status">UniqueKeyPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_Resource">SqlContainerGetProperties_Status_Resource</a>)
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
<code>uniqueKeys</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKey_Status">
[]UniqueKey_Status
</a>
</em>
</td>
<td>
<p>UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKeyPolicy_StatusARM">UniqueKeyPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.SqlContainerGetProperties_Status_ResourceARM">SqlContainerGetProperties_Status_ResourceARM</a>)
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
<code>uniqueKeys</code><br/>
<em>
<a href="#documentdb.azure.com/v1beta20210515.UniqueKey_StatusARM">
[]UniqueKey_StatusARM
</a>
</em>
</td>
<td>
<p>UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
Cosmos DB service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKey_Status">UniqueKey_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicy_Status">UniqueKeyPolicy_Status</a>)
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
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths must be unique for each document in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.UniqueKey_StatusARM">UniqueKey_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.UniqueKeyPolicy_StatusARM">UniqueKeyPolicy_StatusARM</a>)
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
<code>paths</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Paths: List of paths must be unique for each document in the Azure Cosmos DB service</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.VirtualNetworkRule">VirtualNetworkRule
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccounts_Spec">DatabaseAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/VirtualNetworkRule">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/VirtualNetworkRule</a></p>
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
<code>ignoreMissingVNetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Resource ID of a subnet, for example:
/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.VirtualNetworkRuleARM">VirtualNetworkRuleARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountCreateUpdatePropertiesARM">DatabaseAccountCreateUpdatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/VirtualNetworkRule">https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/VirtualNetworkRule</a></p>
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
</td>
</tr>
<tr>
<td>
<code>ignoreMissingVNetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.VirtualNetworkRule_Status">VirtualNetworkRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetResults_Status">DatabaseAccountGetResults_Status</a>)
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
<p>Id: Resource ID of a subnet, for example:
/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.</p>
</td>
</tr>
<tr>
<td>
<code>ignoreMissingVNetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="documentdb.azure.com/v1beta20210515.VirtualNetworkRule_StatusARM">VirtualNetworkRule_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#documentdb.azure.com/v1beta20210515.DatabaseAccountGetProperties_StatusARM">DatabaseAccountGetProperties_StatusARM</a>)
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
<p>Id: Resource ID of a subnet, for example:
/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.</p>
</td>
</tr>
<tr>
<td>
<code>ignoreMissingVNetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVNetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
