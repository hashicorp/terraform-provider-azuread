package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId{}

// GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId is a struct representing the Resource ID for a Group Id Team Primary Channel Planner Plan Id Bucket Id Task
type GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId struct {
	GroupId         string
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID returns a new GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId struct
func NewGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID(groupId string, plannerPlanId string, plannerBucketId string, plannerTaskId string) GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId {
	return GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId{
		GroupId:         groupId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID parses 'input' into a GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId
func ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID(input string) (*GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskIDInsensitively(input string) (*GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerBucketId, ok = input.Parsed["plannerBucketId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID checks that 'input' can be parsed as a Group Id Team Primary Channel Planner Plan Id Bucket Id Task ID
func ValidateGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Planner Plan Id Bucket Id Task ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Planner Plan Id Bucket Id Task ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Planner Plan Id Bucket Id Task ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Planner Plan Id Bucket Id Task (%s)", strings.Join(components, "\n"))
}
