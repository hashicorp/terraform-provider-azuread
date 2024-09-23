package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnPublicKeyCredentialParameters struct {
	// Specifies the cryptographic signature algorithm used for the new credential. The algorithm identifiers should be
	// values registered in the IANA COSE algorithms registry. For more information, see IANA-COSE-ALGS-REG.
	Alg nullable.Type[int64] `json:"alg,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the type of credential to be created. The only supported value is public-key.
	Type nullable.Type[string] `json:"type,omitempty"`
}
