package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{}

// DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId is a struct representing the Resource ID for a Device Management Device Configurations All Managed Device Certificate State
type DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId struct {
	ManagedAllDeviceCertificateStateId string
}

// NewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID returns a new DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId struct
func NewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(managedAllDeviceCertificateStateId string) DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId {
	return DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{
		ManagedAllDeviceCertificateStateId: managedAllDeviceCertificateStateId,
	}
}

// ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID parses 'input' into a DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId
func ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(input string) (*DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateIDInsensitively(input string) (*DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedAllDeviceCertificateStateId, ok = input.Parsed["managedAllDeviceCertificateStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedAllDeviceCertificateStateId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID checks that 'input' can be parsed as a Device Management Device Configurations All Managed Device Certificate State ID
func ValidateDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configurations All Managed Device Certificate State ID
func (id DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedAllDeviceCertificateStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configurations All Managed Device Certificate State ID
func (id DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurationsAllManagedDeviceCertificateStates", "deviceConfigurationsAllManagedDeviceCertificateStates", "deviceConfigurationsAllManagedDeviceCertificateStates"),
		resourceids.UserSpecifiedSegment("managedAllDeviceCertificateStateId", "managedAllDeviceCertificateStateId"),
	}
}

// String returns a human-readable description of this Device Management Device Configurations All Managed Device Certificate State ID
func (id DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed All Device Certificate State: %q", id.ManagedAllDeviceCertificateStateId),
	}
	return fmt.Sprintf("Device Management Device Configurations All Managed Device Certificate State (%s)", strings.Join(components, "\n"))
}
