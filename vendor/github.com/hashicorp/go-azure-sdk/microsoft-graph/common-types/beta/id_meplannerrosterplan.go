package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerRosterPlanId{}

// MePlannerRosterPlanId is a struct representing the Resource ID for a Me Planner Roster Plan
type MePlannerRosterPlanId struct {
	PlannerPlanId string
}

// NewMePlannerRosterPlanID returns a new MePlannerRosterPlanId struct
func NewMePlannerRosterPlanID(plannerPlanId string) MePlannerRosterPlanId {
	return MePlannerRosterPlanId{
		PlannerPlanId: plannerPlanId,
	}
}

// ParseMePlannerRosterPlanID parses 'input' into a MePlannerRosterPlanId
func ParseMePlannerRosterPlanID(input string) (*MePlannerRosterPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerRosterPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerRosterPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerRosterPlanIDInsensitively parses 'input' case-insensitively into a MePlannerRosterPlanId
// note: this method should only be used for API response data and not user input
func ParseMePlannerRosterPlanIDInsensitively(input string) (*MePlannerRosterPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerRosterPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerRosterPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerRosterPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateMePlannerRosterPlanID checks that 'input' can be parsed as a Me Planner Roster Plan ID
func ValidateMePlannerRosterPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerRosterPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Roster Plan ID
func (id MePlannerRosterPlanId) ID() string {
	fmtString := "/me/planner/rosterPlans/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Roster Plan ID
func (id MePlannerRosterPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("rosterPlans", "rosterPlans", "rosterPlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Me Planner Roster Plan ID
func (id MePlannerRosterPlanId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Me Planner Roster Plan (%s)", strings.Join(components, "\n"))
}
