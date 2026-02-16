package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{}

// DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId is a struct representing the Resource ID for a Device Management Device Compliance Policy Id Device Setting State Summary
type DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId struct {
	DeviceCompliancePolicyId    string
	SettingStateDeviceSummaryId string
}

// NewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID returns a new DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId struct
func NewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(deviceCompliancePolicyId string, settingStateDeviceSummaryId string) DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId {
	return DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
		DeviceCompliancePolicyId:    deviceCompliancePolicyId,
		SettingStateDeviceSummaryId: settingStateDeviceSummaryId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID parses 'input' into a DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId
func ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(input string) (*DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	if id.SettingStateDeviceSummaryId, ok = input.Parsed["settingStateDeviceSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "settingStateDeviceSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID checks that 'input' can be parsed as a Device Management Device Compliance Policy Id Device Setting State Summary ID
func ValidateDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Id Device Setting State Summary ID
func (id DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s/deviceSettingStateSummaries/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId, id.SettingStateDeviceSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Id Device Setting State Summary ID
func (id DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
		resourceids.StaticSegment("deviceSettingStateSummaries", "deviceSettingStateSummaries", "deviceSettingStateSummaries"),
		resourceids.UserSpecifiedSegment("settingStateDeviceSummaryId", "settingStateDeviceSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Id Device Setting State Summary ID
func (id DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
		fmt.Sprintf("Setting State Device Summary: %q", id.SettingStateDeviceSummaryId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Id Device Setting State Summary (%s)", strings.Join(components, "\n"))
}
