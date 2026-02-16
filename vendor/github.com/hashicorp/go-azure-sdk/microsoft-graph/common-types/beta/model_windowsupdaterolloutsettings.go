package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateRolloutSettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The feature update's ending of release date and time to be set, update, and displayed for a feature Update profile
	// for example: 2020-06-09T10:00:00Z.
	OfferEndDateTimeInUTC nullable.Type[string] `json:"offerEndDateTimeInUTC,omitempty"`

	// The number of day(s) between each set of offers to be set, updated, and displayed for a feature update profile, for
	// example: if OfferStartDateTimeInUTC is 2020-06-09T10:00:00Z, and OfferIntervalInDays is 1, then the next two sets of
	// offers will be made consecutively on 2020-06-10T10:00:00Z (next day at the same specified time) and
	// 2020-06-11T10:00:00Z (next next day at the same specified time) with 1 day in between each set of offers.
	OfferIntervalInDays nullable.Type[int64] `json:"offerIntervalInDays,omitempty"`

	// The feature update's starting date and time to be set, update, and displayed for a feature Update profile for
	// example: 2020-06-09T10:00:00Z.
	OfferStartDateTimeInUTC nullable.Type[string] `json:"offerStartDateTimeInUTC,omitempty"`
}
