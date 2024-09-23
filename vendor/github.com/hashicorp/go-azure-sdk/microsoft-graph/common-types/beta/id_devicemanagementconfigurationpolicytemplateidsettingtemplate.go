package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{}

// DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId is a struct representing the Resource ID for a Device Management Configuration Policy Template Id Setting Template
type DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId struct {
	DeviceManagementConfigurationPolicyTemplateId  string
	DeviceManagementConfigurationSettingTemplateId string
}

// NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID returns a new DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId struct
func NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(deviceManagementConfigurationPolicyTemplateId string, deviceManagementConfigurationSettingTemplateId string) DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId {
	return DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{
		DeviceManagementConfigurationPolicyTemplateId:  deviceManagementConfigurationPolicyTemplateId,
		DeviceManagementConfigurationSettingTemplateId: deviceManagementConfigurationSettingTemplateId,
	}
}

// ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID parses 'input' into a DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId
func ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(input string) (*DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIDInsensitively(input string) (*DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyTemplateId, ok = input.Parsed["deviceManagementConfigurationPolicyTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyTemplateId", input)
	}

	if id.DeviceManagementConfigurationSettingTemplateId, ok = input.Parsed["deviceManagementConfigurationSettingTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingTemplateId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID checks that 'input' can be parsed as a Device Management Configuration Policy Template Id Setting Template ID
func ValidateDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy Template Id Setting Template ID
func (id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId) ID() string {
	fmtString := "/deviceManagement/configurationPolicyTemplates/%s/settingTemplates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyTemplateId, id.DeviceManagementConfigurationSettingTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy Template Id Setting Template ID
func (id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicyTemplates", "configurationPolicyTemplates", "configurationPolicyTemplates"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationPolicyTemplateId"),
		resourceids.StaticSegment("settingTemplates", "settingTemplates", "settingTemplates"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingTemplateId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy Template Id Setting Template ID
func (id DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy Template: %q", id.DeviceManagementConfigurationPolicyTemplateId),
		fmt.Sprintf("Device Management Configuration Setting Template: %q", id.DeviceManagementConfigurationSettingTemplateId),
	}
	return fmt.Sprintf("Device Management Configuration Policy Template Id Setting Template (%s)", strings.Join(components, "\n"))
}
