package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdAssignmentId{}

// DeviceManagementConfigurationPolicyIdAssignmentId is a struct representing the Resource ID for a Device Management Configuration Policy Id Assignment
type DeviceManagementConfigurationPolicyIdAssignmentId struct {
	DeviceManagementConfigurationPolicyId           string
	DeviceManagementConfigurationPolicyAssignmentId string
}

// NewDeviceManagementConfigurationPolicyIdAssignmentID returns a new DeviceManagementConfigurationPolicyIdAssignmentId struct
func NewDeviceManagementConfigurationPolicyIdAssignmentID(deviceManagementConfigurationPolicyId string, deviceManagementConfigurationPolicyAssignmentId string) DeviceManagementConfigurationPolicyIdAssignmentId {
	return DeviceManagementConfigurationPolicyIdAssignmentId{
		DeviceManagementConfigurationPolicyId:           deviceManagementConfigurationPolicyId,
		DeviceManagementConfigurationPolicyAssignmentId: deviceManagementConfigurationPolicyAssignmentId,
	}
}

// ParseDeviceManagementConfigurationPolicyIdAssignmentID parses 'input' into a DeviceManagementConfigurationPolicyIdAssignmentId
func ParseDeviceManagementConfigurationPolicyIdAssignmentID(input string) (*DeviceManagementConfigurationPolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationPolicyIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationPolicyIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationPolicyIdAssignmentIDInsensitively(input string) (*DeviceManagementConfigurationPolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationPolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationPolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationPolicyIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId, ok = input.Parsed["deviceManagementConfigurationPolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationPolicyIdAssignmentID checks that 'input' can be parsed as a Device Management Configuration Policy Id Assignment ID
func ValidateDeviceManagementConfigurationPolicyIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationPolicyIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Policy Id Assignment ID
func (id DeviceManagementConfigurationPolicyIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/configurationPolicies/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationPolicyId, id.DeviceManagementConfigurationPolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Policy Id Assignment ID
func (id DeviceManagementConfigurationPolicyIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationPolicies", "configurationPolicies", "configurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyAssignmentId", "deviceManagementConfigurationPolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Policy Id Assignment ID
func (id DeviceManagementConfigurationPolicyIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
		fmt.Sprintf("Device Management Configuration Policy Assignment: %q", id.DeviceManagementConfigurationPolicyAssignmentId),
	}
	return fmt.Sprintf("Device Management Configuration Policy Id Assignment (%s)", strings.Join(components, "\n"))
}
