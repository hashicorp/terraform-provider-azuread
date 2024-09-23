package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{}

// DeviceManagementDeviceCompliancePolicyIdDeviceStatusId is a struct representing the Resource ID for a Device Management Device Compliance Policy Id Device Status
type DeviceManagementDeviceCompliancePolicyIdDeviceStatusId struct {
	DeviceCompliancePolicyId       string
	DeviceComplianceDeviceStatusId string
}

// NewDeviceManagementDeviceCompliancePolicyIdDeviceStatusID returns a new DeviceManagementDeviceCompliancePolicyIdDeviceStatusId struct
func NewDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(deviceCompliancePolicyId string, deviceComplianceDeviceStatusId string) DeviceManagementDeviceCompliancePolicyIdDeviceStatusId {
	return DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{
		DeviceCompliancePolicyId:       deviceCompliancePolicyId,
		DeviceComplianceDeviceStatusId: deviceComplianceDeviceStatusId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusID parses 'input' into a DeviceManagementDeviceCompliancePolicyIdDeviceStatusId
func ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(input string) (*DeviceManagementDeviceCompliancePolicyIdDeviceStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyIdDeviceStatusId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyIdDeviceStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyIdDeviceStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	if id.DeviceComplianceDeviceStatusId, ok = input.Parsed["deviceComplianceDeviceStatusId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceDeviceStatusId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyIdDeviceStatusID checks that 'input' can be parsed as a Device Management Device Compliance Policy Id Device Status ID
func ValidateDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Id Device Status ID
func (id DeviceManagementDeviceCompliancePolicyIdDeviceStatusId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s/deviceStatuses/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId, id.DeviceComplianceDeviceStatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Id Device Status ID
func (id DeviceManagementDeviceCompliancePolicyIdDeviceStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
		resourceids.StaticSegment("deviceStatuses", "deviceStatuses", "deviceStatuses"),
		resourceids.UserSpecifiedSegment("deviceComplianceDeviceStatusId", "deviceComplianceDeviceStatusId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Id Device Status ID
func (id DeviceManagementDeviceCompliancePolicyIdDeviceStatusId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
		fmt.Sprintf("Device Compliance Device Status: %q", id.DeviceComplianceDeviceStatusId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Id Device Status (%s)", strings.Join(components, "\n"))
}
