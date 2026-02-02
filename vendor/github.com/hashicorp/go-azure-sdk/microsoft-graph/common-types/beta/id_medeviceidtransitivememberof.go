package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdTransitiveMemberOfId{}

// MeDeviceIdTransitiveMemberOfId is a struct representing the Resource ID for a Me Device Id Transitive Member Of
type MeDeviceIdTransitiveMemberOfId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceIdTransitiveMemberOfID returns a new MeDeviceIdTransitiveMemberOfId struct
func NewMeDeviceIdTransitiveMemberOfID(deviceId string, directoryObjectId string) MeDeviceIdTransitiveMemberOfId {
	return MeDeviceIdTransitiveMemberOfId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceIdTransitiveMemberOfID parses 'input' into a MeDeviceIdTransitiveMemberOfId
func ParseMeDeviceIdTransitiveMemberOfID(input string) (*MeDeviceIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a MeDeviceIdTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdTransitiveMemberOfIDInsensitively(input string) (*MeDeviceIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdTransitiveMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeDeviceIdTransitiveMemberOfID checks that 'input' can be parsed as a Me Device Id Transitive Member Of ID
func ValidateMeDeviceIdTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Transitive Member Of ID
func (id MeDeviceIdTransitiveMemberOfId) ID() string {
	fmtString := "/me/devices/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Transitive Member Of ID
func (id MeDeviceIdTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("transitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Device Id Transitive Member Of ID
func (id MeDeviceIdTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Id Transitive Member Of (%s)", strings.Join(components, "\n"))
}
