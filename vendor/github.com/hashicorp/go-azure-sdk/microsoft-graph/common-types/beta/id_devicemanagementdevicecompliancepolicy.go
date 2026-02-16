package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyId{}

// DeviceManagementDeviceCompliancePolicyId is a struct representing the Resource ID for a Device Management Device Compliance Policy
type DeviceManagementDeviceCompliancePolicyId struct {
	DeviceCompliancePolicyId string
}

// NewDeviceManagementDeviceCompliancePolicyID returns a new DeviceManagementDeviceCompliancePolicyId struct
func NewDeviceManagementDeviceCompliancePolicyID(deviceCompliancePolicyId string) DeviceManagementDeviceCompliancePolicyId {
	return DeviceManagementDeviceCompliancePolicyId{
		DeviceCompliancePolicyId: deviceCompliancePolicyId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyID parses 'input' into a DeviceManagementDeviceCompliancePolicyId
func ParseDeviceManagementDeviceCompliancePolicyID(input string) (*DeviceManagementDeviceCompliancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyID checks that 'input' can be parsed as a Device Management Device Compliance Policy ID
func ValidateDeviceManagementDeviceCompliancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy ID
func (id DeviceManagementDeviceCompliancePolicyId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy ID
func (id DeviceManagementDeviceCompliancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy ID
func (id DeviceManagementDeviceCompliancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy (%s)", strings.Join(components, "\n"))
}
