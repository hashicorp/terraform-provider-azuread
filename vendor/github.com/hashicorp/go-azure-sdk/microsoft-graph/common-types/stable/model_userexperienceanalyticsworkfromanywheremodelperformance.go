package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsWorkFromAnywhereModelPerformance{}

type UserExperienceAnalyticsWorkFromAnywhereModelPerformance struct {
	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The model name of the device. Supports: $select, $OrderBy. Read-only.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The devices count for the model. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
	ModelDeviceCount *int64 `json:"modelDeviceCount,omitempty"`

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

func (s UserExperienceAnalyticsWorkFromAnywhereModelPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsWorkFromAnywhereModelPerformance{}

func (s UserExperienceAnalyticsWorkFromAnywhereModelPerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsWorkFromAnywhereModelPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsWorkFromAnywhereModelPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsWorkFromAnywhereModelPerformance: %+v", err)
	}

	delete(decoded, "manufacturer")
	delete(decoded, "model")
	delete(decoded, "modelDeviceCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereModelPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsWorkFromAnywhereModelPerformance: %+v", err)
	}

	return encoded, nil
}
