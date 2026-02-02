package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkableIdentifiers struct {
	// Represents a unique identifier for the device from which a user is interacting with an application.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents a unique identifier for an entire session and is generated when a user does interactive authentication.
	// This ID helps link all authentication artifacts issued from a single root authentication.
	SessionId nullable.Type[string] `json:"sessionId,omitempty"`

	// Property that represents an access token's unique identifier and the time when the token was issued.
	TokenDetails *TokenDetails `json:"tokenDetails,omitempty"`
}
