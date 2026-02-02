package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceActionResult = RotateBitLockerKeysDeviceActionResult{}

type RotateBitLockerKeysDeviceActionResult struct {
	// RotateBitLockerKeys action error code
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// Fields inherited from DeviceActionResult

	// Action name
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	// State of the action on the device
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

func (s RotateBitLockerKeysDeviceActionResult) DeviceActionResult() BaseDeviceActionResultImpl {
	return BaseDeviceActionResultImpl{
		ActionName:          s.ActionName,
		ActionState:         s.ActionState,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartDateTime:       s.StartDateTime,
	}
}

var _ json.Marshaler = RotateBitLockerKeysDeviceActionResult{}

func (s RotateBitLockerKeysDeviceActionResult) MarshalJSON() ([]byte, error) {
	type wrapper RotateBitLockerKeysDeviceActionResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RotateBitLockerKeysDeviceActionResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RotateBitLockerKeysDeviceActionResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.rotateBitLockerKeysDeviceActionResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RotateBitLockerKeysDeviceActionResult: %+v", err)
	}

	return encoded, nil
}
