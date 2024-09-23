package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerPlanIdBucketId{}

// UserIdPlannerPlanIdBucketId is a struct representing the Resource ID for a User Id Planner Plan Id Bucket
type UserIdPlannerPlanIdBucketId struct {
	UserId          string
	PlannerPlanId   string
	PlannerBucketId string
}

// NewUserIdPlannerPlanIdBucketID returns a new UserIdPlannerPlanIdBucketId struct
func NewUserIdPlannerPlanIdBucketID(userId string, plannerPlanId string, plannerBucketId string) UserIdPlannerPlanIdBucketId {
	return UserIdPlannerPlanIdBucketId{
		UserId:          userId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseUserIdPlannerPlanIdBucketID parses 'input' into a UserIdPlannerPlanIdBucketId
func ParseUserIdPlannerPlanIdBucketID(input string) (*UserIdPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerPlanIdBucketIDInsensitively parses 'input' case-insensitively into a UserIdPlannerPlanIdBucketId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerPlanIdBucketIDInsensitively(input string) (*UserIdPlannerPlanIdBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanIdBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanIdBucketId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerPlanIdBucketId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerBucketId, ok = input.Parsed["plannerBucketId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", input)
	}

	return nil
}

// ValidateUserIdPlannerPlanIdBucketID checks that 'input' can be parsed as a User Id Planner Plan Id Bucket ID
func ValidateUserIdPlannerPlanIdBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerPlanIdBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Plan Id Bucket ID
func (id UserIdPlannerPlanIdBucketId) ID() string {
	fmtString := "/users/%s/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Plan Id Bucket ID
func (id UserIdPlannerPlanIdBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
	}
}

// String returns a human-readable description of this User Id Planner Plan Id Bucket ID
func (id UserIdPlannerPlanIdBucketId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("User Id Planner Plan Id Bucket (%s)", strings.Join(components, "\n"))
}
