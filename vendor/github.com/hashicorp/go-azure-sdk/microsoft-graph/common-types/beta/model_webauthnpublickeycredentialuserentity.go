package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnPublicKeyCredentialUserEntity struct {
	// The display name of the user account bound to the generated credential, as displayed in Microsoft Entra ID.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The Microsoft Entra ID-assigned object ID of the user account bound to the generated credential. The ID is encoded to
	// WebAuthn spec by Microsoft Entra ID and is not represented as a GUID.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The user principal name (UPN) of the user account bound to the generated credential, as displayed in Microsoft Entra
	// ID.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
