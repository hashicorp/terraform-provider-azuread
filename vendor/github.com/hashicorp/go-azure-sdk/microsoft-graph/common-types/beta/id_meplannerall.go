package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerAllId{}

// MePlannerAllId is a struct representing the Resource ID for a Me Planner All
type MePlannerAllId struct {
	PlannerDeltaId string
}

// NewMePlannerAllID returns a new MePlannerAllId struct
func NewMePlannerAllID(plannerDeltaId string) MePlannerAllId {
	return MePlannerAllId{
		PlannerDeltaId: plannerDeltaId,
	}
}

// ParseMePlannerAllID parses 'input' into a MePlannerAllId
func ParseMePlannerAllID(input string) (*MePlannerAllId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerAllId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerAllId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerAllIDInsensitively parses 'input' case-insensitively into a MePlannerAllId
// note: this method should only be used for API response data and not user input
func ParseMePlannerAllIDInsensitively(input string) (*MePlannerAllId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerAllId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerAllId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerAllId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerDeltaId, ok = input.Parsed["plannerDeltaId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerDeltaId", input)
	}

	return nil
}

// ValidateMePlannerAllID checks that 'input' can be parsed as a Me Planner All ID
func ValidateMePlannerAllID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerAllID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner All ID
func (id MePlannerAllId) ID() string {
	fmtString := "/me/planner/all/%s"
	return fmt.Sprintf(fmtString, id.PlannerDeltaId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner All ID
func (id MePlannerAllId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("all", "all", "all"),
		resourceids.UserSpecifiedSegment("plannerDeltaId", "plannerDeltaId"),
	}
}

// String returns a human-readable description of this Me Planner All ID
func (id MePlannerAllId) String() string {
	components := []string{
		fmt.Sprintf("Planner Delta: %q", id.PlannerDeltaId),
	}
	return fmt.Sprintf("Me Planner All (%s)", strings.Join(components, "\n"))
}
