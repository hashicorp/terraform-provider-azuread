package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrivilegedAccessScheduleInstance = PrivilegedAccessGroupAssignmentScheduleInstance{}

type PrivilegedAccessGroupAssignmentScheduleInstance struct {
	// The identifier of the membership or ownership assignment relationship to the group. Required. The possible values
	// are: owner, member, unknownFutureValue. Supports $filter (eq).
	AccessId PrivilegedAccessGroupRelationships `json:"accessId"`

	// When the request activates a membership or ownership in PIM for groups, this object represents the eligibility
	// request for the group. Otherwise, it is null.
	ActivatedUsing *PrivilegedAccessGroupEligibilityScheduleInstance `json:"activatedUsing,omitempty"`

	// The identifier of the privilegedAccessGroupAssignmentSchedule from which this instance was created. Required.
	// Supports $filter (eq, ne).
	AssignmentScheduleId nullable.Type[string] `json:"assignmentScheduleId,omitempty"`

	// Indicates whether the membership or ownership assignment is granted through activation of an eligibility or through
	// direct assignment. Required. The possible values are: assigned, activated, unknownFutureValue. Supports $filter (eq).
	AssignmentType PrivilegedAccessGroupAssignmentType `json:"assignmentType"`

	// References the group that is the scope of the membership or ownership assignment through PIM for groups. Supports
	// $expand.
	Group *Group `json:"group,omitempty"`

	// The identifier of the group representing the scope of the membership or ownership assignment through PIM for groups.
	// Optional. Supports $filter (eq).
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Indicates whether the assignment is derived from a group assignment. It can further imply whether the caller can
	// manage the assignment schedule. Required. The possible values are: direct, group, unknownFutureValue. Supports
	// $filter (eq).
	MemberType PrivilegedAccessGroupMemberType `json:"memberType"`

	// References the principal that's in the scope of the membership or ownership assignment request through the group
	// that's governed by PIM. Supports $expand.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// The identifier of the principal whose membership or ownership assignment to the group is managed through PIM for
	// groups. Required. Supports $filter (eq).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Fields inherited from PrivilegedAccessScheduleInstance

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

func (s PrivilegedAccessGroupAssignmentScheduleInstance) PrivilegedAccessScheduleInstance() BasePrivilegedAccessScheduleInstanceImpl {
	return BasePrivilegedAccessScheduleInstanceImpl{
		EndDateTime:   s.EndDateTime,
		StartDateTime: s.StartDateTime,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s PrivilegedAccessGroupAssignmentScheduleInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroupAssignmentScheduleInstance{}

func (s PrivilegedAccessGroupAssignmentScheduleInstance) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroupAssignmentScheduleInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroupAssignmentScheduleInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroupAssignmentScheduleInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroupAssignmentScheduleInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegedAccessGroupAssignmentScheduleInstance{}

func (s *PrivilegedAccessGroupAssignmentScheduleInstance) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessId             PrivilegedAccessGroupRelationships                `json:"accessId"`
		ActivatedUsing       *PrivilegedAccessGroupEligibilityScheduleInstance `json:"activatedUsing,omitempty"`
		AssignmentScheduleId nullable.Type[string]                             `json:"assignmentScheduleId,omitempty"`
		AssignmentType       PrivilegedAccessGroupAssignmentType               `json:"assignmentType"`
		Group                *Group                                            `json:"group,omitempty"`
		GroupId              nullable.Type[string]                             `json:"groupId,omitempty"`
		MemberType           PrivilegedAccessGroupMemberType                   `json:"memberType"`
		PrincipalId          nullable.Type[string]                             `json:"principalId,omitempty"`
		Principal_ODataBind  *string                                           `json:"principal@odata.bind,omitempty"`
		EndDateTime          nullable.Type[string]                             `json:"endDateTime,omitempty"`
		StartDateTime        nullable.Type[string]                             `json:"startDateTime,omitempty"`
		Id                   *string                                           `json:"id,omitempty"`
		ODataId              *string                                           `json:"@odata.id,omitempty"`
		ODataType            *string                                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessId = decoded.AccessId
	s.ActivatedUsing = decoded.ActivatedUsing
	s.AssignmentScheduleId = decoded.AssignmentScheduleId
	s.AssignmentType = decoded.AssignmentType
	s.Group = decoded.Group
	s.GroupId = decoded.GroupId
	s.MemberType = decoded.MemberType
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.EndDateTime = decoded.EndDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.StartDateTime = decoded.StartDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrivilegedAccessGroupAssignmentScheduleInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'PrivilegedAccessGroupAssignmentScheduleInstance': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
