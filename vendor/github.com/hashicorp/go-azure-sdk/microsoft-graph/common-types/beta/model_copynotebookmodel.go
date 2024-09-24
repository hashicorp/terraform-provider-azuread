package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyNotebookModel struct {
	CreatedBy              nullable.Type[string] `json:"createdBy,omitempty"`
	CreatedByIdentity      IdentitySet           `json:"createdByIdentity"`
	CreatedTime            nullable.Type[string] `json:"createdTime,omitempty"`
	Id                     nullable.Type[string] `json:"id,omitempty"`
	IsDefault              nullable.Type[bool]   `json:"isDefault,omitempty"`
	IsShared               nullable.Type[bool]   `json:"isShared,omitempty"`
	LastModifiedBy         nullable.Type[string] `json:"lastModifiedBy,omitempty"`
	LastModifiedByIdentity IdentitySet           `json:"lastModifiedByIdentity"`
	LastModifiedTime       nullable.Type[string] `json:"lastModifiedTime,omitempty"`
	Links                  *NotebookLinks        `json:"links,omitempty"`
	Name                   nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SectionGroupsUrl nullable.Type[string] `json:"sectionGroupsUrl,omitempty"`
	SectionsUrl      nullable.Type[string] `json:"sectionsUrl,omitempty"`
	Self             nullable.Type[string] `json:"self,omitempty"`
	UserRole         *OnenoteUserRole      `json:"userRole,omitempty"`
}

var _ json.Unmarshaler = &CopyNotebookModel{}

func (s *CopyNotebookModel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedBy        nullable.Type[string] `json:"createdBy,omitempty"`
		CreatedTime      nullable.Type[string] `json:"createdTime,omitempty"`
		Id               nullable.Type[string] `json:"id,omitempty"`
		IsDefault        nullable.Type[bool]   `json:"isDefault,omitempty"`
		IsShared         nullable.Type[bool]   `json:"isShared,omitempty"`
		LastModifiedBy   nullable.Type[string] `json:"lastModifiedBy,omitempty"`
		LastModifiedTime nullable.Type[string] `json:"lastModifiedTime,omitempty"`
		Links            *NotebookLinks        `json:"links,omitempty"`
		Name             nullable.Type[string] `json:"name,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		SectionGroupsUrl nullable.Type[string] `json:"sectionGroupsUrl,omitempty"`
		SectionsUrl      nullable.Type[string] `json:"sectionsUrl,omitempty"`
		Self             nullable.Type[string] `json:"self,omitempty"`
		UserRole         *OnenoteUserRole      `json:"userRole,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedBy = decoded.CreatedBy
	s.CreatedTime = decoded.CreatedTime
	s.Id = decoded.Id
	s.IsDefault = decoded.IsDefault
	s.IsShared = decoded.IsShared
	s.LastModifiedBy = decoded.LastModifiedBy
	s.LastModifiedTime = decoded.LastModifiedTime
	s.Links = decoded.Links
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SectionGroupsUrl = decoded.SectionGroupsUrl
	s.SectionsUrl = decoded.SectionsUrl
	s.Self = decoded.Self
	s.UserRole = decoded.UserRole

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CopyNotebookModel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdByIdentity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedByIdentity' for 'CopyNotebookModel': %+v", err)
		}
		s.CreatedByIdentity = impl
	}

	if v, ok := temp["lastModifiedByIdentity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedByIdentity' for 'CopyNotebookModel': %+v", err)
		}
		s.LastModifiedByIdentity = impl
	}

	return nil
}
