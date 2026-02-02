package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWindows10DevicesSummary struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The count of Windows 10 devices that have unsupported OS versions. Read-only.
	UnsupportedOSversionDeviceCount *int64 `json:"unsupportedOSversionDeviceCount,omitempty"`
}

var _ json.Marshaler = UserExperienceAnalyticsWindows10DevicesSummary{}

func (s UserExperienceAnalyticsWindows10DevicesSummary) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsWindows10DevicesSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsWindows10DevicesSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsWindows10DevicesSummary: %+v", err)
	}

	delete(decoded, "unsupportedOSversionDeviceCount")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsWindows10DevicesSummary: %+v", err)
	}

	return encoded, nil
}
