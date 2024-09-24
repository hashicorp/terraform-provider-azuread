package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyId{}

// DeviceManagementConfigurationPolicyId is a struct representing the Resource ID for a Device Management Configuration Policy
type DeviceManagementConfigurationPolicyId struct {
	DeviceManagementConfigurationPolicyId string
}

// NewDeviceManagementConfigurationPolicyID returns a new DeviceManagementConfigurationPolicyId struct
func NewDeviceManagementConfigurationPolicyID(deviceManagementConfigurationPolicyId string) DeviceManagementConfigurationPolicyId {
	return DeviceManagementConfigurationPolicyId{
		DeviceManagementConfigurationPolicyId: deviceManagementConfigurationPolicyId,
	}
}

// ParseDeviceManagementConfigurationPolicyID parses 'input' into a DeviceManagementConfigurationPolicyId
func ParseDeviceManagementConfigurationPolicyID(input string) (*DeviceManagementConfigurationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyIDInsensitively(input string) (*DeviceManagementConfigurationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyID checks that 'input' can be parsed as a Device Management Configuration Policy ID
func ValidateDeviceManagementConfigurationPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy ID
func (id DeviceManagementConfigurationPolicyId) ID() string {
	fmtString := "/deviceManagement/configurationPolicies/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy ID
func (id DeviceManagementConfigurationPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicies", "configurationPolicies", "configurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy ID
func (id DeviceManagementConfigurationPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
	}
	return fmt.Sprintf("Device Management Configuration Policy (%s)", strings.Join(components, "\n"))
}
