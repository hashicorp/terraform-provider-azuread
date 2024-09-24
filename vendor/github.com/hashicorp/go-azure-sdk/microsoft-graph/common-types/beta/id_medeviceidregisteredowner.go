package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdRegisteredOwnerId{}

// MeDeviceIdRegisteredOwnerId is a struct representing the Resource ID for a Me Device Id Registered Owner
type MeDeviceIdRegisteredOwnerId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceIdRegisteredOwnerID returns a new MeDeviceIdRegisteredOwnerId struct
func NewMeDeviceIdRegisteredOwnerID(deviceId string, directoryObjectId string) MeDeviceIdRegisteredOwnerId {
	return MeDeviceIdRegisteredOwnerId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceIdRegisteredOwnerID parses 'input' into a MeDeviceIdRegisteredOwnerId
func ParseMeDeviceIdRegisteredOwnerID(input string) (*MeDeviceIdRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdRegisteredOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdRegisteredOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdRegisteredOwnerIDInsensitively parses 'input' case-insensitively into a MeDeviceIdRegisteredOwnerId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdRegisteredOwnerIDInsensitively(input string) (*MeDeviceIdRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdRegisteredOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdRegisteredOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdRegisteredOwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeDeviceIdRegisteredOwnerID checks that 'input' can be parsed as a Me Device Id Registered Owner ID
func ValidateMeDeviceIdRegisteredOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdRegisteredOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Registered Owner ID
func (id MeDeviceIdRegisteredOwnerId) ID() string {
	fmtString := "/me/devices/%s/registeredOwners/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Registered Owner ID
func (id MeDeviceIdRegisteredOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("registeredOwners", "registeredOwners", "registeredOwners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Device Id Registered Owner ID
func (id MeDeviceIdRegisteredOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Id Registered Owner (%s)", strings.Join(components, "\n"))
}
