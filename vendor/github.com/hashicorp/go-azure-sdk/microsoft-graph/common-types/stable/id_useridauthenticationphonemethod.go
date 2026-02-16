package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPhoneMethodId{}

// UserIdAuthenticationPhoneMethodId is a struct representing the Resource ID for a User Id Authentication Phone Method
type UserIdAuthenticationPhoneMethodId struct {
	UserId                      string
	PhoneAuthenticationMethodId string
}

// NewUserIdAuthenticationPhoneMethodID returns a new UserIdAuthenticationPhoneMethodId struct
func NewUserIdAuthenticationPhoneMethodID(userId string, phoneAuthenticationMethodId string) UserIdAuthenticationPhoneMethodId {
	return UserIdAuthenticationPhoneMethodId{
		UserId:                      userId,
		PhoneAuthenticationMethodId: phoneAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationPhoneMethodID parses 'input' into a UserIdAuthenticationPhoneMethodId
func ParseUserIdAuthenticationPhoneMethodID(input string) (*UserIdAuthenticationPhoneMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPhoneMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPhoneMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationPhoneMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationPhoneMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationPhoneMethodIDInsensitively(input string) (*UserIdAuthenticationPhoneMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPhoneMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPhoneMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationPhoneMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PhoneAuthenticationMethodId, ok = input.Parsed["phoneAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "phoneAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationPhoneMethodID checks that 'input' can be parsed as a User Id Authentication Phone Method ID
func ValidateUserIdAuthenticationPhoneMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationPhoneMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Phone Method ID
func (id UserIdAuthenticationPhoneMethodId) ID() string {
	fmtString := "/users/%s/authentication/phoneMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PhoneAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Phone Method ID
func (id UserIdAuthenticationPhoneMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("phoneMethods", "phoneMethods", "phoneMethods"),
		resourceids.UserSpecifiedSegment("phoneAuthenticationMethodId", "phoneAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Phone Method ID
func (id UserIdAuthenticationPhoneMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Phone Authentication Method: %q", id.PhoneAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Phone Method (%s)", strings.Join(components, "\n"))
}
