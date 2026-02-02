package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerTaskId{}

// MePlannerTaskId is a struct representing the Resource ID for a Me Planner Task
type MePlannerTaskId struct {
	PlannerTaskId string
}

// NewMePlannerTaskID returns a new MePlannerTaskId struct
func NewMePlannerTaskID(plannerTaskId string) MePlannerTaskId {
	return MePlannerTaskId{
		PlannerTaskId: plannerTaskId,
	}
}

// ParseMePlannerTaskID parses 'input' into a MePlannerTaskId
func ParseMePlannerTaskID(input string) (*MePlannerTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerTaskIDInsensitively parses 'input' case-insensitively into a MePlannerTaskId
// note: this method should only be used for API response data and not user input
func ParseMePlannerTaskIDInsensitively(input string) (*MePlannerTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateMePlannerTaskID checks that 'input' can be parsed as a Me Planner Task ID
func ValidateMePlannerTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Task ID
func (id MePlannerTaskId) ID() string {
	fmtString := "/me/planner/tasks/%s"
	return fmt.Sprintf(fmtString, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Task ID
func (id MePlannerTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this Me Planner Task ID
func (id MePlannerTaskId) String() string {
	components := []string{
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Me Planner Task (%s)", strings.Join(components, "\n"))
}
