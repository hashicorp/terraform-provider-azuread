package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewId{}

// GroupIdCalendarCalendarViewId is a struct representing the Resource ID for a Group Id Calendar Calendar View
type GroupIdCalendarCalendarViewId struct {
	GroupId string
	EventId string
}

// NewGroupIdCalendarCalendarViewID returns a new GroupIdCalendarCalendarViewId struct
func NewGroupIdCalendarCalendarViewID(groupId string, eventId string) GroupIdCalendarCalendarViewId {
	return GroupIdCalendarCalendarViewId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupIdCalendarCalendarViewID parses 'input' into a GroupIdCalendarCalendarViewId
func ParseGroupIdCalendarCalendarViewID(input string) (*GroupIdCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarViewIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarViewIDInsensitively(input string) (*GroupIdCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateGroupIdCalendarCalendarViewID checks that 'input' can be parsed as a Group Id Calendar Calendar View ID
func ValidateGroupIdCalendarCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar View ID
func (id GroupIdCalendarCalendarViewId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar View ID
func (id GroupIdCalendarCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar View ID
func (id GroupIdCalendarCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Id Calendar Calendar View (%s)", strings.Join(components, "\n"))
}
