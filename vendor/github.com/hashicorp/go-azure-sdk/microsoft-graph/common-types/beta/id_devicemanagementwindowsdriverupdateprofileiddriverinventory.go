package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{}

// DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId is a struct representing the Resource ID for a Device Management Windows Driver Update Profile Id Driver Inventory
type DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId struct {
	WindowsDriverUpdateProfileId   string
	WindowsDriverUpdateInventoryId string
}

// NewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID returns a new DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId struct
func NewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(windowsDriverUpdateProfileId string, windowsDriverUpdateInventoryId string) DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId {
	return DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
		WindowsDriverUpdateProfileId:   windowsDriverUpdateProfileId,
		WindowsDriverUpdateInventoryId: windowsDriverUpdateInventoryId,
	}
}

// ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID parses 'input' into a DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId
func ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(input string) (*DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryIDInsensitively(input string) (*DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsDriverUpdateProfileId, ok = input.Parsed["windowsDriverUpdateProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsDriverUpdateProfileId", input)
	}

	if id.WindowsDriverUpdateInventoryId, ok = input.Parsed["windowsDriverUpdateInventoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsDriverUpdateInventoryId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID checks that 'input' can be parsed as a Device Management Windows Driver Update Profile Id Driver Inventory ID
func ValidateDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Driver Update Profile Id Driver Inventory ID
func (id DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId) ID() string {
	fmtString := "/deviceManagement/windowsDriverUpdateProfiles/%s/driverInventories/%s"
	return fmt.Sprintf(fmtString, id.WindowsDriverUpdateProfileId, id.WindowsDriverUpdateInventoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Driver Update Profile Id Driver Inventory ID
func (id DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsDriverUpdateProfiles", "windowsDriverUpdateProfiles", "windowsDriverUpdateProfiles"),
		resourceids.UserSpecifiedSegment("windowsDriverUpdateProfileId", "windowsDriverUpdateProfileId"),
		resourceids.StaticSegment("driverInventories", "driverInventories", "driverInventories"),
		resourceids.UserSpecifiedSegment("windowsDriverUpdateInventoryId", "windowsDriverUpdateInventoryId"),
	}
}

// String returns a human-readable description of this Device Management Windows Driver Update Profile Id Driver Inventory ID
func (id DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId) String() string {
	components := []string{
		fmt.Sprintf("Windows Driver Update Profile: %q", id.WindowsDriverUpdateProfileId),
		fmt.Sprintf("Windows Driver Update Inventory: %q", id.WindowsDriverUpdateInventoryId),
	}
	return fmt.Sprintf("Device Management Windows Driver Update Profile Id Driver Inventory (%s)", strings.Join(components, "\n"))
}
