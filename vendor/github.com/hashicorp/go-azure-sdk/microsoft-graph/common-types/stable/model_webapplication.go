package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplication struct {
	// Home page or landing page of the application.
	HomePageUrl nullable.Type[string] `json:"homePageUrl,omitempty"`

	// Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
	ImplicitGrantSettings *ImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`

	// Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel,
	// back-channel or SAML logout protocols.
	LogoutUrl nullable.Type[string] `json:"logoutUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RedirectUriSettings *[]RedirectUriSettings `json:"redirectUriSettings,omitempty"`

	// Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes
	// and access tokens are sent.
	RedirectUris *[]string `json:"redirectUris,omitempty"`
}
