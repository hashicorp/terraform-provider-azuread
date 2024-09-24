package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileEmailId{}

// UserIdProfileEmailId is a struct representing the Resource ID for a User Id Profile Email
type UserIdProfileEmailId struct {
	UserId      string
	ItemEmailId string
}

// NewUserIdProfileEmailID returns a new UserIdProfileEmailId struct
func NewUserIdProfileEmailID(userId string, itemEmailId string) UserIdProfileEmailId {
	return UserIdProfileEmailId{
		UserId:      userId,
		ItemEmailId: itemEmailId,
	}
}

// ParseUserIdProfileEmailID parses 'input' into a UserIdProfileEmailId
func ParseUserIdProfileEmailID(input string) (*UserIdProfileEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileEmailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileEmailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileEmailIDInsensitively parses 'input' case-insensitively into a UserIdProfileEmailId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileEmailIDInsensitively(input string) (*UserIdProfileEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileEmailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileEmailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileEmailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ItemEmailId, ok = input.Parsed["itemEmailId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemEmailId", input)
	}

	return nil
}

// ValidateUserIdProfileEmailID checks that 'input' can be parsed as a User Id Profile Email ID
func ValidateUserIdProfileEmailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileEmailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Email ID
func (id UserIdProfileEmailId) ID() string {
	fmtString := "/users/%s/profile/emails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemEmailId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Email ID
func (id UserIdProfileEmailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("emails", "emails", "emails"),
		resourceids.UserSpecifiedSegment("itemEmailId", "itemEmailId"),
	}
}

// String returns a human-readable description of this User Id Profile Email ID
func (id UserIdProfileEmailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Email: %q", id.ItemEmailId),
	}
	return fmt.Sprintf("User Id Profile Email (%s)", strings.Join(components, "\n"))
}
