package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileWebAccountId{}

// UserIdProfileWebAccountId is a struct representing the Resource ID for a User Id Profile Web Account
type UserIdProfileWebAccountId struct {
	UserId       string
	WebAccountId string
}

// NewUserIdProfileWebAccountID returns a new UserIdProfileWebAccountId struct
func NewUserIdProfileWebAccountID(userId string, webAccountId string) UserIdProfileWebAccountId {
	return UserIdProfileWebAccountId{
		UserId:       userId,
		WebAccountId: webAccountId,
	}
}

// ParseUserIdProfileWebAccountID parses 'input' into a UserIdProfileWebAccountId
func ParseUserIdProfileWebAccountID(input string) (*UserIdProfileWebAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileWebAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileWebAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileWebAccountIDInsensitively parses 'input' case-insensitively into a UserIdProfileWebAccountId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileWebAccountIDInsensitively(input string) (*UserIdProfileWebAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileWebAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileWebAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileWebAccountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.WebAccountId, ok = input.Parsed["webAccountId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "webAccountId", input)
	}

	return nil
}

// ValidateUserIdProfileWebAccountID checks that 'input' can be parsed as a User Id Profile Web Account ID
func ValidateUserIdProfileWebAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileWebAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Web Account ID
func (id UserIdProfileWebAccountId) ID() string {
	fmtString := "/users/%s/profile/webAccounts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WebAccountId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Web Account ID
func (id UserIdProfileWebAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("webAccounts", "webAccounts", "webAccounts"),
		resourceids.UserSpecifiedSegment("webAccountId", "webAccountId"),
	}
}

// String returns a human-readable description of this User Id Profile Web Account ID
func (id UserIdProfileWebAccountId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Web Account: %q", id.WebAccountId),
	}
	return fmt.Sprintf("User Id Profile Web Account (%s)", strings.Join(components, "\n"))
}
