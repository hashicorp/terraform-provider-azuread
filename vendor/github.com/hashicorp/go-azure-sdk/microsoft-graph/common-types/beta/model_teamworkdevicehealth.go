package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamworkDeviceHealth{}

type TeamworkDeviceHealth struct {
	Connection *TeamworkConnection `json:"connection,omitempty"`

	// Identity of the user who created the device health document.
	CreatedBy IdentitySet `json:"createdBy"`

	// The UTC date and time when the device health document was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Health details about the device hardware.
	HardwareHealth *TeamworkHardwareHealth `json:"hardwareHealth,omitempty"`

	// Identity of the user who last modified the device health details.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The UTC date and time when the device health detail was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The login status of Microsoft Teams, Skype for Business, and Exchange.
	LoginStatus *TeamworkLoginStatus `json:"loginStatus,omitempty"`

	// Health details about all peripherals (for example, speaker and microphone) attached to a device.
	PeripheralsHealth *TeamworkPeripheralsHealth `json:"peripheralsHealth,omitempty"`

	// Software updates available for the device.
	SoftwareUpdateHealth *TeamworkSoftwareUpdateHealth `json:"softwareUpdateHealth,omitempty"`

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

func (s TeamworkDeviceHealth) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamworkDeviceHealth{}

func (s TeamworkDeviceHealth) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkDeviceHealth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkDeviceHealth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkDeviceHealth: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkDeviceHealth"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkDeviceHealth: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamworkDeviceHealth{}

func (s *TeamworkDeviceHealth) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Connection           *TeamworkConnection           `json:"connection,omitempty"`
		CreatedDateTime      nullable.Type[string]         `json:"createdDateTime,omitempty"`
		HardwareHealth       *TeamworkHardwareHealth       `json:"hardwareHealth,omitempty"`
		LastModifiedDateTime nullable.Type[string]         `json:"lastModifiedDateTime,omitempty"`
		LoginStatus          *TeamworkLoginStatus          `json:"loginStatus,omitempty"`
		PeripheralsHealth    *TeamworkPeripheralsHealth    `json:"peripheralsHealth,omitempty"`
		SoftwareUpdateHealth *TeamworkSoftwareUpdateHealth `json:"softwareUpdateHealth,omitempty"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Connection = decoded.Connection
	s.CreatedDateTime = decoded.CreatedDateTime
	s.HardwareHealth = decoded.HardwareHealth
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LoginStatus = decoded.LoginStatus
	s.PeripheralsHealth = decoded.PeripheralsHealth
	s.SoftwareUpdateHealth = decoded.SoftwareUpdateHealth
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamworkDeviceHealth into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamworkDeviceHealth': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TeamworkDeviceHealth': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
