package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceCleanupRuleId{}

// DeviceManagementManagedDeviceCleanupRuleId is a struct representing the Resource ID for a Device Management Managed Device Cleanup Rule
type DeviceManagementManagedDeviceCleanupRuleId struct {
	ManagedDeviceCleanupRuleId string
}

// NewDeviceManagementManagedDeviceCleanupRuleID returns a new DeviceManagementManagedDeviceCleanupRuleId struct
func NewDeviceManagementManagedDeviceCleanupRuleID(managedDeviceCleanupRuleId string) DeviceManagementManagedDeviceCleanupRuleId {
	return DeviceManagementManagedDeviceCleanupRuleId{
		ManagedDeviceCleanupRuleId: managedDeviceCleanupRuleId,
	}
}

// ParseDeviceManagementManagedDeviceCleanupRuleID parses 'input' into a DeviceManagementManagedDeviceCleanupRuleId
func ParseDeviceManagementManagedDeviceCleanupRuleID(input string) (*DeviceManagementManagedDeviceCleanupRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceCleanupRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceCleanupRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceCleanupRuleIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceCleanupRuleId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceCleanupRuleIDInsensitively(input string) (*DeviceManagementManagedDeviceCleanupRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceCleanupRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceCleanupRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceCleanupRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceCleanupRuleId, ok = input.Parsed["managedDeviceCleanupRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceCleanupRuleId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceCleanupRuleID checks that 'input' can be parsed as a Device Management Managed Device Cleanup Rule ID
func ValidateDeviceManagementManagedDeviceCleanupRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceCleanupRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Cleanup Rule ID
func (id DeviceManagementManagedDeviceCleanupRuleId) ID() string {
	fmtString := "/deviceManagement/managedDeviceCleanupRules/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceCleanupRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Cleanup Rule ID
func (id DeviceManagementManagedDeviceCleanupRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDeviceCleanupRules", "managedDeviceCleanupRules", "managedDeviceCleanupRules"),
		resourceids.UserSpecifiedSegment("managedDeviceCleanupRuleId", "managedDeviceCleanupRuleId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Cleanup Rule ID
func (id DeviceManagementManagedDeviceCleanupRuleId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device Cleanup Rule: %q", id.ManagedDeviceCleanupRuleId),
	}
	return fmt.Sprintf("Device Management Managed Device Cleanup Rule (%s)", strings.Join(components, "\n"))
}
