package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAssignmentFilterId{}

// DeviceManagementAssignmentFilterId is a struct representing the Resource ID for a Device Management Assignment Filter
type DeviceManagementAssignmentFilterId struct {
	DeviceAndAppManagementAssignmentFilterId string
}

// NewDeviceManagementAssignmentFilterID returns a new DeviceManagementAssignmentFilterId struct
func NewDeviceManagementAssignmentFilterID(deviceAndAppManagementAssignmentFilterId string) DeviceManagementAssignmentFilterId {
	return DeviceManagementAssignmentFilterId{
		DeviceAndAppManagementAssignmentFilterId: deviceAndAppManagementAssignmentFilterId,
	}
}

// ParseDeviceManagementAssignmentFilterID parses 'input' into a DeviceManagementAssignmentFilterId
func ParseDeviceManagementAssignmentFilterID(input string) (*DeviceManagementAssignmentFilterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAssignmentFilterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAssignmentFilterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAssignmentFilterIDInsensitively parses 'input' case-insensitively into a DeviceManagementAssignmentFilterId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAssignmentFilterIDInsensitively(input string) (*DeviceManagementAssignmentFilterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAssignmentFilterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAssignmentFilterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAssignmentFilterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceAndAppManagementAssignmentFilterId, ok = input.Parsed["deviceAndAppManagementAssignmentFilterId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceAndAppManagementAssignmentFilterId", input)
	}

	return nil
}

// ValidateDeviceManagementAssignmentFilterID checks that 'input' can be parsed as a Device Management Assignment Filter ID
func ValidateDeviceManagementAssignmentFilterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAssignmentFilterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Assignment Filter ID
func (id DeviceManagementAssignmentFilterId) ID() string {
	fmtString := "/deviceManagement/assignmentFilters/%s"
	return fmt.Sprintf(fmtString, id.DeviceAndAppManagementAssignmentFilterId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Assignment Filter ID
func (id DeviceManagementAssignmentFilterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("assignmentFilters", "assignmentFilters", "assignmentFilters"),
		resourceids.UserSpecifiedSegment("deviceAndAppManagementAssignmentFilterId", "deviceAndAppManagementAssignmentFilterId"),
	}
}

// String returns a human-readable description of this Device Management Assignment Filter ID
func (id DeviceManagementAssignmentFilterId) String() string {
	components := []string{
		fmt.Sprintf("Device And App Management Assignment Filter: %q", id.DeviceAndAppManagementAssignmentFilterId),
	}
	return fmt.Sprintf("Device Management Assignment Filter (%s)", strings.Join(components, "\n"))
}
