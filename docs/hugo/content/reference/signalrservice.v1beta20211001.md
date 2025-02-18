---
---
<h2 id="signalrservice.azure.com/v1beta20211001">signalrservice.azure.com/v1beta20211001</h2>
<div>
<p>Package v1beta20211001 contains API Schema definitions for the signalrservice v1beta20211001 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="signalrservice.azure.com/v1beta20211001.ACLAction_Status">ACLAction_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_Status">SignalRNetworkACLs_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_StatusARM">SignalRNetworkACLs_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Allow&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deny&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.FeatureFlags_Status">FeatureFlags_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature_Status">SignalRFeature_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature_StatusARM">SignalRFeature_StatusARM</a>)
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
<tbody><tr><td><p>&#34;EnableConnectivityLogs&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnableLiveTrace&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnableMessagingLogs&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServiceMode&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentity">ManagedIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentity">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentity</a></p>
</div>
<table>
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
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentityType">
ManagedIdentityType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Get or set the user assigned identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentityARM">ManagedIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_SpecARM">SignalR_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentity">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentity</a></p>
</div>
<table>
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
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentityType">
ManagedIdentityType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Get or set the user assigned identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentitySettings">ManagedIdentitySettings
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings">UpstreamAuthSettings</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentitySettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentitySettings</a></p>
</div>
<table>
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
string
</em>
</td>
<td>
<p>Resource: The Resource indicating the App ID URI of the target resource.
It also appears in the aud (audience) claim of the issued token.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentitySettingsARM">ManagedIdentitySettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsARM">UpstreamAuthSettingsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentitySettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentitySettings</a></p>
</div>
<table>
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
string
</em>
</td>
<td>
<p>Resource: The Resource indicating the App ID URI of the target resource.
It also appears in the aud (audience) claim of the issued token.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentitySettings_Status">ManagedIdentitySettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_Status">UpstreamAuthSettings_Status</a>)
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
string
</em>
</td>
<td>
<p>Resource: The Resource indicating the App ID URI of the target resource.
It also appears in the aud (audience) claim of the issued token.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentitySettings_StatusARM">ManagedIdentitySettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_StatusARM">UpstreamAuthSettings_StatusARM</a>)
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
string
</em>
</td>
<td>
<p>Resource: The Resource indicating the App ID URI of the target resource.
It also appears in the aud (audience) claim of the issued token.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentityType">ManagedIdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity">ManagedIdentity</a>, <a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentityARM">ManagedIdentityARM</a>)
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
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentityType_Status">ManagedIdentityType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity_Status">ManagedIdentity_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity_StatusARM">ManagedIdentity_StatusARM</a>)
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
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentity_Status">ManagedIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<p>PrincipalId: Get the principal id for the system assigned identity.
Only be used in response.</p>
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
<p>TenantId: Get the tenant id for the system assigned identity.
Only be used in response</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentityType_Status">
ManagedIdentityType_Status
</a>
</em>
</td>
<td>
<p>Type: Represent the identity type: systemAssigned, userAssigned, None</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UserAssignedIdentityProperty_Status">
map[string]./api/signalrservice/v1beta20211001.UserAssignedIdentityProperty_Status
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Get or set the user assigned identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ManagedIdentity_StatusARM">ManagedIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_StatusARM">SignalRResource_StatusARM</a>)
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
<p>PrincipalId: Get the principal id for the system assigned identity.
Only be used in response.</p>
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
<p>TenantId: Get the tenant id for the system assigned identity.
Only be used in response</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentityType_Status">
ManagedIdentityType_Status
</a>
</em>
</td>
<td>
<p>Type: Represent the identity type: systemAssigned, userAssigned, None</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UserAssignedIdentityProperty_StatusARM">
map[string]./api/signalrservice/v1beta20211001.UserAssignedIdentityProperty_StatusARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Get or set the user assigned identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.NetworkACL">NetworkACL
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs">SignalRNetworkACLs</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/NetworkACL">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/NetworkACL</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACLAllow">
[]NetworkACLAllow
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACLDeny">
[]NetworkACLDeny
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.NetworkACLARM">NetworkACLARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsARM">SignalRNetworkACLsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/NetworkACL">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/NetworkACL</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACLAllow">
[]NetworkACLAllow
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACLDeny">
[]NetworkACLDeny
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.NetworkACLAllow">NetworkACLAllow
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.NetworkACL">NetworkACL</a>, <a href="#signalrservice.azure.com/v1beta20211001.NetworkACLARM">NetworkACLARM</a>)
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
<tbody><tr><td><p>&#34;ClientConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RESTAPI&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServerConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Trace&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.NetworkACLDeny">NetworkACLDeny
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.NetworkACL">NetworkACL</a>, <a href="#signalrservice.azure.com/v1beta20211001.NetworkACLARM">NetworkACLARM</a>)
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
<tbody><tr><td><p>&#34;ClientConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RESTAPI&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServerConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Trace&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.NetworkACL_Status">NetworkACL_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_Status">SignalRNetworkACLs_Status</a>)
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
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.NetworkACL_StatusARM">NetworkACL_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_StatusARM">SignalRNetworkACLs_StatusARM</a>)
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
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointACL">PrivateEndpointACL
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs">SignalRNetworkACLs</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/PrivateEndpointACL">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/PrivateEndpointACL</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLAllow">
[]PrivateEndpointACLAllow
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLDeny">
[]PrivateEndpointACLDeny
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
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
<p>Name: Name of the private endpoint connection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointACLARM">PrivateEndpointACLARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsARM">SignalRNetworkACLsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/PrivateEndpointACL">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/PrivateEndpointACL</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLAllow">
[]PrivateEndpointACLAllow
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLDeny">
[]PrivateEndpointACLDeny
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
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
<p>Name: Name of the private endpoint connection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointACLAllow">PrivateEndpointACLAllow
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL">PrivateEndpointACL</a>, <a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLARM">PrivateEndpointACLARM</a>)
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
<tbody><tr><td><p>&#34;ClientConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RESTAPI&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServerConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Trace&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointACLDeny">PrivateEndpointACLDeny
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL">PrivateEndpointACL</a>, <a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLARM">PrivateEndpointACLARM</a>)
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
<tbody><tr><td><p>&#34;ClientConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RESTAPI&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServerConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Trace&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointACL_Status">PrivateEndpointACL_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_Status">SignalRNetworkACLs_Status</a>)
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
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
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
<p>Name: Name of the private endpoint connection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointACL_StatusARM">PrivateEndpointACL_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_StatusARM">SignalRNetworkACLs_StatusARM</a>)
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
<code>allow</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
</td>
</tr>
<tr>
<td>
<code>deny</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">
[]SignalRRequestType_Status
</a>
</em>
</td>
<td>
<p>Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.</p>
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
<p>Name: Name of the private endpoint connection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded">PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<p>Id: Fully qualified resource Id for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<p>Id: Fully qualified resource Id for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ProvisioningState_Status">ProvisioningState_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
</tr><tr><td><p>&#34;Moving&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unknown&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogCategory">ResourceLogCategory
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration">ResourceLogConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogCategory">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogCategory</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>Enabled: Indicates whether or the resource log category is enabled.
Available values: true, false.
Case insensitive.</p>
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
<p>Name: Gets or sets the resource log category&rsquo;s name.
Available values: ConnectivityLogs, MessagingLogs.
Case insensitive.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogCategoryARM">ResourceLogCategoryARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfigurationARM">ResourceLogConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogCategory">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogCategory</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>Enabled: Indicates whether or the resource log category is enabled.
Available values: true, false.
Case insensitive.</p>
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
<p>Name: Gets or sets the resource log category&rsquo;s name.
Available values: ConnectivityLogs, MessagingLogs.
Case insensitive.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogCategory_Status">ResourceLogCategory_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration_Status">ResourceLogConfiguration_Status</a>)
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
<code>enabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>Enabled: Indicates whether or the resource log category is enabled.
Available values: true, false.
Case insensitive.</p>
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
<p>Name: Gets or sets the resource log category&rsquo;s name.
Available values: ConnectivityLogs, MessagingLogs.
Case insensitive.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogCategory_StatusARM">ResourceLogCategory_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration_StatusARM">ResourceLogConfiguration_StatusARM</a>)
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
<code>enabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>Enabled: Indicates whether or the resource log category is enabled.
Available values: true, false.
Case insensitive.</p>
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
<p>Name: Gets or sets the resource log category&rsquo;s name.
Available values: ConnectivityLogs, MessagingLogs.
Case insensitive.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration">ResourceLogConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogConfiguration">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>categories</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogCategory">
[]ResourceLogCategory
</a>
</em>
</td>
<td>
<p>Categories: Gets or sets the list of category configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogConfigurationARM">ResourceLogConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogConfiguration">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogConfiguration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>categories</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogCategoryARM">
[]ResourceLogCategoryARM
</a>
</em>
</td>
<td>
<p>Categories: Gets or sets the list of category configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration_Status">ResourceLogConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<code>categories</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogCategory_Status">
[]ResourceLogCategory_Status
</a>
</em>
</td>
<td>
<p>Categories: Gets or sets the list of category configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration_StatusARM">ResourceLogConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<code>categories</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogCategory_StatusARM">
[]ResourceLogCategory_StatusARM
</a>
</em>
</td>
<td>
<p>Categories: Gets or sets the list of category configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceSku">ResourceSku
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceSku">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceSku</a></p>
</div>
<table>
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
<p>Capacity: Optional, integer. The unit count of the resource. 1 by default.
If present, following values are allowed:
Free: 1
Standard: 1,2,5,10,20,50,100</p>
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
<p>Name: The name of the SKU. Required.
Allowed values: Standard_S1, Free_F1</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSkuTier">
ResourceSkuTier
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceSkuARM">ResourceSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_SpecARM">SignalR_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceSku">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceSku</a></p>
</div>
<table>
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
<p>Capacity: Optional, integer. The unit count of the resource. 1 by default.
If present, following values are allowed:
Free: 1
Standard: 1,2,5,10,20,50,100</p>
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
<p>Name: The name of the SKU. Required.
Allowed values: Standard_S1, Free_F1</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSkuTier">
ResourceSkuTier
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceSkuTier">ResourceSkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ResourceSku">ResourceSku</a>, <a href="#signalrservice.azure.com/v1beta20211001.ResourceSkuARM">ResourceSkuARM</a>)
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
</tr><tr><td><p>&#34;Free&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceSku_Status">ResourceSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<p>Capacity: Optional, integer. The unit count of the resource. 1 by default.
If present, following values are allowed:
Free: 1
Standard: 1,2,5,10,20,50,100</p>
</td>
</tr>
<tr>
<td>
<code>family</code><br/>
<em>
string
</em>
</td>
<td>
<p>Family: Not used. Retained for future use.</p>
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
<p>Name: The name of the SKU. Required.
Allowed values: Standard_S1, Free_F1</p>
</td>
</tr>
<tr>
<td>
<code>size</code><br/>
<em>
string
</em>
</td>
<td>
<p>Size: Not used. Retained for future use.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRSkuTier_Status">
SignalRSkuTier_Status
</a>
</em>
</td>
<td>
<p>Tier: Optional tier of this particular SKU. &lsquo;Standard&rsquo; or &lsquo;Free&rsquo;.
<code>Basic</code> is deprecated, use <code>Standard</code> instead.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ResourceSku_StatusARM">ResourceSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_StatusARM">SignalRResource_StatusARM</a>)
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
<p>Capacity: Optional, integer. The unit count of the resource. 1 by default.
If present, following values are allowed:
Free: 1
Standard: 1,2,5,10,20,50,100</p>
</td>
</tr>
<tr>
<td>
<code>family</code><br/>
<em>
string
</em>
</td>
<td>
<p>Family: Not used. Retained for future use.</p>
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
<p>Name: The name of the SKU. Required.
Allowed values: Standard_S1, Free_F1</p>
</td>
</tr>
<tr>
<td>
<code>size</code><br/>
<em>
string
</em>
</td>
<td>
<p>Size: Not used. Retained for future use.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRSkuTier_Status">
SignalRSkuTier_Status
</a>
</em>
</td>
<td>
<p>Tier: Optional tier of this particular SKU. &lsquo;Standard&rsquo; or &lsquo;Free&rsquo;.
<code>Basic</code> is deprecated, use <code>Standard</code> instead.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings">ServerlessUpstreamSettings
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ServerlessUpstreamSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ServerlessUpstreamSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>templates</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplate">
[]UpstreamTemplate
</a>
</em>
</td>
<td>
<p>Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettingsARM">ServerlessUpstreamSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ServerlessUpstreamSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ServerlessUpstreamSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>templates</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplateARM">
[]UpstreamTemplateARM
</a>
</em>
</td>
<td>
<p>Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings_Status">ServerlessUpstreamSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<code>templates</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplate_Status">
[]UpstreamTemplate_Status
</a>
</em>
</td>
<td>
<p>Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings_StatusARM">ServerlessUpstreamSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<code>templates</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplate_StatusARM">
[]UpstreamTemplate_StatusARM
</a>
</em>
</td>
<td>
<p>Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.ServiceKind_Status">ServiceKind_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_StatusARM">SignalRResource_StatusARM</a>)
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
<tbody><tr><td><p>&#34;RawWebSockets&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SignalR&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded">SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<p>Id: Fully qualified resource Id for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM">SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<p>Id: Fully qualified resource Id for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalR">SignalR
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/resourceDefinitions/signalR">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/resourceDefinitions/signalR</a></p>
</div>
<table>
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
<a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">
SignalR_Spec
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
<code>cors</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRCorsSettings">
SignalRCorsSettings
</a>
</em>
</td>
<td>
<p>Cors: Cross-Origin Resource Sharing (CORS) settings.</p>
</td>
</tr>
<tr>
<td>
<code>disableAadAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAadAuth: DisableLocalAuth
Enable or disable aad auth
When set as true, connection with AuthType=aad won&rsquo;t work.</p>
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
<p>DisableLocalAuth: DisableLocalAuth
Enable or disable local auth with AccessKey
When set as true, connection with AccessKey=xxx won&rsquo;t work.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature">
[]SignalRFeature
</a>
</em>
</td>
<td>
<p>Features: List of the featureFlags.
FeatureFlags that are not included in the parameters for the update operation will not be modified.
And the response will only include featureFlags that are explicitly set.
When a featureFlag is not explicitly set, its globally default value will be used
But keep in mind, the default value doesn&rsquo;t mean &ldquo;false&rdquo;. It varies in terms of different FeatureFlags.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity">
ManagedIdentity
</a>
</em>
</td>
<td>
<p>Identity: A class represent managed identities used for request and response</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRSpecKind">
SignalRSpecKind
</a>
</em>
</td>
<td>
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
<p>Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.</p>
</td>
</tr>
<tr>
<td>
<code>networkACLs</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs">
SignalRNetworkACLs
</a>
</em>
</td>
<td>
<p>NetworkACLs: Network ACLs for the resource</p>
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
string
</em>
</td>
<td>
<p>PublicNetworkAccess: Enable or disable public network access. Default to &ldquo;Enabled&rdquo;.
When it&rsquo;s Enabled, network ACLs still apply.
When it&rsquo;s Disabled, public network access is always disabled no matter what you set in network ACLs.</p>
</td>
</tr>
<tr>
<td>
<code>resourceLogConfiguration</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration">
ResourceLogConfiguration
</a>
</em>
</td>
<td>
<p>ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSku">
ResourceSku
</a>
</em>
</td>
<td>
<p>Sku: The billing information of the resource.</p>
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
<p>Tags: Tags of the service which is a list of key value pairs that describe the resource.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRTlsSettings">
SignalRTlsSettings
</a>
</em>
</td>
<td>
<p>Tls: TLS settings for the resource</p>
</td>
</tr>
<tr>
<td>
<code>upstream</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings">
ServerlessUpstreamSettings
</a>
</em>
</td>
<td>
<p>Upstream: The settings for the Upstream when the service is in server-less mode.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">
SignalRResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRCorsSettings">SignalRCorsSettings
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRCorsSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRCorsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedOrigins</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
<a href="http://example.com:12345)">http://example.com:12345)</a>. Use &ldquo;*&rdquo; to allow all. If omitted, allow all by default.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRCorsSettingsARM">SignalRCorsSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRCorsSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRCorsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedOrigins</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
<a href="http://example.com:12345)">http://example.com:12345)</a>. Use &ldquo;*&rdquo; to allow all. If omitted, allow all by default.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRCorsSettings_Status">SignalRCorsSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<code>allowedOrigins</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
<a href="http://example.com:12345)">http://example.com:12345)</a>. Use &ldquo;*&rdquo; to allow all. If omitted, allow all by default.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRCorsSettings_StatusARM">SignalRCorsSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<code>allowedOrigins</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
<a href="http://example.com:12345)">http://example.com:12345)</a>. Use &ldquo;*&rdquo; to allow all. If omitted, allow all by default.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRFeature">SignalRFeature
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRFeature">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRFeature</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>flag</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeatureFlag">
SignalRFeatureFlag
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: Optional properties related to this feature.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the feature flag. See Azure SignalR service document <a href="https://docs.microsoft.com/azure/azure-signalr/">https://docs.microsoft.com/azure/azure-signalr/</a> for
allowed values.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRFeatureARM">SignalRFeatureARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRFeature">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRFeature</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>flag</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeatureFlag">
SignalRFeatureFlag
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: Optional properties related to this feature.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the feature flag. See Azure SignalR service document <a href="https://docs.microsoft.com/azure/azure-signalr/">https://docs.microsoft.com/azure/azure-signalr/</a> for
allowed values.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRFeatureFlag">SignalRFeatureFlag
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature">SignalRFeature</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRFeatureARM">SignalRFeatureARM</a>)
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
<tbody><tr><td><p>&#34;EnableConnectivityLogs&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnableLiveTrace&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EnableMessagingLogs&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServiceMode&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRFeature_Status">SignalRFeature_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<code>flag</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.FeatureFlags_Status">
FeatureFlags_Status
</a>
</em>
</td>
<td>
<p>Flag: FeatureFlags is the supported features of Azure SignalR service.
- ServiceMode: Flag for backend server for SignalR service. Values allowed: &ldquo;Default&rdquo;: have your own backend server;
&ldquo;Serverless&rdquo;: your application doesn&rsquo;t have a backend server; &ldquo;Classic&rdquo;: for backward compatibility. Support both
Default and Serverless mode but not recommended; &ldquo;PredefinedOnly&rdquo;: for future use.
- EnableConnectivityLogs: &ldquo;true&rdquo;/&ldquo;false&rdquo;, to enable/disable the connectivity log category respectively.
- EnableMessagingLogs: &ldquo;true&rdquo;/&ldquo;false&rdquo;, to enable/disable the connectivity log category respectively.
- EnableLiveTrace: Live Trace allows you to know what&rsquo;s happening inside Azure SignalR service, it will give you live
traces in real time, it will be helpful when you developing your own Azure SignalR based web application or
self-troubleshooting some issues. Please note that live traces are counted as outbound messages that will be charged.
Values allowed: &ldquo;true&rdquo;/&ldquo;false&rdquo;, to enable/disable live trace feature.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: Optional properties related to this feature.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the feature flag. See Azure SignalR service document <a href="https://docs.microsoft.com/azure/azure-signalr/">https://docs.microsoft.com/azure/azure-signalr/</a> for
allowed values.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRFeature_StatusARM">SignalRFeature_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<code>flag</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.FeatureFlags_Status">
FeatureFlags_Status
</a>
</em>
</td>
<td>
<p>Flag: FeatureFlags is the supported features of Azure SignalR service.
- ServiceMode: Flag for backend server for SignalR service. Values allowed: &ldquo;Default&rdquo;: have your own backend server;
&ldquo;Serverless&rdquo;: your application doesn&rsquo;t have a backend server; &ldquo;Classic&rdquo;: for backward compatibility. Support both
Default and Serverless mode but not recommended; &ldquo;PredefinedOnly&rdquo;: for future use.
- EnableConnectivityLogs: &ldquo;true&rdquo;/&ldquo;false&rdquo;, to enable/disable the connectivity log category respectively.
- EnableMessagingLogs: &ldquo;true&rdquo;/&ldquo;false&rdquo;, to enable/disable the connectivity log category respectively.
- EnableLiveTrace: Live Trace allows you to know what&rsquo;s happening inside Azure SignalR service, it will give you live
traces in real time, it will be helpful when you developing your own Azure SignalR based web application or
self-troubleshooting some issues. Please note that live traces are counted as outbound messages that will be charged.
Values allowed: &ldquo;true&rdquo;/&ldquo;false&rdquo;, to enable/disable live trace feature.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: Optional properties related to this feature.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the feature flag. See Azure SignalR service document <a href="https://docs.microsoft.com/azure/azure-signalr/">https://docs.microsoft.com/azure/azure-signalr/</a> for
allowed values.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs">SignalRNetworkACLs
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRNetworkACLs">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRNetworkACLs</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsDefaultAction">
SignalRNetworkACLsDefaultAction
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateEndpoints</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL">
[]PrivateEndpointACL
</a>
</em>
</td>
<td>
<p>PrivateEndpoints: ACLs for requests from private endpoints</p>
</td>
</tr>
<tr>
<td>
<code>publicNetwork</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACL">
NetworkACL
</a>
</em>
</td>
<td>
<p>PublicNetwork: Network ACL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsARM">SignalRNetworkACLsARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRNetworkACLs">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRNetworkACLs</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsDefaultAction">
SignalRNetworkACLsDefaultAction
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateEndpoints</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACLARM">
[]PrivateEndpointACLARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoints: ACLs for requests from private endpoints</p>
</td>
</tr>
<tr>
<td>
<code>publicNetwork</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACLARM">
NetworkACLARM
</a>
</em>
</td>
<td>
<p>PublicNetwork: Network ACL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsDefaultAction">SignalRNetworkACLsDefaultAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs">SignalRNetworkACLs</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsARM">SignalRNetworkACLsARM</a>)
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
<tbody><tr><td><p>&#34;Allow&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deny&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_Status">SignalRNetworkACLs_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<code>defaultAction</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ACLAction_Status">
ACLAction_Status
</a>
</em>
</td>
<td>
<p>DefaultAction: Default action when no other rule matches</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoints</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL_Status">
[]PrivateEndpointACL_Status
</a>
</em>
</td>
<td>
<p>PrivateEndpoints: ACLs for requests from private endpoints</p>
</td>
</tr>
<tr>
<td>
<code>publicNetwork</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACL_Status">
NetworkACL_Status
</a>
</em>
</td>
<td>
<p>PublicNetwork: ACL for requests from public network</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_StatusARM">SignalRNetworkACLs_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<code>defaultAction</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ACLAction_Status">
ACLAction_Status
</a>
</em>
</td>
<td>
<p>DefaultAction: Default action when no other rule matches</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoints</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL_StatusARM">
[]PrivateEndpointACL_StatusARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoints: ACLs for requests from private endpoints</p>
</td>
</tr>
<tr>
<td>
<code>publicNetwork</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.NetworkACL_StatusARM">
NetworkACL_StatusARM
</a>
</em>
</td>
<td>
<p>PublicNetwork: ACL for requests from public network</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_SpecARM">SignalR_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRProperties">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cors</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRCorsSettingsARM">
SignalRCorsSettingsARM
</a>
</em>
</td>
<td>
<p>Cors: Cross-Origin Resource Sharing (CORS) settings.</p>
</td>
</tr>
<tr>
<td>
<code>disableAadAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAadAuth: DisableLocalAuth
Enable or disable aad auth
When set as true, connection with AuthType=aad won&rsquo;t work.</p>
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
<p>DisableLocalAuth: DisableLocalAuth
Enable or disable local auth with AccessKey
When set as true, connection with AccessKey=xxx won&rsquo;t work.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeatureARM">
[]SignalRFeatureARM
</a>
</em>
</td>
<td>
<p>Features: List of the featureFlags.
FeatureFlags that are not included in the parameters for the update operation will not be modified.
And the response will only include featureFlags that are explicitly set.
When a featureFlag is not explicitly set, its globally default value will be used
But keep in mind, the default value doesn&rsquo;t mean &ldquo;false&rdquo;. It varies in terms of different FeatureFlags.</p>
</td>
</tr>
<tr>
<td>
<code>networkACLs</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLsARM">
SignalRNetworkACLsARM
</a>
</em>
</td>
<td>
<p>NetworkACLs: Network ACLs for the resource</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicNetworkAccess: Enable or disable public network access. Default to &ldquo;Enabled&rdquo;.
When it&rsquo;s Enabled, network ACLs still apply.
When it&rsquo;s Disabled, public network access is always disabled no matter what you set in network ACLs.</p>
</td>
</tr>
<tr>
<td>
<code>resourceLogConfiguration</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfigurationARM">
ResourceLogConfigurationARM
</a>
</em>
</td>
<td>
<p>ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRTlsSettingsARM">
SignalRTlsSettingsARM
</a>
</em>
</td>
<td>
<p>Tls: TLS settings for the resource</p>
</td>
</tr>
<tr>
<td>
<code>upstream</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettingsARM">
ServerlessUpstreamSettingsARM
</a>
</em>
</td>
<td>
<p>Upstream: The settings for the Upstream when the service is in server-less mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_StatusARM">SignalRResource_StatusARM</a>)
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
<code>cors</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRCorsSettings_StatusARM">
SignalRCorsSettings_StatusARM
</a>
</em>
</td>
<td>
<p>Cors: Cross-Origin Resource Sharing (CORS) settings.</p>
</td>
</tr>
<tr>
<td>
<code>disableAadAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAadAuth: DisableLocalAuth
Enable or disable aad auth
When set as true, connection with AuthType=aad won&rsquo;t work.</p>
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
<p>DisableLocalAuth: DisableLocalAuth
Enable or disable local auth with AccessKey
When set as true, connection with AccessKey=xxx won&rsquo;t work.</p>
</td>
</tr>
<tr>
<td>
<code>externalIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExternalIP: The publicly accessible IP of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature_StatusARM">
[]SignalRFeature_StatusARM
</a>
</em>
</td>
<td>
<p>Features: List of the featureFlags.
FeatureFlags that are not included in the parameters for the update operation will not be modified.
And the response will only include featureFlags that are explicitly set.
When a featureFlag is not explicitly set, its globally default value will be used
But keep in mind, the default value doesn&rsquo;t mean &ldquo;false&rdquo;. It varies in terms of different FeatureFlags.</p>
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
<p>HostName: FQDN of the service instance.</p>
</td>
</tr>
<tr>
<td>
<code>hostNamePrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostNamePrefix: Deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>networkACLs</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_StatusARM">
SignalRNetworkACLs_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkACLs: Network ACLs</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: Private endpoint connections to the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicNetworkAccess: Enable or disable public network access. Default to &ldquo;Enabled&rdquo;.
When it&rsquo;s Enabled, network ACLs still apply.
When it&rsquo;s Disabled, public network access is always disabled no matter what you set in network ACLs.</p>
</td>
</tr>
<tr>
<td>
<code>publicPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>PublicPort: The publicly accessible port of the resource which is designed for browser/client side usage.</p>
</td>
</tr>
<tr>
<td>
<code>resourceLogConfiguration</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration_StatusARM">
ResourceLogConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.
If resourceLogConfiguration isn&rsquo;t null or empty, it will override options &ldquo;EnableConnectivityLog&rdquo; and
&ldquo;EnableMessagingLogs&rdquo; in features.
Otherwise, use options &ldquo;EnableConnectivityLog&rdquo; and &ldquo;EnableMessagingLogs&rdquo; in features.</p>
</td>
</tr>
<tr>
<td>
<code>serverPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>ServerPort: The publicly accessible port of the resource which is designed for customer server side usage.</p>
</td>
</tr>
<tr>
<td>
<code>sharedPrivateLinkResources</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM">
[]SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>SharedPrivateLinkResources: The list of shared private link resources.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRTlsSettings_StatusARM">
SignalRTlsSettings_StatusARM
</a>
</em>
</td>
<td>
<p>Tls: TLS settings.</p>
</td>
</tr>
<tr>
<td>
<code>upstream</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings_StatusARM">
ServerlessUpstreamSettings_StatusARM
</a>
</em>
</td>
<td>
<p>Upstream: Upstream settings when the service is in server-less mode.</p>
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
<p>Version: Version of the resource. Probably you need the same or higher version of client SDKs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRRequestType_Status">SignalRRequestType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.NetworkACL_Status">NetworkACL_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.NetworkACL_StatusARM">NetworkACL_StatusARM</a>, <a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL_Status">PrivateEndpointACL_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointACL_StatusARM">PrivateEndpointACL_StatusARM</a>)
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
<tbody><tr><td><p>&#34;ClientConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RESTAPI&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServerConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Trace&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR">SignalR</a>)
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
<code>cors</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRCorsSettings_Status">
SignalRCorsSettings_Status
</a>
</em>
</td>
<td>
<p>Cors: Cross-Origin Resource Sharing (CORS) settings.</p>
</td>
</tr>
<tr>
<td>
<code>disableAadAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAadAuth: DisableLocalAuth
Enable or disable aad auth
When set as true, connection with AuthType=aad won&rsquo;t work.</p>
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
<p>DisableLocalAuth: DisableLocalAuth
Enable or disable local auth with AccessKey
When set as true, connection with AccessKey=xxx won&rsquo;t work.</p>
</td>
</tr>
<tr>
<td>
<code>externalIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExternalIP: The publicly accessible IP of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature_Status">
[]SignalRFeature_Status
</a>
</em>
</td>
<td>
<p>Features: List of the featureFlags.
FeatureFlags that are not included in the parameters for the update operation will not be modified.
And the response will only include featureFlags that are explicitly set.
When a featureFlag is not explicitly set, its globally default value will be used
But keep in mind, the default value doesn&rsquo;t mean &ldquo;false&rdquo;. It varies in terms of different FeatureFlags.</p>
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
<p>HostName: FQDN of the service instance.</p>
</td>
</tr>
<tr>
<td>
<code>hostNamePrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostNamePrefix: Deprecated.</p>
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
<p>Id: Fully qualified resource Id for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity_Status">
ManagedIdentity_Status
</a>
</em>
</td>
<td>
<p>Identity: The managed identity response</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServiceKind_Status">
ServiceKind_Status
</a>
</em>
</td>
<td>
<p>Kind: The kind of the service - e.g. &ldquo;SignalR&rdquo; for &ldquo;Microsoft.SignalRService/SignalR&rdquo;</p>
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
<p>Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.</p>
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
<p>Name: The name of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>networkACLs</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs_Status">
SignalRNetworkACLs_Status
</a>
</em>
</td>
<td>
<p>NetworkACLs: Network ACLs</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: Private endpoint connections to the resource.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ProvisioningState_Status">
ProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicNetworkAccess: Enable or disable public network access. Default to &ldquo;Enabled&rdquo;.
When it&rsquo;s Enabled, network ACLs still apply.
When it&rsquo;s Disabled, public network access is always disabled no matter what you set in network ACLs.</p>
</td>
</tr>
<tr>
<td>
<code>publicPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>PublicPort: The publicly accessible port of the resource which is designed for browser/client side usage.</p>
</td>
</tr>
<tr>
<td>
<code>resourceLogConfiguration</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration_Status">
ResourceLogConfiguration_Status
</a>
</em>
</td>
<td>
<p>ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.
If resourceLogConfiguration isn&rsquo;t null or empty, it will override options &ldquo;EnableConnectivityLog&rdquo; and
&ldquo;EnableMessagingLogs&rdquo; in features.
Otherwise, use options &ldquo;EnableConnectivityLog&rdquo; and &ldquo;EnableMessagingLogs&rdquo; in features.</p>
</td>
</tr>
<tr>
<td>
<code>serverPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>ServerPort: The publicly accessible port of the resource which is designed for customer server side usage.</p>
</td>
</tr>
<tr>
<td>
<code>sharedPrivateLinkResources</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded">
[]SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>SharedPrivateLinkResources: The list of shared private link resources.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSku_Status">
ResourceSku_Status
</a>
</em>
</td>
<td>
<p>Sku: The billing information of the resource.(e.g. Free, Standard)</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
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
<p>Tags: Tags of the service which is a list of key value pairs that describe the resource.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRTlsSettings_Status">
SignalRTlsSettings_Status
</a>
</em>
</td>
<td>
<p>Tls: TLS settings.</p>
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
<p>Type: The type of the resource - e.g. &ldquo;Microsoft.SignalRService/SignalR&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>upstream</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings_Status">
ServerlessUpstreamSettings_Status
</a>
</em>
</td>
<td>
<p>Upstream: Upstream settings when the service is in server-less mode.</p>
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
<p>Version: Version of the resource. Probably you need the same or higher version of client SDKs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRResource_StatusARM">SignalRResource_StatusARM
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
<p>Id: Fully qualified resource Id for the resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity_StatusARM">
ManagedIdentity_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: The managed identity response</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServiceKind_Status">
ServiceKind_Status
</a>
</em>
</td>
<td>
<p>Kind: The kind of the service - e.g. &ldquo;SignalR&rdquo; for &ldquo;Microsoft.SignalRService/SignalR&rdquo;</p>
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
<p>Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.</p>
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
<p>Name: The name of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">
SignalRProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Settings used to provision or configure the resource</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSku_StatusARM">
ResourceSku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The billing information of the resource.(e.g. Free, Standard)</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
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
<p>Tags: Tags of the service which is a list of key value pairs that describe the resource.</p>
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
<p>Type: The type of the resource - e.g. &ldquo;Microsoft.SignalRService/SignalR&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRSkuTier_Status">SignalRSkuTier_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ResourceSku_Status">ResourceSku_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.ResourceSku_StatusARM">ResourceSku_StatusARM</a>)
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
</tr><tr><td><p>&#34;Free&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRSpecAPIVersion">SignalRSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-10-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRSpecKind">SignalRSpecKind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalR_SpecARM">SignalR_SpecARM</a>)
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
<tbody><tr><td><p>&#34;RawWebSockets&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SignalR&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRTlsSettings">SignalRTlsSettings
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRTlsSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRTlsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientCertEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ClientCertEnabled: Request client certificate during TLS handshake if enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRTlsSettingsARM">SignalRTlsSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">SignalRPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRTlsSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRTlsSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientCertEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ClientCertEnabled: Request client certificate during TLS handshake if enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRTlsSettings_Status">SignalRTlsSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<code>clientCertEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ClientCertEnabled: Request client certificate during TLS handshake if enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalRTlsSettings_StatusARM">SignalRTlsSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalRProperties_StatusARM">SignalRProperties_StatusARM</a>)
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
<code>clientCertEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ClientCertEnabled: Request client certificate during TLS handshake if enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalR_Spec">SignalR_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SignalR">SignalR</a>)
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
<code>cors</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRCorsSettings">
SignalRCorsSettings
</a>
</em>
</td>
<td>
<p>Cors: Cross-Origin Resource Sharing (CORS) settings.</p>
</td>
</tr>
<tr>
<td>
<code>disableAadAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableAadAuth: DisableLocalAuth
Enable or disable aad auth
When set as true, connection with AuthType=aad won&rsquo;t work.</p>
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
<p>DisableLocalAuth: DisableLocalAuth
Enable or disable local auth with AccessKey
When set as true, connection with AccessKey=xxx won&rsquo;t work.</p>
</td>
</tr>
<tr>
<td>
<code>features</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRFeature">
[]SignalRFeature
</a>
</em>
</td>
<td>
<p>Features: List of the featureFlags.
FeatureFlags that are not included in the parameters for the update operation will not be modified.
And the response will only include featureFlags that are explicitly set.
When a featureFlag is not explicitly set, its globally default value will be used
But keep in mind, the default value doesn&rsquo;t mean &ldquo;false&rdquo;. It varies in terms of different FeatureFlags.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity">
ManagedIdentity
</a>
</em>
</td>
<td>
<p>Identity: A class represent managed identities used for request and response</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRSpecKind">
SignalRSpecKind
</a>
</em>
</td>
<td>
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
<p>Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.</p>
</td>
</tr>
<tr>
<td>
<code>networkACLs</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRNetworkACLs">
SignalRNetworkACLs
</a>
</em>
</td>
<td>
<p>NetworkACLs: Network ACLs for the resource</p>
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
string
</em>
</td>
<td>
<p>PublicNetworkAccess: Enable or disable public network access. Default to &ldquo;Enabled&rdquo;.
When it&rsquo;s Enabled, network ACLs still apply.
When it&rsquo;s Disabled, public network access is always disabled no matter what you set in network ACLs.</p>
</td>
</tr>
<tr>
<td>
<code>resourceLogConfiguration</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceLogConfiguration">
ResourceLogConfiguration
</a>
</em>
</td>
<td>
<p>ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSku">
ResourceSku
</a>
</em>
</td>
<td>
<p>Sku: The billing information of the resource.</p>
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
<p>Tags: Tags of the service which is a list of key value pairs that describe the resource.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRTlsSettings">
SignalRTlsSettings
</a>
</em>
</td>
<td>
<p>Tls: TLS settings for the resource</p>
</td>
</tr>
<tr>
<td>
<code>upstream</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings">
ServerlessUpstreamSettings
</a>
</em>
</td>
<td>
<p>Upstream: The settings for the Upstream when the service is in server-less mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SignalR_SpecARM">SignalR_SpecARM
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
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentityARM">
ManagedIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: A class represent managed identities used for request and response</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRSpecKind">
SignalRSpecKind
</a>
</em>
</td>
<td>
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
<p>Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.</p>
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
<p>Name: The name of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.SignalRPropertiesARM">
SignalRPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: A class that describes the properties of the resource</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ResourceSkuARM">
ResourceSkuARM
</a>
</em>
</td>
<td>
<p>Sku: The billing information of the resource.</p>
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
<p>Tags: Tags of the service which is a list of key value pairs that describe the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.SystemDataStatusCreatedByType">SystemDataStatusCreatedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SystemData_Status">SystemData_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="signalrservice.azure.com/v1beta20211001.SystemDataStatusLastModifiedByType">SystemDataStatusLastModifiedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.SystemData_Status">SystemData_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="signalrservice.azure.com/v1beta20211001.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded">PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded</a>, <a href="#signalrservice.azure.com/v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded">SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_Status">SignalRResource_Status</a>)
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
<a href="#signalrservice.azure.com/v1beta20211001.SystemDataStatusCreatedByType">
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
<p>LastModifiedAt: The timestamp of resource last modification (UTC)</p>
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
<a href="#signalrservice.azure.com/v1beta20211001.SystemDataStatusLastModifiedByType">
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
<h3 id="signalrservice.azure.com/v1beta20211001.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM</a>, <a href="#signalrservice.azure.com/v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM">SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM</a>, <a href="#signalrservice.azure.com/v1beta20211001.SignalRResource_StatusARM">SignalRResource_StatusARM</a>)
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
<a href="#signalrservice.azure.com/v1beta20211001.SystemDataStatusCreatedByType">
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
<p>LastModifiedAt: The timestamp of resource last modification (UTC)</p>
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
<a href="#signalrservice.azure.com/v1beta20211001.SystemDataStatusLastModifiedByType">
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
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings">UpstreamAuthSettings
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplate">UpstreamTemplate</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamAuthSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamAuthSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>managedIdentity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentitySettings">
ManagedIdentitySettings
</a>
</em>
</td>
<td>
<p>ManagedIdentity: Managed identity settings for upstream.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsType">
UpstreamAuthSettingsType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsARM">UpstreamAuthSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplateARM">UpstreamTemplateARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamAuthSettings">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamAuthSettings</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>managedIdentity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentitySettingsARM">
ManagedIdentitySettingsARM
</a>
</em>
</td>
<td>
<p>ManagedIdentity: Managed identity settings for upstream.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsType">
UpstreamAuthSettingsType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsType">UpstreamAuthSettingsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings">UpstreamAuthSettings</a>, <a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsARM">UpstreamAuthSettingsARM</a>)
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
<tbody><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_Status">UpstreamAuthSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplate_Status">UpstreamTemplate_Status</a>)
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
<code>managedIdentity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentitySettings_Status">
ManagedIdentitySettings_Status
</a>
</em>
</td>
<td>
<p>ManagedIdentity: Gets or sets the managed identity settings. It&rsquo;s required if the auth type is set to ManagedIdentity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthType_Status">
UpstreamAuthType_Status
</a>
</em>
</td>
<td>
<p>Type: Gets or sets the type of auth. None or ManagedIdentity is supported now.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_StatusARM">UpstreamAuthSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamTemplate_StatusARM">UpstreamTemplate_StatusARM</a>)
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
<code>managedIdentity</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentitySettings_StatusARM">
ManagedIdentitySettings_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedIdentity: Gets or sets the managed identity settings. It&rsquo;s required if the auth type is set to ManagedIdentity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthType_Status">
UpstreamAuthType_Status
</a>
</em>
</td>
<td>
<p>Type: Gets or sets the type of auth. None or ManagedIdentity is supported now.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamAuthType_Status">UpstreamAuthType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_Status">UpstreamAuthSettings_Status</a>, <a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_StatusARM">UpstreamAuthSettings_StatusARM</a>)
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
<tbody><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamTemplate">UpstreamTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings">ServerlessUpstreamSettings</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamTemplate">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamTemplate</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings">
UpstreamAuthSettings
</a>
</em>
</td>
<td>
<p>Auth: Upstream auth settings. If not set, no auth is used for upstream messages.</p>
</td>
</tr>
<tr>
<td>
<code>categoryPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any category name
2. Combine multiple categories with &ldquo;,&rdquo;, for example &ldquo;connections,messages&rdquo;, it matches category &ldquo;connections&rdquo; and
&ldquo;messages&rdquo;
3. The single category name, for example, &ldquo;connections&rdquo;, it matches the category &ldquo;connections&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>eventPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any event name
2. Combine multiple events with &ldquo;,&rdquo;, for example &ldquo;connect,disconnect&rdquo;, it matches event &ldquo;connect&rdquo; and &ldquo;disconnect&rdquo;
3. The single event name, for example, &ldquo;connect&rdquo;, it matches &ldquo;connect&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>hubPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any hub name
2. Combine multiple hubs with &ldquo;,&rdquo;, for example &ldquo;hub1,hub2&rdquo;, it matches &ldquo;hub1&rdquo; and &ldquo;hub2&rdquo;
3. The single hub name, for example, &ldquo;hub1&rdquo;, it matches &ldquo;hub1&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>urlTemplate</code><br/>
<em>
string
</em>
</td>
<td>
<p>UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
For example, if the urlTemplate is <code>http://example.com/{hub}/api/{event}</code>, with a client request from hub <code>chat</code>
connects, it will first POST to this URL: <code>http://example.com/chat/api/connect</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamTemplateARM">UpstreamTemplateARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettingsARM">ServerlessUpstreamSettingsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamTemplate">https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamTemplate</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettingsARM">
UpstreamAuthSettingsARM
</a>
</em>
</td>
<td>
<p>Auth: Upstream auth settings. If not set, no auth is used for upstream messages.</p>
</td>
</tr>
<tr>
<td>
<code>categoryPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any category name
2. Combine multiple categories with &ldquo;,&rdquo;, for example &ldquo;connections,messages&rdquo;, it matches category &ldquo;connections&rdquo; and
&ldquo;messages&rdquo;
3. The single category name, for example, &ldquo;connections&rdquo;, it matches the category &ldquo;connections&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>eventPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any event name
2. Combine multiple events with &ldquo;,&rdquo;, for example &ldquo;connect,disconnect&rdquo;, it matches event &ldquo;connect&rdquo; and &ldquo;disconnect&rdquo;
3. The single event name, for example, &ldquo;connect&rdquo;, it matches &ldquo;connect&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>hubPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any hub name
2. Combine multiple hubs with &ldquo;,&rdquo;, for example &ldquo;hub1,hub2&rdquo;, it matches &ldquo;hub1&rdquo; and &ldquo;hub2&rdquo;
3. The single hub name, for example, &ldquo;hub1&rdquo;, it matches &ldquo;hub1&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>urlTemplate</code><br/>
<em>
string
</em>
</td>
<td>
<p>UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
For example, if the urlTemplate is <code>http://example.com/{hub}/api/{event}</code>, with a client request from hub <code>chat</code>
connects, it will first POST to this URL: <code>http://example.com/chat/api/connect</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamTemplate_Status">UpstreamTemplate_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings_Status">ServerlessUpstreamSettings_Status</a>)
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
<code>auth</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_Status">
UpstreamAuthSettings_Status
</a>
</em>
</td>
<td>
<p>Auth: Gets or sets the auth settings for an upstream. If not set, no auth is used for upstream messages.</p>
</td>
</tr>
<tr>
<td>
<code>categoryPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any category name
2. Combine multiple categories with &ldquo;,&rdquo;, for example &ldquo;connections,messages&rdquo;, it matches category &ldquo;connections&rdquo; and
&ldquo;messages&rdquo;
3. The single category name, for example, &ldquo;connections&rdquo;, it matches the category &ldquo;connections&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>eventPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any event name
2. Combine multiple events with &ldquo;,&rdquo;, for example &ldquo;connect,disconnect&rdquo;, it matches event &ldquo;connect&rdquo; and &ldquo;disconnect&rdquo;
3. The single event name, for example, &ldquo;connect&rdquo;, it matches &ldquo;connect&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>hubPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any hub name
2. Combine multiple hubs with &ldquo;,&rdquo;, for example &ldquo;hub1,hub2&rdquo;, it matches &ldquo;hub1&rdquo; and &ldquo;hub2&rdquo;
3. The single hub name, for example, &ldquo;hub1&rdquo;, it matches &ldquo;hub1&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>urlTemplate</code><br/>
<em>
string
</em>
</td>
<td>
<p>UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
For example, if the urlTemplate is <code>http://example.com/{hub}/api/{event}</code>, with a client request from hub <code>chat</code>
connects, it will first POST to this URL: <code>http://example.com/chat/api/connect</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UpstreamTemplate_StatusARM">UpstreamTemplate_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ServerlessUpstreamSettings_StatusARM">ServerlessUpstreamSettings_StatusARM</a>)
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
<code>auth</code><br/>
<em>
<a href="#signalrservice.azure.com/v1beta20211001.UpstreamAuthSettings_StatusARM">
UpstreamAuthSettings_StatusARM
</a>
</em>
</td>
<td>
<p>Auth: Gets or sets the auth settings for an upstream. If not set, no auth is used for upstream messages.</p>
</td>
</tr>
<tr>
<td>
<code>categoryPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any category name
2. Combine multiple categories with &ldquo;,&rdquo;, for example &ldquo;connections,messages&rdquo;, it matches category &ldquo;connections&rdquo; and
&ldquo;messages&rdquo;
3. The single category name, for example, &ldquo;connections&rdquo;, it matches the category &ldquo;connections&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>eventPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any event name
2. Combine multiple events with &ldquo;,&rdquo;, for example &ldquo;connect,disconnect&rdquo;, it matches event &ldquo;connect&rdquo; and &ldquo;disconnect&rdquo;
3. The single event name, for example, &ldquo;connect&rdquo;, it matches &ldquo;connect&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>hubPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
There are 3 kind of patterns supported:
1. &ldquo;*&rdquo;, it to matches any hub name
2. Combine multiple hubs with &ldquo;,&rdquo;, for example &ldquo;hub1,hub2&rdquo;, it matches &ldquo;hub1&rdquo; and &ldquo;hub2&rdquo;
3. The single hub name, for example, &ldquo;hub1&rdquo;, it matches &ldquo;hub1&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>urlTemplate</code><br/>
<em>
string
</em>
</td>
<td>
<p>UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
For example, if the urlTemplate is <code>http://example.com/{hub}/api/{event}</code>, with a client request from hub <code>chat</code>
connects, it will first POST to this URL: <code>http://example.com/chat/api/connect</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UserAssignedIdentityProperty_Status">UserAssignedIdentityProperty_Status
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity_Status">ManagedIdentity_Status</a>)
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
<p>ClientId: Get the client id for the user assigned identity</p>
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
<p>PrincipalId: Get the principal id for the user assigned identity</p>
</td>
</tr>
</tbody>
</table>
<h3 id="signalrservice.azure.com/v1beta20211001.UserAssignedIdentityProperty_StatusARM">UserAssignedIdentityProperty_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#signalrservice.azure.com/v1beta20211001.ManagedIdentity_StatusARM">ManagedIdentity_StatusARM</a>)
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
<p>ClientId: Get the client id for the user assigned identity</p>
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
<p>PrincipalId: Get the principal id for the user assigned identity</p>
</td>
</tr>
</tbody>
</table>
<hr/>
