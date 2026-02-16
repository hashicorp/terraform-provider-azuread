package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAppHealthApplicationPerformance{}

type UserExperienceAnalyticsAppHealthApplicationPerformance struct {
	// The health score of the application. Valid values 0 to 100. Supports: $filter, $select, $OrderBy. Read-only. Valid
	// values -2147483648 to 2147483647
	ActiveDeviceCount *int64 `json:"activeDeviceCount,omitempty"`

	// The number of crashes for the application. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only.
	// Valid values -2147483648 to 2147483647
	AppCrashCount *int64 `json:"appCrashCount,omitempty"`

	// The friendly name of the application. Possible values are: Outlook, Excel. Supports: $select, $OrderBy. Read-only.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The number of hangs for the application. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to
	// 2147483647
	AppHangCount *int64 `json:"appHangCount,omitempty"`

	// The name of the application. Possible values are: outlook.exe, excel.exe. Supports: $select, $OrderBy. Read-only.
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// The publisher of the application. Supports: $select, $OrderBy. Read-only.
	AppPublisher nullable.Type[string] `json:"appPublisher,omitempty"`

	// The total usage time of the application in minutes. Valid values 0 to 2147483647. Supports: $select, $OrderBy.
	// Read-only. Valid values -2147483648 to 2147483647
	AppUsageDuration *int64 `json:"appUsageDuration,omitempty"`

	// The mean time to failure for the application in minutes. Valid values 0 to 2147483647. Supports: $select, $OrderBy.
	// Read-only. Valid values -2147483648 to 2147483647
	MeanTimeToFailureInMinutes *int64 `json:"meanTimeToFailureInMinutes,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserExperienceAnalyticsAppHealthApplicationPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAppHealthApplicationPerformance{}

func (s UserExperienceAnalyticsAppHealthApplicationPerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAppHealthApplicationPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAppHealthApplicationPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAppHealthApplicationPerformance: %+v", err)
	}

	delete(decoded, "activeDeviceCount")
	delete(decoded, "appCrashCount")
	delete(decoded, "appDisplayName")
	delete(decoded, "appHangCount")
	delete(decoded, "appName")
	delete(decoded, "appPublisher")
	delete(decoded, "appUsageDuration")
	delete(decoded, "meanTimeToFailureInMinutes")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAppHealthApplicationPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAppHealthApplicationPerformance: %+v", err)
	}

	return encoded, nil
}
