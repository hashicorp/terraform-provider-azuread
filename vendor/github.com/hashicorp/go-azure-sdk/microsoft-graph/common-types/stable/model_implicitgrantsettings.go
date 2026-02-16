package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImplicitGrantSettings struct {
	// Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
	EnableAccessTokenIssuance nullable.Type[bool] `json:"enableAccessTokenIssuance,omitempty"`

	// Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
	EnableIdTokenIssuance nullable.Type[bool] `json:"enableIdTokenIssuance,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
