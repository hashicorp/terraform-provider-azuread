package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdSettingId{}

// DeviceManagementConfigurationPolicyIdSettingId is a struct representing the Resource ID for a Device Management Configuration Policy Id Setting
type DeviceManagementConfigurationPolicyIdSettingId struct {
	DeviceManagementConfigurationPolicyId  string
	DeviceManagementConfigurationSettingId string
}

// NewDeviceManagementConfigurationPolicyIdSettingID returns a new DeviceManagementConfigurationPolicyIdSettingId struct
func NewDeviceManagementConfigurationPolicyIdSettingID(deviceManagementConfigurationPolicyId string, deviceManagementConfigurationSettingId string) DeviceManagementConfigurationPolicyIdSettingId {
	return DeviceManagementConfigurationPolicyIdSettingId{
		DeviceManagementConfigurationPolicyId:  deviceManagementConfigurationPolicyId,
		DeviceManagementConfigurationSettingId: deviceManagementConfigurationSettingId,
	}
}

// ParseDeviceManagementConfigurationPolicyIdSettingID parses 'input' into a DeviceManagementConfigurationPolicyIdSettingId
func ParseDeviceManagementConfigurationPolicyIdSettingID(input string) (*DeviceManagementConfigurationPolicyIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyIdSettingIDInsensitively(input string) (*DeviceManagementConfigurationPolicyIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	if id.DeviceManagementConfigurationSettingId, ok = input.Parsed["deviceManagementConfigurationSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyIdSettingID checks that 'input' can be parsed as a Device Management Configuration Policy Id Setting ID
func ValidateDeviceManagementConfigurationPolicyIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy Id Setting ID
func (id DeviceManagementConfigurationPolicyIdSettingId) ID() string {
	fmtString := "/deviceManagement/configurationPolicies/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyId, id.DeviceManagementConfigurationSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy Id Setting ID
func (id DeviceManagementConfigurationPolicyIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicies", "configurationPolicies", "configurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingId", "deviceManagementConfigurationSettingId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy Id Setting ID
func (id DeviceManagementConfigurationPolicyIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
		fmt.Sprintf("Device Management Configuration Setting: %q", id.DeviceManagementConfigurationSettingId),
	}
	return fmt.Sprintf("Device Management Configuration Policy Id Setting (%s)", strings.Join(components, "\n"))
}
