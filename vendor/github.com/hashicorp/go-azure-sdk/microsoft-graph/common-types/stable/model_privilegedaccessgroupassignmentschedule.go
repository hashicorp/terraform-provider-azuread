package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrivilegedAccessSchedule = PrivilegedAccessGroupAssignmentSchedule{}

type PrivilegedAccessGroupAssignmentSchedule struct {
	// The identifier of the membership or ownership assignment to the group that is governed through PIM. Required. The
	// possible values are: owner, member, unknownFutureValue. Supports $filter (eq).
	AccessId PrivilegedAccessGroupRelationships `json:"accessId"`

	// When the request activates an ownership or membership assignment in PIM for groups, this object represents the
	// eligibility relationship. Otherwise, it's null. Supports $expand.
	ActivatedUsing *PrivilegedAccessGroupEligibilitySchedule `json:"activatedUsing,omitempty"`

	// Indicates whether the membership or ownership assignment for the principal is granted through activation or direct
	// assignment. Required. The possible values are: assigned, activated, unknownFutureValue. Supports $filter (eq).
	AssignmentType PrivilegedAccessGroupAssignmentType `json:"assignmentType"`

	// References the group that is the scope of the membership or ownership assignment through PIM for groups. Supports
	// $expand and $select nested in $expand for select properties like id, displayName, and mail.
	Group *Group `json:"group,omitempty"`

	// The identifier of the group representing the scope of the membership or ownership assignment through PIM for groups.
	// Required. Supports $filter (eq).
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Indicates whether the assignment is derived from a direct group assignment or through a transitive assignment. The
	// possible values are: direct, group, unknownFutureValue. Supports $filter (eq).
	MemberType *PrivilegedAccessGroupMemberType `json:"memberType,omitempty"`

	// References the principal that's in the scope of this membership or ownership assignment request to the group that's
	// governed through PIM. Supports $expand and $select nested in $expand for id only.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// The identifier of the principal whose membership or ownership assignment is granted through PIM for groups. Required.
	// Supports $filter (eq).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Fields inherited from PrivilegedAccessSchedule

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

func (s PrivilegedAccessGroupAssignmentSchedule) PrivilegedAccessSchedule() BasePrivilegedAccessScheduleImpl {
	return BasePrivilegedAccessScheduleImpl{
		CreatedDateTime:  s.CreatedDateTime,
		CreatedUsing:     s.CreatedUsing,
		ModifiedDateTime: s.ModifiedDateTime,
		ScheduleInfo:     s.ScheduleInfo,
		Status:           s.Status,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s PrivilegedAccessGroupAssignmentSchedule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroupAssignmentSchedule{}

func (s PrivilegedAccessGroupAssignmentSchedule) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroupAssignmentSchedule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroupAssignmentSchedule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroupAssignmentSchedule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroupAssignmentSchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroupAssignmentSchedule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegedAccessGroupAssignmentSchedule{}

func (s *PrivilegedAccessGroupAssignmentSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessId            PrivilegedAccessGroupRelationships        `json:"accessId"`
		ActivatedUsing      *PrivilegedAccessGroupEligibilitySchedule `json:"activatedUsing,omitempty"`
		AssignmentType      PrivilegedAccessGroupAssignmentType       `json:"assignmentType"`
		Group               *Group                                    `json:"group,omitempty"`
		GroupId             nullable.Type[string]                     `json:"groupId,omitempty"`
		MemberType          *PrivilegedAccessGroupMemberType          `json:"memberType,omitempty"`
		PrincipalId         nullable.Type[string]                     `json:"principalId,omitempty"`
		Principal_ODataBind *string                                   `json:"principal@odata.bind,omitempty"`
		CreatedDateTime     nullable.Type[string]                     `json:"createdDateTime,omitempty"`
		CreatedUsing        nullable.Type[string]                     `json:"createdUsing,omitempty"`
		ModifiedDateTime    nullable.Type[string]                     `json:"modifiedDateTime,omitempty"`
		ScheduleInfo        RequestSchedule                           `json:"scheduleInfo"`
		Status              nullable.Type[string]                     `json:"status,omitempty"`
		Id                  *string                                   `json:"id,omitempty"`
		ODataId             *string                                   `json:"@odata.id,omitempty"`
		ODataType           *string                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessId = decoded.AccessId
	s.ActivatedUsing = decoded.ActivatedUsing
	s.AssignmentType = decoded.AssignmentType
	s.Group = decoded.Group
	s.GroupId = decoded.GroupId
	s.MemberType = decoded.MemberType
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CreatedUsing = decoded.CreatedUsing
	s.Id = decoded.Id
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ScheduleInfo = decoded.ScheduleInfo
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrivilegedAccessGroupAssignmentSchedule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'PrivilegedAccessGroupAssignmentSchedule': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
