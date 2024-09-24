package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMonitoringAlertRuleId{}

// DeviceManagementMonitoringAlertRuleId is a struct representing the Resource ID for a Device Management Monitoring Alert Rule
type DeviceManagementMonitoringAlertRuleId struct {
	AlertRuleId string
}

// NewDeviceManagementMonitoringAlertRuleID returns a new DeviceManagementMonitoringAlertRuleId struct
func NewDeviceManagementMonitoringAlertRuleID(alertRuleId string) DeviceManagementMonitoringAlertRuleId {
	return DeviceManagementMonitoringAlertRuleId{
		AlertRuleId: alertRuleId,
	}
}

// ParseDeviceManagementMonitoringAlertRuleID parses 'input' into a DeviceManagementMonitoringAlertRuleId
func ParseDeviceManagementMonitoringAlertRuleID(input string) (*DeviceManagementMonitoringAlertRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMonitoringAlertRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMonitoringAlertRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMonitoringAlertRuleIDInsensitively parses 'input' case-insensitively into a DeviceManagementMonitoringAlertRuleId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMonitoringAlertRuleIDInsensitively(input string) (*DeviceManagementMonitoringAlertRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMonitoringAlertRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMonitoringAlertRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMonitoringAlertRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AlertRuleId, ok = input.Parsed["alertRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertRuleId", input)
	}

	return nil
}

// ValidateDeviceManagementMonitoringAlertRuleID checks that 'input' can be parsed as a Device Management Monitoring Alert Rule ID
func ValidateDeviceManagementMonitoringAlertRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMonitoringAlertRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Monitoring Alert Rule ID
func (id DeviceManagementMonitoringAlertRuleId) ID() string {
	fmtString := "/deviceManagement/monitoring/alertRules/%s"
	return fmt.Sprintf(fmtString, id.AlertRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Monitoring Alert Rule ID
func (id DeviceManagementMonitoringAlertRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("monitoring", "monitoring", "monitoring"),
		resourceids.StaticSegment("alertRules", "alertRules", "alertRules"),
		resourceids.UserSpecifiedSegment("alertRuleId", "alertRuleId"),
	}
}

// String returns a human-readable description of this Device Management Monitoring Alert Rule ID
func (id DeviceManagementMonitoringAlertRuleId) String() string {
	components := []string{
		fmt.Sprintf("Alert Rule: %q", id.AlertRuleId),
	}
	return fmt.Sprintf("Device Management Monitoring Alert Rule (%s)", strings.Join(components, "\n"))
}
