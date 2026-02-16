package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetails struct {
	// The admin configuration of the policy on the user's authentication app. For a policy that does not impact the
	// success/failure of the authentication, the evaluation defaults to notApplicable. The possible values are:
	// notApplicable, enabled, disabled, unknownFutureValue.
	AdminConfiguration *AuthenticationAppAdminConfiguration `json:"adminConfiguration,omitempty"`

	// Evaluates the success/failure of the authentication based on the admin configuration of the policy on the user's
	// client authentication app. The possible values are: success, failure, unknownFutureValue.
	AuthenticationEvaluation *AuthenticationAppEvaluation `json:"authenticationEvaluation,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the policy enforced on the user's authentication app.
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`

	// Refers to whether the policy executed as expected on the user's client authentication app. The possible values are:
	// unknown, appLockOutOfDate, appLockEnabled, appLockDisabled, appContextOutOfDate, appContextShown, appContextNotShown,
	// locationContextOutOfDate, locationContextShown, locationContextNotShown, numberMatchOutOfDate,
	// numberMatchCorrectNumberEntered, numberMatchIncorrectNumberEntered, numberMatchDeny,
	// tamperResistantHardwareOutOfDate, tamperResistantHardwareUsed, tamperResistantHardwareNotUsed, unknownFutureValue.
	Status *AuthenticationAppPolicyStatus `json:"status,omitempty"`
}
