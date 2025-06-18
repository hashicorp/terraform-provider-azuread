package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelPlannerPlanIdBucketId{}

// GroupIdTeamPrimaryChannelPlannerPlanIdBucketId is a struct representing the Resource ID for a Group Id Team Primary Channel Planner Plan Id Bucket
type GroupIdTeamPrimaryChannelPlannerPlanIdBucketId struct {
	GroupId         string
	PlannerPlanId   string
	PlannerBucketId string
}

// NewGroupIdTeamPrimaryChannelPlannerPlanIdBucketID returns a new GroupIdTeamPrimaryChannelPlannerPlanIdBucketId struct
func NewGroupIdTeamPrimaryChannelPlannerPlanIdBucketID(groupId string, plannerPlanId string, plannerBucketId string) GroupIdTeamPrimaryChannelPlannerPlanIdBucketId {
	return GroupIdTeamPrimaryChannelPlannerPlanIdBucketId{
		GroupId:         groupId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketID parses 'input' into a GroupIdTeamPrimaryChannelPlannerPlanIdBucketId
func ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketID(input string) (*GroupIdTeamPrimaryChannelPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelPlannerPlanIdBucketId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketIDInsensitively(input string) (*GroupIdTeamPrimaryChannelPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelPlannerPlanIdBucketId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerBucketId, ok = input.Parsed["plannerBucketId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelPlannerPlanIdBucketID checks that 'input' can be parsed as a Group Id Team Primary Channel Planner Plan Id Bucket ID
func ValidateGroupIdTeamPrimaryChannelPlannerPlanIdBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelPlannerPlanIdBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Planner Plan Id Bucket ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdBucketId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Planner Plan Id Bucket ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Planner Plan Id Bucket ID
func (id GroupIdTeamPrimaryChannelPlannerPlanIdBucketId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Planner Plan Id Bucket (%s)", strings.Join(components, "\n"))
}
