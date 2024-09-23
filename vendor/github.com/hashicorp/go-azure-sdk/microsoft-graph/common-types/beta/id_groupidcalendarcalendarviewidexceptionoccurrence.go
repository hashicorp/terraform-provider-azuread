package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdExceptionOccurrenceId{}

// GroupIdCalendarCalendarViewIdExceptionOccurrenceId is a struct representing the Resource ID for a Group Id Calendar Calendar View Id Exception Occurrence
type GroupIdCalendarCalendarViewIdExceptionOccurrenceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID returns a new GroupIdCalendarCalendarViewIdExceptionOccurrenceId struct
func NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID(groupId string, eventId string, eventId1 string) GroupIdCalendarCalendarViewIdExceptionOccurrenceId {
	return GroupIdCalendarCalendarViewIdExceptionOccurrenceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceID parses 'input' into a GroupIdCalendarCalendarViewIdExceptionOccurrenceId
func ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceID(input string) (*GroupIdCalendarCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarViewIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIDInsensitively(input string) (*GroupIdCalendarCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarViewIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdCalendarCalendarViewIdExceptionOccurrenceID checks that 'input' can be parsed as a Group Id Calendar Calendar View Id Exception Occurrence ID
func ValidateGroupIdCalendarCalendarViewIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar View Id Exception Occurrence ID
func (id GroupIdCalendarCalendarViewIdExceptionOccurrenceId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar View Id Exception Occurrence ID
func (id GroupIdCalendarCalendarViewIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar View Id Exception Occurrence ID
func (id GroupIdCalendarCalendarViewIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Id Calendar Calendar View Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
