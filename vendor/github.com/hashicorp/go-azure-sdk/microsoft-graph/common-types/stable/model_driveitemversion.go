package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItemVersion = DriveItemVersion{}

type DriveItemVersion struct {
	// The content stream for this version of the item.
	Content nullable.Type[string] `json:"content,omitempty"`

	// Indicates the size of the content stream for this version of the item.
	Size nullable.Type[int64] `json:"size,omitempty"`

	// Fields inherited from BaseItemVersion

	// Identity of the user which last modified the version. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the version was last modified. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Indicates the publication status of this particular version. Read-only.
	Publication *PublicationFacet `json:"publication,omitempty"`

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

func (s DriveItemVersion) BaseItemVersion() BaseBaseItemVersionImpl {
	return BaseBaseItemVersionImpl{
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Publication:          s.Publication,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s DriveItemVersion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DriveItemVersion{}

func (s DriveItemVersion) MarshalJSON() ([]byte, error) {
	type wrapper DriveItemVersion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DriveItemVersion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DriveItemVersion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.driveItemVersion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DriveItemVersion: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DriveItemVersion{}

func (s *DriveItemVersion) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Content              nullable.Type[string] `json:"content,omitempty"`
		Size                 nullable.Type[int64]  `json:"size,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Publication          *PublicationFacet     `json:"publication,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Content = decoded.Content
	s.Size = decoded.Size
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Publication = decoded.Publication

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DriveItemVersion into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'DriveItemVersion': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
