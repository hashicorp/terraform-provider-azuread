package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeRegisteredDeviceId{}

// MeRegisteredDeviceId is a struct representing the Resource ID for a Me Registered Device
type MeRegisteredDeviceId struct {
	DirectoryObjectId string
}

// NewMeRegisteredDeviceID returns a new MeRegisteredDeviceId struct
func NewMeRegisteredDeviceID(directoryObjectId string) MeRegisteredDeviceId {
	return MeRegisteredDeviceId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeRegisteredDeviceID parses 'input' into a MeRegisteredDeviceId
func ParseMeRegisteredDeviceID(input string) (*MeRegisteredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeRegisteredDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeRegisteredDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeRegisteredDeviceIDInsensitively parses 'input' case-insensitively into a MeRegisteredDeviceId
// note: this method should only be used for API response data and not user input
func ParseMeRegisteredDeviceIDInsensitively(input string) (*MeRegisteredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeRegisteredDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeRegisteredDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeRegisteredDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateMeRegisteredDeviceID checks that 'input' can be parsed as a Me Registered Device ID
func ValidateMeRegisteredDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeRegisteredDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Registered Device ID
func (id MeRegisteredDeviceId) ID() string {
	fmtString := "/me/registeredDevices/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Registered Device ID
func (id MeRegisteredDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("registeredDevices", "registeredDevices", "registeredDevices"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Me Registered Device ID
func (id MeRegisteredDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Registered Device (%s)", strings.Join(components, "\n"))
}
