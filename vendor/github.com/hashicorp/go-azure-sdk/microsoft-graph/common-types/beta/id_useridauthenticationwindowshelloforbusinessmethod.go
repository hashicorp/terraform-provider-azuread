package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationWindowsHelloForBusinessMethodId{}

// UserIdAuthenticationWindowsHelloForBusinessMethodId is a struct representing the Resource ID for a User Id Authentication Windows Hello For Business Method
type UserIdAuthenticationWindowsHelloForBusinessMethodId struct {
	UserId                                        string
	WindowsHelloForBusinessAuthenticationMethodId string
}

// NewUserIdAuthenticationWindowsHelloForBusinessMethodID returns a new UserIdAuthenticationWindowsHelloForBusinessMethodId struct
func NewUserIdAuthenticationWindowsHelloForBusinessMethodID(userId string, windowsHelloForBusinessAuthenticationMethodId string) UserIdAuthenticationWindowsHelloForBusinessMethodId {
	return UserIdAuthenticationWindowsHelloForBusinessMethodId{
		UserId: userId,
		WindowsHelloForBusinessAuthenticationMethodId: windowsHelloForBusinessAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationWindowsHelloForBusinessMethodID parses 'input' into a UserIdAuthenticationWindowsHelloForBusinessMethodId
func ParseUserIdAuthenticationWindowsHelloForBusinessMethodID(input string) (*UserIdAuthenticationWindowsHelloForBusinessMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationWindowsHelloForBusinessMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationWindowsHelloForBusinessMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationWindowsHelloForBusinessMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationWindowsHelloForBusinessMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationWindowsHelloForBusinessMethodIDInsensitively(input string) (*UserIdAuthenticationWindowsHelloForBusinessMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationWindowsHelloForBusinessMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationWindowsHelloForBusinessMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationWindowsHelloForBusinessMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.WindowsHelloForBusinessAuthenticationMethodId, ok = input.Parsed["windowsHelloForBusinessAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsHelloForBusinessAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationWindowsHelloForBusinessMethodID checks that 'input' can be parsed as a User Id Authentication Windows Hello For Business Method ID
func ValidateUserIdAuthenticationWindowsHelloForBusinessMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationWindowsHelloForBusinessMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Windows Hello For Business Method ID
func (id UserIdAuthenticationWindowsHelloForBusinessMethodId) ID() string {
	fmtString := "/users/%s/authentication/windowsHelloForBusinessMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WindowsHelloForBusinessAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Windows Hello For Business Method ID
func (id UserIdAuthenticationWindowsHelloForBusinessMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("windowsHelloForBusinessMethods", "windowsHelloForBusinessMethods", "windowsHelloForBusinessMethods"),
		resourceids.UserSpecifiedSegment("windowsHelloForBusinessAuthenticationMethodId", "windowsHelloForBusinessAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Windows Hello For Business Method ID
func (id UserIdAuthenticationWindowsHelloForBusinessMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Windows Hello For Business Authentication Method: %q", id.WindowsHelloForBusinessAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Windows Hello For Business Method (%s)", strings.Join(components, "\n"))
}
