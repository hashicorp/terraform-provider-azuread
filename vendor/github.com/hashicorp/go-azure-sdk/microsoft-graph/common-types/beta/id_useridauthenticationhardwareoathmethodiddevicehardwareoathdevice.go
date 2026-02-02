package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}

// UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId is a struct representing the Resource ID for a User Id Authentication Hardware Oath Method Id Device Hardware Oath Device
type UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId struct {
	UserId                                        string
	HardwareOathAuthenticationMethodId            string
	HardwareOathTokenAuthenticationMethodDeviceId string
}

// NewUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID returns a new UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId struct
func NewUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(userId string, hardwareOathAuthenticationMethodId string, hardwareOathTokenAuthenticationMethodDeviceId string) UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId {
	return UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
		UserId:                             userId,
		HardwareOathAuthenticationMethodId: hardwareOathAuthenticationMethodId,
		HardwareOathTokenAuthenticationMethodDeviceId: hardwareOathTokenAuthenticationMethodDeviceId,
	}
}

// ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID parses 'input' into a UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
func ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(input string) (*UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively parses 'input' case-insensitively into a UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively(input string) (*UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.HardwareOathAuthenticationMethodId, ok = input.Parsed["hardwareOathAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathAuthenticationMethodId", input)
	}

	if id.HardwareOathTokenAuthenticationMethodDeviceId, ok = input.Parsed["hardwareOathTokenAuthenticationMethodDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathTokenAuthenticationMethodDeviceId", input)
	}

	return nil
}

// ValidateUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID checks that 'input' can be parsed as a User Id Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func ValidateUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func (id UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) ID() string {
	fmtString := "/users/%s/authentication/hardwareOathMethods/%s/device/hardwareOathDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.HardwareOathAuthenticationMethodId, id.HardwareOathTokenAuthenticationMethodDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func (id UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("hardwareOathMethods", "hardwareOathMethods", "hardwareOathMethods"),
		resourceids.UserSpecifiedSegment("hardwareOathAuthenticationMethodId", "hardwareOathAuthenticationMethodId"),
		resourceids.StaticSegment("device", "device", "device"),
		resourceids.StaticSegment("hardwareOathDevices", "hardwareOathDevices", "hardwareOathDevices"),
		resourceids.UserSpecifiedSegment("hardwareOathTokenAuthenticationMethodDeviceId", "hardwareOathTokenAuthenticationMethodDeviceId"),
	}
}

// String returns a human-readable description of this User Id Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func (id UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Hardware Oath Authentication Method: %q", id.HardwareOathAuthenticationMethodId),
		fmt.Sprintf("Hardware Oath Token Authentication Method Device: %q", id.HardwareOathTokenAuthenticationMethodDeviceId),
	}
	return fmt.Sprintf("User Id Authentication Hardware Oath Method Id Device Hardware Oath Device (%s)", strings.Join(components, "\n"))
}
