package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimsMapping struct {
	// The claim that provides the display name or full name for the user. It's a required property.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The claim that provides the email address of the user.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The claim that provides the first name of the user.
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The claim that provides the last name of the user.
	Surname nullable.Type[string] `json:"surname,omitempty"`

	// The claim that provides the unique identifier for the signed-in user. It is a required property.
	UserId nullable.Type[string] `json:"userId,omitempty"`
}
