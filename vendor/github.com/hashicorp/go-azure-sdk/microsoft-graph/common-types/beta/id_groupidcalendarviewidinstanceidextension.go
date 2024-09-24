package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdInstanceIdExtensionId{}

// GroupIdCalendarViewIdInstanceIdExtensionId is a struct representing the Resource ID for a Group Id Calendar View Id Instance Id Extension
type GroupIdCalendarViewIdInstanceIdExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupIdCalendarViewIdInstanceIdExtensionID returns a new GroupIdCalendarViewIdInstanceIdExtensionId struct
func NewGroupIdCalendarViewIdInstanceIdExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupIdCalendarViewIdInstanceIdExtensionId {
	return GroupIdCalendarViewIdInstanceIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdCalendarViewIdInstanceIdExtensionID parses 'input' into a GroupIdCalendarViewIdInstanceIdExtensionId
func ParseGroupIdCalendarViewIdInstanceIdExtensionID(input string) (*GroupIdCalendarViewIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarViewIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarViewIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarViewIdInstanceIdExtensionIDInsensitively(input string) (*GroupIdCalendarViewIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarViewIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdCalendarViewIdInstanceIdExtensionID checks that 'input' can be parsed as a Group Id Calendar View Id Instance Id Extension ID
func ValidateGroupIdCalendarViewIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarViewIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar View Id Instance Id Extension ID
func (id GroupIdCalendarViewIdInstanceIdExtensionId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar View Id Instance Id Extension ID
func (id GroupIdCalendarViewIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Calendar View Id Instance Id Extension ID
func (id GroupIdCalendarViewIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Calendar View Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
