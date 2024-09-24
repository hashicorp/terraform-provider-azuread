package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JournalLine struct {
	Account                *Account               `json:"account,omitempty"`
	AccountId              nullable.Type[string]  `json:"accountId,omitempty"`
	AccountNumber          nullable.Type[string]  `json:"accountNumber,omitempty"`
	Amount                 nullable.Type[float64] `json:"amount,omitempty"`
	Comment                nullable.Type[string]  `json:"comment,omitempty"`
	Description            nullable.Type[string]  `json:"description,omitempty"`
	DocumentNumber         nullable.Type[string]  `json:"documentNumber,omitempty"`
	ExternalDocumentNumber nullable.Type[string]  `json:"externalDocumentNumber,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	JournalDisplayName     nullable.Type[string]  `json:"journalDisplayName,omitempty"`
	LastModifiedDateTime   nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`
	LineNumber             nullable.Type[int64]   `json:"lineNumber,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PostingDate nullable.Type[string] `json:"postingDate,omitempty"`
}
