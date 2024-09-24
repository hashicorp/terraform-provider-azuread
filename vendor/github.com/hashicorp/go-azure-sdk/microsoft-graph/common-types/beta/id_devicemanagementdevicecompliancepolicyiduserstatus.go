package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdUserStatusId{}

// DeviceManagementDeviceCompliancePolicyIdUserStatusId is a struct representing the Resource ID for a Device Management Device Compliance Policy Id User Status
type DeviceManagementDeviceCompliancePolicyIdUserStatusId struct {
	DeviceCompliancePolicyId     string
	DeviceComplianceUserStatusId string
}

// NewDeviceManagementDeviceCompliancePolicyIdUserStatusID returns a new DeviceManagementDeviceCompliancePolicyIdUserStatusId struct
func NewDeviceManagementDeviceCompliancePolicyIdUserStatusID(deviceCompliancePolicyId string, deviceComplianceUserStatusId string) DeviceManagementDeviceCompliancePolicyIdUserStatusId {
	return DeviceManagementDeviceCompliancePolicyIdUserStatusId{
		DeviceCompliancePolicyId:     deviceCompliancePolicyId,
		DeviceComplianceUserStatusId: deviceComplianceUserStatusId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyIdUserStatusID parses 'input' into a DeviceManagementDeviceCompliancePolicyIdUserStatusId
func ParseDeviceManagementDeviceCompliancePolicyIdUserStatusID(input string) (*DeviceManagementDeviceCompliancePolicyIdUserStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdUserStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdUserStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIdUserStatusIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyIdUserStatusId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIdUserStatusIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyIdUserStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdUserStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdUserStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyIdUserStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	if id.DeviceComplianceUserStatusId, ok = input.Parsed["deviceComplianceUserStatusId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceUserStatusId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyIdUserStatusID checks that 'input' can be parsed as a Device Management Device Compliance Policy Id User Status ID
func ValidateDeviceManagementDeviceCompliancePolicyIdUserStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyIdUserStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Id User Status ID
func (id DeviceManagementDeviceCompliancePolicyIdUserStatusId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s/userStatuses/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId, id.DeviceComplianceUserStatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Id User Status ID
func (id DeviceManagementDeviceCompliancePolicyIdUserStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
		resourceids.StaticSegment("userStatuses", "userStatuses", "userStatuses"),
		resourceids.UserSpecifiedSegment("deviceComplianceUserStatusId", "deviceComplianceUserStatusId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Id User Status ID
func (id DeviceManagementDeviceCompliancePolicyIdUserStatusId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
		fmt.Sprintf("Device Compliance User Status: %q", id.DeviceComplianceUserStatusId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Id User Status (%s)", strings.Join(components, "\n"))
}
