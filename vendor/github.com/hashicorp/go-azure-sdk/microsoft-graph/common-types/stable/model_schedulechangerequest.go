package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleChangeRequest interface {
	Entity
	ChangeTrackedEntity
	ScheduleChangeRequest() BaseScheduleChangeRequestImpl
}

var _ ScheduleChangeRequest = BaseScheduleChangeRequestImpl{}

type BaseScheduleChangeRequestImpl struct {
	AssignedTo            *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`
	ManagerActionDateTime nullable.Type[string]       `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage  nullable.Type[string]       `json:"managerActionMessage,omitempty"`
	ManagerUserId         nullable.Type[string]       `json:"managerUserId,omitempty"`
	SenderDateTime        nullable.Type[string]       `json:"senderDateTime,omitempty"`
	SenderMessage         nullable.Type[string]       `json:"senderMessage,omitempty"`
	SenderUserId          nullable.Type[string]       `json:"senderUserId,omitempty"`
	State                 *ScheduleChangeState        `json:"state,omitempty"`

	// Fields inherited from ChangeTrackedEntity

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identity of the person who last modified the entity.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
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

func (s BaseScheduleChangeRequestImpl) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return s
}

func (s BaseScheduleChangeRequestImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return BaseChangeTrackedEntityImpl{
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseScheduleChangeRequestImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ScheduleChangeRequest = RawScheduleChangeRequestImpl{}

// RawScheduleChangeRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScheduleChangeRequestImpl struct {
	scheduleChangeRequest BaseScheduleChangeRequestImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawScheduleChangeRequestImpl) ScheduleChangeRequest() BaseScheduleChangeRequestImpl {
	return s.scheduleChangeRequest
}

func (s RawScheduleChangeRequestImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return s.scheduleChangeRequest.ChangeTrackedEntity()
}

func (s RawScheduleChangeRequestImpl) Entity() BaseEntityImpl {
	return s.scheduleChangeRequest.Entity()
}

var _ json.Marshaler = BaseScheduleChangeRequestImpl{}

func (s BaseScheduleChangeRequestImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseScheduleChangeRequestImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseScheduleChangeRequestImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseScheduleChangeRequestImpl: %+v", err)
	}

	delete(decoded, "managerActionDateTime")
	delete(decoded, "managerUserId")
	delete(decoded, "senderDateTime")
	delete(decoded, "senderUserId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.scheduleChangeRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseScheduleChangeRequestImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseScheduleChangeRequestImpl{}

func (s *BaseScheduleChangeRequestImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedTo            *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`
		ManagerActionDateTime nullable.Type[string]       `json:"managerActionDateTime,omitempty"`
		ManagerActionMessage  nullable.Type[string]       `json:"managerActionMessage,omitempty"`
		ManagerUserId         nullable.Type[string]       `json:"managerUserId,omitempty"`
		SenderDateTime        nullable.Type[string]       `json:"senderDateTime,omitempty"`
		SenderMessage         nullable.Type[string]       `json:"senderMessage,omitempty"`
		SenderUserId          nullable.Type[string]       `json:"senderUserId,omitempty"`
		State                 *ScheduleChangeState        `json:"state,omitempty"`
		CreatedDateTime       nullable.Type[string]       `json:"createdDateTime,omitempty"`
		LastModifiedDateTime  nullable.Type[string]       `json:"lastModifiedDateTime,omitempty"`
		Id                    *string                     `json:"id,omitempty"`
		ODataId               *string                     `json:"@odata.id,omitempty"`
		ODataType             *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedTo = decoded.AssignedTo
	s.ManagerActionDateTime = decoded.ManagerActionDateTime
	s.ManagerActionMessage = decoded.ManagerActionMessage
	s.ManagerUserId = decoded.ManagerUserId
	s.SenderDateTime = decoded.SenderDateTime
	s.SenderMessage = decoded.SenderMessage
	s.SenderUserId = decoded.SenderUserId
	s.State = decoded.State
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseScheduleChangeRequestImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseScheduleChangeRequestImpl': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}

func UnmarshalScheduleChangeRequestImplementation(input []byte) (ScheduleChangeRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduleChangeRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.offerShiftRequest") {
		var out OfferShiftRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfferShiftRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openShiftChangeRequest") {
		var out OpenShiftChangeRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenShiftChangeRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOffRequest") {
		var out TimeOffRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOffRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseScheduleChangeRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScheduleChangeRequestImpl: %+v", err)
	}

	return RawScheduleChangeRequestImpl{
		scheduleChangeRequest: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
