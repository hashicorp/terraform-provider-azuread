package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWhoisContact struct {
	// The physical address of the entity.
	Address *PhysicalAddress `json:"address,omitempty"`

	// The email of this WHOIS contact.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The fax of this WHOIS contact. No format is guaranteed.
	Fax nullable.Type[string] `json:"fax,omitempty"`

	// The name of this WHOIS contact.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The organization of this WHOIS contact.
	Organization nullable.Type[string] `json:"organization,omitempty"`

	// The telephone of this WHOIS contact. No format is guaranteed.
	Telephone nullable.Type[string] `json:"telephone,omitempty"`
}
