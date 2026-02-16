package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointCloudPCId{}

// DeviceManagementVirtualEndpointCloudPCId is a struct representing the Resource ID for a Device Management Virtual Endpoint Cloud PC
type DeviceManagementVirtualEndpointCloudPCId struct {
	CloudPCId string
}

// NewDeviceManagementVirtualEndpointCloudPCID returns a new DeviceManagementVirtualEndpointCloudPCId struct
func NewDeviceManagementVirtualEndpointCloudPCID(cloudPCId string) DeviceManagementVirtualEndpointCloudPCId {
	return DeviceManagementVirtualEndpointCloudPCId{
		CloudPCId: cloudPCId,
	}
}

// ParseDeviceManagementVirtualEndpointCloudPCID parses 'input' into a DeviceManagementVirtualEndpointCloudPCId
func ParseDeviceManagementVirtualEndpointCloudPCID(input string) (*DeviceManagementVirtualEndpointCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointCloudPCId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointCloudPCId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointCloudPCIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointCloudPCId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointCloudPCIDInsensitively(input string) (*DeviceManagementVirtualEndpointCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointCloudPCId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointCloudPCId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointCloudPCId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCId, ok = input.Parsed["cloudPCId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointCloudPCID checks that 'input' can be parsed as a Device Management Virtual Endpoint Cloud PC ID
func ValidateDeviceManagementVirtualEndpointCloudPCID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointCloudPCID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Cloud PC ID
func (id DeviceManagementVirtualEndpointCloudPCId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/cloudPCs/%s"
	return fmt.Sprintf(fmtString, id.CloudPCId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Cloud PC ID
func (id DeviceManagementVirtualEndpointCloudPCId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("cloudPCs", "cloudPCs", "cloudPCs"),
		resourceids.UserSpecifiedSegment("cloudPCId", "cloudPCId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Cloud PC ID
func (id DeviceManagementVirtualEndpointCloudPCId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC: %q", id.CloudPCId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Cloud PC (%s)", strings.Join(components, "\n"))
}
