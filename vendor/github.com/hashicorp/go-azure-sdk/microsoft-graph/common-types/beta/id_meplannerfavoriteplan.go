package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerFavoritePlanId{}

// MePlannerFavoritePlanId is a struct representing the Resource ID for a Me Planner Favorite Plan
type MePlannerFavoritePlanId struct {
	PlannerPlanId string
}

// NewMePlannerFavoritePlanID returns a new MePlannerFavoritePlanId struct
func NewMePlannerFavoritePlanID(plannerPlanId string) MePlannerFavoritePlanId {
	return MePlannerFavoritePlanId{
		PlannerPlanId: plannerPlanId,
	}
}

// ParseMePlannerFavoritePlanID parses 'input' into a MePlannerFavoritePlanId
func ParseMePlannerFavoritePlanID(input string) (*MePlannerFavoritePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerFavoritePlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerFavoritePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerFavoritePlanIDInsensitively parses 'input' case-insensitively into a MePlannerFavoritePlanId
// note: this method should only be used for API response data and not user input
func ParseMePlannerFavoritePlanIDInsensitively(input string) (*MePlannerFavoritePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerFavoritePlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerFavoritePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerFavoritePlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateMePlannerFavoritePlanID checks that 'input' can be parsed as a Me Planner Favorite Plan ID
func ValidateMePlannerFavoritePlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerFavoritePlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Favorite Plan ID
func (id MePlannerFavoritePlanId) ID() string {
	fmtString := "/me/planner/favoritePlans/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Favorite Plan ID
func (id MePlannerFavoritePlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("favoritePlans", "favoritePlans", "favoritePlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Me Planner Favorite Plan ID
func (id MePlannerFavoritePlanId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Me Planner Favorite Plan (%s)", strings.Join(components, "\n"))
}
