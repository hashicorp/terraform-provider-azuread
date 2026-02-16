package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails{}

type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails struct {
	// The number of crashes for the app. Valid values -2147483648 to 2147483647
	AppCrashCount *int64 `json:"appCrashCount,omitempty"`

	// The friendly name of the application.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The name of the application.
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// The publisher of the application.
	AppPublisher nullable.Type[string] `json:"appPublisher,omitempty"`

	// The version of the application.
	AppVersion nullable.Type[string] `json:"appVersion,omitempty"`

	// The total number of devices that have reported one or more application crashes for this application and version.
	// Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
	DeviceCountWithCrashes *int64 `json:"deviceCountWithCrashes,omitempty"`

	// When TRUE, indicates the version of application is the latest version for that application that is in use. When
	// FALSE, indicates the version is not the latest version. FALSE by default. Supports: $select, $OrderBy.
	IsLatestUsedVersion *bool `json:"isLatestUsedVersion,omitempty"`

	// When TRUE, indicates the version of application is the most used version for that application. When FALSE, indicates
	// the version is not the most used version. FALSE by default. Supports: $select, $OrderBy. Read-only.
	IsMostUsedVersion *bool `json:"isMostUsedVersion,omitempty"`

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

func (s UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails{}

func (s UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails: %+v", err)
	}

	delete(decoded, "deviceCountWithCrashes")
	delete(decoded, "isMostUsedVersion")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails: %+v", err)
	}

	return encoded, nil
}
