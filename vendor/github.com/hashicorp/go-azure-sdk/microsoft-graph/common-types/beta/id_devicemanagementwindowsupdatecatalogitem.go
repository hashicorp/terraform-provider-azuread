package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsUpdateCatalogItemId{}

// DeviceManagementWindowsUpdateCatalogItemId is a struct representing the Resource ID for a Device Management Windows Update Catalog Item
type DeviceManagementWindowsUpdateCatalogItemId struct {
	WindowsUpdateCatalogItemId string
}

// NewDeviceManagementWindowsUpdateCatalogItemID returns a new DeviceManagementWindowsUpdateCatalogItemId struct
func NewDeviceManagementWindowsUpdateCatalogItemID(windowsUpdateCatalogItemId string) DeviceManagementWindowsUpdateCatalogItemId {
	return DeviceManagementWindowsUpdateCatalogItemId{
		WindowsUpdateCatalogItemId: windowsUpdateCatalogItemId,
	}
}

// ParseDeviceManagementWindowsUpdateCatalogItemID parses 'input' into a DeviceManagementWindowsUpdateCatalogItemId
func ParseDeviceManagementWindowsUpdateCatalogItemID(input string) (*DeviceManagementWindowsUpdateCatalogItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsUpdateCatalogItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsUpdateCatalogItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementWindowsUpdateCatalogItemIDInsensitively parses 'input' case-insensitively into a DeviceManagementWindowsUpdateCatalogItemId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementWindowsUpdateCatalogItemIDInsensitively(input string) (*DeviceManagementWindowsUpdateCatalogItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementWindowsUpdateCatalogItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementWindowsUpdateCatalogItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementWindowsUpdateCatalogItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsUpdateCatalogItemId, ok = input.Parsed["windowsUpdateCatalogItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsUpdateCatalogItemId", input)
	}

	return nil
}

// ValidateDeviceManagementWindowsUpdateCatalogItemID checks that 'input' can be parsed as a Device Management Windows Update Catalog Item ID
func ValidateDeviceManagementWindowsUpdateCatalogItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementWindowsUpdateCatalogItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Windows Update Catalog Item ID
func (id DeviceManagementWindowsUpdateCatalogItemId) ID() string {
	fmtString := "/deviceManagement/windowsUpdateCatalogItems/%s"
	return fmt.Sprintf(fmtString, id.WindowsUpdateCatalogItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Windows Update Catalog Item ID
func (id DeviceManagementWindowsUpdateCatalogItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("windowsUpdateCatalogItems", "windowsUpdateCatalogItems", "windowsUpdateCatalogItems"),
		resourceids.UserSpecifiedSegment("windowsUpdateCatalogItemId", "windowsUpdateCatalogItemId"),
	}
}

// String returns a human-readable description of this Device Management Windows Update Catalog Item ID
func (id DeviceManagementWindowsUpdateCatalogItemId) String() string {
	components := []string{
		fmt.Sprintf("Windows Update Catalog Item: %q", id.WindowsUpdateCatalogItemId),
	}
	return fmt.Sprintf("Device Management Windows Update Catalog Item (%s)", strings.Join(components, "\n"))
}
