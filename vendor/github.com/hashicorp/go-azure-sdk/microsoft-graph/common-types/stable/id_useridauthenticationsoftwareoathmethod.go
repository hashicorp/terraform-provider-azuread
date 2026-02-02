package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationSoftwareOathMethodId{}

// UserIdAuthenticationSoftwareOathMethodId is a struct representing the Resource ID for a User Id Authentication Software Oath Method
type UserIdAuthenticationSoftwareOathMethodId struct {
	UserId                             string
	SoftwareOathAuthenticationMethodId string
}

// NewUserIdAuthenticationSoftwareOathMethodID returns a new UserIdAuthenticationSoftwareOathMethodId struct
func NewUserIdAuthenticationSoftwareOathMethodID(userId string, softwareOathAuthenticationMethodId string) UserIdAuthenticationSoftwareOathMethodId {
	return UserIdAuthenticationSoftwareOathMethodId{
		UserId:                             userId,
		SoftwareOathAuthenticationMethodId: softwareOathAuthenticationMethodId,
	}
}

// ParseUserIdAuthenticationSoftwareOathMethodID parses 'input' into a UserIdAuthenticationSoftwareOathMethodId
func ParseUserIdAuthenticationSoftwareOathMethodID(input string) (*UserIdAuthenticationSoftwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationSoftwareOathMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationSoftwareOathMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationSoftwareOathMethodIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationSoftwareOathMethodId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationSoftwareOathMethodIDInsensitively(input string) (*UserIdAuthenticationSoftwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationSoftwareOathMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationSoftwareOathMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationSoftwareOathMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.SoftwareOathAuthenticationMethodId, ok = input.Parsed["softwareOathAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "softwareOathAuthenticationMethodId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationSoftwareOathMethodID checks that 'input' can be parsed as a User Id Authentication Software Oath Method ID
func ValidateUserIdAuthenticationSoftwareOathMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationSoftwareOathMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Software Oath Method ID
func (id UserIdAuthenticationSoftwareOathMethodId) ID() string {
	fmtString := "/users/%s/authentication/softwareOathMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SoftwareOathAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Software Oath Method ID
func (id UserIdAuthenticationSoftwareOathMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("softwareOathMethods", "softwareOathMethods", "softwareOathMethods"),
		resourceids.UserSpecifiedSegment("softwareOathAuthenticationMethodId", "softwareOathAuthenticationMethodId"),
	}
}

// String returns a human-readable description of this User Id Authentication Software Oath Method ID
func (id UserIdAuthenticationSoftwareOathMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Software Oath Authentication Method: %q", id.SoftwareOathAuthenticationMethodId),
	}
	return fmt.Sprintf("User Id Authentication Software Oath Method (%s)", strings.Join(components, "\n"))
}
