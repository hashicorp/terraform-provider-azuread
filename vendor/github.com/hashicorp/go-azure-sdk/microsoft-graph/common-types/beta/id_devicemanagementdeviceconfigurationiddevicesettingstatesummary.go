package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{}

// DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId is a struct representing the Resource ID for a Device Management Device Configuration Id Device Setting State Summary
type DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId struct {
	DeviceConfigurationId       string
	SettingStateDeviceSummaryId string
}

// NewDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID returns a new DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId struct
func NewDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(deviceConfigurationId string, settingStateDeviceSummaryId string) DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId {
	return DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{
		DeviceConfigurationId:       deviceConfigurationId,
		SettingStateDeviceSummaryId: settingStateDeviceSummaryId,
	}
}

// ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID parses 'input' into a DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId
func ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(input string) (*DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryIDInsensitively(input string) (*DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationId, ok = input.Parsed["deviceConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationId", input)
	}

	if id.SettingStateDeviceSummaryId, ok = input.Parsed["settingStateDeviceSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "settingStateDeviceSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID checks that 'input' can be parsed as a Device Management Device Configuration Id Device Setting State Summary ID
func ValidateDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Id Device Setting State Summary ID
func (id DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurations/%s/deviceSettingStateSummaries/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationId, id.SettingStateDeviceSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Id Device Setting State Summary ID
func (id DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurations", "deviceConfigurations", "deviceConfigurations"),
		resourceids.UserSpecifiedSegment("deviceConfigurationId", "deviceConfigurationId"),
		resourceids.StaticSegment("deviceSettingStateSummaries", "deviceSettingStateSummaries", "deviceSettingStateSummaries"),
		resourceids.UserSpecifiedSegment("settingStateDeviceSummaryId", "settingStateDeviceSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Id Device Setting State Summary ID
func (id DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration: %q", id.DeviceConfigurationId),
		fmt.Sprintf("Setting State Device Summary: %q", id.SettingStateDeviceSummaryId),
	}
	return fmt.Sprintf("Device Management Device Configuration Id Device Setting State Summary (%s)", strings.Join(components, "\n"))
}
