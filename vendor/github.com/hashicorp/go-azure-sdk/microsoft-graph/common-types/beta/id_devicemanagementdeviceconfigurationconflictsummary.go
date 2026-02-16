package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationConflictSummaryId{}

// DeviceManagementDeviceConfigurationConflictSummaryId is a struct representing the Resource ID for a Device Management Device Configuration Conflict Summary
type DeviceManagementDeviceConfigurationConflictSummaryId struct {
	DeviceConfigurationConflictSummaryId string
}

// NewDeviceManagementDeviceConfigurationConflictSummaryID returns a new DeviceManagementDeviceConfigurationConflictSummaryId struct
func NewDeviceManagementDeviceConfigurationConflictSummaryID(deviceConfigurationConflictSummaryId string) DeviceManagementDeviceConfigurationConflictSummaryId {
	return DeviceManagementDeviceConfigurationConflictSummaryId{
		DeviceConfigurationConflictSummaryId: deviceConfigurationConflictSummaryId,
	}
}

// ParseDeviceManagementDeviceConfigurationConflictSummaryID parses 'input' into a DeviceManagementDeviceConfigurationConflictSummaryId
func ParseDeviceManagementDeviceConfigurationConflictSummaryID(input string) (*DeviceManagementDeviceConfigurationConflictSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationConflictSummaryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationConflictSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationConflictSummaryIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationConflictSummaryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationConflictSummaryIDInsensitively(input string) (*DeviceManagementDeviceConfigurationConflictSummaryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationConflictSummaryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationConflictSummaryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationConflictSummaryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationConflictSummaryId, ok = input.Parsed["deviceConfigurationConflictSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationConflictSummaryId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationConflictSummaryID checks that 'input' can be parsed as a Device Management Device Configuration Conflict Summary ID
func ValidateDeviceManagementDeviceConfigurationConflictSummaryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationConflictSummaryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Conflict Summary ID
func (id DeviceManagementDeviceConfigurationConflictSummaryId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurationConflictSummary/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationConflictSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Conflict Summary ID
func (id DeviceManagementDeviceConfigurationConflictSummaryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurationConflictSummary", "deviceConfigurationConflictSummary", "deviceConfigurationConflictSummary"),
		resourceids.UserSpecifiedSegment("deviceConfigurationConflictSummaryId", "deviceConfigurationConflictSummaryId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Conflict Summary ID
func (id DeviceManagementDeviceConfigurationConflictSummaryId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration Conflict Summary: %q", id.DeviceConfigurationConflictSummaryId),
	}
	return fmt.Sprintf("Device Management Device Configuration Conflict Summary (%s)", strings.Join(components, "\n"))
}
