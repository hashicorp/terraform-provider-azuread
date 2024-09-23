package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StrongAuthenticationRequirements struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Sets the per-user MFA state for the user. The possible values are: disabled, enforced, enabled, unknownFutureValue.
	PerUserMfaState *PerUserMfaState `json:"perUserMfaState,omitempty"`
}
