package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceComplianceScheduledActionForRule{}

type DeviceComplianceScheduledActionForRule struct {
	// Name of the rule which this scheduled action applies to. Currently scheduled actions are created per policy instead
	// of per rule, thus RuleName is always set to default value PasswordRequired.
	RuleName nullable.Type[string] `json:"ruleName,omitempty"`

	// The list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one
	// block scheduled action.
	ScheduledActionConfigurations *[]DeviceComplianceActionItem `json:"scheduledActionConfigurations,omitempty"`

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

func (s DeviceComplianceScheduledActionForRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceComplianceScheduledActionForRule{}

func (s DeviceComplianceScheduledActionForRule) MarshalJSON() ([]byte, error) {
	type wrapper DeviceComplianceScheduledActionForRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceComplianceScheduledActionForRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceComplianceScheduledActionForRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceComplianceScheduledActionForRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceComplianceScheduledActionForRule: %+v", err)
	}

	return encoded, nil
}
