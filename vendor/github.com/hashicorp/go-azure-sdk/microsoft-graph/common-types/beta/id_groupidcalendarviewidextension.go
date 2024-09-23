package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdExtensionId{}

// GroupIdCalendarViewIdExtensionId is a struct representing the Resource ID for a Group Id Calendar View Id Extension
type GroupIdCalendarViewIdExtensionId struct {
	GroupId     string
	EventId     string
	ExtensionId string
}

// NewGroupIdCalendarViewIdExtensionID returns a new GroupIdCalendarViewIdExtensionId struct
func NewGroupIdCalendarViewIdExtensionID(groupId string, eventId string, extensionId string) GroupIdCalendarViewIdExtensionId {
	return GroupIdCalendarViewIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdCalendarViewIdExtensionID parses 'input' into a GroupIdCalendarViewIdExtensionId
func ParseGroupIdCalendarViewIdExtensionID(input string) (*GroupIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarViewIdExtensionIDInsensitively(input string) (*GroupIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdCalendarViewIdExtensionID checks that 'input' can be parsed as a Group Id Calendar View Id Extension ID
func ValidateGroupIdCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar View Id Extension ID
func (id GroupIdCalendarViewIdExtensionId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar View Id Extension ID
func (id GroupIdCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Calendar View Id Extension ID
func (id GroupIdCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
