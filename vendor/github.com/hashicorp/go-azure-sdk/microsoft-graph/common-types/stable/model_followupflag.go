package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlag struct {
	// The date and time that the follow-up was finished.
	CompletedDateTime *DateTimeTimeZone `json:"completedDateTime,omitempty"`

	// The date and time that the follow-up is to be finished. Note: To set the due date, you must also specify the
	// startDateTime; otherwise, you get a 400 Bad Request response.
	DueDateTime *DateTimeTimeZone `json:"dueDateTime,omitempty"`

	// The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
	FlagStatus *FollowupFlagStatus `json:"flagStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time that the follow-up is to begin.
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`
}
