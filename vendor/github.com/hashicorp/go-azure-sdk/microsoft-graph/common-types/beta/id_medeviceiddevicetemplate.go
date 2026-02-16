package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdDeviceTemplateId{}

// MeDeviceIdDeviceTemplateId is a struct representing the Resource ID for a Me Device Id Device Template
type MeDeviceIdDeviceTemplateId struct {
	DeviceId         string
	DeviceTemplateId string
}

// NewMeDeviceIdDeviceTemplateID returns a new MeDeviceIdDeviceTemplateId struct
func NewMeDeviceIdDeviceTemplateID(deviceId string, deviceTemplateId string) MeDeviceIdDeviceTemplateId {
	return MeDeviceIdDeviceTemplateId{
		DeviceId:         deviceId,
		DeviceTemplateId: deviceTemplateId,
	}
}

// ParseMeDeviceIdDeviceTemplateID parses 'input' into a MeDeviceIdDeviceTemplateId
func ParseMeDeviceIdDeviceTemplateID(input string) (*MeDeviceIdDeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdDeviceTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdDeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdDeviceTemplateIDInsensitively parses 'input' case-insensitively into a MeDeviceIdDeviceTemplateId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdDeviceTemplateIDInsensitively(input string) (*MeDeviceIdDeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdDeviceTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdDeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdDeviceTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DeviceTemplateId, ok = input.Parsed["deviceTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceTemplateId", input)
	}

	return nil
}

// ValidateMeDeviceIdDeviceTemplateID checks that 'input' can be parsed as a Me Device Id Device Template ID
func ValidateMeDeviceIdDeviceTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdDeviceTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Device Template ID
func (id MeDeviceIdDeviceTemplateId) ID() string {
	fmtString := "/me/devices/%s/deviceTemplate/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DeviceTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Device Template ID
func (id MeDeviceIdDeviceTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("deviceTemplate", "deviceTemplate", "deviceTemplate"),
		resourceids.UserSpecifiedSegment("deviceTemplateId", "deviceTemplateId"),
	}
}

// String returns a human-readable description of this Me Device Id Device Template ID
func (id MeDeviceIdDeviceTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Device Template: %q", id.DeviceTemplateId),
	}
	return fmt.Sprintf("Me Device Id Device Template (%s)", strings.Join(components, "\n"))
}
