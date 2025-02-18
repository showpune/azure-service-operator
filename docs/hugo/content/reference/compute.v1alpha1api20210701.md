---
---
<h2 id="compute.azure.com/v1alpha1api20210701">compute.azure.com/v1alpha1api20210701</h2>
<div>
<p>Package v1alpha1api20210701 contains API Schema definitions for the compute v1alpha1api20210701 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="compute.azure.com/v1alpha1api20210701.DiskEncryptionSetParameters">DiskEncryptionSetParameters
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk</a>)
</p>
<div>
<p>Deprecated version of DiskEncryptionSetParameters. Use v1beta20210701.DiskEncryptionSetParameters instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
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
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.DiskEncryptionSetParametersARM">DiskEncryptionSetParametersARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskARM">ImageDataDiskARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
</p>
<div>
<p>Deprecated version of DiskEncryptionSetParameters. Use v1beta20210701.DiskEncryptionSetParameters instead</p>
</div>
<table>
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
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Deprecated version of ExtendedLocation. Use v1beta20210701.ExtendedLocation instead</p>
</div>
<table>
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
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ExtendedLocationARM">ExtendedLocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Images_SpecARM">Images_SpecARM</a>)
</p>
<div>
<p>Deprecated version of ExtendedLocation. Use v1beta20210701.ExtendedLocation instead</p>
</div>
<table>
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
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation">ExtendedLocation</a>, <a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocationARM">ExtendedLocationARM</a>)
</p>
<div>
<p>Deprecated version of ExtendedLocationType. Use v1beta20210701.ExtendedLocationType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ExtendedLocationType_Status">ExtendedLocationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation_Status">ExtendedLocation_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ExtendedLocationType_Status. Use v1beta20210701.ExtendedLocationType_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ExtendedLocation_Status">ExtendedLocation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Image_Status">Image_Status</a>)
</p>
<div>
<p>Deprecated version of ExtendedLocation_Status. Use v1beta20210701.ExtendedLocation_Status instead</p>
</div>
<table>
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
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Image_StatusARM">Image_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ExtendedLocation_Status. Use v1beta20210701.ExtendedLocation_Status instead</p>
</div>
<table>
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
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.HyperVGenerationType_Status">HyperVGenerationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageProperties_StatusARM">ImageProperties_StatusARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.Image_Status">Image_Status</a>)
</p>
<div>
<p>Deprecated version of HyperVGenerationType_Status. Use v1beta20210701.HyperVGenerationType_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;V1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;V2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.Image">Image
</h3>
<div>
<p>Deprecated version of Image. Use v1beta20210701.Image instead</p>
</div>
<table>
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
<a href="#compute.azure.com/v1alpha1api20210701.Images_Spec">
Images_Spec
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
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesHyperVGeneration">
ImagePropertiesHyperVGeneration
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
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile">
ImageStorageProfile
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
<a href="#compute.azure.com/v1alpha1api20210701.Image_Status">
Image_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDisk">ImageDataDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile">ImageStorageProfile</a>)
</p>
<div>
<p>Deprecated version of ImageDataDisk. Use v1beta20210701.ImageDataDisk instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskCaching">
ImageDataDiskCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.DiskEncryptionSetParameters">
DiskEncryptionSetParameters
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskStorageAccountType">
ImageDataDiskStorageAccountType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDiskARM">ImageDataDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfileARM">ImageStorageProfileARM</a>)
</p>
<div>
<p>Deprecated version of ImageDataDisk. Use v1beta20210701.ImageDataDisk instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskCaching">
ImageDataDiskCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.DiskEncryptionSetParametersARM">
DiskEncryptionSetParametersARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskStorageAccountType">
ImageDataDiskStorageAccountType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDiskCaching">ImageDataDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskARM">ImageDataDiskARM</a>)
</p>
<div>
<p>Deprecated version of ImageDataDiskCaching. Use v1beta20210701.ImageDataDiskCaching instead</p>
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDiskStatusCaching">ImageDataDiskStatusCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_Status">ImageDataDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageDataDiskStatusCaching. Use v1beta20210701.ImageDataDiskStatusCaching instead</p>
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDiskStorageAccountType">ImageDataDiskStorageAccountType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskARM">ImageDataDiskARM</a>)
</p>
<div>
<p>Deprecated version of ImageDataDiskStorageAccountType. Use v1beta20210701.ImageDataDiskStorageAccountType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDisk_Status">ImageDataDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile_Status">ImageStorageProfile_Status</a>)
</p>
<div>
<p>Deprecated version of ImageDataDisk_Status. Use v1beta20210701.ImageDataDisk_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskStatusCaching">
ImageDataDiskStatusCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile_StatusARM">ImageStorageProfile_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageDataDisk_Status. Use v1beta20210701.ImageDataDisk_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskStatusCaching">
ImageDataDiskStatusCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile">ImageStorageProfile</a>)
</p>
<div>
<p>Deprecated version of ImageOSDisk. Use v1beta20210701.ImageOSDisk instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskCaching">
ImageOSDiskCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.DiskEncryptionSetParameters">
DiskEncryptionSetParameters
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osState</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskOsState">
ImageOSDiskOsState
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskOsType">
ImageOSDiskOsType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStorageAccountType">
ImageOSDiskStorageAccountType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfileARM">ImageStorageProfileARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDisk. Use v1beta20210701.ImageOSDisk instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskCaching">
ImageOSDiskCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.DiskEncryptionSetParametersARM">
DiskEncryptionSetParametersARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osState</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskOsState">
ImageOSDiskOsState
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskOsType">
ImageOSDiskOsType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStorageAccountType">
ImageOSDiskStorageAccountType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskCaching">ImageOSDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskCaching. Use v1beta20210701.ImageOSDiskCaching instead</p>
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskOsState">ImageOSDiskOsState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskOsState. Use v1beta20210701.ImageOSDiskOsState instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Generalized&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Specialized&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskOsType">ImageOSDiskOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskOsType. Use v1beta20210701.ImageOSDiskOsType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusCaching">ImageOSDiskStatusCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskStatusCaching. Use v1beta20210701.ImageOSDiskStatusCaching instead</p>
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusOsState">ImageOSDiskStatusOsState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskStatusOsState. Use v1beta20210701.ImageOSDiskStatusOsState instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Generalized&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Specialized&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusOsType">ImageOSDiskStatusOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskStatusOsType. Use v1beta20210701.ImageOSDiskStatusOsType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDiskStorageAccountType">ImageOSDiskStorageAccountType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDiskStorageAccountType. Use v1beta20210701.ImageOSDiskStorageAccountType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">ImageOSDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile_Status">ImageStorageProfile_Status</a>)
</p>
<div>
<p>Deprecated version of ImageOSDisk_Status. Use v1beta20210701.ImageOSDisk_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusCaching">
ImageOSDiskStatusCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osState</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusOsState">
ImageOSDiskStatusOsState
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusOsType">
ImageOSDiskStatusOsType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile_StatusARM">ImageStorageProfile_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageOSDisk_Status. Use v1beta20210701.ImageOSDisk_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusCaching">
ImageOSDiskStatusCaching
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osState</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusOsState">
ImageOSDiskStatusOsState
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskStatusOsType">
ImageOSDiskStatusOsType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImagePropertiesARM">ImagePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Images_SpecARM">Images_SpecARM</a>)
</p>
<div>
<p>Deprecated version of ImageProperties. Use v1beta20210701.ImageProperties instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesHyperVGeneration">
ImagePropertiesHyperVGeneration
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResourceARM">
SubResourceARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfileARM">
ImageStorageProfileARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImagePropertiesHyperVGeneration">ImagePropertiesHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesARM">ImagePropertiesARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Deprecated version of ImagePropertiesHyperVGeneration. Use v1beta20210701.ImagePropertiesHyperVGeneration instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;V1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;V2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageProperties_StatusARM">ImageProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Image_StatusARM">Image_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageProperties_Status. Use v1beta20210701.ImageProperties_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.HyperVGenerationType_Status">
HyperVGenerationType_Status
</a>
</em>
</td>
<td>
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
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile_StatusARM">
ImageStorageProfile_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageStorageProfile">ImageStorageProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Deprecated version of ImageStorageProfile. Use v1beta20210701.ImageStorageProfile instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk">
[]ImageDataDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">
ImageOSDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>zoneResilient</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageStorageProfileARM">ImageStorageProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesARM">ImagePropertiesARM</a>)
</p>
<div>
<p>Deprecated version of ImageStorageProfile. Use v1beta20210701.ImageStorageProfile instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskARM">
[]ImageDataDiskARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">
ImageOSDiskARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>zoneResilient</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageStorageProfile_Status">ImageStorageProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Image_Status">Image_Status</a>)
</p>
<div>
<p>Deprecated version of ImageStorageProfile_Status. Use v1beta20210701.ImageStorageProfile_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_Status">
[]ImageDataDisk_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">
ImageOSDisk_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>zoneResilient</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.ImageStorageProfile_StatusARM">ImageStorageProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageProperties_StatusARM">ImageProperties_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ImageStorageProfile_Status. Use v1beta20210701.ImageStorageProfile_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_StatusARM">
[]ImageDataDisk_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">
ImageOSDisk_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>zoneResilient</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.Image_Status">Image_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Image">Image</a>)
</p>
<div>
<p>Deprecated version of Image_Status. Use v1beta20210701.Image_Status instead</p>
</div>
<table>
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
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.HyperVGenerationType_Status">
HyperVGenerationType_Status
</a>
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
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile_Status">
ImageStorageProfile_Status
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
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.Image_StatusARM">Image_StatusARM
</h3>
<div>
<p>Deprecated version of Image_Status. Use v1beta20210701.Image_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
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
<a href="#compute.azure.com/v1alpha1api20210701.ImageProperties_StatusARM">
ImageProperties_StatusARM
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
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.Images_Spec">Images_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.Image">Image</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesHyperVGeneration">
ImagePropertiesHyperVGeneration
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
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.SubResource">
SubResource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ImageStorageProfile">
ImageStorageProfile
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
<h3 id="compute.azure.com/v1alpha1api20210701.Images_SpecARM">Images_SpecARM
</h3>
<div>
<p>Deprecated version of Images_Spec. Use v1beta20210701.Images_Spec instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1alpha1api20210701.ExtendedLocationARM">
ExtendedLocationARM
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
<a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesARM">
ImagePropertiesARM
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
<h3 id="compute.azure.com/v1alpha1api20210701.StorageAccountType_Status">StorageAccountType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_Status">ImageDataDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
</p>
<div>
<p>Deprecated version of StorageAccountType_Status. Use v1beta20210701.StorageAccountType_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_ZRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.SubResource">SubResource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1alpha1api20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Deprecated version of SubResource. Use v1beta20210701.SubResource instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
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
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.SubResourceARM">SubResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDiskARM">ImageDataDiskARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDiskARM">ImageOSDiskARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImagePropertiesARM">ImagePropertiesARM</a>)
</p>
<div>
<p>Deprecated version of SubResource. Use v1beta20210701.SubResource instead</p>
</div>
<table>
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
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.SubResource_Status">SubResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_Status">ImageDataDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1alpha1api20210701.Image_Status">Image_Status</a>)
</p>
<div>
<p>Deprecated version of SubResource_Status. Use v1beta20210701.SubResource_Status instead</p>
</div>
<table>
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
</tbody>
</table>
<h3 id="compute.azure.com/v1alpha1api20210701.SubResource_StatusARM">SubResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1alpha1api20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>, <a href="#compute.azure.com/v1alpha1api20210701.ImageProperties_StatusARM">ImageProperties_StatusARM</a>)
</p>
<div>
<p>Deprecated version of SubResource_Status. Use v1beta20210701.SubResource_Status instead</p>
</div>
<table>
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
</tbody>
</table>
<hr/>
