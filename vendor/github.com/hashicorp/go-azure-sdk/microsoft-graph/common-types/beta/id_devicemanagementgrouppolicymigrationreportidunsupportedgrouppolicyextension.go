package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{}

// DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId is a struct representing the Resource ID for a Device Management Group Policy Migration Report Id Unsupported Group Policy Extension
type DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId struct {
	GroupPolicyMigrationReportId      string
	UnsupportedGroupPolicyExtensionId string
}

// NewDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID returns a new DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId struct
func NewDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(groupPolicyMigrationReportId string, unsupportedGroupPolicyExtensionId string) DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId {
	return DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{
		GroupPolicyMigrationReportId:      groupPolicyMigrationReportId,
		UnsupportedGroupPolicyExtensionId: unsupportedGroupPolicyExtensionId,
	}
}

// ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID parses 'input' into a DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId
func ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(input string) (*DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionIDInsensitively(input string) (*DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyMigrationReportId, ok = input.Parsed["groupPolicyMigrationReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyMigrationReportId", input)
	}

	if id.UnsupportedGroupPolicyExtensionId, ok = input.Parsed["unsupportedGroupPolicyExtensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unsupportedGroupPolicyExtensionId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID checks that 'input' can be parsed as a Device Management Group Policy Migration Report Id Unsupported Group Policy Extension ID
func ValidateDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Migration Report Id Unsupported Group Policy Extension ID
func (id DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId) ID() string {
	fmtString := "/deviceManagement/groupPolicyMigrationReports/%s/unsupportedGroupPolicyExtensions/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyMigrationReportId, id.UnsupportedGroupPolicyExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Migration Report Id Unsupported Group Policy Extension ID
func (id DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyMigrationReports", "groupPolicyMigrationReports", "groupPolicyMigrationReports"),
		resourceids.UserSpecifiedSegment("groupPolicyMigrationReportId", "groupPolicyMigrationReportId"),
		resourceids.StaticSegment("unsupportedGroupPolicyExtensions", "unsupportedGroupPolicyExtensions", "unsupportedGroupPolicyExtensions"),
		resourceids.UserSpecifiedSegment("unsupportedGroupPolicyExtensionId", "unsupportedGroupPolicyExtensionId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Migration Report Id Unsupported Group Policy Extension ID
func (id DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Migration Report: %q", id.GroupPolicyMigrationReportId),
		fmt.Sprintf("Unsupported Group Policy Extension: %q", id.UnsupportedGroupPolicyExtensionId),
	}
	return fmt.Sprintf("Device Management Group Policy Migration Report Id Unsupported Group Policy Extension (%s)", strings.Join(components, "\n"))
}
