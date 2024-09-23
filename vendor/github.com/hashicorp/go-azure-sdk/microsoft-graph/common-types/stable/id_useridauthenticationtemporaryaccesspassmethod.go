package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationTemporaryAccessPassMethodId{}

// UserIdAuthenticationTemporaryAccessPassMethodId is a struct representing the Resource ID for a User Id Authentication Temporary Access Pass Method
type UserIdAuthenticationTemporaryAccessPassMethodId struct {
	UserId                                    string
	TemporaryAccessPassAuthenticationMethodId string
}

// NewUserIdAuthenticationTemporaryAccessPassMethodID returns a new UserIdAuthenticationTemporaryAccessPassMethodId struct
func NewUserIdAuthenticationTemporaryAccessPassMethodID(userId string, temporaryAccessPassAuthenticationMethodId string) UserIdAuthenticationTemporaryAccessPassMethodId {
	return UserIdAuthenticationTemporaryAccessPassMethodId{
		UserId: userId,
		TemporaryAccessPassAuthenticationMethodId: temporaryAccessPassAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationTemporaryAccessPassMethodID parses 'input' into a UserIdAuthenticationTemporaryAccessPassMethodId
func ParseUserIdAuthenticationTemporaryAccessPassMethodID(input string) (*UserIdAuthenticationTemporaryAccessPassMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationTemporaryAccessPassMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationTemporaryAccessPassMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationTemporaryAccessPassMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationTemporaryAccessPassMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationTemporaryAccessPassMethodIDInsensitively(input string) (*UserIdAuthenticationTemporaryAccessPassMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationTemporaryAccessPassMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationTemporaryAccessPassMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationTemporaryAccessPassMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TemporaryAccessPassAuthenticationMethodId, ok = input.Parsed["temporaryAccessPassAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "temporaryAccessPassAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationTemporaryAccessPassMethodID checks that 'input' can be parsed as a User Id Authentication Temporary Access Pass Method ID
func ValidateUserIdAuthenticationTemporaryAccessPassMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationTemporaryAccessPassMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Temporary Access Pass Method ID
func (id UserIdAuthenticationTemporaryAccessPassMethodId) ID() string {
	fmtString := "/users/%s/authentication/temporaryAccessPassMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TemporaryAccessPassAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Temporary Access Pass Method ID
func (id UserIdAuthenticationTemporaryAccessPassMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("temporaryAccessPassMethods", "temporaryAccessPassMethods", "temporaryAccessPassMethods"),
		resourceids.UserSpecifiedSegment("temporaryAccessPassAuthenticationMethodId", "temporaryAccessPassAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Temporary Access Pass Method ID
func (id UserIdAuthenticationTemporaryAccessPassMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Temporary Access Pass Authentication Method: %q", id.TemporaryAccessPassAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Temporary Access Pass Method (%s)", strings.Join(components, "\n"))
}
