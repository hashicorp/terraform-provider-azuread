package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanIdBucketId{}

// GroupIdPlannerPlanIdBucketId is a struct representing the Resource ID for a Group Id Planner Plan Id Bucket
type GroupIdPlannerPlanIdBucketId struct {
	GroupId         string
	PlannerPlanId   string
	PlannerBucketId string
}

// NewGroupIdPlannerPlanIdBucketID returns a new GroupIdPlannerPlanIdBucketId struct
func NewGroupIdPlannerPlanIdBucketID(groupId string, plannerPlanId string, plannerBucketId string) GroupIdPlannerPlanIdBucketId {
	return GroupIdPlannerPlanIdBucketId{
		GroupId:         groupId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseGroupIdPlannerPlanIdBucketID parses 'input' into a GroupIdPlannerPlanIdBucketId
func ParseGroupIdPlannerPlanIdBucketID(input string) (*GroupIdPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdPlannerPlanIdBucketIDInsensitively parses 'input' case-insensitively into a GroupIdPlannerPlanIdBucketId
// note: this method should only be used for API response data and not user input
func ParseGroupIdPlannerPlanIdBucketIDInsensitively(input string) (*GroupIdPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdPlannerPlanIdBucketId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdPlannerPlanIdBucketID checks that 'input' can be parsed as a Group Id Planner Plan Id Bucket ID
func ValidateGroupIdPlannerPlanIdBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdPlannerPlanIdBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Planner Plan Id Bucket ID
func (id GroupIdPlannerPlanIdBucketId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Planner Plan Id Bucket ID
func (id GroupIdPlannerPlanIdBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
	}
}

// String returns a human-readable description of this Group Id Planner Plan Id Bucket ID
func (id GroupIdPlannerPlanIdBucketId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("Group Id Planner Plan Id Bucket (%s)", strings.Join(components, "\n"))
}
