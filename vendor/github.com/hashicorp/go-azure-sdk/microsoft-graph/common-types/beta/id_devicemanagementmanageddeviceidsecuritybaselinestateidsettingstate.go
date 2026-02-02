package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId{}

// DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId is a struct representing the Resource ID for a Device Management Managed Device Id Security Baseline State Id Setting State
type DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId struct {
	ManagedDeviceId                string
	SecurityBaselineStateId        string
	SecurityBaselineSettingStateId string
}

// NewDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID returns a new DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId struct
func NewDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID(managedDeviceId string, securityBaselineStateId string, securityBaselineSettingStateId string) DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId {
	return DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId{
		ManagedDeviceId:                managedDeviceId,
		SecurityBaselineStateId:        securityBaselineStateId,
		SecurityBaselineSettingStateId: securityBaselineSettingStateId,
	}
}

// ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID parses 'input' into a DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId
func ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID(input string) (*DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(input string) (*DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.SecurityBaselineStateId, ok = input.Parsed["securityBaselineStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", input)
	}

	if id.SecurityBaselineSettingStateId, ok = input.Parsed["securityBaselineSettingStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineSettingStateId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID checks that 'input' can be parsed as a Device Management Managed Device Id Security Baseline State Id Setting State ID
func ValidateDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Security Baseline State Id Setting State ID
func (id DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/securityBaselineStates/%s/settingStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId, id.SecurityBaselineSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Security Baseline State Id Setting State ID
func (id DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
		resourceids.StaticSegment("settingStates", "settingStates", "settingStates"),
		resourceids.UserSpecifiedSegment("securityBaselineSettingStateId", "securityBaselineSettingStateId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Security Baseline State Id Setting State ID
func (id DeviceManagementManagedDeviceIdSecurityBaselineStateIdSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
		fmt.Sprintf("Security Baseline Setting State: %q", id.SecurityBaselineSettingStateId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Security Baseline State Id Setting State (%s)", strings.Join(components, "\n"))
}
