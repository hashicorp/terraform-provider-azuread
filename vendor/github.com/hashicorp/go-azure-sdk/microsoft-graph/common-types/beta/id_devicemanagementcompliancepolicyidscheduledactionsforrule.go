package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{}

// DeviceManagementCompliancePolicyIdScheduledActionsForRuleId is a struct representing the Resource ID for a Device Management Compliance Policy Id Scheduled Actions For Rule
type DeviceManagementCompliancePolicyIdScheduledActionsForRuleId struct {
	DeviceManagementCompliancePolicyId                 string
	DeviceManagementComplianceScheduledActionForRuleId string
}

// NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleID returns a new DeviceManagementCompliancePolicyIdScheduledActionsForRuleId struct
func NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(deviceManagementCompliancePolicyId string, deviceManagementComplianceScheduledActionForRuleId string) DeviceManagementCompliancePolicyIdScheduledActionsForRuleId {
	return DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{
		DeviceManagementCompliancePolicyId:                 deviceManagementCompliancePolicyId,
		DeviceManagementComplianceScheduledActionForRuleId: deviceManagementComplianceScheduledActionForRuleId,
	}
}

// ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleID parses 'input' into a DeviceManagementCompliancePolicyIdScheduledActionsForRuleId
func ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(input string) (*DeviceManagementCompliancePolicyIdScheduledActionsForRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIDInsensitively parses 'input' case-insensitively into a DeviceManagementCompliancePolicyIdScheduledActionsForRuleId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIDInsensitively(input string) (*DeviceManagementCompliancePolicyIdScheduledActionsForRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdScheduledActionsForRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCompliancePolicyIdScheduledActionsForRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCompliancePolicyId, ok = input.Parsed["deviceManagementCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCompliancePolicyId", input)
	}

	if id.DeviceManagementComplianceScheduledActionForRuleId, ok = input.Parsed["deviceManagementComplianceScheduledActionForRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementComplianceScheduledActionForRuleId", input)
	}

	return nil
}

// ValidateDeviceManagementCompliancePolicyIdScheduledActionsForRuleID checks that 'input' can be parsed as a Device Management Compliance Policy Id Scheduled Actions For Rule ID
func ValidateDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Policy Id Scheduled Actions For Rule ID
func (id DeviceManagementCompliancePolicyIdScheduledActionsForRuleId) ID() string {
	fmtString := "/deviceManagement/compliancePolicies/%s/scheduledActionsForRule/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCompliancePolicyId, id.DeviceManagementComplianceScheduledActionForRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Policy Id Scheduled Actions For Rule ID
func (id DeviceManagementCompliancePolicyIdScheduledActionsForRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("compliancePolicies", "compliancePolicies", "compliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementCompliancePolicyId", "deviceManagementCompliancePolicyId"),
		resourceids.StaticSegment("scheduledActionsForRule", "scheduledActionsForRule", "scheduledActionsForRule"),
		resourceids.UserSpecifiedSegment("deviceManagementComplianceScheduledActionForRuleId", "deviceManagementComplianceScheduledActionForRuleId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Policy Id Scheduled Actions For Rule ID
func (id DeviceManagementCompliancePolicyIdScheduledActionsForRuleId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Compliance Policy: %q", id.DeviceManagementCompliancePolicyId),
		fmt.Sprintf("Device Management Compliance Scheduled Action For Rule: %q", id.DeviceManagementComplianceScheduledActionForRuleId),
	}
	return fmt.Sprintf("Device Management Compliance Policy Id Scheduled Actions For Rule (%s)", strings.Join(components, "\n"))
}
