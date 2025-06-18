package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeTrackedEntity = SchedulingGroup{}

type SchedulingGroup struct {
	// The code for the schedulingGroup to represent an external identifier. This field must be unique within the team in
	// Microsoft Teams and uses an alphanumeric format, with a maximum of 100 characters.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The display name for the schedulingGroup. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// The list of user IDs that are a member of the schedulingGroup. Required.
	UserIds []string `json:"userIds"`

	// Fields inherited from ChangeTrackedEntity

	// Identity of the creator of the entity.
	CreatedBy IdentitySet `json:"createdBy"`

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

func (s SchedulingGroup) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
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

func (s SchedulingGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SchedulingGroup{}

func (s SchedulingGroup) MarshalJSON() ([]byte, error) {
	type wrapper SchedulingGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SchedulingGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SchedulingGroup: %+v", err)
	}

	delete(decoded, "isActive")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.schedulingGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SchedulingGroup: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SchedulingGroup{}

func (s *SchedulingGroup) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Code                 nullable.Type[string] `json:"code,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		IsActive             nullable.Type[bool]   `json:"isActive,omitempty"`
		UserIds              []string              `json:"userIds"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Code = decoded.Code
	s.DisplayName = decoded.DisplayName
	s.IsActive = decoded.IsActive
	s.UserIds = decoded.UserIds
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SchedulingGroup into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SchedulingGroup': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SchedulingGroup': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
