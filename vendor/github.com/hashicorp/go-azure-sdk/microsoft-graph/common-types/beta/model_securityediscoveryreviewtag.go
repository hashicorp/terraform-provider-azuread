package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityTag = SecurityEdiscoveryReviewTag{}

type SecurityEdiscoveryReviewTag struct {
	// Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.
	// This value controls whether the UX presents the tags as checkboxes or a radio button group.
	ChildSelectability *SecurityChildSelectability `json:"childSelectability,omitempty"`

	// Returns the tags that are a child of a tag.
	ChildTags *[]SecurityEdiscoveryReviewTag `json:"childTags,omitempty"`

	// Returns the parent tag of the specified tag.
	Parent *SecurityEdiscoveryReviewTag `json:"parent,omitempty"`

	// Fields inherited from SecurityTag

	CreatedBy            IdentitySet           `json:"createdBy"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s SecurityEdiscoveryReviewTag) SecurityTag() BaseSecurityTagImpl {
	return BaseSecurityTagImpl{
		CreatedBy:            s.CreatedBy,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryReviewTag) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryReviewTag{}

func (s SecurityEdiscoveryReviewTag) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryReviewTag
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryReviewTag: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryReviewTag: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryReviewTag"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryReviewTag: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryReviewTag{}

func (s *SecurityEdiscoveryReviewTag) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChildSelectability   *SecurityChildSelectability    `json:"childSelectability,omitempty"`
		ChildTags            *[]SecurityEdiscoveryReviewTag `json:"childTags,omitempty"`
		Parent               *SecurityEdiscoveryReviewTag   `json:"parent,omitempty"`
		Description          nullable.Type[string]          `json:"description,omitempty"`
		DisplayName          nullable.Type[string]          `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string]          `json:"lastModifiedDateTime,omitempty"`
		Id                   *string                        `json:"id,omitempty"`
		ODataId              *string                        `json:"@odata.id,omitempty"`
		ODataType            *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChildSelectability = decoded.ChildSelectability
	s.ChildTags = decoded.ChildTags
	s.Parent = decoded.Parent
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryReviewTag into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryReviewTag': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
