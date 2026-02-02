package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Request = UnifiedRoleEligibilityScheduleRequest{}

type UnifiedRoleEligibilityScheduleRequest struct {
	// Represents the type of operation on the role eligibility request. The possible values are: adminAssign, adminUpdate,
	// adminRemove, selfActivate, selfDeactivate, adminExtend, adminRenew, selfExtend, selfRenew, unknownFutureValue.
	// adminAssign: For administrators to assign eligible roles to principals.adminRemove: For administrators to remove
	// eligible roles from principals. adminUpdate: For administrators to change existing role eligibilities.adminExtend:
	// For administrators to extend expiring role eligibilities.adminRenew: For administrators to renew expired
	// eligibilities.selfActivate: For users to activate their assignments.selfDeactivate: For users to deactivate their
	// active assignments.selfExtend: For users to request to extend their expiring assignments.selfRenew: For users to
	// request to renew their expired assignments.
	Action *UnifiedRoleScheduleRequestActions `json:"action,omitempty"`

	// Read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable.
	// Supports $expand.
	AppScope *AppScope `json:"appScope,omitempty"`

	// Identifier of the app-specific scope when the role eligibility is scoped to an app. The scope of a role eligibility
	// determines the set of resources for which the principal is eligible to access. App scopes are scopes that are defined
	// and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to
	// particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
	AppScopeId nullable.Type[string] `json:"appScopeId,omitempty"`

	// The directory object that is the scope of the role eligibility. Read-only. Supports $expand.
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`

	// Identifier of the directory object representing the scope of the role eligibility. The scope of a role eligibility
	// determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes
	// stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to
	// limit the scope to an application only. Supports $filter (eq, ne, and on null values).
	DirectoryScopeId nullable.Type[string] `json:"directoryScopeId,omitempty"`

	// OData ID for `DirectoryScope` to bind to this entity
	DirectoryScope_ODataBind *string `json:"directoryScope@odata.bind,omitempty"`

	// Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an
	// activation is subject to additional rules like MFA before actually submitting the request.
	IsValidationOnly nullable.Type[bool] `json:"isValidationOnly,omitempty"`

	// A message provided by users and administrators when create they create the unifiedRoleEligibilityScheduleRequest
	// object.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The principal that's getting a role eligibility through the request. Supports $expand.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// Identifier of the principal that has been granted the role eligibility. Can be a user or a role-assignable group. You
	// can grant only active assignments service principals.Supports $filter (eq, ne).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property.
	// Supports $expand.
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`

	// Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The period of the role eligibility. Recurring schedules are currently unsupported.
	ScheduleInfo *RequestSchedule `json:"scheduleInfo,omitempty"`

	// The schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
	TargetSchedule *UnifiedRoleEligibilitySchedule `json:"targetSchedule,omitempty"`

	// Identifier of the schedule object that's linked to the eligibility request. Supports $filter (eq, ne).
	TargetScheduleId nullable.Type[string] `json:"targetScheduleId,omitempty"`

	// Ticket details linked to the role eligibility request including details of the ticket number and ticket system.
	// Optional.
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

func (s UnifiedRoleEligibilityScheduleRequest) Request() BaseRequestImpl {
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

func (s UnifiedRoleEligibilityScheduleRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleEligibilityScheduleRequest{}

func (s UnifiedRoleEligibilityScheduleRequest) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleEligibilityScheduleRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleEligibilityScheduleRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleEligibilityScheduleRequest: %+v", err)
	}

	delete(decoded, "directoryScope")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleEligibilityScheduleRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleEligibilityScheduleRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleEligibilityScheduleRequest{}

func (s *UnifiedRoleEligibilityScheduleRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action                   *UnifiedRoleScheduleRequestActions `json:"action,omitempty"`
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
		TargetSchedule           *UnifiedRoleEligibilitySchedule    `json:"targetSchedule,omitempty"`
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
		return fmt.Errorf("unmarshaling UnifiedRoleEligibilityScheduleRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'UnifiedRoleEligibilityScheduleRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["directoryScope"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DirectoryScope' for 'UnifiedRoleEligibilityScheduleRequest': %+v", err)
		}
		s.DirectoryScope = &impl
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'UnifiedRoleEligibilityScheduleRequest': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
