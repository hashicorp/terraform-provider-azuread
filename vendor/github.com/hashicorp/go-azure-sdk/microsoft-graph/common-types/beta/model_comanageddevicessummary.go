package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDevicesSummary struct {
	// Number of devices with CompliancePolicy swung-over. This property is read-only.
	CompliancePolicyCount *int64 `json:"compliancePolicyCount,omitempty"`

	// Number of devices with ConfigurationSettings swung-over. This property is read-only.
	ConfigurationSettingsCount *int64 `json:"configurationSettingsCount,omitempty"`

	// Number of devices with EndpointProtection swung-over. This property is read-only.
	EndpointProtectionCount *int64 `json:"endpointProtectionCount,omitempty"`

	// Number of devices with Inventory swung-over. This property is read-only.
	InventoryCount *int64 `json:"inventoryCount,omitempty"`

	// Number of devices with ModernApps swung-over. This property is read-only.
	ModernAppsCount *int64 `json:"modernAppsCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of devices with OfficeApps swung-over. This property is read-only.
	OfficeAppsCount *int64 `json:"officeAppsCount,omitempty"`

	// Number of devices with ResourceAccess swung-over. This property is read-only.
	ResourceAccessCount *int64 `json:"resourceAccessCount,omitempty"`

	// Number of Co-Managed Devices. This property is read-only.
	TotalComanagedCount *int64 `json:"totalComanagedCount,omitempty"`

	// Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
	WindowsUpdateForBusinessCount *int64 `json:"windowsUpdateForBusinessCount,omitempty"`
}

var _ json.Marshaler = ComanagedDevicesSummary{}

func (s ComanagedDevicesSummary) MarshalJSON() ([]byte, error) {
	type wrapper ComanagedDevicesSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ComanagedDevicesSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ComanagedDevicesSummary: %+v", err)
	}

	delete(decoded, "compliancePolicyCount")
	delete(decoded, "configurationSettingsCount")
	delete(decoded, "endpointProtectionCount")
	delete(decoded, "inventoryCount")
	delete(decoded, "modernAppsCount")
	delete(decoded, "officeAppsCount")
	delete(decoded, "resourceAccessCount")
	delete(decoded, "totalComanagedCount")
	delete(decoded, "windowsUpdateForBusinessCount")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ComanagedDevicesSummary: %+v", err)
	}

	return encoded, nil
}
