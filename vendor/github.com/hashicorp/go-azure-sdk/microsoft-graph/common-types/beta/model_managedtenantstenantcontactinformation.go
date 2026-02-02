package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantContactInformation struct {
	// The email address for the contact. Optional
	Email nullable.Type[string] `json:"email,omitempty"`

	// The name for the contact. Required.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The notes associated with the contact. Optional
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The phone number for the contact. Optional.
	Phone nullable.Type[string] `json:"phone,omitempty"`

	// The title for the contact. Required.
	Title nullable.Type[string] `json:"title,omitempty"`
}
