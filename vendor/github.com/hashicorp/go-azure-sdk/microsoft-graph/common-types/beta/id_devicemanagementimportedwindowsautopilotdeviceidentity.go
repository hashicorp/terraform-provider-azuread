package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementImportedWindowsAutopilotDeviceIdentityId{}

// DeviceManagementImportedWindowsAutopilotDeviceIdentityId is a struct representing the Resource ID for a Device Management Imported Windows Autopilot Device Identity
type DeviceManagementImportedWindowsAutopilotDeviceIdentityId struct {
	ImportedWindowsAutopilotDeviceIdentityId string
}

// NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID returns a new DeviceManagementImportedWindowsAutopilotDeviceIdentityId struct
func NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID(importedWindowsAutopilotDeviceIdentityId string) DeviceManagementImportedWindowsAutopilotDeviceIdentityId {
	return DeviceManagementImportedWindowsAutopilotDeviceIdentityId{
		ImportedWindowsAutopilotDeviceIdentityId: importedWindowsAutopilotDeviceIdentityId,
	}
}

// ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityID parses 'input' into a DeviceManagementImportedWindowsAutopilotDeviceIdentityId
func ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityID(input string) (*DeviceManagementImportedWindowsAutopilotDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementImportedWindowsAutopilotDeviceIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementImportedWindowsAutopilotDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityIDInsensitively parses 'input' case-insensitively into a DeviceManagementImportedWindowsAutopilotDeviceIdentityId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityIDInsensitively(input string) (*DeviceManagementImportedWindowsAutopilotDeviceIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementImportedWindowsAutopilotDeviceIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementImportedWindowsAutopilotDeviceIdentityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementImportedWindowsAutopilotDeviceIdentityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ImportedWindowsAutopilotDeviceIdentityId, ok = input.Parsed["importedWindowsAutopilotDeviceIdentityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "importedWindowsAutopilotDeviceIdentityId", input)
	}

	return nil
}

// ValidateDeviceManagementImportedWindowsAutopilotDeviceIdentityID checks that 'input' can be parsed as a Device Management Imported Windows Autopilot Device Identity ID
func ValidateDeviceManagementImportedWindowsAutopilotDeviceIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Imported Windows Autopilot Device Identity ID
func (id DeviceManagementImportedWindowsAutopilotDeviceIdentityId) ID() string {
	fmtString := "/deviceManagement/importedWindowsAutopilotDeviceIdentities/%s"
	return fmt.Sprintf(fmtString, id.ImportedWindowsAutopilotDeviceIdentityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Imported Windows Autopilot Device Identity ID
func (id DeviceManagementImportedWindowsAutopilotDeviceIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("importedWindowsAutopilotDeviceIdentities", "importedWindowsAutopilotDeviceIdentities", "importedWindowsAutopilotDeviceIdentities"),
		resourceids.UserSpecifiedSegment("importedWindowsAutopilotDeviceIdentityId", "importedWindowsAutopilotDeviceIdentityId"),
	}
}

// String returns a human-readable description of this Device Management Imported Windows Autopilot Device Identity ID
func (id DeviceManagementImportedWindowsAutopilotDeviceIdentityId) String() string {
	components := []string{
		fmt.Sprintf("Imported Windows Autopilot Device Identity: %q", id.ImportedWindowsAutopilotDeviceIdentityId),
	}
	return fmt.Sprintf("Device Management Imported Windows Autopilot Device Identity (%s)", strings.Join(components, "\n"))
}
