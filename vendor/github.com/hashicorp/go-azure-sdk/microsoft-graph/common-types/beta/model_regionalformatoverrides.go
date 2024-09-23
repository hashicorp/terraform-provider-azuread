package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionalFormatOverrides struct {
	// The calendar to use, e.g., Gregorian Calendar.Returned by default.
	Calendar nullable.Type[string] `json:"calendar,omitempty"`

	// The first day of the week to use, e.g., Sunday.Returned by default.
	FirstDayOfWeek nullable.Type[string] `json:"firstDayOfWeek,omitempty"`

	// The long date time format to be used for displaying dates.Returned by default.
	LongDateFormat nullable.Type[string] `json:"longDateFormat,omitempty"`

	// The long time format to be used for displaying time.Returned by default.
	LongTimeFormat nullable.Type[string] `json:"longTimeFormat,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The short date time format to be used for displaying dates.Returned by default.
	ShortDateFormat nullable.Type[string] `json:"shortDateFormat,omitempty"`

	// The short time format to be used for displaying time.Returned by default.
	ShortTimeFormat nullable.Type[string] `json:"shortTimeFormat,omitempty"`

	// The timezone to be used for displaying time.Returned by default.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`
}
