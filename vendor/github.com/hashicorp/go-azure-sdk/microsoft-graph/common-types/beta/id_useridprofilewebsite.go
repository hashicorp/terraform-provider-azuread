package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileWebsiteId{}

// UserIdProfileWebsiteId is a struct representing the Resource ID for a User Id Profile Website
type UserIdProfileWebsiteId struct {
	UserId          string
	PersonWebsiteId string
}

// NewUserIdProfileWebsiteID returns a new UserIdProfileWebsiteId struct
func NewUserIdProfileWebsiteID(userId string, personWebsiteId string) UserIdProfileWebsiteId {
	return UserIdProfileWebsiteId{
		UserId:          userId,
		PersonWebsiteId: personWebsiteId,
	}
}

// ParseUserIdProfileWebsiteID parses 'input' into a UserIdProfileWebsiteId
func ParseUserIdProfileWebsiteID(input string) (*UserIdProfileWebsiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileWebsiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileWebsiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileWebsiteIDInsensitively parses 'input' case-insensitively into a UserIdProfileWebsiteId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileWebsiteIDInsensitively(input string) (*UserIdProfileWebsiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileWebsiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileWebsiteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileWebsiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PersonWebsiteId, ok = input.Parsed["personWebsiteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personWebsiteId", input)
	}

	return nil
}

// ValidateUserIdProfileWebsiteID checks that 'input' can be parsed as a User Id Profile Website ID
func ValidateUserIdProfileWebsiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileWebsiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Website ID
func (id UserIdProfileWebsiteId) ID() string {
	fmtString := "/users/%s/profile/websites/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonWebsiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Website ID
func (id UserIdProfileWebsiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("websites", "websites", "websites"),
		resourceids.UserSpecifiedSegment("personWebsiteId", "personWebsiteId"),
	}
}

// String returns a human-readable description of this User Id Profile Website ID
func (id UserIdProfileWebsiteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Website: %q", id.PersonWebsiteId),
	}
	return fmt.Sprintf("User Id Profile Website (%s)", strings.Join(components, "\n"))
}
