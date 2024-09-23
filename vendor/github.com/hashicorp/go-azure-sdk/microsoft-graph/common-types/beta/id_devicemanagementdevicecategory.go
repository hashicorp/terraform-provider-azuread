package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCategoryId{}

// DeviceManagementDeviceCategoryId is a struct representing the Resource ID for a Device Management Device Category
type DeviceManagementDeviceCategoryId struct {
	DeviceCategoryId string
}

// NewDeviceManagementDeviceCategoryID returns a new DeviceManagementDeviceCategoryId struct
func NewDeviceManagementDeviceCategoryID(deviceCategoryId string) DeviceManagementDeviceCategoryId {
	return DeviceManagementDeviceCategoryId{
		DeviceCategoryId: deviceCategoryId,
	}
}

// ParseDeviceManagementDeviceCategoryID parses 'input' into a DeviceManagementDeviceCategoryId
func ParseDeviceManagementDeviceCategoryID(input string) (*DeviceManagementDeviceCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCategoryIDInsensitively(input string) (*DeviceManagementDeviceCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCategoryId, ok = input.Parsed["deviceCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCategoryID checks that 'input' can be parsed as a Device Management Device Category ID
func ValidateDeviceManagementDeviceCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Category ID
func (id DeviceManagementDeviceCategoryId) ID() string {
	fmtString := "/deviceManagement/deviceCategories/%s"
	return fmt.Sprintf(fmtString, id.DeviceCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Category ID
func (id DeviceManagementDeviceCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCategories", "deviceCategories", "deviceCategories"),
		resourceids.UserSpecifiedSegment("deviceCategoryId", "deviceCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Device Category ID
func (id DeviceManagementDeviceCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Device Category: %q", id.DeviceCategoryId),
	}
	return fmt.Sprintf("Device Management Device Category (%s)", strings.Join(components, "\n"))
}
