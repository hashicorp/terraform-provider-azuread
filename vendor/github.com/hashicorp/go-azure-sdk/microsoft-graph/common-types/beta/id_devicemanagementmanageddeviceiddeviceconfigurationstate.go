package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdDeviceConfigurationStateId{}

// DeviceManagementManagedDeviceIdDeviceConfigurationStateId is a struct representing the Resource ID for a Device Management Managed Device Id Device Configuration State
type DeviceManagementManagedDeviceIdDeviceConfigurationStateId struct {
	ManagedDeviceId            string
	DeviceConfigurationStateId string
}

// NewDeviceManagementManagedDeviceIdDeviceConfigurationStateID returns a new DeviceManagementManagedDeviceIdDeviceConfigurationStateId struct
func NewDeviceManagementManagedDeviceIdDeviceConfigurationStateID(managedDeviceId string, deviceConfigurationStateId string) DeviceManagementManagedDeviceIdDeviceConfigurationStateId {
	return DeviceManagementManagedDeviceIdDeviceConfigurationStateId{
		ManagedDeviceId:            managedDeviceId,
		DeviceConfigurationStateId: deviceConfigurationStateId,
	}
}

// ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateID parses 'input' into a DeviceManagementManagedDeviceIdDeviceConfigurationStateId
func ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateID(input string) (*DeviceManagementManagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdDeviceConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateIDInsensitively(input string) (*DeviceManagementManagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdDeviceConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceConfigurationStateId, ok = input.Parsed["deviceConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdDeviceConfigurationStateID checks that 'input' can be parsed as a Device Management Managed Device Id Device Configuration State ID
func ValidateDeviceManagementManagedDeviceIdDeviceConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Device Configuration State ID
func (id DeviceManagementManagedDeviceIdDeviceConfigurationStateId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/deviceConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Device Configuration State ID
func (id DeviceManagementManagedDeviceIdDeviceConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceConfigurationStates", "deviceConfigurationStates", "deviceConfigurationStates"),
		resourceids.UserSpecifiedSegment("deviceConfigurationStateId", "deviceConfigurationStateId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Device Configuration State ID
func (id DeviceManagementManagedDeviceIdDeviceConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Configuration State: %q", id.DeviceConfigurationStateId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Device Configuration State (%s)", strings.Join(components, "\n"))
}
