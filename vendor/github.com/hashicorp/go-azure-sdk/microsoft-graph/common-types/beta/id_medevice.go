package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceId{}

// MeDeviceId is a struct representing the Resource ID for a Me Device
type MeDeviceId struct {
	DeviceId string
}

// NewMeDeviceID returns a new MeDeviceId struct
func NewMeDeviceID(deviceId string) MeDeviceId {
	return MeDeviceId{
		DeviceId: deviceId,
	}
}

// ParseMeDeviceID parses 'input' into a MeDeviceId
func ParseMeDeviceID(input string) (*MeDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIDInsensitively parses 'input' case-insensitively into a MeDeviceId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIDInsensitively(input string) (*MeDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	return nil
}

// ValidateMeDeviceID checks that 'input' can be parsed as a Me Device ID
func ValidateMeDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device ID
func (id MeDeviceId) ID() string {
	fmtString := "/me/devices/%s"
	return fmt.Sprintf(fmtString, id.DeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device ID
func (id MeDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
	}
}

// String returns a human-readable description of this Me Device ID
func (id MeDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
	}
	return fmt.Sprintf("Me Device (%s)", strings.Join(components, "\n"))
}
