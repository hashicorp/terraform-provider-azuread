package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationMethodId{}

// UserIdAuthenticationMethodId is a struct representing the Resource ID for a User Id Authentication Method
type UserIdAuthenticationMethodId struct {
	UserId                 string
	AuthenticationMethodId string
}

// NewUserIdAuthenticationMethodID returns a new UserIdAuthenticationMethodId struct
func NewUserIdAuthenticationMethodID(userId string, authenticationMethodId string) UserIdAuthenticationMethodId {
	return UserIdAuthenticationMethodId{
		UserId:                 userId,
		AuthenticationMethodId: authenticationMethodId,
	}
}

// ParseUserIdAuthenticationMethodID parses 'input' into a UserIdAuthenticationMethodId
func ParseUserIdAuthenticationMethodID(input string) (*UserIdAuthenticationMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationMethodIDInsensitively(input string) (*UserIdAuthenticationMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AuthenticationMethodId, ok = input.Parsed["authenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationMethodID checks that 'input' can be parsed as a User Id Authentication Method ID
func ValidateUserIdAuthenticationMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Method ID
func (id UserIdAuthenticationMethodId) ID() string {
	fmtString := "/users/%s/authentication/methods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Method ID
func (id UserIdAuthenticationMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("methods", "methods", "methods"),
		resourceids.UserSpecifiedSegment("authenticationMethodId", "authenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Method ID
func (id UserIdAuthenticationMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Authentication Method: %q", id.AuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Method (%s)", strings.Join(components, "\n"))
}
