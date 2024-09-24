package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EdiscoveryTag{}

type EdiscoveryTag struct {
	// Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.
	// This value controls whether the UX presents the tags as checkboxes or a radio button group.
	ChildSelectability *EdiscoveryChildSelectability `json:"childSelectability,omitempty"`

	// Returns the tags that are a child of a tag.
	ChildTags *[]EdiscoveryTag `json:"childTags,omitempty"`

	// The user who created the tag.
	CreatedBy IdentitySet `json:"createdBy"`

	// The description for the tag.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the tag.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time the tag was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Returns the parent tag of the specified tag.
	Parent *EdiscoveryTag `json:"parent,omitempty"`

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

func (s EdiscoveryTag) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryTag{}

func (s EdiscoveryTag) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryTag
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryTag: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryTag: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.tag"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryTag: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryTag{}

func (s *EdiscoveryTag) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChildSelectability   *EdiscoveryChildSelectability `json:"childSelectability,omitempty"`
		ChildTags            *[]EdiscoveryTag              `json:"childTags,omitempty"`
		Description          nullable.Type[string]         `json:"description,omitempty"`
		DisplayName          nullable.Type[string]         `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string]         `json:"lastModifiedDateTime,omitempty"`
		Parent               *EdiscoveryTag                `json:"parent,omitempty"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChildSelectability = decoded.ChildSelectability
	s.ChildTags = decoded.ChildTags
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Parent = decoded.Parent
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryTag into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoveryTag': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
