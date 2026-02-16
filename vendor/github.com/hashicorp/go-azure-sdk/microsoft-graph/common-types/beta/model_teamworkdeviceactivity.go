package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamworkDeviceActivity{}

type TeamworkDeviceActivity struct {
	// The active peripheral devices attached to the device.
	ActivePeripherals *TeamworkActivePeripherals `json:"activePeripherals,omitempty"`

	// Identity of the user who created the device activity document.
	CreatedBy IdentitySet `json:"createdBy"`

	// The UTC date and time when the device activity document was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the user who last modified the device activity details.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The UTC date and time when the device activity detail was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s TeamworkDeviceActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamworkDeviceActivity{}

func (s TeamworkDeviceActivity) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkDeviceActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkDeviceActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkDeviceActivity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkDeviceActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkDeviceActivity: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamworkDeviceActivity{}

func (s *TeamworkDeviceActivity) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActivePeripherals    *TeamworkActivePeripherals `json:"activePeripherals,omitempty"`
		CreatedDateTime      nullable.Type[string]      `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string]      `json:"lastModifiedDateTime,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActivePeripherals = decoded.ActivePeripherals
	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamworkDeviceActivity into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamworkDeviceActivity': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TeamworkDeviceActivity': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
