package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPasswordMethodId{}

// UserIdAuthenticationPasswordMethodId is a struct representing the Resource ID for a User Id Authentication Password Method
type UserIdAuthenticationPasswordMethodId struct {
	UserId                         string
	PasswordAuthenticationMethodId string
}

// NewUserIdAuthenticationPasswordMethodID returns a new UserIdAuthenticationPasswordMethodId struct
func NewUserIdAuthenticationPasswordMethodID(userId string, passwordAuthenticationMethodId string) UserIdAuthenticationPasswordMethodId {
	return UserIdAuthenticationPasswordMethodId{
		UserId:                         userId,
		PasswordAuthenticationMethodId: passwordAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationPasswordMethodID parses 'input' into a UserIdAuthenticationPasswordMethodId
func ParseUserIdAuthenticationPasswordMethodID(input string) (*UserIdAuthenticationPasswordMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPasswordMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPasswordMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationPasswordMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationPasswordMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationPasswordMethodIDInsensitively(input string) (*UserIdAuthenticationPasswordMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPasswordMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPasswordMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationPasswordMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PasswordAuthenticationMethodId, ok = input.Parsed["passwordAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "passwordAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationPasswordMethodID checks that 'input' can be parsed as a User Id Authentication Password Method ID
func ValidateUserIdAuthenticationPasswordMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationPasswordMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Password Method ID
func (id UserIdAuthenticationPasswordMethodId) ID() string {
	fmtString := "/users/%s/authentication/passwordMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PasswordAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Password Method ID
func (id UserIdAuthenticationPasswordMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("passwordMethods", "passwordMethods", "passwordMethods"),
		resourceids.UserSpecifiedSegment("passwordAuthenticationMethodId", "passwordAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Password Method ID
func (id UserIdAuthenticationPasswordMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Password Authentication Method: %q", id.PasswordAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Password Method (%s)", strings.Join(components, "\n"))
}
