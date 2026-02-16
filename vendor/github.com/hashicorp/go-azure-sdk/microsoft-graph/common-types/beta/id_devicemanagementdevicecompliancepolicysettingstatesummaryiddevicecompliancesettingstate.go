package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{}

// DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId is a struct representing the Resource ID for a Device Management Device Compliance Policy Setting State Summary Id Device Compliance Setting State
type DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId struct {
	DeviceCompliancePolicySettingStateSummaryId string
	DeviceComplianceSettingStateId              string
}

// NewDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID returns a new DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId struct
func NewDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(deviceCompliancePolicySettingStateSummaryId string, deviceComplianceSettingStateId string) DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId {
	return DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{
		DeviceCompliancePolicySettingStateSummaryId: deviceCompliancePolicySettingStateSummaryId,
		DeviceComplianceSettingStateId:              deviceComplianceSettingStateId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID parses 'input' into a DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId
func ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(input string) (*DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicySettingStateSummaryId, ok = input.Parsed["deviceCompliancePolicySettingStateSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicySettingStateSummaryId", input)
	}

	if id.DeviceComplianceSettingStateId, ok = input.Parsed["deviceComplianceSettingStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceSettingStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID checks that 'input' can be parsed as a Device Management Device Compliance Policy Setting State Summary Id Device Compliance Setting State ID
func ValidateDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Setting State Summary Id Device Compliance Setting State ID
func (id DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicySettingStateSummaries/%s/deviceComplianceSettingStates/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicySettingStateSummaryId, id.DeviceComplianceSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Setting State Summary Id Device Compliance Setting State ID
func (id DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicySettingStateSummaries", "deviceCompliancePolicySettingStateSummaries", "deviceCompliancePolicySettingStateSummaries"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicySettingStateSummaryId", "deviceCompliancePolicySettingStateSummaryId"),
		resourceids.StaticSegment("deviceComplianceSettingStates", "deviceComplianceSettingStates", "deviceComplianceSettingStates"),
		resourceids.UserSpecifiedSegment("deviceComplianceSettingStateId", "deviceComplianceSettingStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Setting State Summary Id Device Compliance Setting State ID
func (id DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy Setting State Summary: %q", id.DeviceCompliancePolicySettingStateSummaryId),
		fmt.Sprintf("Device Compliance Setting State: %q", id.DeviceComplianceSettingStateId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Setting State Summary Id Device Compliance Setting State (%s)", strings.Join(components, "\n"))
}
