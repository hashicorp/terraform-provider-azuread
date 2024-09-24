package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInActivity struct {
	// The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client
	// attempted (either successfully or unsuccessfully) to sign in to the directory on behalf of a user. Because some users
	// may use clients to access tenant resources rather than signing into your tenant directly, you can use the
	// non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014 is 2014-01-01T00:00:00Z. Microsoft Entra ID maintains non-interactive sign-ins going back to May 2020. For more
	// information about using the value of this property, see Manage inactive user accounts in Microsoft Entra ID.
	LastNonInteractiveSignInDateTime nullable.Type[string] `json:"lastNonInteractiveSignInDateTime,omitempty"`

	// Request identifier of the last non-interactive sign-in performed by this user.
	LastNonInteractiveSignInRequestId nullable.Type[string] `json:"lastNonInteractiveSignInRequestId,omitempty"`

	// The last interactive sign-in date and time for a specific user. You can use this field to calculate the last time a
	// user attempted (either successfully or unsuccessfully) to sign in to the directory the directory with an interactive
	// authentication method. This field can be used to build reports, such as inactive users. The timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Microsoft Entra ID maintains interactive sign-ins going back to April 2020. For more
	// information about using the value of this property, see Manage inactive user accounts in Microsoft Entra ID.
	LastSignInDateTime nullable.Type[string] `json:"lastSignInDateTime,omitempty"`

	// Request identifier of the last interactive sign-in performed by this user.
	LastSignInRequestId nullable.Type[string] `json:"lastSignInRequestId,omitempty"`

	// The date and time of the user's most recent successful sign-in activity. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastSuccessfulSignInDateTime nullable.Type[string] `json:"lastSuccessfulSignInDateTime,omitempty"`

	// The request ID of the last successful sign-in.
	LastSuccessfulSignInRequestId nullable.Type[string] `json:"lastSuccessfulSignInRequestId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
