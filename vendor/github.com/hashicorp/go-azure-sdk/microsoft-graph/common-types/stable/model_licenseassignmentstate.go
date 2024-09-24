package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseAssignmentState struct {
	AssignedByGroup     nullable.Type[string] `json:"assignedByGroup,omitempty"`
	DisabledPlans       *[]string             `json:"disabledPlans,omitempty"`
	Error               nullable.Type[string] `json:"error,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SkuId nullable.Type[string] `json:"skuId,omitempty"`
	State nullable.Type[string] `json:"state,omitempty"`
}
