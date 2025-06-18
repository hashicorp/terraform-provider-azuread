package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxItemImportSession struct {
	// The date and time in UTC when the import session expires. The date and time information uses ISO 8601 format and is
	// always in UTC. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The URL endpoint that accepts POST requests for FastTransfer stream format of the item.
	ImportUrl nullable.Type[string] `json:"importUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
