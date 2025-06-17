package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageAttachment struct {
	// The content of the attachment. If the attachment is a rich card, set the property to the rich card object. This
	// property and contentUrl are mutually exclusive.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The media type of the content attachment. The possible values are: reference: The attachment is a link to another
	// file. Populate the contentURL with the link to the object.forwardedMessageReference: The attachment is a reference to
	// a forwarded message. Populate the content with the original message context.Any contentType that is supported by the
	// Bot Framework's Attachment object.application/vnd.microsoft.card.codesnippet: Either a code snippet or place holder.
	// application/vnd.microsoft.card.announcement: An announcement header. application/vnd.microsoft.card.fluidEmbedCard: A
	// Microsoft Loop component.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// The URL for the content of the attachment.
	ContentUrl nullable.Type[string] `json:"contentUrl,omitempty"`

	// Read-only. The unique ID of the attachment.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Name of the attachment.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the Teams app that is associated with the attachment. The property is used to attribute a Teams message
	// card to the specified app.
	TeamsAppId nullable.Type[string] `json:"teamsAppId,omitempty"`

	// The URL to a thumbnail image that the channel can use if it supports using an alternative, smaller form of content or
	// contentUrl. For example, if you set contentType to application/word and set contentUrl to the location of the Word
	// document, you might include a thumbnail image that represents the document. The channel could display the thumbnail
	// image instead of the document. When the user selects the image, the channel would open the document.
	ThumbnailUrl nullable.Type[string] `json:"thumbnailUrl,omitempty"`
}

var _ json.Marshaler = ChatMessageAttachment{}

func (s ChatMessageAttachment) MarshalJSON() ([]byte, error) {
	type wrapper ChatMessageAttachment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMessageAttachment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMessageAttachment: %+v", err)
	}

	delete(decoded, "id")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMessageAttachment: %+v", err)
	}

	return encoded, nil
}
