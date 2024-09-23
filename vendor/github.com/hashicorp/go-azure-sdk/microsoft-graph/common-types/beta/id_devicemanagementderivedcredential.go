package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDerivedCredentialId{}

// DeviceManagementDerivedCredentialId is a struct representing the Resource ID for a Device Management Derived Credential
type DeviceManagementDerivedCredentialId struct {
	DeviceManagementDerivedCredentialSettingsId string
}

// NewDeviceManagementDerivedCredentialID returns a new DeviceManagementDerivedCredentialId struct
func NewDeviceManagementDerivedCredentialID(deviceManagementDerivedCredentialSettingsId string) DeviceManagementDerivedCredentialId {
	return DeviceManagementDerivedCredentialId{
		DeviceManagementDerivedCredentialSettingsId: deviceManagementDerivedCredentialSettingsId,
	}
}

// ParseDeviceManagementDerivedCredentialID parses 'input' into a DeviceManagementDerivedCredentialId
func ParseDeviceManagementDerivedCredentialID(input string) (*DeviceManagementDerivedCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDerivedCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDerivedCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDerivedCredentialIDInsensitively parses 'input' case-insensitively into a DeviceManagementDerivedCredentialId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDerivedCredentialIDInsensitively(input string) (*DeviceManagementDerivedCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDerivedCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDerivedCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDerivedCredentialId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementDerivedCredentialSettingsId, ok = input.Parsed["deviceManagementDerivedCredentialSettingsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementDerivedCredentialSettingsId", input)
	}

	return nil
}

// ValidateDeviceManagementDerivedCredentialID checks that 'input' can be parsed as a Device Management Derived Credential ID
func ValidateDeviceManagementDerivedCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDerivedCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Derived Credential ID
func (id DeviceManagementDerivedCredentialId) ID() string {
	fmtString := "/deviceManagement/derivedCredentials/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementDerivedCredentialSettingsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Derived Credential ID
func (id DeviceManagementDerivedCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("derivedCredentials", "derivedCredentials", "derivedCredentials"),
		resourceids.UserSpecifiedSegment("deviceManagementDerivedCredentialSettingsId", "deviceManagementDerivedCredentialSettingsId"),
	}
}

// String returns a human-readable description of this Device Management Derived Credential ID
func (id DeviceManagementDerivedCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Derived Credential Settings: %q", id.DeviceManagementDerivedCredentialSettingsId),
	}
	return fmt.Sprintf("Device Management Derived Credential (%s)", strings.Join(components, "\n"))
}
