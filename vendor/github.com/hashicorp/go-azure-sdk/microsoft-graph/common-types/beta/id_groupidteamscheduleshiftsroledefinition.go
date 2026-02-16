package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleShiftsRoleDefinitionId{}

// GroupIdTeamScheduleShiftsRoleDefinitionId is a struct representing the Resource ID for a Group Id Team Schedule Shifts Role Definition
type GroupIdTeamScheduleShiftsRoleDefinitionId struct {
	GroupId                string
	ShiftsRoleDefinitionId string
}

// NewGroupIdTeamScheduleShiftsRoleDefinitionID returns a new GroupIdTeamScheduleShiftsRoleDefinitionId struct
func NewGroupIdTeamScheduleShiftsRoleDefinitionID(groupId string, shiftsRoleDefinitionId string) GroupIdTeamScheduleShiftsRoleDefinitionId {
	return GroupIdTeamScheduleShiftsRoleDefinitionId{
		GroupId:                groupId,
		ShiftsRoleDefinitionId: shiftsRoleDefinitionId,
	}
}

// ParseGroupIdTeamScheduleShiftsRoleDefinitionID parses 'input' into a GroupIdTeamScheduleShiftsRoleDefinitionId
func ParseGroupIdTeamScheduleShiftsRoleDefinitionID(input string) (*GroupIdTeamScheduleShiftsRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleShiftsRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleShiftsRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleShiftsRoleDefinitionIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleShiftsRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleShiftsRoleDefinitionIDInsensitively(input string) (*GroupIdTeamScheduleShiftsRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleShiftsRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleShiftsRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleShiftsRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ShiftsRoleDefinitionId, ok = input.Parsed["shiftsRoleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "shiftsRoleDefinitionId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleShiftsRoleDefinitionID checks that 'input' can be parsed as a Group Id Team Schedule Shifts Role Definition ID
func ValidateGroupIdTeamScheduleShiftsRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleShiftsRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Shifts Role Definition ID
func (id GroupIdTeamScheduleShiftsRoleDefinitionId) ID() string {
	fmtString := "/groups/%s/team/schedule/shiftsRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ShiftsRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Shifts Role Definition ID
func (id GroupIdTeamScheduleShiftsRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("shiftsRoleDefinitions", "shiftsRoleDefinitions", "shiftsRoleDefinitions"),
		resourceids.UserSpecifiedSegment("shiftsRoleDefinitionId", "shiftsRoleDefinitionId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Shifts Role Definition ID
func (id GroupIdTeamScheduleShiftsRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shifts Role Definition: %q", id.ShiftsRoleDefinitionId),
	}
	return fmt.Sprintf("Group Id Team Schedule Shifts Role Definition (%s)", strings.Join(components, "\n"))
}
