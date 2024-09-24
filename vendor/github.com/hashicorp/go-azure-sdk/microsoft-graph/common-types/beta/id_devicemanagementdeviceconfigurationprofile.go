package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationProfileId{}

// DeviceManagementDeviceConfigurationProfileId is a struct representing the Resource ID for a Device Management Device Configuration Profile
type DeviceManagementDeviceConfigurationProfileId struct {
	DeviceConfigurationProfileId string
}

// NewDeviceManagementDeviceConfigurationProfileID returns a new DeviceManagementDeviceConfigurationProfileId struct
func NewDeviceManagementDeviceConfigurationProfileID(deviceConfigurationProfileId string) DeviceManagementDeviceConfigurationProfileId {
	return DeviceManagementDeviceConfigurationProfileId{
		DeviceConfigurationProfileId: deviceConfigurationProfileId,
	}
}

// ParseDeviceManagementDeviceConfigurationProfileID parses 'input' into a DeviceManagementDeviceConfigurationProfileId
func ParseDeviceManagementDeviceConfigurationProfileID(input string) (*DeviceManagementDeviceConfigurationProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationProfileIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationProfileId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationProfileIDInsensitively(input string) (*DeviceManagementDeviceConfigurationProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationProfileId, ok = input.Parsed["deviceConfigurationProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationProfileId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationProfileID checks that 'input' can be parsed as a Device Management Device Configuration Profile ID
func ValidateDeviceManagementDeviceConfigurationProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Profile ID
func (id DeviceManagementDeviceConfigurationProfileId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurationProfiles/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Profile ID
func (id DeviceManagementDeviceConfigurationProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurationProfiles", "deviceConfigurationProfiles", "deviceConfigurationProfiles"),
		resourceids.UserSpecifiedSegment("deviceConfigurationProfileId", "deviceConfigurationProfileId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Profile ID
func (id DeviceManagementDeviceConfigurationProfileId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration Profile: %q", id.DeviceConfigurationProfileId),
	}
	return fmt.Sprintf("Device Management Device Configuration Profile (%s)", strings.Join(components, "\n"))
}
