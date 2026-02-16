package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdExtensionId{}

// UserIdEventIdExtensionId is a struct representing the Resource ID for a User Id Event Id Extension
type UserIdEventIdExtensionId struct {
	UserId      string
	EventId     string
	ExtensionId string
}

// NewUserIdEventIdExtensionID returns a new UserIdEventIdExtensionId struct
func NewUserIdEventIdExtensionID(userId string, eventId string, extensionId string) UserIdEventIdExtensionId {
	return UserIdEventIdExtensionId{
		UserId:      userId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdEventIdExtensionID parses 'input' into a UserIdEventIdExtensionId
func ParseUserIdEventIdExtensionID(input string) (*UserIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdEventIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdEventIdExtensionIDInsensitively(input string) (*UserIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdEventIdExtensionID checks that 'input' can be parsed as a User Id Event Id Extension ID
func ValidateUserIdEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Event Id Extension ID
func (id UserIdEventIdExtensionId) ID() string {
	fmtString := "/users/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Event Id Extension ID
func (id UserIdEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Event Id Extension ID
func (id UserIdEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Event Id Extension (%s)", strings.Join(components, "\n"))
}
