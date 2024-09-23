package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessScheduleRequest interface {
	Entity
	Request
	PrivilegedAccessScheduleRequest() BasePrivilegedAccessScheduleRequestImpl
}

var _ PrivilegedAccessScheduleRequest = BasePrivilegedAccessScheduleRequestImpl{}

type BasePrivilegedAccessScheduleRequestImpl struct {
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

func (s BasePrivilegedAccessScheduleRequestImpl) PrivilegedAccessScheduleRequest() BasePrivilegedAccessScheduleRequestImpl {
	return s
}

func (s BasePrivilegedAccessScheduleRequestImpl) Request() BaseRequestImpl {
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

func (s BasePrivilegedAccessScheduleRequestImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrivilegedAccessScheduleRequest = RawPrivilegedAccessScheduleRequestImpl{}

// RawPrivilegedAccessScheduleRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrivilegedAccessScheduleRequestImpl struct {
	privilegedAccessScheduleRequest BasePrivilegedAccessScheduleRequestImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawPrivilegedAccessScheduleRequestImpl) PrivilegedAccessScheduleRequest() BasePrivilegedAccessScheduleRequestImpl {
	return s.privilegedAccessScheduleRequest
}

func (s RawPrivilegedAccessScheduleRequestImpl) Request() BaseRequestImpl {
	return s.privilegedAccessScheduleRequest.Request()
}

func (s RawPrivilegedAccessScheduleRequestImpl) Entity() BaseEntityImpl {
	return s.privilegedAccessScheduleRequest.Entity()
}

var _ json.Marshaler = BasePrivilegedAccessScheduleRequestImpl{}

func (s BasePrivilegedAccessScheduleRequestImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrivilegedAccessScheduleRequestImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrivilegedAccessScheduleRequestImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrivilegedAccessScheduleRequestImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedAccessScheduleRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrivilegedAccessScheduleRequestImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BasePrivilegedAccessScheduleRequestImpl{}

func (s *BasePrivilegedAccessScheduleRequestImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action            *ScheduleRequestActions `json:"action,omitempty"`
		IsValidationOnly  nullable.Type[bool]     `json:"isValidationOnly,omitempty"`
		Justification     nullable.Type[string]   `json:"justification,omitempty"`
		ScheduleInfo      *RequestSchedule        `json:"scheduleInfo,omitempty"`
		TicketInfo        *TicketInfo             `json:"ticketInfo,omitempty"`
		ApprovalId        nullable.Type[string]   `json:"approvalId,omitempty"`
		CompletedDateTime nullable.Type[string]   `json:"completedDateTime,omitempty"`
		CreatedDateTime   nullable.Type[string]   `json:"createdDateTime,omitempty"`
		CustomData        nullable.Type[string]   `json:"customData,omitempty"`
		Status            *string                 `json:"status,omitempty"`
		Id                *string                 `json:"id,omitempty"`
		ODataId           *string                 `json:"@odata.id,omitempty"`
		ODataType         *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.IsValidationOnly = decoded.IsValidationOnly
	s.Justification = decoded.Justification
	s.ScheduleInfo = decoded.ScheduleInfo
	s.TicketInfo = decoded.TicketInfo
	s.ApprovalId = decoded.ApprovalId
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomData = decoded.CustomData
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BasePrivilegedAccessScheduleRequestImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BasePrivilegedAccessScheduleRequestImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalPrivilegedAccessScheduleRequestImplementation(input []byte) (PrivilegedAccessScheduleRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessScheduleRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest") {
		var out PrivilegedAccessGroupAssignmentScheduleRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupAssignmentScheduleRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest") {
		var out PrivilegedAccessGroupEligibilityScheduleRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilityScheduleRequest: %+v", err)
		}
		return out, nil
	}

	var parent BasePrivilegedAccessScheduleRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrivilegedAccessScheduleRequestImpl: %+v", err)
	}

	return RawPrivilegedAccessScheduleRequestImpl{
		privilegedAccessScheduleRequest: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
