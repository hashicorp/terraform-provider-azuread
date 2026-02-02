package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBaseline{}

type UserExperienceAnalyticsBaseline struct {
	// The scores and insights for the application health metrics.
	AppHealthMetrics *UserExperienceAnalyticsCategory `json:"appHealthMetrics,omitempty"`

	// The scores and insights for the battery health metrics.
	BatteryHealthMetrics *UserExperienceAnalyticsCategory `json:"batteryHealthMetrics,omitempty"`

	// The scores and insights for the best practices metrics.
	BestPracticesMetrics *UserExperienceAnalyticsCategory `json:"bestPracticesMetrics,omitempty"`

	// The date the custom baseline was created. The value cannot be modified and is automatically populated when the
	// baseline is created. The Timestamp type represents date and time information using ISO 8601 format and is always in
	// UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The scores and insights for the device boot performance metrics.
	DeviceBootPerformanceMetrics *UserExperienceAnalyticsCategory `json:"deviceBootPerformanceMetrics,omitempty"`

	// The name of the baseline.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// When TRUE, indicates the current baseline is the commercial median baseline. When FALSE, indicates it is a custom
	// baseline. FALSE by default.
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`

	// The scores and insights for the reboot analytics metrics.
	RebootAnalyticsMetrics *UserExperienceAnalyticsCategory `json:"rebootAnalyticsMetrics,omitempty"`

	// The scores and insights for the resource performance metrics.
	ResourcePerformanceMetrics *UserExperienceAnalyticsCategory `json:"resourcePerformanceMetrics,omitempty"`

	// The scores and insights for the work from anywhere metrics.
	WorkFromAnywhereMetrics *UserExperienceAnalyticsCategory `json:"workFromAnywhereMetrics,omitempty"`

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

func (s UserExperienceAnalyticsBaseline) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBaseline{}

func (s UserExperienceAnalyticsBaseline) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBaseline
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBaseline: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBaseline: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBaseline"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBaseline: %+v", err)
	}

	return encoded, nil
}
