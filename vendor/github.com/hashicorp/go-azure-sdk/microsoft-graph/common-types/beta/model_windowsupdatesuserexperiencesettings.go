package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUserExperienceSettings struct {
	// Specifies the number of days after an update is installed, during which the user of the device can control when the
	// device restarts.
	DaysUntilForcedReboot nullable.Type[int64] `json:"daysUntilForcedReboot,omitempty"`

	// Specifies whether the update is offered as a hotpatch. It can only be set to true on automatic policies that target
	// monthly security updates.
	IsHotpatchEnabled nullable.Type[bool] `json:"isHotpatchEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies whether the update is offered as Optional rather than Required.
	OfferAsOptional nullable.Type[bool] `json:"offerAsOptional,omitempty"`
}
