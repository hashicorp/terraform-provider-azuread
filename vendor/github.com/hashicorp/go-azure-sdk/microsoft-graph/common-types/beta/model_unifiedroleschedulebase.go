package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleScheduleBase interface {
	Entity
	UnifiedRoleScheduleBase() BaseUnifiedRoleScheduleBaseImpl
}

var _ UnifiedRoleScheduleBase = BaseUnifiedRoleScheduleBaseImpl{}

type BaseUnifiedRoleScheduleBaseImpl struct {
	// Read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an
	// app. Nullable.
	AppScope *AppScope `json:"appScope,omitempty"`

	// Identifier of the app-specific scope when the assignment or eligibility is scoped to an app. The scope of an
	// assignment or eligibility determines the set of resources for which the principal has been granted access. App scopes
	// are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use
	// directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
	AppScopeId nullable.Type[string] `json:"appScopeId,omitempty"`

	// When the schedule was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Identifier of the object through which this schedule was created.
	CreatedUsing nullable.Type[string] `json:"createdUsing,omitempty"`

	// The directory object that is the scope of the role eligibility or assignment. Read-only.
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`

	// Identifier of the directory object representing the scope of the assignment or eligibility. The scope of an
	// assignment or eligibility determines the set of resources for which the principal has been granted access. Directory
	// scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide
	// scope. Use appScopeId to limit the scope to an application only.
	DirectoryScopeId nullable.Type[string] `json:"directoryScopeId,omitempty"`

	// OData ID for `DirectoryScope` to bind to this entity
	DirectoryScope_ODataBind *string `json:"directoryScope@odata.bind,omitempty"`

	// When the schedule was last modified.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// The principal that's getting a role assignment or that's eligible for a role through the request.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// Identifier of the principal that has been granted the role assignment or eligibility.
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`

	// Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that a principal is
	// eligible for.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The status of the role assignment or eligibility request.
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

func (s BaseUnifiedRoleScheduleBaseImpl) UnifiedRoleScheduleBase() BaseUnifiedRoleScheduleBaseImpl {
	return s
}

func (s BaseUnifiedRoleScheduleBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ UnifiedRoleScheduleBase = RawUnifiedRoleScheduleBaseImpl{}

// RawUnifiedRoleScheduleBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUnifiedRoleScheduleBaseImpl struct {
	unifiedRoleScheduleBase BaseUnifiedRoleScheduleBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawUnifiedRoleScheduleBaseImpl) UnifiedRoleScheduleBase() BaseUnifiedRoleScheduleBaseImpl {
	return s.unifiedRoleScheduleBase
}

func (s RawUnifiedRoleScheduleBaseImpl) Entity() BaseEntityImpl {
	return s.unifiedRoleScheduleBase.Entity()
}

var _ json.Marshaler = BaseUnifiedRoleScheduleBaseImpl{}

func (s BaseUnifiedRoleScheduleBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseUnifiedRoleScheduleBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseUnifiedRoleScheduleBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseUnifiedRoleScheduleBaseImpl: %+v", err)
	}

	delete(decoded, "directoryScope")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleScheduleBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseUnifiedRoleScheduleBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseUnifiedRoleScheduleBaseImpl{}

func (s *BaseUnifiedRoleScheduleBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppScopeId               nullable.Type[string]  `json:"appScopeId,omitempty"`
		CreatedDateTime          nullable.Type[string]  `json:"createdDateTime,omitempty"`
		CreatedUsing             nullable.Type[string]  `json:"createdUsing,omitempty"`
		DirectoryScopeId         nullable.Type[string]  `json:"directoryScopeId,omitempty"`
		DirectoryScope_ODataBind *string                `json:"directoryScope@odata.bind,omitempty"`
		ModifiedDateTime         nullable.Type[string]  `json:"modifiedDateTime,omitempty"`
		PrincipalId              nullable.Type[string]  `json:"principalId,omitempty"`
		Principal_ODataBind      *string                `json:"principal@odata.bind,omitempty"`
		RoleDefinition           *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
		RoleDefinitionId         nullable.Type[string]  `json:"roleDefinitionId,omitempty"`
		Status                   nullable.Type[string]  `json:"status,omitempty"`
		Id                       *string                `json:"id,omitempty"`
		ODataId                  *string                `json:"@odata.id,omitempty"`
		ODataType                *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppScopeId = decoded.AppScopeId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CreatedUsing = decoded.CreatedUsing
	s.DirectoryScopeId = decoded.DirectoryScopeId
	s.DirectoryScope_ODataBind = decoded.DirectoryScope_ODataBind
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.RoleDefinition = decoded.RoleDefinition
	s.RoleDefinitionId = decoded.RoleDefinitionId
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseUnifiedRoleScheduleBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appScope"]; ok {
		impl, err := UnmarshalAppScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppScope' for 'BaseUnifiedRoleScheduleBaseImpl': %+v", err)
		}
		s.AppScope = &impl
	}

	if v, ok := temp["directoryScope"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DirectoryScope' for 'BaseUnifiedRoleScheduleBaseImpl': %+v", err)
		}
		s.DirectoryScope = &impl
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'BaseUnifiedRoleScheduleBaseImpl': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}

func UnmarshalUnifiedRoleScheduleBaseImplementation(input []byte) (UnifiedRoleScheduleBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleScheduleBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleAssignmentSchedule") {
		var out UnifiedRoleAssignmentSchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleAssignmentSchedule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleEligibilitySchedule") {
		var out UnifiedRoleEligibilitySchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleEligibilitySchedule: %+v", err)
		}
		return out, nil
	}

	var parent BaseUnifiedRoleScheduleBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUnifiedRoleScheduleBaseImpl: %+v", err)
	}

	return RawUnifiedRoleScheduleBaseImpl{
		unifiedRoleScheduleBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
