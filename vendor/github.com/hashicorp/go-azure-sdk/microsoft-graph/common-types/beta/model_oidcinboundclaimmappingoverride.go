package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OidcInboundClaimMappingOverride struct {
	// End-user's preferred postal address. The value of the address member is a JSON RFC8259 structure containing some or
	// all of the members defined in the resource type
	Address *OidcAddressInboundClaims `json:"address,omitempty"`

	// End-user's preferred e-mail address. Its value MUST conform to the RFC 5322 addr-spec syntax.
	Email nullable.Type[string] `json:"email,omitempty"`

	// True if the end-user's e-mail address has been verified by the identity provider; otherwise, false. When this claim
	// value is true, this means that your identity provider took affirmative steps to ensure that this e-mail address was
	// controlled by the end-user at the time the verification was performed. If this claim value is false, or not mapped
	// with any claim of the identity provider, the user is asked to verify email during sign-up if email is required in the
	// user flow.
	Emailverified nullable.Type[string] `json:"email_verified,omitempty"`

	// Surname(s) or family name of the end-user.
	Familyname nullable.Type[string] `json:"family_name,omitempty"`

	// Given name(s) or first name(s) of the end-user.
	Givenname nullable.Type[string] `json:"given_name,omitempty"`

	// End-user's full name in displayable form including all name parts, possibly including titles and suffixes, ordered
	// according to the end-user's locale and preferences.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The claim provides the phone number for the user.
	Phonenumber nullable.Type[string] `json:"phone_number,omitempty"`

	// True if the end-user's phone number has been verified; otherwise, false. When this claim value is true, this means
	// that your identity provider took affirmative steps to verify the phone number.
	Phonenumberverified nullable.Type[string] `json:"phone_number_verified,omitempty"`

	// Subject - Identifier for the end-user at the Issuer.
	Sub nullable.Type[string] `json:"sub,omitempty"`
}
