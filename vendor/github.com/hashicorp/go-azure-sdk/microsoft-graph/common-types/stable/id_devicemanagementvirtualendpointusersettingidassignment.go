package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointUserSettingIdAssignmentId{}

// DeviceManagementVirtualEndpointUserSettingIdAssignmentId is a struct representing the Resource ID for a Device Management Virtual Endpoint User Setting Id Assignment
type DeviceManagementVirtualEndpointUserSettingIdAssignmentId struct {
	CloudPCUserSettingId           string
	CloudPCUserSettingAssignmentId string
}

// NewDeviceManagementVirtualEndpointUserSettingIdAssignmentID returns a new DeviceManagementVirtualEndpointUserSettingIdAssignmentId struct
func NewDeviceManagementVirtualEndpointUserSettingIdAssignmentID(cloudPCUserSettingId string, cloudPCUserSettingAssignmentId string) DeviceManagementVirtualEndpointUserSettingIdAssignmentId {
	return DeviceManagementVirtualEndpointUserSettingIdAssignmentId{
		CloudPCUserSettingId:           cloudPCUserSettingId,
		CloudPCUserSettingAssignmentId: cloudPCUserSettingAssignmentId,
	}
}

// ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentID parses 'input' into a DeviceManagementVirtualEndpointUserSettingIdAssignmentId
func ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentID(input string) (*DeviceManagementVirtualEndpointUserSettingIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointUserSettingIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointUserSettingIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointUserSettingIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentIDInsensitively(input string) (*DeviceManagementVirtualEndpointUserSettingIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointUserSettingIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointUserSettingIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointUserSettingIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCUserSettingId, ok = input.Parsed["cloudPCUserSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCUserSettingId", input)
	}

	if id.CloudPCUserSettingAssignmentId, ok = input.Parsed["cloudPCUserSettingAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCUserSettingAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointUserSettingIdAssignmentID checks that 'input' can be parsed as a Device Management Virtual Endpoint User Setting Id Assignment ID
func ValidateDeviceManagementVirtualEndpointUserSettingIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointUserSettingIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint User Setting Id Assignment ID
func (id DeviceManagementVirtualEndpointUserSettingIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/userSettings/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.CloudPCUserSettingId, id.CloudPCUserSettingAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint User Setting Id Assignment ID
func (id DeviceManagementVirtualEndpointUserSettingIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("userSettings", "userSettings", "userSettings"),
		resourceids.UserSpecifiedSegment("cloudPCUserSettingId", "cloudPCUserSettingId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("cloudPCUserSettingAssignmentId", "cloudPCUserSettingAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint User Setting Id Assignment ID
func (id DeviceManagementVirtualEndpointUserSettingIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC User Setting: %q", id.CloudPCUserSettingId),
		fmt.Sprintf("Cloud PC User Setting Assignment: %q", id.CloudPCUserSettingAssignmentId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint User Setting Id Assignment (%s)", strings.Join(components, "\n"))
}
