package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationMembership interface {
	DeviceRegistrationMembership() BaseDeviceRegistrationMembershipImpl
}

var _ DeviceRegistrationMembership = BaseDeviceRegistrationMembershipImpl{}

type BaseDeviceRegistrationMembershipImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceRegistrationMembershipImpl) DeviceRegistrationMembership() BaseDeviceRegistrationMembershipImpl {
	return s
}

var _ DeviceRegistrationMembership = RawDeviceRegistrationMembershipImpl{}

// RawDeviceRegistrationMembershipImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceRegistrationMembershipImpl struct {
	deviceRegistrationMembership BaseDeviceRegistrationMembershipImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawDeviceRegistrationMembershipImpl) DeviceRegistrationMembership() BaseDeviceRegistrationMembershipImpl {
	return s.deviceRegistrationMembership
}

func UnmarshalDeviceRegistrationMembershipImplementation(input []byte) (DeviceRegistrationMembership, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceRegistrationMembership into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allDeviceRegistrationMembership") {
		var out AllDeviceRegistrationMembership
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllDeviceRegistrationMembership: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enumeratedDeviceRegistrationMembership") {
		var out EnumeratedDeviceRegistrationMembership
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnumeratedDeviceRegistrationMembership: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noDeviceRegistrationMembership") {
		var out NoDeviceRegistrationMembership
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoDeviceRegistrationMembership: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceRegistrationMembershipImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceRegistrationMembershipImpl: %+v", err)
	}

	return RawDeviceRegistrationMembershipImpl{
		deviceRegistrationMembership: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
