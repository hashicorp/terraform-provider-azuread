package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchedulingGroupInfo struct {
	// The code for the schedulingGroup.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The display name for the schedulingGroup. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// ID of the schedulingGroup.
	SchedulingGroupId nullable.Type[string] `json:"schedulingGroupId,omitempty"`
}
