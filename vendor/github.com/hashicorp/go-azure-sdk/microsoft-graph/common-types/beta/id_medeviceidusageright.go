package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdUsageRightId{}

// MeDeviceIdUsageRightId is a struct representing the Resource ID for a Me Device Id Usage Right
type MeDeviceIdUsageRightId struct {
	DeviceId     string
	UsageRightId string
}

// NewMeDeviceIdUsageRightID returns a new MeDeviceIdUsageRightId struct
func NewMeDeviceIdUsageRightID(deviceId string, usageRightId string) MeDeviceIdUsageRightId {
	return MeDeviceIdUsageRightId{
		DeviceId:     deviceId,
		UsageRightId: usageRightId,
	}
}

// ParseMeDeviceIdUsageRightID parses 'input' into a MeDeviceIdUsageRightId
func ParseMeDeviceIdUsageRightID(input string) (*MeDeviceIdUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdUsageRightIDInsensitively parses 'input' case-insensitively into a MeDeviceIdUsageRightId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdUsageRightIDInsensitively(input string) (*MeDeviceIdUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdUsageRightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.UsageRightId, ok = input.Parsed["usageRightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", input)
	}

	return nil
}

// ValidateMeDeviceIdUsageRightID checks that 'input' can be parsed as a Me Device Id Usage Right ID
func ValidateMeDeviceIdUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Usage Right ID
func (id MeDeviceIdUsageRightId) ID() string {
	fmtString := "/me/devices/%s/usageRights/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Usage Right ID
func (id MeDeviceIdUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("usageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightId"),
	}
}

// String returns a human-readable description of this Me Device Id Usage Right ID
func (id MeDeviceIdUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("Me Device Id Usage Right (%s)", strings.Join(components, "\n"))
}
