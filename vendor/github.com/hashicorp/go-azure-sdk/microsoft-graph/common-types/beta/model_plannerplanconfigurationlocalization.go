package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerPlanConfigurationLocalization{}

type PlannerPlanConfigurationLocalization struct {
	// Localized names for configured buckets in the plan configuration.
	Buckets *[]PlannerPlanConfigurationBucketLocalization `json:"buckets,omitempty"`

	// The language code associated with the localized names in this object.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// Localized title of the plan.
	PlanTitle nullable.Type[string] `json:"planTitle,omitempty"`

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

func (s PlannerPlanConfigurationLocalization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerPlanConfigurationLocalization{}

func (s PlannerPlanConfigurationLocalization) MarshalJSON() ([]byte, error) {
	type wrapper PlannerPlanConfigurationLocalization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerPlanConfigurationLocalization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlanConfigurationLocalization: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerPlanConfigurationLocalization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerPlanConfigurationLocalization: %+v", err)
	}

	return encoded, nil
}
