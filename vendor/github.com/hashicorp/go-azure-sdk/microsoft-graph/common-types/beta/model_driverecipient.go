package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRecipient struct {
	// The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
	Alias nullable.Type[string] `json:"alias,omitempty"`

	// The email address for the recipient, if the recipient has an associated email address.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the recipient in the directory.
	ObjectId nullable.Type[string] `json:"objectId,omitempty"`
}
