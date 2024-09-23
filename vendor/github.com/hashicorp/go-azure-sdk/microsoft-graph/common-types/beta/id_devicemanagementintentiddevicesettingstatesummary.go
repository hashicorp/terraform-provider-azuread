package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdDeviceSettingStateSummaryId{}

// DeviceManagementIntentIdDeviceSettingStateSummaryId is a struct representing the Resource ID for a Device Management Intent Id Device Setting State Summary
type DeviceManagementIntentIdDeviceSettingStateSummaryId struct {
	DeviceManagementIntentId                          string
	DeviceManagementIntentDeviceSettingStateSummaryId string
}

// NewDeviceManagementIntentIdDeviceSettingStateSummaryID returns a new DeviceManagementIntentIdDeviceSettingStateSummaryId struct
func NewDeviceManagementIntentIdDeviceSettingStateSummaryID(deviceManagementIntentId string, deviceManagementIntentDeviceSettingStateSummaryId string) DeviceManagementIntentIdDeviceSettingStateSummaryId {
	return DeviceManagementIntentIdDeviceSettingStateSummaryId{
		DeviceManagementIntentId:                          deviceManagementIntentId,
		DeviceManagementIntentDeviceSettingStateSummaryId: deviceManagementIntentDeviceSettingStateSummaryId,
	}
}

// ParseDeviceManagementIntentIdDeviceSettingStateSummaryID parses 'input' into a DeviceManagementIntentIdDeviceSettingStateSummaryId
func ParseDeviceManagementIntentIdDeviceSettingStateSummaryID(input string) (*DeviceManagementIntentIdDeviceSettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdDeviceSettingStateSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdDeviceSettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdDeviceSettingStateSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdDeviceSettingStateSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdDeviceSettingStateSummaryIDInsensitively(input string) (*DeviceManagementIntentIdDeviceSettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdDeviceSettingStateSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdDeviceSettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdDeviceSettingStateSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentDeviceSettingStateSummaryId, ok = input.Parsed["deviceManagementIntentDeviceSettingStateSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentDeviceSettingStateSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdDeviceSettingStateSummaryID checks that 'input' can be parsed as a Device Management Intent Id Device Setting State Summary ID
func ValidateDeviceManagementIntentIdDeviceSettingStateSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdDeviceSettingStateSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Device Setting State Summary ID
func (id DeviceManagementIntentIdDeviceSettingStateSummaryId) ID() string {
	fmtString := "/deviceManagement/intents/%s/deviceSettingStateSummaries/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentDeviceSettingStateSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Device Setting State Summary ID
func (id DeviceManagementIntentIdDeviceSettingStateSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("deviceSettingStateSummaries", "deviceSettingStateSummaries", "deviceSettingStateSummaries"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentDeviceSettingStateSummaryId", "deviceManagementIntentDeviceSettingStateSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Device Setting State Summary ID
func (id DeviceManagementIntentIdDeviceSettingStateSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent Device Setting State Summary: %q", id.DeviceManagementIntentDeviceSettingStateSummaryId),
	}
	return fmt.Sprintf("Device Management Intent Id Device Setting State Summary (%s)", strings.Join(components, "\n"))
}
