package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanIdBucketIdTaskId{}

// GroupIdPlannerPlanIdBucketIdTaskId is a struct representing the Resource ID for a Group Id Planner Plan Id Bucket Id Task
type GroupIdPlannerPlanIdBucketIdTaskId struct {
	GroupId         string
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewGroupIdPlannerPlanIdBucketIdTaskID returns a new GroupIdPlannerPlanIdBucketIdTaskId struct
func NewGroupIdPlannerPlanIdBucketIdTaskID(groupId string, plannerPlanId string, plannerBucketId string, plannerTaskId string) GroupIdPlannerPlanIdBucketIdTaskId {
	return GroupIdPlannerPlanIdBucketIdTaskId{
		GroupId:         groupId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseGroupIdPlannerPlanIdBucketIdTaskID parses 'input' into a GroupIdPlannerPlanIdBucketIdTaskId
func ParseGroupIdPlannerPlanIdBucketIdTaskID(input string) (*GroupIdPlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdPlannerPlanIdBucketIdTaskIDInsensitively parses 'input' case-insensitively into a GroupIdPlannerPlanIdBucketIdTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupIdPlannerPlanIdBucketIdTaskIDInsensitively(input string) (*GroupIdPlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdPlannerPlanIdBucketIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdPlannerPlanIdBucketIdTaskID checks that 'input' can be parsed as a Group Id Planner Plan Id Bucket Id Task ID
func ValidateGroupIdPlannerPlanIdBucketIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdPlannerPlanIdBucketIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Planner Plan Id Bucket Id Task ID
func (id GroupIdPlannerPlanIdBucketIdTaskId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Planner Plan Id Bucket Id Task ID
func (id GroupIdPlannerPlanIdBucketIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Group Id Planner Plan Id Bucket Id Task ID
func (id GroupIdPlannerPlanIdBucketIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Id Planner Plan Id Bucket Id Task (%s)", strings.Join(components, "\n"))
}
