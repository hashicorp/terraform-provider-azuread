package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicySettingStateSummaryId{}

// DeviceManagementDeviceCompliancePolicySettingStateSummaryId is a struct representing the Resource ID for a Device Management Device Compliance Policy Setting State Summary
type DeviceManagementDeviceCompliancePolicySettingStateSummaryId struct {
	DeviceCompliancePolicySettingStateSummaryId string
}

// NewDeviceManagementDeviceCompliancePolicySettingStateSummaryID returns a new DeviceManagementDeviceCompliancePolicySettingStateSummaryId struct
func NewDeviceManagementDeviceCompliancePolicySettingStateSummaryID(deviceCompliancePolicySettingStateSummaryId string) DeviceManagementDeviceCompliancePolicySettingStateSummaryId {
	return DeviceManagementDeviceCompliancePolicySettingStateSummaryId{
		DeviceCompliancePolicySettingStateSummaryId: deviceCompliancePolicySettingStateSummaryId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryID parses 'input' into a DeviceManagementDeviceCompliancePolicySettingStateSummaryId
func ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryID(input string) (*DeviceManagementDeviceCompliancePolicySettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicySettingStateSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicySettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicySettingStateSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicySettingStateSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicySettingStateSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicySettingStateSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicySettingStateSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicySettingStateSummaryId, ok = input.Parsed["deviceCompliancePolicySettingStateSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicySettingStateSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicySettingStateSummaryID checks that 'input' can be parsed as a Device Management Device Compliance Policy Setting State Summary ID
func ValidateDeviceManagementDeviceCompliancePolicySettingStateSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Setting State Summary ID
func (id DeviceManagementDeviceCompliancePolicySettingStateSummaryId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicySettingStateSummaries/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicySettingStateSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Setting State Summary ID
func (id DeviceManagementDeviceCompliancePolicySettingStateSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicySettingStateSummaries", "deviceCompliancePolicySettingStateSummaries", "deviceCompliancePolicySettingStateSummaries"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicySettingStateSummaryId", "deviceCompliancePolicySettingStateSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Setting State Summary ID
func (id DeviceManagementDeviceCompliancePolicySettingStateSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy Setting State Summary: %q", id.DeviceCompliancePolicySettingStateSummaryId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Setting State Summary (%s)", strings.Join(components, "\n"))
}
