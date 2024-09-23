package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaInfo struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. Used to uniquely identity the resource. If passed in, the prompt uri is against this resourceId as a key.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Path to the prompt that will be played. Currently supports only Wave file (.wav) format, single-channel, 16-bit
	// samples with a 16,000 (16 KHz) sampling rate.
	Uri *string `json:"uri,omitempty"`
}
