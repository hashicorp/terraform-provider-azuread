package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdRegistrationCustomQuestionId{}

// UserIdOnlineMeetingIdRegistrationCustomQuestionId is a struct representing the Resource ID for a User Id Online Meeting Id Registration Custom Question
type UserIdOnlineMeetingIdRegistrationCustomQuestionId struct {
	UserId                        string
	OnlineMeetingId               string
	MeetingRegistrationQuestionId string
}

// NewUserIdOnlineMeetingIdRegistrationCustomQuestionID returns a new UserIdOnlineMeetingIdRegistrationCustomQuestionId struct
func NewUserIdOnlineMeetingIdRegistrationCustomQuestionID(userId string, onlineMeetingId string, meetingRegistrationQuestionId string) UserIdOnlineMeetingIdRegistrationCustomQuestionId {
	return UserIdOnlineMeetingIdRegistrationCustomQuestionId{
		UserId:                        userId,
		OnlineMeetingId:               onlineMeetingId,
		MeetingRegistrationQuestionId: meetingRegistrationQuestionId,
	}
}

// ParseUserIdOnlineMeetingIdRegistrationCustomQuestionID parses 'input' into a UserIdOnlineMeetingIdRegistrationCustomQuestionId
func ParseUserIdOnlineMeetingIdRegistrationCustomQuestionID(input string) (*UserIdOnlineMeetingIdRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdRegistrationCustomQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdRegistrationCustomQuestionIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdRegistrationCustomQuestionId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdRegistrationCustomQuestionIDInsensitively(input string) (*UserIdOnlineMeetingIdRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdRegistrationCustomQuestionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdRegistrationCustomQuestionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.MeetingRegistrationQuestionId, ok = input.Parsed["meetingRegistrationQuestionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrationQuestionId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingIdRegistrationCustomQuestionID checks that 'input' can be parsed as a User Id Online Meeting Id Registration Custom Question ID
func ValidateUserIdOnlineMeetingIdRegistrationCustomQuestionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdRegistrationCustomQuestionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Registration Custom Question ID
func (id UserIdOnlineMeetingIdRegistrationCustomQuestionId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/registration/customQuestions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingRegistrationQuestionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Registration Custom Question ID
func (id UserIdOnlineMeetingIdRegistrationCustomQuestionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("registration", "registration", "registration"),
		resourceids.StaticSegment("customQuestions", "customQuestions", "customQuestions"),
		resourceids.UserSpecifiedSegment("meetingRegistrationQuestionId", "meetingRegistrationQuestionId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Registration Custom Question ID
func (id UserIdOnlineMeetingIdRegistrationCustomQuestionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registration Question: %q", id.MeetingRegistrationQuestionId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Registration Custom Question (%s)", strings.Join(components, "\n"))
}
