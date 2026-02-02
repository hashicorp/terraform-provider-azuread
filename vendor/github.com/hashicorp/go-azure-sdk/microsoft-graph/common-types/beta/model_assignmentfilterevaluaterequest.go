package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluateRequest struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Order the devices should be sorted in. Default is ascending on device name.
	OrderBy *[]string `json:"orderBy,omitempty"`

	// Supported platform types.
	Platform *DevicePlatformType `json:"platform,omitempty"`

	// Rule definition of the Assignment Filter.
	Rule *string `json:"rule,omitempty"`

	// Search keyword applied to scope found devices.
	Search nullable.Type[string] `json:"search,omitempty"`

	// Number of records to skip. Default value is 0
	Skip *int64 `json:"skip,omitempty"`

	// Limit of records per request. Default value is 100, if provided less than 0 or greater than 100
	Top *int64 `json:"top,omitempty"`
}
