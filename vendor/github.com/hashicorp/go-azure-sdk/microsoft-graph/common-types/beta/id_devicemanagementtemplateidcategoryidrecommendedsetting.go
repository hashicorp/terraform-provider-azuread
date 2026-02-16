package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdCategoryIdRecommendedSettingId{}

// DeviceManagementTemplateIdCategoryIdRecommendedSettingId is a struct representing the Resource ID for a Device Management Template Id Category Id Recommended Setting
type DeviceManagementTemplateIdCategoryIdRecommendedSettingId struct {
	DeviceManagementTemplateId                string
	DeviceManagementTemplateSettingCategoryId string
	DeviceManagementSettingInstanceId         string
}

// NewDeviceManagementTemplateIdCategoryIdRecommendedSettingID returns a new DeviceManagementTemplateIdCategoryIdRecommendedSettingId struct
func NewDeviceManagementTemplateIdCategoryIdRecommendedSettingID(deviceManagementTemplateId string, deviceManagementTemplateSettingCategoryId string, deviceManagementSettingInstanceId string) DeviceManagementTemplateIdCategoryIdRecommendedSettingId {
	return DeviceManagementTemplateIdCategoryIdRecommendedSettingId{
		DeviceManagementTemplateId:                deviceManagementTemplateId,
		DeviceManagementTemplateSettingCategoryId: deviceManagementTemplateSettingCategoryId,
		DeviceManagementSettingInstanceId:         deviceManagementSettingInstanceId,
	}
}

// ParseDeviceManagementTemplateIdCategoryIdRecommendedSettingID parses 'input' into a DeviceManagementTemplateIdCategoryIdRecommendedSettingId
func ParseDeviceManagementTemplateIdCategoryIdRecommendedSettingID(input string) (*DeviceManagementTemplateIdCategoryIdRecommendedSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdCategoryIdRecommendedSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdCategoryIdRecommendedSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdCategoryIdRecommendedSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdCategoryIdRecommendedSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdCategoryIdRecommendedSettingIDInsensitively(input string) (*DeviceManagementTemplateIdCategoryIdRecommendedSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdCategoryIdRecommendedSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdCategoryIdRecommendedSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdCategoryIdRecommendedSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementTemplateSettingCategoryId, ok = input.Parsed["deviceManagementTemplateSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateSettingCategoryId", input)
	}

	if id.DeviceManagementSettingInstanceId, ok = input.Parsed["deviceManagementSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingInstanceId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdCategoryIdRecommendedSettingID checks that 'input' can be parsed as a Device Management Template Id Category Id Recommended Setting ID
func ValidateDeviceManagementTemplateIdCategoryIdRecommendedSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdCategoryIdRecommendedSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Category Id Recommended Setting ID
func (id DeviceManagementTemplateIdCategoryIdRecommendedSettingId) ID() string {
	fmtString := "/deviceManagement/templates/%s/categories/%s/recommendedSettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateSettingCategoryId, id.DeviceManagementSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Category Id Recommended Setting ID
func (id DeviceManagementTemplateIdCategoryIdRecommendedSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateSettingCategoryId", "deviceManagementTemplateSettingCategoryId"),
		resourceids.StaticSegment("recommendedSettings", "recommendedSettings", "recommendedSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingInstanceId", "deviceManagementSettingInstanceId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Category Id Recommended Setting ID
func (id DeviceManagementTemplateIdCategoryIdRecommendedSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Setting Category: %q", id.DeviceManagementTemplateSettingCategoryId),
		fmt.Sprintf("Device Management Setting Instance: %q", id.DeviceManagementSettingInstanceId),
	}
	return fmt.Sprintf("Device Management Template Id Category Id Recommended Setting (%s)", strings.Join(components, "\n"))
}
