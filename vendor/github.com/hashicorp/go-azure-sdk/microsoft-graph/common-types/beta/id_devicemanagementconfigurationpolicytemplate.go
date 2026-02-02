package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateId{}

// DeviceManagementConfigurationPolicyTemplateId is a struct representing the Resource ID for a Device Management Configuration Policy Template
type DeviceManagementConfigurationPolicyTemplateId struct {
	DeviceManagementConfigurationPolicyTemplateId string
}

// NewDeviceManagementConfigurationPolicyTemplateID returns a new DeviceManagementConfigurationPolicyTemplateId struct
func NewDeviceManagementConfigurationPolicyTemplateID(deviceManagementConfigurationPolicyTemplateId string) DeviceManagementConfigurationPolicyTemplateId {
	return DeviceManagementConfigurationPolicyTemplateId{
		DeviceManagementConfigurationPolicyTemplateId: deviceManagementConfigurationPolicyTemplateId,
	}
}

// ParseDeviceManagementConfigurationPolicyTemplateID parses 'input' into a DeviceManagementConfigurationPolicyTemplateId
func ParseDeviceManagementConfigurationPolicyTemplateID(input string) (*DeviceManagementConfigurationPolicyTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyTemplateIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyTemplateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyTemplateIDInsensitively(input string) (*DeviceManagementConfigurationPolicyTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyTemplateId, ok = input.Parsed["deviceManagementConfigurationPolicyTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyTemplateId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyTemplateID checks that 'input' can be parsed as a Device Management Configuration Policy Template ID
func ValidateDeviceManagementConfigurationPolicyTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy Template ID
func (id DeviceManagementConfigurationPolicyTemplateId) ID() string {
	fmtString := "/deviceManagement/configurationPolicyTemplates/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy Template ID
func (id DeviceManagementConfigurationPolicyTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicyTemplates", "configurationPolicyTemplates", "configurationPolicyTemplates"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationPolicyTemplateId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy Template ID
func (id DeviceManagementConfigurationPolicyTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy Template: %q", id.DeviceManagementConfigurationPolicyTemplateId),
	}
	return fmt.Sprintf("Device Management Configuration Policy Template (%s)", strings.Join(components, "\n"))
}
