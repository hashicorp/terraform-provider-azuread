package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerPlanIdTaskId{}

// UserIdPlannerPlanIdTaskId is a struct representing the Resource ID for a User Id Planner Plan Id Task
type UserIdPlannerPlanIdTaskId struct {
	UserId        string
	PlannerPlanId string
	PlannerTaskId string
}

// NewUserIdPlannerPlanIdTaskID returns a new UserIdPlannerPlanIdTaskId struct
func NewUserIdPlannerPlanIdTaskID(userId string, plannerPlanId string, plannerTaskId string) UserIdPlannerPlanIdTaskId {
	return UserIdPlannerPlanIdTaskId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseUserIdPlannerPlanIdTaskID parses 'input' into a UserIdPlannerPlanIdTaskId
func ParseUserIdPlannerPlanIdTaskID(input string) (*UserIdPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerPlanIdTaskIDInsensitively parses 'input' case-insensitively into a UserIdPlannerPlanIdTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerPlanIdTaskIDInsensitively(input string) (*UserIdPlannerPlanIdTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanIdTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanIdTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerPlanIdTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateUserIdPlannerPlanIdTaskID checks that 'input' can be parsed as a User Id Planner Plan Id Task ID
func ValidateUserIdPlannerPlanIdTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerPlanIdTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Plan Id Task ID
func (id UserIdPlannerPlanIdTaskId) ID() string {
	fmtString := "/users/%s/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Plan Id Task ID
func (id UserIdPlannerPlanIdTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this User Id Planner Plan Id Task ID
func (id UserIdPlannerPlanIdTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Id Planner Plan Id Task (%s)", strings.Join(components, "\n"))
}
