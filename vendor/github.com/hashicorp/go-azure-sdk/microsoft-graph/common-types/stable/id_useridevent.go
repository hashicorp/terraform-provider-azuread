package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventId{}

// UserIdEventId is a struct representing the Resource ID for a User Id Event
type UserIdEventId struct {
	UserId  string
	EventId string
}

// NewUserIdEventID returns a new UserIdEventId struct
func NewUserIdEventID(userId string, eventId string) UserIdEventId {
	return UserIdEventId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserIdEventID parses 'input' into a UserIdEventId
func ParseUserIdEventID(input string) (*UserIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEventIDInsensitively parses 'input' case-insensitively into a UserIdEventId
// note: this method should only be used for API response data and not user input
func ParseUserIdEventIDInsensitively(input string) (*UserIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdEventID checks that 'input' can be parsed as a User Id Event ID
func ValidateUserIdEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Event ID
func (id UserIdEventId) ID() string {
	fmtString := "/users/%s/events/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Event ID
func (id UserIdEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
	}
}

// String returns a human-readable description of this User Id Event ID
func (id UserIdEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Event (%s)", strings.Join(components, "\n"))
}
