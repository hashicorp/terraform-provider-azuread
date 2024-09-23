package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaStream struct {
	Direction *MediaDirection `json:"direction,omitempty"`

	// The media stream label.
	Label nullable.Type[string] `json:"label,omitempty"`

	MediaType *Modality `json:"mediaType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the server has muted the media.
	ServerMuted *bool `json:"serverMuted,omitempty"`

	// The source ID.
	SourceId *string `json:"sourceId,omitempty"`
}
