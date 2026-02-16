package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementSettingDefinitionId{}

// DeviceManagementSettingDefinitionId is a struct representing the Resource ID for a Device Management Setting Definition
type DeviceManagementSettingDefinitionId struct {
	DeviceManagementSettingDefinitionId string
}

// NewDeviceManagementSettingDefinitionID returns a new DeviceManagementSettingDefinitionId struct
func NewDeviceManagementSettingDefinitionID(deviceManagementSettingDefinitionId string) DeviceManagementSettingDefinitionId {
	return DeviceManagementSettingDefinitionId{
		DeviceManagementSettingDefinitionId: deviceManagementSettingDefinitionId,
	}
}

// ParseDeviceManagementSettingDefinitionID parses 'input' into a DeviceManagementSettingDefinitionId
func ParseDeviceManagementSettingDefinitionID(input string) (*DeviceManagementSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementSettingDefinitionIDInsensitively(input string) (*DeviceManagementSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementSettingDefinitionId, ok = input.Parsed["deviceManagementSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementSettingDefinitionID checks that 'input' can be parsed as a Device Management Setting Definition ID
func ValidateDeviceManagementSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Setting Definition ID
func (id DeviceManagementSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Setting Definition ID
func (id DeviceManagementSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementSettingDefinitionId", "deviceManagementSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Setting Definition ID
func (id DeviceManagementSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Setting Definition: %q", id.DeviceManagementSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Setting Definition (%s)", strings.Join(components, "\n"))
}
