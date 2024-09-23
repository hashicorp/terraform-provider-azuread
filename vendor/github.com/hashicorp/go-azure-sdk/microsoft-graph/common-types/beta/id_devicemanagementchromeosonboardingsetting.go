package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementChromeOSOnboardingSettingId{}

// DeviceManagementChromeOSOnboardingSettingId is a struct representing the Resource ID for a Device Management Chrome O S Onboarding Setting
type DeviceManagementChromeOSOnboardingSettingId struct {
	ChromeOSOnboardingSettingsId string
}

// NewDeviceManagementChromeOSOnboardingSettingID returns a new DeviceManagementChromeOSOnboardingSettingId struct
func NewDeviceManagementChromeOSOnboardingSettingID(chromeOSOnboardingSettingsId string) DeviceManagementChromeOSOnboardingSettingId {
	return DeviceManagementChromeOSOnboardingSettingId{
		ChromeOSOnboardingSettingsId: chromeOSOnboardingSettingsId,
	}
}

// ParseDeviceManagementChromeOSOnboardingSettingID parses 'input' into a DeviceManagementChromeOSOnboardingSettingId
func ParseDeviceManagementChromeOSOnboardingSettingID(input string) (*DeviceManagementChromeOSOnboardingSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementChromeOSOnboardingSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementChromeOSOnboardingSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementChromeOSOnboardingSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementChromeOSOnboardingSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementChromeOSOnboardingSettingIDInsensitively(input string) (*DeviceManagementChromeOSOnboardingSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementChromeOSOnboardingSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementChromeOSOnboardingSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementChromeOSOnboardingSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChromeOSOnboardingSettingsId, ok = input.Parsed["chromeOSOnboardingSettingsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chromeOSOnboardingSettingsId", input)
	}

	return nil
}

// ValidateDeviceManagementChromeOSOnboardingSettingID checks that 'input' can be parsed as a Device Management Chrome O S Onboarding Setting ID
func ValidateDeviceManagementChromeOSOnboardingSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementChromeOSOnboardingSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Chrome O S Onboarding Setting ID
func (id DeviceManagementChromeOSOnboardingSettingId) ID() string {
	fmtString := "/deviceManagement/chromeOSOnboardingSettings/%s"
	return fmt.Sprintf(fmtString, id.ChromeOSOnboardingSettingsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Chrome O S Onboarding Setting ID
func (id DeviceManagementChromeOSOnboardingSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("chromeOSOnboardingSettings", "chromeOSOnboardingSettings", "chromeOSOnboardingSettings"),
		resourceids.UserSpecifiedSegment("chromeOSOnboardingSettingsId", "chromeOSOnboardingSettingsId"),
	}
}

// String returns a human-readable description of this Device Management Chrome O S Onboarding Setting ID
func (id DeviceManagementChromeOSOnboardingSettingId) String() string {
	components := []string{
		fmt.Sprintf("Chrome O S Onboarding Settings: %q", id.ChromeOSOnboardingSettingsId),
	}
	return fmt.Sprintf("Device Management Chrome O S Onboarding Setting (%s)", strings.Join(components, "\n"))
}
