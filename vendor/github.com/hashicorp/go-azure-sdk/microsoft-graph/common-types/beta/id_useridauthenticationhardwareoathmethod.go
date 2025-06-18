package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationHardwareOathMethodId{}

// UserIdAuthenticationHardwareOathMethodId is a struct representing the Resource ID for a User Id Authentication Hardware Oath Method
type UserIdAuthenticationHardwareOathMethodId struct {
	UserId                             string
	HardwareOathAuthenticationMethodId string
}

// NewUserIdAuthenticationHardwareOathMethodID returns a new UserIdAuthenticationHardwareOathMethodId struct
func NewUserIdAuthenticationHardwareOathMethodID(userId string, hardwareOathAuthenticationMethodId string) UserIdAuthenticationHardwareOathMethodId {
	return UserIdAuthenticationHardwareOathMethodId{
		UserId:                             userId,
		HardwareOathAuthenticationMethodId: hardwareOathAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationHardwareOathMethodID parses 'input' into a UserIdAuthenticationHardwareOathMethodId
func ParseUserIdAuthenticationHardwareOathMethodID(input string) (*UserIdAuthenticationHardwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationHardwareOathMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationHardwareOathMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationHardwareOathMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationHardwareOathMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationHardwareOathMethodIDInsensitively(input string) (*UserIdAuthenticationHardwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationHardwareOathMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationHardwareOathMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationHardwareOathMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.HardwareOathAuthenticationMethodId, ok = input.Parsed["hardwareOathAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationHardwareOathMethodID checks that 'input' can be parsed as a User Id Authentication Hardware Oath Method ID
func ValidateUserIdAuthenticationHardwareOathMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationHardwareOathMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Hardware Oath Method ID
func (id UserIdAuthenticationHardwareOathMethodId) ID() string {
	fmtString := "/users/%s/authentication/hardwareOathMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.HardwareOathAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Hardware Oath Method ID
func (id UserIdAuthenticationHardwareOathMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("hardwareOathMethods", "hardwareOathMethods", "hardwareOathMethods"),
		resourceids.UserSpecifiedSegment("hardwareOathAuthenticationMethodId", "hardwareOathAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Hardware Oath Method ID
func (id UserIdAuthenticationHardwareOathMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Hardware Oath Authentication Method: %q", id.HardwareOathAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Hardware Oath Method (%s)", strings.Join(components, "\n"))
}
