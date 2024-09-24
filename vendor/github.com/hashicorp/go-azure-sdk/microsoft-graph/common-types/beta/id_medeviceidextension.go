package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdExtensionId{}

// MeDeviceIdExtensionId is a struct representing the Resource ID for a Me Device Id Extension
type MeDeviceIdExtensionId struct {
	DeviceId    string
	ExtensionId string
}

// NewMeDeviceIdExtensionID returns a new MeDeviceIdExtensionId struct
func NewMeDeviceIdExtensionID(deviceId string, extensionId string) MeDeviceIdExtensionId {
	return MeDeviceIdExtensionId{
		DeviceId:    deviceId,
		ExtensionId: extensionId,
	}
}

// ParseMeDeviceIdExtensionID parses 'input' into a MeDeviceIdExtensionId
func ParseMeDeviceIdExtensionID(input string) (*MeDeviceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeDeviceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdExtensionIDInsensitively(input string) (*MeDeviceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeDeviceIdExtensionID checks that 'input' can be parsed as a Me Device Id Extension ID
func ValidateMeDeviceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Extension ID
func (id MeDeviceIdExtensionId) ID() string {
	fmtString := "/me/devices/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Extension ID
func (id MeDeviceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Device Id Extension ID
func (id MeDeviceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Device Id Extension (%s)", strings.Join(components, "\n"))
}
