package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCoachmark struct {
	// The coachmark location.
	CoachmarkLocation *CoachmarkLocation `json:"coachmarkLocation,omitempty"`

	// The description about the coachmark.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The coachmark indicator.
	Indicator nullable.Type[string] `json:"indicator,omitempty"`

	// Indicates whether the coachmark is valid or not.
	IsValid nullable.Type[bool] `json:"isValid,omitempty"`

	// The coachmark language.
	Language nullable.Type[string] `json:"language,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The coachmark order.
	Order nullable.Type[string] `json:"order,omitempty"`
}
