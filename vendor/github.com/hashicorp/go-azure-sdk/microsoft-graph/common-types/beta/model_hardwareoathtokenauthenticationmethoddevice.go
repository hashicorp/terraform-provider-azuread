package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodDevice = HardwareOathTokenAuthenticationMethodDevice{}

type HardwareOathTokenAuthenticationMethodDevice struct {
	// Assign the hardware OATH token to a user.
	AssignTo *User `json:"assignTo,omitempty"`

	// User the token is assigned to. Nullable. Supports $filter (eq).
	AssignedTo *Identity `json:"assignedTo"`

	// Hash function of the hardrware token. The possible values are: hmacsha1 or hmacsha256. Default value is: hmacsha1.
	// Supports $filter (eq).
	HashFunction *HardwareOathTokenHashFunction `json:"hashFunction,omitempty"`

	// Manufacturer name of the hardware token. Supports $filter (eq).
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Model name of the hardware token. Supports $filter (eq).
	Model *string `json:"model,omitempty"`

	// Secret key of the specific hardware token, provided by the vendor.
	SecretKey nullable.Type[string] `json:"secretKey,omitempty"`

	// Serial number of the specific hardware token, often found on the back of the device. Supports $select and $filter
	// (eq).
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Status of the hardware OATH token.The possible values are: available, assigned, activated, failedActivation. Supports
	// $filter(eq).
	Status *HardwareOathTokenStatus `json:"status,omitempty"`

	// Refresh interval of the 6-digit verification code, in seconds. The possible values are: 30 or 60. Supports $filter
	// (eq).
	TimeIntervalInSeconds *int64 `json:"timeIntervalInSeconds,omitempty"`

	// Fields inherited from AuthenticationMethodDevice

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

func (s HardwareOathTokenAuthenticationMethodDevice) AuthenticationMethodDevice() BaseAuthenticationMethodDeviceImpl {
	return BaseAuthenticationMethodDeviceImpl{
		DisplayName:         s.DisplayName,
		HardwareOathDevices: s.HardwareOathDevices,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s HardwareOathTokenAuthenticationMethodDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwareOathTokenAuthenticationMethodDevice{}

func (s HardwareOathTokenAuthenticationMethodDevice) MarshalJSON() ([]byte, error) {
	type wrapper HardwareOathTokenAuthenticationMethodDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwareOathTokenAuthenticationMethodDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwareOathTokenAuthenticationMethodDevice: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwareOathTokenAuthenticationMethodDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwareOathTokenAuthenticationMethodDevice: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &HardwareOathTokenAuthenticationMethodDevice{}

func (s *HardwareOathTokenAuthenticationMethodDevice) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignTo              *User                                          `json:"assignTo,omitempty"`
		HashFunction          *HardwareOathTokenHashFunction                 `json:"hashFunction,omitempty"`
		Manufacturer          *string                                        `json:"manufacturer,omitempty"`
		Model                 *string                                        `json:"model,omitempty"`
		SecretKey             nullable.Type[string]                          `json:"secretKey,omitempty"`
		SerialNumber          *string                                        `json:"serialNumber,omitempty"`
		Status                *HardwareOathTokenStatus                       `json:"status,omitempty"`
		TimeIntervalInSeconds *int64                                         `json:"timeIntervalInSeconds,omitempty"`
		DisplayName           nullable.Type[string]                          `json:"displayName,omitempty"`
		HardwareOathDevices   *[]HardwareOathTokenAuthenticationMethodDevice `json:"hardwareOathDevices,omitempty"`
		Id                    *string                                        `json:"id,omitempty"`
		ODataId               *string                                        `json:"@odata.id,omitempty"`
		ODataType             *string                                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignTo = decoded.AssignTo
	s.HashFunction = decoded.HashFunction
	s.Manufacturer = decoded.Manufacturer
	s.Model = decoded.Model
	s.SecretKey = decoded.SecretKey
	s.SerialNumber = decoded.SerialNumber
	s.Status = decoded.Status
	s.TimeIntervalInSeconds = decoded.TimeIntervalInSeconds
	s.DisplayName = decoded.DisplayName
	s.HardwareOathDevices = decoded.HardwareOathDevices
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling HardwareOathTokenAuthenticationMethodDevice into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignedTo"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssignedTo' for 'HardwareOathTokenAuthenticationMethodDevice': %+v", err)
		}
		s.AssignedTo = &impl
	}

	return nil
}
