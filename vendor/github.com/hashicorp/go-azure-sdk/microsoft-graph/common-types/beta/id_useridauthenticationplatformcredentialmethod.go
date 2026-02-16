package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPlatformCredentialMethodId{}

// UserIdAuthenticationPlatformCredentialMethodId is a struct representing the Resource ID for a User Id Authentication Platform Credential Method
type UserIdAuthenticationPlatformCredentialMethodId struct {
	UserId                                   string
	PlatformCredentialAuthenticationMethodId string
}

// NewUserIdAuthenticationPlatformCredentialMethodID returns a new UserIdAuthenticationPlatformCredentialMethodId struct
func NewUserIdAuthenticationPlatformCredentialMethodID(userId string, platformCredentialAuthenticationMethodId string) UserIdAuthenticationPlatformCredentialMethodId {
	return UserIdAuthenticationPlatformCredentialMethodId{
		UserId:                                   userId,
		PlatformCredentialAuthenticationMethodId: platformCredentialAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationPlatformCredentialMethodID parses 'input' into a UserIdAuthenticationPlatformCredentialMethodId
func ParseUserIdAuthenticationPlatformCredentialMethodID(input string) (*UserIdAuthenticationPlatformCredentialMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPlatformCredentialMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPlatformCredentialMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationPlatformCredentialMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationPlatformCredentialMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationPlatformCredentialMethodIDInsensitively(input string) (*UserIdAuthenticationPlatformCredentialMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationPlatformCredentialMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationPlatformCredentialMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationPlatformCredentialMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.PlatformCredentialAuthenticationMethodId, ok = input.Parsed["platformCredentialAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "platformCredentialAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationPlatformCredentialMethodID checks that 'input' can be parsed as a User Id Authentication Platform Credential Method ID
func ValidateUserIdAuthenticationPlatformCredentialMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationPlatformCredentialMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Platform Credential Method ID
func (id UserIdAuthenticationPlatformCredentialMethodId) ID() string {
	fmtString := "/users/%s/authentication/platformCredentialMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlatformCredentialAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Platform Credential Method ID
func (id UserIdAuthenticationPlatformCredentialMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("platformCredentialMethods", "platformCredentialMethods", "platformCredentialMethods"),
		resourceids.UserSpecifiedSegment("platformCredentialAuthenticationMethodId", "platformCredentialAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Platform Credential Method ID
func (id UserIdAuthenticationPlatformCredentialMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Platform Credential Authentication Method: %q", id.PlatformCredentialAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Platform Credential Method (%s)", strings.Join(components, "\n"))
}
