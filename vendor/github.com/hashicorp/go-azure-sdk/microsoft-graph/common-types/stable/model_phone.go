package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Phone struct {
	Language nullable.Type[string] `json:"language,omitempty"`

	// The phone number.
	Number nullable.Type[string] `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Region nullable.Type[string] `json:"region,omitempty"`

	// The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax,
	// otherFax, pager, radio.
	Type *PhoneType `json:"type,omitempty"`
}
