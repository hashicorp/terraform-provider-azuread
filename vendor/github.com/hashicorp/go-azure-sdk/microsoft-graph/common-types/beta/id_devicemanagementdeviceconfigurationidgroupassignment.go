package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdGroupAssignmentId{}

// DeviceManagementDeviceConfigurationIdGroupAssignmentId is a struct representing the Resource ID for a Device Management Device Configuration Id Group Assignment
type DeviceManagementDeviceConfigurationIdGroupAssignmentId struct {
	DeviceConfigurationId                string
	DeviceConfigurationGroupAssignmentId string
}

// NewDeviceManagementDeviceConfigurationIdGroupAssignmentID returns a new DeviceManagementDeviceConfigurationIdGroupAssignmentId struct
func NewDeviceManagementDeviceConfigurationIdGroupAssignmentID(deviceConfigurationId string, deviceConfigurationGroupAssignmentId string) DeviceManagementDeviceConfigurationIdGroupAssignmentId {
	return DeviceManagementDeviceConfigurationIdGroupAssignmentId{
		DeviceConfigurationId:                deviceConfigurationId,
		DeviceConfigurationGroupAssignmentId: deviceConfigurationGroupAssignmentId,
	}
}

// ParseDeviceManagementDeviceConfigurationIdGroupAssignmentID parses 'input' into a DeviceManagementDeviceConfigurationIdGroupAssignmentId
func ParseDeviceManagementDeviceConfigurationIdGroupAssignmentID(input string) (*DeviceManagementDeviceConfigurationIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationIdGroupAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationIdGroupAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationIdGroupAssignmentIDInsensitively(input string) (*DeviceManagementDeviceConfigurationIdGroupAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdGroupAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdGroupAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationIdGroupAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationId, ok = input.Parsed["deviceConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationId", input)
	}

	if id.DeviceConfigurationGroupAssignmentId, ok = input.Parsed["deviceConfigurationGroupAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationGroupAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationIdGroupAssignmentID checks that 'input' can be parsed as a Device Management Device Configuration Id Group Assignment ID
func ValidateDeviceManagementDeviceConfigurationIdGroupAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationIdGroupAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Id Group Assignment ID
func (id DeviceManagementDeviceConfigurationIdGroupAssignmentId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurations/%s/groupAssignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationId, id.DeviceConfigurationGroupAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Id Group Assignment ID
func (id DeviceManagementDeviceConfigurationIdGroupAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurations", "deviceConfigurations", "deviceConfigurations"),
		resourceids.UserSpecifiedSegment("deviceConfigurationId", "deviceConfigurationId"),
		resourceids.StaticSegment("groupAssignments", "groupAssignments", "groupAssignments"),
		resourceids.UserSpecifiedSegment("deviceConfigurationGroupAssignmentId", "deviceConfigurationGroupAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Id Group Assignment ID
func (id DeviceManagementDeviceConfigurationIdGroupAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration: %q", id.DeviceConfigurationId),
		fmt.Sprintf("Device Configuration Group Assignment: %q", id.DeviceConfigurationGroupAssignmentId),
	}
	return fmt.Sprintf("Device Management Device Configuration Id Group Assignment (%s)", strings.Join(components, "\n"))
}
