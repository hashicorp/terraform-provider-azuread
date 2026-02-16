package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDevicePerformance{}

type UserExperienceAnalyticsDevicePerformance struct {
	// Number of Blue Screens in the last 30 days. Valid values 0 to 9999999
	BlueScreenCount *int64 `json:"blueScreenCount,omitempty"`

	// The user experience analytics device boot score.
	BootScore *int64 `json:"bootScore,omitempty"`

	// The user experience analytics device core boot time in milliseconds.
	CoreBootTimeInMs *int64 `json:"coreBootTimeInMs,omitempty"`

	// The user experience analytics device core login time in milliseconds.
	CoreLoginTimeInMs *int64 `json:"coreLoginTimeInMs,omitempty"`

	// User experience analytics summarized device count.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The user experience analytics device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	DiskType *DiskType `json:"diskType,omitempty"`

	// The user experience analytics device group policy boot time in milliseconds.
	GroupPolicyBootTimeInMs *int64 `json:"groupPolicyBootTimeInMs,omitempty"`

	// The user experience analytics device group policy login time in milliseconds.
	GroupPolicyLoginTimeInMs *int64 `json:"groupPolicyLoginTimeInMs,omitempty"`

	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// The user experience analytics device login score.
	LoginScore *int64 `json:"loginScore,omitempty"`

	// The user experience analytics device manufacturer.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The user experience analytics device model.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The user experience analytics device Operating System version.
	OperatingSystemVersion nullable.Type[string] `json:"operatingSystemVersion,omitempty"`

	// The user experience analytics responsive desktop time in milliseconds.
	ResponsiveDesktopTimeInMs *int64 `json:"responsiveDesktopTimeInMs,omitempty"`

	// Number of Restarts in the last 30 days. Valid values 0 to 9999999
	RestartCount *int64 `json:"restartCount,omitempty"`

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

func (s UserExperienceAnalyticsDevicePerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDevicePerformance{}

func (s UserExperienceAnalyticsDevicePerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDevicePerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDevicePerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDevicePerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDevicePerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDevicePerformance: %+v", err)
	}

	return encoded, nil
}
