package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeTrackedEntity = TimeCard{}

type TimeCard struct {
	// The list of breaks associated with the timeCard.
	Breaks *[]TimeCardBreak `json:"breaks,omitempty"`

	// The clock-in event of the timeCard.
	ClockInEvent *TimeCardEvent `json:"clockInEvent,omitempty"`

	// The clock-out event of the timeCard.
	ClockOutEvent *TimeCardEvent `json:"clockOutEvent,omitempty"`

	// Indicates whether this timeCard entry is confirmed. Possible values are none, user, manager, unknownFutureValue.
	ConfirmedBy *ConfirmedBy `json:"confirmedBy,omitempty"`

	// Notes about the timeCard.
	Notes *ItemBody `json:"notes,omitempty"`

	// The original timeCardEntry of the timeCard, before user edits.
	OriginalEntry *TimeCardEntry `json:"originalEntry,omitempty"`

	// The current state of the timeCard during its life cycle.Possible values are: clockedIn, onBreak, clockedOut,
	// unknownFutureValue.
	State *TimeCardState `json:"state,omitempty"`

	// User ID to which the timeCard belongs.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s TimeCard) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
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

func (s TimeCard) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TimeCard{}

func (s TimeCard) MarshalJSON() ([]byte, error) {
	type wrapper TimeCard
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TimeCard: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TimeCard: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.timeCard"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TimeCard: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TimeCard{}

func (s *TimeCard) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Breaks               *[]TimeCardBreak      `json:"breaks,omitempty"`
		ClockInEvent         *TimeCardEvent        `json:"clockInEvent,omitempty"`
		ClockOutEvent        *TimeCardEvent        `json:"clockOutEvent,omitempty"`
		ConfirmedBy          *ConfirmedBy          `json:"confirmedBy,omitempty"`
		Notes                *ItemBody             `json:"notes,omitempty"`
		OriginalEntry        *TimeCardEntry        `json:"originalEntry,omitempty"`
		State                *TimeCardState        `json:"state,omitempty"`
		UserId               nullable.Type[string] `json:"userId,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Breaks = decoded.Breaks
	s.ClockInEvent = decoded.ClockInEvent
	s.ClockOutEvent = decoded.ClockOutEvent
	s.ConfirmedBy = decoded.ConfirmedBy
	s.Notes = decoded.Notes
	s.OriginalEntry = decoded.OriginalEntry
	s.State = decoded.State
	s.UserId = decoded.UserId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TimeCard into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TimeCard': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TimeCard': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
