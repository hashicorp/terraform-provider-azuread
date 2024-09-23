package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryId{}

// DeviceManagementIntentIdCategoryId is a struct representing the Resource ID for a Device Management Intent Id Category
type DeviceManagementIntentIdCategoryId struct {
	DeviceManagementIntentId                string
	DeviceManagementIntentSettingCategoryId string
}

// NewDeviceManagementIntentIdCategoryID returns a new DeviceManagementIntentIdCategoryId struct
func NewDeviceManagementIntentIdCategoryID(deviceManagementIntentId string, deviceManagementIntentSettingCategoryId string) DeviceManagementIntentIdCategoryId {
	return DeviceManagementIntentIdCategoryId{
		DeviceManagementIntentId:                deviceManagementIntentId,
		DeviceManagementIntentSettingCategoryId: deviceManagementIntentSettingCategoryId,
	}
}

// ParseDeviceManagementIntentIdCategoryID parses 'input' into a DeviceManagementIntentIdCategoryId
func ParseDeviceManagementIntentIdCategoryID(input string) (*DeviceManagementIntentIdCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdCategoryIDInsensitively(input string) (*DeviceManagementIntentIdCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentSettingCategoryId, ok = input.Parsed["deviceManagementIntentSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentSettingCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdCategoryID checks that 'input' can be parsed as a Device Management Intent Id Category ID
func ValidateDeviceManagementIntentIdCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Category ID
func (id DeviceManagementIntentIdCategoryId) ID() string {
	fmtString := "/deviceManagement/intents/%s/categories/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentSettingCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Category ID
func (id DeviceManagementIntentIdCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentSettingCategoryId", "deviceManagementIntentSettingCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Category ID
func (id DeviceManagementIntentIdCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent Setting Category: %q", id.DeviceManagementIntentSettingCategoryId),
	}
	return fmt.Sprintf("Device Management Intent Id Category (%s)", strings.Join(components, "\n"))
}
