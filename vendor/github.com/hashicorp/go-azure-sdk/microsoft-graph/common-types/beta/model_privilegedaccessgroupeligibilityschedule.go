package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrivilegedAccessSchedule = PrivilegedAccessGroupEligibilitySchedule{}

type PrivilegedAccessGroupEligibilitySchedule struct {
	// The identifier of the membership or ownership eligibility to the group that is governed by PIM. Required. The
	// possible values are: owner, member. Supports $filter (eq).
	AccessId PrivilegedAccessGroupRelationships `json:"accessId"`

	// References the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports
	// $expand.
	Group *Group `json:"group,omitempty"`

	// The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups.
	// Required. Supports $filter (eq).
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Indicates whether the assignment is derived from a group assignment. It can further imply whether the caller can
	// manage the schedule. Required. The possible values are: direct, group, unknownFutureValue. Supports $filter (eq).
	MemberType PrivilegedAccessGroupMemberType `json:"memberType"`

	// References the principal that's in the scope of this membership or ownership eligibility request to the group that's
	// governed by PIM. Supports $expand.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// The identifier of the principal whose membership or ownership eligibility is granted through PIM for groups.
	// Required. Supports $filter (eq).
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

func (s PrivilegedAccessGroupEligibilitySchedule) PrivilegedAccessSchedule() BasePrivilegedAccessScheduleImpl {
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

func (s PrivilegedAccessGroupEligibilitySchedule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroupEligibilitySchedule{}

func (s PrivilegedAccessGroupEligibilitySchedule) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroupEligibilitySchedule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroupEligibilitySchedule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroupEligibilitySchedule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroupEligibilitySchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroupEligibilitySchedule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegedAccessGroupEligibilitySchedule{}

func (s *PrivilegedAccessGroupEligibilitySchedule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessId            PrivilegedAccessGroupRelationships `json:"accessId"`
		Group               *Group                             `json:"group,omitempty"`
		GroupId             nullable.Type[string]              `json:"groupId,omitempty"`
		MemberType          PrivilegedAccessGroupMemberType    `json:"memberType"`
		PrincipalId         nullable.Type[string]              `json:"principalId,omitempty"`
		Principal_ODataBind *string                            `json:"principal@odata.bind,omitempty"`
		CreatedDateTime     nullable.Type[string]              `json:"createdDateTime,omitempty"`
		CreatedUsing        nullable.Type[string]              `json:"createdUsing,omitempty"`
		ModifiedDateTime    nullable.Type[string]              `json:"modifiedDateTime,omitempty"`
		ScheduleInfo        RequestSchedule                    `json:"scheduleInfo"`
		Status              nullable.Type[string]              `json:"status,omitempty"`
		Id                  *string                            `json:"id,omitempty"`
		ODataId             *string                            `json:"@odata.id,omitempty"`
		ODataType           *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessId = decoded.AccessId
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
		return fmt.Errorf("unmarshaling PrivilegedAccessGroupEligibilitySchedule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'PrivilegedAccessGroupEligibilitySchedule': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
