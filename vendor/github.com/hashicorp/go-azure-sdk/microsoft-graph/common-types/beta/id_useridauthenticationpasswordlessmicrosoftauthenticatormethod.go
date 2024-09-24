package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

// UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId is a struct representing the Resource ID for a User Id Authentication Passwordless Microsoft Authenticator Method
type UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId struct {
	UserId                                                   string
	PasswordlessMicrosoftAuthenticatorAuthenticationMethodId string
}

// NewUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID returns a new UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId struct
func NewUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(userId string, passwordlessMicrosoftAuthenticatorAuthenticationMethodId string) UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId {
	return UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
		UserId: userId,
		PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: passwordlessMicrosoftAuthenticatorAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID parses 'input' into a UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
func ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(input string) (*UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(input string) (*UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, ok = input.Parsed["passwordlessMicrosoftAuthenticatorAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "passwordlessMicrosoftAuthenticatorAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID checks that 'input' can be parsed as a User Id Authentication Passwordless Microsoft Authenticator Method ID
func ValidateUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Passwordless Microsoft Authenticator Method ID
func (id UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) ID() string {
	fmtString := "/users/%s/authentication/passwordlessMicrosoftAuthenticatorMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Passwordless Microsoft Authenticator Method ID
func (id UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("passwordlessMicrosoftAuthenticatorMethods", "passwordlessMicrosoftAuthenticatorMethods", "passwordlessMicrosoftAuthenticatorMethods"),
		resourceids.UserSpecifiedSegment("passwordlessMicrosoftAuthenticatorAuthenticationMethodId", "passwordlessMicrosoftAuthenticatorAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Passwordless Microsoft Authenticator Method ID
func (id UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Passwordless Microsoft Authenticator Authentication Method: %q", id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Passwordless Microsoft Authenticator Method (%s)", strings.Join(components, "\n"))
}
