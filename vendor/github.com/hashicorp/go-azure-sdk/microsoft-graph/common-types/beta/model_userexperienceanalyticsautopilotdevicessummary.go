package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAutopilotDevicesSummary struct {
	// The count of intune devices that are not autopilot registerd. Read-only.
	DevicesNotAutopilotRegistered *int64 `json:"devicesNotAutopilotRegistered,omitempty"`

	// The count of intune devices not autopilot profile assigned. Read-only.
	DevicesWithoutAutopilotProfileAssigned *int64 `json:"devicesWithoutAutopilotProfileAssigned,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The count of windows 10 devices that are Intune and co-managed. Read-only.
	TotalWindows10DevicesWithoutTenantAttached *int64 `json:"totalWindows10DevicesWithoutTenantAttached,omitempty"`
}

var _ json.Marshaler = UserExperienceAnalyticsAutopilotDevicesSummary{}

func (s UserExperienceAnalyticsAutopilotDevicesSummary) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAutopilotDevicesSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAutopilotDevicesSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAutopilotDevicesSummary: %+v", err)
	}

	delete(decoded, "devicesNotAutopilotRegistered")
	delete(decoded, "devicesWithoutAutopilotProfileAssigned")
	delete(decoded, "totalWindows10DevicesWithoutTenantAttached")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAutopilotDevicesSummary: %+v", err)
	}

	return encoded, nil
}
