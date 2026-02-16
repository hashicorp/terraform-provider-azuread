package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileInterestId{}

// UserIdProfileInterestId is a struct representing the Resource ID for a User Id Profile Interest
type UserIdProfileInterestId struct {
	UserId           string
	PersonInterestId string
}

// NewUserIdProfileInterestID returns a new UserIdProfileInterestId struct
func NewUserIdProfileInterestID(userId string, personInterestId string) UserIdProfileInterestId {
	return UserIdProfileInterestId{
		UserId:           userId,
		PersonInterestId: personInterestId,
	}
}

// ParseUserIdProfileInterestID parses 'input' into a UserIdProfileInterestId
func ParseUserIdProfileInterestID(input string) (*UserIdProfileInterestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileInterestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileInterestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileInterestIDInsensitively parses 'input' case-insensitively into a UserIdProfileInterestId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileInterestIDInsensitively(input string) (*UserIdProfileInterestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileInterestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileInterestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileInterestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonInterestId, ok = input.Parsed["personInterestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personInterestId", input)
	}

	return nil
}

// ValidateUserIdProfileInterestID checks that 'input' can be parsed as a User Id Profile Interest ID
func ValidateUserIdProfileInterestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileInterestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Interest ID
func (id UserIdProfileInterestId) ID() string {
	fmtString := "/users/%s/profile/interests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonInterestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Interest ID
func (id UserIdProfileInterestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("interests", "interests", "interests"),
		resourceids.UserSpecifiedSegment("personInterestId", "personInterestId"),
	}
}

// String returns a human-readable description of this User Id Profile Interest ID
func (id UserIdProfileInterestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Interest: %q", id.PersonInterestId),
	}
	return fmt.Sprintf("User Id Profile Interest (%s)", strings.Join(components, "\n"))
}
