package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerApprovalRequirement struct {
	// Specifies whether approval is required to complete the plannerTask. When this property is set to true, the task can
	// only be marked complete if an approval is created for the task and approved.
	IsApprovalRequired nullable.Type[bool] `json:"isApprovalRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
