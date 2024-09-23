package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdSecurityBaselineStateId{}

// DeviceManagementManagedDeviceIdSecurityBaselineStateId is a struct representing the Resource ID for a Device Management Managed Device Id Security Baseline State
type DeviceManagementManagedDeviceIdSecurityBaselineStateId struct {
	ManagedDeviceId         string
	SecurityBaselineStateId string
}

// NewDeviceManagementManagedDeviceIdSecurityBaselineStateID returns a new DeviceManagementManagedDeviceIdSecurityBaselineStateId struct
func NewDeviceManagementManagedDeviceIdSecurityBaselineStateID(managedDeviceId string, securityBaselineStateId string) DeviceManagementManagedDeviceIdSecurityBaselineStateId {
	return DeviceManagementManagedDeviceIdSecurityBaselineStateId{
		ManagedDeviceId:         managedDeviceId,
		SecurityBaselineStateId: securityBaselineStateId,
	}
}

// ParseDeviceManagementManagedDeviceIdSecurityBaselineStateID parses 'input' into a DeviceManagementManagedDeviceIdSecurityBaselineStateId
func ParseDeviceManagementManagedDeviceIdSecurityBaselineStateID(input string) (*DeviceManagementManagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdSecurityBaselineStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIDInsensitively(input string) (*DeviceManagementManagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdSecurityBaselineStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.SecurityBaselineStateId, ok = input.Parsed["securityBaselineStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdSecurityBaselineStateID checks that 'input' can be parsed as a Device Management Managed Device Id Security Baseline State ID
func ValidateDeviceManagementManagedDeviceIdSecurityBaselineStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdSecurityBaselineStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Security Baseline State ID
func (id DeviceManagementManagedDeviceIdSecurityBaselineStateId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/securityBaselineStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Security Baseline State ID
func (id DeviceManagementManagedDeviceIdSecurityBaselineStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Security Baseline State ID
func (id DeviceManagementManagedDeviceIdSecurityBaselineStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Security Baseline State (%s)", strings.Join(components, "\n"))
}
