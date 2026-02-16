package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConstraint interface {
	DeviceManagementConstraint() BaseDeviceManagementConstraintImpl
}

var _ DeviceManagementConstraint = BaseDeviceManagementConstraintImpl{}

type BaseDeviceManagementConstraintImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConstraintImpl) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return s
}

var _ DeviceManagementConstraint = RawDeviceManagementConstraintImpl{}

// RawDeviceManagementConstraintImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConstraintImpl struct {
	deviceManagementConstraint BaseDeviceManagementConstraintImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawDeviceManagementConstraintImpl) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return s.deviceManagementConstraint
}

func UnmarshalDeviceManagementConstraintImplementation(input []byte) (DeviceManagementConstraint, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConstraint into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementEnumConstraint") {
		var out DeviceManagementEnumConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementEnumConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentSettingSecretConstraint") {
		var out DeviceManagementIntentSettingSecretConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentSettingSecretConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingAbstractImplementationConstraint") {
		var out DeviceManagementSettingAbstractImplementationConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingAbstractImplementationConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingAppConstraint") {
		var out DeviceManagementSettingAppConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingAppConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingBooleanConstraint") {
		var out DeviceManagementSettingBooleanConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingBooleanConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingCollectionConstraint") {
		var out DeviceManagementSettingCollectionConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingCollectionConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingEnrollmentTypeConstraint") {
		var out DeviceManagementSettingEnrollmentTypeConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingEnrollmentTypeConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingFileConstraint") {
		var out DeviceManagementSettingFileConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingFileConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingIntegerConstraint") {
		var out DeviceManagementSettingIntegerConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingIntegerConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingProfileConstraint") {
		var out DeviceManagementSettingProfileConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingProfileConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingRegexConstraint") {
		var out DeviceManagementSettingRegexConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingRegexConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingRequiredConstraint") {
		var out DeviceManagementSettingRequiredConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingRequiredConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingSddlConstraint") {
		var out DeviceManagementSettingSddlConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingSddlConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingStringLengthConstraint") {
		var out DeviceManagementSettingStringLengthConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingStringLengthConstraint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementSettingXmlConstraint") {
		var out DeviceManagementSettingXmlConstraint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementSettingXmlConstraint: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConstraintImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConstraintImpl: %+v", err)
	}

	return RawDeviceManagementConstraintImpl{
		deviceManagementConstraint: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
