package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationModule{}

type EducationModule struct {
	// The display name of the user that created the module.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Date time the module was created. The timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the module.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the module.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the module is pinned.
	IsPinned nullable.Type[bool] `json:"isPinned,omitempty"`

	// Specifies the language in which UI notifications for the assignment are displayed. If languageTag isn't provided, the
	// default language is en-US. Optional.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The last user that modified the module.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date time the module was last modified. The timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Learning objects that are associated with this module. Only teachers can modify this list. Nullable.
	Resources *[]EducationModuleResource `json:"resources,omitempty"`

	// Folder URL where all the file resources for this module are stored.
	ResourcesFolderUrl nullable.Type[string] `json:"resourcesFolderUrl,omitempty"`

	// Status of the module. You can't use a PATCH operation to update this value. Possible values are: draft and published.
	Status *EducationModuleStatus `json:"status,omitempty"`

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

func (s EducationModule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationModule{}

func (s EducationModule) MarshalJSON() ([]byte, error) {
	type wrapper EducationModule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationModule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationModule: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "resourcesFolderUrl")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationModule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationModule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationModule{}

func (s *EducationModule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string]      `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]      `json:"description,omitempty"`
		DisplayName          nullable.Type[string]      `json:"displayName,omitempty"`
		IsPinned             nullable.Type[bool]        `json:"isPinned,omitempty"`
		LanguageTag          nullable.Type[string]      `json:"languageTag,omitempty"`
		LastModifiedDateTime nullable.Type[string]      `json:"lastModifiedDateTime,omitempty"`
		Resources            *[]EducationModuleResource `json:"resources,omitempty"`
		ResourcesFolderUrl   nullable.Type[string]      `json:"resourcesFolderUrl,omitempty"`
		Status               *EducationModuleStatus     `json:"status,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsPinned = decoded.IsPinned
	s.LanguageTag = decoded.LanguageTag
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Resources = decoded.Resources
	s.ResourcesFolderUrl = decoded.ResourcesFolderUrl
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationModule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationModule': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationModule': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
