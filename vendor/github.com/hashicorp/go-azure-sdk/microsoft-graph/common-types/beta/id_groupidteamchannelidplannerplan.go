package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdPlannerPlanId{}

// GroupIdTeamChannelIdPlannerPlanId is a struct representing the Resource ID for a Group Id Team Channel Id Planner Plan
type GroupIdTeamChannelIdPlannerPlanId struct {
	GroupId       string
	ChannelId     string
	PlannerPlanId string
}

// NewGroupIdTeamChannelIdPlannerPlanID returns a new GroupIdTeamChannelIdPlannerPlanId struct
func NewGroupIdTeamChannelIdPlannerPlanID(groupId string, channelId string, plannerPlanId string) GroupIdTeamChannelIdPlannerPlanId {
	return GroupIdTeamChannelIdPlannerPlanId{
		GroupId:       groupId,
		ChannelId:     channelId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseGroupIdTeamChannelIdPlannerPlanID parses 'input' into a GroupIdTeamChannelIdPlannerPlanId
func ParseGroupIdTeamChannelIdPlannerPlanID(input string) (*GroupIdTeamChannelIdPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdPlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdPlannerPlanIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdPlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdPlannerPlanIDInsensitively(input string) (*GroupIdTeamChannelIdPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdPlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdPlannerPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdPlannerPlanID checks that 'input' can be parsed as a Group Id Team Channel Id Planner Plan ID
func ValidateGroupIdTeamChannelIdPlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdPlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Planner Plan ID
func (id GroupIdTeamChannelIdPlannerPlanId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Planner Plan ID
func (id GroupIdTeamChannelIdPlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Planner Plan ID
func (id GroupIdTeamChannelIdPlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Planner Plan (%s)", strings.Join(components, "\n"))
}
