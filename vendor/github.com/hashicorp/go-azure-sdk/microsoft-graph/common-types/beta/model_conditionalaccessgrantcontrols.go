package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGrantControls struct {
	AuthenticationStrength *AuthenticationStrengthPolicy `json:"authenticationStrength,omitempty"`

	// List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice,
	// domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
	BuiltInControls *[]ConditionalAccessGrantControl `json:"builtInControls,omitempty"`

	// List of custom controls IDs required by the policy. To learn more about custom control, see Custom controls
	// (preview).
	CustomAuthenticationFactors *[]string `json:"customAuthenticationFactors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the relationship of the grant controls. Possible values: AND, OR.
	Operator nullable.Type[string] `json:"operator,omitempty"`

	// List of terms of use IDs required by the policy.
	TermsOfUse *[]string `json:"termsOfUse,omitempty"`
}
