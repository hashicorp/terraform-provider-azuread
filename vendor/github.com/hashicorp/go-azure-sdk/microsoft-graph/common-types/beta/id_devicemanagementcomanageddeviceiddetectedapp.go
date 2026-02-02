package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdDetectedAppId{}

// DeviceManagementComanagedDeviceIdDetectedAppId is a struct representing the Resource ID for a Device Management Comanaged Device Id Detected App
type DeviceManagementComanagedDeviceIdDetectedAppId struct {
	ManagedDeviceId string
	DetectedAppId   string
}

// NewDeviceManagementComanagedDeviceIdDetectedAppID returns a new DeviceManagementComanagedDeviceIdDetectedAppId struct
func NewDeviceManagementComanagedDeviceIdDetectedAppID(managedDeviceId string, detectedAppId string) DeviceManagementComanagedDeviceIdDetectedAppId {
	return DeviceManagementComanagedDeviceIdDetectedAppId{
		ManagedDeviceId: managedDeviceId,
		DetectedAppId:   detectedAppId,
	}
}

// ParseDeviceManagementComanagedDeviceIdDetectedAppID parses 'input' into a DeviceManagementComanagedDeviceIdDetectedAppId
func ParseDeviceManagementComanagedDeviceIdDetectedAppID(input string) (*DeviceManagementComanagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdDetectedAppIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdDetectedAppIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdDetectedAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DetectedAppId, ok = input.Parsed["detectedAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdDetectedAppID checks that 'input' can be parsed as a Device Management Comanaged Device Id Detected App ID
func ValidateDeviceManagementComanagedDeviceIdDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Detected App ID
func (id DeviceManagementComanagedDeviceIdDetectedAppId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Detected App ID
func (id DeviceManagementComanagedDeviceIdDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("detectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Detected App ID
func (id DeviceManagementComanagedDeviceIdDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Detected App (%s)", strings.Join(components, "\n"))
}
