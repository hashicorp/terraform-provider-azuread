package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessThirdPartyTokenDetails struct {
	// Time the token will expire.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Time the token was issued at.
	IssuedAtDateTime nullable.Type[string] `json:"issuedAtDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique token identifier.
	UniqueTokenIdentifier *string `json:"uniqueTokenIdentifier,omitempty"`

	// Time the token is valid from.
	ValidFromDateTime *string `json:"validFromDateTime,omitempty"`
}
