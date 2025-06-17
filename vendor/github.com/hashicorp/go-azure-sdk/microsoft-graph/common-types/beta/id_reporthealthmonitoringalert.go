package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportHealthMonitoringAlertId{}

// ReportHealthMonitoringAlertId is a struct representing the Resource ID for a Report Health Monitoring Alert
type ReportHealthMonitoringAlertId struct {
	AlertId string
}

// NewReportHealthMonitoringAlertID returns a new ReportHealthMonitoringAlertId struct
func NewReportHealthMonitoringAlertID(alertId string) ReportHealthMonitoringAlertId {
	return ReportHealthMonitoringAlertId{
		AlertId: alertId,
	}
}

// ParseReportHealthMonitoringAlertID parses 'input' into a ReportHealthMonitoringAlertId
func ParseReportHealthMonitoringAlertID(input string) (*ReportHealthMonitoringAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportHealthMonitoringAlertId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportHealthMonitoringAlertId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportHealthMonitoringAlertIDInsensitively parses 'input' case-insensitively into a ReportHealthMonitoringAlertId
// note: this method should only be used for API response data and not user input
func ParseReportHealthMonitoringAlertIDInsensitively(input string) (*ReportHealthMonitoringAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportHealthMonitoringAlertId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportHealthMonitoringAlertId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportHealthMonitoringAlertId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AlertId, ok = input.Parsed["alertId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertId", input)
	}

	return nil
}

// ValidateReportHealthMonitoringAlertID checks that 'input' can be parsed as a Report Health Monitoring Alert ID
func ValidateReportHealthMonitoringAlertID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportHealthMonitoringAlertID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Health Monitoring Alert ID
func (id ReportHealthMonitoringAlertId) ID() string {
	fmtString := "/reports/healthMonitoring/alerts/%s"
	return fmt.Sprintf(fmtString, id.AlertId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Health Monitoring Alert ID
func (id ReportHealthMonitoringAlertId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("healthMonitoring", "healthMonitoring", "healthMonitoring"),
		resourceids.StaticSegment("alerts", "alerts", "alerts"),
		resourceids.UserSpecifiedSegment("alertId", "alertId"),
	}
}

// String returns a human-readable description of this Report Health Monitoring Alert ID
func (id ReportHealthMonitoringAlertId) String() string {
	components := []string{
		fmt.Sprintf("Alert: %q", id.AlertId),
	}
	return fmt.Sprintf("Report Health Monitoring Alert (%s)", strings.Join(components, "\n"))
}
