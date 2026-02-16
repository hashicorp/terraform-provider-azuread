package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateInsightId{}

// DeviceManagementTemplateInsightId is a struct representing the Resource ID for a Device Management Template Insight
type DeviceManagementTemplateInsightId struct {
	DeviceManagementTemplateInsightsDefinitionId string
}

// NewDeviceManagementTemplateInsightID returns a new DeviceManagementTemplateInsightId struct
func NewDeviceManagementTemplateInsightID(deviceManagementTemplateInsightsDefinitionId string) DeviceManagementTemplateInsightId {
	return DeviceManagementTemplateInsightId{
		DeviceManagementTemplateInsightsDefinitionId: deviceManagementTemplateInsightsDefinitionId,
	}
}

// ParseDeviceManagementTemplateInsightID parses 'input' into a DeviceManagementTemplateInsightId
func ParseDeviceManagementTemplateInsightID(input string) (*DeviceManagementTemplateInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementTemplateInsightIDInsensitively parses 'input' case-insensitively into a DeviceManagementTemplateInsightId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementTemplateInsightIDInsensitively(input string) (*DeviceManagementTemplateInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementTemplateInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementTemplateInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementTemplateInsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementTemplateInsightsDefinitionId, ok = input.Parsed["deviceManagementTemplateInsightsDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTemplateInsightsDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementTemplateInsightID checks that 'input' can be parsed as a Device Management Template Insight ID
func ValidateDeviceManagementTemplateInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementTemplateInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Template Insight ID
func (id DeviceManagementTemplateInsightId) ID() string {
	fmtString := "/deviceManagement/templateInsights/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTemplateInsightsDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Template Insight ID
func (id DeviceManagementTemplateInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("templateInsights", "templateInsights", "templateInsights"),
		resourceids.UserSpecifiedSegment("deviceManagementTemplateInsightsDefinitionId", "deviceManagementTemplateInsightsDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Template Insight ID
func (id DeviceManagementTemplateInsightId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Template Insights Definition: %q", id.DeviceManagementTemplateInsightsDefinitionId),
	}
	return fmt.Sprintf("Device Management Template Insight (%s)", strings.Join(components, "\n"))
}
