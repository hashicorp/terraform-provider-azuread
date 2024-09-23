package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceLifecycleManagementSettings{}

type IdentityGovernanceLifecycleManagementSettings struct {
	EmailSettings *EmailSettings `json:"emailSettings,omitempty"`

	// The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval
	// has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
	WorkflowScheduleIntervalInHours *int64 `json:"workflowScheduleIntervalInHours,omitempty"`

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

func (s IdentityGovernanceLifecycleManagementSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceLifecycleManagementSettings{}

func (s IdentityGovernanceLifecycleManagementSettings) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceLifecycleManagementSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceLifecycleManagementSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceLifecycleManagementSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.lifecycleManagementSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceLifecycleManagementSettings: %+v", err)
	}

	return encoded, nil
}
