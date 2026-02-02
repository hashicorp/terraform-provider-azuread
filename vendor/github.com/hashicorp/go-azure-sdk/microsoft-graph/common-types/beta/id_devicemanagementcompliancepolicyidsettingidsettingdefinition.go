package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId{}

// DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Compliance Policy Id Setting Id Setting Definition
type DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId struct {
	DeviceManagementCompliancePolicyId               string
	DeviceManagementConfigurationSettingId           string
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID returns a new DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId struct
func NewDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID(deviceManagementCompliancePolicyId string, deviceManagementConfigurationSettingId string, deviceManagementConfigurationSettingDefinitionId string) DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId {
	return DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId{
		DeviceManagementCompliancePolicyId:               deviceManagementCompliancePolicyId,
		DeviceManagementConfigurationSettingId:           deviceManagementConfigurationSettingId,
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID parses 'input' into a DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId
func ParseDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID(input string) (*DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCompliancePolicyId, ok = input.Parsed["deviceManagementCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCompliancePolicyId", input)
	}

	if id.DeviceManagementConfigurationSettingId, ok = input.Parsed["deviceManagementConfigurationSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingId", input)
	}

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Compliance Policy Id Setting Id Setting Definition ID
func ValidateDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCompliancePolicyIdSettingIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Policy Id Setting Id Setting Definition ID
func (id DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/compliancePolicies/%s/settings/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCompliancePolicyId, id.DeviceManagementConfigurationSettingId, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Policy Id Setting Id Setting Definition ID
func (id DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("compliancePolicies", "compliancePolicies", "compliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementCompliancePolicyId", "deviceManagementCompliancePolicyId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingId", "deviceManagementConfigurationSettingId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Policy Id Setting Id Setting Definition ID
func (id DeviceManagementCompliancePolicyIdSettingIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Compliance Policy: %q", id.DeviceManagementCompliancePolicyId),
		fmt.Sprintf("Device Management Configuration Setting: %q", id.DeviceManagementConfigurationSettingId),
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Compliance Policy Id Setting Id Setting Definition (%s)", strings.Join(components, "\n"))
}
