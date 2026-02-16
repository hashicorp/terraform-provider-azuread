package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodDevice interface {
	Entity
	AuthenticationMethodDevice() BaseAuthenticationMethodDeviceImpl
}

var _ AuthenticationMethodDevice = BaseAuthenticationMethodDeviceImpl{}

type BaseAuthenticationMethodDeviceImpl struct {
	// Optional name given to the hardware OATH device.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Exposes the hardware OATH method in the directory.
	HardwareOathDevices *[]HardwareOathTokenAuthenticationMethodDevice `json:"hardwareOathDevices,omitempty"`

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

func (s BaseAuthenticationMethodDeviceImpl) AuthenticationMethodDevice() BaseAuthenticationMethodDeviceImpl {
	return s
}

func (s BaseAuthenticationMethodDeviceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationMethodDevice = RawAuthenticationMethodDeviceImpl{}

// RawAuthenticationMethodDeviceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationMethodDeviceImpl struct {
	authenticationMethodDevice BaseAuthenticationMethodDeviceImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawAuthenticationMethodDeviceImpl) AuthenticationMethodDevice() BaseAuthenticationMethodDeviceImpl {
	return s.authenticationMethodDevice
}

func (s RawAuthenticationMethodDeviceImpl) Entity() BaseEntityImpl {
	return s.authenticationMethodDevice.Entity()
}

var _ json.Marshaler = BaseAuthenticationMethodDeviceImpl{}

func (s BaseAuthenticationMethodDeviceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationMethodDeviceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationMethodDeviceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationMethodDeviceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationMethodDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationMethodDeviceImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationMethodDeviceImplementation(input []byte) (AuthenticationMethodDevice, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationMethodDevice into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareOathTokenAuthenticationMethodDevice") {
		var out HardwareOathTokenAuthenticationMethodDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareOathTokenAuthenticationMethodDevice: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationMethodDeviceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationMethodDeviceImpl: %+v", err)
	}

	return RawAuthenticationMethodDeviceImpl{
		authenticationMethodDevice: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
