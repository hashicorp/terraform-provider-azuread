package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPersonId{}

// UserIdPersonId is a struct representing the Resource ID for a User Id Person
type UserIdPersonId struct {
	UserId   string
	PersonId string
}

// NewUserIdPersonID returns a new UserIdPersonId struct
func NewUserIdPersonID(userId string, personId string) UserIdPersonId {
	return UserIdPersonId{
		UserId:   userId,
		PersonId: personId,
	}
}

// ParseUserIdPersonID parses 'input' into a UserIdPersonId
func ParseUserIdPersonID(input string) (*UserIdPersonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPersonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPersonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPersonIDInsensitively parses 'input' case-insensitively into a UserIdPersonId
// note: this method should only be used for API response data and not user input
func ParseUserIdPersonIDInsensitively(input string) (*UserIdPersonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPersonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPersonId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPersonId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonId, ok = input.Parsed["personId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personId", input)
	}

	return nil
}

// ValidateUserIdPersonID checks that 'input' can be parsed as a User Id Person ID
func ValidateUserIdPersonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPersonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Person ID
func (id UserIdPersonId) ID() string {
	fmtString := "/users/%s/people/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Person ID
func (id UserIdPersonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("people", "people", "people"),
		resourceids.UserSpecifiedSegment("personId", "personId"),
	}
}

// String returns a human-readable description of this User Id Person ID
func (id UserIdPersonId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person: %q", id.PersonId),
	}
	return fmt.Sprintf("User Id Person (%s)", strings.Join(components, "\n"))
}
