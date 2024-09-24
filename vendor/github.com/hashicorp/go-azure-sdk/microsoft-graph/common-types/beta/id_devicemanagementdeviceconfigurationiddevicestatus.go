package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdDeviceStatusId{}

// DeviceManagementDeviceConfigurationIdDeviceStatusId is a struct representing the Resource ID for a Device Management Device Configuration Id Device Status
type DeviceManagementDeviceConfigurationIdDeviceStatusId struct {
	DeviceConfigurationId             string
	DeviceConfigurationDeviceStatusId string
}

// NewDeviceManagementDeviceConfigurationIdDeviceStatusID returns a new DeviceManagementDeviceConfigurationIdDeviceStatusId struct
func NewDeviceManagementDeviceConfigurationIdDeviceStatusID(deviceConfigurationId string, deviceConfigurationDeviceStatusId string) DeviceManagementDeviceConfigurationIdDeviceStatusId {
	return DeviceManagementDeviceConfigurationIdDeviceStatusId{
		DeviceConfigurationId:             deviceConfigurationId,
		DeviceConfigurationDeviceStatusId: deviceConfigurationDeviceStatusId,
	}
}

// ParseDeviceManagementDeviceConfigurationIdDeviceStatusID parses 'input' into a DeviceManagementDeviceConfigurationIdDeviceStatusId
func ParseDeviceManagementDeviceConfigurationIdDeviceStatusID(input string) (*DeviceManagementDeviceConfigurationIdDeviceStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdDeviceStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdDeviceStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationIdDeviceStatusIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationIdDeviceStatusId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationIdDeviceStatusIDInsensitively(input string) (*DeviceManagementDeviceConfigurationIdDeviceStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdDeviceStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdDeviceStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationIdDeviceStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationId, ok = input.Parsed["deviceConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationId", input)
	}

	if id.DeviceConfigurationDeviceStatusId, ok = input.Parsed["deviceConfigurationDeviceStatusId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationDeviceStatusId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationIdDeviceStatusID checks that 'input' can be parsed as a Device Management Device Configuration Id Device Status ID
func ValidateDeviceManagementDeviceConfigurationIdDeviceStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationIdDeviceStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Id Device Status ID
func (id DeviceManagementDeviceConfigurationIdDeviceStatusId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurations/%s/deviceStatuses/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationId, id.DeviceConfigurationDeviceStatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Id Device Status ID
func (id DeviceManagementDeviceConfigurationIdDeviceStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurations", "deviceConfigurations", "deviceConfigurations"),
		resourceids.UserSpecifiedSegment("deviceConfigurationId", "deviceConfigurationId"),
		resourceids.StaticSegment("deviceStatuses", "deviceStatuses", "deviceStatuses"),
		resourceids.UserSpecifiedSegment("deviceConfigurationDeviceStatusId", "deviceConfigurationDeviceStatusId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Id Device Status ID
func (id DeviceManagementDeviceConfigurationIdDeviceStatusId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration: %q", id.DeviceConfigurationId),
		fmt.Sprintf("Device Configuration Device Status: %q", id.DeviceConfigurationDeviceStatusId),
	}
	return fmt.Sprintf("Device Management Device Configuration Id Device Status (%s)", strings.Join(components, "\n"))
}
