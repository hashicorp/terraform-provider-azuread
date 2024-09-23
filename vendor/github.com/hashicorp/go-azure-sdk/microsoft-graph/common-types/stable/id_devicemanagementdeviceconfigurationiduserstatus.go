package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdUserStatusId{}

// DeviceManagementDeviceConfigurationIdUserStatusId is a struct representing the Resource ID for a Device Management Device Configuration Id User Status
type DeviceManagementDeviceConfigurationIdUserStatusId struct {
	DeviceConfigurationId           string
	DeviceConfigurationUserStatusId string
}

// NewDeviceManagementDeviceConfigurationIdUserStatusID returns a new DeviceManagementDeviceConfigurationIdUserStatusId struct
func NewDeviceManagementDeviceConfigurationIdUserStatusID(deviceConfigurationId string, deviceConfigurationUserStatusId string) DeviceManagementDeviceConfigurationIdUserStatusId {
	return DeviceManagementDeviceConfigurationIdUserStatusId{
		DeviceConfigurationId:           deviceConfigurationId,
		DeviceConfigurationUserStatusId: deviceConfigurationUserStatusId,
	}
}

// ParseDeviceManagementDeviceConfigurationIdUserStatusID parses 'input' into a DeviceManagementDeviceConfigurationIdUserStatusId
func ParseDeviceManagementDeviceConfigurationIdUserStatusID(input string) (*DeviceManagementDeviceConfigurationIdUserStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdUserStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdUserStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceConfigurationIdUserStatusIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceConfigurationIdUserStatusId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceConfigurationIdUserStatusIDInsensitively(input string) (*DeviceManagementDeviceConfigurationIdUserStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceConfigurationIdUserStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceConfigurationIdUserStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceConfigurationIdUserStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceConfigurationId, ok = input.Parsed["deviceConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationId", input)
	}

	if id.DeviceConfigurationUserStatusId, ok = input.Parsed["deviceConfigurationUserStatusId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationUserStatusId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceConfigurationIdUserStatusID checks that 'input' can be parsed as a Device Management Device Configuration Id User Status ID
func ValidateDeviceManagementDeviceConfigurationIdUserStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceConfigurationIdUserStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Configuration Id User Status ID
func (id DeviceManagementDeviceConfigurationIdUserStatusId) ID() string {
	fmtString := "/deviceManagement/deviceConfigurations/%s/userStatuses/%s"
	return fmt.Sprintf(fmtString, id.DeviceConfigurationId, id.DeviceConfigurationUserStatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Configuration Id User Status ID
func (id DeviceManagementDeviceConfigurationIdUserStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceConfigurations", "deviceConfigurations", "deviceConfigurations"),
		resourceids.UserSpecifiedSegment("deviceConfigurationId", "deviceConfigurationId"),
		resourceids.StaticSegment("userStatuses", "userStatuses", "userStatuses"),
		resourceids.UserSpecifiedSegment("deviceConfigurationUserStatusId", "deviceConfigurationUserStatusId"),
	}
}

// String returns a human-readable description of this Device Management Device Configuration Id User Status ID
func (id DeviceManagementDeviceConfigurationIdUserStatusId) String() string {
	components := []string{
		fmt.Sprintf("Device Configuration: %q", id.DeviceConfigurationId),
		fmt.Sprintf("Device Configuration User Status: %q", id.DeviceConfigurationUserStatusId),
	}
	return fmt.Sprintf("Device Management Device Configuration Id User Status (%s)", strings.Join(components, "\n"))
}
