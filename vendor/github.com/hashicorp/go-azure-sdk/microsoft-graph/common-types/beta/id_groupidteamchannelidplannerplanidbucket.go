package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdPlannerPlanIdBucketId{}

// GroupIdTeamChannelIdPlannerPlanIdBucketId is a struct representing the Resource ID for a Group Id Team Channel Id Planner Plan Id Bucket
type GroupIdTeamChannelIdPlannerPlanIdBucketId struct {
	GroupId         string
	ChannelId       string
	PlannerPlanId   string
	PlannerBucketId string
}

// NewGroupIdTeamChannelIdPlannerPlanIdBucketID returns a new GroupIdTeamChannelIdPlannerPlanIdBucketId struct
func NewGroupIdTeamChannelIdPlannerPlanIdBucketID(groupId string, channelId string, plannerPlanId string, plannerBucketId string) GroupIdTeamChannelIdPlannerPlanIdBucketId {
	return GroupIdTeamChannelIdPlannerPlanIdBucketId{
		GroupId:         groupId,
		ChannelId:       channelId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseGroupIdTeamChannelIdPlannerPlanIdBucketID parses 'input' into a GroupIdTeamChannelIdPlannerPlanIdBucketId
func ParseGroupIdTeamChannelIdPlannerPlanIdBucketID(input string) (*GroupIdTeamChannelIdPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamChannelIdPlannerPlanIdBucketIDInsensitively parses 'input' case-insensitively into a GroupIdTeamChannelIdPlannerPlanIdBucketId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamChannelIdPlannerPlanIdBucketIDInsensitively(input string) (*GroupIdTeamChannelIdPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamChannelIdPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamChannelIdPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamChannelIdPlannerPlanIdBucketId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PlannerBucketId, ok = input.Parsed["plannerBucketId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", input)
	}

	return nil
}

// ValidateGroupIdTeamChannelIdPlannerPlanIdBucketID checks that 'input' can be parsed as a Group Id Team Channel Id Planner Plan Id Bucket ID
func ValidateGroupIdTeamChannelIdPlannerPlanIdBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamChannelIdPlannerPlanIdBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Channel Id Planner Plan Id Bucket ID
func (id GroupIdTeamChannelIdPlannerPlanIdBucketId) ID() string {
	fmtString := "/groups/%s/team/channels/%s/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChannelId, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Channel Id Planner Plan Id Bucket ID
func (id GroupIdTeamChannelIdPlannerPlanIdBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
	}
}

// String returns a human-readable description of this Group Id Team Channel Id Planner Plan Id Bucket ID
func (id GroupIdTeamChannelIdPlannerPlanIdBucketId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("Group Id Team Channel Id Planner Plan Id Bucket (%s)", strings.Join(components, "\n"))
}
