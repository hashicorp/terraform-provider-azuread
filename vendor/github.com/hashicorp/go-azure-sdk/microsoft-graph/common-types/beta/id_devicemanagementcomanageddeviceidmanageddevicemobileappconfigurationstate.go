package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}

// DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId is a struct representing the Resource ID for a Device Management Comanaged Device Id Managed Device Mobile App Configuration State
type DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct {
	ManagedDeviceId                            string
	ManagedDeviceMobileAppConfigurationStateId string
}

// NewDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID returns a new DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct
func NewDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(managedDeviceId string, managedDeviceMobileAppConfigurationStateId string) DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId {
	return DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
		ManagedDeviceId: managedDeviceId,
		ManagedDeviceMobileAppConfigurationStateId: managedDeviceMobileAppConfigurationStateId,
	}
}

// ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID parses 'input' into a DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId
func ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input string) (*DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = input.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID checks that 'input' can be parsed as a Device Management Comanaged Device Id Managed Device Mobile App Configuration State ID
func ValidateDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Managed Device Mobile App Configuration State ID
func (id DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/managedDeviceMobileAppConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.ManagedDeviceMobileAppConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Managed Device Mobile App Configuration State ID
func (id DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates"),
		resourceids.UserSpecifiedSegment("managedDeviceMobileAppConfigurationStateId", "managedDeviceMobileAppConfigurationStateId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Managed Device Mobile App Configuration State ID
func (id DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Managed Device Mobile App Configuration State: %q", id.ManagedDeviceMobileAppConfigurationStateId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Managed Device Mobile App Configuration State (%s)", strings.Join(components, "\n"))
}
