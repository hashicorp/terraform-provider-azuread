package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelPlannerPlanId{}

// GroupIdTeamPrimaryChannelPlannerPlanId is a struct representing the Resource ID for a Group Id Team Primary Channel Planner Plan
type GroupIdTeamPrimaryChannelPlannerPlanId struct {
	GroupId       string
	PlannerPlanId string
}

// NewGroupIdTeamPrimaryChannelPlannerPlanID returns a new GroupIdTeamPrimaryChannelPlannerPlanId struct
func NewGroupIdTeamPrimaryChannelPlannerPlanID(groupId string, plannerPlanId string) GroupIdTeamPrimaryChannelPlannerPlanId {
	return GroupIdTeamPrimaryChannelPlannerPlanId{
		GroupId:       groupId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanID parses 'input' into a GroupIdTeamPrimaryChannelPlannerPlanId
func ParseGroupIdTeamPrimaryChannelPlannerPlanID(input string) (*GroupIdTeamPrimaryChannelPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelPlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelPlannerPlanIDInsensitively(input string) (*GroupIdTeamPrimaryChannelPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelPlannerPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelPlannerPlanID checks that 'input' can be parsed as a Group Id Team Primary Channel Planner Plan ID
func ValidateGroupIdTeamPrimaryChannelPlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelPlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Planner Plan ID
func (id GroupIdTeamPrimaryChannelPlannerPlanId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Planner Plan ID
func (id GroupIdTeamPrimaryChannelPlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Planner Plan ID
func (id GroupIdTeamPrimaryChannelPlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Planner Plan (%s)", strings.Join(components, "\n"))
}
