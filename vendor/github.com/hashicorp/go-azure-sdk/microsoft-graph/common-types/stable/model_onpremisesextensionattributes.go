package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesExtensionAttributes struct {
	// First customizable extension attribute.
	ExtensionAttribute1 nullable.Type[string] `json:"extensionAttribute1,omitempty"`

	// Tenth customizable extension attribute.
	ExtensionAttribute10 nullable.Type[string] `json:"extensionAttribute10,omitempty"`

	// Eleventh customizable extension attribute.
	ExtensionAttribute11 nullable.Type[string] `json:"extensionAttribute11,omitempty"`

	// Twelfth customizable extension attribute.
	ExtensionAttribute12 nullable.Type[string] `json:"extensionAttribute12,omitempty"`

	// Thirteenth customizable extension attribute.
	ExtensionAttribute13 nullable.Type[string] `json:"extensionAttribute13,omitempty"`

	// Fourteenth customizable extension attribute.
	ExtensionAttribute14 nullable.Type[string] `json:"extensionAttribute14,omitempty"`

	// Fifteenth customizable extension attribute.
	ExtensionAttribute15 nullable.Type[string] `json:"extensionAttribute15,omitempty"`

	// Second customizable extension attribute.
	ExtensionAttribute2 nullable.Type[string] `json:"extensionAttribute2,omitempty"`

	// Third customizable extension attribute.
	ExtensionAttribute3 nullable.Type[string] `json:"extensionAttribute3,omitempty"`

	// Fourth customizable extension attribute.
	ExtensionAttribute4 nullable.Type[string] `json:"extensionAttribute4,omitempty"`

	// Fifth customizable extension attribute.
	ExtensionAttribute5 nullable.Type[string] `json:"extensionAttribute5,omitempty"`

	// Sixth customizable extension attribute.
	ExtensionAttribute6 nullable.Type[string] `json:"extensionAttribute6,omitempty"`

	// Seventh customizable extension attribute.
	ExtensionAttribute7 nullable.Type[string] `json:"extensionAttribute7,omitempty"`

	// Eighth customizable extension attribute.
	ExtensionAttribute8 nullable.Type[string] `json:"extensionAttribute8,omitempty"`

	// Ninth customizable extension attribute.
	ExtensionAttribute9 nullable.Type[string] `json:"extensionAttribute9,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
