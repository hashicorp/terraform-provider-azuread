package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryIdSettingDefinitionId{}

// DeviceManagementIntentIdCategoryIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Intent Id Category Id Setting Definition
type DeviceManagementIntentIdCategoryIdSettingDefinitionId struct {
	DeviceManagementIntentId                string
	DeviceManagementIntentSettingCategoryId string
	DeviceManagementSettingDefinitionId     string
}

// NewDeviceManagementIntentIdCategoryIdSettingDefinitionID returns a new DeviceManagementIntentIdCategoryIdSettingDefinitionId struct
func NewDeviceManagementIntentIdCategoryIdSettingDefinitionID(deviceManagementIntentId string, deviceManagementIntentSettingCategoryId string, deviceManagementSettingDefinitionId string) DeviceManagementIntentIdCategoryIdSettingDefinitionId {
	return DeviceManagementIntentIdCategoryIdSettingDefinitionId{
		DeviceManagementIntentId:                deviceManagementIntentId,
		DeviceManagementIntentSettingCategoryId: deviceManagementIntentSettingCategoryId,
		DeviceManagementSettingDefinitionId:     deviceManagementSettingDefinitionId,
	}
}

// ParseDeviceManagementIntentIdCategoryIdSettingDefinitionID parses 'input' into a DeviceManagementIntentIdCategoryIdSettingDefinitionId
func ParseDeviceManagementIntentIdCategoryIdSettingDefinitionID(input string) (*DeviceManagementIntentIdCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementIntentIdCategoryIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementIntentIdCategoryIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementIntentIdCategoryIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementIntentIdCategoryIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementIntentIdCategoryIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementIntentIdCategoryIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementIntentIdCategoryIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementIntentId, ok = input.Parsed["deviceManagementIntentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentId", input)
	}

	if id.DeviceManagementIntentSettingCategoryId, ok = input.Parsed["deviceManagementIntentSettingCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementIntentSettingCategoryId", input)
	}

	if id.DeviceManagementSettingDefinitionId, ok = input.Parsed["deviceManagementSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementIntentIdCategoryIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Intent Id Category Id Setting Definition ID
func ValidateDeviceManagementIntentIdCategoryIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementIntentIdCategoryIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Intent Id Category Id Setting Definition ID
func (id DeviceManagementIntentIdCategoryIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/intents/%s/categories/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementIntentId, id.DeviceManagementIntentSettingCategoryId, id.DeviceManagementSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Intent Id Category Id Setting Definition ID
func (id DeviceManagementIntentIdCategoryIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("intents", "intents", "intents"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentId", "deviceManagementIntentId"),
		resourceids.StaticSegment("categories", "categories", "categories"),
		resourceids.UserSpecifiedSegment("deviceManagementIntentSettingCategoryId", "deviceManagementIntentSettingCategoryId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingDefinitionId", "deviceManagementSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Intent Id Category Id Setting Definition ID
func (id DeviceManagementIntentIdCategoryIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Intent: %q", id.DeviceManagementIntentId),
		fmt.Sprintf("Device Management Intent Setting Category: %q", id.DeviceManagementIntentSettingCategoryId),
		fmt.Sprintf("Device Management Setting Definition: %q", id.DeviceManagementSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Intent Id Category Id Setting Definition (%s)", strings.Join(components, "\n"))
}
