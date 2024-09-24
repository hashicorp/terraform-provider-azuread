package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAccountId{}

// UserIdProfileAccountId is a struct representing the Resource ID for a User Id Profile Account
type UserIdProfileAccountId struct {
	UserId                   string
	UserAccountInformationId string
}

// NewUserIdProfileAccountID returns a new UserIdProfileAccountId struct
func NewUserIdProfileAccountID(userId string, userAccountInformationId string) UserIdProfileAccountId {
	return UserIdProfileAccountId{
		UserId:                   userId,
		UserAccountInformationId: userAccountInformationId,
	}
}

// ParseUserIdProfileAccountID parses 'input' into a UserIdProfileAccountId
func ParseUserIdProfileAccountID(input string) (*UserIdProfileAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileAccountIDInsensitively parses 'input' case-insensitively into a UserIdProfileAccountId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileAccountIDInsensitively(input string) (*UserIdProfileAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileAccountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.UserAccountInformationId, ok = input.Parsed["userAccountInformationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userAccountInformationId", input)
	}

	return nil
}

// ValidateUserIdProfileAccountID checks that 'input' can be parsed as a User Id Profile Account ID
func ValidateUserIdProfileAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Account ID
func (id UserIdProfileAccountId) ID() string {
	fmtString := "/users/%s/profile/account/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserAccountInformationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Account ID
func (id UserIdProfileAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("account", "account", "account"),
		resourceids.UserSpecifiedSegment("userAccountInformationId", "userAccountInformationId"),
	}
}

// String returns a human-readable description of this User Id Profile Account ID
func (id UserIdProfileAccountId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Account Information: %q", id.UserAccountInformationId),
	}
	return fmt.Sprintf("User Id Profile Account (%s)", strings.Join(components, "\n"))
}
