package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessSchedule interface {
	Entity
	PrivilegedAccessSchedule() BasePrivilegedAccessScheduleImpl
}

var _ PrivilegedAccessSchedule = BasePrivilegedAccessScheduleImpl{}

type BasePrivilegedAccessScheduleImpl struct {
	// When the schedule was created. Optional.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The identifier of the access assignment or eligibility request that created this schedule. Optional.
	CreatedUsing nullable.Type[string] `json:"createdUsing,omitempty"`

	// When the schedule was last modified. Optional.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// Represents the period of the access assignment or eligibility. The scheduleInfo can represent a single occurrence or
	// multiple recurring instances. Required.
	ScheduleInfo RequestSchedule `json:"scheduleInfo"`

	// The status of the access assignment or eligibility request. The possible values are: Canceled, Denied, Failed,
	// Granted, PendingAdminDecision, PendingApproval, PendingProvisioning, PendingScheduleCreation, Provisioned, Revoked,
	// and ScheduleCreated. Not nullable. Optional.
	Status nullable.Type[string] `json:"status,omitempty"`

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

func (s BasePrivilegedAccessScheduleImpl) PrivilegedAccessSchedule() BasePrivilegedAccessScheduleImpl {
	return s
}

func (s BasePrivilegedAccessScheduleImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrivilegedAccessSchedule = RawPrivilegedAccessScheduleImpl{}

// RawPrivilegedAccessScheduleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrivilegedAccessScheduleImpl struct {
	privilegedAccessSchedule BasePrivilegedAccessScheduleImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawPrivilegedAccessScheduleImpl) PrivilegedAccessSchedule() BasePrivilegedAccessScheduleImpl {
	return s.privilegedAccessSchedule
}

func (s RawPrivilegedAccessScheduleImpl) Entity() BaseEntityImpl {
	return s.privilegedAccessSchedule.Entity()
}

var _ json.Marshaler = BasePrivilegedAccessScheduleImpl{}

func (s BasePrivilegedAccessScheduleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrivilegedAccessScheduleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrivilegedAccessScheduleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrivilegedAccessScheduleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessSchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrivilegedAccessScheduleImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPrivilegedAccessScheduleImplementation(input []byte) (PrivilegedAccessSchedule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessSchedule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupAssignmentSchedule") {
		var out PrivilegedAccessGroupAssignmentSchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupAssignmentSchedule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupEligibilitySchedule") {
		var out PrivilegedAccessGroupEligibilitySchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilitySchedule: %+v", err)
		}
		return out, nil
	}

	var parent BasePrivilegedAccessScheduleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrivilegedAccessScheduleImpl: %+v", err)
	}

	return RawPrivilegedAccessScheduleImpl{
		privilegedAccessSchedule: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
