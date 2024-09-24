package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMonitoringAlertRecordId{}

// DeviceManagementMonitoringAlertRecordId is a struct representing the Resource ID for a Device Management Monitoring Alert Record
type DeviceManagementMonitoringAlertRecordId struct {
	AlertRecordId string
}

// NewDeviceManagementMonitoringAlertRecordID returns a new DeviceManagementMonitoringAlertRecordId struct
func NewDeviceManagementMonitoringAlertRecordID(alertRecordId string) DeviceManagementMonitoringAlertRecordId {
	return DeviceManagementMonitoringAlertRecordId{
		AlertRecordId: alertRecordId,
	}
}

// ParseDeviceManagementMonitoringAlertRecordID parses 'input' into a DeviceManagementMonitoringAlertRecordId
func ParseDeviceManagementMonitoringAlertRecordID(input string) (*DeviceManagementMonitoringAlertRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMonitoringAlertRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMonitoringAlertRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMonitoringAlertRecordIDInsensitively parses 'input' case-insensitively into a DeviceManagementMonitoringAlertRecordId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMonitoringAlertRecordIDInsensitively(input string) (*DeviceManagementMonitoringAlertRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMonitoringAlertRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMonitoringAlertRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMonitoringAlertRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AlertRecordId, ok = input.Parsed["alertRecordId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertRecordId", input)
	}

	return nil
}

// ValidateDeviceManagementMonitoringAlertRecordID checks that 'input' can be parsed as a Device Management Monitoring Alert Record ID
func ValidateDeviceManagementMonitoringAlertRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMonitoringAlertRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Monitoring Alert Record ID
func (id DeviceManagementMonitoringAlertRecordId) ID() string {
	fmtString := "/deviceManagement/monitoring/alertRecords/%s"
	return fmt.Sprintf(fmtString, id.AlertRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Monitoring Alert Record ID
func (id DeviceManagementMonitoringAlertRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("monitoring", "monitoring", "monitoring"),
		resourceids.StaticSegment("alertRecords", "alertRecords", "alertRecords"),
		resourceids.UserSpecifiedSegment("alertRecordId", "alertRecordId"),
	}
}

// String returns a human-readable description of this Device Management Monitoring Alert Record ID
func (id DeviceManagementMonitoringAlertRecordId) String() string {
	components := []string{
		fmt.Sprintf("Alert Record: %q", id.AlertRecordId),
	}
	return fmt.Sprintf("Device Management Monitoring Alert Record (%s)", strings.Join(components, "\n"))
}
