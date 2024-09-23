package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentTarget interface {
	DeviceAndAppManagementAssignmentTarget() BaseDeviceAndAppManagementAssignmentTargetImpl
}

var _ DeviceAndAppManagementAssignmentTarget = BaseDeviceAndAppManagementAssignmentTargetImpl{}

type BaseDeviceAndAppManagementAssignmentTargetImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceAndAppManagementAssignmentTargetImpl) DeviceAndAppManagementAssignmentTarget() BaseDeviceAndAppManagementAssignmentTargetImpl {
	return s
}

var _ DeviceAndAppManagementAssignmentTarget = RawDeviceAndAppManagementAssignmentTargetImpl{}

// RawDeviceAndAppManagementAssignmentTargetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceAndAppManagementAssignmentTargetImpl struct {
	deviceAndAppManagementAssignmentTarget BaseDeviceAndAppManagementAssignmentTargetImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawDeviceAndAppManagementAssignmentTargetImpl) DeviceAndAppManagementAssignmentTarget() BaseDeviceAndAppManagementAssignmentTargetImpl {
	return s.deviceAndAppManagementAssignmentTarget
}

func UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(input []byte) (DeviceAndAppManagementAssignmentTarget, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementAssignmentTarget into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allDevicesAssignmentTarget") {
		var out AllDevicesAssignmentTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllDevicesAssignmentTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.allLicensedUsersAssignmentTarget") {
		var out AllLicensedUsersAssignmentTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllLicensedUsersAssignmentTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.configurationManagerCollectionAssignmentTarget") {
		var out ConfigurationManagerCollectionAssignmentTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConfigurationManagerCollectionAssignmentTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupAssignmentTarget") {
		var out GroupAssignmentTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupAssignmentTarget: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceAndAppManagementAssignmentTargetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceAndAppManagementAssignmentTargetImpl: %+v", err)
	}

	return RawDeviceAndAppManagementAssignmentTargetImpl{
		deviceAndAppManagementAssignmentTarget: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
