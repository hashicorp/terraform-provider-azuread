package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdExtensionId{}

// GroupIdCalendarCalendarViewIdExtensionId is a struct representing the Resource ID for a Group Id Calendar Calendar View Id Extension
type GroupIdCalendarCalendarViewIdExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupIdCalendarCalendarViewIdExtensionID returns a new GroupIdCalendarCalendarViewIdExtensionId struct
func NewGroupIdCalendarCalendarViewIdExtensionID(groupId string, eventId string, extensionId string) GroupIdCalendarCalendarViewIdExtensionId {
	return GroupIdCalendarCalendarViewIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdCalendarCalendarViewIdExtensionID parses 'input' into a GroupIdCalendarCalendarViewIdExtensionId
func ParseGroupIdCalendarCalendarViewIdExtensionID(input string) (*GroupIdCalendarCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarViewIdExtensionIDInsensitively(input string) (*GroupIdCalendarCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdCalendarCalendarViewIdExtensionID checks that 'input' can be parsed as a Group Id Calendar Calendar View Id Extension ID
func ValidateGroupIdCalendarCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar View Id Extension ID
func (id GroupIdCalendarCalendarViewIdExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar View Id Extension ID
func (id GroupIdCalendarCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar View Id Extension ID
func (id GroupIdCalendarCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Calendar Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
