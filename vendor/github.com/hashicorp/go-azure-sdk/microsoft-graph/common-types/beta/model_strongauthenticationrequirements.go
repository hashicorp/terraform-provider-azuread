package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StrongAuthenticationRequirements struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Sets the per-user MFA state for the user. The possible values are: disabled, enforced, enabled, unknownFutureValue.
	// When you update a user's MFA state to enabled and the user has already registered an MFA method, their state changes
	// automatically to enforced.
	PerUserMfaState *PerUserMfaState `json:"perUserMfaState,omitempty"`
}
