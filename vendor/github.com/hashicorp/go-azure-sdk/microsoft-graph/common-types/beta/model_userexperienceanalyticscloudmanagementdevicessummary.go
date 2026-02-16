package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCloudManagementDevicesSummary struct {
	// Total number of co-managed devices. Read-only.
	CoManagedDeviceCount *int64 `json:"coManagedDeviceCount,omitempty"`

	// The count of intune devices that are not autopilot registerd. Read-only.
	IntuneDeviceCount *int64 `json:"intuneDeviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total count of tenant attach devices. Read-only.
	TenantAttachDeviceCount *int64 `json:"tenantAttachDeviceCount,omitempty"`
}

var _ json.Marshaler = UserExperienceAnalyticsCloudManagementDevicesSummary{}

func (s UserExperienceAnalyticsCloudManagementDevicesSummary) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsCloudManagementDevicesSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsCloudManagementDevicesSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsCloudManagementDevicesSummary: %+v", err)
	}

	delete(decoded, "coManagedDeviceCount")
	delete(decoded, "intuneDeviceCount")
	delete(decoded, "tenantAttachDeviceCount")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsCloudManagementDevicesSummary: %+v", err)
	}

	return encoded, nil
}
