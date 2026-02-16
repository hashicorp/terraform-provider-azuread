package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowOrBlockListAction struct {
	// Specifies whether the tenant allow-or-block list is an allow or block. The possible values are: allow, block, and
	// unkownFutureValue.
	Action *SecurityTenantAllowBlockListAction `json:"action,omitempty"`

	// Specifies when the tenant allow-block-list expires in date time.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// Specifies the note added to the tenant allow-or-block list entry in the format of string.
	Note nullable.Type[string] `json:"note,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains the result of the submission that lead to the tenant allow-block-list entry creation.
	Results *[]SecurityTenantAllowBlockListEntryResult `json:"results,omitempty"`
}
