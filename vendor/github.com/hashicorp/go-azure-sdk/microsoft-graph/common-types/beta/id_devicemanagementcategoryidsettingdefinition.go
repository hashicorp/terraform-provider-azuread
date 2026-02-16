package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCategoryIdSettingDefinitionId{}

// DeviceManagementCategoryIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Category Id Setting Definition
type DeviceManagementCategoryIdSettingDefinitionId struct {
	DeviceManagementSettingCategoryId   string
	DeviceManagementSettingDefinitionId string
}

// NewDeviceManagementCategoryIdSettingDefinitionID returns a new DeviceManagementCategoryIdSettingDefinitionId struct
func NewDeviceManagementCategoryIdSettingDefinitionID(deviceManagementSettingCategoryId string, deviceManagementSettingDefinitionId string) DeviceManagementCategoryIdSettingDefinitionId {
	return DeviceManagementCategoryIdSettingDefinitionId{
		DeviceManagementSettingCategoryId:   deviceManagementSettingCategoryId,
		DeviceManagementSettingDefinitionId: deviceManagementSettingDefinitionId,
	}
}

// ParseDeviceManagementCategoryIdSettingDefinitionID parses 'input' into a DeviceManagementCategoryIdSettingDefinitionId
func ParseDeviceManagementCategoryIdSettingDefinitionID(input string) (*DeviceManagementCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCategoryIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementCategoryIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCategoryIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCategoryIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementSettingCategoryId, ok = input.Parsed["deviceManagementSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingCategoryId", input)
	}

	if id.DeviceManagementSettingDefinitionId, ok = input.Parsed["deviceManagementSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementCategoryIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Category Id Setting Definition ID
func ValidateDeviceManagementCategoryIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCategoryIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Category Id Setting Definition ID
func (id DeviceManagementCategoryIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/categories/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementSettingCategoryId, id.DeviceManagementSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Category Id Setting Definition ID
func (id DeviceManagementCategoryIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingCategoryId", "deviceManagementSettingCategoryId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingDefinitionId", "deviceManagementSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Category Id Setting Definition ID
func (id DeviceManagementCategoryIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Setting Category: %q", id.DeviceManagementSettingCategoryId),
		fmt.Sprintf("Device Management Setting Definition: %q", id.DeviceManagementSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Category Id Setting Definition (%s)", strings.Join(components, "\n"))
}
