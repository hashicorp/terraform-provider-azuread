package user

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FindMeetingTimesRequest struct {
	Attendees               *[]stable.AttendeeBase     `json:"attendees,omitempty"`
	IsOrganizerOptional     nullable.Type[bool]        `json:"isOrganizerOptional,omitempty"`
	LocationConstraint      *stable.LocationConstraint `json:"locationConstraint,omitempty"`
	MaxCandidates           nullable.Type[int64]       `json:"maxCandidates,omitempty"`
	MeetingDuration         nullable.Type[string]      `json:"meetingDuration,omitempty"`
	ReturnSuggestionReasons nullable.Type[bool]        `json:"returnSuggestionReasons,omitempty"`
	TimeConstraint          *stable.TimeConstraint     `json:"timeConstraint,omitempty"`
}
