package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdAssignmentId{}

// DeviceManagementIntentIdAssignmentId is a struct representing the Resource ID for a Device Management Intent Id Assignment
type DeviceManagementIntentIdAssignmentId struct {
	DeviceManagementIntentId           string
	DeviceManagementIntentAssignmentId string
}

// NewDeviceManagementIntentIdAssignmentID returns a new DeviceManagementIntentIdAssignmentId struct
func NewDeviceManagementIntentIdAssignmentID(deviceManagementIntentId string, deviceManagementIntentAssignmentId string) DeviceManagementIntentIdAssignmentId {
	return DeviceManagementIntentIdAssignmentId{
		DeviceManagementIntentId:           deviceManagementIntentId,
		DeviceManagementIntentAssignmentId: deviceManagementIntentAssignmentId,
	}
}

// ParseDeviceManagementIntentIdAssignmentID parses 'input' into a DeviceManagementIntentIdAssignmentId
func ParseDeviceManagementIntentIdAssignmentID(input string) (*DeviceManagementIntentIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdAssignmentIDInsensitively(input string) (*DeviceManagementIntentIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentAssignmentId, ok = input.Parsed["deviceManagementIntentAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdAssignmentID checks that 'input' can be parsed as a Device Management Intent Id Assignment ID
func ValidateDeviceManagementIntentIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Assignment ID
func (id DeviceManagementIntentIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/intents/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Assignment ID
func (id DeviceManagementIntentIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentAssignmentId", "deviceManagementIntentAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Assignment ID
func (id DeviceManagementIntentIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent Assignment: %q", id.DeviceManagementIntentAssignmentId),
	}
	return fmt.Sprintf("Device Management Intent Id Assignment (%s)", strings.Join(components, "\n"))
}
