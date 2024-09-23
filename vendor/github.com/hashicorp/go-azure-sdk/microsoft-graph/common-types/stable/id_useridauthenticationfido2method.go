package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationFido2MethodId{}

// UserIdAuthenticationFido2MethodId is a struct representing the Resource ID for a User Id Authentication Fido 2 Method
type UserIdAuthenticationFido2MethodId struct {
	UserId                      string
	Fido2AuthenticationMethodId string
}

// NewUserIdAuthenticationFido2MethodID returns a new UserIdAuthenticationFido2MethodId struct
func NewUserIdAuthenticationFido2MethodID(userId string, fido2AuthenticationMethodId string) UserIdAuthenticationFido2MethodId {
	return UserIdAuthenticationFido2MethodId{
		UserId:                      userId,
		Fido2AuthenticationMethodId: fido2AuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationFido2MethodID parses 'input' into a UserIdAuthenticationFido2MethodId
func ParseUserIdAuthenticationFido2MethodID(input string) (*UserIdAuthenticationFido2MethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationFido2MethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationFido2MethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationFido2MethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationFido2MethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationFido2MethodIDInsensitively(input string) (*UserIdAuthenticationFido2MethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationFido2MethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationFido2MethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationFido2MethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.Fido2AuthenticationMethodId, ok = input.Parsed["fido2AuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fido2AuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationFido2MethodID checks that 'input' can be parsed as a User Id Authentication Fido 2 Method ID
func ValidateUserIdAuthenticationFido2MethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationFido2MethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Fido 2 Method ID
func (id UserIdAuthenticationFido2MethodId) ID() string {
	fmtString := "/users/%s/authentication/fido2Methods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.Fido2AuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Fido 2 Method ID
func (id UserIdAuthenticationFido2MethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("fido2Methods", "fido2Methods", "fido2Methods"),
		resourceids.UserSpecifiedSegment("fido2AuthenticationMethodId", "fido2AuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Fido 2 Method ID
func (id UserIdAuthenticationFido2MethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Fido 2 Authentication Method: %q", id.Fido2AuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Fido 2 Method (%s)", strings.Join(components, "\n"))
}
