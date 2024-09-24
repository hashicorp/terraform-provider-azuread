package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerPlanId{}

// UserIdPlannerPlanId is a struct representing the Resource ID for a User Id Planner Plan
type UserIdPlannerPlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserIdPlannerPlanID returns a new UserIdPlannerPlanId struct
func NewUserIdPlannerPlanID(userId string, plannerPlanId string) UserIdPlannerPlanId {
	return UserIdPlannerPlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserIdPlannerPlanID parses 'input' into a UserIdPlannerPlanId
func ParseUserIdPlannerPlanID(input string) (*UserIdPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerPlanIDInsensitively parses 'input' case-insensitively into a UserIdPlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerPlanIDInsensitively(input string) (*UserIdPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateUserIdPlannerPlanID checks that 'input' can be parsed as a User Id Planner Plan ID
func ValidateUserIdPlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Plan ID
func (id UserIdPlannerPlanId) ID() string {
	fmtString := "/users/%s/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Plan ID
func (id UserIdPlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("plans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this User Id Planner Plan ID
func (id UserIdPlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Id Planner Plan (%s)", strings.Join(components, "\n"))
}
