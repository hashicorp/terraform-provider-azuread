package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdRegisteredUserId{}

// MeDeviceIdRegisteredUserId is a struct representing the Resource ID for a Me Device Id Registered User
type MeDeviceIdRegisteredUserId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceIdRegisteredUserID returns a new MeDeviceIdRegisteredUserId struct
func NewMeDeviceIdRegisteredUserID(deviceId string, directoryObjectId string) MeDeviceIdRegisteredUserId {
	return MeDeviceIdRegisteredUserId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceIdRegisteredUserID parses 'input' into a MeDeviceIdRegisteredUserId
func ParseMeDeviceIdRegisteredUserID(input string) (*MeDeviceIdRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdRegisteredUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdRegisteredUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDeviceIdRegisteredUserIDInsensitively parses 'input' case-insensitively into a MeDeviceIdRegisteredUserId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceIdRegisteredUserIDInsensitively(input string) (*MeDeviceIdRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDeviceIdRegisteredUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDeviceIdRegisteredUserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDeviceIdRegisteredUserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeDeviceIdRegisteredUserID checks that 'input' can be parsed as a Me Device Id Registered User ID
func ValidateMeDeviceIdRegisteredUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceIdRegisteredUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Id Registered User ID
func (id MeDeviceIdRegisteredUserId) ID() string {
	fmtString := "/me/devices/%s/registeredUsers/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Id Registered User ID
func (id MeDeviceIdRegisteredUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("registeredUsers", "registeredUsers", "registeredUsers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Device Id Registered User ID
func (id MeDeviceIdRegisteredUserId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Id Registered User (%s)", strings.Join(components, "\n"))
}
