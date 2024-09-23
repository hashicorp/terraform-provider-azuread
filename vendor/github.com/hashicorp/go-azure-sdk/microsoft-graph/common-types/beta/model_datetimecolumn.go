package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeColumn struct {
	// How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more
	// details. If unspecified, treated as default.
	DisplayAs nullable.Type[string] `json:"displayAs,omitempty"`

	// Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or
	// dateTime
	Format nullable.Type[string] `json:"format,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
