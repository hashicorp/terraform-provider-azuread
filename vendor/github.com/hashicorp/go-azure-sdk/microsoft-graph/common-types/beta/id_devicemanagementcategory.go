package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCategoryId{}

// DeviceManagementCategoryId is a struct representing the Resource ID for a Device Management Category
type DeviceManagementCategoryId struct {
	DeviceManagementSettingCategoryId string
}

// NewDeviceManagementCategoryID returns a new DeviceManagementCategoryId struct
func NewDeviceManagementCategoryID(deviceManagementSettingCategoryId string) DeviceManagementCategoryId {
	return DeviceManagementCategoryId{
		DeviceManagementSettingCategoryId: deviceManagementSettingCategoryId,
	}
}

// ParseDeviceManagementCategoryID parses 'input' into a DeviceManagementCategoryId
func ParseDeviceManagementCategoryID(input string) (*DeviceManagementCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCategoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementCategoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCategoryIDInsensitively(input string) (*DeviceManagementCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCategoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCategoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementSettingCategoryId, ok = input.Parsed["deviceManagementSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingCategoryId", input)
	}

	return nil
}

// ValidateDeviceManagementCategoryID checks that 'input' can be parsed as a Device Management Category ID
func ValidateDeviceManagementCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Category ID
func (id DeviceManagementCategoryId) ID() string {
	fmtString := "/deviceManagement/categories/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementSettingCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Category ID
func (id DeviceManagementCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingCategoryId", "deviceManagementSettingCategoryId"),
	}
}

// String returns a human-readable description of this Device Management Category ID
func (id DeviceManagementCategoryId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Setting Category: %q", id.DeviceManagementSettingCategoryId),
	}
	return fmt.Sprintf("Device Management Category (%s)", strings.Join(components, "\n"))
}
