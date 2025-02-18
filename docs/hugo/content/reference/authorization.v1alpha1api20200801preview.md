---
---
<h2 id="authorization.azure.com/v1alpha1api20200801preview">authorization.azure.com/v1alpha1api20200801preview</h2>
<div>
<p>Package v1alpha1api20200801preview contains API Schema definitions for the authorization v1alpha1api20200801preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignment">RoleAssignment
</h3>
<div>
<p>Deprecated version of RoleAssignment. Use v1beta20200801preview.RoleAssignment instead</p>
</div>
<table>
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
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_Spec">
RoleAssignments_Spec
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
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
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ArbitraryOwnerReference">
genruntime.ArbitraryOwnerReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
extension resource, which means that any other Azure resource can be its owner.</p>
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
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">
RoleAssignmentPropertiesPrincipalType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
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
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_Status">
RoleAssignment_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesARM">RoleAssignmentPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_SpecARM">RoleAssignments_SpecARM</a>)
</p>
<div>
<p>Deprecated version of RoleAssignmentProperties. Use v1beta20200801preview.RoleAssignmentProperties instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
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
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">
RoleAssignmentPropertiesPrincipalType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">RoleAssignmentPropertiesPrincipalType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesARM">RoleAssignmentPropertiesARM</a>, <a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_Spec">RoleAssignments_Spec</a>)
</p>
<div>
<p>Deprecated version of RoleAssignmentPropertiesPrincipalType. Use
v1beta20200801preview.RoleAssignmentPropertiesPrincipalType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ForeignGroup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Group&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServicePrincipal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesStatusPrincipalType">RoleAssignmentPropertiesStatusPrincipalType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentProperties_StatusARM">RoleAssignmentProperties_StatusARM</a>, <a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_Status">RoleAssignment_Status</a>)
</p>
<div>
<p>Deprecated version of RoleAssignmentPropertiesStatusPrincipalType. Use
v1beta20200801preview.RoleAssignmentPropertiesStatusPrincipalType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ForeignGroup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Group&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServicePrincipal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentProperties_StatusARM">RoleAssignmentProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_StatusARM">RoleAssignment_StatusARM</a>)
</p>
<div>
<p>Deprecated version of RoleAssignmentProperties_Status. Use v1beta20200801preview.RoleAssignmentProperties_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
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
</td>
</tr>
<tr>
<td>
<code>createdOn</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
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
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesStatusPrincipalType">
RoleAssignmentPropertiesStatusPrincipalType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>updatedBy</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>updatedOn</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_Status">RoleAssignment_Status
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment">RoleAssignment</a>)
</p>
<div>
<p>Deprecated version of RoleAssignment_Status. Use v1beta20200801preview.RoleAssignment_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
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
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>createdOn</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
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
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesStatusPrincipalType">
RoleAssignmentPropertiesStatusPrincipalType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
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
</td>
</tr>
<tr>
<td>
<code>updatedBy</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>updatedOn</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_StatusARM">RoleAssignment_StatusARM
</h3>
<div>
<p>Deprecated version of RoleAssignment_Status. Use v1beta20200801preview.RoleAssignment_Status instead</p>
</div>
<table>
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentProperties_StatusARM">
RoleAssignmentProperties_StatusARM
</a>
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
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_Spec">RoleAssignments_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment">RoleAssignment</a>)
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
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
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ArbitraryOwnerReference">
genruntime.ArbitraryOwnerReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
extension resource, which means that any other Azure resource can be its owner.</p>
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
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">
RoleAssignmentPropertiesPrincipalType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
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
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_SpecARM">RoleAssignments_SpecARM
</h3>
<div>
<p>Deprecated version of RoleAssignments_Spec. Use v1beta20200801preview.RoleAssignments_Spec instead</p>
</div>
<table>
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
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesARM">
RoleAssignmentPropertiesARM
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
</tbody>
</table>
<hr/>
