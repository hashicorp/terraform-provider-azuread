package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConvertIdResult struct {
	// An error object indicating the reason for the conversion failure. This value isn't present if the conversion
	// succeeded.
	ErrorDetails *GenericError `json:"errorDetails,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier that was converted. This value is the original, un-converted identifier.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// The converted identifier. This value isn't present if the conversion failed.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`
}
