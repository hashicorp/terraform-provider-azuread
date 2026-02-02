package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdCategoryId{}

// DeviceManagementTemplateIdCategoryId is a struct representing the Resource ID for a Device Management Template Id Category
type DeviceManagementTemplateIdCategoryId struct {
	DeviceManagementTemplateId                string
	DeviceManagementTemplateSettingCategoryId string
}

// NewDeviceManagementTemplateIdCategoryID returns a new DeviceManagementTemplateIdCategoryId struct
func NewDeviceManagementTemplateIdCategoryID(deviceManagementTemplateId string, deviceManagementTemplateSettingCategoryId string) DeviceManagementTemplateIdCategoryId {
	return DeviceManagementTemplateIdCategoryId{
		DeviceManagementTemplateId:                deviceManagementTemplateId,
		DeviceManagementTemplateSettingCategoryId: deviceManagementTemplateSettingCategoryId,
	}
}

// ParseDeviceManagementTemplateIdCategoryID parses 'input' into a DeviceManagementTemplateIdCategoryId
func ParseDeviceManagementTemplateIdCategoryID(input string) (*DeviceManagementTemplateIdCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateIdCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateIdCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateIdCategoryIDInsensitively(input string) (*DeviceManagementTemplateIdCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateIdCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateIdCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateIdCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateId, ok = input.Parsed["deviceManagementTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateId", input)
	}

	if id.DeviceManagementTemplateSettingCategoryId, ok = input.Parsed["deviceManagementTemplateSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateSettingCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateIdCategoryID checks that 'input' can be parsed as a Device Management Template Id Category ID
func ValidateDeviceManagementTemplateIdCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateIdCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Id Category ID
func (id DeviceManagementTemplateIdCategoryId) ID() string {
	fmtString := "/deviceManagement/templates/%s/categories/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateId, id.DeviceManagementTemplateSettingCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Id Category ID
func (id DeviceManagementTemplateIdCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateId", "deviceManagementTemplateId"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateSettingCategoryId", "deviceManagementTemplateSettingCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Template Id Category ID
func (id DeviceManagementTemplateIdCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template: %q", id.DeviceManagementTemplateId),
		fmt.Sprintf("Device Management Template Setting Category: %q", id.DeviceManagementTemplateSettingCategoryId),
	}
	return fmt.Sprintf("Device Management Template Id Category (%s)", strings.Join(components, "\n"))
}
