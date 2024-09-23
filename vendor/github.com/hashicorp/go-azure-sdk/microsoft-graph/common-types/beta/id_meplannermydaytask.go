package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerMyDayTaskId{}

// MePlannerMyDayTaskId is a struct representing the Resource ID for a Me Planner My Day Task
type MePlannerMyDayTaskId struct {
	PlannerTaskId string
}

// NewMePlannerMyDayTaskID returns a new MePlannerMyDayTaskId struct
func NewMePlannerMyDayTaskID(plannerTaskId string) MePlannerMyDayTaskId {
	return MePlannerMyDayTaskId{
		PlannerTaskId: plannerTaskId,
	}
}

// ParseMePlannerMyDayTaskID parses 'input' into a MePlannerMyDayTaskId
func ParseMePlannerMyDayTaskID(input string) (*MePlannerMyDayTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerMyDayTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerMyDayTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerMyDayTaskIDInsensitively parses 'input' case-insensitively into a MePlannerMyDayTaskId
// note: this method should only be used for API response data and not user input
func ParseMePlannerMyDayTaskIDInsensitively(input string) (*MePlannerMyDayTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerMyDayTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerMyDayTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerMyDayTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateMePlannerMyDayTaskID checks that 'input' can be parsed as a Me Planner My Day Task ID
func ValidateMePlannerMyDayTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerMyDayTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner My Day Task ID
func (id MePlannerMyDayTaskId) ID() string {
	fmtString := "/me/planner/myDayTasks/%s"
	return fmt.Sprintf(fmtString, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner My Day Task ID
func (id MePlannerMyDayTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("myDayTasks", "myDayTasks", "myDayTasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Me Planner My Day Task ID
func (id MePlannerMyDayTaskId) String() string {
	components := []string{
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Me Planner My Day Task (%s)", strings.Join(components, "\n"))
}
