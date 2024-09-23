package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Attachment interface {
	Entity
	Attachment() BaseAttachmentImpl
}

var _ Attachment = BaseAttachmentImpl{}

type BaseAttachmentImpl struct {
	// The MIME type.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// true if the attachment is an inline attachment; otherwise, false.
	IsInline *bool `json:"isInline,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The display name of the attachment. This does not need to be the actual file name.
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

func (s BaseAttachmentImpl) Attachment() BaseAttachmentImpl {
	return s
}

func (s BaseAttachmentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Attachment = RawAttachmentImpl{}

// RawAttachmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAttachmentImpl struct {
	attachment BaseAttachmentImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawAttachmentImpl) Attachment() BaseAttachmentImpl {
	return s.attachment
}

func (s RawAttachmentImpl) Entity() BaseEntityImpl {
	return s.attachment.Entity()
}

var _ json.Marshaler = BaseAttachmentImpl{}

func (s BaseAttachmentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAttachmentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAttachmentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAttachmentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attachment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAttachmentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAttachmentImplementation(input []byte) (Attachment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Attachment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.fileAttachment") {
		var out FileAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileAttachment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.itemAttachment") {
		var out ItemAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemAttachment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.referenceAttachment") {
		var out ReferenceAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ReferenceAttachment: %+v", err)
		}
		return out, nil
	}

	var parent BaseAttachmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAttachmentImpl: %+v", err)
	}

	return RawAttachmentImpl{
		attachment: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
