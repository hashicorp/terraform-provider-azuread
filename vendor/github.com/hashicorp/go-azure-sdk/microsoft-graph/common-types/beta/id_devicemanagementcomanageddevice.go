package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceId{}

// DeviceManagementComanagedDeviceId is a struct representing the Resource ID for a Device Management Comanaged Device
type DeviceManagementComanagedDeviceId struct {
	ManagedDeviceId string
}

// NewDeviceManagementComanagedDeviceID returns a new DeviceManagementComanagedDeviceId struct
func NewDeviceManagementComanagedDeviceID(managedDeviceId string) DeviceManagementComanagedDeviceId {
	return DeviceManagementComanagedDeviceId{
		ManagedDeviceId: managedDeviceId,
	}
}

// ParseDeviceManagementComanagedDeviceID parses 'input' into a DeviceManagementComanagedDeviceId
func ParseDeviceManagementComanagedDeviceID(input string) (*DeviceManagementComanagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIDInsensitively(input string) (*DeviceManagementComanagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceID checks that 'input' can be parsed as a Device Management Comanaged Device ID
func ValidateDeviceManagementComanagedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device ID
func (id DeviceManagementComanagedDeviceId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device ID
func (id DeviceManagementComanagedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device ID
func (id DeviceManagementComanagedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
	}
	return fmt.Sprintf("Device Management Comanaged Device (%s)", strings.Join(components, "\n"))
}
