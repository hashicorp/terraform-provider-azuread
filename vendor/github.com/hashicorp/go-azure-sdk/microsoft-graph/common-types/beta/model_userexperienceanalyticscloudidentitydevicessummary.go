package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCloudIdentityDevicesSummary struct {
	// The count of devices that are not cloud identity. Read-only.
	DeviceWithoutCloudIdentityCount *int64 `json:"deviceWithoutCloudIdentityCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = UserExperienceAnalyticsCloudIdentityDevicesSummary{}

func (s UserExperienceAnalyticsCloudIdentityDevicesSummary) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsCloudIdentityDevicesSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsCloudIdentityDevicesSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsCloudIdentityDevicesSummary: %+v", err)
	}

	delete(decoded, "deviceWithoutCloudIdentityCount")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsCloudIdentityDevicesSummary: %+v", err)
	}

	return encoded, nil
}
