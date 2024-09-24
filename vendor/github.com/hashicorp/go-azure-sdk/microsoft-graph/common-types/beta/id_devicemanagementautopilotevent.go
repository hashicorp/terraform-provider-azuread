package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAutopilotEventId{}

// DeviceManagementAutopilotEventId is a struct representing the Resource ID for a Device Management Autopilot Event
type DeviceManagementAutopilotEventId struct {
	DeviceManagementAutopilotEventId string
}

// NewDeviceManagementAutopilotEventID returns a new DeviceManagementAutopilotEventId struct
func NewDeviceManagementAutopilotEventID(deviceManagementAutopilotEventId string) DeviceManagementAutopilotEventId {
	return DeviceManagementAutopilotEventId{
		DeviceManagementAutopilotEventId: deviceManagementAutopilotEventId,
	}
}

// ParseDeviceManagementAutopilotEventID parses 'input' into a DeviceManagementAutopilotEventId
func ParseDeviceManagementAutopilotEventID(input string) (*DeviceManagementAutopilotEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAutopilotEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAutopilotEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAutopilotEventIDInsensitively parses 'input' case-insensitively into a DeviceManagementAutopilotEventId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAutopilotEventIDInsensitively(input string) (*DeviceManagementAutopilotEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAutopilotEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAutopilotEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAutopilotEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementAutopilotEventId, ok = input.Parsed["deviceManagementAutopilotEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementAutopilotEventId", input)
	}

	return nil
}

// ValidateDeviceManagementAutopilotEventID checks that 'input' can be parsed as a Device Management Autopilot Event ID
func ValidateDeviceManagementAutopilotEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAutopilotEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Autopilot Event ID
func (id DeviceManagementAutopilotEventId) ID() string {
	fmtString := "/deviceManagement/autopilotEvents/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementAutopilotEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Autopilot Event ID
func (id DeviceManagementAutopilotEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("autopilotEvents", "autopilotEvents", "autopilotEvents"),
		resourceids.UserSpecifiedSegment("deviceManagementAutopilotEventId", "deviceManagementAutopilotEventId"),
	}
}

// String returns a human-readable description of this Device Management Autopilot Event ID
func (id DeviceManagementAutopilotEventId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Autopilot Event: %q", id.DeviceManagementAutopilotEventId),
	}
	return fmt.Sprintf("Device Management Autopilot Event (%s)", strings.Join(components, "\n"))
}
