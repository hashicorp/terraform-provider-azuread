package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceOperatingSystemSummary struct {
	// The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid
	// values -1 to 2147483647
	AndroidCorporateWorkProfileCount *int64 `json:"androidCorporateWorkProfileCount,omitempty"`

	// Number of android device count.
	AndroidCount *int64 `json:"androidCount,omitempty"`

	// Number of dedicated Android devices.
	AndroidDedicatedCount *int64 `json:"androidDedicatedCount,omitempty"`

	// Number of device admin Android devices.
	AndroidDeviceAdminCount *int64 `json:"androidDeviceAdminCount,omitempty"`

	// Number of fully managed Android devices.
	AndroidFullyManagedCount *int64 `json:"androidFullyManagedCount,omitempty"`

	// Number of work profile Android devices.
	AndroidWorkProfileCount *int64 `json:"androidWorkProfileCount,omitempty"`

	// Number of ConfigMgr managed devices.
	ConfigMgrDeviceCount *int64 `json:"configMgrDeviceCount,omitempty"`

	// Number of iOS device count.
	IosCount *int64 `json:"iosCount,omitempty"`

	// Number of Mac OS X device count.
	MacOSCount *int64 `json:"macOSCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of unknown device count.
	UnknownCount *int64 `json:"unknownCount,omitempty"`

	// Number of Windows device count.
	WindowsCount *int64 `json:"windowsCount,omitempty"`

	// Number of Windows mobile device count.
	WindowsMobileCount *int64 `json:"windowsMobileCount,omitempty"`
}
