package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipantInfo interface {
	MeetingParticipantInfo() BaseMeetingParticipantInfoImpl
}

var _ MeetingParticipantInfo = BaseMeetingParticipantInfoImpl{}

type BaseMeetingParticipantInfoImpl struct {
	// Identity information of the participant. Only the user property is used for onlineMeeting participants.
	Identity IdentitySet `json:"identity"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the participant's role in the meeting.
	Role *OnlineMeetingRole `json:"role,omitempty"`

	// User principal name of the participant.
	Upn nullable.Type[string] `json:"upn,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMeetingParticipantInfoImpl) MeetingParticipantInfo() BaseMeetingParticipantInfoImpl {
	return s
}

var _ MeetingParticipantInfo = RawMeetingParticipantInfoImpl{}

// RawMeetingParticipantInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMeetingParticipantInfoImpl struct {
	meetingParticipantInfo BaseMeetingParticipantInfoImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawMeetingParticipantInfoImpl) MeetingParticipantInfo() BaseMeetingParticipantInfoImpl {
	return s.meetingParticipantInfo
}

var _ json.Unmarshaler = &BaseMeetingParticipantInfoImpl{}

func (s *BaseMeetingParticipantInfoImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
		Role      *OnlineMeetingRole    `json:"role,omitempty"`
		Upn       nullable.Type[string] `json:"upn,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Role = decoded.Role
	s.Upn = decoded.Upn

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseMeetingParticipantInfoImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'BaseMeetingParticipantInfoImpl': %+v", err)
		}
		s.Identity = impl
	}

	return nil
}

func UnmarshalMeetingParticipantInfoImplementation(input []byte) (MeetingParticipantInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingParticipantInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventPresenterInfo") {
		var out VirtualEventPresenterInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventPresenterInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseMeetingParticipantInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMeetingParticipantInfoImpl: %+v", err)
	}

	return RawMeetingParticipantInfoImpl{
		meetingParticipantInfo: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
