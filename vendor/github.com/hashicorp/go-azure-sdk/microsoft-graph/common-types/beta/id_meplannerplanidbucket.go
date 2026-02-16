package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanIdBucketId{}

// MePlannerPlanIdBucketId is a struct representing the Resource ID for a Me Planner Plan Id Bucket
type MePlannerPlanIdBucketId struct {
	PlannerPlanId   string
	PlannerBucketId string
}

// NewMePlannerPlanIdBucketID returns a new MePlannerPlanIdBucketId struct
func NewMePlannerPlanIdBucketID(plannerPlanId string, plannerBucketId string) MePlannerPlanIdBucketId {
	return MePlannerPlanIdBucketId{
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseMePlannerPlanIdBucketID parses 'input' into a MePlannerPlanIdBucketId
func ParseMePlannerPlanIdBucketID(input string) (*MePlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePlannerPlanIdBucketIDInsensitively parses 'input' case-insensitively into a MePlannerPlanIdBucketId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanIdBucketIDInsensitively(input string) (*MePlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePlannerPlanIdBucketId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerBucketId, ok = input.Parsed["plannerBucketId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", input)
	}

	return nil
}

// ValidateMePlannerPlanIdBucketID checks that 'input' can be parsed as a Me Planner Plan Id Bucket ID
func ValidateMePlannerPlanIdBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanIdBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan Id Bucket ID
func (id MePlannerPlanIdBucketId) ID() string {
	fmtString := "/me/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan Id Bucket ID
func (id MePlannerPlanIdBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
	}
}

// String returns a human-readable description of this Me Planner Plan Id Bucket ID
func (id MePlannerPlanIdBucketId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("Me Planner Plan Id Bucket (%s)", strings.Join(components, "\n"))
}
