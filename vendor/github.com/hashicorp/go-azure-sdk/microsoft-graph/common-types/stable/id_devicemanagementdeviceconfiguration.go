package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationId{}

// DeviceManagementDeviceConfigurationId is a struct representing the Resource ID for a Device Management Device Configuration
type DeviceManagementDeviceConfigurationId struct {
	DeviceConfigurationId string
}

// NewDeviceManagementDeviceConfigurationID returns a new DeviceManagementDeviceConfigurationId struct
func NewDeviceManagementDeviceConfigurationID(deviceConfigurationId string) DeviceManagementDeviceConfigurationId {
	return DeviceManagementDeviceConfigurationId{
		DeviceConfigurationId: deviceConfigurationId,
	}
}

// ParseDeviceManagementDeviceConfigurationID parses 'input' into a DeviceManagementDeviceConfigurationId
func ParseDeviceManagementDeviceConfigurationID(input string) (*DeviceManagementDeviceConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationIDInsensitively(input string) (*DeviceManagementDeviceConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationId, ok = input.Parsed["deviceConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationID checks that 'input' can be parsed as a Device Management Device Configuration ID
func ValidateDeviceManagementDeviceConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration ID
func (id DeviceManagementDeviceConfigurationId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurations/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration ID
func (id DeviceManagementDeviceConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurations", "deviceConfigurations", "deviceConfigurations"),
		resourceids.UserSpecifiedSegment("deviceConfigurationId", "deviceConfigurationId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration ID
func (id DeviceManagementDeviceConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration: %q", id.DeviceConfigurationId),
	}
	return fmt.Sprintf("Device Management Device Configuration (%s)", strings.Join(components, "\n"))
}
