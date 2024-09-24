package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentId{}

// DeviceManagementIntentId is a struct representing the Resource ID for a Device Management Intent
type DeviceManagementIntentId struct {
	DeviceManagementIntentId string
}

// NewDeviceManagementIntentID returns a new DeviceManagementIntentId struct
func NewDeviceManagementIntentID(deviceManagementIntentId string) DeviceManagementIntentId {
	return DeviceManagementIntentId{
		DeviceManagementIntentId: deviceManagementIntentId,
	}
}

// ParseDeviceManagementIntentID parses 'input' into a DeviceManagementIntentId
func ParseDeviceManagementIntentID(input string) (*DeviceManagementIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIDInsensitively(input string) (*DeviceManagementIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentID checks that 'input' can be parsed as a Device Management Intent ID
func ValidateDeviceManagementIntentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent ID
func (id DeviceManagementIntentId) ID() string {
	fmtString := "/deviceManagement/intents/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent ID
func (id DeviceManagementIntentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
	}
}

// String returns a human-readable description of this Device Management Intent ID
func (id DeviceManagementIntentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
	}
	return fmt.Sprintf("Device Management Intent (%s)", strings.Join(components, "\n"))
}
