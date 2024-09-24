package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdUserStateId{}

// DeviceManagementIntentIdUserStateId is a struct representing the Resource ID for a Device Management Intent Id User State
type DeviceManagementIntentIdUserStateId struct {
	DeviceManagementIntentId          string
	DeviceManagementIntentUserStateId string
}

// NewDeviceManagementIntentIdUserStateID returns a new DeviceManagementIntentIdUserStateId struct
func NewDeviceManagementIntentIdUserStateID(deviceManagementIntentId string, deviceManagementIntentUserStateId string) DeviceManagementIntentIdUserStateId {
	return DeviceManagementIntentIdUserStateId{
		DeviceManagementIntentId:          deviceManagementIntentId,
		DeviceManagementIntentUserStateId: deviceManagementIntentUserStateId,
	}
}

// ParseDeviceManagementIntentIdUserStateID parses 'input' into a DeviceManagementIntentIdUserStateId
func ParseDeviceManagementIntentIdUserStateID(input string) (*DeviceManagementIntentIdUserStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdUserStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdUserStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdUserStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdUserStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdUserStateIDInsensitively(input string) (*DeviceManagementIntentIdUserStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdUserStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdUserStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdUserStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentUserStateId, ok = input.Parsed["deviceManagementIntentUserStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentUserStateId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdUserStateID checks that 'input' can be parsed as a Device Management Intent Id User State ID
func ValidateDeviceManagementIntentIdUserStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdUserStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id User State ID
func (id DeviceManagementIntentIdUserStateId) ID() string {
	fmtString := "/deviceManagement/intents/%s/userStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentUserStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id User State ID
func (id DeviceManagementIntentIdUserStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("userStates", "userStates", "userStates"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentUserStateId", "deviceManagementIntentUserStateId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id User State ID
func (id DeviceManagementIntentIdUserStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent User State: %q", id.DeviceManagementIntentUserStateId),
	}
	return fmt.Sprintf("Device Management Intent Id User State (%s)", strings.Join(components, "\n"))
}
