package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdDeviceCompliancePolicyStateId{}

// MeManagedDeviceIdDeviceCompliancePolicyStateId is a struct representing the Resource ID for a Me Managed Device Id Device Compliance Policy State
type MeManagedDeviceIdDeviceCompliancePolicyStateId struct {
	ManagedDeviceId               string
	DeviceCompliancePolicyStateId string
}

// NewMeManagedDeviceIdDeviceCompliancePolicyStateID returns a new MeManagedDeviceIdDeviceCompliancePolicyStateId struct
func NewMeManagedDeviceIdDeviceCompliancePolicyStateID(managedDeviceId string, deviceCompliancePolicyStateId string) MeManagedDeviceIdDeviceCompliancePolicyStateId {
	return MeManagedDeviceIdDeviceCompliancePolicyStateId{
		ManagedDeviceId:               managedDeviceId,
		DeviceCompliancePolicyStateId: deviceCompliancePolicyStateId,
	}
}

// ParseMeManagedDeviceIdDeviceCompliancePolicyStateID parses 'input' into a MeManagedDeviceIdDeviceCompliancePolicyStateId
func ParseMeManagedDeviceIdDeviceCompliancePolicyStateID(input string) (*MeManagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdDeviceCompliancePolicyStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively(input string) (*MeManagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdDeviceCompliancePolicyStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceCompliancePolicyStateId, ok = input.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdDeviceCompliancePolicyStateID checks that 'input' can be parsed as a Me Managed Device Id Device Compliance Policy State ID
func ValidateMeManagedDeviceIdDeviceCompliancePolicyStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdDeviceCompliancePolicyStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Device Compliance Policy State ID
func (id MeManagedDeviceIdDeviceCompliancePolicyStateId) ID() string {
	fmtString := "/me/managedDevices/%s/deviceCompliancePolicyStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceCompliancePolicyStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Device Compliance Policy State ID
func (id MeManagedDeviceIdDeviceCompliancePolicyStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceCompliancePolicyStates", "deviceCompliancePolicyStates", "deviceCompliancePolicyStates"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyStateId", "deviceCompliancePolicyStateId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Device Compliance Policy State ID
func (id MeManagedDeviceIdDeviceCompliancePolicyStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Compliance Policy State: %q", id.DeviceCompliancePolicyStateId),
	}
	return fmt.Sprintf("Me Managed Device Id Device Compliance Policy State (%s)", strings.Join(components, "\n"))
}
