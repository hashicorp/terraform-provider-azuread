package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdSettingId{}

// DeviceManagementIntentIdSettingId is a struct representing the Resource ID for a Device Management Intent Id Setting
type DeviceManagementIntentIdSettingId struct {
	DeviceManagementIntentId          string
	DeviceManagementSettingInstanceId string
}

// NewDeviceManagementIntentIdSettingID returns a new DeviceManagementIntentIdSettingId struct
func NewDeviceManagementIntentIdSettingID(deviceManagementIntentId string, deviceManagementSettingInstanceId string) DeviceManagementIntentIdSettingId {
	return DeviceManagementIntentIdSettingId{
		DeviceManagementIntentId:          deviceManagementIntentId,
		DeviceManagementSettingInstanceId: deviceManagementSettingInstanceId,
	}
}

// ParseDeviceManagementIntentIdSettingID parses 'input' into a DeviceManagementIntentIdSettingId
func ParseDeviceManagementIntentIdSettingID(input string) (*DeviceManagementIntentIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdSettingIDInsensitively(input string) (*DeviceManagementIntentIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementSettingInstanceId, ok = input.Parsed["deviceManagementSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingInstanceId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdSettingID checks that 'input' can be parsed as a Device Management Intent Id Setting ID
func ValidateDeviceManagementIntentIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Setting ID
func (id DeviceManagementIntentIdSettingId) ID() string {
	fmtString := "/deviceManagement/intents/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Setting ID
func (id DeviceManagementIntentIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingInstanceId", "deviceManagementSettingInstanceId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Setting ID
func (id DeviceManagementIntentIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Setting Instance: %q", id.DeviceManagementSettingInstanceId),
	}
	return fmt.Sprintf("Device Management Intent Id Setting (%s)", strings.Join(components, "\n"))
}
