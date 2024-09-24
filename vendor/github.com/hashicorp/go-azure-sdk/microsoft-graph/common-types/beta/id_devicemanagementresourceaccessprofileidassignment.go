package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceAccessProfileIdAssignmentId{}

// DeviceManagementResourceAccessProfileIdAssignmentId is a struct representing the Resource ID for a Device Management Resource Access Profile Id Assignment
type DeviceManagementResourceAccessProfileIdAssignmentId struct {
	DeviceManagementResourceAccessProfileBaseId       string
	DeviceManagementResourceAccessProfileAssignmentId string
}

// NewDeviceManagementResourceAccessProfileIdAssignmentID returns a new DeviceManagementResourceAccessProfileIdAssignmentId struct
func NewDeviceManagementResourceAccessProfileIdAssignmentID(deviceManagementResourceAccessProfileBaseId string, deviceManagementResourceAccessProfileAssignmentId string) DeviceManagementResourceAccessProfileIdAssignmentId {
	return DeviceManagementResourceAccessProfileIdAssignmentId{
		DeviceManagementResourceAccessProfileBaseId:       deviceManagementResourceAccessProfileBaseId,
		DeviceManagementResourceAccessProfileAssignmentId: deviceManagementResourceAccessProfileAssignmentId,
	}
}

// ParseDeviceManagementResourceAccessProfileIdAssignmentID parses 'input' into a DeviceManagementResourceAccessProfileIdAssignmentId
func ParseDeviceManagementResourceAccessProfileIdAssignmentID(input string) (*DeviceManagementResourceAccessProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementResourceAccessProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementResourceAccessProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementResourceAccessProfileIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementResourceAccessProfileIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementResourceAccessProfileIdAssignmentIDInsensitively(input string) (*DeviceManagementResourceAccessProfileIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementResourceAccessProfileIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementResourceAccessProfileIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementResourceAccessProfileIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementResourceAccessProfileBaseId, ok = input.Parsed["deviceManagementResourceAccessProfileBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementResourceAccessProfileBaseId", input)
	}

	if id.DeviceManagementResourceAccessProfileAssignmentId, ok = input.Parsed["deviceManagementResourceAccessProfileAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementResourceAccessProfileAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementResourceAccessProfileIdAssignmentID checks that 'input' can be parsed as a Device Management Resource Access Profile Id Assignment ID
func ValidateDeviceManagementResourceAccessProfileIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementResourceAccessProfileIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Resource Access Profile Id Assignment ID
func (id DeviceManagementResourceAccessProfileIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/resourceAccessProfiles/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementResourceAccessProfileBaseId, id.DeviceManagementResourceAccessProfileAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Resource Access Profile Id Assignment ID
func (id DeviceManagementResourceAccessProfileIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("resourceAccessProfiles", "resourceAccessProfiles", "resourceAccessProfiles"),
		resourceids.UserSpecifiedSegment("deviceManagementResourceAccessProfileBaseId", "deviceManagementResourceAccessProfileBaseId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementResourceAccessProfileAssignmentId", "deviceManagementResourceAccessProfileAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Resource Access Profile Id Assignment ID
func (id DeviceManagementResourceAccessProfileIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Resource Access Profile Base: %q", id.DeviceManagementResourceAccessProfileBaseId),
		fmt.Sprintf("Device Management Resource Access Profile Assignment: %q", id.DeviceManagementResourceAccessProfileAssignmentId),
	}
	return fmt.Sprintf("Device Management Resource Access Profile Id Assignment (%s)", strings.Join(components, "\n"))
}
