package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryDeviceLocalCredentialId{}

// DirectoryDeviceLocalCredentialId is a struct representing the Resource ID for a Directory Device Local Credential
type DirectoryDeviceLocalCredentialId struct {
	DeviceLocalCredentialInfoId string
}

// NewDirectoryDeviceLocalCredentialID returns a new DirectoryDeviceLocalCredentialId struct
func NewDirectoryDeviceLocalCredentialID(deviceLocalCredentialInfoId string) DirectoryDeviceLocalCredentialId {
	return DirectoryDeviceLocalCredentialId{
		DeviceLocalCredentialInfoId: deviceLocalCredentialInfoId,
	}
}

// ParseDirectoryDeviceLocalCredentialID parses 'input' into a DirectoryDeviceLocalCredentialId
func ParseDirectoryDeviceLocalCredentialID(input string) (*DirectoryDeviceLocalCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryDeviceLocalCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryDeviceLocalCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryDeviceLocalCredentialIDInsensitively parses 'input' case-insensitively into a DirectoryDeviceLocalCredentialId
// note: this method should only be used for API response data and not user input
func ParseDirectoryDeviceLocalCredentialIDInsensitively(input string) (*DirectoryDeviceLocalCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryDeviceLocalCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryDeviceLocalCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryDeviceLocalCredentialId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceLocalCredentialInfoId, ok = input.Parsed["deviceLocalCredentialInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceLocalCredentialInfoId", input)
	}

	return nil
}

// ValidateDirectoryDeviceLocalCredentialID checks that 'input' can be parsed as a Directory Device Local Credential ID
func ValidateDirectoryDeviceLocalCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryDeviceLocalCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Device Local Credential ID
func (id DirectoryDeviceLocalCredentialId) ID() string {
	fmtString := "/directory/deviceLocalCredentials/%s"
	return fmt.Sprintf(fmtString, id.DeviceLocalCredentialInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Device Local Credential ID
func (id DirectoryDeviceLocalCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("deviceLocalCredentials", "deviceLocalCredentials", "deviceLocalCredentials"),
		resourceids.UserSpecifiedSegment("deviceLocalCredentialInfoId", "deviceLocalCredentialInfoId"),
	}
}

// String returns a human-readable description of this Directory Device Local Credential ID
func (id DirectoryDeviceLocalCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Device Local Credential Info: %q", id.DeviceLocalCredentialInfoId),
	}
	return fmt.Sprintf("Directory Device Local Credential (%s)", strings.Join(components, "\n"))
}
