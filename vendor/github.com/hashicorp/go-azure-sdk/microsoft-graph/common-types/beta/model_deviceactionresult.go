package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceActionResult interface {
	DeviceActionResult() BaseDeviceActionResultImpl
}

var _ DeviceActionResult = BaseDeviceActionResultImpl{}

type BaseDeviceActionResultImpl struct {
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

func (s BaseDeviceActionResultImpl) DeviceActionResult() BaseDeviceActionResultImpl {
	return s
}

var _ DeviceActionResult = RawDeviceActionResultImpl{}

// RawDeviceActionResultImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceActionResultImpl struct {
	deviceActionResult BaseDeviceActionResultImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawDeviceActionResultImpl) DeviceActionResult() BaseDeviceActionResultImpl {
	return s.deviceActionResult
}

func UnmarshalDeviceActionResultImplementation(input []byte) (DeviceActionResult, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceActionResult into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.activateDeviceEsimActionResult") {
		var out ActivateDeviceEsimActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivateDeviceEsimActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.changeAssignmentsActionResult") {
		var out ChangeAssignmentsActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChangeAssignmentsActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.configurationManagerActionResult") {
		var out ConfigurationManagerActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConfigurationManagerActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deleteUserFromSharedAppleDeviceActionResult") {
		var out DeleteUserFromSharedAppleDeviceActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeleteUserFromSharedAppleDeviceActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.locateDeviceActionResult") {
		var out LocateDeviceActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocateDeviceActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.remoteLockActionResult") {
		var out RemoteLockActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RemoteLockActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resetPasscodeActionResult") {
		var out ResetPasscodeActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResetPasscodeActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.revokeAppleVppLicensesActionResult") {
		var out RevokeAppleVppLicensesActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RevokeAppleVppLicensesActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.rotateBitLockerKeysDeviceActionResult") {
		var out RotateBitLockerKeysDeviceActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RotateBitLockerKeysDeviceActionResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDefenderScanActionResult") {
		var out WindowsDefenderScanActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDefenderScanActionResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceActionResultImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceActionResultImpl: %+v", err)
	}

	return RawDeviceActionResultImpl{
		deviceActionResult: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
