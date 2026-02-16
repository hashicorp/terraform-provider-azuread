package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateSettingId{}

// DeviceManagementTemplateSettingId is a struct representing the Resource ID for a Device Management Template Setting
type DeviceManagementTemplateSettingId struct {
	DeviceManagementConfigurationSettingTemplateId string
}

// NewDeviceManagementTemplateSettingID returns a new DeviceManagementTemplateSettingId struct
func NewDeviceManagementTemplateSettingID(deviceManagementConfigurationSettingTemplateId string) DeviceManagementTemplateSettingId {
	return DeviceManagementTemplateSettingId{
		DeviceManagementConfigurationSettingTemplateId: deviceManagementConfigurationSettingTemplateId,
	}
}

// ParseDeviceManagementTemplateSettingID parses 'input' into a DeviceManagementTemplateSettingId
func ParseDeviceManagementTemplateSettingID(input string) (*DeviceManagementTemplateSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateSettingIDInsensitively(input string) (*DeviceManagementTemplateSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationSettingTemplateId, ok = input.Parsed["deviceManagementConfigurationSettingTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingTemplateId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateSettingID checks that 'input' can be parsed as a Device Management Template Setting ID
func ValidateDeviceManagementTemplateSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Setting ID
func (id DeviceManagementTemplateSettingId) ID() string {
	fmtString := "/deviceManagement/templateSettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationSettingTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Setting ID
func (id DeviceManagementTemplateSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templateSettings", "templateSettings", "templateSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingTemplateId"),
	}
}

// String returns a human-readable description of this Device Management Template Setting ID
func (id DeviceManagementTemplateSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Setting Template: %q", id.DeviceManagementConfigurationSettingTemplateId),
	}
	return fmt.Sprintf("Device Management Template Setting (%s)", strings.Join(components, "\n"))
}
