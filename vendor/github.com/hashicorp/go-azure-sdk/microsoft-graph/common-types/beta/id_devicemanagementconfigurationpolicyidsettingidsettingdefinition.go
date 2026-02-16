package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{}

// DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Configuration Policy Id Setting Id Setting Definition
type DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId struct {
	DeviceManagementConfigurationPolicyId            string
	DeviceManagementConfigurationSettingId           string
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID returns a new DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId struct
func NewDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(deviceManagementConfigurationPolicyId string, deviceManagementConfigurationSettingId string, deviceManagementConfigurationSettingDefinitionId string) DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId {
	return DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{
		DeviceManagementConfigurationPolicyId:            deviceManagementConfigurationPolicyId,
		DeviceManagementConfigurationSettingId:           deviceManagementConfigurationSettingId,
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID parses 'input' into a DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId
func ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(input string) (*DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	if id.DeviceManagementConfigurationSettingId, ok = input.Parsed["deviceManagementConfigurationSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingId", input)
	}

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Configuration Policy Id Setting Id Setting Definition ID
func ValidateDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy Id Setting Id Setting Definition ID
func (id DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/configurationPolicies/%s/settings/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyId, id.DeviceManagementConfigurationSettingId, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy Id Setting Id Setting Definition ID
func (id DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicies", "configurationPolicies", "configurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingId", "deviceManagementConfigurationSettingId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy Id Setting Id Setting Definition ID
func (id DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
		fmt.Sprintf("Device Management Configuration Setting: %q", id.DeviceManagementConfigurationSettingId),
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Configuration Policy Id Setting Id Setting Definition (%s)", strings.Join(components, "\n"))
}
