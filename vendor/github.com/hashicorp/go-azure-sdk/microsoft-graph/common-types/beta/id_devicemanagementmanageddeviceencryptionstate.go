package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceEncryptionStateId{}

// DeviceManagementManagedDeviceEncryptionStateId is a struct representing the Resource ID for a Device Management Managed Device Encryption State
type DeviceManagementManagedDeviceEncryptionStateId struct {
	ManagedDeviceEncryptionStateId string
}

// NewDeviceManagementManagedDeviceEncryptionStateID returns a new DeviceManagementManagedDeviceEncryptionStateId struct
func NewDeviceManagementManagedDeviceEncryptionStateID(managedDeviceEncryptionStateId string) DeviceManagementManagedDeviceEncryptionStateId {
	return DeviceManagementManagedDeviceEncryptionStateId{
		ManagedDeviceEncryptionStateId: managedDeviceEncryptionStateId,
	}
}

// ParseDeviceManagementManagedDeviceEncryptionStateID parses 'input' into a DeviceManagementManagedDeviceEncryptionStateId
func ParseDeviceManagementManagedDeviceEncryptionStateID(input string) (*DeviceManagementManagedDeviceEncryptionStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceEncryptionStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceEncryptionStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceEncryptionStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceEncryptionStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceEncryptionStateIDInsensitively(input string) (*DeviceManagementManagedDeviceEncryptionStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceEncryptionStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceEncryptionStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceEncryptionStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceEncryptionStateId, ok = input.Parsed["managedDeviceEncryptionStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceEncryptionStateId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceEncryptionStateID checks that 'input' can be parsed as a Device Management Managed Device Encryption State ID
func ValidateDeviceManagementManagedDeviceEncryptionStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceEncryptionStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Encryption State ID
func (id DeviceManagementManagedDeviceEncryptionStateId) ID() string {
	fmtString := "/deviceManagement/managedDeviceEncryptionStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceEncryptionStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Encryption State ID
func (id DeviceManagementManagedDeviceEncryptionStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDeviceEncryptionStates", "managedDeviceEncryptionStates", "managedDeviceEncryptionStates"),
		resourceids.UserSpecifiedSegment("managedDeviceEncryptionStateId", "managedDeviceEncryptionStateId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Encryption State ID
func (id DeviceManagementManagedDeviceEncryptionStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device Encryption State: %q", id.ManagedDeviceEncryptionStateId),
	}
	return fmt.Sprintf("Device Management Managed Device Encryption State (%s)", strings.Join(components, "\n"))
}
