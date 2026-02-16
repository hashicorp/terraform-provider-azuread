package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceStartupHistory{}

type UserExperienceAnalyticsDeviceStartupHistory struct {
	// The device core boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
	CoreBootTimeInMs *int64 `json:"coreBootTimeInMs,omitempty"`

	// The device core login time in milliseconds. Supports: $select, $OrderBy. Read-only.
	CoreLoginTimeInMs *int64 `json:"coreLoginTimeInMs,omitempty"`

	// The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The impact of device feature updates on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
	FeatureUpdateBootTimeInMs *int64 `json:"featureUpdateBootTimeInMs,omitempty"`

	// The impact of device group policy client on boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
	GroupPolicyBootTimeInMs *int64 `json:"groupPolicyBootTimeInMs,omitempty"`

	// The impact of device group policy client on login time in milliseconds. Supports: $select, $OrderBy. Read-only.
	GroupPolicyLoginTimeInMs *int64 `json:"groupPolicyLoginTimeInMs,omitempty"`

	// When TRUE, indicates the device boot record is associated with feature updates. When FALSE, indicates the device boot
	// record is not associated with feature updates. Supports: $select, $OrderBy. Read-only.
	IsFeatureUpdate *bool `json:"isFeatureUpdate,omitempty"`

	// When TRUE, indicates the device login is the first login after a reboot. When FALSE, indicates the device login is
	// not the first login after a reboot. Supports: $select, $OrderBy. Read-only.
	IsFirstLogin *bool `json:"isFirstLogin,omitempty"`

	// The user experience analytics device boot record's operating system version. Supports: $select, $OrderBy. Read-only.
	OperatingSystemVersion nullable.Type[string] `json:"operatingSystemVersion,omitempty"`

	// The time for desktop to become responsive during login process in milliseconds. Supports: $select, $OrderBy.
	// Read-only.
	ResponsiveDesktopTimeInMs *int64 `json:"responsiveDesktopTimeInMs,omitempty"`

	// Operating System restart category.
	RestartCategory *UserExperienceAnalyticsOperatingSystemRestartCategory `json:"restartCategory,omitempty"`

	// OS restart fault bucket. The fault bucket is used to find additional information about a system crash. Supports:
	// $select, $OrderBy. Read-only.
	RestartFaultBucket nullable.Type[string] `json:"restartFaultBucket,omitempty"`

	// OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason. Supports:
	// $select, $OrderBy. Read-only.
	RestartStopCode nullable.Type[string] `json:"restartStopCode,omitempty"`

	// The device boot start time. The value cannot be modified and is automatically populated when the device performs a
	// reboot. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For
	// example, midnight UTC on Jan 1, 2022 would look like this: '2022-01-01T00:00:00Z'. Returned by default. Read-only.
	StartTime *string `json:"startTime,omitempty"`

	// The device total boot time in milliseconds. Supports: $select, $OrderBy. Read-only.
	TotalBootTimeInMs *int64 `json:"totalBootTimeInMs,omitempty"`

	// The device total login time in milliseconds. Supports: $select, $OrderBy. Read-only.
	TotalLoginTimeInMs *int64 `json:"totalLoginTimeInMs,omitempty"`

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

func (s UserExperienceAnalyticsDeviceStartupHistory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceStartupHistory{}

func (s UserExperienceAnalyticsDeviceStartupHistory) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceStartupHistory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceStartupHistory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceStartupHistory: %+v", err)
	}

	delete(decoded, "coreBootTimeInMs")
	delete(decoded, "coreLoginTimeInMs")
	delete(decoded, "deviceId")
	delete(decoded, "featureUpdateBootTimeInMs")
	delete(decoded, "groupPolicyBootTimeInMs")
	delete(decoded, "groupPolicyLoginTimeInMs")
	delete(decoded, "isFeatureUpdate")
	delete(decoded, "isFirstLogin")
	delete(decoded, "operatingSystemVersion")
	delete(decoded, "responsiveDesktopTimeInMs")
	delete(decoded, "restartFaultBucket")
	delete(decoded, "restartStopCode")
	delete(decoded, "startTime")
	delete(decoded, "totalBootTimeInMs")
	delete(decoded, "totalLoginTimeInMs")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceStartupHistory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceStartupHistory: %+v", err)
	}

	return encoded, nil
}
