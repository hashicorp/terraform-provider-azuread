package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlternativeSecurityId struct {
	// For internal use only.
	IdentityProvider nullable.Type[string] `json:"identityProvider,omitempty"`

	// For internal use only.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// For internal use only.
	Type nullable.Type[int64] `json:"type,omitempty"`
}
