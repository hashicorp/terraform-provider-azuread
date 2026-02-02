package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsImpactingProcess{}

type UserExperienceAnalyticsImpactingProcess struct {
	// The category of impacting process.
	Category nullable.Type[string] `json:"category,omitempty"`

	// The description of process.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The unique identifier of the impacted device.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The process name.
	ProcessName nullable.Type[string] `json:"processName,omitempty"`

	// The publisher of the process.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

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

func (s UserExperienceAnalyticsImpactingProcess) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsImpactingProcess{}

func (s UserExperienceAnalyticsImpactingProcess) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsImpactingProcess
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsImpactingProcess: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsImpactingProcess: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsImpactingProcess"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsImpactingProcess: %+v", err)
	}

	return encoded, nil
}
