package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleSchedulingGroupId{}

// GroupIdTeamScheduleSchedulingGroupId is a struct representing the Resource ID for a Group Id Team Schedule Scheduling Group
type GroupIdTeamScheduleSchedulingGroupId struct {
	GroupId           string
	SchedulingGroupId string
}

// NewGroupIdTeamScheduleSchedulingGroupID returns a new GroupIdTeamScheduleSchedulingGroupId struct
func NewGroupIdTeamScheduleSchedulingGroupID(groupId string, schedulingGroupId string) GroupIdTeamScheduleSchedulingGroupId {
	return GroupIdTeamScheduleSchedulingGroupId{
		GroupId:           groupId,
		SchedulingGroupId: schedulingGroupId,
	}
}

// ParseGroupIdTeamScheduleSchedulingGroupID parses 'input' into a GroupIdTeamScheduleSchedulingGroupId
func ParseGroupIdTeamScheduleSchedulingGroupID(input string) (*GroupIdTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleSchedulingGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamScheduleSchedulingGroupIDInsensitively parses 'input' case-insensitively into a GroupIdTeamScheduleSchedulingGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamScheduleSchedulingGroupIDInsensitively(input string) (*GroupIdTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamScheduleSchedulingGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamScheduleSchedulingGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SchedulingGroupId, ok = input.Parsed["schedulingGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", input)
	}

	return nil
}

// ValidateGroupIdTeamScheduleSchedulingGroupID checks that 'input' can be parsed as a Group Id Team Schedule Scheduling Group ID
func ValidateGroupIdTeamScheduleSchedulingGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamScheduleSchedulingGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Schedule Scheduling Group ID
func (id GroupIdTeamScheduleSchedulingGroupId) ID() string {
	fmtString := "/groups/%s/team/schedule/schedulingGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SchedulingGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Schedule Scheduling Group ID
func (id GroupIdTeamScheduleSchedulingGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("schedule", "schedule", "schedule"),
		resourceids.StaticSegment("schedulingGroups", "schedulingGroups", "schedulingGroups"),
		resourceids.UserSpecifiedSegment("schedulingGroupId", "schedulingGroupId"),
	}
}

// String returns a human-readable description of this Group Id Team Schedule Scheduling Group ID
func (id GroupIdTeamScheduleSchedulingGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Scheduling Group: %q", id.SchedulingGroupId),
	}
	return fmt.Sprintf("Group Id Team Schedule Scheduling Group (%s)", strings.Join(components, "\n"))
}
