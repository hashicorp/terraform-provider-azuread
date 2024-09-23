package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorFeatureSettings struct {
	// Determines whether the user's Authenticator app shows them the client app they're signing into.
	DisplayAppInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayAppInformationRequiredState,omitempty"`

	// Determines whether the user's Authenticator app shows them the geographic location of where the authentication
	// request originated from.
	DisplayLocationInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayLocationInformationRequiredState,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
