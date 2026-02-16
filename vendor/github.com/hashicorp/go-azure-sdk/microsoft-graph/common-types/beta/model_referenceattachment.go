package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Attachment = ReferenceAttachment{}

type ReferenceAttachment struct {
	// Specifies whether the attachment is a link to a folder. You must set this property to true if sourceUrl is a link to
	// a folder. Optional.
	IsFolder nullable.Type[bool] `json:"isFolder,omitempty"`

	// Specifies the permissions granted for the attachment by the type of provider in providerType. Possible values are:
	// other, view, edit, anonymousView, anonymousEdit, organizationView, organizationEdit. Optional.
	Permission *ReferenceAttachmentPermission `json:"permission,omitempty"`

	// Applies to only a reference attachment of an image - URL to get a preview image. Use thumbnailUrl and previewUrl only
	// when sourceUrl identifies an image file. Optional.
	PreviewUrl nullable.Type[string] `json:"previewUrl,omitempty"`

	// The type of provider that supports an attachment of this contentType. Possible values are: other, oneDriveBusiness,
	// oneDriveConsumer, dropbox. Optional.
	ProviderType *ReferenceAttachmentProvider `json:"providerType,omitempty"`

	// URL to get the attachment content. If this value is a URL to a folder, then for the folder to be displayed correctly
	// in Outlook or Outlook on the web, set isFolder to true. Required.
	SourceUrl nullable.Type[string] `json:"sourceUrl,omitempty"`

	// Applies to only a reference attachment of an image - URL to get a thumbnail image. Use thumbnailUrl and previewUrl
	// only when sourceUrl identifies an image file. Optional.
	ThumbnailUrl nullable.Type[string] `json:"thumbnailUrl,omitempty"`

	// Fields inherited from Attachment

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

func (s ReferenceAttachment) Attachment() BaseAttachmentImpl {
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

func (s ReferenceAttachment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ReferenceAttachment{}

func (s ReferenceAttachment) MarshalJSON() ([]byte, error) {
	type wrapper ReferenceAttachment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ReferenceAttachment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ReferenceAttachment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.referenceAttachment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ReferenceAttachment: %+v", err)
	}

	return encoded, nil
}
