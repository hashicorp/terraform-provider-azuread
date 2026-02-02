package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{}

// DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId is a struct representing the Resource ID for a Device Management Group Policy Migration Report Id Group Policy Setting Mapping
type DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId struct {
	GroupPolicyMigrationReportId string
	GroupPolicySettingMappingId  string
}

// NewDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID returns a new DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId struct
func NewDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(groupPolicyMigrationReportId string, groupPolicySettingMappingId string) DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId {
	return DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{
		GroupPolicyMigrationReportId: groupPolicyMigrationReportId,
		GroupPolicySettingMappingId:  groupPolicySettingMappingId,
	}
}

// ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID parses 'input' into a DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId
func ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(input string) (*DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingIDInsensitively(input string) (*DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyMigrationReportId, ok = input.Parsed["groupPolicyMigrationReportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyMigrationReportId", input)
	}

	if id.GroupPolicySettingMappingId, ok = input.Parsed["groupPolicySettingMappingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicySettingMappingId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID checks that 'input' can be parsed as a Device Management Group Policy Migration Report Id Group Policy Setting Mapping ID
func ValidateDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Migration Report Id Group Policy Setting Mapping ID
func (id DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId) ID() string {
	fmtString := "/deviceManagement/groupPolicyMigrationReports/%s/groupPolicySettingMappings/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyMigrationReportId, id.GroupPolicySettingMappingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Migration Report Id Group Policy Setting Mapping ID
func (id DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyMigrationReports", "groupPolicyMigrationReports", "groupPolicyMigrationReports"),
		resourceids.UserSpecifiedSegment("groupPolicyMigrationReportId", "groupPolicyMigrationReportId"),
		resourceids.StaticSegment("groupPolicySettingMappings", "groupPolicySettingMappings", "groupPolicySettingMappings"),
		resourceids.UserSpecifiedSegment("groupPolicySettingMappingId", "groupPolicySettingMappingId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Migration Report Id Group Policy Setting Mapping ID
func (id DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Migration Report: %q", id.GroupPolicyMigrationReportId),
		fmt.Sprintf("Group Policy Setting Mapping: %q", id.GroupPolicySettingMappingId),
	}
	return fmt.Sprintf("Device Management Group Policy Migration Report Id Group Policy Setting Mapping (%s)", strings.Join(components, "\n"))
}
