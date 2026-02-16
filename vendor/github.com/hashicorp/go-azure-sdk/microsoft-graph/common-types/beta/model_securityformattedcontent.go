package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFormattedContent struct {
	// The content of this formattedContent.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The format of the content. The possible values are: text, html, markdown, unknownFutureValue.
	Format *SecurityContentFormat `json:"format,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
