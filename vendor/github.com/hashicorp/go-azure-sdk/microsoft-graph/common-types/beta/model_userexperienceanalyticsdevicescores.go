package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceScores{}

type UserExperienceAnalyticsDeviceScores struct {
	// The name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy.
	// Read-only.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The model name of the device. Supports: $select, $OrderBy. Read-only.
	Model nullable.Type[string] `json:"model,omitempty"`

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

func (s UserExperienceAnalyticsDeviceScores) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceScores{}

func (s UserExperienceAnalyticsDeviceScores) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceScores
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceScores: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceScores: %+v", err)
	}

	delete(decoded, "deviceName")
	delete(decoded, "manufacturer")
	delete(decoded, "model")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceScores"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceScores: %+v", err)
	}

	return encoded, nil
}
