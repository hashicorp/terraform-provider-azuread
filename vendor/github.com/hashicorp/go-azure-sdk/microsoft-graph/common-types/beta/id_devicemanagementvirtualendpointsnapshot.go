package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointSnapshotId{}

// DeviceManagementVirtualEndpointSnapshotId is a struct representing the Resource ID for a Device Management Virtual Endpoint Snapshot
type DeviceManagementVirtualEndpointSnapshotId struct {
	CloudPCSnapshotId string
}

// NewDeviceManagementVirtualEndpointSnapshotID returns a new DeviceManagementVirtualEndpointSnapshotId struct
func NewDeviceManagementVirtualEndpointSnapshotID(cloudPCSnapshotId string) DeviceManagementVirtualEndpointSnapshotId {
	return DeviceManagementVirtualEndpointSnapshotId{
		CloudPCSnapshotId: cloudPCSnapshotId,
	}
}

// ParseDeviceManagementVirtualEndpointSnapshotID parses 'input' into a DeviceManagementVirtualEndpointSnapshotId
func ParseDeviceManagementVirtualEndpointSnapshotID(input string) (*DeviceManagementVirtualEndpointSnapshotId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointSnapshotId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointSnapshotId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointSnapshotIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointSnapshotId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointSnapshotIDInsensitively(input string) (*DeviceManagementVirtualEndpointSnapshotId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointSnapshotId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointSnapshotId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointSnapshotId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCSnapshotId, ok = input.Parsed["cloudPCSnapshotId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCSnapshotId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointSnapshotID checks that 'input' can be parsed as a Device Management Virtual Endpoint Snapshot ID
func ValidateDeviceManagementVirtualEndpointSnapshotID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointSnapshotID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Snapshot ID
func (id DeviceManagementVirtualEndpointSnapshotId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/snapshots/%s"
	return fmt.Sprintf(fmtString, id.CloudPCSnapshotId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Snapshot ID
func (id DeviceManagementVirtualEndpointSnapshotId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("snapshots", "snapshots", "snapshots"),
		resourceids.UserSpecifiedSegment("cloudPCSnapshotId", "cloudPCSnapshotId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Snapshot ID
func (id DeviceManagementVirtualEndpointSnapshotId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Snapshot: %q", id.CloudPCSnapshotId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Snapshot (%s)", strings.Join(components, "\n"))
}
