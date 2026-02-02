package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnenoteEntityHierarchyModel = SectionGroup{}

type SectionGroup struct {
	// The notebook that contains the section group. Read-only.
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`

	// The section group that contains the section group. Read-only.
	ParentSectionGroup *SectionGroup `json:"parentSectionGroup,omitempty"`

	// The section groups in the section. Read-only. Nullable.
	SectionGroups *[]SectionGroup `json:"sectionGroups,omitempty"`

	// The URL for the sectionGroups navigation property, which returns all the section groups in the section group.
	// Read-only.
	SectionGroupsUrl nullable.Type[string] `json:"sectionGroupsUrl,omitempty"`

	// The sections in the section group. Read-only. Nullable.
	Sections *[]OnenoteSection `json:"sections,omitempty"`

	// The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
	SectionsUrl nullable.Type[string] `json:"sectionsUrl,omitempty"`

	// Fields inherited from OnenoteEntityHierarchyModel

	CreatedBy            IdentitySet           `json:"createdBy"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from OnenoteEntitySchemaObjectModel

	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Fields inherited from OnenoteEntityBaseModel

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

func (s SectionGroup) OnenoteEntityHierarchyModel() BaseOnenoteEntityHierarchyModelImpl {
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

func (s SectionGroup) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return BaseOnenoteEntitySchemaObjectModelImpl{
		CreatedDateTime: s.CreatedDateTime,
		Self:            s.Self,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SectionGroup) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return BaseOnenoteEntityBaseModelImpl{
		Self:      s.Self,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SectionGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SectionGroup{}

func (s SectionGroup) MarshalJSON() ([]byte, error) {
	type wrapper SectionGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SectionGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SectionGroup: %+v", err)
	}

	delete(decoded, "parentNotebook")
	delete(decoded, "parentSectionGroup")
	delete(decoded, "sectionGroups")
	delete(decoded, "sectionGroupsUrl")
	delete(decoded, "sections")
	delete(decoded, "sectionsUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sectionGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SectionGroup: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SectionGroup{}

func (s *SectionGroup) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ParentNotebook       *Notebook             `json:"parentNotebook,omitempty"`
		ParentSectionGroup   *SectionGroup         `json:"parentSectionGroup,omitempty"`
		SectionGroups        *[]SectionGroup       `json:"sectionGroups,omitempty"`
		SectionGroupsUrl     nullable.Type[string] `json:"sectionGroupsUrl,omitempty"`
		Sections             *[]OnenoteSection     `json:"sections,omitempty"`
		SectionsUrl          nullable.Type[string] `json:"sectionsUrl,omitempty"`
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

	s.ParentNotebook = decoded.ParentNotebook
	s.ParentSectionGroup = decoded.ParentSectionGroup
	s.SectionGroups = decoded.SectionGroups
	s.SectionGroupsUrl = decoded.SectionGroupsUrl
	s.Sections = decoded.Sections
	s.SectionsUrl = decoded.SectionsUrl
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Self = decoded.Self

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SectionGroup into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SectionGroup': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SectionGroup': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
