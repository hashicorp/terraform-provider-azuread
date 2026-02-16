package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdRegistrationRegistrantId{}

// UserIdOnlineMeetingIdRegistrationRegistrantId is a struct representing the Resource ID for a User Id Online Meeting Id Registration Registrant
type UserIdOnlineMeetingIdRegistrationRegistrantId struct {
	UserId                  string
	OnlineMeetingId         string
	MeetingRegistrantBaseId string
}

// NewUserIdOnlineMeetingIdRegistrationRegistrantID returns a new UserIdOnlineMeetingIdRegistrationRegistrantId struct
func NewUserIdOnlineMeetingIdRegistrationRegistrantID(userId string, onlineMeetingId string, meetingRegistrantBaseId string) UserIdOnlineMeetingIdRegistrationRegistrantId {
	return UserIdOnlineMeetingIdRegistrationRegistrantId{
		UserId:                  userId,
		OnlineMeetingId:         onlineMeetingId,
		MeetingRegistrantBaseId: meetingRegistrantBaseId,
	}
}

// ParseUserIdOnlineMeetingIdRegistrationRegistrantID parses 'input' into a UserIdOnlineMeetingIdRegistrationRegistrantId
func ParseUserIdOnlineMeetingIdRegistrationRegistrantID(input string) (*UserIdOnlineMeetingIdRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdRegistrationRegistrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdRegistrationRegistrantIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdRegistrationRegistrantId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdRegistrationRegistrantIDInsensitively(input string) (*UserIdOnlineMeetingIdRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdRegistrationRegistrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdRegistrationRegistrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.MeetingRegistrantBaseId, ok = input.Parsed["meetingRegistrantBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrantBaseId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingIdRegistrationRegistrantID checks that 'input' can be parsed as a User Id Online Meeting Id Registration Registrant ID
func ValidateUserIdOnlineMeetingIdRegistrationRegistrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdRegistrationRegistrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Registration Registrant ID
func (id UserIdOnlineMeetingIdRegistrationRegistrantId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/registration/registrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingRegistrantBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Registration Registrant ID
func (id UserIdOnlineMeetingIdRegistrationRegistrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("registration", "registration", "registration"),
		resourceids.StaticSegment("registrants", "registrants", "registrants"),
		resourceids.UserSpecifiedSegment("meetingRegistrantBaseId", "meetingRegistrantBaseId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Registration Registrant ID
func (id UserIdOnlineMeetingIdRegistrationRegistrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registrant Base: %q", id.MeetingRegistrantBaseId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Registration Registrant (%s)", strings.Join(components, "\n"))
}
