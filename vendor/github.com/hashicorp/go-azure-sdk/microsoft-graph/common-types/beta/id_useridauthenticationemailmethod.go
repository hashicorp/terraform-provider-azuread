package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationEmailMethodId{}

// UserIdAuthenticationEmailMethodId is a struct representing the Resource ID for a User Id Authentication Email Method
type UserIdAuthenticationEmailMethodId struct {
	UserId                      string
	EmailAuthenticationMethodId string
}

// NewUserIdAuthenticationEmailMethodID returns a new UserIdAuthenticationEmailMethodId struct
func NewUserIdAuthenticationEmailMethodID(userId string, emailAuthenticationMethodId string) UserIdAuthenticationEmailMethodId {
	return UserIdAuthenticationEmailMethodId{
		UserId:                      userId,
		EmailAuthenticationMethodId: emailAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationEmailMethodID parses 'input' into a UserIdAuthenticationEmailMethodId
func ParseUserIdAuthenticationEmailMethodID(input string) (*UserIdAuthenticationEmailMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationEmailMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationEmailMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationEmailMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationEmailMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationEmailMethodIDInsensitively(input string) (*UserIdAuthenticationEmailMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationEmailMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationEmailMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationEmailMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EmailAuthenticationMethodId, ok = input.Parsed["emailAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "emailAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationEmailMethodID checks that 'input' can be parsed as a User Id Authentication Email Method ID
func ValidateUserIdAuthenticationEmailMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationEmailMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Email Method ID
func (id UserIdAuthenticationEmailMethodId) ID() string {
	fmtString := "/users/%s/authentication/emailMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EmailAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Email Method ID
func (id UserIdAuthenticationEmailMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("emailMethods", "emailMethods", "emailMethods"),
		resourceids.UserSpecifiedSegment("emailAuthenticationMethodId", "emailAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Email Method ID
func (id UserIdAuthenticationEmailMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Email Authentication Method: %q", id.EmailAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Email Method (%s)", strings.Join(components, "\n"))
}
