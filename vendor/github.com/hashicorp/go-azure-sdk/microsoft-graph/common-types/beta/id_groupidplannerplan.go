package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanId{}

// GroupIdPlannerPlanId is a struct representing the Resource ID for a Group Id Planner Plan
type GroupIdPlannerPlanId struct {
	GroupId       string
	PlannerPlanId string
}

// NewGroupIdPlannerPlanID returns a new GroupIdPlannerPlanId struct
func NewGroupIdPlannerPlanID(groupId string, plannerPlanId string) GroupIdPlannerPlanId {
	return GroupIdPlannerPlanId{
		GroupId:       groupId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseGroupIdPlannerPlanID parses 'input' into a GroupIdPlannerPlanId
func ParseGroupIdPlannerPlanID(input string) (*GroupIdPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdPlannerPlanIDInsensitively parses 'input' case-insensitively into a GroupIdPlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseGroupIdPlannerPlanIDInsensitively(input string) (*GroupIdPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdPlannerPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateGroupIdPlannerPlanID checks that 'input' can be parsed as a Group Id Planner Plan ID
func ValidateGroupIdPlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdPlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Planner Plan ID
func (id GroupIdPlannerPlanId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Planner Plan ID
func (id GroupIdPlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Group Id Planner Plan ID
func (id GroupIdPlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Group Id Planner Plan (%s)", strings.Join(components, "\n"))
}
