package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryIdSettingId{}

// DeviceManagementIntentIdCategoryIdSettingId is a struct representing the Resource ID for a Device Management Intent Id Category Id Setting
type DeviceManagementIntentIdCategoryIdSettingId struct {
	DeviceManagementIntentId                string
	DeviceManagementIntentSettingCategoryId string
	DeviceManagementSettingInstanceId       string
}

// NewDeviceManagementIntentIdCategoryIdSettingID returns a new DeviceManagementIntentIdCategoryIdSettingId struct
func NewDeviceManagementIntentIdCategoryIdSettingID(deviceManagementIntentId string, deviceManagementIntentSettingCategoryId string, deviceManagementSettingInstanceId string) DeviceManagementIntentIdCategoryIdSettingId {
	return DeviceManagementIntentIdCategoryIdSettingId{
		DeviceManagementIntentId:                deviceManagementIntentId,
		DeviceManagementIntentSettingCategoryId: deviceManagementIntentSettingCategoryId,
		DeviceManagementSettingInstanceId:       deviceManagementSettingInstanceId,
	}
}

// ParseDeviceManagementIntentIdCategoryIdSettingID parses 'input' into a DeviceManagementIntentIdCategoryIdSettingId
func ParseDeviceManagementIntentIdCategoryIdSettingID(input string) (*DeviceManagementIntentIdCategoryIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdCategoryIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdCategoryIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdCategoryIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdCategoryIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdCategoryIdSettingIDInsensitively(input string) (*DeviceManagementIntentIdCategoryIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdCategoryIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdCategoryIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdCategoryIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentSettingCategoryId, ok = input.Parsed["deviceManagementIntentSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentSettingCategoryId", input)
	}

	if id.DeviceManagementSettingInstanceId, ok = input.Parsed["deviceManagementSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingInstanceId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdCategoryIdSettingID checks that 'input' can be parsed as a Device Management Intent Id Category Id Setting ID
func ValidateDeviceManagementIntentIdCategoryIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdCategoryIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Category Id Setting ID
func (id DeviceManagementIntentIdCategoryIdSettingId) ID() string {
	fmtString := "/deviceManagement/intents/%s/categories/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentSettingCategoryId, id.DeviceManagementSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Category Id Setting ID
func (id DeviceManagementIntentIdCategoryIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentSettingCategoryId", "deviceManagementIntentSettingCategoryId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingInstanceId", "deviceManagementSettingInstanceId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Category Id Setting ID
func (id DeviceManagementIntentIdCategoryIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent Setting Category: %q", id.DeviceManagementIntentSettingCategoryId),
		fmt.Sprintf("Device Management Setting Instance: %q", id.DeviceManagementSettingInstanceId),
	}
	return fmt.Sprintf("Device Management Intent Id Category Id Setting (%s)", strings.Join(components, "\n"))
}
