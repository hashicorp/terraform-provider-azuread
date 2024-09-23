package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{}

// DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Template Id Migratable To Id Category Id Setting Definition
type DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId struct {
	DeviceManagementTemplateId                string
	DeviceManagementTemplateId1               string
	DeviceManagementTemplateSettingCategoryId string
	DeviceManagementSettingDefinitionId       string
}

// NewDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID returns a new DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId struct
func NewDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(deviceManagementTemplateId string, deviceManagementTemplateId1 string, deviceManagementTemplateSettingCategoryId string, deviceManagementSettingDefinitionId string) DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId {
	return DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{
		DeviceManagementTemplateId:                deviceManagementTemplateId,
		DeviceManagementTemplateId1:               deviceManagementTemplateId1,
		DeviceManagementTemplateSettingCategoryId: deviceManagementTemplateSettingCategoryId,
		DeviceManagementSettingDefinitionId:       deviceManagementSettingDefinitionId,
	}
}

// ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID parses 'input' into a DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId
func ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(input string) (*DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementTemplateId1, ok = input.Parsed["deviceManagementTemplateId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId1", input)
	}

	if id.DeviceManagementTemplateSettingCategoryId, ok = input.Parsed["deviceManagementTemplateSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateSettingCategoryId", input)
	}

	if id.DeviceManagementSettingDefinitionId, ok = input.Parsed["deviceManagementSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Template Id Migratable To Id Category Id Setting Definition ID
func ValidateDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Migratable To Id Category Id Setting Definition ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/templates/%s/migratableTo/%s/categories/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateId1, id.DeviceManagementTemplateSettingCategoryId, id.DeviceManagementSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Migratable To Id Category Id Setting Definition ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("migratableTo", "migratableTo", "migratableTo"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId1", "deviceManagementTemplateId1"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateSettingCategoryId", "deviceManagementTemplateSettingCategoryId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingDefinitionId", "deviceManagementSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Migratable To Id Category Id Setting Definition ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Id 1: %q", id.DeviceManagementTemplateId1),
		fmt.Sprintf("Device Management Template Setting Category: %q", id.DeviceManagementTemplateSettingCategoryId),
		fmt.Sprintf("Device Management Setting Definition: %q", id.DeviceManagementSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Template Id Migratable To Id Category Id Setting Definition (%s)", strings.Join(components, "\n"))
}
