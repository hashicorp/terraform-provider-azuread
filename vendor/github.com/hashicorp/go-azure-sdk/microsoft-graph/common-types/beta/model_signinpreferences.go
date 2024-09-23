package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInPreferences struct {
	// Indicates whether the credential preferences of the system are enabled.
	IsSystemPreferredAuthenticationMethodEnabled nullable.Type[bool] `json:"isSystemPreferredAuthenticationMethodEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The default second-factor method used by the user when signing in. If a user is enabled for system-preferred
	// authentication, then this value is ignored except for a few scenarios where a user is authenticating via NPS
	// extension or ADFS adapter. Possible values are push, oath, voiceMobile, voiceAlternateMobile, voiceOffice, sms, and
	// unknownFutureValue
	UserPreferredMethodForSecondaryAuthentication *UserDefaultAuthenticationMethodType `json:"userPreferredMethodForSecondaryAuthentication,omitempty"`
}
