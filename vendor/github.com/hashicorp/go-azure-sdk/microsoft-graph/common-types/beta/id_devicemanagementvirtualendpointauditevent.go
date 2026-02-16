package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointAuditEventId{}

// DeviceManagementVirtualEndpointAuditEventId is a struct representing the Resource ID for a Device Management Virtual Endpoint Audit Event
type DeviceManagementVirtualEndpointAuditEventId struct {
	CloudPCAuditEventId string
}

// NewDeviceManagementVirtualEndpointAuditEventID returns a new DeviceManagementVirtualEndpointAuditEventId struct
func NewDeviceManagementVirtualEndpointAuditEventID(cloudPCAuditEventId string) DeviceManagementVirtualEndpointAuditEventId {
	return DeviceManagementVirtualEndpointAuditEventId{
		CloudPCAuditEventId: cloudPCAuditEventId,
	}
}

// ParseDeviceManagementVirtualEndpointAuditEventID parses 'input' into a DeviceManagementVirtualEndpointAuditEventId
func ParseDeviceManagementVirtualEndpointAuditEventID(input string) (*DeviceManagementVirtualEndpointAuditEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointAuditEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointAuditEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointAuditEventIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointAuditEventId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointAuditEventIDInsensitively(input string) (*DeviceManagementVirtualEndpointAuditEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointAuditEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointAuditEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointAuditEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCAuditEventId, ok = input.Parsed["cloudPCAuditEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCAuditEventId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointAuditEventID checks that 'input' can be parsed as a Device Management Virtual Endpoint Audit Event ID
func ValidateDeviceManagementVirtualEndpointAuditEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointAuditEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Audit Event ID
func (id DeviceManagementVirtualEndpointAuditEventId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/auditEvents/%s"
	return fmt.Sprintf(fmtString, id.CloudPCAuditEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Audit Event ID
func (id DeviceManagementVirtualEndpointAuditEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("auditEvents", "auditEvents", "auditEvents"),
		resourceids.UserSpecifiedSegment("cloudPCAuditEventId", "cloudPCAuditEventId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Audit Event ID
func (id DeviceManagementVirtualEndpointAuditEventId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Audit Event: %q", id.CloudPCAuditEventId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Audit Event (%s)", strings.Join(components, "\n"))
}
