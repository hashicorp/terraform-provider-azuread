package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfServiceSignUpAuthenticationFlowConfiguration struct {
	// Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property isn't a
	// key. Required.
	IsEnabled bool `json:"isEnabled"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
