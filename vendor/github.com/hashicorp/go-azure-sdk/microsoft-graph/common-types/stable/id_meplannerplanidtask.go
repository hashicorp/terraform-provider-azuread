package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanIdTaskId{}

// MePlannerPlanIdTaskId is a struct representing the Resource ID for a Me Planner Plan Id Task
type MePlannerPlanIdTaskId struct {
	PlannerPlanId string
	PlannerTaskId string
}

// NewMePlannerPlanIdTaskID returns a new MePlannerPlanIdTaskId struct
func NewMePlannerPlanIdTaskID(plannerPlanId string, plannerTaskId string) MePlannerPlanIdTaskId {
	return MePlannerPlanIdTaskId{
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseMePlannerPlanIdTaskID parses 'input' into a MePlannerPlanIdTaskId
func ParseMePlannerPlanIdTaskID(input string) (*MePlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerPlanIdTaskIDInsensitively parses 'input' case-insensitively into a MePlannerPlanIdTaskId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanIdTaskIDInsensitively(input string) (*MePlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerPlanIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateMePlannerPlanIdTaskID checks that 'input' can be parsed as a Me Planner Plan Id Task ID
func ValidateMePlannerPlanIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan Id Task ID
func (id MePlannerPlanIdTaskId) ID() string {
	fmtString := "/me/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan Id Task ID
func (id MePlannerPlanIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Me Planner Plan Id Task ID
func (id MePlannerPlanIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Me Planner Plan Id Task (%s)", strings.Join(components, "\n"))
}
