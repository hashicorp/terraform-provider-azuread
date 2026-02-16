package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordProfile struct {
	// true if the user must change their password on the next sign-in; otherwise false. If not set, default is false.
	ForceChangePasswordNextSignIn nullable.Type[bool] `json:"forceChangePasswordNextSignIn,omitempty"`

	// If true, at next sign-in, the user must perform a multifactor authentication (MFA) before being forced to change
	// their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first
	// perform a multifactor authentication before password change. After a password change, this property will be
	// automatically reset to false. If not set, default is false.
	ForceChangePasswordNextSignInWithMfa nullable.Type[bool] `json:"forceChangePasswordNextSignInWithMfa,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The password for the user. This property is required when a user is created. It can be updated, but the user will be
	// required to change the password on the next sign-in. The password must satisfy minimum requirements as specified by
	// the user's passwordPolicies property. By default, a strong password is required.
	Password nullable.Type[string] `json:"password,omitempty"`
}
