package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdDeviceConfigurationStateId{}

// MeManagedDeviceIdDeviceConfigurationStateId is a struct representing the Resource ID for a Me Managed Device Id Device Configuration State
type MeManagedDeviceIdDeviceConfigurationStateId struct {
	ManagedDeviceId            string
	DeviceConfigurationStateId string
}

// NewMeManagedDeviceIdDeviceConfigurationStateID returns a new MeManagedDeviceIdDeviceConfigurationStateId struct
func NewMeManagedDeviceIdDeviceConfigurationStateID(managedDeviceId string, deviceConfigurationStateId string) MeManagedDeviceIdDeviceConfigurationStateId {
	return MeManagedDeviceIdDeviceConfigurationStateId{
		ManagedDeviceId:            managedDeviceId,
		DeviceConfigurationStateId: deviceConfigurationStateId,
	}
}

// ParseMeManagedDeviceIdDeviceConfigurationStateID parses 'input' into a MeManagedDeviceIdDeviceConfigurationStateId
func ParseMeManagedDeviceIdDeviceConfigurationStateID(input string) (*MeManagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIdDeviceConfigurationStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceIdDeviceConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIdDeviceConfigurationStateIDInsensitively(input string) (*MeManagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceIdDeviceConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceConfigurationStateId, ok = input.Parsed["deviceConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", input)
	}

	return nil
}

// ValidateMeManagedDeviceIdDeviceConfigurationStateID checks that 'input' can be parsed as a Me Managed Device Id Device Configuration State ID
func ValidateMeManagedDeviceIdDeviceConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceIdDeviceConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Id Device Configuration State ID
func (id MeManagedDeviceIdDeviceConfigurationStateId) ID() string {
	fmtString := "/me/managedDevices/%s/deviceConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Id Device Configuration State ID
func (id MeManagedDeviceIdDeviceConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceConfigurationStates", "deviceConfigurationStates", "deviceConfigurationStates"),
		resourceids.UserSpecifiedSegment("deviceConfigurationStateId", "deviceConfigurationStateId"),
	}
}

// String returns a human-readable description of this Me Managed Device Id Device Configuration State ID
func (id MeManagedDeviceIdDeviceConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Configuration State: %q", id.DeviceConfigurationStateId),
	}
	return fmt.Sprintf("Me Managed Device Id Device Configuration State (%s)", strings.Join(components, "\n"))
}
