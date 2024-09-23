package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialType struct {
	// The type of credential issued, for example, BusinessCardCredential.
	CredentialType nullable.Type[string] `json:"credentialType,omitempty"`

	// List of the accepted issuers authority as identified by the Microsoft Entra Verified ID service, for example,
	// did:ion:EiAlrenrtD3Lsw0GlbzS1O2YFdy3Xtu8yo35W/<SNIP/>....
	Issuers *[]string `json:"issuers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
