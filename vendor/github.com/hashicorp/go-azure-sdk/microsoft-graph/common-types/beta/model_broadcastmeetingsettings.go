package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingSettings struct {
	// Defines who can join the Teams live event. Possible values are listed in the following table.
	AllowedAudience *BroadcastMeetingAudience `json:"allowedAudience,omitempty"`

	// Caption settings of a Teams live event.
	Captions *BroadcastMeetingCaptionSettings `json:"captions,omitempty"`

	// Indicates whether attendee report is enabled for this Teams live event. Default value is false.
	IsAttendeeReportEnabled nullable.Type[bool] `json:"isAttendeeReportEnabled,omitempty"`

	// Indicates whether Q&A is enabled for this Teams live event. Default value is false.
	IsQuestionAndAnswerEnabled nullable.Type[bool] `json:"isQuestionAndAnswerEnabled,omitempty"`

	// Indicates whether recording is enabled for this Teams live event. Default value is false.
	IsRecordingEnabled nullable.Type[bool] `json:"isRecordingEnabled,omitempty"`

	// Indicates whether video on demand is enabled for this Teams live event. Default value is false.
	IsVideoOnDemandEnabled nullable.Type[bool] `json:"isVideoOnDemandEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
