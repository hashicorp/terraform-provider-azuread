package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeTrackedEntity = TimeOff{}

type TimeOff struct {
	// The draft version of this timeOff item that is viewable by managers. It must be shared before it is visible to team
	// members. Required.
	DraftTimeOff TimeOffItem `json:"draftTimeOff"`

	// The timeOff is marked for deletion, a process that is finalized when the schedule is shared.
	IsStagedForDeletion nullable.Type[bool] `json:"isStagedForDeletion,omitempty"`

	// The shared version of this timeOff that is viewable by both employees and managers. Updates to the sharedTimeOff
	// property send notifications to users in the Teams client. Required.
	SharedTimeOff TimeOffItem `json:"sharedTimeOff"`

	// Information of the team that the timeOff is in.
	TeamInfo *ShiftsTeamInfo `json:"teamInfo,omitempty"`

	// ID of the user assigned to the timeOff. Required.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Information of the user assigned to the timeOff.
	UserInfo *ShiftsUserInfo `json:"userInfo,omitempty"`

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

func (s TimeOff) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
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

func (s TimeOff) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TimeOff{}

func (s TimeOff) MarshalJSON() ([]byte, error) {
	type wrapper TimeOff
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TimeOff: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TimeOff: %+v", err)
	}

	delete(decoded, "teamInfo")
	delete(decoded, "userInfo")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.timeOff"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TimeOff: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TimeOff{}

func (s *TimeOff) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DraftTimeOff         TimeOffItem           `json:"draftTimeOff"`
		IsStagedForDeletion  nullable.Type[bool]   `json:"isStagedForDeletion,omitempty"`
		SharedTimeOff        TimeOffItem           `json:"sharedTimeOff"`
		TeamInfo             *ShiftsTeamInfo       `json:"teamInfo,omitempty"`
		UserId               nullable.Type[string] `json:"userId,omitempty"`
		UserInfo             *ShiftsUserInfo       `json:"userInfo,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DraftTimeOff = decoded.DraftTimeOff
	s.IsStagedForDeletion = decoded.IsStagedForDeletion
	s.SharedTimeOff = decoded.SharedTimeOff
	s.TeamInfo = decoded.TeamInfo
	s.UserId = decoded.UserId
	s.UserInfo = decoded.UserInfo
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TimeOff into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TimeOff': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TimeOff': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
