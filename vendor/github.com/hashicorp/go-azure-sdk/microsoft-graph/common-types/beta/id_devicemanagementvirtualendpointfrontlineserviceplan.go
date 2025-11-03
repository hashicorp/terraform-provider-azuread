package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointFrontLineServicePlanId{}

// DeviceManagementVirtualEndpointFrontLineServicePlanId is a struct representing the Resource ID for a Device Management Virtual Endpoint Front Line Service Plan
type DeviceManagementVirtualEndpointFrontLineServicePlanId struct {
	CloudPCFrontLineServicePlanId string
}

// NewDeviceManagementVirtualEndpointFrontLineServicePlanID returns a new DeviceManagementVirtualEndpointFrontLineServicePlanId struct
func NewDeviceManagementVirtualEndpointFrontLineServicePlanID(cloudPCFrontLineServicePlanId string) DeviceManagementVirtualEndpointFrontLineServicePlanId {
	return DeviceManagementVirtualEndpointFrontLineServicePlanId{
		CloudPCFrontLineServicePlanId: cloudPCFrontLineServicePlanId,
	}
}

// ParseDeviceManagementVirtualEndpointFrontLineServicePlanID parses 'input' into a DeviceManagementVirtualEndpointFrontLineServicePlanId
func ParseDeviceManagementVirtualEndpointFrontLineServicePlanID(input string) (*DeviceManagementVirtualEndpointFrontLineServicePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointFrontLineServicePlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointFrontLineServicePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointFrontLineServicePlanIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointFrontLineServicePlanId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointFrontLineServicePlanIDInsensitively(input string) (*DeviceManagementVirtualEndpointFrontLineServicePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointFrontLineServicePlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointFrontLineServicePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointFrontLineServicePlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCFrontLineServicePlanId, ok = input.Parsed["cloudPCFrontLineServicePlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCFrontLineServicePlanId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointFrontLineServicePlanID checks that 'input' can be parsed as a Device Management Virtual Endpoint Front Line Service Plan ID
func ValidateDeviceManagementVirtualEndpointFrontLineServicePlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointFrontLineServicePlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Front Line Service Plan ID
func (id DeviceManagementVirtualEndpointFrontLineServicePlanId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/frontLineServicePlans/%s"
	return fmt.Sprintf(fmtString, id.CloudPCFrontLineServicePlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Front Line Service Plan ID
func (id DeviceManagementVirtualEndpointFrontLineServicePlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("frontLineServicePlans", "frontLineServicePlans", "frontLineServicePlans"),
		resourceids.UserSpecifiedSegment("cloudPCFrontLineServicePlanId", "cloudPCFrontLineServicePlanId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Front Line Service Plan ID
func (id DeviceManagementVirtualEndpointFrontLineServicePlanId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Front Line Service Plan: %q", id.CloudPCFrontLineServicePlanId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Front Line Service Plan (%s)", strings.Join(components, "\n"))
}
