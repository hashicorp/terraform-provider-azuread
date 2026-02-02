package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessScheduleInstance interface {
	Entity
	PrivilegedAccessScheduleInstance() BasePrivilegedAccessScheduleInstanceImpl
}

var _ PrivilegedAccessScheduleInstance = BasePrivilegedAccessScheduleInstanceImpl{}

type BasePrivilegedAccessScheduleInstanceImpl struct {
	// When the schedule instance ends. Required.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// When this instance starts. Required.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

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

func (s BasePrivilegedAccessScheduleInstanceImpl) PrivilegedAccessScheduleInstance() BasePrivilegedAccessScheduleInstanceImpl {
	return s
}

func (s BasePrivilegedAccessScheduleInstanceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrivilegedAccessScheduleInstance = RawPrivilegedAccessScheduleInstanceImpl{}

// RawPrivilegedAccessScheduleInstanceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrivilegedAccessScheduleInstanceImpl struct {
	privilegedAccessScheduleInstance BasePrivilegedAccessScheduleInstanceImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawPrivilegedAccessScheduleInstanceImpl) PrivilegedAccessScheduleInstance() BasePrivilegedAccessScheduleInstanceImpl {
	return s.privilegedAccessScheduleInstance
}

func (s RawPrivilegedAccessScheduleInstanceImpl) Entity() BaseEntityImpl {
	return s.privilegedAccessScheduleInstance.Entity()
}

var _ json.Marshaler = BasePrivilegedAccessScheduleInstanceImpl{}

func (s BasePrivilegedAccessScheduleInstanceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrivilegedAccessScheduleInstanceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrivilegedAccessScheduleInstanceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrivilegedAccessScheduleInstanceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessScheduleInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrivilegedAccessScheduleInstanceImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPrivilegedAccessScheduleInstanceImplementation(input []byte) (PrivilegedAccessScheduleInstance, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessScheduleInstance into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance") {
		var out PrivilegedAccessGroupAssignmentScheduleInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupAssignmentScheduleInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance") {
		var out PrivilegedAccessGroupEligibilityScheduleInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
		}
		return out, nil
	}

	var parent BasePrivilegedAccessScheduleInstanceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrivilegedAccessScheduleInstanceImpl: %+v", err)
	}

	return RawPrivilegedAccessScheduleInstanceImpl{
		privilegedAccessScheduleInstance: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
