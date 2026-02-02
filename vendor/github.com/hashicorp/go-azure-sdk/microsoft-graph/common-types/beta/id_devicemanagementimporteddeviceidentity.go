package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementImportedDeviceIdentityId{}

// DeviceManagementImportedDeviceIdentityId is a struct representing the Resource ID for a Device Management Imported Device Identity
type DeviceManagementImportedDeviceIdentityId struct {
	ImportedDeviceIdentityId string
}

// NewDeviceManagementImportedDeviceIdentityID returns a new DeviceManagementImportedDeviceIdentityId struct
func NewDeviceManagementImportedDeviceIdentityID(importedDeviceIdentityId string) DeviceManagementImportedDeviceIdentityId {
	return DeviceManagementImportedDeviceIdentityId{
		ImportedDeviceIdentityId: importedDeviceIdentityId,
	}
}

// ParseDeviceManagementImportedDeviceIdentityID parses 'input' into a DeviceManagementImportedDeviceIdentityId
func ParseDeviceManagementImportedDeviceIdentityID(input string) (*DeviceManagementImportedDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementImportedDeviceIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementImportedDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementImportedDeviceIdentityIDInsensitively parses 'input' case-insensitively into a DeviceManagementImportedDeviceIdentityId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementImportedDeviceIdentityIDInsensitively(input string) (*DeviceManagementImportedDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementImportedDeviceIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementImportedDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementImportedDeviceIdentityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ImportedDeviceIdentityId, ok = input.Parsed["importedDeviceIdentityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "importedDeviceIdentityId", input)
	}

	return nil
}

// ValidateDeviceManagementImportedDeviceIdentityID checks that 'input' can be parsed as a Device Management Imported Device Identity ID
func ValidateDeviceManagementImportedDeviceIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementImportedDeviceIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Imported Device Identity ID
func (id DeviceManagementImportedDeviceIdentityId) ID() string {
	fmtString := "/deviceManagement/importedDeviceIdentities/%s"
	return fmt.Sprintf(fmtString, id.ImportedDeviceIdentityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Imported Device Identity ID
func (id DeviceManagementImportedDeviceIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("importedDeviceIdentities", "importedDeviceIdentities", "importedDeviceIdentities"),
		resourceids.UserSpecifiedSegment("importedDeviceIdentityId", "importedDeviceIdentityId"),
	}
}

// String returns a human-readable description of this Device Management Imported Device Identity ID
func (id DeviceManagementImportedDeviceIdentityId) String() string {
	components := []string{
		fmt.Sprintf("Imported Device Identity: %q", id.ImportedDeviceIdentityId),
	}
	return fmt.Sprintf("Device Management Imported Device Identity (%s)", strings.Join(components, "\n"))
}
