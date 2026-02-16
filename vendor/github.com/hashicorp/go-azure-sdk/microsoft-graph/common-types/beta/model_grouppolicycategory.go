package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicyCategory{}

type GroupPolicyCategory struct {
	// The children categories
	Children *[]GroupPolicyCategory `json:"children,omitempty"`

	// The id of the definition file the category came from
	DefinitionFile *GroupPolicyDefinitionFile `json:"definitionFile,omitempty"`

	// The immediate GroupPolicyDefinition children of the category
	Definitions *[]GroupPolicyDefinition `json:"definitions,omitempty"`

	// The string id of the category's display name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Category Ingestion source
	IngestionSource *IngestionSource `json:"ingestionSource,omitempty"`

	// Defines if the category is a root category
	IsRoot *bool `json:"isRoot,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The parent category
	Parent *GroupPolicyCategory `json:"parent,omitempty"`

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

func (s GroupPolicyCategory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyCategory{}

func (s GroupPolicyCategory) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyCategory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyCategory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyCategory: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GroupPolicyCategory{}

func (s *GroupPolicyCategory) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Children             *[]GroupPolicyCategory   `json:"children,omitempty"`
		Definitions          *[]GroupPolicyDefinition `json:"definitions,omitempty"`
		DisplayName          nullable.Type[string]    `json:"displayName,omitempty"`
		IngestionSource      *IngestionSource         `json:"ingestionSource,omitempty"`
		IsRoot               *bool                    `json:"isRoot,omitempty"`
		LastModifiedDateTime *string                  `json:"lastModifiedDateTime,omitempty"`
		Parent               *GroupPolicyCategory     `json:"parent,omitempty"`
		Id                   *string                  `json:"id,omitempty"`
		ODataId              *string                  `json:"@odata.id,omitempty"`
		ODataType            *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Children = decoded.Children
	s.Definitions = decoded.Definitions
	s.DisplayName = decoded.DisplayName
	s.IngestionSource = decoded.IngestionSource
	s.IsRoot = decoded.IsRoot
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Parent = decoded.Parent
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GroupPolicyCategory into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["definitionFile"]; ok {
		impl, err := UnmarshalGroupPolicyDefinitionFileImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefinitionFile' for 'GroupPolicyCategory': %+v", err)
		}
		s.DefinitionFile = &impl
	}

	return nil
}
