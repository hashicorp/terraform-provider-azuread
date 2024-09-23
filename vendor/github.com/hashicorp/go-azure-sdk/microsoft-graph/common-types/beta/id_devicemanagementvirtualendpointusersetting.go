package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointUserSettingId{}

// DeviceManagementVirtualEndpointUserSettingId is a struct representing the Resource ID for a Device Management Virtual Endpoint User Setting
type DeviceManagementVirtualEndpointUserSettingId struct {
	CloudPCUserSettingId string
}

// NewDeviceManagementVirtualEndpointUserSettingID returns a new DeviceManagementVirtualEndpointUserSettingId struct
func NewDeviceManagementVirtualEndpointUserSettingID(cloudPCUserSettingId string) DeviceManagementVirtualEndpointUserSettingId {
	return DeviceManagementVirtualEndpointUserSettingId{
		CloudPCUserSettingId: cloudPCUserSettingId,
	}
}

// ParseDeviceManagementVirtualEndpointUserSettingID parses 'input' into a DeviceManagementVirtualEndpointUserSettingId
func ParseDeviceManagementVirtualEndpointUserSettingID(input string) (*DeviceManagementVirtualEndpointUserSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointUserSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointUserSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointUserSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointUserSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointUserSettingIDInsensitively(input string) (*DeviceManagementVirtualEndpointUserSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointUserSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointUserSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointUserSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCUserSettingId, ok = input.Parsed["cloudPCUserSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCUserSettingId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointUserSettingID checks that 'input' can be parsed as a Device Management Virtual Endpoint User Setting ID
func ValidateDeviceManagementVirtualEndpointUserSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointUserSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint User Setting ID
func (id DeviceManagementVirtualEndpointUserSettingId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/userSettings/%s"
	return fmt.Sprintf(fmtString, id.CloudPCUserSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint User Setting ID
func (id DeviceManagementVirtualEndpointUserSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("userSettings", "userSettings", "userSettings"),
		resourceids.UserSpecifiedSegment("cloudPCUserSettingId", "cloudPCUserSettingId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint User Setting ID
func (id DeviceManagementVirtualEndpointUserSettingId) String() string {
	components := []string{
		fmt.Sprintf("Cloud P C User Setting: %q", id.CloudPCUserSettingId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint User Setting (%s)", strings.Join(components, "\n"))
}
