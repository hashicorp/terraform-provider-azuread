package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerPlanIdBucketIdTaskId{}

// UserIdPlannerPlanIdBucketIdTaskId is a struct representing the Resource ID for a User Id Planner Plan Id Bucket Id Task
type UserIdPlannerPlanIdBucketIdTaskId struct {
	UserId          string
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewUserIdPlannerPlanIdBucketIdTaskID returns a new UserIdPlannerPlanIdBucketIdTaskId struct
func NewUserIdPlannerPlanIdBucketIdTaskID(userId string, plannerPlanId string, plannerBucketId string, plannerTaskId string) UserIdPlannerPlanIdBucketIdTaskId {
	return UserIdPlannerPlanIdBucketIdTaskId{
		UserId:          userId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseUserIdPlannerPlanIdBucketIdTaskID parses 'input' into a UserIdPlannerPlanIdBucketIdTaskId
func ParseUserIdPlannerPlanIdBucketIdTaskID(input string) (*UserIdPlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerPlanIdBucketIdTaskIDInsensitively parses 'input' case-insensitively into a UserIdPlannerPlanIdBucketIdTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerPlanIdBucketIdTaskIDInsensitively(input string) (*UserIdPlannerPlanIdBucketIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanIdBucketIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanIdBucketIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerPlanIdBucketIdTaskId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateUserIdPlannerPlanIdBucketIdTaskID checks that 'input' can be parsed as a User Id Planner Plan Id Bucket Id Task ID
func ValidateUserIdPlannerPlanIdBucketIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerPlanIdBucketIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Plan Id Bucket Id Task ID
func (id UserIdPlannerPlanIdBucketIdTaskId) ID() string {
	fmtString := "/users/%s/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Plan Id Bucket Id Task ID
func (id UserIdPlannerPlanIdBucketIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("buckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this User Id Planner Plan Id Bucket Id Task ID
func (id UserIdPlannerPlanIdBucketIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Id Planner Plan Id Bucket Id Task (%s)", strings.Join(components, "\n"))
}
