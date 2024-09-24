package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityDetails struct {
	// A date specifiying when the Identity was created, could be null
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// A date specifiying when the Identity was active last time, could be null
	LastActiveDateTime nullable.Type[string] `json:"lastActiveDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
