package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Request = UnifiedRoleAssignmentScheduleRequest{}

type UnifiedRoleAssignmentScheduleRequest struct {
	// Represents the type of the operation on the role assignment request. The possible values are: adminAssign,
	// adminUpdate, adminRemove, selfActivate, selfDeactivate, adminExtend, adminRenew, selfExtend, selfRenew,
	// unknownFutureValue. adminAssign: For administrators to assign roles to principals.adminRemove: For administrators to
	// remove principals from roles. adminUpdate: For administrators to change existing role assignments.adminExtend: For
	// administrators to extend expiring assignments.adminRenew: For administrators to renew expired
	// assignments.selfActivate: For principals to activate their assignments.selfDeactivate: For principals to deactivate
	// their active assignments.selfExtend: For principals to request to extend their expiring assignments.selfRenew: For
	// principals to request to renew their expired assignments.
	Action *UnifiedRoleScheduleRequestActions `json:"action,omitempty"`

	// If the request is from an eligible administrator to activate a role, this parameter will show the related eligible
	// assignment for that activation. Otherwise, it's null. Supports $expand and $select nested in $expand.
	ActivatedUsing *UnifiedRoleEligibilitySchedule `json:"activatedUsing,omitempty"`

	// Read-only property with details of the app-specific scope when the assignment is scoped to an app. Nullable. Supports
	// $expand.
	AppScope *AppScope `json:"appScope,omitempty"`

	// Identifier of the app-specific scope when the assignment is scoped to an app. The scope of an assignment determines
	// the set of resources for which the principal has been granted access. App scopes are scopes that are defined and
	// understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to
	// particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
	AppScopeId nullable.Type[string] `json:"appScopeId,omitempty"`

	// The directory object that is the scope of the assignment. Read-only. Supports $expand.
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`

	// Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines
	// the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in
	// the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the
	// scope to an application only. Supports $filter (eq, ne, and on null values).
	DirectoryScopeId nullable.Type[string] `json:"directoryScopeId,omitempty"`

	// OData ID for `DirectoryScope` to bind to this entity
	DirectoryScope_ODataBind *string `json:"directoryScope@odata.bind,omitempty"`

	// Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an
	// activation is subject to additional rules like MFA before actually submitting the request.
	IsValidationOnly nullable.Type[bool] `json:"isValidationOnly,omitempty"`

	// A message provided by users and administrators when create they create the unifiedRoleAssignmentScheduleRequest
	// object.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The principal that's getting a role assignment through the request. Supports $expand and $select nested in $expand
	// for id only.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// Identifier of the principal that has been granted the assignment. Can be a user, role-assignable group, or a service
	// principal. Supports $filter (eq, ne).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property.
	// Supports $expand and $select nested in $expand.
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`

	// Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The period of the role assignment. Recurring schedules are currently unsupported.
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`

	// The schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports
	// $expand and $select nested in $expand.
	TargetSchedule *UnifiedRoleAssignmentSchedule `json:"targetSchedule,omitempty"`

	// Identifier of the schedule object that's linked to the assignment request. Supports $filter (eq, ne).
	TargetScheduleId nullable.Type[string] `json:"targetScheduleId,omitempty"`

	// Ticket details linked to the role assignment request including details of the ticket number and ticket system.
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

func (s UnifiedRoleAssignmentScheduleRequest) Request() BaseRequestImpl {
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

func (s UnifiedRoleAssignmentScheduleRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleAssignmentScheduleRequest{}

func (s UnifiedRoleAssignmentScheduleRequest) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleAssignmentScheduleRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleAssignmentScheduleRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleAssignmentScheduleRequest: %+v", err)
	}

	delete(decoded, "directoryScope")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleAssignmentScheduleRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleAssignmentScheduleRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleAssignmentScheduleRequest{}

func (s *UnifiedRoleAssignmentScheduleRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action                   *UnifiedRoleScheduleRequestActions `json:"action,omitempty"`
		ActivatedUsing           *UnifiedRoleEligibilitySchedule    `json:"activatedUsing,omitempty"`
		AppScope                 *AppScope                          `json:"appScope,omitempty"`
		AppScopeId               nullable.Type[string]              `json:"appScopeId,omitempty"`
		DirectoryScopeId         nullable.Type[string]              `json:"directoryScopeId,omitempty"`
		DirectoryScope_ODataBind *string                            `json:"directoryScope@odata.bind,omitempty"`
		IsValidationOnly         nullable.Type[bool]                `json:"isValidationOnly,omitempty"`
		Justification            nullable.Type[string]              `json:"justification,omitempty"`
		PrincipalId              nullable.Type[string]              `json:"principalId,omitempty"`
		Principal_ODataBind      *string                            `json:"principal@odata.bind,omitempty"`
		RoleDefinition           *UnifiedRoleDefinition             `json:"roleDefinition,omitempty"`
		RoleDefinitionId         nullable.Type[string]              `json:"roleDefinitionId,omitempty"`
		ScheduleInfo             *RequestSchedule                   `json:"scheduleInfo,omitempty"`
		TargetSchedule           *UnifiedRoleAssignmentSchedule     `json:"targetSchedule,omitempty"`
		TargetScheduleId         nullable.Type[string]              `json:"targetScheduleId,omitempty"`
		TicketInfo               *TicketInfo                        `json:"ticketInfo,omitempty"`
		ApprovalId               nullable.Type[string]              `json:"approvalId,omitempty"`
		CompletedDateTime        nullable.Type[string]              `json:"completedDateTime,omitempty"`
		CreatedDateTime          nullable.Type[string]              `json:"createdDateTime,omitempty"`
		CustomData               nullable.Type[string]              `json:"customData,omitempty"`
		Status                   *string                            `json:"status,omitempty"`
		Id                       *string                            `json:"id,omitempty"`
		ODataId                  *string                            `json:"@odata.id,omitempty"`
		ODataType                *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.ActivatedUsing = decoded.ActivatedUsing
	s.AppScope = decoded.AppScope
	s.AppScopeId = decoded.AppScopeId
	s.DirectoryScopeId = decoded.DirectoryScopeId
	s.DirectoryScope_ODataBind = decoded.DirectoryScope_ODataBind
	s.IsValidationOnly = decoded.IsValidationOnly
	s.Justification = decoded.Justification
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.RoleDefinition = decoded.RoleDefinition
	s.RoleDefinitionId = decoded.RoleDefinitionId
	s.ScheduleInfo = decoded.ScheduleInfo
	s.TargetSchedule = decoded.TargetSchedule
	s.TargetScheduleId = decoded.TargetScheduleId
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
		return fmt.Errorf("unmarshaling UnifiedRoleAssignmentScheduleRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'UnifiedRoleAssignmentScheduleRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["directoryScope"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DirectoryScope' for 'UnifiedRoleAssignmentScheduleRequest': %+v", err)
		}
		s.DirectoryScope = &impl
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'UnifiedRoleAssignmentScheduleRequest': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
