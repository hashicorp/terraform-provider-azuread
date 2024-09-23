package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdDeviceStateId{}

// DeviceManagementIntentIdDeviceStateId is a struct representing the Resource ID for a Device Management Intent Id Device State
type DeviceManagementIntentIdDeviceStateId struct {
	DeviceManagementIntentId            string
	DeviceManagementIntentDeviceStateId string
}

// NewDeviceManagementIntentIdDeviceStateID returns a new DeviceManagementIntentIdDeviceStateId struct
func NewDeviceManagementIntentIdDeviceStateID(deviceManagementIntentId string, deviceManagementIntentDeviceStateId string) DeviceManagementIntentIdDeviceStateId {
	return DeviceManagementIntentIdDeviceStateId{
		DeviceManagementIntentId:            deviceManagementIntentId,
		DeviceManagementIntentDeviceStateId: deviceManagementIntentDeviceStateId,
	}
}

// ParseDeviceManagementIntentIdDeviceStateID parses 'input' into a DeviceManagementIntentIdDeviceStateId
func ParseDeviceManagementIntentIdDeviceStateID(input string) (*DeviceManagementIntentIdDeviceStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdDeviceStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdDeviceStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdDeviceStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdDeviceStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdDeviceStateIDInsensitively(input string) (*DeviceManagementIntentIdDeviceStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdDeviceStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdDeviceStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdDeviceStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentDeviceStateId, ok = input.Parsed["deviceManagementIntentDeviceStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentDeviceStateId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdDeviceStateID checks that 'input' can be parsed as a Device Management Intent Id Device State ID
func ValidateDeviceManagementIntentIdDeviceStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdDeviceStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Device State ID
func (id DeviceManagementIntentIdDeviceStateId) ID() string {
	fmtString := "/deviceManagement/intents/%s/deviceStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentDeviceStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Device State ID
func (id DeviceManagementIntentIdDeviceStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("deviceStates", "deviceStates", "deviceStates"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentDeviceStateId", "deviceManagementIntentDeviceStateId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Device State ID
func (id DeviceManagementIntentIdDeviceStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent Device State: %q", id.DeviceManagementIntentDeviceStateId),
	}
	return fmt.Sprintf("Device Management Intent Id Device State (%s)", strings.Join(components, "\n"))
}
