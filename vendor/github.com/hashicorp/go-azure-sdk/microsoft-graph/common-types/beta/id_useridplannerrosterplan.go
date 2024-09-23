package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerRosterPlanId{}

// UserIdPlannerRosterPlanId is a struct representing the Resource ID for a User Id Planner Roster Plan
type UserIdPlannerRosterPlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserIdPlannerRosterPlanID returns a new UserIdPlannerRosterPlanId struct
func NewUserIdPlannerRosterPlanID(userId string, plannerPlanId string) UserIdPlannerRosterPlanId {
	return UserIdPlannerRosterPlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserIdPlannerRosterPlanID parses 'input' into a UserIdPlannerRosterPlanId
func ParseUserIdPlannerRosterPlanID(input string) (*UserIdPlannerRosterPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerRosterPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerRosterPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerRosterPlanIDInsensitively parses 'input' case-insensitively into a UserIdPlannerRosterPlanId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerRosterPlanIDInsensitively(input string) (*UserIdPlannerRosterPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerRosterPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerRosterPlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerRosterPlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerPlanId, ok = input.Parsed["plannerPlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", input)
	}

	return nil
}

// ValidateUserIdPlannerRosterPlanID checks that 'input' can be parsed as a User Id Planner Roster Plan ID
func ValidateUserIdPlannerRosterPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerRosterPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner Roster Plan ID
func (id UserIdPlannerRosterPlanId) ID() string {
	fmtString := "/users/%s/planner/rosterPlans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner Roster Plan ID
func (id UserIdPlannerRosterPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("rosterPlans", "rosterPlans", "rosterPlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanId"),
	}
}

// String returns a human-readable description of this User Id Planner Roster Plan ID
func (id UserIdPlannerRosterPlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Id Planner Roster Plan (%s)", strings.Join(components, "\n"))
}
