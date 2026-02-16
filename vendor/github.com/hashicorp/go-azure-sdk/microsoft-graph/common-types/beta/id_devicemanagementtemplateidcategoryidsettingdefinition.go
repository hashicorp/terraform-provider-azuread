package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdCategoryIdSettingDefinitionId{}

// DeviceManagementTemplateIdCategoryIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Template Id Category Id Setting Definition
type DeviceManagementTemplateIdCategoryIdSettingDefinitionId struct {
	DeviceManagementTemplateId                string
	DeviceManagementTemplateSettingCategoryId string
	DeviceManagementSettingDefinitionId       string
}

// NewDeviceManagementTemplateIdCategoryIdSettingDefinitionID returns a new DeviceManagementTemplateIdCategoryIdSettingDefinitionId struct
func NewDeviceManagementTemplateIdCategoryIdSettingDefinitionID(deviceManagementTemplateId string, deviceManagementTemplateSettingCategoryId string, deviceManagementSettingDefinitionId string) DeviceManagementTemplateIdCategoryIdSettingDefinitionId {
	return DeviceManagementTemplateIdCategoryIdSettingDefinitionId{
		DeviceManagementTemplateId:                deviceManagementTemplateId,
		DeviceManagementTemplateSettingCategoryId: deviceManagementTemplateSettingCategoryId,
		DeviceManagementSettingDefinitionId:       deviceManagementSettingDefinitionId,
	}
}

// ParseDeviceManagementTemplateIdCategoryIdSettingDefinitionID parses 'input' into a DeviceManagementTemplateIdCategoryIdSettingDefinitionId
func ParseDeviceManagementTemplateIdCategoryIdSettingDefinitionID(input string) (*DeviceManagementTemplateIdCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdCategoryIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdCategoryIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdCategoryIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementTemplateIdCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdCategoryIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementTemplateSettingCategoryId, ok = input.Parsed["deviceManagementTemplateSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateSettingCategoryId", input)
	}

	if id.DeviceManagementSettingDefinitionId, ok = input.Parsed["deviceManagementSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdCategoryIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Template Id Category Id Setting Definition ID
func ValidateDeviceManagementTemplateIdCategoryIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdCategoryIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Category Id Setting Definition ID
func (id DeviceManagementTemplateIdCategoryIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/templates/%s/categories/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateSettingCategoryId, id.DeviceManagementSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Category Id Setting Definition ID
func (id DeviceManagementTemplateIdCategoryIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateSettingCategoryId", "deviceManagementTemplateSettingCategoryId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingDefinitionId", "deviceManagementSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Category Id Setting Definition ID
func (id DeviceManagementTemplateIdCategoryIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Setting Category: %q", id.DeviceManagementTemplateSettingCategoryId),
		fmt.Sprintf("Device Management Setting Definition: %q", id.DeviceManagementSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Template Id Category Id Setting Definition (%s)", strings.Join(components, "\n"))
}
