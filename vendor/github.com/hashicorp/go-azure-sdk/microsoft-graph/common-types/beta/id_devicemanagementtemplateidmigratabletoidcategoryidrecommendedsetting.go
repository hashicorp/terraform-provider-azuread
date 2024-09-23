package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{}

// DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId is a struct representing the Resource ID for a Device Management Template Id Migratable To Id Category Id Recommended Setting
type DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId struct {
	DeviceManagementTemplateId                string
	DeviceManagementTemplateId1               string
	DeviceManagementTemplateSettingCategoryId string
	DeviceManagementSettingInstanceId         string
}

// NewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID returns a new DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId struct
func NewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(deviceManagementTemplateId string, deviceManagementTemplateId1 string, deviceManagementTemplateSettingCategoryId string, deviceManagementSettingInstanceId string) DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId {
	return DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
		DeviceManagementTemplateId:                deviceManagementTemplateId,
		DeviceManagementTemplateId1:               deviceManagementTemplateId1,
		DeviceManagementTemplateSettingCategoryId: deviceManagementTemplateSettingCategoryId,
		DeviceManagementSettingInstanceId:         deviceManagementSettingInstanceId,
	}
}

// ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID parses 'input' into a DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId
func ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(input string) (*DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingIDInsensitively(input string) (*DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DeviceManagementSettingInstanceId, ok = input.Parsed["deviceManagementSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingInstanceId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID checks that 'input' can be parsed as a Device Management Template Id Migratable To Id Category Id Recommended Setting ID
func ValidateDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Migratable To Id Category Id Recommended Setting ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId) ID() string {
	fmtString := "/deviceManagement/templates/%s/migratableTo/%s/categories/%s/recommendedSettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateId1, id.DeviceManagementTemplateSettingCategoryId, id.DeviceManagementSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Migratable To Id Category Id Recommended Setting ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("migratableTo", "migratableTo", "migratableTo"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId1", "deviceManagementTemplateId1"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateSettingCategoryId", "deviceManagementTemplateSettingCategoryId"),
		resourceids.StaticSegment("recommendedSettings", "recommendedSettings", "recommendedSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingInstanceId", "deviceManagementSettingInstanceId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Migratable To Id Category Id Recommended Setting ID
func (id DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Id 1: %q", id.DeviceManagementTemplateId1),
		fmt.Sprintf("Device Management Template Setting Category: %q", id.DeviceManagementTemplateSettingCategoryId),
		fmt.Sprintf("Device Management Setting Instance: %q", id.DeviceManagementSettingInstanceId),
	}
	return fmt.Sprintf("Device Management Template Id Migratable To Id Category Id Recommended Setting (%s)", strings.Join(components, "\n"))
}
