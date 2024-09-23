package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdInstanceId{}

// UserIdEventIdInstanceId is a struct representing the Resource ID for a User Id Event Id Instance
type UserIdEventIdInstanceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserIdEventIdInstanceID returns a new UserIdEventIdInstanceId struct
func NewUserIdEventIdInstanceID(userId string, eventId string, eventId1 string) UserIdEventIdInstanceId {
	return UserIdEventIdInstanceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserIdEventIdInstanceID parses 'input' into a UserIdEventIdInstanceId
func ParseUserIdEventIdInstanceID(input string) (*UserIdEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEventIdInstanceIDInsensitively parses 'input' case-insensitively into a UserIdEventIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdEventIdInstanceIDInsensitively(input string) (*UserIdEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEventIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdEventIdInstanceID checks that 'input' can be parsed as a User Id Event Id Instance ID
func ValidateUserIdEventIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEventIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Event Id Instance ID
func (id UserIdEventIdInstanceId) ID() string {
	fmtString := "/users/%s/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Event Id Instance ID
func (id UserIdEventIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this User Id Event Id Instance ID
func (id UserIdEventIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Event Id Instance (%s)", strings.Join(components, "\n"))
}
