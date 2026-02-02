package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaffAvailabilityItem struct {
	// Each item in this collection indicates a slot and the status of the staff member.
	AvailabilityItems *[]AvailabilityItem `json:"availabilityItems,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the staff member.
	StaffId nullable.Type[string] `json:"staffId,omitempty"`
}
