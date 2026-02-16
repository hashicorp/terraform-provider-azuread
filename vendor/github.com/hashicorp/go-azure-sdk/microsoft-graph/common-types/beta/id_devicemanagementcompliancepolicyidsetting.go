package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdSettingId{}

// DeviceManagementCompliancePolicyIdSettingId is a struct representing the Resource ID for a Device Management Compliance Policy Id Setting
type DeviceManagementCompliancePolicyIdSettingId struct {
	DeviceManagementCompliancePolicyId     string
	DeviceManagementConfigurationSettingId string
}

// NewDeviceManagementCompliancePolicyIdSettingID returns a new DeviceManagementCompliancePolicyIdSettingId struct
func NewDeviceManagementCompliancePolicyIdSettingID(deviceManagementCompliancePolicyId string, deviceManagementConfigurationSettingId string) DeviceManagementCompliancePolicyIdSettingId {
	return DeviceManagementCompliancePolicyIdSettingId{
		DeviceManagementCompliancePolicyId:     deviceManagementCompliancePolicyId,
		DeviceManagementConfigurationSettingId: deviceManagementConfigurationSettingId,
	}
}

// ParseDeviceManagementCompliancePolicyIdSettingID parses 'input' into a DeviceManagementCompliancePolicyIdSettingId
func ParseDeviceManagementCompliancePolicyIdSettingID(input string) (*DeviceManagementCompliancePolicyIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCompliancePolicyIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementCompliancePolicyIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCompliancePolicyIdSettingIDInsensitively(input string) (*DeviceManagementCompliancePolicyIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCompliancePolicyIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCompliancePolicyId, ok = input.Parsed["deviceManagementCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCompliancePolicyId", input)
	}

	if id.DeviceManagementConfigurationSettingId, ok = input.Parsed["deviceManagementConfigurationSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingId", input)
	}

	return nil
}

// ValidateDeviceManagementCompliancePolicyIdSettingID checks that 'input' can be parsed as a Device Management Compliance Policy Id Setting ID
func ValidateDeviceManagementCompliancePolicyIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCompliancePolicyIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Policy Id Setting ID
func (id DeviceManagementCompliancePolicyIdSettingId) ID() string {
	fmtString := "/deviceManagement/compliancePolicies/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCompliancePolicyId, id.DeviceManagementConfigurationSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Policy Id Setting ID
func (id DeviceManagementCompliancePolicyIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("compliancePolicies", "compliancePolicies", "compliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementCompliancePolicyId", "deviceManagementCompliancePolicyId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingId", "deviceManagementConfigurationSettingId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Policy Id Setting ID
func (id DeviceManagementCompliancePolicyIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Compliance Policy: %q", id.DeviceManagementCompliancePolicyId),
		fmt.Sprintf("Device Management Configuration Setting: %q", id.DeviceManagementConfigurationSettingId),
	}
	return fmt.Sprintf("Device Management Compliance Policy Id Setting (%s)", strings.Join(components, "\n"))
}
