package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusableSettingId{}

// DeviceManagementReusableSettingId is a struct representing the Resource ID for a Device Management Reusable Setting
type DeviceManagementReusableSettingId struct {
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementReusableSettingID returns a new DeviceManagementReusableSettingId struct
func NewDeviceManagementReusableSettingID(deviceManagementConfigurationSettingDefinitionId string) DeviceManagementReusableSettingId {
	return DeviceManagementReusableSettingId{
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementReusableSettingID parses 'input' into a DeviceManagementReusableSettingId
func ParseDeviceManagementReusableSettingID(input string) (*DeviceManagementReusableSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusableSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusableSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReusableSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementReusableSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReusableSettingIDInsensitively(input string) (*DeviceManagementReusableSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusableSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusableSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReusableSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementReusableSettingID checks that 'input' can be parsed as a Device Management Reusable Setting ID
func ValidateDeviceManagementReusableSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReusableSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Reusable Setting ID
func (id DeviceManagementReusableSettingId) ID() string {
	fmtString := "/deviceManagement/reusableSettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Reusable Setting ID
func (id DeviceManagementReusableSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reusableSettings", "reusableSettings", "reusableSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Reusable Setting ID
func (id DeviceManagementReusableSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Reusable Setting (%s)", strings.Join(components, "\n"))
}
