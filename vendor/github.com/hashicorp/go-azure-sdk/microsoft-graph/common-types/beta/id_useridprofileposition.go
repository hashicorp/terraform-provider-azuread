package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePositionId{}

// UserIdProfilePositionId is a struct representing the Resource ID for a User Id Profile Position
type UserIdProfilePositionId struct {
	UserId         string
	WorkPositionId string
}

// NewUserIdProfilePositionID returns a new UserIdProfilePositionId struct
func NewUserIdProfilePositionID(userId string, workPositionId string) UserIdProfilePositionId {
	return UserIdProfilePositionId{
		UserId:         userId,
		WorkPositionId: workPositionId,
	}
}

// ParseUserIdProfilePositionID parses 'input' into a UserIdProfilePositionId
func ParseUserIdProfilePositionID(input string) (*UserIdProfilePositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePositionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfilePositionIDInsensitively parses 'input' case-insensitively into a UserIdProfilePositionId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfilePositionIDInsensitively(input string) (*UserIdProfilePositionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePositionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePositionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfilePositionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.WorkPositionId, ok = input.Parsed["workPositionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workPositionId", input)
	}

	return nil
}

// ValidateUserIdProfilePositionID checks that 'input' can be parsed as a User Id Profile Position ID
func ValidateUserIdProfilePositionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfilePositionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Position ID
func (id UserIdProfilePositionId) ID() string {
	fmtString := "/users/%s/profile/positions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WorkPositionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Position ID
func (id UserIdProfilePositionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("positions", "positions", "positions"),
		resourceids.UserSpecifiedSegment("workPositionId", "workPositionId"),
	}
}

// String returns a human-readable description of this User Id Profile Position ID
func (id UserIdProfilePositionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Work Position: %q", id.WorkPositionId),
	}
	return fmt.Sprintf("User Id Profile Position (%s)", strings.Join(components, "\n"))
}
