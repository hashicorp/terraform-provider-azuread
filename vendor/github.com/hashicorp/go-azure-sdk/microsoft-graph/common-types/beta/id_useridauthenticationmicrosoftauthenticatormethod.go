package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationMicrosoftAuthenticatorMethodId{}

// UserIdAuthenticationMicrosoftAuthenticatorMethodId is a struct representing the Resource ID for a User Id Authentication Microsoft Authenticator Method
type UserIdAuthenticationMicrosoftAuthenticatorMethodId struct {
	UserId                                       string
	MicrosoftAuthenticatorAuthenticationMethodId string
}

// NewUserIdAuthenticationMicrosoftAuthenticatorMethodID returns a new UserIdAuthenticationMicrosoftAuthenticatorMethodId struct
func NewUserIdAuthenticationMicrosoftAuthenticatorMethodID(userId string, microsoftAuthenticatorAuthenticationMethodId string) UserIdAuthenticationMicrosoftAuthenticatorMethodId {
	return UserIdAuthenticationMicrosoftAuthenticatorMethodId{
		UserId: userId,
		MicrosoftAuthenticatorAuthenticationMethodId: microsoftAuthenticatorAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationMicrosoftAuthenticatorMethodID parses 'input' into a UserIdAuthenticationMicrosoftAuthenticatorMethodId
func ParseUserIdAuthenticationMicrosoftAuthenticatorMethodID(input string) (*UserIdAuthenticationMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationMicrosoftAuthenticatorMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationMicrosoftAuthenticatorMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationMicrosoftAuthenticatorMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(input string) (*UserIdAuthenticationMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationMicrosoftAuthenticatorMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationMicrosoftAuthenticatorMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MicrosoftAuthenticatorAuthenticationMethodId, ok = input.Parsed["microsoftAuthenticatorAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftAuthenticatorAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationMicrosoftAuthenticatorMethodID checks that 'input' can be parsed as a User Id Authentication Microsoft Authenticator Method ID
func ValidateUserIdAuthenticationMicrosoftAuthenticatorMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationMicrosoftAuthenticatorMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Microsoft Authenticator Method ID
func (id UserIdAuthenticationMicrosoftAuthenticatorMethodId) ID() string {
	fmtString := "/users/%s/authentication/microsoftAuthenticatorMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MicrosoftAuthenticatorAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Microsoft Authenticator Method ID
func (id UserIdAuthenticationMicrosoftAuthenticatorMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("microsoftAuthenticatorMethods", "microsoftAuthenticatorMethods", "microsoftAuthenticatorMethods"),
		resourceids.UserSpecifiedSegment("microsoftAuthenticatorAuthenticationMethodId", "microsoftAuthenticatorAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Microsoft Authenticator Method ID
func (id UserIdAuthenticationMicrosoftAuthenticatorMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Microsoft Authenticator Authentication Method: %q", id.MicrosoftAuthenticatorAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Microsoft Authenticator Method (%s)", strings.Join(components, "\n"))
}
