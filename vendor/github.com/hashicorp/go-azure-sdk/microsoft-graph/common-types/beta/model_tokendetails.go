package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenDetails struct {
	// Represents when the authentication for this token occurred.
	IssuedAtDateTime nullable.Type[string] `json:"issuedAtDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the token identifier claim. This ID is a unique per-token identifier that is case-sensitive.
	UniqueTokenIdentifier nullable.Type[string] `json:"uniqueTokenIdentifier,omitempty"`
}
