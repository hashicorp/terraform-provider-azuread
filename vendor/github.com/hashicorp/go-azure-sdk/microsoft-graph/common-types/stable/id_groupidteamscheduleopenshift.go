package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleOpenShiftId{}

// GroupIdTeamScheduleOpenShiftId is a struct representing the Resource ID for a Group Id Team Schedule Open Shift
type GroupIdTeamScheduleOpenShiftId struct {
	GroupId     string
	OpenShiftId string
}

// NewGroupIdTeamScheduleOpenShiftID returns a new GroupIdTeamScheduleOpenShiftId struct
func NewGroupIdTeamScheduleOpenShiftID(groupId string, openShiftId string) GroupIdTeamScheduleOpenShiftId {
	return GroupIdTeamScheduleOpenShiftId{
		GroupId:     groupId,
		OpenShiftId: openShiftId,
	}
}

// ParseGroupIdTeamScheduleOpenShiftID parses 'input' into a GroupIdTeamScheduleOpenShiftId
func ParseGroupIdTeamScheduleOpenShiftID(input string) (*GroupIdTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleOpenShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleOpenShiftIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleOpenShiftId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleOpenShiftIDInsensitively(input string) (*GroupIdTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleOpenShiftId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleOpenShiftId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.OpenShiftId, ok = input.Parsed["openShiftId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleOpenShiftID checks that 'input' can be parsed as a Group Id Team Schedule Open Shift ID
func ValidateGroupIdTeamScheduleOpenShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleOpenShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Open Shift ID
func (id GroupIdTeamScheduleOpenShiftId) ID() string {
	fmtString := "/groups/%s/team/schedule/openShifts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OpenShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Open Shift ID
func (id GroupIdTeamScheduleOpenShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("openShifts", "openShifts", "openShifts"),
		resourceids.UserSpecifiedSegment("openShiftId", "openShiftId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Open Shift ID
func (id GroupIdTeamScheduleOpenShiftId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Open Shift: %q", id.OpenShiftId),
	}
	return fmt.Sprintf("Group Id Team Schedule Open Shift (%s)", strings.Join(components, "\n"))
}
