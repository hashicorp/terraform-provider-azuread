package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerRecentPlanId{}

// UserIdPlannerRecentPlanId is a struct representing the Resource ID for a User Id Planner Recent Plan
type UserIdPlannerRecentPlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserIdPlannerRecentPlanID returns a new UserIdPlannerRecentPlanId struct
func NewUserIdPlannerRecentPlanID(userId string, plannerPlanId string) UserIdPlannerRecentPlanId {
	return UserIdPlannerRecentPlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserIdPlannerRecentPlanID parses 'input' into a UserIdPlannerRecentPlanId
func ParseUserIdPlannerRecentPlanID(input string) (*UserIdPlannerRecentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerRecentPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerRecentPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerRecentPlanIDInsensitively parses 'input' case-insensitively into a UserIdPlannerRecentPlanId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerRecentPlanIDInsensitively(input string) (*UserIdPlannerRecentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerRecentPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerRecentPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerRecentPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateUserIdPlannerRecentPlanID checks that 'input' can be parsed as a User Id Planner Recent Plan ID
func ValidateUserIdPlannerRecentPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerRecentPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Recent Plan ID
func (id UserIdPlannerRecentPlanId) ID() string {
	fmtString := "/users/%s/planner/recentPlans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Recent Plan ID
func (id UserIdPlannerRecentPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("recentPlans", "recentPlans", "recentPlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this User Id Planner Recent Plan ID
func (id UserIdPlannerRecentPlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Id Planner Recent Plan (%s)", strings.Join(components, "\n"))
}
