package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = OrganizationalUnit{}

type OrganizationalUnit struct {
	Children                 *[]OrganizationalUnit `json:"children,omitempty"`
	Description              nullable.Type[string] `json:"description,omitempty"`
	DisplayName              nullable.Type[string] `json:"displayName,omitempty"`
	OrganizationalUnitParent *OrganizationalUnit   `json:"organizationalUnitParent,omitempty"`
	Resources                *[]DirectoryObject    `json:"resources,omitempty"`

	// List of OData IDs for `Resources` to bind to this entity
	Resources_ODataBind *[]string `json:"resources@odata.bind,omitempty"`

	TransitiveChildren  *[]OrganizationalUnit `json:"transitiveChildren,omitempty"`
	TransitiveResources *[]DirectoryObject    `json:"transitiveResources,omitempty"`

	// List of OData IDs for `TransitiveResources` to bind to this entity
	TransitiveResources_ODataBind *[]string `json:"transitiveResources@odata.bind,omitempty"`

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

func (s OrganizationalUnit) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s OrganizationalUnit) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OrganizationalUnit{}

func (s OrganizationalUnit) MarshalJSON() ([]byte, error) {
	type wrapper OrganizationalUnit
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OrganizationalUnit: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OrganizationalUnit: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.organizationalUnit"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OrganizationalUnit: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OrganizationalUnit{}

func (s *OrganizationalUnit) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Children                      *[]OrganizationalUnit `json:"children,omitempty"`
		Description                   nullable.Type[string] `json:"description,omitempty"`
		DisplayName                   nullable.Type[string] `json:"displayName,omitempty"`
		OrganizationalUnitParent      *OrganizationalUnit   `json:"organizationalUnitParent,omitempty"`
		Resources_ODataBind           *[]string             `json:"resources@odata.bind,omitempty"`
		TransitiveChildren            *[]OrganizationalUnit `json:"transitiveChildren,omitempty"`
		TransitiveResources_ODataBind *[]string             `json:"transitiveResources@odata.bind,omitempty"`
		DeletedDateTime               nullable.Type[string] `json:"deletedDateTime,omitempty"`
		Id                            *string               `json:"id,omitempty"`
		ODataId                       *string               `json:"@odata.id,omitempty"`
		ODataType                     *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Children = decoded.Children
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.OrganizationalUnitParent = decoded.OrganizationalUnitParent
	s.Resources_ODataBind = decoded.Resources_ODataBind
	s.TransitiveChildren = decoded.TransitiveChildren
	s.TransitiveResources_ODataBind = decoded.TransitiveResources_ODataBind
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OrganizationalUnit into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Resources into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Resources' for 'OrganizationalUnit': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Resources = &output
	}

	if v, ok := temp["transitiveResources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveResources into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveResources' for 'OrganizationalUnit': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveResources = &output
	}

	return nil
}
