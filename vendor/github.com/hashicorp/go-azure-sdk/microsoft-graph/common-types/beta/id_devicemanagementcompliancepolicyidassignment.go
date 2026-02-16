package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdAssignmentId{}

// DeviceManagementCompliancePolicyIdAssignmentId is a struct representing the Resource ID for a Device Management Compliance Policy Id Assignment
type DeviceManagementCompliancePolicyIdAssignmentId struct {
	DeviceManagementCompliancePolicyId              string
	DeviceManagementConfigurationPolicyAssignmentId string
}

// NewDeviceManagementCompliancePolicyIdAssignmentID returns a new DeviceManagementCompliancePolicyIdAssignmentId struct
func NewDeviceManagementCompliancePolicyIdAssignmentID(deviceManagementCompliancePolicyId string, deviceManagementConfigurationPolicyAssignmentId string) DeviceManagementCompliancePolicyIdAssignmentId {
	return DeviceManagementCompliancePolicyIdAssignmentId{
		DeviceManagementCompliancePolicyId:              deviceManagementCompliancePolicyId,
		DeviceManagementConfigurationPolicyAssignmentId: deviceManagementConfigurationPolicyAssignmentId,
	}
}

// ParseDeviceManagementCompliancePolicyIdAssignmentID parses 'input' into a DeviceManagementCompliancePolicyIdAssignmentId
func ParseDeviceManagementCompliancePolicyIdAssignmentID(input string) (*DeviceManagementCompliancePolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCompliancePolicyIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementCompliancePolicyIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCompliancePolicyIdAssignmentIDInsensitively(input string) (*DeviceManagementCompliancePolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCompliancePolicyIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCompliancePolicyId, ok = input.Parsed["deviceManagementCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCompliancePolicyId", input)
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId, ok = input.Parsed["deviceManagementConfigurationPolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementCompliancePolicyIdAssignmentID checks that 'input' can be parsed as a Device Management Compliance Policy Id Assignment ID
func ValidateDeviceManagementCompliancePolicyIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCompliancePolicyIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Policy Id Assignment ID
func (id DeviceManagementCompliancePolicyIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/compliancePolicies/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCompliancePolicyId, id.DeviceManagementConfigurationPolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Policy Id Assignment ID
func (id DeviceManagementCompliancePolicyIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("compliancePolicies", "compliancePolicies", "compliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementCompliancePolicyId", "deviceManagementCompliancePolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyAssignmentId", "deviceManagementConfigurationPolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Policy Id Assignment ID
func (id DeviceManagementCompliancePolicyIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Compliance Policy: %q", id.DeviceManagementCompliancePolicyId),
		fmt.Sprintf("Device Management Configuration Policy Assignment: %q", id.DeviceManagementConfigurationPolicyAssignmentId),
	}
	return fmt.Sprintf("Device Management Compliance Policy Id Assignment (%s)", strings.Join(components, "\n"))
}
