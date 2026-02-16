package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}

// MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId is a struct representing the Resource ID for a Me Authentication Hardware Oath Method Id Device Hardware Oath Device
type MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId struct {
	HardwareOathAuthenticationMethodId            string
	HardwareOathTokenAuthenticationMethodDeviceId string
}

// NewMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID returns a new MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId struct
func NewMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(hardwareOathAuthenticationMethodId string, hardwareOathTokenAuthenticationMethodDeviceId string) MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId {
	return MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
		HardwareOathAuthenticationMethodId:            hardwareOathAuthenticationMethodId,
		HardwareOathTokenAuthenticationMethodDeviceId: hardwareOathTokenAuthenticationMethodDeviceId,
	}
}

// ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID parses 'input' into a MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
func ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(input string) (*MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively parses 'input' case-insensitively into a MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively(input string) (*MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.HardwareOathAuthenticationMethodId, ok = input.Parsed["hardwareOathAuthenticationMethodId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathAuthenticationMethodId", input)
	}

	if id.HardwareOathTokenAuthenticationMethodDeviceId, ok = input.Parsed["hardwareOathTokenAuthenticationMethodDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hardwareOathTokenAuthenticationMethodDeviceId", input)
	}

	return nil
}

// ValidateMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID checks that 'input' can be parsed as a Me Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func ValidateMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func (id MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) ID() string {
	fmtString := "/me/authentication/hardwareOathMethods/%s/device/hardwareOathDevices/%s"
	return fmt.Sprintf(fmtString, id.HardwareOathAuthenticationMethodId, id.HardwareOathTokenAuthenticationMethodDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func (id MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("authentication", "authentication", "authentication"),
		resourceids.StaticSegment("hardwareOathMethods", "hardwareOathMethods", "hardwareOathMethods"),
		resourceids.UserSpecifiedSegment("hardwareOathAuthenticationMethodId", "hardwareOathAuthenticationMethodId"),
		resourceids.StaticSegment("device", "device", "device"),
		resourceids.StaticSegment("hardwareOathDevices", "hardwareOathDevices", "hardwareOathDevices"),
		resourceids.UserSpecifiedSegment("hardwareOathTokenAuthenticationMethodDeviceId", "hardwareOathTokenAuthenticationMethodDeviceId"),
	}
}

// String returns a human-readable description of this Me Authentication Hardware Oath Method Id Device Hardware Oath Device ID
func (id MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Hardware Oath Authentication Method: %q", id.HardwareOathAuthenticationMethodId),
		fmt.Sprintf("Hardware Oath Token Authentication Method Device: %q", id.HardwareOathTokenAuthenticationMethodDeviceId),
	}
	return fmt.Sprintf("Me Authentication Hardware Oath Method Id Device Hardware Oath Device (%s)", strings.Join(components, "\n"))
}
