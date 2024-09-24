package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileNameId{}

// UserIdProfileNameId is a struct representing the Resource ID for a User Id Profile Name
type UserIdProfileNameId struct {
	UserId       string
	PersonNameId string
}

// NewUserIdProfileNameID returns a new UserIdProfileNameId struct
func NewUserIdProfileNameID(userId string, personNameId string) UserIdProfileNameId {
	return UserIdProfileNameId{
		UserId:       userId,
		PersonNameId: personNameId,
	}
}

// ParseUserIdProfileNameID parses 'input' into a UserIdProfileNameId
func ParseUserIdProfileNameID(input string) (*UserIdProfileNameId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileNameId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileNameId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileNameIDInsensitively parses 'input' case-insensitively into a UserIdProfileNameId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileNameIDInsensitively(input string) (*UserIdProfileNameId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileNameId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileNameId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileNameId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonNameId, ok = input.Parsed["personNameId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personNameId", input)
	}

	return nil
}

// ValidateUserIdProfileNameID checks that 'input' can be parsed as a User Id Profile Name ID
func ValidateUserIdProfileNameID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileNameID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Name ID
func (id UserIdProfileNameId) ID() string {
	fmtString := "/users/%s/profile/names/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonNameId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Name ID
func (id UserIdProfileNameId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("names", "names", "names"),
		resourceids.UserSpecifiedSegment("personNameId", "personNameId"),
	}
}

// String returns a human-readable description of this User Id Profile Name ID
func (id UserIdProfileNameId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Name: %q", id.PersonNameId),
	}
	return fmt.Sprintf("User Id Profile Name (%s)", strings.Join(components, "\n"))
}
