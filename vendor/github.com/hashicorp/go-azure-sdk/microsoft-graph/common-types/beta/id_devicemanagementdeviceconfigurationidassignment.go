package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdAssignmentId{}

// DeviceManagementDeviceConfigurationIdAssignmentId is a struct representing the Resource ID for a Device Management Device Configuration Id Assignment
type DeviceManagementDeviceConfigurationIdAssignmentId struct {
	DeviceConfigurationId           string
	DeviceConfigurationAssignmentId string
}

// NewDeviceManagementDeviceConfigurationIdAssignmentID returns a new DeviceManagementDeviceConfigurationIdAssignmentId struct
func NewDeviceManagementDeviceConfigurationIdAssignmentID(deviceConfigurationId string, deviceConfigurationAssignmentId string) DeviceManagementDeviceConfigurationIdAssignmentId {
	return DeviceManagementDeviceConfigurationIdAssignmentId{
		DeviceConfigurationId:           deviceConfigurationId,
		DeviceConfigurationAssignmentId: deviceConfigurationAssignmentId,
	}
}

// ParseDeviceManagementDeviceConfigurationIdAssignmentID parses 'input' into a DeviceManagementDeviceConfigurationIdAssignmentId
func ParseDeviceManagementDeviceConfigurationIdAssignmentID(input string) (*DeviceManagementDeviceConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceConfigurationIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationId, ok = input.Parsed["deviceConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationId", input)
	}

	if id.DeviceConfigurationAssignmentId, ok = input.Parsed["deviceConfigurationAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationIdAssignmentID checks that 'input' can be parsed as a Device Management Device Configuration Id Assignment ID
func ValidateDeviceManagementDeviceConfigurationIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Id Assignment ID
func (id DeviceManagementDeviceConfigurationIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurations/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationId, id.DeviceConfigurationAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Id Assignment ID
func (id DeviceManagementDeviceConfigurationIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurations", "deviceConfigurations", "deviceConfigurations"),
		resourceids.UserSpecifiedSegment("deviceConfigurationId", "deviceConfigurationId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceConfigurationAssignmentId", "deviceConfigurationAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Id Assignment ID
func (id DeviceManagementDeviceConfigurationIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration: %q", id.DeviceConfigurationId),
		fmt.Sprintf("Device Configuration Assignment: %q", id.DeviceConfigurationAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Configuration Id Assignment (%s)", strings.Join(components, "\n"))
}
