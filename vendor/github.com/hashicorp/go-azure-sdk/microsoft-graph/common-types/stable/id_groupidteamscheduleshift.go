package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleShiftId{}

// GroupIdTeamScheduleShiftId is a struct representing the Resource ID for a Group Id Team Schedule Shift
type GroupIdTeamScheduleShiftId struct {
	GroupId string
	ShiftId string
}

// NewGroupIdTeamScheduleShiftID returns a new GroupIdTeamScheduleShiftId struct
func NewGroupIdTeamScheduleShiftID(groupId string, shiftId string) GroupIdTeamScheduleShiftId {
	return GroupIdTeamScheduleShiftId{
		GroupId: groupId,
		ShiftId: shiftId,
	}
}

// ParseGroupIdTeamScheduleShiftID parses 'input' into a GroupIdTeamScheduleShiftId
func ParseGroupIdTeamScheduleShiftID(input string) (*GroupIdTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleShiftIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleShiftId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleShiftIDInsensitively(input string) (*GroupIdTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleShiftId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ShiftId, ok = input.Parsed["shiftId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "shiftId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleShiftID checks that 'input' can be parsed as a Group Id Team Schedule Shift ID
func ValidateGroupIdTeamScheduleShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Shift ID
func (id GroupIdTeamScheduleShiftId) ID() string {
	fmtString := "/groups/%s/team/schedule/shifts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Shift ID
func (id GroupIdTeamScheduleShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("shifts", "shifts", "shifts"),
		resourceids.UserSpecifiedSegment("shiftId", "shiftId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Shift ID
func (id GroupIdTeamScheduleShiftId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shift: %q", id.ShiftId),
	}
	return fmt.Sprintf("Group Id Team Schedule Shift (%s)", strings.Join(components, "\n"))
}
