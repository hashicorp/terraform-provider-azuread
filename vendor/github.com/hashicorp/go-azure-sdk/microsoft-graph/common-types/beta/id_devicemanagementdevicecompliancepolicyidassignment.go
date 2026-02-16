package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdAssignmentId{}

// DeviceManagementDeviceCompliancePolicyIdAssignmentId is a struct representing the Resource ID for a Device Management Device Compliance Policy Id Assignment
type DeviceManagementDeviceCompliancePolicyIdAssignmentId struct {
	DeviceCompliancePolicyId           string
	DeviceCompliancePolicyAssignmentId string
}

// NewDeviceManagementDeviceCompliancePolicyIdAssignmentID returns a new DeviceManagementDeviceCompliancePolicyIdAssignmentId struct
func NewDeviceManagementDeviceCompliancePolicyIdAssignmentID(deviceCompliancePolicyId string, deviceCompliancePolicyAssignmentId string) DeviceManagementDeviceCompliancePolicyIdAssignmentId {
	return DeviceManagementDeviceCompliancePolicyIdAssignmentId{
		DeviceCompliancePolicyId:           deviceCompliancePolicyId,
		DeviceCompliancePolicyAssignmentId: deviceCompliancePolicyAssignmentId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyIdAssignmentID parses 'input' into a DeviceManagementDeviceCompliancePolicyIdAssignmentId
func ParseDeviceManagementDeviceCompliancePolicyIdAssignmentID(input string) (*DeviceManagementDeviceCompliancePolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIdAssignmentIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	if id.DeviceCompliancePolicyAssignmentId, ok = input.Parsed["deviceCompliancePolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyIdAssignmentID checks that 'input' can be parsed as a Device Management Device Compliance Policy Id Assignment ID
func ValidateDeviceManagementDeviceCompliancePolicyIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Id Assignment ID
func (id DeviceManagementDeviceCompliancePolicyIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId, id.DeviceCompliancePolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Id Assignment ID
func (id DeviceManagementDeviceCompliancePolicyIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyAssignmentId", "deviceCompliancePolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Id Assignment ID
func (id DeviceManagementDeviceCompliancePolicyIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
		fmt.Sprintf("Device Compliance Policy Assignment: %q", id.DeviceCompliancePolicyAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Id Assignment (%s)", strings.Join(components, "\n"))
}
