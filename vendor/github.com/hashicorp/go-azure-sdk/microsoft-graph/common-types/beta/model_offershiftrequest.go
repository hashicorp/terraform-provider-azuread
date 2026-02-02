package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfferShiftRequest interface {
	Entity
	ChangeTrackedEntity
	ScheduleChangeRequest
	OfferShiftRequest() BaseOfferShiftRequestImpl
}

var _ OfferShiftRequest = BaseOfferShiftRequestImpl{}

type BaseOfferShiftRequestImpl struct {
	// The date and time when the recipient approved or declined the request.
	RecipientActionDateTime nullable.Type[string] `json:"recipientActionDateTime,omitempty"`

	// The message sent by the recipient regarding the request.
	RecipientActionMessage nullable.Type[string] `json:"recipientActionMessage,omitempty"`

	// The recipient's user ID.
	RecipientUserId nullable.Type[string] `json:"recipientUserId,omitempty"`

	// The sender's shift ID.
	SenderShiftId nullable.Type[string] `json:"senderShiftId,omitempty"`

	// Fields inherited from ScheduleChangeRequest

	// Indicates who the request is assigned to. Possible values are: sender, recipient, manager, system,
	// unknownFutureValue.
	AssignedTo *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`

	// The date and time when the manager approved or declined the scheduleChangeRequest. The timestamp type represents date
	// and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	ManagerActionDateTime nullable.Type[string] `json:"managerActionDateTime,omitempty"`

	// The message sent by the manager regarding the scheduleChangeRequest. Optional.
	ManagerActionMessage nullable.Type[string] `json:"managerActionMessage,omitempty"`

	// The user ID of the manager who approved or declined the scheduleChangeRequest.
	ManagerUserId nullable.Type[string] `json:"managerUserId,omitempty"`

	// The date and time when the sender sent the scheduleChangeRequest. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	SenderDateTime nullable.Type[string] `json:"senderDateTime,omitempty"`

	// The message sent by the sender of the scheduleChangeRequest. Optional.
	SenderMessage nullable.Type[string] `json:"senderMessage,omitempty"`

	// The user ID of the sender of the scheduleChangeRequest.
	SenderUserId nullable.Type[string] `json:"senderUserId,omitempty"`

	// The state of the scheduleChangeRequest. Possible values are: pending, approved, declined, unknownFutureValue.
	State *ScheduleChangeState `json:"state,omitempty"`

	// Fields inherited from ChangeTrackedEntity

	// Identity of the user who created the entity.
	CreatedBy IdentitySet `json:"createdBy"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the user who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
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

func (s BaseOfferShiftRequestImpl) OfferShiftRequest() BaseOfferShiftRequestImpl {
	return s
}

func (s BaseOfferShiftRequestImpl) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return BaseScheduleChangeRequestImpl{
		AssignedTo:            s.AssignedTo,
		ManagerActionDateTime: s.ManagerActionDateTime,
		ManagerActionMessage:  s.ManagerActionMessage,
		ManagerUserId:         s.ManagerUserId,
		SenderDateTime:        s.SenderDateTime,
		SenderMessage:         s.SenderMessage,
		SenderUserId:          s.SenderUserId,
		State:                 s.State,
		CreatedBy:             s.CreatedBy,
		CreatedDateTime:       s.CreatedDateTime,
		LastModifiedBy:        s.LastModifiedBy,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s BaseOfferShiftRequestImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return BaseChangeTrackedEntityImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseOfferShiftRequestImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OfferShiftRequest = RawOfferShiftRequestImpl{}

// RawOfferShiftRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOfferShiftRequestImpl struct {
	offerShiftRequest BaseOfferShiftRequestImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawOfferShiftRequestImpl) OfferShiftRequest() BaseOfferShiftRequestImpl {
	return s.offerShiftRequest
}

func (s RawOfferShiftRequestImpl) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return s.offerShiftRequest.ScheduleChangeRequest()
}

func (s RawOfferShiftRequestImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return s.offerShiftRequest.ChangeTrackedEntity()
}

func (s RawOfferShiftRequestImpl) Entity() BaseEntityImpl {
	return s.offerShiftRequest.Entity()
}

var _ json.Marshaler = BaseOfferShiftRequestImpl{}

func (s BaseOfferShiftRequestImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOfferShiftRequestImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOfferShiftRequestImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOfferShiftRequestImpl: %+v", err)
	}

	delete(decoded, "recipientActionDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.offerShiftRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOfferShiftRequestImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseOfferShiftRequestImpl{}

func (s *BaseOfferShiftRequestImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		RecipientActionDateTime nullable.Type[string]       `json:"recipientActionDateTime,omitempty"`
		RecipientActionMessage  nullable.Type[string]       `json:"recipientActionMessage,omitempty"`
		RecipientUserId         nullable.Type[string]       `json:"recipientUserId,omitempty"`
		SenderShiftId           nullable.Type[string]       `json:"senderShiftId,omitempty"`
		AssignedTo              *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`
		ManagerActionDateTime   nullable.Type[string]       `json:"managerActionDateTime,omitempty"`
		ManagerActionMessage    nullable.Type[string]       `json:"managerActionMessage,omitempty"`
		ManagerUserId           nullable.Type[string]       `json:"managerUserId,omitempty"`
		SenderDateTime          nullable.Type[string]       `json:"senderDateTime,omitempty"`
		SenderMessage           nullable.Type[string]       `json:"senderMessage,omitempty"`
		SenderUserId            nullable.Type[string]       `json:"senderUserId,omitempty"`
		State                   *ScheduleChangeState        `json:"state,omitempty"`
		CreatedDateTime         nullable.Type[string]       `json:"createdDateTime,omitempty"`
		LastModifiedDateTime    nullable.Type[string]       `json:"lastModifiedDateTime,omitempty"`
		Id                      *string                     `json:"id,omitempty"`
		ODataId                 *string                     `json:"@odata.id,omitempty"`
		ODataType               *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.RecipientActionDateTime = decoded.RecipientActionDateTime
	s.RecipientActionMessage = decoded.RecipientActionMessage
	s.RecipientUserId = decoded.RecipientUserId
	s.SenderShiftId = decoded.SenderShiftId
	s.AssignedTo = decoded.AssignedTo
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ManagerActionDateTime = decoded.ManagerActionDateTime
	s.ManagerActionMessage = decoded.ManagerActionMessage
	s.ManagerUserId = decoded.ManagerUserId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SenderDateTime = decoded.SenderDateTime
	s.SenderMessage = decoded.SenderMessage
	s.SenderUserId = decoded.SenderUserId
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseOfferShiftRequestImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseOfferShiftRequestImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseOfferShiftRequestImpl': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}

func UnmarshalOfferShiftRequestImplementation(input []byte) (OfferShiftRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OfferShiftRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.swapShiftsChangeRequest") {
		var out SwapShiftsChangeRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SwapShiftsChangeRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseOfferShiftRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOfferShiftRequestImpl: %+v", err)
	}

	return RawOfferShiftRequestImpl{
		offerShiftRequest: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
