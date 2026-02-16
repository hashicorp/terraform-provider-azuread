package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDetectedAppIdManagedDeviceId{}

// DeviceManagementDetectedAppIdManagedDeviceId is a struct representing the Resource ID for a Device Management Detected App Id Managed Device
type DeviceManagementDetectedAppIdManagedDeviceId struct {
	DetectedAppId   string
	ManagedDeviceId string
}

// NewDeviceManagementDetectedAppIdManagedDeviceID returns a new DeviceManagementDetectedAppIdManagedDeviceId struct
func NewDeviceManagementDetectedAppIdManagedDeviceID(detectedAppId string, managedDeviceId string) DeviceManagementDetectedAppIdManagedDeviceId {
	return DeviceManagementDetectedAppIdManagedDeviceId{
		DetectedAppId:   detectedAppId,
		ManagedDeviceId: managedDeviceId,
	}
}

// ParseDeviceManagementDetectedAppIdManagedDeviceID parses 'input' into a DeviceManagementDetectedAppIdManagedDeviceId
func ParseDeviceManagementDetectedAppIdManagedDeviceID(input string) (*DeviceManagementDetectedAppIdManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDetectedAppIdManagedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDetectedAppIdManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDetectedAppIdManagedDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementDetectedAppIdManagedDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDetectedAppIdManagedDeviceIDInsensitively(input string) (*DeviceManagementDetectedAppIdManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDetectedAppIdManagedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDetectedAppIdManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDetectedAppIdManagedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DetectedAppId, ok = input.Parsed["detectedAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementDetectedAppIdManagedDeviceID checks that 'input' can be parsed as a Device Management Detected App Id Managed Device ID
func ValidateDeviceManagementDetectedAppIdManagedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDetectedAppIdManagedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Detected App Id Managed Device ID
func (id DeviceManagementDetectedAppIdManagedDeviceId) ID() string {
	fmtString := "/deviceManagement/detectedApps/%s/managedDevices/%s"
	return fmt.Sprintf(fmtString, id.DetectedAppId, id.ManagedDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Detected App Id Managed Device ID
func (id DeviceManagementDetectedAppIdManagedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("detectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
	}
}

// String returns a human-readable description of this Device Management Detected App Id Managed Device ID
func (id DeviceManagementDetectedAppIdManagedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
	}
	return fmt.Sprintf("Device Management Detected App Id Managed Device (%s)", strings.Join(components, "\n"))
}
