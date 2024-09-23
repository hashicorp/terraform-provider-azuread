package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{}

// DeviceManagementComanagedDeviceIdDeviceConfigurationStateId is a struct representing the Resource ID for a Device Management Comanaged Device Id Device Configuration State
type DeviceManagementComanagedDeviceIdDeviceConfigurationStateId struct {
	ManagedDeviceId            string
	DeviceConfigurationStateId string
}

// NewDeviceManagementComanagedDeviceIdDeviceConfigurationStateID returns a new DeviceManagementComanagedDeviceIdDeviceConfigurationStateId struct
func NewDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(managedDeviceId string, deviceConfigurationStateId string) DeviceManagementComanagedDeviceIdDeviceConfigurationStateId {
	return DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{
		ManagedDeviceId:            managedDeviceId,
		DeviceConfigurationStateId: deviceConfigurationStateId,
	}
}

// ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateID parses 'input' into a DeviceManagementComanagedDeviceIdDeviceConfigurationStateId
func ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(input string) (*DeviceManagementComanagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdDeviceConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdDeviceConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceConfigurationStateId, ok = input.Parsed["deviceConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdDeviceConfigurationStateID checks that 'input' can be parsed as a Device Management Comanaged Device Id Device Configuration State ID
func ValidateDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Device Configuration State ID
func (id DeviceManagementComanagedDeviceIdDeviceConfigurationStateId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/deviceConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Device Configuration State ID
func (id DeviceManagementComanagedDeviceIdDeviceConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceConfigurationStates", "deviceConfigurationStates", "deviceConfigurationStates"),
		resourceids.UserSpecifiedSegment("deviceConfigurationStateId", "deviceConfigurationStateId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Device Configuration State ID
func (id DeviceManagementComanagedDeviceIdDeviceConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Configuration State: %q", id.DeviceConfigurationStateId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Device Configuration State (%s)", strings.Join(components, "\n"))
}
