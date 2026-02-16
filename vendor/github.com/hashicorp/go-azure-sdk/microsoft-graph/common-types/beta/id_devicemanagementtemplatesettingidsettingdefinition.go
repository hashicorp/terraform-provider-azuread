package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateSettingIdSettingDefinitionId{}

// DeviceManagementTemplateSettingIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Template Setting Id Setting Definition
type DeviceManagementTemplateSettingIdSettingDefinitionId struct {
	DeviceManagementConfigurationSettingTemplateId   string
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementTemplateSettingIdSettingDefinitionID returns a new DeviceManagementTemplateSettingIdSettingDefinitionId struct
func NewDeviceManagementTemplateSettingIdSettingDefinitionID(deviceManagementConfigurationSettingTemplateId string, deviceManagementConfigurationSettingDefinitionId string) DeviceManagementTemplateSettingIdSettingDefinitionId {
	return DeviceManagementTemplateSettingIdSettingDefinitionId{
		DeviceManagementConfigurationSettingTemplateId:   deviceManagementConfigurationSettingTemplateId,
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementTemplateSettingIdSettingDefinitionID parses 'input' into a DeviceManagementTemplateSettingIdSettingDefinitionId
func ParseDeviceManagementTemplateSettingIdSettingDefinitionID(input string) (*DeviceManagementTemplateSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateSettingIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateSettingIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateSettingIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementTemplateSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateSettingIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationSettingTemplateId, ok = input.Parsed["deviceManagementConfigurationSettingTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingTemplateId", input)
	}

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateSettingIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Template Setting Id Setting Definition ID
func ValidateDeviceManagementTemplateSettingIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateSettingIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Setting Id Setting Definition ID
func (id DeviceManagementTemplateSettingIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/templateSettings/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationSettingTemplateId, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Setting Id Setting Definition ID
func (id DeviceManagementTemplateSettingIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templateSettings", "templateSettings", "templateSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingTemplateId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Template Setting Id Setting Definition ID
func (id DeviceManagementTemplateSettingIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Setting Template: %q", id.DeviceManagementConfigurationSettingTemplateId),
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Template Setting Id Setting Definition (%s)", strings.Join(components, "\n"))
}
