package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceId{}

// DeviceManagementManagedDeviceId is a struct representing the Resource ID for a Device Management Managed Device
type DeviceManagementManagedDeviceId struct {
	ManagedDeviceId string
}

// NewDeviceManagementManagedDeviceID returns a new DeviceManagementManagedDeviceId struct
func NewDeviceManagementManagedDeviceID(managedDeviceId string) DeviceManagementManagedDeviceId {
	return DeviceManagementManagedDeviceId{
		ManagedDeviceId: managedDeviceId,
	}
}

// ParseDeviceManagementManagedDeviceID parses 'input' into a DeviceManagementManagedDeviceId
func ParseDeviceManagementManagedDeviceID(input string) (*DeviceManagementManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIDInsensitively(input string) (*DeviceManagementManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceID checks that 'input' can be parsed as a Device Management Managed Device ID
func ValidateDeviceManagementManagedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device ID
func (id DeviceManagementManagedDeviceId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device ID
func (id DeviceManagementManagedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device ID
func (id DeviceManagementManagedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
	}
	return fmt.Sprintf("Device Management Managed Device (%s)", strings.Join(components, "\n"))
}
