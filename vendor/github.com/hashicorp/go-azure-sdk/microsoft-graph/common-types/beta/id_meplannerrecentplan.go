package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerRecentPlanId{}

// MePlannerRecentPlanId is a struct representing the Resource ID for a Me Planner Recent Plan
type MePlannerRecentPlanId struct {
	PlannerPlanId string
}

// NewMePlannerRecentPlanID returns a new MePlannerRecentPlanId struct
func NewMePlannerRecentPlanID(plannerPlanId string) MePlannerRecentPlanId {
	return MePlannerRecentPlanId{
		PlannerPlanId: plannerPlanId,
	}
}

// ParseMePlannerRecentPlanID parses 'input' into a MePlannerRecentPlanId
func ParseMePlannerRecentPlanID(input string) (*MePlannerRecentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerRecentPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerRecentPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerRecentPlanIDInsensitively parses 'input' case-insensitively into a MePlannerRecentPlanId
// note: this method should only be used for API response data and not user input
func ParseMePlannerRecentPlanIDInsensitively(input string) (*MePlannerRecentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerRecentPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerRecentPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerRecentPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateMePlannerRecentPlanID checks that 'input' can be parsed as a Me Planner Recent Plan ID
func ValidateMePlannerRecentPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerRecentPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Recent Plan ID
func (id MePlannerRecentPlanId) ID() string {
	fmtString := "/me/planner/recentPlans/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Recent Plan ID
func (id MePlannerRecentPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("recentPlans", "recentPlans", "recentPlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Me Planner Recent Plan ID
func (id MePlannerRecentPlanId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Me Planner Recent Plan (%s)", strings.Join(components, "\n"))
}
