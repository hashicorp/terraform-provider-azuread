package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewId{}

// GroupIdCalendarViewId is a struct representing the Resource ID for a Group Id Calendar View
type GroupIdCalendarViewId struct {
	GroupId string
	EventId string
}

// NewGroupIdCalendarViewID returns a new GroupIdCalendarViewId struct
func NewGroupIdCalendarViewID(groupId string, eventId string) GroupIdCalendarViewId {
	return GroupIdCalendarViewId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupIdCalendarViewID parses 'input' into a GroupIdCalendarViewId
func ParseGroupIdCalendarViewID(input string) (*GroupIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarViewIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarViewIDInsensitively(input string) (*GroupIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateGroupIdCalendarViewID checks that 'input' can be parsed as a Group Id Calendar View ID
func ValidateGroupIdCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar View ID
func (id GroupIdCalendarViewId) ID() string {
	fmtString := "/groups/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar View ID
func (id GroupIdCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
	}
}

// String returns a human-readable description of this Group Id Calendar View ID
func (id GroupIdCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Id Calendar View (%s)", strings.Join(components, "\n"))
}
