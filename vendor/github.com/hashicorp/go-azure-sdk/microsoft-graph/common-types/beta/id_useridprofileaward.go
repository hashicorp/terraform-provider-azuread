package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAwardId{}

// UserIdProfileAwardId is a struct representing the Resource ID for a User Id Profile Award
type UserIdProfileAwardId struct {
	UserId        string
	PersonAwardId string
}

// NewUserIdProfileAwardID returns a new UserIdProfileAwardId struct
func NewUserIdProfileAwardID(userId string, personAwardId string) UserIdProfileAwardId {
	return UserIdProfileAwardId{
		UserId:        userId,
		PersonAwardId: personAwardId,
	}
}

// ParseUserIdProfileAwardID parses 'input' into a UserIdProfileAwardId
func ParseUserIdProfileAwardID(input string) (*UserIdProfileAwardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAwardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAwardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileAwardIDInsensitively parses 'input' case-insensitively into a UserIdProfileAwardId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileAwardIDInsensitively(input string) (*UserIdProfileAwardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAwardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAwardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileAwardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonAwardId, ok = input.Parsed["personAwardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personAwardId", input)
	}

	return nil
}

// ValidateUserIdProfileAwardID checks that 'input' can be parsed as a User Id Profile Award ID
func ValidateUserIdProfileAwardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileAwardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Award ID
func (id UserIdProfileAwardId) ID() string {
	fmtString := "/users/%s/profile/awards/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonAwardId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Award ID
func (id UserIdProfileAwardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("awards", "awards", "awards"),
		resourceids.UserSpecifiedSegment("personAwardId", "personAwardId"),
	}
}

// String returns a human-readable description of this User Id Profile Award ID
func (id UserIdProfileAwardId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Award: %q", id.PersonAwardId),
	}
	return fmt.Sprintf("User Id Profile Award (%s)", strings.Join(components, "\n"))
}
