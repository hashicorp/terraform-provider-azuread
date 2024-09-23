package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId{}

// DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId is a struct representing the Resource ID for a Device Management Comanaged Device Id Device Compliance Policy State
type DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId struct {
	ManagedDeviceId               string
	DeviceCompliancePolicyStateId string
}

// NewDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID returns a new DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId struct
func NewDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID(managedDeviceId string, deviceCompliancePolicyStateId string) DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId {
	return DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId{
		ManagedDeviceId:               managedDeviceId,
		DeviceCompliancePolicyStateId: deviceCompliancePolicyStateId,
	}
}

// ParseDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID parses 'input' into a DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId
func ParseDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID(input string) (*DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceCompliancePolicyStateId, ok = input.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID checks that 'input' can be parsed as a Device Management Comanaged Device Id Device Compliance Policy State ID
func ValidateDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Device Compliance Policy State ID
func (id DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/deviceCompliancePolicyStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceCompliancePolicyStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Device Compliance Policy State ID
func (id DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceCompliancePolicyStates", "deviceCompliancePolicyStates", "deviceCompliancePolicyStates"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyStateId", "deviceCompliancePolicyStateId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Device Compliance Policy State ID
func (id DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Compliance Policy State: %q", id.DeviceCompliancePolicyStateId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Device Compliance Policy State (%s)", strings.Join(components, "\n"))
}
