package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDetectedAppId{}

// DeviceManagementDetectedAppId is a struct representing the Resource ID for a Device Management Detected App
type DeviceManagementDetectedAppId struct {
	DetectedAppId string
}

// NewDeviceManagementDetectedAppID returns a new DeviceManagementDetectedAppId struct
func NewDeviceManagementDetectedAppID(detectedAppId string) DeviceManagementDetectedAppId {
	return DeviceManagementDetectedAppId{
		DetectedAppId: detectedAppId,
	}
}

// ParseDeviceManagementDetectedAppID parses 'input' into a DeviceManagementDetectedAppId
func ParseDeviceManagementDetectedAppID(input string) (*DeviceManagementDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDetectedAppIDInsensitively parses 'input' case-insensitively into a DeviceManagementDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDetectedAppIDInsensitively(input string) (*DeviceManagementDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDetectedAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DetectedAppId, ok = input.Parsed["detectedAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", input)
	}

	return nil
}

// ValidateDeviceManagementDetectedAppID checks that 'input' can be parsed as a Device Management Detected App ID
func ValidateDeviceManagementDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Detected App ID
func (id DeviceManagementDetectedAppId) ID() string {
	fmtString := "/deviceManagement/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Detected App ID
func (id DeviceManagementDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("detectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppId"),
	}
}

// String returns a human-readable description of this Device Management Detected App ID
func (id DeviceManagementDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("Device Management Detected App (%s)", strings.Join(components, "\n"))
}
