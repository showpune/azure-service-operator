---
---
<h2 id="compute.azure.com/v1beta20210701">compute.azure.com/v1beta20210701</h2>
<div>
<p>Package v1beta20210701 contains API Schema definitions for the compute v1beta20210701 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="compute.azure.com/v1beta20210701.DiskEncryptionSetParameters">DiskEncryptionSetParameters
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters</a></p>
</div>
<table>
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
<p>Reference: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.DiskEncryptionSetParametersARM">DiskEncryptionSetParametersARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDiskARM">ImageDataDiskARM</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters</a></p>
</div>
<table>
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
<h3 id="compute.azure.com/v1beta20210701.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ExtendedLocationARM">ExtendedLocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Images_SpecARM">Images_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ExtendedLocation">ExtendedLocation</a>, <a href="#compute.azure.com/v1beta20210701.ExtendedLocationARM">ExtendedLocationARM</a>)
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
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ExtendedLocationType_Status">ExtendedLocationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ExtendedLocation_Status">ExtendedLocation_Status</a>, <a href="#compute.azure.com/v1beta20210701.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM</a>)
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
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ExtendedLocation_Status">ExtendedLocation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Image_Status">Image_Status</a>)
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Image_StatusARM">Image_StatusARM</a>)
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.HyperVGenerationType_Status">HyperVGenerationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageProperties_StatusARM">ImageProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20210701.Image_Status">Image_Status</a>)
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
<tbody><tr><td><p>&#34;V1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;V2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.Image">Image
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/resourceDefinitions/images">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/resourceDefinitions/images</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20210701.Images_Spec">
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
<a href="#compute.azure.com/v1beta20210701.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImagePropertiesHyperVGeneration">
ImagePropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
resource.</p>
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
<a href="#compute.azure.com/v1beta20210701.SubResource">
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
<a href="#compute.azure.com/v1beta20210701.ImageStorageProfile">
ImageStorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Describes a storage profile.</p>
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
<a href="#compute.azure.com/v1beta20210701.Image_Status">
Image_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageDataDisk">ImageDataDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfile">ImageStorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageDataDisk">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageDataDisk</a></p>
</div>
<table>
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
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskCaching">
ImageDataDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.DiskEncryptionSetParameters">
DiskEncryptionSetParameters
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
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
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource">
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
<a href="#compute.azure.com/v1beta20210701.SubResource">
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
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskStorageAccountType">
ImageDataDiskStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageDataDiskARM">ImageDataDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfileARM">ImageStorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageDataDisk">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageDataDisk</a></p>
</div>
<table>
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
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskCaching">
ImageDataDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.DiskEncryptionSetParametersARM">
DiskEncryptionSetParametersARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
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
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResourceARM">
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
<a href="#compute.azure.com/v1beta20210701.SubResourceARM">
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
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskStorageAccountType">
ImageDataDiskStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageDataDiskCaching">ImageDataDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageDataDiskARM">ImageDataDiskARM</a>)
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageDataDiskStatusCaching">ImageDataDiskStatusCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk_Status">ImageDataDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM</a>)
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageDataDiskStorageAccountType">ImageDataDiskStorageAccountType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageDataDiskARM">ImageDataDiskARM</a>)
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
<h3 id="compute.azure.com/v1beta20210701.ImageDataDisk_Status">ImageDataDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfile_Status">ImageStorageProfile_Status</a>)
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
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskStatusCaching">
ImageDataDiskStatusCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
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
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managedDisk.</p>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>Snapshot: The snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfile_StatusARM">ImageStorageProfile_StatusARM</a>)
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
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskStatusCaching">
ImageDataDiskStatusCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
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
<p>Lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and
therefore must be unique for each data disk attached to a VM.</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managedDisk.</p>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>Snapshot: The snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfile">ImageStorageProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageOSDisk">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageOSDisk</a></p>
</div>
<table>
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
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskCaching">
ImageOSDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.DiskEncryptionSetParameters">
DiskEncryptionSetParameters
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource">
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
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskOsState">
ImageOSDiskOsState
</a>
</em>
</td>
<td>
<p>OsState: The OS State.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskOsType">
ImageOSDiskOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
custom image.
Possible values are:
Windows
Linux.</p>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource">
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
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStorageAccountType">
ImageOSDiskStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfileARM">ImageStorageProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageOSDisk">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageOSDisk</a></p>
</div>
<table>
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
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskCaching">
ImageOSDiskCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.DiskEncryptionSetParametersARM">
DiskEncryptionSetParametersARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Describes the parameter of customer managed disk encryption set resource id that can be specified for
disk.
NOTE: The disk encryption set resource id can only be specified for managed disk. Please refer
<a href="https://aka.ms/mdssewithcmkoverview">https://aka.ms/mdssewithcmkoverview</a> for more details.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResourceARM">
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
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskOsState">
ImageOSDiskOsState
</a>
</em>
</td>
<td>
<p>OsState: The OS State.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskOsType">
ImageOSDiskOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
custom image.
Possible values are:
Windows
Linux.</p>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResourceARM">
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
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStorageAccountType">
ImageOSDiskStorageAccountType
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskCaching">ImageOSDiskCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskOsState">ImageOSDiskOsState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
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
<tbody><tr><td><p>&#34;Generalized&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Specialized&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskOsType">ImageOSDiskOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
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
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskStatusCaching">ImageOSDiskStatusCaching
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
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
</tr><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskStatusOsState">ImageOSDiskStatusOsState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Generalized&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Specialized&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskStatusOsType">ImageOSDiskStatusOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDiskStorageAccountType">ImageOSDiskStorageAccountType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM</a>)
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
<h3 id="compute.azure.com/v1beta20210701.ImageOSDisk_Status">ImageOSDisk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfile_Status">ImageStorageProfile_Status</a>)
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
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStatusCaching">
ImageOSDiskStatusCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managedDisk.</p>
</td>
</tr>
<tr>
<td>
<code>osState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStatusOsState">
ImageOSDiskStatusOsState
</a>
</em>
</td>
<td>
<p>OsState: The OS State.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStatusOsType">
ImageOSDiskStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
custom image.
Possible values are:
Windows
Linux</p>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>Snapshot: The snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageStorageProfile_StatusARM">ImageStorageProfile_StatusARM</a>)
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
<code>blobUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobUri: The Virtual Hard Disk.</p>
</td>
</tr>
<tr>
<td>
<code>caching</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStatusCaching">
ImageOSDiskStatusCaching
</a>
</em>
</td>
<td>
<p>Caching: Specifies the caching requirements.
Possible values are:
None
ReadOnly
ReadWrite
Default: None for Standard storage. ReadOnly for Premium storage</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSet</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionSet: Specifies the customer managed disk encryption set resource id for the managed image disk.</p>
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
<p>DiskSizeGB: Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the
disk in a virtual machine image.
This value cannot be larger than 1023 GB</p>
</td>
</tr>
<tr>
<td>
<code>managedDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>ManagedDisk: The managedDisk.</p>
</td>
</tr>
<tr>
<td>
<code>osState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStatusOsState">
ImageOSDiskStatusOsState
</a>
</em>
</td>
<td>
<p>OsState: The OS State.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskStatusOsType">
ImageOSDiskStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: This property allows you to specify the type of the OS that is included in the disk if creating a VM from a
custom image.
Possible values are:
Windows
Linux</p>
</td>
</tr>
<tr>
<td>
<code>snapshot</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>Snapshot: The snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountType</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.StorageAccountType_Status">
StorageAccountType_Status
</a>
</em>
</td>
<td>
<p>StorageAccountType: Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with
data disks, it cannot be used with OS Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImagePropertiesARM">ImagePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Images_SpecARM">Images_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageProperties">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageProperties</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20210701.ImagePropertiesHyperVGeneration">
ImagePropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
resource.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResourceARM">
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
<a href="#compute.azure.com/v1beta20210701.ImageStorageProfileARM">
ImageStorageProfileARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Describes a storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImagePropertiesHyperVGeneration">ImagePropertiesHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImagePropertiesARM">ImagePropertiesARM</a>, <a href="#compute.azure.com/v1beta20210701.Images_Spec">Images_Spec</a>)
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
<tbody><tr><td><p>&#34;V1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;V2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageProperties_StatusARM">ImageProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Image_StatusARM">Image_StatusARM</a>)
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
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.HyperVGenerationType_Status">
HyperVGenerationType_Status
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
resource.</p>
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
<p>ProvisioningState: The provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_StatusARM">
SubResource_StatusARM
</a>
</em>
</td>
<td>
<p>SourceVirtualMachine: The source virtual machine from which Image is created.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageStorageProfile_StatusARM">
ImageStorageProfile_StatusARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageStorageProfile">ImageStorageProfile
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageStorageProfile">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageStorageProfile</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20210701.ImageDataDisk">
[]ImageDataDisk
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDisk">
ImageOSDisk
</a>
</em>
</td>
<td>
<p>OsDisk: Describes an Operating System disk.</p>
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
<p>ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
created only in regions that provide Zone Redundant Storage (ZRS).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageStorageProfileARM">ImageStorageProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImagePropertiesARM">ImagePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageStorageProfile">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageStorageProfile</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20210701.ImageDataDiskARM">
[]ImageDataDiskARM
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">
ImageOSDiskARM
</a>
</em>
</td>
<td>
<p>OsDisk: Describes an Operating System disk.</p>
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
<p>ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
created only in regions that provide Zone Redundant Storage (ZRS).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageStorageProfile_Status">ImageStorageProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Image_Status">Image_Status</a>)
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
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageDataDisk_Status">
[]ImageDataDisk_Status
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDisk_Status">
ImageOSDisk_Status
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview">About disks and VHDs for Azure virtual
machines</a>.</p>
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
<p>ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
created only in regions that provide Zone Redundant Storage (ZRS).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.ImageStorageProfile_StatusARM">ImageStorageProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageProperties_StatusARM">ImageProperties_StatusARM</a>)
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
<code>dataDisks</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageDataDisk_StatusARM">
[]ImageDataDisk_StatusARM
</a>
</em>
</td>
<td>
<p>DataDisks: Specifies the parameters that are used to add a data disk to a virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview">About disks and VHDs for Azure virtual
machines</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDisk</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">
ImageOSDisk_StatusARM
</a>
</em>
</td>
<td>
<p>OsDisk: Specifies information about the operating system disk used by the virtual machine.
For more information about disks, see <a href="https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview">About disks and VHDs for Azure virtual
machines</a>.</p>
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
<p>ZoneResilient: Specifies whether an image is zone resilient or not. Default is false. Zone resilient images can be
created only in regions that provide Zone Redundant Storage (ZRS).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.Image_Status">Image_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Image">Image</a>)
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
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Image.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.HyperVGenerationType_Status">
HyperVGenerationType_Status
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
resource.</p>
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
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
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>sourceVirtualMachine</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.SubResource_Status">
SubResource_Status
</a>
</em>
</td>
<td>
<p>SourceVirtualMachine: The source virtual machine from which Image is created.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImageStorageProfile_Status">
ImageStorageProfile_Status
</a>
</em>
</td>
<td>
<p>StorageProfile: Specifies the storage settings for the virtual machine disks.</p>
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
<h3 id="compute.azure.com/v1beta20210701.Image_StatusARM">Image_StatusARM
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
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Image.</p>
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location</p>
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
<a href="#compute.azure.com/v1beta20210701.ImageProperties_StatusARM">
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
<h3 id="compute.azure.com/v1beta20210701.ImagesSpecAPIVersion">ImagesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-07-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.Images_Spec">Images_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.Image">Image</a>)
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
<a href="#compute.azure.com/v1beta20210701.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImagePropertiesHyperVGeneration">
ImagePropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: Specifies the HyperVGenerationType of the VirtualMachine created from the image. From API Version
2019-03-01 if the image source is a blob, then we need the user to specify the value, if the source is managed resource
like disk or snapshot, we may require the user to specify the property if we cannot deduce it from the source managed
resource.</p>
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
<a href="#compute.azure.com/v1beta20210701.SubResource">
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
<a href="#compute.azure.com/v1beta20210701.ImageStorageProfile">
ImageStorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Describes a storage profile.</p>
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
<h3 id="compute.azure.com/v1beta20210701.Images_SpecARM">Images_SpecARM
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
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the image.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20210701.ImagePropertiesARM">
ImagePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Describes the properties of an Image.</p>
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
<h3 id="compute.azure.com/v1beta20210701.StorageAccountType_Status">StorageAccountType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk_Status">ImageDataDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>)
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
<h3 id="compute.azure.com/v1beta20210701.SubResource">SubResource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk">ImageDataDisk</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk">ImageOSDisk</a>, <a href="#compute.azure.com/v1beta20210701.Images_Spec">Images_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/SubResource">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/SubResource</a></p>
</div>
<table>
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
<p>Reference: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.SubResourceARM">SubResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDiskARM">ImageDataDiskARM</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDiskARM">ImageOSDiskARM</a>, <a href="#compute.azure.com/v1beta20210701.ImagePropertiesARM">ImagePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/SubResource">https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/SubResource</a></p>
</div>
<table>
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
<h3 id="compute.azure.com/v1beta20210701.SubResource_Status">SubResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk_Status">ImageDataDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_Status">ImageOSDisk_Status</a>, <a href="#compute.azure.com/v1beta20210701.Image_Status">Image_Status</a>)
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
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20210701.SubResource_StatusARM">SubResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20210701.ImageDataDisk_StatusARM">ImageDataDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20210701.ImageOSDisk_StatusARM">ImageOSDisk_StatusARM</a>, <a href="#compute.azure.com/v1beta20210701.ImageProperties_StatusARM">ImageProperties_StatusARM</a>)
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
</tbody>
</table>
<hr/>
