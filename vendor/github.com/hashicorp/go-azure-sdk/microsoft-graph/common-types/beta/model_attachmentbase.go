package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentBase interface {
	Entity
	AttachmentBase() BaseAttachmentBaseImpl
}

var _ AttachmentBase = BaseAttachmentBaseImpl{}

type BaseAttachmentBaseImpl struct {
	// The MIME type.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The display name of the attachment. This doesn't need to be the actual file name.
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

func (s BaseAttachmentBaseImpl) AttachmentBase() BaseAttachmentBaseImpl {
	return s
}

func (s BaseAttachmentBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AttachmentBase = RawAttachmentBaseImpl{}

// RawAttachmentBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAttachmentBaseImpl struct {
	attachmentBase BaseAttachmentBaseImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawAttachmentBaseImpl) AttachmentBase() BaseAttachmentBaseImpl {
	return s.attachmentBase
}

func (s RawAttachmentBaseImpl) Entity() BaseEntityImpl {
	return s.attachmentBase.Entity()
}

var _ json.Marshaler = BaseAttachmentBaseImpl{}

func (s BaseAttachmentBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAttachmentBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAttachmentBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAttachmentBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attachmentBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAttachmentBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAttachmentBaseImplementation(input []byte) (AttachmentBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AttachmentBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.taskFileAttachment") {
		var out TaskFileAttachment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TaskFileAttachment: %+v", err)
		}
		return out, nil
	}

	var parent BaseAttachmentBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAttachmentBaseImpl: %+v", err)
	}

	return RawAttachmentBaseImpl{
		attachmentBase: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
