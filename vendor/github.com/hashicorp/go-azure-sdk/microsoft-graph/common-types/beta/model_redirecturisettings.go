package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedirectUriSettings struct {
	// Identifies the specific URI within the redirectURIs collection in SAML SSO flows. Defaults to null. The index is
	// unique across all the redirectUris for the application.
	Index nullable.Type[int64] `json:"index,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the URI that tokens are sent to.
	Uri nullable.Type[string] `json:"uri,omitempty"`
}
