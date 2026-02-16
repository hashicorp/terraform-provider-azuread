package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestionsResult struct {
	// A reason for not returning any meeting suggestions. Possible values are: attendeesUnavailable,
	// attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty
	// string if the meetingTimeSuggestions property does include any meeting suggestions.
	EmptySuggestionsReason nullable.Type[string] `json:"emptySuggestionsReason,omitempty"`

	// An array of meeting suggestions.
	MeetingTimeSuggestions *[]MeetingTimeSuggestion `json:"meetingTimeSuggestions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
