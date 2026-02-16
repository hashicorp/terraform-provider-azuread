package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrivilegedAccessScheduleInstance = PrivilegedAccessGroupEligibilityScheduleInstance{}

type PrivilegedAccessGroupEligibilityScheduleInstance struct {
	// The identifier of the membership or ownership eligibility relationship to the group. Required. The possible values
	// are: owner, member. Supports $filter (eq).
	AccessId PrivilegedAccessGroupRelationships `json:"accessId"`

	// The identifier of the privilegedAccessGroupEligibilitySchedule from which this instance was created. Required.
	// Supports $filter (eq, ne).
	EligibilityScheduleId nullable.Type[string] `json:"eligibilityScheduleId,omitempty"`

	// References the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports
	// $expand.
	Group *Group `json:"group,omitempty"`

	// The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups.
	// Required. Supports $filter (eq).
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Indicates whether the assignment is derived from a group assignment. It can further imply whether the calling
	// principal can manage the assignment schedule. Required. The possible values are: direct, group, unknownFutureValue.
	// Supports $filter (eq).
	MemberType PrivilegedAccessGroupMemberType `json:"memberType"`

	// References the principal that's in the scope of the membership or ownership eligibility request through the group
	// that's governed by PIM. Supports $expand.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for
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

func (s PrivilegedAccessGroupEligibilityScheduleInstance) PrivilegedAccessScheduleInstance() BasePrivilegedAccessScheduleInstanceImpl {
	return BasePrivilegedAccessScheduleInstanceImpl{
		EndDateTime:   s.EndDateTime,
		StartDateTime: s.StartDateTime,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s PrivilegedAccessGroupEligibilityScheduleInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroupEligibilityScheduleInstance{}

func (s PrivilegedAccessGroupEligibilityScheduleInstance) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroupEligibilityScheduleInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegedAccessGroupEligibilityScheduleInstance{}

func (s *PrivilegedAccessGroupEligibilityScheduleInstance) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessId              PrivilegedAccessGroupRelationships `json:"accessId"`
		EligibilityScheduleId nullable.Type[string]              `json:"eligibilityScheduleId,omitempty"`
		Group                 *Group                             `json:"group,omitempty"`
		GroupId               nullable.Type[string]              `json:"groupId,omitempty"`
		MemberType            PrivilegedAccessGroupMemberType    `json:"memberType"`
		PrincipalId           nullable.Type[string]              `json:"principalId,omitempty"`
		Principal_ODataBind   *string                            `json:"principal@odata.bind,omitempty"`
		EndDateTime           nullable.Type[string]              `json:"endDateTime,omitempty"`
		StartDateTime         nullable.Type[string]              `json:"startDateTime,omitempty"`
		Id                    *string                            `json:"id,omitempty"`
		ODataId               *string                            `json:"@odata.id,omitempty"`
		ODataType             *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessId = decoded.AccessId
	s.EligibilityScheduleId = decoded.EligibilityScheduleId
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
		return fmt.Errorf("unmarshaling PrivilegedAccessGroupEligibilityScheduleInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'PrivilegedAccessGroupEligibilityScheduleInstance': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
