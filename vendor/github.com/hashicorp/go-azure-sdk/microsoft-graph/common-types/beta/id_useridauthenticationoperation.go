package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationOperationId{}

// UserIdAuthenticationOperationId is a struct representing the Resource ID for a User Id Authentication Operation
type UserIdAuthenticationOperationId struct {
	UserId                 string
	LongRunningOperationId string
}

// NewUserIdAuthenticationOperationID returns a new UserIdAuthenticationOperationId struct
func NewUserIdAuthenticationOperationID(userId string, longRunningOperationId string) UserIdAuthenticationOperationId {
	return UserIdAuthenticationOperationId{
		UserId:                 userId,
		LongRunningOperationId: longRunningOperationId,
	}
}

// ParseUserIdAuthenticationOperationID parses 'input' into a UserIdAuthenticationOperationId
func ParseUserIdAuthenticationOperationID(input string) (*UserIdAuthenticationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationOperationIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationOperationIDInsensitively(input string) (*UserIdAuthenticationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.LongRunningOperationId, ok = input.Parsed["longRunningOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longRunningOperationId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationOperationID checks that 'input' can be parsed as a User Id Authentication Operation ID
func ValidateUserIdAuthenticationOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Operation ID
func (id UserIdAuthenticationOperationId) ID() string {
	fmtString := "/users/%s/authentication/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Operation ID
func (id UserIdAuthenticationOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("longRunningOperationId", "longRunningOperationId"),
	}
}

// String returns a human-readable description of this User Id Authentication Operation ID
func (id UserIdAuthenticationOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Long Running Operation: %q", id.LongRunningOperationId),
	}
	return fmt.Sprintf("User Id Authentication Operation (%s)", strings.Join(components, "\n"))
}
