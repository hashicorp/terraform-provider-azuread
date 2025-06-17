package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{}

// DirectoryAuthenticationMethodDeviceHardwareOathDeviceId is a struct representing the Resource ID for a Directory Authentication Method Device Hardware Oath Device
type DirectoryAuthenticationMethodDeviceHardwareOathDeviceId struct {
	HardwareOathTokenAuthenticationMethodDeviceId string
}

// NewDirectoryAuthenticationMethodDeviceHardwareOathDeviceID returns a new DirectoryAuthenticationMethodDeviceHardwareOathDeviceId struct
func NewDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(hardwareOathTokenAuthenticationMethodDeviceId string) DirectoryAuthenticationMethodDeviceHardwareOathDeviceId {
	return DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{
		HardwareOathTokenAuthenticationMethodDeviceId: hardwareOathTokenAuthenticationMethodDeviceId,
	}
}

// ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceID parses 'input' into a DirectoryAuthenticationMethodDeviceHardwareOathDeviceId
func ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(input string) (*DirectoryAuthenticationMethodDeviceHardwareOathDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceIDInsensitively parses 'input' case-insensitively into a DirectoryAuthenticationMethodDeviceHardwareOathDeviceId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceIDInsensitively(input string) (*DirectoryAuthenticationMethodDeviceHardwareOathDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryAuthenticationMethodDeviceHardwareOathDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareOathTokenAuthenticationMethodDeviceId, ok = input.Parsed["hardwareOathTokenAuthenticationMethodDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathTokenAuthenticationMethodDeviceId", input)
	}

	return nil
}

// ValidateDirectoryAuthenticationMethodDeviceHardwareOathDeviceID checks that 'input' can be parsed as a Directory Authentication Method Device Hardware Oath Device ID
func ValidateDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Authentication Method Device Hardware Oath Device ID
func (id DirectoryAuthenticationMethodDeviceHardwareOathDeviceId) ID() string {
	fmtString := "/directory/authenticationMethodDevices/hardwareOathDevices/%s"
	return fmt.Sprintf(fmtString, id.HardwareOathTokenAuthenticationMethodDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Authentication Method Device Hardware Oath Device ID
func (id DirectoryAuthenticationMethodDeviceHardwareOathDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("authenticationMethodDevices", "authenticationMethodDevices", "authenticationMethodDevices"),
		resourceids.StaticSegment("hardwareOathDevices", "hardwareOathDevices", "hardwareOathDevices"),
		resourceids.UserSpecifiedSegment("hardwareOathTokenAuthenticationMethodDeviceId", "hardwareOathTokenAuthenticationMethodDeviceId"),
	}
}

// String returns a human-readable description of this Directory Authentication Method Device Hardware Oath Device ID
func (id DirectoryAuthenticationMethodDeviceHardwareOathDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Oath Token Authentication Method Device: %q", id.HardwareOathTokenAuthenticationMethodDeviceId),
	}
	return fmt.Sprintf("Directory Authentication Method Device Hardware Oath Device (%s)", strings.Join(components, "\n"))
}
