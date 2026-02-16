package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReportCachedReportConfigurationId{}

// DeviceManagementReportCachedReportConfigurationId is a struct representing the Resource ID for a Device Management Report Cached Report Configuration
type DeviceManagementReportCachedReportConfigurationId struct {
	DeviceManagementCachedReportConfigurationId string
}

// NewDeviceManagementReportCachedReportConfigurationID returns a new DeviceManagementReportCachedReportConfigurationId struct
func NewDeviceManagementReportCachedReportConfigurationID(deviceManagementCachedReportConfigurationId string) DeviceManagementReportCachedReportConfigurationId {
	return DeviceManagementReportCachedReportConfigurationId{
		DeviceManagementCachedReportConfigurationId: deviceManagementCachedReportConfigurationId,
	}
}

// ParseDeviceManagementReportCachedReportConfigurationID parses 'input' into a DeviceManagementReportCachedReportConfigurationId
func ParseDeviceManagementReportCachedReportConfigurationID(input string) (*DeviceManagementReportCachedReportConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReportCachedReportConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReportCachedReportConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReportCachedReportConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementReportCachedReportConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReportCachedReportConfigurationIDInsensitively(input string) (*DeviceManagementReportCachedReportConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReportCachedReportConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReportCachedReportConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReportCachedReportConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCachedReportConfigurationId, ok = input.Parsed["deviceManagementCachedReportConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCachedReportConfigurationId", input)
	}

	return nil
}

// ValidateDeviceManagementReportCachedReportConfigurationID checks that 'input' can be parsed as a Device Management Report Cached Report Configuration ID
func ValidateDeviceManagementReportCachedReportConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReportCachedReportConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Report Cached Report Configuration ID
func (id DeviceManagementReportCachedReportConfigurationId) ID() string {
	fmtString := "/deviceManagement/reports/cachedReportConfigurations/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCachedReportConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Report Cached Report Configuration ID
func (id DeviceManagementReportCachedReportConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("cachedReportConfigurations", "cachedReportConfigurations", "cachedReportConfigurations"),
		resourceids.UserSpecifiedSegment("deviceManagementCachedReportConfigurationId", "deviceManagementCachedReportConfigurationId"),
	}
}

// String returns a human-readable description of this Device Management Report Cached Report Configuration ID
func (id DeviceManagementReportCachedReportConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Cached Report Configuration: %q", id.DeviceManagementCachedReportConfigurationId),
	}
	return fmt.Sprintf("Device Management Report Cached Report Configuration (%s)", strings.Join(components, "\n"))
}
