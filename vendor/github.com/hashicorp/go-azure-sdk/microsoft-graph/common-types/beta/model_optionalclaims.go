package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OptionalClaims struct {
	// The optional claims returned in the JWT access token.
	AccessToken *[]OptionalClaim `json:"accessToken,omitempty"`

	// The optional claims returned in the JWT ID token.
	IdToken *[]OptionalClaim `json:"idToken,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The optional claims returned in the SAML token.
	Saml2Token *[]OptionalClaim `json:"saml2Token,omitempty"`
}
