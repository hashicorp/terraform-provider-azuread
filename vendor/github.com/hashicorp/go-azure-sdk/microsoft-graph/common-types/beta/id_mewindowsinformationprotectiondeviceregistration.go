package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeWindowsInformationProtectionDeviceRegistrationId{}

// MeWindowsInformationProtectionDeviceRegistrationId is a struct representing the Resource ID for a Me Windows Information Protection Device Registration
type MeWindowsInformationProtectionDeviceRegistrationId struct {
	WindowsInformationProtectionDeviceRegistrationId string
}

// NewMeWindowsInformationProtectionDeviceRegistrationID returns a new MeWindowsInformationProtectionDeviceRegistrationId struct
func NewMeWindowsInformationProtectionDeviceRegistrationID(windowsInformationProtectionDeviceRegistrationId string) MeWindowsInformationProtectionDeviceRegistrationId {
	return MeWindowsInformationProtectionDeviceRegistrationId{
		WindowsInformationProtectionDeviceRegistrationId: windowsInformationProtectionDeviceRegistrationId,
	}
}

// ParseMeWindowsInformationProtectionDeviceRegistrationID parses 'input' into a MeWindowsInformationProtectionDeviceRegistrationId
func ParseMeWindowsInformationProtectionDeviceRegistrationID(input string) (*MeWindowsInformationProtectionDeviceRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeWindowsInformationProtectionDeviceRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeWindowsInformationProtectionDeviceRegistrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeWindowsInformationProtectionDeviceRegistrationIDInsensitively parses 'input' case-insensitively into a MeWindowsInformationProtectionDeviceRegistrationId
// note: this method should only be used for API response data and not user input
func ParseMeWindowsInformationProtectionDeviceRegistrationIDInsensitively(input string) (*MeWindowsInformationProtectionDeviceRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeWindowsInformationProtectionDeviceRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeWindowsInformationProtectionDeviceRegistrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeWindowsInformationProtectionDeviceRegistrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.WindowsInformationProtectionDeviceRegistrationId, ok = input.Parsed["windowsInformationProtectionDeviceRegistrationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsInformationProtectionDeviceRegistrationId", input)
	}

	return nil
}

// ValidateMeWindowsInformationProtectionDeviceRegistrationID checks that 'input' can be parsed as a Me Windows Information Protection Device Registration ID
func ValidateMeWindowsInformationProtectionDeviceRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeWindowsInformationProtectionDeviceRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Windows Information Protection Device Registration ID
func (id MeWindowsInformationProtectionDeviceRegistrationId) ID() string {
	fmtString := "/me/windowsInformationProtectionDeviceRegistrations/%s"
	return fmt.Sprintf(fmtString, id.WindowsInformationProtectionDeviceRegistrationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Windows Information Protection Device Registration ID
func (id MeWindowsInformationProtectionDeviceRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("windowsInformationProtectionDeviceRegistrations", "windowsInformationProtectionDeviceRegistrations", "windowsInformationProtectionDeviceRegistrations"),
		resourceids.UserSpecifiedSegment("windowsInformationProtectionDeviceRegistrationId", "windowsInformationProtectionDeviceRegistrationId"),
	}
}

// String returns a human-readable description of this Me Windows Information Protection Device Registration ID
func (id MeWindowsInformationProtectionDeviceRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("Windows Information Protection Device Registration: %q", id.WindowsInformationProtectionDeviceRegistrationId),
	}
	return fmt.Sprintf("Me Windows Information Protection Device Registration (%s)", strings.Join(components, "\n"))
}
