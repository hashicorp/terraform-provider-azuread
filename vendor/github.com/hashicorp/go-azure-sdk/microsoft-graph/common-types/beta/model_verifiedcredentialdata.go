package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedCredentialData struct {
	// The authority ID for the issuer.
	Authority nullable.Type[string] `json:"authority,omitempty"`

	// Key-value pair of claims retrieved from the credential that the user presented, and the service verified.
	Claims *VerifiedCredentialClaims `json:"claims,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of credential types provided by the issuer.
	Type *[]string `json:"type,omitempty"`
}
