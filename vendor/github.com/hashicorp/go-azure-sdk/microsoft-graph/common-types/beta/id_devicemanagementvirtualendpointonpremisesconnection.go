package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointOnPremisesConnectionId{}

// DeviceManagementVirtualEndpointOnPremisesConnectionId is a struct representing the Resource ID for a Device Management Virtual Endpoint On Premises Connection
type DeviceManagementVirtualEndpointOnPremisesConnectionId struct {
	CloudPCOnPremisesConnectionId string
}

// NewDeviceManagementVirtualEndpointOnPremisesConnectionID returns a new DeviceManagementVirtualEndpointOnPremisesConnectionId struct
func NewDeviceManagementVirtualEndpointOnPremisesConnectionID(cloudPCOnPremisesConnectionId string) DeviceManagementVirtualEndpointOnPremisesConnectionId {
	return DeviceManagementVirtualEndpointOnPremisesConnectionId{
		CloudPCOnPremisesConnectionId: cloudPCOnPremisesConnectionId,
	}
}

// ParseDeviceManagementVirtualEndpointOnPremisesConnectionID parses 'input' into a DeviceManagementVirtualEndpointOnPremisesConnectionId
func ParseDeviceManagementVirtualEndpointOnPremisesConnectionID(input string) (*DeviceManagementVirtualEndpointOnPremisesConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointOnPremisesConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointOnPremisesConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointOnPremisesConnectionIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointOnPremisesConnectionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointOnPremisesConnectionIDInsensitively(input string) (*DeviceManagementVirtualEndpointOnPremisesConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointOnPremisesConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointOnPremisesConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointOnPremisesConnectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCOnPremisesConnectionId, ok = input.Parsed["cloudPCOnPremisesConnectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCOnPremisesConnectionId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointOnPremisesConnectionID checks that 'input' can be parsed as a Device Management Virtual Endpoint On Premises Connection ID
func ValidateDeviceManagementVirtualEndpointOnPremisesConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointOnPremisesConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint On Premises Connection ID
func (id DeviceManagementVirtualEndpointOnPremisesConnectionId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/onPremisesConnections/%s"
	return fmt.Sprintf(fmtString, id.CloudPCOnPremisesConnectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint On Premises Connection ID
func (id DeviceManagementVirtualEndpointOnPremisesConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("onPremisesConnections", "onPremisesConnections", "onPremisesConnections"),
		resourceids.UserSpecifiedSegment("cloudPCOnPremisesConnectionId", "cloudPCOnPremisesConnectionId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint On Premises Connection ID
func (id DeviceManagementVirtualEndpointOnPremisesConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC On Premises Connection: %q", id.CloudPCOnPremisesConnectionId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint On Premises Connection (%s)", strings.Join(components, "\n"))
}
