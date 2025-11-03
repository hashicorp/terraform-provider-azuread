package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointServicePlanId{}

// DeviceManagementVirtualEndpointServicePlanId is a struct representing the Resource ID for a Device Management Virtual Endpoint Service Plan
type DeviceManagementVirtualEndpointServicePlanId struct {
	CloudPCServicePlanId string
}

// NewDeviceManagementVirtualEndpointServicePlanID returns a new DeviceManagementVirtualEndpointServicePlanId struct
func NewDeviceManagementVirtualEndpointServicePlanID(cloudPCServicePlanId string) DeviceManagementVirtualEndpointServicePlanId {
	return DeviceManagementVirtualEndpointServicePlanId{
		CloudPCServicePlanId: cloudPCServicePlanId,
	}
}

// ParseDeviceManagementVirtualEndpointServicePlanID parses 'input' into a DeviceManagementVirtualEndpointServicePlanId
func ParseDeviceManagementVirtualEndpointServicePlanID(input string) (*DeviceManagementVirtualEndpointServicePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointServicePlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointServicePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointServicePlanIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointServicePlanId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointServicePlanIDInsensitively(input string) (*DeviceManagementVirtualEndpointServicePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointServicePlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointServicePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointServicePlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCServicePlanId, ok = input.Parsed["cloudPCServicePlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCServicePlanId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointServicePlanID checks that 'input' can be parsed as a Device Management Virtual Endpoint Service Plan ID
func ValidateDeviceManagementVirtualEndpointServicePlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointServicePlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Service Plan ID
func (id DeviceManagementVirtualEndpointServicePlanId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/servicePlans/%s"
	return fmt.Sprintf(fmtString, id.CloudPCServicePlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Service Plan ID
func (id DeviceManagementVirtualEndpointServicePlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("servicePlans", "servicePlans", "servicePlans"),
		resourceids.UserSpecifiedSegment("cloudPCServicePlanId", "cloudPCServicePlanId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Service Plan ID
func (id DeviceManagementVirtualEndpointServicePlanId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Service Plan: %q", id.CloudPCServicePlanId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Service Plan (%s)", strings.Join(components, "\n"))
}
