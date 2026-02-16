package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StringKeyLongValuePair struct {
	// The mapping of the user type from the source system to the target system. For example:User to User - For Microsoft
	// Entra ID to Microsoft Entra synchronization worker to user - For Workday to Microsoft Entra synchronization.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total number of synchronized objects.
	Value *int64 `json:"value,omitempty"`
}
