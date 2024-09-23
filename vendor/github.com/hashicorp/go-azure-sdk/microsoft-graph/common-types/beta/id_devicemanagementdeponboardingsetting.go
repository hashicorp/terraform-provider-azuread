package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingId{}

// DeviceManagementDepOnboardingSettingId is a struct representing the Resource ID for a Device Management Dep Onboarding Setting
type DeviceManagementDepOnboardingSettingId struct {
	DepOnboardingSettingId string
}

// NewDeviceManagementDepOnboardingSettingID returns a new DeviceManagementDepOnboardingSettingId struct
func NewDeviceManagementDepOnboardingSettingID(depOnboardingSettingId string) DeviceManagementDepOnboardingSettingId {
	return DeviceManagementDepOnboardingSettingId{
		DepOnboardingSettingId: depOnboardingSettingId,
	}
}

// ParseDeviceManagementDepOnboardingSettingID parses 'input' into a DeviceManagementDepOnboardingSettingId
func ParseDeviceManagementDepOnboardingSettingID(input string) (*DeviceManagementDepOnboardingSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDepOnboardingSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDepOnboardingSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDepOnboardingSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementDepOnboardingSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDepOnboardingSettingIDInsensitively(input string) (*DeviceManagementDepOnboardingSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDepOnboardingSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDepOnboardingSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDepOnboardingSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DepOnboardingSettingId, ok = input.Parsed["depOnboardingSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "depOnboardingSettingId", input)
	}

	return nil
}

// ValidateDeviceManagementDepOnboardingSettingID checks that 'input' can be parsed as a Device Management Dep Onboarding Setting ID
func ValidateDeviceManagementDepOnboardingSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDepOnboardingSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Dep Onboarding Setting ID
func (id DeviceManagementDepOnboardingSettingId) ID() string {
	fmtString := "/deviceManagement/depOnboardingSettings/%s"
	return fmt.Sprintf(fmtString, id.DepOnboardingSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Dep Onboarding Setting ID
func (id DeviceManagementDepOnboardingSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("depOnboardingSettings", "depOnboardingSettings", "depOnboardingSettings"),
		resourceids.UserSpecifiedSegment("depOnboardingSettingId", "depOnboardingSettingId"),
	}
}

// String returns a human-readable description of this Device Management Dep Onboarding Setting ID
func (id DeviceManagementDepOnboardingSettingId) String() string {
	components := []string{
		fmt.Sprintf("Dep Onboarding Setting: %q", id.DepOnboardingSettingId),
	}
	return fmt.Sprintf("Device Management Dep Onboarding Setting (%s)", strings.Join(components, "\n"))
}
