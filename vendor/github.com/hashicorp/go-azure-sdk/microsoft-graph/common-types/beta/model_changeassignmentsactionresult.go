package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceActionResult = ChangeAssignmentsActionResult{}

type ChangeAssignmentsActionResult struct {
	// Indicates the list of applications or configuration to report live results during their changeAssignments action
	// execution process. The result for each individual application or configuration can contain whether it's being removed
	// or restored, what's the current status with potential message or error code, and when any changes happen on it.
	// Read-Only. This collection can contain a maximum of 30 elements.
	DeviceAssignmentItems *[]DeviceAssignmentItem `json:"deviceAssignmentItems,omitempty"`

	// Fields inherited from DeviceActionResult

	// Action name
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	ActionState *ActionState `json:"actionState,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time the action was initiated
	StartDateTime *string `json:"startDateTime,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChangeAssignmentsActionResult) DeviceActionResult() BaseDeviceActionResultImpl {
	return BaseDeviceActionResultImpl{
		ActionName:          s.ActionName,
		ActionState:         s.ActionState,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartDateTime:       s.StartDateTime,
	}
}

var _ json.Marshaler = ChangeAssignmentsActionResult{}

func (s ChangeAssignmentsActionResult) MarshalJSON() ([]byte, error) {
	type wrapper ChangeAssignmentsActionResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChangeAssignmentsActionResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChangeAssignmentsActionResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.changeAssignmentsActionResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChangeAssignmentsActionResult: %+v", err)
	}

	return encoded, nil
}
