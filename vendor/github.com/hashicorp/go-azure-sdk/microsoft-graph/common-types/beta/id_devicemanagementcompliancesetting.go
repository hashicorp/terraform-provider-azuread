package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComplianceSettingId{}

// DeviceManagementComplianceSettingId is a struct representing the Resource ID for a Device Management Compliance Setting
type DeviceManagementComplianceSettingId struct {
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementComplianceSettingID returns a new DeviceManagementComplianceSettingId struct
func NewDeviceManagementComplianceSettingID(deviceManagementConfigurationSettingDefinitionId string) DeviceManagementComplianceSettingId {
	return DeviceManagementComplianceSettingId{
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementComplianceSettingID parses 'input' into a DeviceManagementComplianceSettingId
func ParseDeviceManagementComplianceSettingID(input string) (*DeviceManagementComplianceSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComplianceSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComplianceSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComplianceSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementComplianceSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComplianceSettingIDInsensitively(input string) (*DeviceManagementComplianceSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComplianceSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComplianceSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComplianceSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementComplianceSettingID checks that 'input' can be parsed as a Device Management Compliance Setting ID
func ValidateDeviceManagementComplianceSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComplianceSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Setting ID
func (id DeviceManagementComplianceSettingId) ID() string {
	fmtString := "/deviceManagement/complianceSettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Setting ID
func (id DeviceManagementComplianceSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("complianceSettings", "complianceSettings", "complianceSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Setting ID
func (id DeviceManagementComplianceSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Compliance Setting (%s)", strings.Join(components, "\n"))
}
