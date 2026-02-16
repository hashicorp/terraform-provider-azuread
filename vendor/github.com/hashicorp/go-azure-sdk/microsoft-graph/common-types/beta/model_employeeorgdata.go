package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeOrgData struct {
	// The cost center associated with the user. Returned only on $select. Supports $filter.
	CostCenter nullable.Type[string] `json:"costCenter,omitempty"`

	// The name of the division in which the user works. Returned only on $select. Supports $filter.
	Division nullable.Type[string] `json:"division,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
