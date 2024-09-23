package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = DirectoryRole{}

type DirectoryRole struct {
	// The description for the directory role. Read-only. Supports $filter (eq), $search, $select.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the directory role. Read-only. Supports $filter (eq), $search, $select.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports
	// $expand.
	Members *[]DirectoryObject `json:"members,omitempty"`

	// List of OData IDs for `Members` to bind to this entity
	Members_ODataBind *[]string `json:"members@odata.bind,omitempty"`

	// The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a
	// directory role in a tenant with a POST operation. After the directory role has been activated, the property is read
	// only. Supports $filter (eq), $select.
	RoleTemplateId nullable.Type[string] `json:"roleTemplateId,omitempty"`

	// Members of this directory role that are scoped to administrative units. Read-only. Nullable.
	ScopedMembers *[]ScopedRoleMembership `json:"scopedMembers,omitempty"`

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

func (s DirectoryRole) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s DirectoryRole) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DirectoryRole{}

func (s DirectoryRole) MarshalJSON() ([]byte, error) {
	type wrapper DirectoryRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DirectoryRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectoryRole: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "members")
	delete(decoded, "scopedMembers")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directoryRole"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DirectoryRole: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DirectoryRole{}

func (s *DirectoryRole) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description       nullable.Type[string]   `json:"description,omitempty"`
		DisplayName       nullable.Type[string]   `json:"displayName,omitempty"`
		Members_ODataBind *[]string               `json:"members@odata.bind,omitempty"`
		RoleTemplateId    nullable.Type[string]   `json:"roleTemplateId,omitempty"`
		ScopedMembers     *[]ScopedRoleMembership `json:"scopedMembers,omitempty"`
		DeletedDateTime   nullable.Type[string]   `json:"deletedDateTime,omitempty"`
		Id                *string                 `json:"id,omitempty"`
		ODataId           *string                 `json:"@odata.id,omitempty"`
		ODataType         *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Members_ODataBind = decoded.Members_ODataBind
	s.RoleTemplateId = decoded.RoleTemplateId
	s.ScopedMembers = decoded.ScopedMembers
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DirectoryRole into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'DirectoryRole': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	return nil
}
