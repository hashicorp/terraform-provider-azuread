package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeTrackedEntity = OpenShift{}

type OpenShift struct {
	// Draft changes in the openShift are only visible to managers until they're shared.
	DraftOpenShift *OpenShiftItem `json:"draftOpenShift,omitempty"`

	// The openShift is marked for deletion, a process that is finalized when the schedule is shared.
	IsStagedForDeletion nullable.Type[bool] `json:"isStagedForDeletion,omitempty"`

	// The ID of the schedulingGroup that contains the openShift.
	SchedulingGroupId nullable.Type[string] `json:"schedulingGroupId,omitempty"`

	// The shared version of this openShift that is viewable by both employees and managers.
	SharedOpenShift *OpenShiftItem `json:"sharedOpenShift,omitempty"`

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

func (s OpenShift) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
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

func (s OpenShift) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OpenShift{}

func (s OpenShift) MarshalJSON() ([]byte, error) {
	type wrapper OpenShift
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenShift: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenShift: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.openShift"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenShift: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OpenShift{}

func (s *OpenShift) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DraftOpenShift       *OpenShiftItem        `json:"draftOpenShift,omitempty"`
		IsStagedForDeletion  nullable.Type[bool]   `json:"isStagedForDeletion,omitempty"`
		SchedulingGroupId    nullable.Type[string] `json:"schedulingGroupId,omitempty"`
		SharedOpenShift      *OpenShiftItem        `json:"sharedOpenShift,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DraftOpenShift = decoded.DraftOpenShift
	s.IsStagedForDeletion = decoded.IsStagedForDeletion
	s.SchedulingGroupId = decoded.SchedulingGroupId
	s.SharedOpenShift = decoded.SharedOpenShift
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OpenShift into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'OpenShift': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'OpenShift': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
