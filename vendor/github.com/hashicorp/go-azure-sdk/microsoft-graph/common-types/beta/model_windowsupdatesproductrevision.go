package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesProductRevision{}

type WindowsUpdatesProductRevision struct {
	CatalogEntry *WindowsUpdatesCatalogEntry `json:"catalogEntry,omitempty"`

	// The display name of the content. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	// True indicates that the content is hotpatchable; otherwise, false. For more information, see Deploy a hotpatch
	// quality update using Windows Autopatch. Read-only.
	IsHotpatchUpdate *bool `json:"isHotpatchUpdate,omitempty"`

	// The knowledge base article associated with the product revision.
	KnowledgeBaseArticle *WindowsUpdatesKnowledgeBaseArticle `json:"knowledgeBaseArticle,omitempty"`

	OsBuild *WindowsUpdatesBuildVersionDetails `json:"osBuild,omitempty"`

	// The product of the revision. Possible values are: Windows 10, Windows 11. Read-only.
	Product *string `json:"product,omitempty"`

	// The release date for the content. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ReleaseDateTime *string `json:"releaseDateTime,omitempty"`

	// The version of the feature update. Read-only.
	Version *string `json:"version,omitempty"`

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

func (s WindowsUpdatesProductRevision) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesProductRevision{}

func (s WindowsUpdatesProductRevision) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesProductRevision
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesProductRevision: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesProductRevision: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "isHotpatchUpdate")
	delete(decoded, "product")
	delete(decoded, "releaseDateTime")
	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.productRevision"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesProductRevision: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesProductRevision{}

func (s *WindowsUpdatesProductRevision) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName          *string                             `json:"displayName,omitempty"`
		IsHotpatchUpdate     *bool                               `json:"isHotpatchUpdate,omitempty"`
		KnowledgeBaseArticle *WindowsUpdatesKnowledgeBaseArticle `json:"knowledgeBaseArticle,omitempty"`
		OsBuild              *WindowsUpdatesBuildVersionDetails  `json:"osBuild,omitempty"`
		Product              *string                             `json:"product,omitempty"`
		ReleaseDateTime      *string                             `json:"releaseDateTime,omitempty"`
		Version              *string                             `json:"version,omitempty"`
		Id                   *string                             `json:"id,omitempty"`
		ODataId              *string                             `json:"@odata.id,omitempty"`
		ODataType            *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.IsHotpatchUpdate = decoded.IsHotpatchUpdate
	s.KnowledgeBaseArticle = decoded.KnowledgeBaseArticle
	s.OsBuild = decoded.OsBuild
	s.Product = decoded.Product
	s.ReleaseDateTime = decoded.ReleaseDateTime
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesProductRevision into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["catalogEntry"]; ok {
		impl, err := UnmarshalWindowsUpdatesCatalogEntryImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CatalogEntry' for 'WindowsUpdatesProductRevision': %+v", err)
		}
		s.CatalogEntry = &impl
	}

	return nil
}
