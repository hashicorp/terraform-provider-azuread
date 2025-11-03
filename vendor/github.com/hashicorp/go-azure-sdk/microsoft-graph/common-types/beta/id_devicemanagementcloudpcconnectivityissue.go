package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudPCConnectivityIssueId{}

// DeviceManagementCloudPCConnectivityIssueId is a struct representing the Resource ID for a Device Management Cloud PC Connectivity Issue
type DeviceManagementCloudPCConnectivityIssueId struct {
	CloudPCConnectivityIssueId string
}

// NewDeviceManagementCloudPCConnectivityIssueID returns a new DeviceManagementCloudPCConnectivityIssueId struct
func NewDeviceManagementCloudPCConnectivityIssueID(cloudPCConnectivityIssueId string) DeviceManagementCloudPCConnectivityIssueId {
	return DeviceManagementCloudPCConnectivityIssueId{
		CloudPCConnectivityIssueId: cloudPCConnectivityIssueId,
	}
}

// ParseDeviceManagementCloudPCConnectivityIssueID parses 'input' into a DeviceManagementCloudPCConnectivityIssueId
func ParseDeviceManagementCloudPCConnectivityIssueID(input string) (*DeviceManagementCloudPCConnectivityIssueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudPCConnectivityIssueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudPCConnectivityIssueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCloudPCConnectivityIssueIDInsensitively parses 'input' case-insensitively into a DeviceManagementCloudPCConnectivityIssueId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCloudPCConnectivityIssueIDInsensitively(input string) (*DeviceManagementCloudPCConnectivityIssueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCloudPCConnectivityIssueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCloudPCConnectivityIssueId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCloudPCConnectivityIssueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCConnectivityIssueId, ok = input.Parsed["cloudPCConnectivityIssueId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCConnectivityIssueId", input)
	}

	return nil
}

// ValidateDeviceManagementCloudPCConnectivityIssueID checks that 'input' can be parsed as a Device Management Cloud PC Connectivity Issue ID
func ValidateDeviceManagementCloudPCConnectivityIssueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCloudPCConnectivityIssueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Cloud PC Connectivity Issue ID
func (id DeviceManagementCloudPCConnectivityIssueId) ID() string {
	fmtString := "/deviceManagement/cloudPCConnectivityIssues/%s"
	return fmt.Sprintf(fmtString, id.CloudPCConnectivityIssueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Cloud PC Connectivity Issue ID
func (id DeviceManagementCloudPCConnectivityIssueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("cloudPCConnectivityIssues", "cloudPCConnectivityIssues", "cloudPCConnectivityIssues"),
		resourceids.UserSpecifiedSegment("cloudPCConnectivityIssueId", "cloudPCConnectivityIssueId"),
	}
}

// String returns a human-readable description of this Device Management Cloud PC Connectivity Issue ID
func (id DeviceManagementCloudPCConnectivityIssueId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC Connectivity Issue: %q", id.CloudPCConnectivityIssueId),
	}
	return fmt.Sprintf("Device Management Cloud PC Connectivity Issue (%s)", strings.Join(components, "\n"))
}
