package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{}

// DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId is a struct representing the Resource ID for a Device Management Device Compliance Policy Id Scheduled Actions For Rule
type DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId struct {
	DeviceCompliancePolicyId                 string
	DeviceComplianceScheduledActionForRuleId string
}

// NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID returns a new DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId struct
func NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(deviceCompliancePolicyId string, deviceComplianceScheduledActionForRuleId string) DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId {
	return DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{
		DeviceCompliancePolicyId:                 deviceCompliancePolicyId,
		DeviceComplianceScheduledActionForRuleId: deviceComplianceScheduledActionForRuleId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID parses 'input' into a DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId
func ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(input string) (*DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	if id.DeviceComplianceScheduledActionForRuleId, ok = input.Parsed["deviceComplianceScheduledActionForRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceScheduledActionForRuleId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID checks that 'input' can be parsed as a Device Management Device Compliance Policy Id Scheduled Actions For Rule ID
func ValidateDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Id Scheduled Actions For Rule ID
func (id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s/scheduledActionsForRule/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId, id.DeviceComplianceScheduledActionForRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Id Scheduled Actions For Rule ID
func (id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
		resourceids.StaticSegment("scheduledActionsForRule", "scheduledActionsForRule", "scheduledActionsForRule"),
		resourceids.UserSpecifiedSegment("deviceComplianceScheduledActionForRuleId", "deviceComplianceScheduledActionForRuleId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Id Scheduled Actions For Rule ID
func (id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
		fmt.Sprintf("Device Compliance Scheduled Action For Rule: %q", id.DeviceComplianceScheduledActionForRuleId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Id Scheduled Actions For Rule (%s)", strings.Join(components, "\n"))
}
