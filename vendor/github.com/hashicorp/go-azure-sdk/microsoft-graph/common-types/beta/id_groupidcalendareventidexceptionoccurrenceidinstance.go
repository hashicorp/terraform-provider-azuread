package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId{}

// GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId is a struct representing the Resource ID for a Group Id Calendar Event Id Exception Occurrence Id Instance
type GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID returns a new GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId struct
func NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID(groupId string, eventId string, eventId1 string, eventId2 string) GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId {
	return GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID parses 'input' into a GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId
func ParseGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID(input string) (*GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarEventIdExceptionOccurrenceIdInstanceIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarEventIdExceptionOccurrenceIdInstanceIDInsensitively(input string) (*GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	return nil
}

// ValidateGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID checks that 'input' can be parsed as a Group Id Calendar Event Id Exception Occurrence Id Instance ID
func ValidateGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Event Id Exception Occurrence Id Instance ID
func (id GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Event Id Exception Occurrence Id Instance ID
func (id GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2"),
	}
}

// String returns a human-readable description of this Group Id Calendar Event Id Exception Occurrence Id Instance ID
func (id GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Group Id Calendar Event Id Exception Occurrence Id Instance (%s)", strings.Join(components, "\n"))
}
