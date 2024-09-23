package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationDetail struct {
	// The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice,
	// Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
	AuthenticationMethod nullable.Type[string] `json:"authenticationMethod,omitempty"`

	// Details about the authentication method used to perform this authentication step. For example, phone number (for SMS
	// and voice), device name (for Authenticator app), and password source (for example, cloud, AD FS, PTA, PHS).
	AuthenticationMethodDetail nullable.Type[string] `json:"authenticationMethodDetail,omitempty"`

	// Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z.
	AuthenticationStepDateTime nullable.Type[string] `json:"authenticationStepDateTime,omitempty"`

	// The step of authentication that this satisfied. For example, primary authentication, or multifactor authentication.
	AuthenticationStepRequirement nullable.Type[string] `json:"authenticationStepRequirement,omitempty"`

	// Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input -
	// timed out, phone unreachable, or claim in token.
	AuthenticationStepResultDetail nullable.Type[string] `json:"authenticationStepResultDetail,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the status of the authentication step. Possible values: succeeded, failed.
	Succeeded nullable.Type[bool] `json:"succeeded,omitempty"`
}
