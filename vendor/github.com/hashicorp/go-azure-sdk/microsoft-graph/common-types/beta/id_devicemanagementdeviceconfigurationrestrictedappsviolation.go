package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationRestrictedAppsViolationId{}

// DeviceManagementDeviceConfigurationRestrictedAppsViolationId is a struct representing the Resource ID for a Device Management Device Configuration Restricted Apps Violation
type DeviceManagementDeviceConfigurationRestrictedAppsViolationId struct {
	RestrictedAppsViolationId string
}

// NewDeviceManagementDeviceConfigurationRestrictedAppsViolationID returns a new DeviceManagementDeviceConfigurationRestrictedAppsViolationId struct
func NewDeviceManagementDeviceConfigurationRestrictedAppsViolationID(restrictedAppsViolationId string) DeviceManagementDeviceConfigurationRestrictedAppsViolationId {
	return DeviceManagementDeviceConfigurationRestrictedAppsViolationId{
		RestrictedAppsViolationId: restrictedAppsViolationId,
	}
}

// ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationID parses 'input' into a DeviceManagementDeviceConfigurationRestrictedAppsViolationId
func ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationID(input string) (*DeviceManagementDeviceConfigurationRestrictedAppsViolationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationRestrictedAppsViolationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationRestrictedAppsViolationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationRestrictedAppsViolationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationIDInsensitively(input string) (*DeviceManagementDeviceConfigurationRestrictedAppsViolationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationRestrictedAppsViolationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationRestrictedAppsViolationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationRestrictedAppsViolationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RestrictedAppsViolationId, ok = input.Parsed["restrictedAppsViolationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "restrictedAppsViolationId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationRestrictedAppsViolationID checks that 'input' can be parsed as a Device Management Device Configuration Restricted Apps Violation ID
func ValidateDeviceManagementDeviceConfigurationRestrictedAppsViolationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationRestrictedAppsViolationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Restricted Apps Violation ID
func (id DeviceManagementDeviceConfigurationRestrictedAppsViolationId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurationRestrictedAppsViolations/%s"
	return fmt.Sprintf(fmtString, id.RestrictedAppsViolationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Restricted Apps Violation ID
func (id DeviceManagementDeviceConfigurationRestrictedAppsViolationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurationRestrictedAppsViolations", "deviceConfigurationRestrictedAppsViolations", "deviceConfigurationRestrictedAppsViolations"),
		resourceids.UserSpecifiedSegment("restrictedAppsViolationId", "restrictedAppsViolationId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Restricted Apps Violation ID
func (id DeviceManagementDeviceConfigurationRestrictedAppsViolationId) String() string {
	components := []string{
		fmt.Sprintf("Restricted Apps Violation: %q", id.RestrictedAppsViolationId),
	}
	return fmt.Sprintf("Device Management Device Configuration Restricted Apps Violation (%s)", strings.Join(components, "\n"))
}
