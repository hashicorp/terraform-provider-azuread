package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdCommandId{}

// MeDeviceIdCommandId is a struct representing the Resource ID for a Me Device Id Command
type MeDeviceIdCommandId struct {
	DeviceId  string
	CommandId string
}

// NewMeDeviceIdCommandID returns a new MeDeviceIdCommandId struct
func NewMeDeviceIdCommandID(deviceId string, commandId string) MeDeviceIdCommandId {
	return MeDeviceIdCommandId{
		DeviceId:  deviceId,
		CommandId: commandId,
	}
}

// ParseMeDeviceIdCommandID parses 'input' into a MeDeviceIdCommandId
func ParseMeDeviceIdCommandID(input string) (*MeDeviceIdCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdCommandIDInsensitively parses 'input' case-insensitively into a MeDeviceIdCommandId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdCommandIDInsensitively(input string) (*MeDeviceIdCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdCommandId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.CommandId, ok = input.Parsed["commandId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "commandId", input)
	}

	return nil
}

// ValidateMeDeviceIdCommandID checks that 'input' can be parsed as a Me Device Id Command ID
func ValidateMeDeviceIdCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Command ID
func (id MeDeviceIdCommandId) ID() string {
	fmtString := "/me/devices/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.CommandId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Command ID
func (id MeDeviceIdCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("commands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandId", "commandId"),
	}
}

// String returns a human-readable description of this Me Device Id Command ID
func (id MeDeviceIdCommandId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Command: %q", id.CommandId),
	}
	return fmt.Sprintf("Me Device Id Command (%s)", strings.Join(components, "\n"))
}
