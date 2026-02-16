package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerTaskId{}

// UserIdPlannerTaskId is a struct representing the Resource ID for a User Id Planner Task
type UserIdPlannerTaskId struct {
	UserId        string
	PlannerTaskId string
}

// NewUserIdPlannerTaskID returns a new UserIdPlannerTaskId struct
func NewUserIdPlannerTaskID(userId string, plannerTaskId string) UserIdPlannerTaskId {
	return UserIdPlannerTaskId{
		UserId:        userId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseUserIdPlannerTaskID parses 'input' into a UserIdPlannerTaskId
func ParseUserIdPlannerTaskID(input string) (*UserIdPlannerTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerTaskIDInsensitively parses 'input' case-insensitively into a UserIdPlannerTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerTaskIDInsensitively(input string) (*UserIdPlannerTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateUserIdPlannerTaskID checks that 'input' can be parsed as a User Id Planner Task ID
func ValidateUserIdPlannerTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Task ID
func (id UserIdPlannerTaskId) ID() string {
	fmtString := "/users/%s/planner/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Task ID
func (id UserIdPlannerTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("tasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this User Id Planner Task ID
func (id UserIdPlannerTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Id Planner Task (%s)", strings.Join(components, "\n"))
}
