package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanIdBucketIdTaskId{}

// MePlannerPlanIdBucketIdTaskId is a struct representing the Resource ID for a Me Planner Plan Id Bucket Id Task
type MePlannerPlanIdBucketIdTaskId struct {
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewMePlannerPlanIdBucketIdTaskID returns a new MePlannerPlanIdBucketIdTaskId struct
func NewMePlannerPlanIdBucketIdTaskID(plannerPlanId string, plannerBucketId string, plannerTaskId string) MePlannerPlanIdBucketIdTaskId {
	return MePlannerPlanIdBucketIdTaskId{
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseMePlannerPlanIdBucketIdTaskID parses 'input' into a MePlannerPlanIdBucketIdTaskId
func ParseMePlannerPlanIdBucketIdTaskID(input string) (*MePlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerPlanIdBucketIdTaskIDInsensitively parses 'input' case-insensitively into a MePlannerPlanIdBucketIdTaskId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanIdBucketIdTaskIDInsensitively(input string) (*MePlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerPlanIdBucketIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMePlannerPlanIdBucketIdTaskID checks that 'input' can be parsed as a Me Planner Plan Id Bucket Id Task ID
func ValidateMePlannerPlanIdBucketIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanIdBucketIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan Id Bucket Id Task ID
func (id MePlannerPlanIdBucketIdTaskId) ID() string {
	fmtString := "/me/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan Id Bucket Id Task ID
func (id MePlannerPlanIdBucketIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Me Planner Plan Id Bucket Id Task ID
func (id MePlannerPlanIdBucketIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Me Planner Plan Id Bucket Id Task (%s)", strings.Join(components, "\n"))
}
