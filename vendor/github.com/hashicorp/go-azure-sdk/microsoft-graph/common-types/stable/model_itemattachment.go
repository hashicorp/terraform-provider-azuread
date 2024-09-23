package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Attachment = ItemAttachment{}

type ItemAttachment struct {
	// The attached message or event. Navigation property.
	Item *OutlookItem `json:"item,omitempty"`

	// Fields inherited from Attachment

	// The MIME type.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// true if the attachment is an inline attachment; otherwise, false.
	IsInline *bool `json:"isInline,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The attachment's file name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The length of the attachment in bytes.
	Size *int64 `json:"size,omitempty"`

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

func (s ItemAttachment) Attachment() BaseAttachmentImpl {
	return BaseAttachmentImpl{
		ContentType:          s.ContentType,
		IsInline:             s.IsInline,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		Size:                 s.Size,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ItemAttachment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemAttachment{}

func (s ItemAttachment) MarshalJSON() ([]byte, error) {
	type wrapper ItemAttachment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemAttachment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemAttachment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemAttachment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemAttachment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ItemAttachment{}

func (s *ItemAttachment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContentType          nullable.Type[string] `json:"contentType,omitempty"`
		IsInline             *bool                 `json:"isInline,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Name                 nullable.Type[string] `json:"name,omitempty"`
		Size                 *int64                `json:"size,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ContentType = decoded.ContentType
	s.Id = decoded.Id
	s.IsInline = decoded.IsInline
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Size = decoded.Size

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ItemAttachment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["item"]; ok {
		impl, err := UnmarshalOutlookItemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Item' for 'ItemAttachment': %+v", err)
		}
		s.Item = &impl
	}

	return nil
}
