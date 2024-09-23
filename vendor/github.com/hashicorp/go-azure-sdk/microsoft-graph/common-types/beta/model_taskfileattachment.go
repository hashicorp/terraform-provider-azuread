package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AttachmentBase = TaskFileAttachment{}

type TaskFileAttachment struct {
	// The base64-encoded contents of the file.
	ContentBytes *string `json:"contentBytes,omitempty"`

	// Fields inherited from AttachmentBase

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

func (s TaskFileAttachment) AttachmentBase() BaseAttachmentBaseImpl {
	return BaseAttachmentBaseImpl{
		ContentType:          s.ContentType,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		Size:                 s.Size,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s TaskFileAttachment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TaskFileAttachment{}

func (s TaskFileAttachment) MarshalJSON() ([]byte, error) {
	type wrapper TaskFileAttachment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TaskFileAttachment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TaskFileAttachment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.taskFileAttachment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TaskFileAttachment: %+v", err)
	}

	return encoded, nil
}
