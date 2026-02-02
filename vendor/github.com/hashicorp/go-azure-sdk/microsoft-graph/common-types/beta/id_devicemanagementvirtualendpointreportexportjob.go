package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointReportExportJobId{}

// DeviceManagementVirtualEndpointReportExportJobId is a struct representing the Resource ID for a Device Management Virtual Endpoint Report Export Job
type DeviceManagementVirtualEndpointReportExportJobId struct {
	CloudPCExportJobId string
}

// NewDeviceManagementVirtualEndpointReportExportJobID returns a new DeviceManagementVirtualEndpointReportExportJobId struct
func NewDeviceManagementVirtualEndpointReportExportJobID(cloudPCExportJobId string) DeviceManagementVirtualEndpointReportExportJobId {
	return DeviceManagementVirtualEndpointReportExportJobId{
		CloudPCExportJobId: cloudPCExportJobId,
	}
}

// ParseDeviceManagementVirtualEndpointReportExportJobID parses 'input' into a DeviceManagementVirtualEndpointReportExportJobId
func ParseDeviceManagementVirtualEndpointReportExportJobID(input string) (*DeviceManagementVirtualEndpointReportExportJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointReportExportJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointReportExportJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointReportExportJobIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointReportExportJobId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointReportExportJobIDInsensitively(input string) (*DeviceManagementVirtualEndpointReportExportJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointReportExportJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointReportExportJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointReportExportJobId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCExportJobId, ok = input.Parsed["cloudPCExportJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCExportJobId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointReportExportJobID checks that 'input' can be parsed as a Device Management Virtual Endpoint Report Export Job ID
func ValidateDeviceManagementVirtualEndpointReportExportJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointReportExportJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Report Export Job ID
func (id DeviceManagementVirtualEndpointReportExportJobId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/reports/exportJobs/%s"
	return fmt.Sprintf(fmtString, id.CloudPCExportJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Report Export Job ID
func (id DeviceManagementVirtualEndpointReportExportJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("exportJobs", "exportJobs", "exportJobs"),
		resourceids.UserSpecifiedSegment("cloudPCExportJobId", "cloudPCExportJobId"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Report Export Job ID
func (id DeviceManagementVirtualEndpointReportExportJobId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Export Job: %q", id.CloudPCExportJobId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Report Export Job (%s)", strings.Join(components, "\n"))
}
