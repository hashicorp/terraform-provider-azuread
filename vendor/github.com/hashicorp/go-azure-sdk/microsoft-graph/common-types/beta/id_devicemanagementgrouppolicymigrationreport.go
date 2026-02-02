package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyMigrationReportId{}

// DeviceManagementGroupPolicyMigrationReportId is a struct representing the Resource ID for a Device Management Group Policy Migration Report
type DeviceManagementGroupPolicyMigrationReportId struct {
	GroupPolicyMigrationReportId string
}

// NewDeviceManagementGroupPolicyMigrationReportID returns a new DeviceManagementGroupPolicyMigrationReportId struct
func NewDeviceManagementGroupPolicyMigrationReportID(groupPolicyMigrationReportId string) DeviceManagementGroupPolicyMigrationReportId {
	return DeviceManagementGroupPolicyMigrationReportId{
		GroupPolicyMigrationReportId: groupPolicyMigrationReportId,
	}
}

// ParseDeviceManagementGroupPolicyMigrationReportID parses 'input' into a DeviceManagementGroupPolicyMigrationReportId
func ParseDeviceManagementGroupPolicyMigrationReportID(input string) (*DeviceManagementGroupPolicyMigrationReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyMigrationReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyMigrationReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyMigrationReportIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyMigrationReportId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyMigrationReportIDInsensitively(input string) (*DeviceManagementGroupPolicyMigrationReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyMigrationReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyMigrationReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyMigrationReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyMigrationReportId, ok = input.Parsed["groupPolicyMigrationReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyMigrationReportId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyMigrationReportID checks that 'input' can be parsed as a Device Management Group Policy Migration Report ID
func ValidateDeviceManagementGroupPolicyMigrationReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyMigrationReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Migration Report ID
func (id DeviceManagementGroupPolicyMigrationReportId) ID() string {
	fmtString := "/deviceManagement/groupPolicyMigrationReports/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyMigrationReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Migration Report ID
func (id DeviceManagementGroupPolicyMigrationReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyMigrationReports", "groupPolicyMigrationReports", "groupPolicyMigrationReports"),
		resourceids.UserSpecifiedSegment("groupPolicyMigrationReportId", "groupPolicyMigrationReportId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Migration Report ID
func (id DeviceManagementGroupPolicyMigrationReportId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Migration Report: %q", id.GroupPolicyMigrationReportId),
	}
	return fmt.Sprintf("Device Management Group Policy Migration Report (%s)", strings.Join(components, "\n"))
}
