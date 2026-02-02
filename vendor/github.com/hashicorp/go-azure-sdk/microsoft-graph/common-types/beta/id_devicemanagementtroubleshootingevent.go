package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTroubleshootingEventId{}

// DeviceManagementTroubleshootingEventId is a struct representing the Resource ID for a Device Management Troubleshooting Event
type DeviceManagementTroubleshootingEventId struct {
	DeviceManagementTroubleshootingEventId string
}

// NewDeviceManagementTroubleshootingEventID returns a new DeviceManagementTroubleshootingEventId struct
func NewDeviceManagementTroubleshootingEventID(deviceManagementTroubleshootingEventId string) DeviceManagementTroubleshootingEventId {
	return DeviceManagementTroubleshootingEventId{
		DeviceManagementTroubleshootingEventId: deviceManagementTroubleshootingEventId,
	}
}

// ParseDeviceManagementTroubleshootingEventID parses 'input' into a DeviceManagementTroubleshootingEventId
func ParseDeviceManagementTroubleshootingEventID(input string) (*DeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a DeviceManagementTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTroubleshootingEventIDInsensitively(input string) (*DeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTroubleshootingEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTroubleshootingEventId, ok = input.Parsed["deviceManagementTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTroubleshootingEventId", input)
	}

	return nil
}

// ValidateDeviceManagementTroubleshootingEventID checks that 'input' can be parsed as a Device Management Troubleshooting Event ID
func ValidateDeviceManagementTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Troubleshooting Event ID
func (id DeviceManagementTroubleshootingEventId) ID() string {
	fmtString := "/deviceManagement/troubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Troubleshooting Event ID
func (id DeviceManagementTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("troubleshootingEvents", "troubleshootingEvents", "troubleshootingEvents"),
		resourceids.UserSpecifiedSegment("deviceManagementTroubleshootingEventId", "deviceManagementTroubleshootingEventId"),
	}
}

// String returns a human-readable description of this Device Management Troubleshooting Event ID
func (id DeviceManagementTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Troubleshooting Event: %q", id.DeviceManagementTroubleshootingEventId),
	}
	return fmt.Sprintf("Device Management Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
