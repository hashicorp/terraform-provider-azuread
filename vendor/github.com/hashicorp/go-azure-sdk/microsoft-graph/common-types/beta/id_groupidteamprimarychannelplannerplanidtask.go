package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelPlannerPlanIdTaskId{}

// GroupIdTeamPrimaryChannelPlannerPlanIdTaskId is a struct representing the Resource ID for a Group Id Team Primary Channel Planner Plan Id Task
type GroupIdTeamPrimaryChannelPlannerPlanIdTaskId struct {
	GroupId       string
	PlannerPlanId string
	PlannerTaskId string
}

// NewGroupIdTeamPrimaryChannelPlannerPlanIdTaskID returns a new GroupIdTeamPrimaryChannelPlannerPlanIdTaskId struct
func NewGroupIdTeamPrimaryChannelPlannerPlanIdTaskID(groupId string, plannerPlanId string, plannerTaskId string) GroupIdTeamPrimaryChannelPlannerPlanIdTaskId {
	return GroupIdTeamPrimaryChannelPlannerPlanIdTaskId{
		GroupId:       groupId,
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIdTaskID parses 'input' into a GroupIdTeamPrimaryChannelPlannerPlanIdTaskId
func ParseGroupIdTeamPrimaryChannelPlannerPlanIdTaskID(input string) (*GroupIdTeamPrimaryChannelPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIdTaskIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelPlannerPlanIdTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelPlannerPlanIdTaskIDInsensitively(input string) (*GroupIdTeamPrimaryChannelPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelPlannerPlanIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelPlannerPlanIdTaskID checks that 'input' can be parsed as a Group Id Team Primary Channel Planner Plan Id Task ID
func ValidateGroupIdTeamPrimaryChannelPlannerPlanIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelPlannerPlanIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Planner Plan Id Task ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdTaskId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Planner Plan Id Task ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Planner Plan Id Task ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Planner Plan Id Task (%s)", strings.Join(components, "\n"))
}
