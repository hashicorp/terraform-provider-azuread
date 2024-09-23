package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarEventId{}

// GroupIdCalendarEventId is a struct representing the Resource ID for a Group Id Calendar Event
type GroupIdCalendarEventId struct {
	GroupId string
	EventId string
}

// NewGroupIdCalendarEventID returns a new GroupIdCalendarEventId struct
func NewGroupIdCalendarEventID(groupId string, eventId string) GroupIdCalendarEventId {
	return GroupIdCalendarEventId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupIdCalendarEventID parses 'input' into a GroupIdCalendarEventId
func ParseGroupIdCalendarEventID(input string) (*GroupIdCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarEventIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarEventIDInsensitively(input string) (*GroupIdCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateGroupIdCalendarEventID checks that 'input' can be parsed as a Group Id Calendar Event ID
func ValidateGroupIdCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Event ID
func (id GroupIdCalendarEventId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Event ID
func (id GroupIdCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Event ID
func (id GroupIdCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Id Calendar Event (%s)", strings.Join(components, "\n"))
}
