package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdExceptionOccurrenceId{}

// UserIdEventIdExceptionOccurrenceId is a struct representing the Resource ID for a User Id Event Id Exception Occurrence
type UserIdEventIdExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserIdEventIdExceptionOccurrenceID returns a new UserIdEventIdExceptionOccurrenceId struct
func NewUserIdEventIdExceptionOccurrenceID(userId string, eventId string, eventId1 string) UserIdEventIdExceptionOccurrenceId {
	return UserIdEventIdExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserIdEventIdExceptionOccurrenceID parses 'input' into a UserIdEventIdExceptionOccurrenceId
func ParseUserIdEventIdExceptionOccurrenceID(input string) (*UserIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEventIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserIdEventIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserIdEventIdExceptionOccurrenceIDInsensitively(input string) (*UserIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEventIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateUserIdEventIdExceptionOccurrenceID checks that 'input' can be parsed as a User Id Event Id Exception Occurrence ID
func ValidateUserIdEventIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEventIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Event Id Exception Occurrence ID
func (id UserIdEventIdExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Event Id Exception Occurrence ID
func (id UserIdEventIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this User Id Event Id Exception Occurrence ID
func (id UserIdEventIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Event Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
