package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdExceptionOccurrenceId{}

// GroupIdCalendarViewIdExceptionOccurrenceId is a struct representing the Resource ID for a Group Id Calendar View Id Exception Occurrence
type GroupIdCalendarViewIdExceptionOccurrenceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupIdCalendarViewIdExceptionOccurrenceID returns a new GroupIdCalendarViewIdExceptionOccurrenceId struct
func NewGroupIdCalendarViewIdExceptionOccurrenceID(groupId string, eventId string, eventId1 string) GroupIdCalendarViewIdExceptionOccurrenceId {
	return GroupIdCalendarViewIdExceptionOccurrenceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupIdCalendarViewIdExceptionOccurrenceID parses 'input' into a GroupIdCalendarViewIdExceptionOccurrenceId
func ParseGroupIdCalendarViewIdExceptionOccurrenceID(input string) (*GroupIdCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarViewIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarViewIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarViewIdExceptionOccurrenceIDInsensitively(input string) (*GroupIdCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarViewIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdCalendarViewIdExceptionOccurrenceID checks that 'input' can be parsed as a Group Id Calendar View Id Exception Occurrence ID
func ValidateGroupIdCalendarViewIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarViewIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar View Id Exception Occurrence ID
func (id GroupIdCalendarViewIdExceptionOccurrenceId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar View Id Exception Occurrence ID
func (id GroupIdCalendarViewIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this Group Id Calendar View Id Exception Occurrence ID
func (id GroupIdCalendarViewIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Id Calendar View Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
