package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDevicesSummary struct {
	// The user experience analytics work from anywhere Autopilot devices summary. Read-only.
	AutopilotDevicesSummary *UserExperienceAnalyticsAutopilotDevicesSummary `json:"autopilotDevicesSummary,omitempty"`

	// The user experience analytics work from anywhere Cloud Identity devices summary. Read-only.
	CloudIdentityDevicesSummary *UserExperienceAnalyticsCloudIdentityDevicesSummary `json:"cloudIdentityDevicesSummary,omitempty"`

	// The user experience analytics work from anywhere Cloud management devices summary. Read-only.
	CloudManagementDevicesSummary *UserExperienceAnalyticsCloudManagementDevicesSummary `json:"cloudManagementDevicesSummary,omitempty"`

	// Total number of co-managed devices. Read-only. Valid values -2147483648 to 2147483647
	CoManagedDevices *int64 `json:"coManagedDevices,omitempty"`

	// The count of intune devices that are not autopilot registerd. Read-only. Valid values -2147483648 to 2147483647
	DevicesNotAutopilotRegistered *int64 `json:"devicesNotAutopilotRegistered,omitempty"`

	// The count of intune devices not autopilot profile assigned. Read-only. Valid values -2147483648 to 2147483647
	DevicesWithoutAutopilotProfileAssigned *int64 `json:"devicesWithoutAutopilotProfileAssigned,omitempty"`

	// The count of devices that are not cloud identity. Read-only. Valid values -2147483648 to 2147483647
	DevicesWithoutCloudIdentity *int64 `json:"devicesWithoutCloudIdentity,omitempty"`

	// The count of intune devices that are not autopilot registerd. Read-only. Valid values -2147483648 to 2147483647
	IntuneDevices *int64 `json:"intuneDevices,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total count of tenant attach devices. Read-only. Valid values -2147483648 to 2147483647
	TenantAttachDevices *int64 `json:"tenantAttachDevices,omitempty"`

	// The total count of devices. Read-only. Valid values -2147483648 to 2147483647
	TotalDevices *int64 `json:"totalDevices,omitempty"`

	// The count of Windows 10 devices that have unsupported OS versions. Read-only. Valid values -2147483648 to 2147483647
	UnsupportedOSversionDevices *int64 `json:"unsupportedOSversionDevices,omitempty"`

	// The count of windows 10 devices. Read-only. Valid values -2147483648 to 2147483647
	Windows10Devices *int64 `json:"windows10Devices,omitempty"`

	// The user experience analytics work from anywhere Windows 10 devices summary. Read-only.
	Windows10DevicesSummary *UserExperienceAnalyticsWindows10DevicesSummary `json:"windows10DevicesSummary,omitempty"`

	// The count of windows 10 devices that are Intune and co-managed. Read-only. Valid values -2147483648 to 2147483647
	Windows10DevicesWithoutTenantAttach *int64 `json:"windows10DevicesWithoutTenantAttach,omitempty"`
}

var _ json.Marshaler = UserExperienceAnalyticsWorkFromAnywhereDevicesSummary{}

func (s UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsWorkFromAnywhereDevicesSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsWorkFromAnywhereDevicesSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsWorkFromAnywhereDevicesSummary: %+v", err)
	}

	delete(decoded, "autopilotDevicesSummary")
	delete(decoded, "cloudIdentityDevicesSummary")
	delete(decoded, "cloudManagementDevicesSummary")
	delete(decoded, "coManagedDevices")
	delete(decoded, "devicesNotAutopilotRegistered")
	delete(decoded, "devicesWithoutAutopilotProfileAssigned")
	delete(decoded, "devicesWithoutCloudIdentity")
	delete(decoded, "intuneDevices")
	delete(decoded, "tenantAttachDevices")
	delete(decoded, "totalDevices")
	delete(decoded, "unsupportedOSversionDevices")
	delete(decoded, "windows10Devices")
	delete(decoded, "windows10DevicesSummary")
	delete(decoded, "windows10DevicesWithoutTenantAttach")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsWorkFromAnywhereDevicesSummary: %+v", err)
	}

	return encoded, nil
}
