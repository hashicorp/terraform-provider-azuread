package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdInstanceId{}

// GroupIdCalendarViewIdInstanceId is a struct representing the Resource ID for a Group Id Calendar View Id Instance
type GroupIdCalendarViewIdInstanceId struct {
	GroupId  string
	EventId  string
	EventId1 string
}

// NewGroupIdCalendarViewIdInstanceID returns a new GroupIdCalendarViewIdInstanceId struct
func NewGroupIdCalendarViewIdInstanceID(groupId string, eventId string, eventId1 string) GroupIdCalendarViewIdInstanceId {
	return GroupIdCalendarViewIdInstanceId{
		GroupId:  groupId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseGroupIdCalendarViewIdInstanceID parses 'input' into a GroupIdCalendarViewIdInstanceId
func ParseGroupIdCalendarViewIdInstanceID(input string) (*GroupIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarViewIdInstanceIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarViewIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarViewIdInstanceIDInsensitively(input string) (*GroupIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarViewIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdCalendarViewIdInstanceID checks that 'input' can be parsed as a Group Id Calendar View Id Instance ID
func ValidateGroupIdCalendarViewIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarViewIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar View Id Instance ID
func (id GroupIdCalendarViewIdInstanceId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar View Id Instance ID
func (id GroupIdCalendarViewIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
	}
}

// String returns a human-readable description of this Group Id Calendar View Id Instance ID
func (id GroupIdCalendarViewIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Group Id Calendar View Id Instance (%s)", strings.Join(components, "\n"))
}
