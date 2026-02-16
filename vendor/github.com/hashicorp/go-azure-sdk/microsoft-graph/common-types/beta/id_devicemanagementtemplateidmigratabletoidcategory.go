package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryId{}

// DeviceManagementTemplateIdMigratableToIdCategoryId is a struct representing the Resource ID for a Device Management Template Id Migratable To Id Category
type DeviceManagementTemplateIdMigratableToIdCategoryId struct {
	DeviceManagementTemplateId                string
	DeviceManagementTemplateId1               string
	DeviceManagementTemplateSettingCategoryId string
}

// NewDeviceManagementTemplateIdMigratableToIdCategoryID returns a new DeviceManagementTemplateIdMigratableToIdCategoryId struct
func NewDeviceManagementTemplateIdMigratableToIdCategoryID(deviceManagementTemplateId string, deviceManagementTemplateId1 string, deviceManagementTemplateSettingCategoryId string) DeviceManagementTemplateIdMigratableToIdCategoryId {
	return DeviceManagementTemplateIdMigratableToIdCategoryId{
		DeviceManagementTemplateId:                deviceManagementTemplateId,
		DeviceManagementTemplateId1:               deviceManagementTemplateId1,
		DeviceManagementTemplateSettingCategoryId: deviceManagementTemplateSettingCategoryId,
	}
}

// ParseDeviceManagementTemplateIdMigratableToIdCategoryID parses 'input' into a DeviceManagementTemplateIdMigratableToIdCategoryId
func ParseDeviceManagementTemplateIdMigratableToIdCategoryID(input string) (*DeviceManagementTemplateIdMigratableToIdCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdMigratableToIdCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdMigratableToIdCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdMigratableToIdCategoryIDInsensitively(input string) (*DeviceManagementTemplateIdMigratableToIdCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdMigratableToIdCategoryId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateDeviceManagementTemplateIdMigratableToIdCategoryID checks that 'input' can be parsed as a Device Management Template Id Migratable To Id Category ID
func ValidateDeviceManagementTemplateIdMigratableToIdCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Migratable To Id Category ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryId) ID() string {
	fmtString := "/deviceManagement/templates/%s/migratableTo/%s/categories/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateId1, id.DeviceManagementTemplateSettingCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Migratable To Id Category ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("migratableTo", "migratableTo", "migratableTo"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId1", "deviceManagementTemplateId1"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateSettingCategoryId", "deviceManagementTemplateSettingCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Migratable To Id Category ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Id 1: %q", id.DeviceManagementTemplateId1),
		fmt.Sprintf("Device Management Template Setting Category: %q", id.DeviceManagementTemplateSettingCategoryId),
	}
	return fmt.Sprintf("Device Management Template Id Migratable To Id Category (%s)", strings.Join(components, "\n"))
}
