package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarPermissionId{}

// GroupIdCalendarCalendarPermissionId is a struct representing the Resource ID for a Group Id Calendar Calendar Permission
type GroupIdCalendarCalendarPermissionId struct {
	GroupId              string
	CalendarPermissionId string
}

// NewGroupIdCalendarCalendarPermissionID returns a new GroupIdCalendarCalendarPermissionId struct
func NewGroupIdCalendarCalendarPermissionID(groupId string, calendarPermissionId string) GroupIdCalendarCalendarPermissionId {
	return GroupIdCalendarCalendarPermissionId{
		GroupId:              groupId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseGroupIdCalendarCalendarPermissionID parses 'input' into a GroupIdCalendarCalendarPermissionId
func ParseGroupIdCalendarCalendarPermissionID(input string) (*GroupIdCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarPermissionIDInsensitively(input string) (*GroupIdCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.CalendarPermissionId, ok = input.Parsed["calendarPermissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", input)
	}

	return nil
}

// ValidateGroupIdCalendarCalendarPermissionID checks that 'input' can be parsed as a Group Id Calendar Calendar Permission ID
func ValidateGroupIdCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar Permission ID
func (id GroupIdCalendarCalendarPermissionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar Permission ID
func (id GroupIdCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar Permission ID
func (id GroupIdCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("Group Id Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
