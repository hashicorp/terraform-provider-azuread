package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorFeatureSettings struct {
	// Determines whether users are able to approve push notifications on other Microsoft applications such as Outlook
	// Mobile.
	CompanionAppAllowedState *AuthenticationMethodFeatureConfiguration `json:"companionAppAllowedState,omitempty"`

	// Determines whether the user's Authenticator app shows them the client app they're signing into.
	DisplayAppInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayAppInformationRequiredState,omitempty"`

	// Determines whether the user's Authenticator app shows them the geographic location of where the authentication
	// request originated from.
	DisplayLocationInformationRequiredState *AuthenticationMethodFeatureConfiguration `json:"displayLocationInformationRequiredState,omitempty"`

	// Specifies whether the user needs to enter a number in the Authenticator app from the login screen to complete their
	// login. Value is ignored for phone sign-in notifications.
	NumberMatchingRequiredState *AuthenticationMethodFeatureConfiguration `json:"numberMatchingRequiredState,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
