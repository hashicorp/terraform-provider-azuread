package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}

// MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId is a struct representing the Resource ID for a Me Managed Device Id Managed Device Mobile App Configuration State
type MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct {
	ManagedDeviceId                            string
	ManagedDeviceMobileAppConfigurationStateId string
}

// NewMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID returns a new MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct
func NewMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(managedDeviceId string, managedDeviceMobileAppConfigurationStateId string) MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId {
	return MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
		ManagedDeviceId: managedDeviceId,
		ManagedDeviceMobileAppConfigurationStateId: managedDeviceMobileAppConfigurationStateId,
	}
}

// ParseMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID parses 'input' into a MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
func ParseMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input string) (*MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(input string) (*MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = input.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID checks that 'input' can be parsed as a Me Managed Device Id Managed Device Mobile App Configuration State ID
func ValidateMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Managed Device Mobile App Configuration State ID
func (id MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) ID() string {
	fmtString := "/me/managedDevices/%s/managedDeviceMobileAppConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.ManagedDeviceMobileAppConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Managed Device Mobile App Configuration State ID
func (id MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates"),
		resourceids.UserSpecifiedSegment("managedDeviceMobileAppConfigurationStateId", "managedDeviceMobileAppConfigurationStateId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Managed Device Mobile App Configuration State ID
func (id MeManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Managed Device Mobile App Configuration State: %q", id.ManagedDeviceMobileAppConfigurationStateId),
	}
	return fmt.Sprintf("Me Managed Device Id Managed Device Mobile App Configuration State (%s)", strings.Join(components, "\n"))
}
