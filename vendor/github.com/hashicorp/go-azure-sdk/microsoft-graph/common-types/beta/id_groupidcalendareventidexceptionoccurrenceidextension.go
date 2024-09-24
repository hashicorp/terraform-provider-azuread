package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId{}

// GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a Group Id Calendar Event Id Exception Occurrence Id Extension
type GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID returns a new GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId struct
func NewGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId {
	return GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID parses 'input' into a GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId
func ParseGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID(input string) (*GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarEventIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarEventIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a Group Id Calendar Event Id Exception Occurrence Id Extension ID
func ValidateGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarEventIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Event Id Exception Occurrence Id Extension ID
func (id GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Event Id Exception Occurrence Id Extension ID
func (id GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Event Id Exception Occurrence Id Extension ID
func (id GroupIdCalendarEventIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Calendar Event Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
