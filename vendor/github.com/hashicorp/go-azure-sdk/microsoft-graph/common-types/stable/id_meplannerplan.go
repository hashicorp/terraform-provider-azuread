package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanId{}

// MePlannerPlanId is a struct representing the Resource ID for a Me Planner Plan
type MePlannerPlanId struct {
	PlannerPlanId string
}

// NewMePlannerPlanID returns a new MePlannerPlanId struct
func NewMePlannerPlanID(plannerPlanId string) MePlannerPlanId {
	return MePlannerPlanId{
		PlannerPlanId: plannerPlanId,
	}
}

// ParseMePlannerPlanID parses 'input' into a MePlannerPlanId
func ParseMePlannerPlanID(input string) (*MePlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerPlanIDInsensitively parses 'input' case-insensitively into a MePlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanIDInsensitively(input string) (*MePlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateMePlannerPlanID checks that 'input' can be parsed as a Me Planner Plan ID
func ValidateMePlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan ID
func (id MePlannerPlanId) ID() string {
	fmtString := "/me/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan ID
func (id MePlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Me Planner Plan ID
func (id MePlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Me Planner Plan (%s)", strings.Join(components, "\n"))
}
