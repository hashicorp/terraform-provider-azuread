package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonType struct {
	// The type of data source, such as Person.
	Class nullable.Type[string] `json:"class,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The secondary type of data source, such as OrganizationUser.
	Subclass nullable.Type[string] `json:"subclass,omitempty"`
}
