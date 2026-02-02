package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageTextVisibilitySettings struct {
	// Option to hide the self-service password reset (SSPR) hyperlinks such as 'Can't access your account?', 'Forgot my
	// password' and 'Reset it now' on the sign-in form.
	HideAccountResetCredentials nullable.Type[bool] `json:"hideAccountResetCredentials,omitempty"`

	// Option to hide the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in form.
	HideCannotAccessYourAccount nullable.Type[bool] `json:"hideCannotAccessYourAccount,omitempty"`

	// Option to hide the self-service password reset (SSPR) 'Forgot my password' hyperlink on the sign-in form.
	HideForgotMyPassword nullable.Type[bool] `json:"hideForgotMyPassword,omitempty"`

	// Option to hide the 'Privacy & Cookies' hyperlink in the footer.
	HidePrivacyAndCookies nullable.Type[bool] `json:"hidePrivacyAndCookies,omitempty"`

	// Option to hide the self-service password reset (SSPR) 'reset it now' hyperlink on the sign-in form.
	HideResetItNow nullable.Type[bool] `json:"hideResetItNow,omitempty"`

	// Option to hide the 'Terms of Use' hyperlink in the footer.
	HideTermsOfUse nullable.Type[bool] `json:"hideTermsOfUse,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
