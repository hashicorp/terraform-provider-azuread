package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationUser struct {
	// Display name of the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Email address of the user.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of days the user is OOF during a simulation journey/course of a campaign.
	OutOfOfficeDays nullable.Type[int64] `json:"outOfOfficeDays,omitempty"`

	// The id property value of the user resource that represents the user in the Microsoft Entra tenant.
	UserId nullable.Type[string] `json:"userId,omitempty"`
}
