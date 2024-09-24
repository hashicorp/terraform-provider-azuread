package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerFavoritePlanId{}

// UserIdPlannerFavoritePlanId is a struct representing the Resource ID for a User Id Planner Favorite Plan
type UserIdPlannerFavoritePlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserIdPlannerFavoritePlanID returns a new UserIdPlannerFavoritePlanId struct
func NewUserIdPlannerFavoritePlanID(userId string, plannerPlanId string) UserIdPlannerFavoritePlanId {
	return UserIdPlannerFavoritePlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserIdPlannerFavoritePlanID parses 'input' into a UserIdPlannerFavoritePlanId
func ParseUserIdPlannerFavoritePlanID(input string) (*UserIdPlannerFavoritePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerFavoritePlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerFavoritePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerFavoritePlanIDInsensitively parses 'input' case-insensitively into a UserIdPlannerFavoritePlanId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerFavoritePlanIDInsensitively(input string) (*UserIdPlannerFavoritePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerFavoritePlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerFavoritePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerFavoritePlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateUserIdPlannerFavoritePlanID checks that 'input' can be parsed as a User Id Planner Favorite Plan ID
func ValidateUserIdPlannerFavoritePlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerFavoritePlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Favorite Plan ID
func (id UserIdPlannerFavoritePlanId) ID() string {
	fmtString := "/users/%s/planner/favoritePlans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Favorite Plan ID
func (id UserIdPlannerFavoritePlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("favoritePlans", "favoritePlans", "favoritePlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this User Id Planner Favorite Plan ID
func (id UserIdPlannerFavoritePlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Id Planner Favorite Plan (%s)", strings.Join(components, "\n"))
}
