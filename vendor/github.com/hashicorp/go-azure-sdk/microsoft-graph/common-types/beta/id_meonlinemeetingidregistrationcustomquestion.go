package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdRegistrationCustomQuestionId{}

// MeOnlineMeetingIdRegistrationCustomQuestionId is a struct representing the Resource ID for a Me Online Meeting Id Registration Custom Question
type MeOnlineMeetingIdRegistrationCustomQuestionId struct {
	OnlineMeetingId               string
	MeetingRegistrationQuestionId string
}

// NewMeOnlineMeetingIdRegistrationCustomQuestionID returns a new MeOnlineMeetingIdRegistrationCustomQuestionId struct
func NewMeOnlineMeetingIdRegistrationCustomQuestionID(onlineMeetingId string, meetingRegistrationQuestionId string) MeOnlineMeetingIdRegistrationCustomQuestionId {
	return MeOnlineMeetingIdRegistrationCustomQuestionId{
		OnlineMeetingId:               onlineMeetingId,
		MeetingRegistrationQuestionId: meetingRegistrationQuestionId,
	}
}

// ParseMeOnlineMeetingIdRegistrationCustomQuestionID parses 'input' into a MeOnlineMeetingIdRegistrationCustomQuestionId
func ParseMeOnlineMeetingIdRegistrationCustomQuestionID(input string) (*MeOnlineMeetingIdRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdRegistrationCustomQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdRegistrationCustomQuestionIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdRegistrationCustomQuestionId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdRegistrationCustomQuestionIDInsensitively(input string) (*MeOnlineMeetingIdRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdRegistrationCustomQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdRegistrationCustomQuestionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.MeetingRegistrationQuestionId, ok = input.Parsed["meetingRegistrationQuestionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrationQuestionId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdRegistrationCustomQuestionID checks that 'input' can be parsed as a Me Online Meeting Id Registration Custom Question ID
func ValidateMeOnlineMeetingIdRegistrationCustomQuestionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdRegistrationCustomQuestionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Registration Custom Question ID
func (id MeOnlineMeetingIdRegistrationCustomQuestionId) ID() string {
	fmtString := "/me/onlineMeetings/%s/registration/customQuestions/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingRegistrationQuestionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Registration Custom Question ID
func (id MeOnlineMeetingIdRegistrationCustomQuestionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("registration", "registration", "registration"),
		resourceids.StaticSegment("customQuestions", "customQuestions", "customQuestions"),
		resourceids.UserSpecifiedSegment("meetingRegistrationQuestionId", "meetingRegistrationQuestionId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Registration Custom Question ID
func (id MeOnlineMeetingIdRegistrationCustomQuestionId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registration Question: %q", id.MeetingRegistrationQuestionId),
	}
	return fmt.Sprintf("Me Online Meeting Id Registration Custom Question (%s)", strings.Join(components, "\n"))
}
