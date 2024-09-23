package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = AdministrativeUnit{}

type AdministrativeUnit struct {
	Description nullable.Type[string] `json:"description,omitempty"`
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The collection of open extensions defined for this administrative unit. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	IsMemberManagementRestricted nullable.Type[bool] `json:"isMemberManagementRestricted,omitempty"`

	// Users and groups that are members of this administrative unit. Supports $expand.
	Members *[]DirectoryObject `json:"members,omitempty"`

	// List of OData IDs for `Members` to bind to this entity
	Members_ODataBind *[]string `json:"members@odata.bind,omitempty"`

	MembershipRule                nullable.Type[string] `json:"membershipRule,omitempty"`
	MembershipRuleProcessingState nullable.Type[string] `json:"membershipRuleProcessingState,omitempty"`
	MembershipType                nullable.Type[string] `json:"membershipType,omitempty"`

	// Scoped-role members of this administrative unit.
	ScopedRoleMembers *[]ScopedRoleMembership `json:"scopedRoleMembers,omitempty"`

	Visibility nullable.Type[string] `json:"visibility,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s AdministrativeUnit) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AdministrativeUnit) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AdministrativeUnit{}

func (s AdministrativeUnit) MarshalJSON() ([]byte, error) {
	type wrapper AdministrativeUnit
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AdministrativeUnit: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AdministrativeUnit: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.administrativeUnit"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AdministrativeUnit: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AdministrativeUnit{}

func (s *AdministrativeUnit) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description                   nullable.Type[string]   `json:"description,omitempty"`
		DisplayName                   nullable.Type[string]   `json:"displayName,omitempty"`
		IsMemberManagementRestricted  nullable.Type[bool]     `json:"isMemberManagementRestricted,omitempty"`
		Members_ODataBind             *[]string               `json:"members@odata.bind,omitempty"`
		MembershipRule                nullable.Type[string]   `json:"membershipRule,omitempty"`
		MembershipRuleProcessingState nullable.Type[string]   `json:"membershipRuleProcessingState,omitempty"`
		MembershipType                nullable.Type[string]   `json:"membershipType,omitempty"`
		ScopedRoleMembers             *[]ScopedRoleMembership `json:"scopedRoleMembers,omitempty"`
		Visibility                    nullable.Type[string]   `json:"visibility,omitempty"`
		DeletedDateTime               nullable.Type[string]   `json:"deletedDateTime,omitempty"`
		Id                            *string                 `json:"id,omitempty"`
		ODataId                       *string                 `json:"@odata.id,omitempty"`
		ODataType                     *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsMemberManagementRestricted = decoded.IsMemberManagementRestricted
	s.Members_ODataBind = decoded.Members_ODataBind
	s.MembershipRule = decoded.MembershipRule
	s.MembershipRuleProcessingState = decoded.MembershipRuleProcessingState
	s.MembershipType = decoded.MembershipType
	s.ScopedRoleMembers = decoded.ScopedRoleMembers
	s.Visibility = decoded.Visibility
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AdministrativeUnit into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'AdministrativeUnit': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["members"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Members into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'AdministrativeUnit': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	return nil
}
