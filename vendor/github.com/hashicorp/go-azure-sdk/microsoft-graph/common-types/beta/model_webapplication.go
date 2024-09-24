package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplication struct {
	// Home page or landing page of the application.
	HomePageUrl nullable.Type[string] `json:"homePageUrl,omitempty"`

	// Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
	ImplicitGrantSettings *ImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`

	// Specifies the URL that will be used by Microsoft's authorization service to logout a user using front-channel,
	// back-channel or SAML logout protocols.
	LogoutUrl nullable.Type[string] `json:"logoutUrl,omitempty"`

	OAuth2AllowImplicitFlow nullable.Type[bool] `json:"oauth2AllowImplicitFlow,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the index of the URLs where user tokens are sent for sign-in. This is only valid for applications using
	// SAML.
	RedirectUriSettings *[]RedirectUriSettings `json:"redirectUriSettings,omitempty"`

	// Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes
	// and access tokens are sent.
	RedirectUris *[]string `json:"redirectUris,omitempty"`
}
