package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceId{}

// MeManagedDeviceId is a struct representing the Resource ID for a Me Managed Device
type MeManagedDeviceId struct {
	ManagedDeviceId string
}

// NewMeManagedDeviceID returns a new MeManagedDeviceId struct
func NewMeManagedDeviceID(managedDeviceId string) MeManagedDeviceId {
	return MeManagedDeviceId{
		ManagedDeviceId: managedDeviceId,
	}
}

// ParseMeManagedDeviceID parses 'input' into a MeManagedDeviceId
func ParseMeManagedDeviceID(input string) (*MeManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedDeviceIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceIDInsensitively(input string) (*MeManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	return nil
}

// ValidateMeManagedDeviceID checks that 'input' can be parsed as a Me Managed Device ID
func ValidateMeManagedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device ID
func (id MeManagedDeviceId) ID() string {
	fmtString := "/me/managedDevices/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device ID
func (id MeManagedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
	}
}

// String returns a human-readable description of this Me Managed Device ID
func (id MeManagedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
	}
	return fmt.Sprintf("Me Managed Device (%s)", strings.Join(components, "\n"))
}
