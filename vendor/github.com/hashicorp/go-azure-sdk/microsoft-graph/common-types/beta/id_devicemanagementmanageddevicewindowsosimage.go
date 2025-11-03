package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceWindowsOSImageId{}

// DeviceManagementManagedDeviceWindowsOSImageId is a struct representing the Resource ID for a Device Management Managed Device Windows OS Image
type DeviceManagementManagedDeviceWindowsOSImageId struct {
	ManagedDeviceWindowsOperatingSystemImageId string
}

// NewDeviceManagementManagedDeviceWindowsOSImageID returns a new DeviceManagementManagedDeviceWindowsOSImageId struct
func NewDeviceManagementManagedDeviceWindowsOSImageID(managedDeviceWindowsOperatingSystemImageId string) DeviceManagementManagedDeviceWindowsOSImageId {
	return DeviceManagementManagedDeviceWindowsOSImageId{
		ManagedDeviceWindowsOperatingSystemImageId: managedDeviceWindowsOperatingSystemImageId,
	}
}

// ParseDeviceManagementManagedDeviceWindowsOSImageID parses 'input' into a DeviceManagementManagedDeviceWindowsOSImageId
func ParseDeviceManagementManagedDeviceWindowsOSImageID(input string) (*DeviceManagementManagedDeviceWindowsOSImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceWindowsOSImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceWindowsOSImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceWindowsOSImageIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceWindowsOSImageId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceWindowsOSImageIDInsensitively(input string) (*DeviceManagementManagedDeviceWindowsOSImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceWindowsOSImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceWindowsOSImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceWindowsOSImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceWindowsOperatingSystemImageId, ok = input.Parsed["managedDeviceWindowsOperatingSystemImageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceWindowsOperatingSystemImageId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceWindowsOSImageID checks that 'input' can be parsed as a Device Management Managed Device Windows OS Image ID
func ValidateDeviceManagementManagedDeviceWindowsOSImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceWindowsOSImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Windows OS Image ID
func (id DeviceManagementManagedDeviceWindowsOSImageId) ID() string {
	fmtString := "/deviceManagement/managedDeviceWindowsOSImages/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceWindowsOperatingSystemImageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Windows OS Image ID
func (id DeviceManagementManagedDeviceWindowsOSImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDeviceWindowsOSImages", "managedDeviceWindowsOSImages", "managedDeviceWindowsOSImages"),
		resourceids.UserSpecifiedSegment("managedDeviceWindowsOperatingSystemImageId", "managedDeviceWindowsOperatingSystemImageId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Windows OS Image ID
func (id DeviceManagementManagedDeviceWindowsOSImageId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device Windows Operating System Image: %q", id.ManagedDeviceWindowsOperatingSystemImageId),
	}
	return fmt.Sprintf("Device Management Managed Device Windows OS Image (%s)", strings.Join(components, "\n"))
}
