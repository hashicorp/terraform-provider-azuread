package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdMemberOfId{}

// MeDeviceIdMemberOfId is a struct representing the Resource ID for a Me Device Id Member Of
type MeDeviceIdMemberOfId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceIdMemberOfID returns a new MeDeviceIdMemberOfId struct
func NewMeDeviceIdMemberOfID(deviceId string, directoryObjectId string) MeDeviceIdMemberOfId {
	return MeDeviceIdMemberOfId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceIdMemberOfID parses 'input' into a MeDeviceIdMemberOfId
func ParseMeDeviceIdMemberOfID(input string) (*MeDeviceIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdMemberOfIDInsensitively parses 'input' case-insensitively into a MeDeviceIdMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdMemberOfIDInsensitively(input string) (*MeDeviceIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeDeviceIdMemberOfID checks that 'input' can be parsed as a Me Device Id Member Of ID
func ValidateMeDeviceIdMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Member Of ID
func (id MeDeviceIdMemberOfId) ID() string {
	fmtString := "/me/devices/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Member Of ID
func (id MeDeviceIdMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("memberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Device Id Member Of ID
func (id MeDeviceIdMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Id Member Of (%s)", strings.Join(components, "\n"))
}
