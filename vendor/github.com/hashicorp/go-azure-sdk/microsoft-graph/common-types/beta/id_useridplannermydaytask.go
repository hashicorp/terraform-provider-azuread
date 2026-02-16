package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerMyDayTaskId{}

// UserIdPlannerMyDayTaskId is a struct representing the Resource ID for a User Id Planner My Day Task
type UserIdPlannerMyDayTaskId struct {
	UserId        string
	PlannerTaskId string
}

// NewUserIdPlannerMyDayTaskID returns a new UserIdPlannerMyDayTaskId struct
func NewUserIdPlannerMyDayTaskID(userId string, plannerTaskId string) UserIdPlannerMyDayTaskId {
	return UserIdPlannerMyDayTaskId{
		UserId:        userId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseUserIdPlannerMyDayTaskID parses 'input' into a UserIdPlannerMyDayTaskId
func ParseUserIdPlannerMyDayTaskID(input string) (*UserIdPlannerMyDayTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerMyDayTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerMyDayTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerMyDayTaskIDInsensitively parses 'input' case-insensitively into a UserIdPlannerMyDayTaskId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerMyDayTaskIDInsensitively(input string) (*UserIdPlannerMyDayTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerMyDayTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerMyDayTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerMyDayTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerTaskId, ok = input.Parsed["plannerTaskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", input)
	}

	return nil
}

// ValidateUserIdPlannerMyDayTaskID checks that 'input' can be parsed as a User Id Planner My Day Task ID
func ValidateUserIdPlannerMyDayTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerMyDayTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner My Day Task ID
func (id UserIdPlannerMyDayTaskId) ID() string {
	fmtString := "/users/%s/planner/myDayTasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner My Day Task ID
func (id UserIdPlannerMyDayTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("myDayTasks", "myDayTasks", "myDayTasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskId"),
	}
}

// String returns a human-readable description of this User Id Planner My Day Task ID
func (id UserIdPlannerMyDayTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Id Planner My Day Task (%s)", strings.Join(components, "\n"))
}
