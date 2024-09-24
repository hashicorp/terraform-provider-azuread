package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}

// DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId is a struct representing the Resource ID for a Device Management Device Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration
type DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId struct {
	DeviceCompliancePolicyId                 string
	DeviceComplianceScheduledActionForRuleId string
	DeviceComplianceActionItemId             string
}

// NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID returns a new DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId struct
func NewDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(deviceCompliancePolicyId string, deviceComplianceScheduledActionForRuleId string, deviceComplianceActionItemId string) DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId {
	return DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
		DeviceCompliancePolicyId:                 deviceCompliancePolicyId,
		DeviceComplianceScheduledActionForRuleId: deviceComplianceScheduledActionForRuleId,
		DeviceComplianceActionItemId:             deviceComplianceActionItemId,
	}
}

// ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID parses 'input' into a DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
func ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(input string) (*DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively(input string) (*DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceCompliancePolicyId, ok = input.Parsed["deviceCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyId", input)
	}

	if id.DeviceComplianceScheduledActionForRuleId, ok = input.Parsed["deviceComplianceScheduledActionForRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceScheduledActionForRuleId", input)
	}

	if id.DeviceComplianceActionItemId, ok = input.Parsed["deviceComplianceActionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceComplianceActionItemId", input)
	}

	return nil
}

// ValidateDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID checks that 'input' can be parsed as a Device Management Device Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func ValidateDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Device Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func (id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) ID() string {
	fmtString := "/deviceManagement/deviceCompliancePolicies/%s/scheduledActionsForRule/%s/scheduledActionConfigurations/%s"
	return fmt.Sprintf(fmtString, id.DeviceCompliancePolicyId, id.DeviceComplianceScheduledActionForRuleId, id.DeviceComplianceActionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Device Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func (id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("deviceCompliancePolicies", "deviceCompliancePolicies", "deviceCompliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyId", "deviceCompliancePolicyId"),
		resourceids.StaticSegment("scheduledActionsForRule", "scheduledActionsForRule", "scheduledActionsForRule"),
		resourceids.UserSpecifiedSegment("deviceComplianceScheduledActionForRuleId", "deviceComplianceScheduledActionForRuleId"),
		resourceids.StaticSegment("scheduledActionConfigurations", "scheduledActionConfigurations", "scheduledActionConfigurations"),
		resourceids.UserSpecifiedSegment("deviceComplianceActionItemId", "deviceComplianceActionItemId"),
	}
}

// String returns a human-readable description of this Device Management Device Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func (id DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Device Compliance Policy: %q", id.DeviceCompliancePolicyId),
		fmt.Sprintf("Device Compliance Scheduled Action For Rule: %q", id.DeviceComplianceScheduledActionForRuleId),
		fmt.Sprintf("Device Compliance Action Item: %q", id.DeviceComplianceActionItemId),
	}
	return fmt.Sprintf("Device Management Device Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration (%s)", strings.Join(components, "\n"))
}
