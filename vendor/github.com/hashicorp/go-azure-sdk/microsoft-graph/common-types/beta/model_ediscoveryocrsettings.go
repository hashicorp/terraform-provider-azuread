package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryOcrSettings struct {
	// Indicates whether or not OCR is enabled for the case.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Maximum image size that will be processed in KB).
	MaxImageSize nullable.Type[int64] `json:"maxImageSize,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The timeout duration for the OCR engine. A longer timeout may increase success of OCR, but may add to the total
	// processing time.
	Timeout nullable.Type[string] `json:"timeout,omitempty"`
}
