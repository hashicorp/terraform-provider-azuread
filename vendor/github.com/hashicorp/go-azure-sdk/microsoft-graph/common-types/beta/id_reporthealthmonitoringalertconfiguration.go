package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportHealthMonitoringAlertConfigurationId{}

// ReportHealthMonitoringAlertConfigurationId is a struct representing the Resource ID for a Report Health Monitoring Alert Configuration
type ReportHealthMonitoringAlertConfigurationId struct {
	AlertConfigurationId string
}

// NewReportHealthMonitoringAlertConfigurationID returns a new ReportHealthMonitoringAlertConfigurationId struct
func NewReportHealthMonitoringAlertConfigurationID(alertConfigurationId string) ReportHealthMonitoringAlertConfigurationId {
	return ReportHealthMonitoringAlertConfigurationId{
		AlertConfigurationId: alertConfigurationId,
	}
}

// ParseReportHealthMonitoringAlertConfigurationID parses 'input' into a ReportHealthMonitoringAlertConfigurationId
func ParseReportHealthMonitoringAlertConfigurationID(input string) (*ReportHealthMonitoringAlertConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportHealthMonitoringAlertConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportHealthMonitoringAlertConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportHealthMonitoringAlertConfigurationIDInsensitively parses 'input' case-insensitively into a ReportHealthMonitoringAlertConfigurationId
// note: this method should only be used for API response data and not user input
func ParseReportHealthMonitoringAlertConfigurationIDInsensitively(input string) (*ReportHealthMonitoringAlertConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportHealthMonitoringAlertConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportHealthMonitoringAlertConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportHealthMonitoringAlertConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AlertConfigurationId, ok = input.Parsed["alertConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertConfigurationId", input)
	}

	return nil
}

// ValidateReportHealthMonitoringAlertConfigurationID checks that 'input' can be parsed as a Report Health Monitoring Alert Configuration ID
func ValidateReportHealthMonitoringAlertConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportHealthMonitoringAlertConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report Health Monitoring Alert Configuration ID
func (id ReportHealthMonitoringAlertConfigurationId) ID() string {
	fmtString := "/reports/healthMonitoring/alertConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AlertConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report Health Monitoring Alert Configuration ID
func (id ReportHealthMonitoringAlertConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("healthMonitoring", "healthMonitoring", "healthMonitoring"),
		resourceids.StaticSegment("alertConfigurations", "alertConfigurations", "alertConfigurations"),
		resourceids.UserSpecifiedSegment("alertConfigurationId", "alertConfigurationId"),
	}
}

// String returns a human-readable description of this Report Health Monitoring Alert Configuration ID
func (id ReportHealthMonitoringAlertConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Alert Configuration: %q", id.AlertConfigurationId),
	}
	return fmt.Sprintf("Report Health Monitoring Alert Configuration (%s)", strings.Join(components, "\n"))
}
