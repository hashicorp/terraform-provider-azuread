package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeneralLedgerEntry struct {
	Account              *Account               `json:"account,omitempty"`
	AccountId            nullable.Type[string]  `json:"accountId,omitempty"`
	AccountNumber        nullable.Type[string]  `json:"accountNumber,omitempty"`
	CreditAmount         nullable.Type[float64] `json:"creditAmount,omitempty"`
	DebitAmount          nullable.Type[float64] `json:"debitAmount,omitempty"`
	Description          nullable.Type[string]  `json:"description,omitempty"`
	DocumentNumber       nullable.Type[string]  `json:"documentNumber,omitempty"`
	DocumentType         nullable.Type[string]  `json:"documentType,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	LastModifiedDateTime nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PostingDate nullable.Type[string] `json:"postingDate,omitempty"`
}
