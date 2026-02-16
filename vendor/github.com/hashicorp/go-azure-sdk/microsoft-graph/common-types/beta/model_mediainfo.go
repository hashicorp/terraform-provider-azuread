package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaInfo struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional, used to uniquely identity the resource. If passed, the prompt uri is cached against this resourceId as key.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Path to the prompt to be played. Currently only Wave file (.wav) format, single-channel, 16-bit samples with a 16,000
	// (16 KHz) sampling rate is only supported.
	Uri *string `json:"uri,omitempty"`
}
