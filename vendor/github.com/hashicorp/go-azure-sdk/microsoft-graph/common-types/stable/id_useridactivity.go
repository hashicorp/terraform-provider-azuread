package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdActivityId{}

// UserIdActivityId is a struct representing the Resource ID for a User Id Activity
type UserIdActivityId struct {
	UserId         string
	UserActivityId string
}

// NewUserIdActivityID returns a new UserIdActivityId struct
func NewUserIdActivityID(userId string, userActivityId string) UserIdActivityId {
	return UserIdActivityId{
		UserId:         userId,
		UserActivityId: userActivityId,
	}
}

// ParseUserIdActivityID parses 'input' into a UserIdActivityId
func ParseUserIdActivityID(input string) (*UserIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdActivityIDInsensitively parses 'input' case-insensitively into a UserIdActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdActivityIDInsensitively(input string) (*UserIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.UserActivityId, ok = input.Parsed["userActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userActivityId", input)
	}

	return nil
}

// ValidateUserIdActivityID checks that 'input' can be parsed as a User Id Activity ID
func ValidateUserIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Activity ID
func (id UserIdActivityId) ID() string {
	fmtString := "/users/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Activity ID
func (id UserIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("userActivityId", "userActivityId"),
	}
}

// String returns a human-readable description of this User Id Activity ID
func (id UserIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Activity: %q", id.UserActivityId),
	}
	return fmt.Sprintf("User Id Activity (%s)", strings.Join(components, "\n"))
}
