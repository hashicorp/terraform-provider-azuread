package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsParticipantBase interface {
	Entity
	CallRecordsParticipantBase() BaseCallRecordsParticipantBaseImpl
}

var _ CallRecordsParticipantBase = BaseCallRecordsParticipantBaseImpl{}

type BaseCallRecordsParticipantBaseImpl struct {
	// List of administrativeUnitInfo of the call participant.
	AdministrativeUnitInfos *[]CallRecordsAdministrativeUnitInfo `json:"administrativeUnitInfos,omitempty"`

	// The identity of the call participant.
	Identity *CommunicationsIdentitySet `json:"identity,omitempty"`

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

func (s BaseCallRecordsParticipantBaseImpl) CallRecordsParticipantBase() BaseCallRecordsParticipantBaseImpl {
	return s
}

func (s BaseCallRecordsParticipantBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CallRecordsParticipantBase = RawCallRecordsParticipantBaseImpl{}

// RawCallRecordsParticipantBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCallRecordsParticipantBaseImpl struct {
	callRecordsParticipantBase BaseCallRecordsParticipantBaseImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawCallRecordsParticipantBaseImpl) CallRecordsParticipantBase() BaseCallRecordsParticipantBaseImpl {
	return s.callRecordsParticipantBase
}

func (s RawCallRecordsParticipantBaseImpl) Entity() BaseEntityImpl {
	return s.callRecordsParticipantBase.Entity()
}

var _ json.Marshaler = BaseCallRecordsParticipantBaseImpl{}

func (s BaseCallRecordsParticipantBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCallRecordsParticipantBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCallRecordsParticipantBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCallRecordsParticipantBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.participantBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCallRecordsParticipantBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalCallRecordsParticipantBaseImplementation(input []byte) (CallRecordsParticipantBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsParticipantBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.organizer") {
		var out CallRecordsOrganizer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsOrganizer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.participant") {
		var out CallRecordsParticipant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsParticipant: %+v", err)
		}
		return out, nil
	}

	var parent BaseCallRecordsParticipantBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCallRecordsParticipantBaseImpl: %+v", err)
	}

	return RawCallRecordsParticipantBaseImpl{
		callRecordsParticipantBase: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
