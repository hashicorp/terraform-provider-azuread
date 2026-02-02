package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedAccessGroup{}

type PrivilegedAccessGroup struct {
	AssignmentApprovals *[]Approval `json:"assignmentApprovals,omitempty"`

	// The instances of assignment schedules to activate a just-in-time access.
	AssignmentScheduleInstances *[]PrivilegedAccessGroupAssignmentScheduleInstance `json:"assignmentScheduleInstances,omitempty"`

	// The schedule requests for operations to create, update, delete, extend, and renew an assignment.
	AssignmentScheduleRequests *[]PrivilegedAccessGroupAssignmentScheduleRequest `json:"assignmentScheduleRequests,omitempty"`

	// The assignment schedules to activate a just-in-time access.
	AssignmentSchedules *[]PrivilegedAccessGroupAssignmentSchedule `json:"assignmentSchedules,omitempty"`

	// The instances of eligibility schedules to activate a just-in-time access.
	EligibilityScheduleInstances *[]PrivilegedAccessGroupEligibilityScheduleInstance `json:"eligibilityScheduleInstances,omitempty"`

	// The schedule requests for operations to create, update, delete, extend, and renew an eligibility.
	EligibilityScheduleRequests *[]PrivilegedAccessGroupEligibilityScheduleRequest `json:"eligibilityScheduleRequests,omitempty"`

	// The eligibility schedules to activate a just-in-time access.
	EligibilitySchedules *[]PrivilegedAccessGroupEligibilitySchedule `json:"eligibilitySchedules,omitempty"`

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

func (s PrivilegedAccessGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroup{}

func (s PrivilegedAccessGroup) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroup: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroup: %+v", err)
	}

	return encoded, nil
}
