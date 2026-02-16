package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportItemResponse struct {
	// The version of the item.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// Data that represents an item in a base64 encoded FastTransfer stream format.
	Data nullable.Type[string] `json:"data,omitempty"`

	// An error that occurs during an action.
	Error *MailTipsError `json:"error,omitempty"`

	// The unique identifier of the item.
	ItemId nullable.Type[string] `json:"itemId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
