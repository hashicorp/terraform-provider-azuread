package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleScheduleInstanceBase = UnifiedRoleEligibilityScheduleInstance{}

type UnifiedRoleEligibilityScheduleInstance struct {
	// Time that the roleEligibilityScheduleInstance will expire.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Membership type of the assignment. It can either be Inherited, Direct, or Group.
	MemberType nullable.Type[string] `json:"memberType,omitempty"`

	// Identifier of the parent roleEligibilitySchedule for this instance.
	RoleEligibilityScheduleId nullable.Type[string] `json:"roleEligibilityScheduleId,omitempty"`

	// Time that the roleEligibilityScheduleInstance will start.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Fields inherited from UnifiedRoleScheduleInstanceBase

	// Read-only property with details of the app-specific scope when the assignment or role eligibility is scoped to an
	// app. Nullable.
	AppScope *AppScope `json:"appScope,omitempty"`

	// Identifier of the app-specific scope when the assignment or role eligibility is scoped to an app. The scope of an
	// assignment or role eligibility determines the set of resources for which the principal has been granted access. App
	// scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use
	// directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
	AppScopeId nullable.Type[string] `json:"appScopeId,omitempty"`

	// The directory object that is the scope of the assignment or role eligibility. Read-only.
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`

	// Identifier of the directory object representing the scope of the assignment or role eligibility. The scope of an
	// assignment or role eligibility determines the set of resources for which the principal has been granted access.
	// Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for
	// tenant-wide scope. Use appScopeId to limit the scope to an application only.
	DirectoryScopeId nullable.Type[string] `json:"directoryScopeId,omitempty"`

	// OData ID for `DirectoryScope` to bind to this entity
	DirectoryScope_ODataBind *string `json:"directoryScope@odata.bind,omitempty"`

	// The principal that's getting a role assignment or role eligibility through the request.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// Identifier of the principal that has been granted the role assignment or that's eligible for a role.
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`

	// Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that the principal is
	// eligible for.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

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

func (s UnifiedRoleEligibilityScheduleInstance) UnifiedRoleScheduleInstanceBase() BaseUnifiedRoleScheduleInstanceBaseImpl {
	return BaseUnifiedRoleScheduleInstanceBaseImpl{
		AppScope:                 s.AppScope,
		AppScopeId:               s.AppScopeId,
		DirectoryScope:           s.DirectoryScope,
		DirectoryScopeId:         s.DirectoryScopeId,
		DirectoryScope_ODataBind: s.DirectoryScope_ODataBind,
		Principal:                s.Principal,
		PrincipalId:              s.PrincipalId,
		Principal_ODataBind:      s.Principal_ODataBind,
		RoleDefinition:           s.RoleDefinition,
		RoleDefinitionId:         s.RoleDefinitionId,
		Id:                       s.Id,
		ODataId:                  s.ODataId,
		ODataType:                s.ODataType,
	}
}

func (s UnifiedRoleEligibilityScheduleInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleEligibilityScheduleInstance{}

func (s UnifiedRoleEligibilityScheduleInstance) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleEligibilityScheduleInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleEligibilityScheduleInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleEligibilityScheduleInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleEligibilityScheduleInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleEligibilityScheduleInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleEligibilityScheduleInstance{}

func (s *UnifiedRoleEligibilityScheduleInstance) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EndDateTime               nullable.Type[string]  `json:"endDateTime,omitempty"`
		MemberType                nullable.Type[string]  `json:"memberType,omitempty"`
		RoleEligibilityScheduleId nullable.Type[string]  `json:"roleEligibilityScheduleId,omitempty"`
		StartDateTime             nullable.Type[string]  `json:"startDateTime,omitempty"`
		AppScopeId                nullable.Type[string]  `json:"appScopeId,omitempty"`
		DirectoryScopeId          nullable.Type[string]  `json:"directoryScopeId,omitempty"`
		DirectoryScope_ODataBind  *string                `json:"directoryScope@odata.bind,omitempty"`
		PrincipalId               nullable.Type[string]  `json:"principalId,omitempty"`
		Principal_ODataBind       *string                `json:"principal@odata.bind,omitempty"`
		RoleDefinition            *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
		RoleDefinitionId          nullable.Type[string]  `json:"roleDefinitionId,omitempty"`
		Id                        *string                `json:"id,omitempty"`
		ODataId                   *string                `json:"@odata.id,omitempty"`
		ODataType                 *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EndDateTime = decoded.EndDateTime
	s.MemberType = decoded.MemberType
	s.RoleEligibilityScheduleId = decoded.RoleEligibilityScheduleId
	s.StartDateTime = decoded.StartDateTime
	s.AppScopeId = decoded.AppScopeId
	s.DirectoryScopeId = decoded.DirectoryScopeId
	s.DirectoryScope_ODataBind = decoded.DirectoryScope_ODataBind
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.RoleDefinition = decoded.RoleDefinition
	s.RoleDefinitionId = decoded.RoleDefinitionId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UnifiedRoleEligibilityScheduleInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appScope"]; ok {
		impl, err := UnmarshalAppScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppScope' for 'UnifiedRoleEligibilityScheduleInstance': %+v", err)
		}
		s.AppScope = &impl
	}

	if v, ok := temp["directoryScope"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DirectoryScope' for 'UnifiedRoleEligibilityScheduleInstance': %+v", err)
		}
		s.DirectoryScope = &impl
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'UnifiedRoleEligibilityScheduleInstance': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
