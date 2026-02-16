package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdPlannerPlanIdTaskId{}

// GroupIdTeamChannelIdPlannerPlanIdTaskId is a struct representing the Resource ID for a Group Id Team Channel Id Planner Plan Id Task
type GroupIdTeamChannelIdPlannerPlanIdTaskId struct {
	GroupId       string
	ChannelId     string
	PlannerPlanId string
	PlannerTaskId string
}

// NewGroupIdTeamChannelIdPlannerPlanIdTaskID returns a new GroupIdTeamChannelIdPlannerPlanIdTaskId struct
func NewGroupIdTeamChannelIdPlannerPlanIdTaskID(groupId string, channelId string, plannerPlanId string, plannerTaskId string) GroupIdTeamChannelIdPlannerPlanIdTaskId {
	return GroupIdTeamChannelIdPlannerPlanIdTaskId{
		GroupId:       groupId,
		ChannelId:     channelId,
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseGroupIdTeamChannelIdPlannerPlanIdTaskID parses 'input' into a GroupIdTeamChannelIdPlannerPlanIdTaskId
func ParseGroupIdTeamChannelIdPlannerPlanIdTaskID(input string) (*GroupIdTeamChannelIdPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdPlannerPlanIdTaskIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdPlannerPlanIdTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdPlannerPlanIdTaskIDInsensitively(input string) (*GroupIdTeamChannelIdPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdPlannerPlanIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdPlannerPlanIdTaskID checks that 'input' can be parsed as a Group Id Team Channel Id Planner Plan Id Task ID
func ValidateGroupIdTeamChannelIdPlannerPlanIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdPlannerPlanIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Planner Plan Id Task ID
func (id GroupIdTeamChannelIdPlannerPlanIdTaskId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Planner Plan Id Task ID
func (id GroupIdTeamChannelIdPlannerPlanIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Planner Plan Id Task ID
func (id GroupIdTeamChannelIdPlannerPlanIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Planner Plan Id Task (%s)", strings.Join(components, "\n"))
}
