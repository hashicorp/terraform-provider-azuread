package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeTrackedEntity interface {
	Entity
	ChangeTrackedEntity() BaseChangeTrackedEntityImpl
}

var _ ChangeTrackedEntity = BaseChangeTrackedEntityImpl{}

type BaseChangeTrackedEntityImpl struct {
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

func (s BaseChangeTrackedEntityImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return s
}

func (s BaseChangeTrackedEntityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ChangeTrackedEntity = RawChangeTrackedEntityImpl{}

// RawChangeTrackedEntityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawChangeTrackedEntityImpl struct {
	changeTrackedEntity BaseChangeTrackedEntityImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawChangeTrackedEntityImpl) ChangeTrackedEntity() BaseChangeTrackedEntityImpl {
	return s.changeTrackedEntity
}

func (s RawChangeTrackedEntityImpl) Entity() BaseEntityImpl {
	return s.changeTrackedEntity.Entity()
}

var _ json.Marshaler = BaseChangeTrackedEntityImpl{}

func (s BaseChangeTrackedEntityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseChangeTrackedEntityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseChangeTrackedEntityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseChangeTrackedEntityImpl: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.changeTrackedEntity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseChangeTrackedEntityImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseChangeTrackedEntityImpl{}

func (s *BaseChangeTrackedEntityImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseChangeTrackedEntityImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseChangeTrackedEntityImpl': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}

func UnmarshalChangeTrackedEntityImplementation(input []byte) (ChangeTrackedEntity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChangeTrackedEntity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.openShift") {
		var out OpenShift
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenShift: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.scheduleChangeRequest") {
		var out ScheduleChangeRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScheduleChangeRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.schedulingGroup") {
		var out SchedulingGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SchedulingGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shift") {
		var out Shift
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Shift: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.shiftPreferences") {
		var out ShiftPreferences
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftPreferences: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOff") {
		var out TimeOff
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOff: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOffReason") {
		var out TimeOffReason
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOffReason: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.workforceIntegration") {
		var out WorkforceIntegration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkforceIntegration: %+v", err)
		}
		return out, nil
	}

	var parent BaseChangeTrackedEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseChangeTrackedEntityImpl: %+v", err)
	}

	return RawChangeTrackedEntityImpl{
		changeTrackedEntity: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
