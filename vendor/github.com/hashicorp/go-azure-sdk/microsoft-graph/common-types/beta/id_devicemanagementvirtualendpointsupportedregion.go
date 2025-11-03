package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointSupportedRegionId{}

// DeviceManagementVirtualEndpointSupportedRegionId is a struct representing the Resource ID for a Device Management Virtual Endpoint Supported Region
type DeviceManagementVirtualEndpointSupportedRegionId struct {
	CloudPCSupportedRegionId string
}

// NewDeviceManagementVirtualEndpointSupportedRegionID returns a new DeviceManagementVirtualEndpointSupportedRegionId struct
func NewDeviceManagementVirtualEndpointSupportedRegionID(cloudPCSupportedRegionId string) DeviceManagementVirtualEndpointSupportedRegionId {
	return DeviceManagementVirtualEndpointSupportedRegionId{
		CloudPCSupportedRegionId: cloudPCSupportedRegionId,
	}
}

// ParseDeviceManagementVirtualEndpointSupportedRegionID parses 'input' into a DeviceManagementVirtualEndpointSupportedRegionId
func ParseDeviceManagementVirtualEndpointSupportedRegionID(input string) (*DeviceManagementVirtualEndpointSupportedRegionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointSupportedRegionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointSupportedRegionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointSupportedRegionIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointSupportedRegionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointSupportedRegionIDInsensitively(input string) (*DeviceManagementVirtualEndpointSupportedRegionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointSupportedRegionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointSupportedRegionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointSupportedRegionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCSupportedRegionId, ok = input.Parsed["cloudPCSupportedRegionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCSupportedRegionId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointSupportedRegionID checks that 'input' can be parsed as a Device Management Virtual Endpoint Supported Region ID
func ValidateDeviceManagementVirtualEndpointSupportedRegionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointSupportedRegionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Supported Region ID
func (id DeviceManagementVirtualEndpointSupportedRegionId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/supportedRegions/%s"
	return fmt.Sprintf(fmtString, id.CloudPCSupportedRegionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Supported Region ID
func (id DeviceManagementVirtualEndpointSupportedRegionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("supportedRegions", "supportedRegions", "supportedRegions"),
		resourceids.UserSpecifiedSegment("cloudPCSupportedRegionId", "cloudPCSupportedRegionId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Supported Region ID
func (id DeviceManagementVirtualEndpointSupportedRegionId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Supported Region: %q", id.CloudPCSupportedRegionId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Supported Region (%s)", strings.Join(components, "\n"))
}
