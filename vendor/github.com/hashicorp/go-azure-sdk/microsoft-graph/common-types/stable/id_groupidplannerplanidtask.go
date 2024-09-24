package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanIdTaskId{}

// GroupIdPlannerPlanIdTaskId is a struct representing the Resource ID for a Group Id Planner Plan Id Task
type GroupIdPlannerPlanIdTaskId struct {
	GroupId       string
	PlannerPlanId string
	PlannerTaskId string
}

// NewGroupIdPlannerPlanIdTaskID returns a new GroupIdPlannerPlanIdTaskId struct
func NewGroupIdPlannerPlanIdTaskID(groupId string, plannerPlanId string, plannerTaskId string) GroupIdPlannerPlanIdTaskId {
	return GroupIdPlannerPlanIdTaskId{
		GroupId:       groupId,
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseGroupIdPlannerPlanIdTaskID parses 'input' into a GroupIdPlannerPlanIdTaskId
func ParseGroupIdPlannerPlanIdTaskID(input string) (*GroupIdPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdPlannerPlanIdTaskIDInsensitively parses 'input' case-insensitively into a GroupIdPlannerPlanIdTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupIdPlannerPlanIdTaskIDInsensitively(input string) (*GroupIdPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdPlannerPlanIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdPlannerPlanIdTaskID checks that 'input' can be parsed as a Group Id Planner Plan Id Task ID
func ValidateGroupIdPlannerPlanIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdPlannerPlanIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Planner Plan Id Task ID
func (id GroupIdPlannerPlanIdTaskId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Planner Plan Id Task ID
func (id GroupIdPlannerPlanIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Group Id Planner Plan Id Task ID
func (id GroupIdPlannerPlanIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Id Planner Plan Id Task (%s)", strings.Join(components, "\n"))
}
