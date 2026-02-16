package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileEducationalActivityId{}

// UserIdProfileEducationalActivityId is a struct representing the Resource ID for a User Id Profile Educational Activity
type UserIdProfileEducationalActivityId struct {
	UserId                string
	EducationalActivityId string
}

// NewUserIdProfileEducationalActivityID returns a new UserIdProfileEducationalActivityId struct
func NewUserIdProfileEducationalActivityID(userId string, educationalActivityId string) UserIdProfileEducationalActivityId {
	return UserIdProfileEducationalActivityId{
		UserId:                userId,
		EducationalActivityId: educationalActivityId,
	}
}

// ParseUserIdProfileEducationalActivityID parses 'input' into a UserIdProfileEducationalActivityId
func ParseUserIdProfileEducationalActivityID(input string) (*UserIdProfileEducationalActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileEducationalActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileEducationalActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileEducationalActivityIDInsensitively parses 'input' case-insensitively into a UserIdProfileEducationalActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileEducationalActivityIDInsensitively(input string) (*UserIdProfileEducationalActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileEducationalActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileEducationalActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileEducationalActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EducationalActivityId, ok = input.Parsed["educationalActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "educationalActivityId", input)
	}

	return nil
}

// ValidateUserIdProfileEducationalActivityID checks that 'input' can be parsed as a User Id Profile Educational Activity ID
func ValidateUserIdProfileEducationalActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileEducationalActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Educational Activity ID
func (id UserIdProfileEducationalActivityId) ID() string {
	fmtString := "/users/%s/profile/educationalActivities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EducationalActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Educational Activity ID
func (id UserIdProfileEducationalActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("educationalActivities", "educationalActivities", "educationalActivities"),
		resourceids.UserSpecifiedSegment("educationalActivityId", "educationalActivityId"),
	}
}

// String returns a human-readable description of this User Id Profile Educational Activity ID
func (id UserIdProfileEducationalActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Educational Activity: %q", id.EducationalActivityId),
	}
	return fmt.Sprintf("User Id Profile Educational Activity (%s)", strings.Join(components, "\n"))
}
