package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdSecurityBaselineStateId{}

// MeManagedDeviceIdSecurityBaselineStateId is a struct representing the Resource ID for a Me Managed Device Id Security Baseline State
type MeManagedDeviceIdSecurityBaselineStateId struct {
	ManagedDeviceId         string
	SecurityBaselineStateId string
}

// NewMeManagedDeviceIdSecurityBaselineStateID returns a new MeManagedDeviceIdSecurityBaselineStateId struct
func NewMeManagedDeviceIdSecurityBaselineStateID(managedDeviceId string, securityBaselineStateId string) MeManagedDeviceIdSecurityBaselineStateId {
	return MeManagedDeviceIdSecurityBaselineStateId{
		ManagedDeviceId:         managedDeviceId,
		SecurityBaselineStateId: securityBaselineStateId,
	}
}

// ParseMeManagedDeviceIdSecurityBaselineStateID parses 'input' into a MeManagedDeviceIdSecurityBaselineStateId
func ParseMeManagedDeviceIdSecurityBaselineStateID(input string) (*MeManagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdSecurityBaselineStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdSecurityBaselineStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdSecurityBaselineStateIDInsensitively(input string) (*MeManagedDeviceIdSecurityBaselineStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdSecurityBaselineStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdSecurityBaselineStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdSecurityBaselineStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.SecurityBaselineStateId, ok = input.Parsed["securityBaselineStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityBaselineStateId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdSecurityBaselineStateID checks that 'input' can be parsed as a Me Managed Device Id Security Baseline State ID
func ValidateMeManagedDeviceIdSecurityBaselineStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdSecurityBaselineStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Security Baseline State ID
func (id MeManagedDeviceIdSecurityBaselineStateId) ID() string {
	fmtString := "/me/managedDevices/%s/securityBaselineStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.SecurityBaselineStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Security Baseline State ID
func (id MeManagedDeviceIdSecurityBaselineStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("securityBaselineStates", "securityBaselineStates", "securityBaselineStates"),
		resourceids.UserSpecifiedSegment("securityBaselineStateId", "securityBaselineStateId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Security Baseline State ID
func (id MeManagedDeviceIdSecurityBaselineStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Security Baseline State: %q", id.SecurityBaselineStateId),
	}
	return fmt.Sprintf("Me Managed Device Id Security Baseline State (%s)", strings.Join(components, "\n"))
}
