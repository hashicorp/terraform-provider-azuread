package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{}

// DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId is a struct representing the Resource ID for a Device Management Managed Device Id Device Compliance Policy State
type DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId struct {
	ManagedDeviceId               string
	DeviceCompliancePolicyStateId string
}

// NewDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID returns a new DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId struct
func NewDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(managedDeviceId string, deviceCompliancePolicyStateId string) DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId {
	return DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{
		ManagedDeviceId:               managedDeviceId,
		DeviceCompliancePolicyStateId: deviceCompliancePolicyStateId,
	}
}

// ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID parses 'input' into a DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId
func ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(input string) (*DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively(input string) (*DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceCompliancePolicyStateId, ok = input.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID checks that 'input' can be parsed as a Device Management Managed Device Id Device Compliance Policy State ID
func ValidateDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Device Compliance Policy State ID
func (id DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/deviceCompliancePolicyStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceCompliancePolicyStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Device Compliance Policy State ID
func (id DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceCompliancePolicyStates", "deviceCompliancePolicyStates", "deviceCompliancePolicyStates"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyStateId", "deviceCompliancePolicyStateId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Device Compliance Policy State ID
func (id DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Compliance Policy State: %q", id.DeviceCompliancePolicyStateId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Device Compliance Policy State (%s)", strings.Join(components, "\n"))
}
