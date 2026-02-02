package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrivilegedAccessScheduleRequest = PrivilegedAccessGroupAssignmentScheduleRequest{}

type PrivilegedAccessGroupAssignmentScheduleRequest struct {
	// The identifier of a membership or ownership assignment relationship to the group. Required. The possible values are:
	// owner, member, unknownFutureValue.
	AccessId PrivilegedAccessGroupRelationships `json:"accessId"`

	// When the request activates a membership or ownership assignment in PIM for groups, this object represents the
	// eligibility policy for the group. Otherwise, it is null. Supports $expand.
	ActivatedUsing *PrivilegedAccessGroupEligibilitySchedule `json:"activatedUsing,omitempty"`

	// References the group that is the scope of the membership or ownership assignment request through PIM for groups.
	// Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
	Group *Group `json:"group,omitempty"`

	// The identifier of the group representing the scope of the membership or ownership assignment through PIM for groups.
	// Required.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// References the principal that's in the scope of this membership or ownership assignment request through the group
	// that's governed by PIM. Supports $expand and $select nested in $expand for id only.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// The identifier of the principal whose membership or ownership assignment to the group is managed through PIM for
	// groups. Supports $filter (eq, ne).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Schedule created by this request. Supports $expand.
	TargetSchedule *PrivilegedAccessGroupEligibilitySchedule `json:"targetSchedule,omitempty"`

	// The identifier of the schedule that's created from the membership or ownership assignment request. Supports $filter
	// (eq, ne).
	TargetScheduleId nullable.Type[string] `json:"targetScheduleId,omitempty"`

	// Fields inherited from PrivilegedAccessScheduleRequest

	// Represents the type of operation on the group membership or ownership assignment request. The possible values are:
	// adminAssign, adminUpdate, adminRemove, selfActivate, selfDeactivate, adminExtend, adminRenew. adminAssign: For
	// administrators to assign group membership or ownership to principals.adminRemove: For administrators to remove
	// principals from group membership or ownership. adminUpdate: For administrators to change existing group membership or
	// ownership assignments.adminExtend: For administrators to extend expiring assignments.adminRenew: For administrators
	// to renew expired assignments.selfActivate: For principals to activate their assignments.selfDeactivate: For
	// principals to deactivate their active assignments.
	Action *ScheduleRequestActions `json:"action,omitempty"`

	// Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an
	// activation is subject to additional rules like MFA before actually submitting the request.
	IsValidationOnly nullable.Type[bool] `json:"isValidationOnly,omitempty"`

	// A message provided by users and administrators when create they create the
	// privilegedAccessGroupAssignmentScheduleRequest object.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The period of the group membership or ownership assignment. Recurring schedules are currently unsupported.
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`

	// Ticket details linked to the group membership or ownership assignment request including details of the ticket number
	// and ticket system.
	TicketInfo *TicketInfo `json:"ticketInfo,omitempty"`

	// Fields inherited from Request

	// The identifier of the approval of the request.
	ApprovalId nullable.Type[string] `json:"approvalId,omitempty"`

	// The request completion date time.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The principal that created the request.
	CreatedBy IdentitySet `json:"createdBy"`

	// The request creation date time.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Free text field to define any custom data for the request. Not used.
	CustomData nullable.Type[string] `json:"customData,omitempty"`

	// The status of the request. Not nullable. The possible values are: Canceled, Denied, Failed, Granted,
	// PendingAdminDecision, PendingApproval, PendingProvisioning, PendingScheduleCreation, Provisioned, Revoked, and
	// ScheduleCreated. Not nullable.
	Status *string `json:"status,omitempty"`

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

func (s PrivilegedAccessGroupAssignmentScheduleRequest) PrivilegedAccessScheduleRequest() BasePrivilegedAccessScheduleRequestImpl {
	return BasePrivilegedAccessScheduleRequestImpl{
		Action:            s.Action,
		IsValidationOnly:  s.IsValidationOnly,
		Justification:     s.Justification,
		ScheduleInfo:      s.ScheduleInfo,
		TicketInfo:        s.TicketInfo,
		ApprovalId:        s.ApprovalId,
		CompletedDateTime: s.CompletedDateTime,
		CreatedBy:         s.CreatedBy,
		CreatedDateTime:   s.CreatedDateTime,
		CustomData:        s.CustomData,
		Status:            s.Status,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s PrivilegedAccessGroupAssignmentScheduleRequest) Request() BaseRequestImpl {
	return BaseRequestImpl{
		ApprovalId:        s.ApprovalId,
		CompletedDateTime: s.CompletedDateTime,
		CreatedBy:         s.CreatedBy,
		CreatedDateTime:   s.CreatedDateTime,
		CustomData:        s.CustomData,
		Status:            s.Status,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s PrivilegedAccessGroupAssignmentScheduleRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroupAssignmentScheduleRequest{}

func (s PrivilegedAccessGroupAssignmentScheduleRequest) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroupAssignmentScheduleRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroupAssignmentScheduleRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroupAssignmentScheduleRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroupAssignmentScheduleRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegedAccessGroupAssignmentScheduleRequest{}

func (s *PrivilegedAccessGroupAssignmentScheduleRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessId            PrivilegedAccessGroupRelationships        `json:"accessId"`
		ActivatedUsing      *PrivilegedAccessGroupEligibilitySchedule `json:"activatedUsing,omitempty"`
		Group               *Group                                    `json:"group,omitempty"`
		GroupId             nullable.Type[string]                     `json:"groupId,omitempty"`
		PrincipalId         nullable.Type[string]                     `json:"principalId,omitempty"`
		Principal_ODataBind *string                                   `json:"principal@odata.bind,omitempty"`
		TargetSchedule      *PrivilegedAccessGroupEligibilitySchedule `json:"targetSchedule,omitempty"`
		TargetScheduleId    nullable.Type[string]                     `json:"targetScheduleId,omitempty"`
		Action              *ScheduleRequestActions                   `json:"action,omitempty"`
		IsValidationOnly    nullable.Type[bool]                       `json:"isValidationOnly,omitempty"`
		Justification       nullable.Type[string]                     `json:"justification,omitempty"`
		ScheduleInfo        *RequestSchedule                          `json:"scheduleInfo,omitempty"`
		TicketInfo          *TicketInfo                               `json:"ticketInfo,omitempty"`
		ApprovalId          nullable.Type[string]                     `json:"approvalId,omitempty"`
		CompletedDateTime   nullable.Type[string]                     `json:"completedDateTime,omitempty"`
		CreatedDateTime     nullable.Type[string]                     `json:"createdDateTime,omitempty"`
		CustomData          nullable.Type[string]                     `json:"customData,omitempty"`
		Status              *string                                   `json:"status,omitempty"`
		Id                  *string                                   `json:"id,omitempty"`
		ODataId             *string                                   `json:"@odata.id,omitempty"`
		ODataType           *string                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessId = decoded.AccessId
	s.ActivatedUsing = decoded.ActivatedUsing
	s.Group = decoded.Group
	s.GroupId = decoded.GroupId
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.TargetSchedule = decoded.TargetSchedule
	s.TargetScheduleId = decoded.TargetScheduleId
	s.Action = decoded.Action
	s.ApprovalId = decoded.ApprovalId
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomData = decoded.CustomData
	s.Id = decoded.Id
	s.IsValidationOnly = decoded.IsValidationOnly
	s.Justification = decoded.Justification
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ScheduleInfo = decoded.ScheduleInfo
	s.Status = decoded.Status
	s.TicketInfo = decoded.TicketInfo

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrivilegedAccessGroupAssignmentScheduleRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PrivilegedAccessGroupAssignmentScheduleRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'PrivilegedAccessGroupAssignmentScheduleRequest': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
