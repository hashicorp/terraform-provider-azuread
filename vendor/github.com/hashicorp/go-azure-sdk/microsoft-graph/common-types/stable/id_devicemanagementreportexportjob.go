package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReportExportJobId{}

// DeviceManagementReportExportJobId is a struct representing the Resource ID for a Device Management Report Export Job
type DeviceManagementReportExportJobId struct {
	DeviceManagementExportJobId string
}

// NewDeviceManagementReportExportJobID returns a new DeviceManagementReportExportJobId struct
func NewDeviceManagementReportExportJobID(deviceManagementExportJobId string) DeviceManagementReportExportJobId {
	return DeviceManagementReportExportJobId{
		DeviceManagementExportJobId: deviceManagementExportJobId,
	}
}

// ParseDeviceManagementReportExportJobID parses 'input' into a DeviceManagementReportExportJobId
func ParseDeviceManagementReportExportJobID(input string) (*DeviceManagementReportExportJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReportExportJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReportExportJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReportExportJobIDInsensitively parses 'input' case-insensitively into a DeviceManagementReportExportJobId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReportExportJobIDInsensitively(input string) (*DeviceManagementReportExportJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReportExportJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReportExportJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReportExportJobId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementExportJobId, ok = input.Parsed["deviceManagementExportJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementExportJobId", input)
	}

	return nil
}

// ValidateDeviceManagementReportExportJobID checks that 'input' can be parsed as a Device Management Report Export Job ID
func ValidateDeviceManagementReportExportJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReportExportJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Report Export Job ID
func (id DeviceManagementReportExportJobId) ID() string {
	fmtString := "/deviceManagement/reports/exportJobs/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementExportJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Report Export Job ID
func (id DeviceManagementReportExportJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("exportJobs", "exportJobs", "exportJobs"),
		resourceids.UserSpecifiedSegment("deviceManagementExportJobId", "deviceManagementExportJobId"),
	}
}

// String returns a human-readable description of this Device Management Report Export Job ID
func (id DeviceManagementReportExportJobId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Export Job: %q", id.DeviceManagementExportJobId),
	}
	return fmt.Sprintf("Device Management Report Export Job (%s)", strings.Join(components, "\n"))
}
