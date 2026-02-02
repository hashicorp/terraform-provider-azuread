package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementComplianceScheduledActionForRule{}

type DeviceManagementComplianceScheduledActionForRule struct {
	// Name of the rule which this scheduled action applies to.
	RuleName nullable.Type[string] `json:"ruleName,omitempty"`

	// The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100
	// elements.
	ScheduledActionConfigurations *[]DeviceManagementComplianceActionItem `json:"scheduledActionConfigurations,omitempty"`

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

func (s DeviceManagementComplianceScheduledActionForRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementComplianceScheduledActionForRule{}

func (s DeviceManagementComplianceScheduledActionForRule) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementComplianceScheduledActionForRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementComplianceScheduledActionForRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementComplianceScheduledActionForRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementComplianceScheduledActionForRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementComplianceScheduledActionForRule: %+v", err)
	}

	return encoded, nil
}
