---
---
<h2 id="compute.azure.com/v1beta20200930">compute.azure.com/v1beta20200930</h2>
<div>
<p>Package v1beta20200930 contains API Schema definitions for the compute v1beta20200930 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="compute.azure.com/v1beta20200930.CreationData">CreationData
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationDataCreateOption">
CreationDataCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: This enumerates the possible sources of a disk&rsquo;s creation.</p>
</td>
</tr>
<tr>
<td>
<code>galleryImageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReference">
ImageDiskReference
</a>
</em>
</td>
<td>
<p>GalleryImageReference: The source image used for creating the disk.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReference">
ImageDiskReference
</a>
</em>
</td>
<td>
<p>ImageReference: The source image used for creating the disk.</p>
</td>
</tr>
<tr>
<td>
<code>logicalSectorSize</code><br/>
<em>
int
</em>
</td>
<td>
<p>LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.</p>
</td>
</tr>
<tr>
<td>
<code>sourceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SourceResourceReference: If createOption is Copy, this is the ARM id of the source snapshot or disk.</p>
</td>
</tr>
<tr>
<td>
<code>sourceUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
containing the blob to import as a disk.</p>
</td>
</tr>
<tr>
<td>
<code>uploadSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
bytes for the VHD footer).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.CreationDataARM">CreationDataARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationDataCreateOption">
CreationDataCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: This enumerates the possible sources of a disk&rsquo;s creation.</p>
</td>
</tr>
<tr>
<td>
<code>galleryImageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReferenceARM">
ImageDiskReferenceARM
</a>
</em>
</td>
<td>
<p>GalleryImageReference: The source image used for creating the disk.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReferenceARM">
ImageDiskReferenceARM
</a>
</em>
</td>
<td>
<p>ImageReference: The source image used for creating the disk.</p>
</td>
</tr>
<tr>
<td>
<code>logicalSectorSize</code><br/>
<em>
int
</em>
</td>
<td>
<p>LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.</p>
</td>
</tr>
<tr>
<td>
<code>sourceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sourceUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
containing the blob to import as a disk.</p>
</td>
</tr>
<tr>
<td>
<code>uploadSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
bytes for the VHD footer).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.CreationDataCreateOption">CreationDataCreateOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.CreationData">CreationData</a>, <a href="#compute.azure.com/v1beta20200930.CreationDataARM">CreationDataARM</a>)
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
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Copy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Import&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Upload&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.CreationDataStatusCreateOption">CreationDataStatusCreateOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.CreationData_Status">CreationData_Status</a>, <a href="#compute.azure.com/v1beta20200930.CreationData_StatusARM">CreationData_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Attach&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Copy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Empty&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FromImage&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Import&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Upload&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.CreationData_Status">CreationData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationDataStatusCreateOption">
CreationDataStatusCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: This enumerates the possible sources of a disk&rsquo;s creation.</p>
</td>
</tr>
<tr>
<td>
<code>galleryImageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReference_Status">
ImageDiskReference_Status
</a>
</em>
</td>
<td>
<p>GalleryImageReference: Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of
the shared galley image version from which to create a disk.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReference_Status">
ImageDiskReference_Status
</a>
</em>
</td>
<td>
<p>ImageReference: Disk source information.</p>
</td>
</tr>
<tr>
<td>
<code>logicalSectorSize</code><br/>
<em>
int
</em>
</td>
<td>
<p>LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.</p>
</td>
</tr>
<tr>
<td>
<code>sourceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceResourceId: If createOption is Copy, this is the ARM id of the source snapshot or disk.</p>
</td>
</tr>
<tr>
<td>
<code>sourceUniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceUniqueId: If this field is set, this is the unique id identifying the source of this resource.</p>
</td>
</tr>
<tr>
<td>
<code>sourceUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
containing the blob to import as a disk.</p>
</td>
</tr>
<tr>
<td>
<code>uploadSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
bytes for the VHD footer).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.CreationData_StatusARM">CreationData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>)
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
<code>createOption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationDataStatusCreateOption">
CreationDataStatusCreateOption
</a>
</em>
</td>
<td>
<p>CreateOption: This enumerates the possible sources of a disk&rsquo;s creation.</p>
</td>
</tr>
<tr>
<td>
<code>galleryImageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReference_StatusARM">
ImageDiskReference_StatusARM
</a>
</em>
</td>
<td>
<p>GalleryImageReference: Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of
the shared galley image version from which to create a disk.</p>
</td>
</tr>
<tr>
<td>
<code>imageReference</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ImageDiskReference_StatusARM">
ImageDiskReference_StatusARM
</a>
</em>
</td>
<td>
<p>ImageReference: Disk source information.</p>
</td>
</tr>
<tr>
<td>
<code>logicalSectorSize</code><br/>
<em>
int
</em>
</td>
<td>
<p>LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.</p>
</td>
</tr>
<tr>
<td>
<code>sourceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceResourceId: If createOption is Copy, this is the ARM id of the source snapshot or disk.</p>
</td>
</tr>
<tr>
<td>
<code>sourceUniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceUniqueId: If this field is set, this is the unique id identifying the source of this resource.</p>
</td>
</tr>
<tr>
<td>
<code>sourceUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
containing the blob to import as a disk.</p>
</td>
</tr>
<tr>
<td>
<code>uploadSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
bytes for the VHD footer).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Disk">Disk
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/disks">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/disks</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.Disks_Spec">
Disks_Spec
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
<code>burstingEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
disabled by default. Does not apply to Ultra disks.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: Data used when creating a disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskAccessReference: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
operation can transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
per second - MB here uses the ISO notation, of powers of 10.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Encryption at rest settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection">
EncryptionSettingsCollection
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation">
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesHyperVGeneration">
DiskPropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
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
<code>maxShares</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
disk that can be mounted on multiple VMs at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesNetworkAccessPolicy">
DiskPropertiesNetworkAccessPolicy
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesOsType">
DiskPropertiesOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan">
PurchasePlan
</a>
</em>
</td>
<td>
<p>PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskSku">
DiskSku
</a>
</em>
</td>
<td>
<p>Sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.</p>
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
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Performance tier of the disk (e.g, P4, S10) as described here:
<a href="https://azure.microsoft.com/en-us/pricing/details/managed-disks/">https://azure.microsoft.com/en-us/pricing/details/managed-disks/</a>. Does not apply to Ultra disks.</p>
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
<p>Zones: The Logical zone list for Disk.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Disk_Status">
Disk_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_SpecARM">Disks_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskProperties">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>burstingEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
disabled by default. Does not apply to Ultra disks.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationDataARM">
CreationDataARM
</a>
</em>
</td>
<td>
<p>CreationData: Data used when creating a disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
operation can transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
per second - MB here uses the ISO notation, of powers of 10.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionARM">
EncryptionARM
</a>
</em>
</td>
<td>
<p>Encryption: Encryption at rest settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollectionARM">
EncryptionSettingsCollectionARM
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesHyperVGeneration">
DiskPropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
</td>
</tr>
<tr>
<td>
<code>maxShares</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
disk that can be mounted on multiple VMs at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesNetworkAccessPolicy">
DiskPropertiesNetworkAccessPolicy
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesOsType">
DiskPropertiesOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
</td>
</tr>
<tr>
<td>
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlanARM">
PurchasePlanARM
</a>
</em>
</td>
<td>
<p>PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Performance tier of the disk (e.g, P4, S10) as described here:
<a href="https://azure.microsoft.com/en-us/pricing/details/managed-disks/">https://azure.microsoft.com/en-us/pricing/details/managed-disks/</a>. Does not apply to Ultra disks.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskPropertiesHyperVGeneration">DiskPropertiesHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>)
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
<h3 id="compute.azure.com/v1beta20200930.DiskPropertiesNetworkAccessPolicy">DiskPropertiesNetworkAccessPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>)
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
<tbody><tr><td><p>&#34;AllowAll&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllowPrivate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DenyAll&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskPropertiesOsType">DiskPropertiesOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>)
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
<h3 id="compute.azure.com/v1beta20200930.DiskPropertiesStatusHyperVGeneration">DiskPropertiesStatusHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>)
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
<h3 id="compute.azure.com/v1beta20200930.DiskPropertiesStatusOsType">DiskPropertiesStatusOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>)
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
<h3 id="compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_StatusARM">Disk_StatusARM</a>)
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
<code>burstingEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
disabled by default. Does not apply to Ultra disks.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData_StatusARM">
CreationData_StatusARM
</a>
</em>
</td>
<td>
<p>CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
operation can transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeBytes: The size of the disk in bytes. This field is read only.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskState_Status">
DiskState_Status
</a>
</em>
</td>
<td>
<p>DiskState: The state of the disk.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption_StatusARM">
Encryption_StatusARM
</a>
</em>
</td>
<td>
<p>Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection_StatusARM">
EncryptionSettingsCollection_StatusARM
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings collection used for Azure Disk Encryption, can contain multiple
encryption settings per disk or snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesStatusHyperVGeneration">
DiskPropertiesStatusHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
</td>
</tr>
<tr>
<td>
<code>maxShares</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
disk that can be mounted on multiple VMs at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.NetworkAccessPolicy_Status">
NetworkAccessPolicy_Status
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesStatusOsType">
DiskPropertiesStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<p>ProvisioningState: The disk provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan_StatusARM">
PurchasePlan_StatusARM
</a>
</em>
</td>
<td>
<p>PurchasePlan: Purchase plan information for the the image from which the OS disk was created. E.g. - {name:
2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}</p>
</td>
</tr>
<tr>
<td>
<code>shareInfo</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ShareInfoElement_StatusARM">
[]ShareInfoElement_StatusARM
</a>
</em>
</td>
<td>
<p>ShareInfo: Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than
one for disks to allow attaching them to multiple VMs.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Performance tier of the disk (e.g, P4, S10) as described here:
<a href="https://azure.microsoft.com/en-us/pricing/details/managed-disks/">https://azure.microsoft.com/en-us/pricing/details/managed-disks/</a>. Does not apply to Ultra disks.</p>
</td>
</tr>
<tr>
<td>
<code>timeCreated</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeCreated: The time when the disk was created.</p>
</td>
</tr>
<tr>
<td>
<code>uniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>UniqueId: Unique Guid identifying the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskSku">DiskSku
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.DiskSkuName">
DiskSkuName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskSkuARM">DiskSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_SpecARM">Disks_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.DiskSkuName">
DiskSkuName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskSkuName">DiskSkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskSku">DiskSku</a>, <a href="#compute.azure.com/v1beta20200930.DiskSkuARM">DiskSkuARM</a>)
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
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskSkuStatusName">DiskSkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskSku_Status">DiskSku_Status</a>, <a href="#compute.azure.com/v1beta20200930.DiskSku_StatusARM">DiskSku_StatusARM</a>)
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
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardSSD_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UltraSSD_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskSku_Status">DiskSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>)
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
<a href="#compute.azure.com/v1beta20200930.DiskSkuStatusName">
DiskSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: The sku tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskSku_StatusARM">DiskSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_StatusARM">Disk_StatusARM</a>)
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
<a href="#compute.azure.com/v1beta20200930.DiskSkuStatusName">
DiskSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: The sku tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DiskState_Status">DiskState_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<tbody><tr><td><p>&#34;ActiveSAS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ActiveUpload&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Attached&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadyToUpload&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Reserved&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unattached&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Disk_Status">Disk_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk">Disk</a>)
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
<code>burstingEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
disabled by default. Does not apply to Ultra disks.</p>
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
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData_Status">
CreationData_Status
</a>
</em>
</td>
<td>
<p>CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
operation can transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeBytes: The size of the disk in bytes. This field is read only.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskState_Status">
DiskState_Status
</a>
</em>
</td>
<td>
<p>DiskState: The state of the disk.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption_Status">
Encryption_Status
</a>
</em>
</td>
<td>
<p>Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection_Status">
EncryptionSettingsCollection_Status
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings collection used for Azure Disk Encryption, can contain multiple
encryption settings per disk or snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location where the disk will be created. Extended location cannot be changed.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesStatusHyperVGeneration">
DiskPropertiesStatusHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
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
<code>managedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>ManagedBy: A relative URI containing the ID of the VM that has the disk attached.</p>
</td>
</tr>
<tr>
<td>
<code>managedByExtended</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ManagedByExtended: List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be
set to a value greater than one for disks to allow attaching them to multiple VMs.</p>
</td>
</tr>
<tr>
<td>
<code>maxShares</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
disk that can be mounted on multiple VMs at the same time.</p>
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
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.NetworkAccessPolicy_Status">
NetworkAccessPolicy_Status
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesStatusOsType">
DiskPropertiesStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<p>ProvisioningState: The disk provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan_Status">
PurchasePlan_Status
</a>
</em>
</td>
<td>
<p>PurchasePlan: Purchase plan information for the the image from which the OS disk was created. E.g. - {name:
2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}</p>
</td>
</tr>
<tr>
<td>
<code>shareInfo</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ShareInfoElement_Status">
[]ShareInfoElement_Status
</a>
</em>
</td>
<td>
<p>ShareInfo: Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than
one for disks to allow attaching them to multiple VMs.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskSku_Status">
DiskSku_Status
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
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Performance tier of the disk (e.g, P4, S10) as described here:
<a href="https://azure.microsoft.com/en-us/pricing/details/managed-disks/">https://azure.microsoft.com/en-us/pricing/details/managed-disks/</a>. Does not apply to Ultra disks.</p>
</td>
</tr>
<tr>
<td>
<code>timeCreated</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeCreated: The time when the disk was created.</p>
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
<code>uniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>UniqueId: Unique Guid identifying the resource.</p>
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
<p>Zones: The Logical zone list for Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Disk_StatusARM">Disk_StatusARM
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location where the disk will be created. Extended location cannot be changed.</p>
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
<code>managedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>ManagedBy: A relative URI containing the ID of the VM that has the disk attached.</p>
</td>
</tr>
<tr>
<td>
<code>managedByExtended</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ManagedByExtended: List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be
set to a value greater than one for disks to allow attaching them to multiple VMs.</p>
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
<a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">
DiskProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskSku_StatusARM">
DiskSku_StatusARM
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
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: The Logical zone list for Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.DisksSpecAPIVersion">DisksSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-09-30&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk">Disk</a>)
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
<code>burstingEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
disabled by default. Does not apply to Ultra disks.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: Data used when creating a disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskAccessReference: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
operation can transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskIOPSReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
transfer between 4k and 256k bytes.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadOnly</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.</p>
</td>
</tr>
<tr>
<td>
<code>diskMBpsReadWrite</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
per second - MB here uses the ISO notation, of powers of 10.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Encryption at rest settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection">
EncryptionSettingsCollection
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation">
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesHyperVGeneration">
DiskPropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
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
<code>maxShares</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
disk that can be mounted on multiple VMs at the same time.</p>
</td>
</tr>
<tr>
<td>
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesNetworkAccessPolicy">
DiskPropertiesNetworkAccessPolicy
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
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesOsType">
DiskPropertiesOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan">
PurchasePlan
</a>
</em>
</td>
<td>
<p>PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskSku">
DiskSku
</a>
</em>
</td>
<td>
<p>Sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.</p>
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
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: Performance tier of the disk (e.g, P4, S10) as described here:
<a href="https://azure.microsoft.com/en-us/pricing/details/managed-disks/">https://azure.microsoft.com/en-us/pricing/details/managed-disks/</a>. Does not apply to Ultra disks.</p>
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
<p>Zones: The Logical zone list for Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Disks_SpecARM">Disks_SpecARM
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocationARM">
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
<p>Name: The name of the managed disk that is being created. The name can&rsquo;t be changed after the disk is created. Supported
characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">
DiskPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Disk resource properties.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskSkuARM">
DiskSkuARM
</a>
</em>
</td>
<td>
<p>Sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.</p>
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
<p>Zones: The Logical zone list for Disk.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Encryption">Encryption
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionSetReference: ResourceId of the disk encryption set to use for enabling encryption at rest.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionType">
EncryptionType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionARM">EncryptionARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionSetId</code><br/>
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
<a href="#compute.azure.com/v1beta20200930.EncryptionType">
EncryptionType
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsCollection">EncryptionSettingsCollection
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection</a></p>
</div>
<table>
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
bool
</em>
</td>
<td>
<p>Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
null in the request object, the existing settings remain unchanged.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement">
[]EncryptionSettingsElement
</a>
</em>
</td>
<td>
<p>EncryptionSettings: A collection of encryption settings, one for each disk volume.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
overwritten. &lsquo;1.0&rsquo; corresponds to Azure Disk Encryption with AAD app.&lsquo;1.1&rsquo; corresponds to Azure Disk Encryption.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsCollectionARM">EncryptionSettingsCollectionARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection</a></p>
</div>
<table>
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
bool
</em>
</td>
<td>
<p>Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
null in the request object, the existing settings remain unchanged.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElementARM">
[]EncryptionSettingsElementARM
</a>
</em>
</td>
<td>
<p>EncryptionSettings: A collection of encryption settings, one for each disk volume.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
overwritten. &lsquo;1.0&rsquo; corresponds to Azure Disk Encryption with AAD app.&lsquo;1.1&rsquo; corresponds to Azure Disk Encryption.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsCollection_Status">EncryptionSettingsCollection_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
bool
</em>
</td>
<td>
<p>Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
null in the request object, the existing settings remain unchanged.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement_Status">
[]EncryptionSettingsElement_Status
</a>
</em>
</td>
<td>
<p>EncryptionSettings: A collection of encryption settings, one for each disk volume.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
overwritten. &lsquo;1.0&rsquo; corresponds to Azure Disk Encryption with AAD app.&lsquo;1.1&rsquo; corresponds to Azure Disk Encryption.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsCollection_StatusARM">EncryptionSettingsCollection_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>)
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
bool
</em>
</td>
<td>
<p>Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
null in the request object, the existing settings remain unchanged.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettings</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement_StatusARM">
[]EncryptionSettingsElement_StatusARM
</a>
</em>
</td>
<td>
<p>EncryptionSettings: A collection of encryption settings, one for each disk volume.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
overwritten. &lsquo;1.0&rsquo; corresponds to Azure Disk Encryption with AAD app.&lsquo;1.1&rsquo; corresponds to Azure Disk Encryption.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsElement">EncryptionSettingsElement
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection">EncryptionSettingsCollection</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReference">
KeyVaultAndSecretReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Key Vault Secret Url and vault id of the encryption key</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReference">
KeyVaultAndKeyReference
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the
encryptionKey</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsElementARM">EncryptionSettingsElementARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollectionARM">EncryptionSettingsCollectionARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReferenceARM">
KeyVaultAndSecretReferenceARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Key Vault Secret Url and vault id of the encryption key</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReferenceARM">
KeyVaultAndKeyReferenceARM
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the
encryptionKey</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsElement_Status">EncryptionSettingsElement_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection_Status">EncryptionSettingsCollection_Status</a>)
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
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReference_Status">
KeyVaultAndSecretReference_Status
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Key Vault Secret Url and vault id of the disk encryption key</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReference_Status">
KeyVaultAndKeyReference_Status
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when
provided is used to unwrap the disk encryption key.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionSettingsElement_StatusARM">EncryptionSettingsElement_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection_StatusARM">EncryptionSettingsCollection_StatusARM</a>)
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
<code>diskEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReference_StatusARM">
KeyVaultAndSecretReference_StatusARM
</a>
</em>
</td>
<td>
<p>DiskEncryptionKey: Key Vault Secret Url and vault id of the disk encryption key</p>
</td>
</tr>
<tr>
<td>
<code>keyEncryptionKey</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReference_StatusARM">
KeyVaultAndKeyReference_StatusARM
</a>
</em>
</td>
<td>
<p>KeyEncryptionKey: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when
provided is used to unwrap the disk encryption key.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionType">EncryptionType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Encryption">Encryption</a>, <a href="#compute.azure.com/v1beta20200930.EncryptionARM">EncryptionARM</a>)
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
<tbody><tr><td><p>&#34;EncryptionAtRestWithCustomerKey&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EncryptionAtRestWithPlatformAndCustomerKeys&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EncryptionAtRestWithPlatformKey&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.EncryptionType_Status">EncryptionType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Encryption_Status">Encryption_Status</a>, <a href="#compute.azure.com/v1beta20200930.Encryption_StatusARM">Encryption_StatusARM</a>)
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
<tbody><tr><td><p>&#34;EncryptionAtRestWithCustomerKey&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EncryptionAtRestWithPlatformAndCustomerKeys&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;EncryptionAtRestWithPlatformKey&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Encryption_Status">Encryption_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<code>diskEncryptionSetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskEncryptionSetId: ResourceId of the disk encryption set to use for enabling encryption at rest.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionType_Status">
EncryptionType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Encryption_StatusARM">Encryption_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>)
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
<code>diskEncryptionSetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskEncryptionSetId: ResourceId of the disk encryption set to use for enabling encryption at rest.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionType_Status">
EncryptionType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocationType">
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
<h3 id="compute.azure.com/v1beta20200930.ExtendedLocationARM">ExtendedLocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_SpecARM">Disks_SpecARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_SpecARM">Snapshots_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocationType">
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
<h3 id="compute.azure.com/v1beta20200930.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.ExtendedLocation">ExtendedLocation</a>, <a href="#compute.azure.com/v1beta20200930.ExtendedLocationARM">ExtendedLocationARM</a>)
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
<h3 id="compute.azure.com/v1beta20200930.ExtendedLocationType_Status">ExtendedLocationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.ExtendedLocation_Status">ExtendedLocation_Status</a>, <a href="#compute.azure.com/v1beta20200930.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM</a>)
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
<h3 id="compute.azure.com/v1beta20200930.ExtendedLocation_Status">ExtendedLocation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocationType_Status">
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
<h3 id="compute.azure.com/v1beta20200930.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_StatusARM">Disk_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_StatusARM">Snapshot_StatusARM</a>)
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocationType_Status">
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
<h3 id="compute.azure.com/v1beta20200930.ImageDiskReference">ImageDiskReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.CreationData">CreationData</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: If the disk is created from an image&rsquo;s data disk, this is an index that indicates which of the data disks in the
image to use. For OS disks, this field is null.</p>
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
<p>Reference: A relative uri containing either a Platform Image Repository or user image reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.ImageDiskReferenceARM">ImageDiskReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.CreationDataARM">CreationDataARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference</a></p>
</div>
<table>
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
<code>lun</code><br/>
<em>
int
</em>
</td>
<td>
<p>Lun: If the disk is created from an image&rsquo;s data disk, this is an index that indicates which of the data disks in the
image to use. For OS disks, this field is null.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.ImageDiskReference_Status">ImageDiskReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.CreationData_Status">CreationData_Status</a>)
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
<p>Id: A relative uri containing either a Platform Image Repository or user image reference.</p>
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
<p>Lun: If the disk is created from an image&rsquo;s data disk, this is an index that indicates which of the data disks in the
image to use. For OS disks, this field is null.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.ImageDiskReference_StatusARM">ImageDiskReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.CreationData_StatusARM">CreationData_StatusARM</a>)
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
<p>Id: A relative uri containing either a Platform Image Repository or user image reference.</p>
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
<p>Lun: If the disk is created from an image&rsquo;s data disk, this is an index that indicates which of the data disks in the
image to use. For OS disks, this field is null.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndKeyReference">KeyVaultAndKeyReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement">EncryptionSettingsElement</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVault">
SourceVault
</a>
</em>
</td>
<td>
<p>SourceVault: The vault id is an Azure Resource Manager Resource id in the form
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndKeyReferenceARM">KeyVaultAndKeyReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElementARM">EncryptionSettingsElementARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVaultARM">
SourceVaultARM
</a>
</em>
</td>
<td>
<p>SourceVault: The vault id is an Azure Resource Manager Resource id in the form
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndKeyReference_Status">KeyVaultAndKeyReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement_Status">EncryptionSettingsElement_Status</a>)
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
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVault_Status">
SourceVault_Status
</a>
</em>
</td>
<td>
<p>SourceVault: Resource id of the KeyVault containing the key or secret</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndKeyReference_StatusARM">KeyVaultAndKeyReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement_StatusARM">EncryptionSettingsElement_StatusARM</a>)
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
<code>keyUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVault_StatusARM">
SourceVault_StatusARM
</a>
</em>
</td>
<td>
<p>SourceVault: Resource id of the KeyVault containing the key or secret</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndSecretReference">KeyVaultAndSecretReference
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement">EncryptionSettingsElement</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVault">
SourceVault
</a>
</em>
</td>
<td>
<p>SourceVault: The vault id is an Azure Resource Manager Resource id in the form
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndSecretReferenceARM">KeyVaultAndSecretReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElementARM">EncryptionSettingsElementARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVaultARM">
SourceVaultARM
</a>
</em>
</td>
<td>
<p>SourceVault: The vault id is an Azure Resource Manager Resource id in the form
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndSecretReference_Status">KeyVaultAndSecretReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement_Status">EncryptionSettingsElement_Status</a>)
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
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVault_Status">
SourceVault_Status
</a>
</em>
</td>
<td>
<p>SourceVault: Resource id of the KeyVault containing the key or secret</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.KeyVaultAndSecretReference_StatusARM">KeyVaultAndSecretReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.EncryptionSettingsElement_StatusARM">EncryptionSettingsElement_StatusARM</a>)
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
<code>secretUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretUrl: Url pointing to a key or secret in KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>sourceVault</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SourceVault_StatusARM">
SourceVault_StatusARM
</a>
</em>
</td>
<td>
<p>SourceVault: Resource id of the KeyVault containing the key or secret</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.NetworkAccessPolicy_Status">NetworkAccessPolicy_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<tbody><tr><td><p>&#34;AllowAll&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllowPrivate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DenyAll&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.PurchasePlan">PurchasePlan
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disks_Spec">Disks_Spec</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan</a></p>
</div>
<table>
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
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The Offer Promotion Code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.PurchasePlanARM">PurchasePlanARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskPropertiesARM">DiskPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan</a></p>
</div>
<table>
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
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The Offer Promotion Code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.PurchasePlan_Status">PurchasePlan_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The Offer Promotion Code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.PurchasePlan_StatusARM">PurchasePlan_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>)
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
<p>Name: The plan ID.</p>
</td>
</tr>
<tr>
<td>
<code>product</code><br/>
<em>
string
</em>
</td>
<td>
<p>Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
imageReference element.</p>
</td>
</tr>
<tr>
<td>
<code>promotionCode</code><br/>
<em>
string
</em>
</td>
<td>
<p>PromotionCode: The Offer Promotion Code.</p>
</td>
</tr>
<tr>
<td>
<code>publisher</code><br/>
<em>
string
</em>
</td>
<td>
<p>Publisher: The publisher ID.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.ShareInfoElement_Status">ShareInfoElement_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Disk_Status">Disk_Status</a>)
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
<code>vmUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmUri: A relative URI containing the ID of the VM that has the disk attached.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.ShareInfoElement_StatusARM">ShareInfoElement_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.DiskProperties_StatusARM">DiskProperties_StatusARM</a>)
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
<code>vmUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmUri: A relative URI containing the ID of the VM that has the disk attached.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Snapshot">Snapshot
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/snapshots">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/resourceDefinitions/snapshots</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">
Snapshots_Spec
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
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: Data used when creating a disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskAccessReference: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesDiskState">
SnapshotPropertiesDiskState
</a>
</em>
</td>
<td>
<p>DiskState: The state of the snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Encryption at rest settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection">
EncryptionSettingsCollection
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation">
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesHyperVGeneration">
SnapshotPropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
</td>
</tr>
<tr>
<td>
<code>incremental</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
snapshots and can be diffed.</p>
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
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesNetworkAccessPolicy">
SnapshotPropertiesNetworkAccessPolicy
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesOsType">
SnapshotPropertiesOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan">
PurchasePlan
</a>
</em>
</td>
<td>
<p>PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotSku">
SnapshotSku
</a>
</em>
</td>
<td>
<p>Sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for
incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot</p>
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
<a href="#compute.azure.com/v1beta20200930.Snapshot_Status">
Snapshot_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshots_SpecARM">Snapshots_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotProperties">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationDataARM">
CreationDataARM
</a>
</em>
</td>
<td>
<p>CreationData: Data used when creating a disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessId</code><br/>
<em>
string
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesDiskState">
SnapshotPropertiesDiskState
</a>
</em>
</td>
<td>
<p>DiskState: The state of the snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionARM">
EncryptionARM
</a>
</em>
</td>
<td>
<p>Encryption: Encryption at rest settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollectionARM">
EncryptionSettingsCollectionARM
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesHyperVGeneration">
SnapshotPropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
</td>
</tr>
<tr>
<td>
<code>incremental</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
snapshots and can be diffed.</p>
</td>
</tr>
<tr>
<td>
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesNetworkAccessPolicy">
SnapshotPropertiesNetworkAccessPolicy
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesOsType">
SnapshotPropertiesOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
</td>
</tr>
<tr>
<td>
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlanARM">
PurchasePlanARM
</a>
</em>
</td>
<td>
<p>PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesDiskState">SnapshotPropertiesDiskState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
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
<tbody><tr><td><p>&#34;ActiveSAS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ActiveUpload&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Attached&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadyToUpload&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Reserved&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unattached&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesHyperVGeneration">SnapshotPropertiesHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
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
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesNetworkAccessPolicy">SnapshotPropertiesNetworkAccessPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
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
<tbody><tr><td><p>&#34;AllowAll&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AllowPrivate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DenyAll&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesOsType">SnapshotPropertiesOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">SnapshotPropertiesARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
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
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesStatusHyperVGeneration">SnapshotPropertiesStatusHyperVGeneration
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<h3 id="compute.azure.com/v1beta20200930.SnapshotPropertiesStatusOsType">SnapshotPropertiesStatusOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<h3 id="compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">SnapshotProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshot_StatusARM">Snapshot_StatusARM</a>)
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
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData_StatusARM">
CreationData_StatusARM
</a>
</em>
</td>
<td>
<p>CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeBytes: The size of the disk in bytes. This field is read only.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskState_Status">
DiskState_Status
</a>
</em>
</td>
<td>
<p>DiskState: The state of the snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption_StatusARM">
Encryption_StatusARM
</a>
</em>
</td>
<td>
<p>Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection_StatusARM">
EncryptionSettingsCollection_StatusARM
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings collection used be Azure Disk Encryption, can contain multiple
encryption settings per disk or snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesStatusHyperVGeneration">
SnapshotPropertiesStatusHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
</td>
</tr>
<tr>
<td>
<code>incremental</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
snapshots and can be diffed.</p>
</td>
</tr>
<tr>
<td>
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.NetworkAccessPolicy_Status">
NetworkAccessPolicy_Status
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesStatusOsType">
SnapshotPropertiesStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<p>ProvisioningState: The disk provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan_StatusARM">
PurchasePlan_StatusARM
</a>
</em>
</td>
<td>
<p>PurchasePlan: Purchase plan information for the image from which the source disk for the snapshot was originally created.</p>
</td>
</tr>
<tr>
<td>
<code>timeCreated</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeCreated: The time when the snapshot was created.</p>
</td>
</tr>
<tr>
<td>
<code>uniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>UniqueId: Unique Guid identifying the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotSku">SnapshotSku
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotSku">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotSku</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.SnapshotSkuName">
SnapshotSkuName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotSkuARM">SnapshotSkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshots_SpecARM">Snapshots_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotSku">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SnapshotSku</a></p>
</div>
<table>
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
<a href="#compute.azure.com/v1beta20200930.SnapshotSkuName">
SnapshotSkuName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotSkuName">SnapshotSkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotSku">SnapshotSku</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotSkuARM">SnapshotSkuARM</a>)
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
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ZRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotSkuStatusName">SnapshotSkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.SnapshotSku_Status">SnapshotSku_Status</a>, <a href="#compute.azure.com/v1beta20200930.SnapshotSku_StatusARM">SnapshotSku_StatusARM</a>)
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
</tr><tr><td><p>&#34;Standard_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard_ZRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotSku_Status">SnapshotSku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status</a>)
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
<a href="#compute.azure.com/v1beta20200930.SnapshotSkuStatusName">
SnapshotSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: The sku tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.SnapshotSku_StatusARM">SnapshotSku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshot_StatusARM">Snapshot_StatusARM</a>)
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
<a href="#compute.azure.com/v1beta20200930.SnapshotSkuStatusName">
SnapshotSkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The sku name.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tier: The sku tier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Snapshot_Status">Snapshot_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshot">Snapshot</a>)
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
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData_Status">
CreationData_Status
</a>
</em>
</td>
<td>
<p>CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
</td>
</tr>
<tr>
<td>
<code>diskSizeBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>DiskSizeBytes: The size of the disk in bytes. This field is read only.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.DiskState_Status">
DiskState_Status
</a>
</em>
</td>
<td>
<p>DiskState: The state of the snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption_Status">
Encryption_Status
</a>
</em>
</td>
<td>
<p>Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection_Status">
EncryptionSettingsCollection_Status
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings collection used be Azure Disk Encryption, can contain multiple
encryption settings per disk or snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation_Status">
ExtendedLocation_Status
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location where the snapshot will be created. Extended location cannot be changed.</p>
</td>
</tr>
<tr>
<td>
<code>hyperVGeneration</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesStatusHyperVGeneration">
SnapshotPropertiesStatusHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
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
<code>incremental</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
snapshots and can be diffed.</p>
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
<code>managedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>ManagedBy: Unused. Always Null.</p>
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
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.NetworkAccessPolicy_Status">
NetworkAccessPolicy_Status
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesStatusOsType">
SnapshotPropertiesStatusOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<p>ProvisioningState: The disk provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan_Status">
PurchasePlan_Status
</a>
</em>
</td>
<td>
<p>PurchasePlan: Purchase plan information for the image from which the source disk for the snapshot was originally created.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotSku_Status">
SnapshotSku_Status
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
<code>timeCreated</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeCreated: The time when the snapshot was created.</p>
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
<code>uniqueId</code><br/>
<em>
string
</em>
</td>
<td>
<p>UniqueId: Unique Guid identifying the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Snapshot_StatusARM">Snapshot_StatusARM
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location where the snapshot will be created. Extended location cannot be changed.</p>
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
<code>managedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>ManagedBy: Unused. Always Null.</p>
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
<a href="#compute.azure.com/v1beta20200930.SnapshotProperties_StatusARM">
SnapshotProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotSku_StatusARM">
SnapshotSku_StatusARM
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
<h3 id="compute.azure.com/v1beta20200930.SnapshotsSpecAPIVersion">SnapshotsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-09-30&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="compute.azure.com/v1beta20200930.Snapshots_Spec">Snapshots_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.Snapshot">Snapshot</a>)
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
<code>creationData</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: Data used when creating a disk.</p>
</td>
</tr>
<tr>
<td>
<code>diskAccessReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskAccessReference: ARM id of the DiskAccess resource for using private endpoints on disks.</p>
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
<p>DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
allowed if the disk is not attached to a running VM, and can only increase the disk&rsquo;s size.</p>
</td>
</tr>
<tr>
<td>
<code>diskState</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesDiskState">
SnapshotPropertiesDiskState
</a>
</em>
</td>
<td>
<p>DiskState: The state of the snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Encryption at rest settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>encryptionSettingsCollection</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.EncryptionSettingsCollection">
EncryptionSettingsCollection
</a>
</em>
</td>
<td>
<p>EncryptionSettingsCollection: Encryption settings for disk or snapshot</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.ExtendedLocation">
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesHyperVGeneration">
SnapshotPropertiesHyperVGeneration
</a>
</em>
</td>
<td>
<p>HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.</p>
</td>
</tr>
<tr>
<td>
<code>incremental</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
snapshots and can be diffed.</p>
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
<code>networkAccessPolicy</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesNetworkAccessPolicy">
SnapshotPropertiesNetworkAccessPolicy
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
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesOsType">
SnapshotPropertiesOsType
</a>
</em>
</td>
<td>
<p>OsType: The Operating System type.</p>
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
<code>purchasePlan</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.PurchasePlan">
PurchasePlan
</a>
</em>
</td>
<td>
<p>PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotSku">
SnapshotSku
</a>
</em>
</td>
<td>
<p>Sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for
incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot</p>
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
<h3 id="compute.azure.com/v1beta20200930.Snapshots_SpecARM">Snapshots_SpecARM
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
<a href="#compute.azure.com/v1beta20200930.ExtendedLocationARM">
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
<p>Name: The name of the snapshot that is being created. The name can&rsquo;t be changed after the snapshot is created. Supported
characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotPropertiesARM">
SnapshotPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Snapshot resource properties.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#compute.azure.com/v1beta20200930.SnapshotSkuARM">
SnapshotSkuARM
</a>
</em>
</td>
<td>
<p>Sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for
incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot</p>
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
<h3 id="compute.azure.com/v1beta20200930.SourceVault">SourceVault
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReference">KeyVaultAndKeyReference</a>, <a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReference">KeyVaultAndSecretReference</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault</a></p>
</div>
<table>
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
<h3 id="compute.azure.com/v1beta20200930.SourceVaultARM">SourceVaultARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReferenceARM">KeyVaultAndKeyReferenceARM</a>, <a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReferenceARM">KeyVaultAndSecretReferenceARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault">https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault</a></p>
</div>
<table>
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
<h3 id="compute.azure.com/v1beta20200930.SourceVault_Status">SourceVault_Status
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReference_Status">KeyVaultAndKeyReference_Status</a>, <a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReference_Status">KeyVaultAndSecretReference_Status</a>)
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
<h3 id="compute.azure.com/v1beta20200930.SourceVault_StatusARM">SourceVault_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#compute.azure.com/v1beta20200930.KeyVaultAndKeyReference_StatusARM">KeyVaultAndKeyReference_StatusARM</a>, <a href="#compute.azure.com/v1beta20200930.KeyVaultAndSecretReference_StatusARM">KeyVaultAndSecretReference_StatusARM</a>)
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
