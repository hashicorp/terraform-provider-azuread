package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePhoneId{}

// UserIdProfilePhoneId is a struct representing the Resource ID for a User Id Profile Phone
type UserIdProfilePhoneId struct {
	UserId      string
	ItemPhoneId string
}

// NewUserIdProfilePhoneID returns a new UserIdProfilePhoneId struct
func NewUserIdProfilePhoneID(userId string, itemPhoneId string) UserIdProfilePhoneId {
	return UserIdProfilePhoneId{
		UserId:      userId,
		ItemPhoneId: itemPhoneId,
	}
}

// ParseUserIdProfilePhoneID parses 'input' into a UserIdProfilePhoneId
func ParseUserIdProfilePhoneID(input string) (*UserIdProfilePhoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePhoneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePhoneId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfilePhoneIDInsensitively parses 'input' case-insensitively into a UserIdProfilePhoneId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfilePhoneIDInsensitively(input string) (*UserIdProfilePhoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePhoneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePhoneId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfilePhoneId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ItemPhoneId, ok = input.Parsed["itemPhoneId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemPhoneId", input)
	}

	return nil
}

// ValidateUserIdProfilePhoneID checks that 'input' can be parsed as a User Id Profile Phone ID
func ValidateUserIdProfilePhoneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfilePhoneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Phone ID
func (id UserIdProfilePhoneId) ID() string {
	fmtString := "/users/%s/profile/phones/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemPhoneId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Phone ID
func (id UserIdProfilePhoneId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("phones", "phones", "phones"),
		resourceids.UserSpecifiedSegment("itemPhoneId", "itemPhoneId"),
	}
}

// String returns a human-readable description of this User Id Profile Phone ID
func (id UserIdProfilePhoneId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Phone: %q", id.ItemPhoneId),
	}
	return fmt.Sprintf("User Id Profile Phone (%s)", strings.Join(components, "\n"))
}
