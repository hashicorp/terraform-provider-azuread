package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrength struct {
	// Identifier of the authentication strength.
	AuthenticationStrengthId nullable.Type[string] `json:"authenticationStrengthId,omitempty"`

	// The result of the authentication strength. The possible values are: notSet, skippedForProofUp, satisfied,
	// singleChallengeRequired, multipleChallengesRequired, singleRegistrationRequired, multipleRegistrationsRequired,
	// cannotSatisfyDueToCombinationConfiguration, cannotSatisfy, unknownFutureValue.
	AuthenticationStrengthResult *AuthenticationStrengthResult `json:"authenticationStrengthResult,omitempty"`

	// The name of the authentication strength.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
