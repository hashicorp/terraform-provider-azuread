package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{}

// DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Configuration Policy Template Id Setting Template Id Setting Definition
type DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId struct {
	DeviceManagementConfigurationPolicyTemplateId    string
	DeviceManagementConfigurationSettingTemplateId   string
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID returns a new DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId struct
func NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(deviceManagementConfigurationPolicyTemplateId string, deviceManagementConfigurationSettingTemplateId string, deviceManagementConfigurationSettingDefinitionId string) DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId {
	return DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
		DeviceManagementConfigurationPolicyTemplateId:    deviceManagementConfigurationPolicyTemplateId,
		DeviceManagementConfigurationSettingTemplateId:   deviceManagementConfigurationSettingTemplateId,
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID parses 'input' into a DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId
func ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(input string) (*DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyTemplateId, ok = input.Parsed["deviceManagementConfigurationPolicyTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyTemplateId", input)
	}

	if id.DeviceManagementConfigurationSettingTemplateId, ok = input.Parsed["deviceManagementConfigurationSettingTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingTemplateId", input)
	}

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Configuration Policy Template Id Setting Template Id Setting Definition ID
func ValidateDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy Template Id Setting Template Id Setting Definition ID
func (id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/configurationPolicyTemplates/%s/settingTemplates/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyTemplateId, id.DeviceManagementConfigurationSettingTemplateId, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy Template Id Setting Template Id Setting Definition ID
func (id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicyTemplates", "configurationPolicyTemplates", "configurationPolicyTemplates"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationPolicyTemplateId"),
		resourceids.StaticSegment("settingTemplates", "settingTemplates", "settingTemplates"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingTemplateId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy Template Id Setting Template Id Setting Definition ID
func (id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy Template: %q", id.DeviceManagementConfigurationPolicyTemplateId),
		fmt.Sprintf("Device Management Configuration Setting Template: %q", id.DeviceManagementConfigurationSettingTemplateId),
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Configuration Policy Template Id Setting Template Id Setting Definition (%s)", strings.Join(components, "\n"))
}
