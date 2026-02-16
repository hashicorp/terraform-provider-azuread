package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Account struct {
	Blocked              nullable.Type[bool]   `json:"blocked,omitempty"`
	Category             nullable.Type[string] `json:"category,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	Number               nullable.Type[string] `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SubCategory nullable.Type[string] `json:"subCategory,omitempty"`
}
