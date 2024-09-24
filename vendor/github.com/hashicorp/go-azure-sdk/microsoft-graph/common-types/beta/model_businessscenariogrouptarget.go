package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BusinessScenarioTaskTargetBase = BusinessScenarioGroupTarget{}

type BusinessScenarioGroupTarget struct {
	// The unique identifier for the group.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Fields inherited from BusinessScenarioTaskTargetBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TaskTargetKind *PlannerTaskTargetKind `json:"taskTargetKind,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BusinessScenarioGroupTarget) BusinessScenarioTaskTargetBase() BaseBusinessScenarioTaskTargetBaseImpl {
	return BaseBusinessScenarioTaskTargetBaseImpl{
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		TaskTargetKind: s.TaskTargetKind,
	}
}

var _ json.Marshaler = BusinessScenarioGroupTarget{}

func (s BusinessScenarioGroupTarget) MarshalJSON() ([]byte, error) {
	type wrapper BusinessScenarioGroupTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BusinessScenarioGroupTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BusinessScenarioGroupTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.businessScenarioGroupTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BusinessScenarioGroupTarget: %+v", err)
	}

	return encoded, nil
}
