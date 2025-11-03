package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointBulkActionId{}

// DeviceManagementVirtualEndpointBulkActionId is a struct representing the Resource ID for a Device Management Virtual Endpoint Bulk Action
type DeviceManagementVirtualEndpointBulkActionId struct {
	CloudPCBulkActionId string
}

// NewDeviceManagementVirtualEndpointBulkActionID returns a new DeviceManagementVirtualEndpointBulkActionId struct
func NewDeviceManagementVirtualEndpointBulkActionID(cloudPCBulkActionId string) DeviceManagementVirtualEndpointBulkActionId {
	return DeviceManagementVirtualEndpointBulkActionId{
		CloudPCBulkActionId: cloudPCBulkActionId,
	}
}

// ParseDeviceManagementVirtualEndpointBulkActionID parses 'input' into a DeviceManagementVirtualEndpointBulkActionId
func ParseDeviceManagementVirtualEndpointBulkActionID(input string) (*DeviceManagementVirtualEndpointBulkActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointBulkActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointBulkActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointBulkActionIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointBulkActionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointBulkActionIDInsensitively(input string) (*DeviceManagementVirtualEndpointBulkActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointBulkActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointBulkActionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointBulkActionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCBulkActionId, ok = input.Parsed["cloudPCBulkActionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCBulkActionId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointBulkActionID checks that 'input' can be parsed as a Device Management Virtual Endpoint Bulk Action ID
func ValidateDeviceManagementVirtualEndpointBulkActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointBulkActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Bulk Action ID
func (id DeviceManagementVirtualEndpointBulkActionId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/bulkActions/%s"
	return fmt.Sprintf(fmtString, id.CloudPCBulkActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Bulk Action ID
func (id DeviceManagementVirtualEndpointBulkActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("bulkActions", "bulkActions", "bulkActions"),
		resourceids.UserSpecifiedSegment("cloudPCBulkActionId", "cloudPCBulkActionId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Bulk Action ID
func (id DeviceManagementVirtualEndpointBulkActionId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Bulk Action: %q", id.CloudPCBulkActionId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Bulk Action (%s)", strings.Join(components, "\n"))
}
