package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointDeviceImageId{}

// DeviceManagementVirtualEndpointDeviceImageId is a struct representing the Resource ID for a Device Management Virtual Endpoint Device Image
type DeviceManagementVirtualEndpointDeviceImageId struct {
	CloudPCDeviceImageId string
}

// NewDeviceManagementVirtualEndpointDeviceImageID returns a new DeviceManagementVirtualEndpointDeviceImageId struct
func NewDeviceManagementVirtualEndpointDeviceImageID(cloudPCDeviceImageId string) DeviceManagementVirtualEndpointDeviceImageId {
	return DeviceManagementVirtualEndpointDeviceImageId{
		CloudPCDeviceImageId: cloudPCDeviceImageId,
	}
}

// ParseDeviceManagementVirtualEndpointDeviceImageID parses 'input' into a DeviceManagementVirtualEndpointDeviceImageId
func ParseDeviceManagementVirtualEndpointDeviceImageID(input string) (*DeviceManagementVirtualEndpointDeviceImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointDeviceImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointDeviceImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointDeviceImageIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointDeviceImageId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointDeviceImageIDInsensitively(input string) (*DeviceManagementVirtualEndpointDeviceImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointDeviceImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointDeviceImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointDeviceImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCDeviceImageId, ok = input.Parsed["cloudPCDeviceImageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCDeviceImageId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointDeviceImageID checks that 'input' can be parsed as a Device Management Virtual Endpoint Device Image ID
func ValidateDeviceManagementVirtualEndpointDeviceImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointDeviceImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Device Image ID
func (id DeviceManagementVirtualEndpointDeviceImageId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/deviceImages/%s"
	return fmt.Sprintf(fmtString, id.CloudPCDeviceImageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Device Image ID
func (id DeviceManagementVirtualEndpointDeviceImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("deviceImages", "deviceImages", "deviceImages"),
		resourceids.UserSpecifiedSegment("cloudPCDeviceImageId", "cloudPCDeviceImageId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Device Image ID
func (id DeviceManagementVirtualEndpointDeviceImageId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Device Image: %q", id.CloudPCDeviceImageId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Device Image (%s)", strings.Join(components, "\n"))
}
