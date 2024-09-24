package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdSecurityBaselineStateIdSettingStateId{}

// MeManagedDeviceIdSecurityBaselineStateIdSettingStateId is a struct representing the Resource ID for a Me Managed Device Id Security Baseline State Id Setting State
type MeManagedDeviceIdSecurityBaselineStateIdSettingStateId struct {
	ManagedDeviceId                string
	SecurityBaselineStateId        string
	SecurityBaselineSettingStateId string
}

// NewMeManagedDeviceIdSecurityBaselineStateIdSettingStateID returns a new MeManagedDeviceIdSecurityBaselineStateIdSettingStateId struct
func NewMeManagedDeviceIdSecurityBaselineStateIdSettingStateID(managedDeviceId string, securityBaselineStateId string, securityBaselineSettingStateId string) MeManagedDeviceIdSecurityBaselineStateIdSettingStateId {
	return MeManagedDeviceIdSecurityBaselineStateIdSettingStateId{
		ManagedDeviceId:                managedDeviceId,
		SecurityBaselineStateId:        securityBaselineStateId,
		SecurityBaselineSettingStateId: securityBaselineSettingStateId,
	}
}

// ParseMeManagedDeviceIdSecurityBaselineStateIdSettingStateID parses 'input' into a MeManagedDeviceIdSecurityBaselineStateIdSettingStateId
func ParseMeManagedDeviceIdSecurityBaselineStateIdSettingStateID(input string) (*MeManagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdSecurityBaselineStateIdSettingStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdSecurityBaselineStateIdSettingStateIDInsensitively(input string) (*MeManagedDeviceIdSecurityBaselineStateIdSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdSecurityBaselineStateIdSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdSecurityBaselineStateIdSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdSecurityBaselineStateIdSettingStateId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeManagedDeviceIdSecurityBaselineStateIdSettingStateID checks that 'input' can be parsed as a Me Managed Device Id Security Baseline State Id Setting State ID
func ValidateMeManagedDeviceIdSecurityBaselineStateIdSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdSecurityBaselineStateIdSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Security Baseline State Id Setting State ID
func (id MeManagedDeviceIdSecurityBaselineStateIdSettingStateId) ID() string {
	fmtString := "/me/managedDevices/%s/securityBaselineStates/%s/settingStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId, id.SecurityBaselineSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Security Baseline State Id Setting State ID
func (id MeManagedDeviceIdSecurityBaselineStateIdSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
		resourceids.StaticSegment("settingStates", "settingStates", "settingStates"),
		resourceids.UserSpecifiedSegment("securityBaselineSettingStateId", "securityBaselineSettingStateId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Security Baseline State Id Setting State ID
func (id MeManagedDeviceIdSecurityBaselineStateIdSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
		fmt.Sprintf("Security Baseline Setting State: %q", id.SecurityBaselineSettingStateId),
	}
	return fmt.Sprintf("Me Managed Device Id Security Baseline State Id Setting State (%s)", strings.Join(components, "\n"))
}
