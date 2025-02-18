---
---
<h2 id="servicebus.azure.com/v1beta20210101preview">servicebus.azure.com/v1beta20210101preview</h2>
<div>
<p>Package v1beta20210101preview contains API Schema definitions for the servicebus v1beta20210101preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="servicebus.azure.com/v1beta20210101preview.DictionaryValue_Status">DictionaryValue_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Identity_Status">Identity_Status</a>)
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
<p>ClientId: Client Id of user assigned identity</p>
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
<p>PrincipalId: Principal Id of user assigned identity</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.DictionaryValue_StatusARM">DictionaryValue_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Identity_StatusARM">Identity_StatusARM</a>)
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
<p>ClientId: Client Id of user assigned identity</p>
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
<p>PrincipalId: Principal Id of user assigned identity</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Encryption">Encryption
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_Spec">Namespaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Encryption">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Encryption</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keySource</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EncryptionKeySource">
EncryptionKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultProperties">
[]KeyVaultProperties
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.EncryptionARM">EncryptionARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_Spec_PropertiesARM">Namespaces_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Encryption">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Encryption</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keySource</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EncryptionKeySource">
EncryptionKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultPropertiesARM">
[]KeyVaultPropertiesARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.EncryptionKeySource">EncryptionKeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Encryption">Encryption</a>, <a href="#servicebus.azure.com/v1beta20210101preview.EncryptionARM">EncryptionARM</a>)
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
<tbody><tr><td><p>&#34;Microsoft.KeyVault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.EncryptionStatusKeySource">EncryptionStatusKeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Encryption_Status">Encryption_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.Encryption_StatusARM">Encryption_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Microsoft.KeyVault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Encryption_Status">Encryption_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">SBNamespace_Status</a>)
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
<code>keySource</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EncryptionStatusKeySource">
EncryptionStatusKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultProperties_Status">
[]KeyVaultProperties_Status
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Encryption_StatusARM">Encryption_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespaceProperties_StatusARM">SBNamespaceProperties_StatusARM</a>)
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
<code>keySource</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EncryptionStatusKeySource">
EncryptionStatusKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultProperties_StatusARM">
[]KeyVaultProperties_StatusARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.EntityStatus_Status">EntityStatus_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBQueueProperties_StatusARM">SBQueueProperties_StatusARM</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBQueue_Status">SBQueue_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBTopicProperties_StatusARM">SBTopicProperties_StatusARM</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBTopic_Status">SBTopic_Status</a>)
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
<tbody><tr><td><p>&#34;Active&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReceiveDisabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Renaming&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Restoring&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SendDisabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unknown&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Identity">Identity
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_Spec">Namespaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Identity">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Identity</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.IdentityType">
IdentityType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.IdentityARM">IdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_SpecARM">Namespaces_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Identity">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/Identity</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.IdentityType">
IdentityType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.IdentityStatusType">IdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Identity_Status">Identity_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.Identity_StatusARM">Identity_StatusARM</a>)
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
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.IdentityType">IdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Identity">Identity</a>, <a href="#servicebus.azure.com/v1beta20210101preview.IdentityARM">IdentityARM</a>)
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
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Identity_Status">Identity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">SBNamespace_Status</a>)
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
<p>PrincipalId: ObjectId from the KeyVault</p>
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
<p>TenantId: TenantId from the KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.IdentityStatusType">
IdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.DictionaryValue_Status">
map[string]./api/servicebus/v1beta20210101preview.DictionaryValue_Status
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Identity_StatusARM">Identity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_StatusARM">SBNamespace_StatusARM</a>)
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
<p>PrincipalId: ObjectId from the KeyVault</p>
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
<p>TenantId: TenantId from the KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.IdentityStatusType">
IdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.DictionaryValue_StatusARM">
map[string]./api/servicebus/v1beta20210101preview.DictionaryValue_StatusARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.KeyVaultProperties">KeyVaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Encryption">Encryption</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/KeyVaultProperties</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityProperties">
UserAssignedIdentityProperties
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.KeyVaultPropertiesARM">KeyVaultPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.EncryptionARM">EncryptionARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/KeyVaultProperties</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityPropertiesARM">
UserAssignedIdentityPropertiesARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.KeyVaultProperties_Status">KeyVaultProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Encryption_Status">Encryption_Status</a>)
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
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityProperties_Status">
UserAssignedIdentityProperties_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.KeyVaultProperties_StatusARM">KeyVaultProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Encryption_StatusARM">Encryption_StatusARM</a>)
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
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityProperties_StatusARM">
UserAssignedIdentityProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.MessageCountDetails_Status">MessageCountDetails_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBQueue_Status">SBQueue_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBTopic_Status">SBTopic_Status</a>)
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
<code>activeMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ActiveMessageCount: Number of active messages in the queue, topic, or subscription.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>DeadLetterMessageCount: Number of messages that are dead lettered.</p>
</td>
</tr>
<tr>
<td>
<code>scheduledMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ScheduledMessageCount: Number of scheduled messages.</p>
</td>
</tr>
<tr>
<td>
<code>transferDeadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferDeadLetterMessageCount: Number of messages transferred into dead letters.</p>
</td>
</tr>
<tr>
<td>
<code>transferMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferMessageCount: Number of messages transferred to another queue, topic, or subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.MessageCountDetails_StatusARM">MessageCountDetails_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBQueueProperties_StatusARM">SBQueueProperties_StatusARM</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBTopicProperties_StatusARM">SBTopicProperties_StatusARM</a>)
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
<code>activeMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ActiveMessageCount: Number of active messages in the queue, topic, or subscription.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>DeadLetterMessageCount: Number of messages that are dead lettered.</p>
</td>
</tr>
<tr>
<td>
<code>scheduledMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ScheduledMessageCount: Number of scheduled messages.</p>
</td>
</tr>
<tr>
<td>
<code>transferDeadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferDeadLetterMessageCount: Number of messages transferred into dead letters.</p>
</td>
</tr>
<tr>
<td>
<code>transferMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferMessageCount: Number of messages transferred to another queue, topic, or subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Namespace">Namespace
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_Spec">
Namespaces_Spec
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
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Properties to configure Encryption</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Identity">
Identity
</a>
</em>
</td>
<td>
<p>Identity: Properties to configure User Assigned Identities for Bring your Own Keys</p>
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
<p>Location: The Geo-location where the resource lives</p>
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
<a href="#servicebus.azure.com/v1beta20210101preview.SBSku">
SBSku
</a>
</em>
</td>
<td>
<p>Sku: SKU of the namespace.</p>
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
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">
SBNamespace_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesQueue">NamespacesQueue
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_queues">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_queues</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.NamespacesQueues_Spec">
NamespacesQueues_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
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
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
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
<a href="#servicebus.azure.com/v1beta20210101preview.SBQueue_Status">
SBQueue_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesQueuesSpecAPIVersion">NamespacesQueuesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-01-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesQueues_Spec">NamespacesQueues_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.NamespacesQueue">NamespacesQueue</a>)
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
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
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
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
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
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesQueues_SpecARM">NamespacesQueues_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBQueuePropertiesARM">
SBQueuePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The Queue Properties definition.</p>
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
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesSpecAPIVersion">NamespacesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-01-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesTopic">NamespacesTopic
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_topics">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_topics</a></p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1beta20210101preview.NamespacesTopics_Spec">
NamespacesTopics_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
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
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
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
<a href="#servicebus.azure.com/v1beta20210101preview.SBTopic_Status">
SBTopic_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesTopicsSpecAPIVersion">NamespacesTopicsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-01-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesTopics_Spec">NamespacesTopics_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.NamespacesTopic">NamespacesTopic</a>)
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
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
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
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
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
<h3 id="servicebus.azure.com/v1beta20210101preview.NamespacesTopics_SpecARM">NamespacesTopics_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBTopicPropertiesARM">
SBTopicPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The Topic Properties definition.</p>
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
<h3 id="servicebus.azure.com/v1beta20210101preview.Namespaces_Spec">Namespaces_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespace">Namespace</a>)
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
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Properties to configure Encryption</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Identity">
Identity
</a>
</em>
</td>
<td>
<p>Identity: Properties to configure User Assigned Identities for Bring your Own Keys</p>
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
<p>Location: The Geo-location where the resource lives</p>
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
<a href="#servicebus.azure.com/v1beta20210101preview.SBSku">
SBSku
</a>
</em>
</td>
<td>
<p>Sku: SKU of the namespace.</p>
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
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.Namespaces_SpecARM">Namespaces_SpecARM
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
<a href="#servicebus.azure.com/v1beta20210101preview.IdentityARM">
IdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Properties to configure User Assigned Identities for Bring your Own Keys</p>
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
<p>Location: The Geo-location where the resource lives</p>
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
<a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_Spec_PropertiesARM">
Namespaces_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuARM">
SBSkuARM
</a>
</em>
</td>
<td>
<p>Sku: SKU of the namespace.</p>
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
<h3 id="servicebus.azure.com/v1beta20210101preview.Namespaces_Spec_PropertiesARM">Namespaces_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_SpecARM">Namespaces_SpecARM</a>)
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
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EncryptionARM">
EncryptionARM
</a>
</em>
</td>
<td>
<p>Encryption: Properties to configure Encryption</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">SBNamespace_Status</a>)
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespaceProperties_StatusARM">SBNamespaceProperties_StatusARM</a>)
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBNamespaceProperties_StatusARM">SBNamespaceProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_StatusARM">SBNamespace_StatusARM</a>)
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
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The time the namespace was created</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Encryption_StatusARM">
Encryption_StatusARM
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
</td>
</tr>
<tr>
<td>
<code>metricId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricId: Identifier for Azure Insights metrics</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
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
<p>ProvisioningState: Provisioning state of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>serviceBusEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
string
</em>
</td>
<td>
<p>Status: Status of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The time the namespace was updated.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">SBNamespace_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespace">Namespace</a>)
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
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The time the namespace was created</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Encryption_Status">
Encryption_Status
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Identity_Status">
Identity_Status
</a>
</em>
</td>
<td>
<p>Identity: Properties of BYOK Identity description</p>
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
<p>Location: The Geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>metricId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricId: Identifier for Azure Insights metrics</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
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
<p>ProvisioningState: Provisioning state of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>serviceBusEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSku_Status">
SBSku_Status
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
string
</em>
</td>
<td>
<p>Status: Status of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Tags: Resource tags</p>
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
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The time the namespace was updated.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBNamespace_StatusARM">SBNamespace_StatusARM
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.Identity_StatusARM">
Identity_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: Properties of BYOK Identity description</p>
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
<p>Location: The Geo-location where the resource lives</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBNamespaceProperties_StatusARM">
SBNamespaceProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSku_StatusARM">
SBSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Tags: Resource tags</p>
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBQueuePropertiesARM">SBQueuePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.NamespacesQueues_SpecARM">NamespacesQueues_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBQueueProperties">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBQueueProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBQueueProperties_StatusARM">SBQueueProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBQueue_StatusARM">SBQueue_StatusARM</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time a message was sent, or the last time there was a receive request to this queue.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.MessageCountDetails_StatusARM">
MessageCountDetails_StatusARM
</a>
</em>
</td>
<td>
<p>CountDetails: Message Count Details.</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>messageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageCount: The number of messages in the queue.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: The size of the queue, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EntityStatus_Status">
EntityStatus_Status
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBQueue_Status">SBQueue_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.NamespacesQueue">NamespacesQueue</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time a message was sent, or the last time there was a receive request to this queue.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.MessageCountDetails_Status">
MessageCountDetails_Status
</a>
</em>
</td>
<td>
<p>CountDetails: Message Count Details.</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>messageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageCount: The number of messages in the queue.</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: The size of the queue, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EntityStatus_Status">
EntityStatus_Status
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBQueue_StatusARM">SBQueue_StatusARM
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
<p>Id: Resource Id</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBQueueProperties_StatusARM">
SBQueueProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Queue Properties</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSku">SBSku
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_Spec">Namespaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBSku">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBSku</a></p>
</div>
<table>
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuName">
SBSkuName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuTier">
SBSkuTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSkuARM">SBSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.Namespaces_SpecARM">Namespaces_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBSku">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBSku</a></p>
</div>
<table>
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuName">
SBSkuName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuTier">
SBSkuTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSkuName">SBSkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBSku">SBSku</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBSkuARM">SBSkuARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSkuStatusName">SBSkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBSku_Status">SBSku_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBSku_StatusARM">SBSku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSkuStatusTier">SBSkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBSku_Status">SBSku_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBSku_StatusARM">SBSku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSkuTier">SBSkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBSku">SBSku</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBSkuARM">SBSkuARM</a>)
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
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSku_Status">SBSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">SBNamespace_Status</a>)
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuStatusName">
SBSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuStatusTier">
SBSkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBSku_StatusARM">SBSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_StatusARM">SBNamespace_StatusARM</a>)
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuStatusName">
SBSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBSkuStatusTier">
SBSkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBTopicPropertiesARM">SBTopicPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.NamespacesTopics_SpecARM">NamespacesTopics_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBTopicProperties">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBTopicProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBTopicProperties_StatusARM">SBTopicProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SBTopic_StatusARM">SBTopic_StatusARM</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time the message was sent, or a request was received, for this topic.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.MessageCountDetails_StatusARM">
MessageCountDetails_StatusARM
</a>
</em>
</td>
<td>
<p>CountDetails: Message count details</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: Exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: Size of the topic, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EntityStatus_Status">
EntityStatus_Status
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionCount: Number of subscriptions.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBTopic_Status">SBTopic_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.NamespacesTopic">NamespacesTopic</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time the message was sent, or a request was received, for this topic.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.MessageCountDetails_Status">
MessageCountDetails_Status
</a>
</em>
</td>
<td>
<p>CountDetails: Message count details</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: Exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: Size of the topic, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.EntityStatus_Status">
EntityStatus_Status
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionCount: Number of subscriptions.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SBTopic_StatusARM">SBTopic_StatusARM
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
<p>Id: Resource Id</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SBTopicProperties_StatusARM">
SBTopicProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of topic resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SystemDataStatusCreatedByType">SystemDataStatusCreatedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SystemData_Status">SystemData_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Application&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Key&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SystemDataStatusLastModifiedByType">SystemDataStatusLastModifiedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.SystemData_Status">SystemData_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Application&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Key&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_Status">SBNamespace_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBQueue_Status">SBQueue_Status</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBTopic_Status">SBTopic_Status</a>)
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
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The timestamp of resource creation (UTC).</p>
</td>
</tr>
<tr>
<td>
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedBy: The identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemDataStatusCreatedByType">
SystemDataStatusCreatedByType
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedAt: The type of identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedBy: The identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemDataStatusLastModifiedByType">
SystemDataStatusLastModifiedByType
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBNamespace_StatusARM">SBNamespace_StatusARM</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBQueue_StatusARM">SBQueue_StatusARM</a>, <a href="#servicebus.azure.com/v1beta20210101preview.SBTopic_StatusARM">SBTopic_StatusARM</a>)
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
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The timestamp of resource creation (UTC).</p>
</td>
</tr>
<tr>
<td>
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedBy: The identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemDataStatusCreatedByType">
SystemDataStatusCreatedByType
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedAt: The type of identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedBy: The identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#servicebus.azure.com/v1beta20210101preview.SystemDataStatusLastModifiedByType">
SystemDataStatusLastModifiedByType
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityProperties">UserAssignedIdentityProperties
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultProperties">KeyVaultProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/UserAssignedIdentityProperties">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/UserAssignedIdentityProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>userAssignedIdentityReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityReference: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityPropertiesARM">UserAssignedIdentityPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultPropertiesARM">KeyVaultPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/UserAssignedIdentityProperties">https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/UserAssignedIdentityProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>userAssignedIdentity</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityProperties_Status">UserAssignedIdentityProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultProperties_Status">KeyVaultProperties_Status</a>)
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
<code>userAssignedIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserAssignedIdentity: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1beta20210101preview.UserAssignedIdentityProperties_StatusARM">UserAssignedIdentityProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1beta20210101preview.KeyVaultProperties_StatusARM">KeyVaultProperties_StatusARM</a>)
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
<code>userAssignedIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserAssignedIdentity: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<hr/>
