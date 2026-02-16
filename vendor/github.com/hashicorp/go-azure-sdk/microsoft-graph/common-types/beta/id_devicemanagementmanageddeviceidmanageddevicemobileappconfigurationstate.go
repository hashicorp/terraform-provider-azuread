package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}

// DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId is a struct representing the Resource ID for a Device Management Managed Device Id Managed Device Mobile App Configuration State
type DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct {
	ManagedDeviceId                            string
	ManagedDeviceMobileAppConfigurationStateId string
}

// NewDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID returns a new DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct
func NewDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(managedDeviceId string, managedDeviceMobileAppConfigurationStateId string) DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId {
	return DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
		ManagedDeviceId: managedDeviceId,
		ManagedDeviceMobileAppConfigurationStateId: managedDeviceMobileAppConfigurationStateId,
	}
}

// ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID parses 'input' into a DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
func ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input string) (*DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(input string) (*DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = input.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID checks that 'input' can be parsed as a Device Management Managed Device Id Managed Device Mobile App Configuration State ID
func ValidateDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Managed Device Mobile App Configuration State ID
func (id DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/managedDeviceMobileAppConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.ManagedDeviceMobileAppConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Managed Device Mobile App Configuration State ID
func (id DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates"),
		resourceids.UserSpecifiedSegment("managedDeviceMobileAppConfigurationStateId", "managedDeviceMobileAppConfigurationStateId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Managed Device Mobile App Configuration State ID
func (id DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Managed Device Mobile App Configuration State: %q", id.ManagedDeviceMobileAppConfigurationStateId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Managed Device Mobile App Configuration State (%s)", strings.Join(components, "\n"))
}
