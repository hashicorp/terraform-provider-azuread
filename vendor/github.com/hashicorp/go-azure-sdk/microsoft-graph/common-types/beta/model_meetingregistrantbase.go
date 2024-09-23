package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrantBase interface {
	Entity
	MeetingRegistrantBase() BaseMeetingRegistrantBaseImpl
}

var _ MeetingRegistrantBase = BaseMeetingRegistrantBaseImpl{}

type BaseMeetingRegistrantBaseImpl struct {
	// A unique web URL for the registrant to join the meeting. Read-only.
	JoinWebUrl nullable.Type[string] `json:"joinWebUrl,omitempty"`

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

func (s BaseMeetingRegistrantBaseImpl) MeetingRegistrantBase() BaseMeetingRegistrantBaseImpl {
	return s
}

func (s BaseMeetingRegistrantBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MeetingRegistrantBase = RawMeetingRegistrantBaseImpl{}

// RawMeetingRegistrantBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMeetingRegistrantBaseImpl struct {
	meetingRegistrantBase BaseMeetingRegistrantBaseImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawMeetingRegistrantBaseImpl) MeetingRegistrantBase() BaseMeetingRegistrantBaseImpl {
	return s.meetingRegistrantBase
}

func (s RawMeetingRegistrantBaseImpl) Entity() BaseEntityImpl {
	return s.meetingRegistrantBase.Entity()
}

var _ json.Marshaler = BaseMeetingRegistrantBaseImpl{}

func (s BaseMeetingRegistrantBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMeetingRegistrantBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMeetingRegistrantBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMeetingRegistrantBaseImpl: %+v", err)
	}

	delete(decoded, "joinWebUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.meetingRegistrantBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMeetingRegistrantBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMeetingRegistrantBaseImplementation(input []byte) (MeetingRegistrantBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingRegistrantBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.externalMeetingRegistrant") {
		var out ExternalMeetingRegistrant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalMeetingRegistrant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingRegistrant") {
		var out MeetingRegistrant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingRegistrant: %+v", err)
		}
		return out, nil
	}

	var parent BaseMeetingRegistrantBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMeetingRegistrantBaseImpl: %+v", err)
	}

	return RawMeetingRegistrantBaseImpl{
		meetingRegistrantBase: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
