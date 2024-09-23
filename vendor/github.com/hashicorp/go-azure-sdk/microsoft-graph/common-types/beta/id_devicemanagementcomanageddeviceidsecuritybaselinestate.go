package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdSecurityBaselineStateId{}

// DeviceManagementComanagedDeviceIdSecurityBaselineStateId is a struct representing the Resource ID for a Device Management Comanaged Device Id Security Baseline State
type DeviceManagementComanagedDeviceIdSecurityBaselineStateId struct {
	ManagedDeviceId         string
	SecurityBaselineStateId string
}

// NewDeviceManagementComanagedDeviceIdSecurityBaselineStateID returns a new DeviceManagementComanagedDeviceIdSecurityBaselineStateId struct
func NewDeviceManagementComanagedDeviceIdSecurityBaselineStateID(managedDeviceId string, securityBaselineStateId string) DeviceManagementComanagedDeviceIdSecurityBaselineStateId {
	return DeviceManagementComanagedDeviceIdSecurityBaselineStateId{
		ManagedDeviceId:         managedDeviceId,
		SecurityBaselineStateId: securityBaselineStateId,
	}
}

// ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateID parses 'input' into a DeviceManagementComanagedDeviceIdSecurityBaselineStateId
func ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateID(input string) (*DeviceManagementComanagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdSecurityBaselineStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdSecurityBaselineStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.SecurityBaselineStateId, ok = input.Parsed["securityBaselineStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", input)
	}

	return nil
}

// ValidateDeviceManagementComanagedDeviceIdSecurityBaselineStateID checks that 'input' can be parsed as a Device Management Comanaged Device Id Security Baseline State ID
func ValidateDeviceManagementComanagedDeviceIdSecurityBaselineStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Security Baseline State ID
func (id DeviceManagementComanagedDeviceIdSecurityBaselineStateId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/securityBaselineStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Security Baseline State ID
func (id DeviceManagementComanagedDeviceIdSecurityBaselineStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Security Baseline State ID
func (id DeviceManagementComanagedDeviceIdSecurityBaselineStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Security Baseline State (%s)", strings.Join(components, "\n"))
}
