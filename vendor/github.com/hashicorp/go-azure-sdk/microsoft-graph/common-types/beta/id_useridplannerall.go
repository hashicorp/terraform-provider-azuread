package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerAllId{}

// UserIdPlannerAllId is a struct representing the Resource ID for a User Id Planner All
type UserIdPlannerAllId struct {
	UserId         string
	PlannerDeltaId string
}

// NewUserIdPlannerAllID returns a new UserIdPlannerAllId struct
func NewUserIdPlannerAllID(userId string, plannerDeltaId string) UserIdPlannerAllId {
	return UserIdPlannerAllId{
		UserId:         userId,
		PlannerDeltaId: plannerDeltaId,
	}
}

// ParseUserIdPlannerAllID parses 'input' into a UserIdPlannerAllId
func ParseUserIdPlannerAllID(input string) (*UserIdPlannerAllId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerAllId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerAllId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPlannerAllIDInsensitively parses 'input' case-insensitively into a UserIdPlannerAllId
// note: this method should only be used for API response data and not user input
func ParseUserIdPlannerAllIDInsensitively(input string) (*UserIdPlannerAllId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPlannerAllId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPlannerAllId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPlannerAllId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlannerDeltaId, ok = input.Parsed["plannerDeltaId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "plannerDeltaId", input)
	}

	return nil
}

// ValidateUserIdPlannerAllID checks that 'input' can be parsed as a User Id Planner All ID
func ValidateUserIdPlannerAllID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPlannerAllID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Planner All ID
func (id UserIdPlannerAllId) ID() string {
	fmtString := "/users/%s/planner/all/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerDeltaId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Planner All ID
func (id UserIdPlannerAllId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("planner", "planner", "planner"),
		resourceids.StaticSegment("all", "all", "all"),
		resourceids.UserSpecifiedSegment("plannerDeltaId", "plannerDeltaId"),
	}
}

// String returns a human-readable description of this User Id Planner All ID
func (id UserIdPlannerAllId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Delta: %q", id.PlannerDeltaId),
	}
	return fmt.Sprintf("User Id Planner All (%s)", strings.Join(components, "\n"))
}
