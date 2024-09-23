package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnenoteEntityHierarchyModel = Notebook{}

type Notebook struct {
	// Indicates whether this is the user's default notebook. Read-only.
	IsDefault nullable.Type[bool] `json:"isDefault,omitempty"`

	// Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the
	// owner. Read-only.
	IsShared nullable.Type[bool] `json:"isShared,omitempty"`

	// Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's
	// installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
	Links *NotebookLinks `json:"links,omitempty"`

	// The section groups in the notebook. Read-only. Nullable.
	SectionGroups *[]SectionGroup `json:"sectionGroups,omitempty"`

	// The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
	SectionGroupsUrl nullable.Type[string] `json:"sectionGroupsUrl,omitempty"`

	// The sections in the notebook. Read-only. Nullable.
	Sections *[]OnenoteSection `json:"sections,omitempty"`

	// The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
	SectionsUrl nullable.Type[string] `json:"sectionsUrl,omitempty"`

	// Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook.
	// Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook.
	// Read-only.
	UserRole *OnenoteUserRole `json:"userRole,omitempty"`

	// Fields inherited from OnenoteEntityHierarchyModel

	// Identity of the user, device, and application that created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The name of the notebook.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Identity of the user, device, and application that created the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The date and time when the notebook was last modified. The timestamp represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from OnenoteEntitySchemaObjectModel

	// The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Fields inherited from OnenoteEntityBaseModel

	// The endpoint where you can get details about the page. Read-only.
	Self nullable.Type[string] `json:"self,omitempty"`

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

func (s Notebook) OnenoteEntityHierarchyModel() BaseOnenoteEntityHierarchyModelImpl {
	return BaseOnenoteEntityHierarchyModelImpl{
		CreatedBy:            s.CreatedBy,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		CreatedDateTime:      s.CreatedDateTime,
		Self:                 s.Self,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Notebook) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return BaseOnenoteEntitySchemaObjectModelImpl{
		CreatedDateTime: s.CreatedDateTime,
		Self:            s.Self,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Notebook) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return BaseOnenoteEntityBaseModelImpl{
		Self:      s.Self,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s Notebook) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Notebook{}

func (s Notebook) MarshalJSON() ([]byte, error) {
	type wrapper Notebook
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Notebook: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Notebook: %+v", err)
	}

	delete(decoded, "isDefault")
	delete(decoded, "isShared")
	delete(decoded, "sectionGroups")
	delete(decoded, "sectionGroupsUrl")
	delete(decoded, "sections")
	delete(decoded, "sectionsUrl")
	delete(decoded, "userRole")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.notebook"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Notebook: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Notebook{}

func (s *Notebook) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsDefault            nullable.Type[bool]   `json:"isDefault,omitempty"`
		IsShared             nullable.Type[bool]   `json:"isShared,omitempty"`
		Links                *NotebookLinks        `json:"links,omitempty"`
		SectionGroups        *[]SectionGroup       `json:"sectionGroups,omitempty"`
		SectionGroupsUrl     nullable.Type[string] `json:"sectionGroupsUrl,omitempty"`
		Sections             *[]OnenoteSection     `json:"sections,omitempty"`
		SectionsUrl          nullable.Type[string] `json:"sectionsUrl,omitempty"`
		UserRole             *OnenoteUserRole      `json:"userRole,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Self                 nullable.Type[string] `json:"self,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsDefault = decoded.IsDefault
	s.IsShared = decoded.IsShared
	s.Links = decoded.Links
	s.SectionGroups = decoded.SectionGroups
	s.SectionGroupsUrl = decoded.SectionGroupsUrl
	s.Sections = decoded.Sections
	s.SectionsUrl = decoded.SectionsUrl
	s.UserRole = decoded.UserRole
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Self = decoded.Self

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Notebook into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'Notebook': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'Notebook': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
