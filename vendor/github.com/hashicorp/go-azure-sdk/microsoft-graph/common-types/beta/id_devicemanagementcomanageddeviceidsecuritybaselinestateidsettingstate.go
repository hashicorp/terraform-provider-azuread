package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{}

// DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId is a struct representing the Resource ID for a Device Management Comanaged Device Id Security Baseline State Id Setting State
type DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId struct {
	ManagedDeviceId                string
	SecurityBaselineStateId        string
	SecurityBaselineSettingStateId string
}

// NewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID returns a new DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId struct
func NewDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(managedDeviceId string, securityBaselineStateId string, securityBaselineSettingStateId string) DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId {
	return DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{
		ManagedDeviceId:                managedDeviceId,
		SecurityBaselineStateId:        securityBaselineStateId,
		SecurityBaselineSettingStateId: securityBaselineSettingStateId,
	}
}

// ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID parses 'input' into a DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId
func ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(input string) (*DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(input string) (*DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID checks that 'input' can be parsed as a Device Management Comanaged Device Id Security Baseline State Id Setting State ID
func ValidateDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Comanaged Device Id Security Baseline State Id Setting State ID
func (id DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId) ID() string {
	fmtString := "/deviceManagement/comanagedDevices/%s/securityBaselineStates/%s/settingStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId, id.SecurityBaselineSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Comanaged Device Id Security Baseline State Id Setting State ID
func (id DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("comanagedDevices", "comanagedDevices", "comanagedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
		resourceids.StaticSegment("settingStates", "settingStates", "settingStates"),
		resourceids.UserSpecifiedSegment("securityBaselineSettingStateId", "securityBaselineSettingStateId"),
	}
}

// String returns a human-readable description of this Device Management Comanaged Device Id Security Baseline State Id Setting State ID
func (id DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
		fmt.Sprintf("Security Baseline Setting State: %q", id.SecurityBaselineSettingStateId),
	}
	return fmt.Sprintf("Device Management Comanaged Device Id Security Baseline State Id Setting State (%s)", strings.Join(components, "\n"))
}
