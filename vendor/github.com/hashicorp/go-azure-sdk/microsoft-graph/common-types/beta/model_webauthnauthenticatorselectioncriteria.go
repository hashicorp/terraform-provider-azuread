package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnAuthenticatorSelectionCriteria struct {
	// Microsoft Entra ID-preferred attachment modality. For more information, see Authenticator Attachment Modality
	AuthenticatorAttachment nullable.Type[string] `json:"authenticatorAttachment,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Microsoft Entra ID-preferred client-side credential discoverability. Currently always true. The WebAuthn
	// authenticator must store the credential identifier on the authenticator.
	RequireResidentKey nullable.Type[bool] `json:"requireResidentKey,omitempty"`

	// Microsoft Entra ID requirement to verify the user is present during credential provisioning. Currently always
	// required.
	UserVerification nullable.Type[string] `json:"userVerification,omitempty"`
}
