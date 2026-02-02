package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}

// DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId is a struct representing the Resource ID for a Device Management Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration
type DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId struct {
	DeviceManagementCompliancePolicyId                 string
	DeviceManagementComplianceScheduledActionForRuleId string
	DeviceManagementComplianceActionItemId             string
}

// NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID returns a new DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId struct
func NewDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(deviceManagementCompliancePolicyId string, deviceManagementComplianceScheduledActionForRuleId string, deviceManagementComplianceActionItemId string) DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId {
	return DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{
		DeviceManagementCompliancePolicyId:                 deviceManagementCompliancePolicyId,
		DeviceManagementComplianceScheduledActionForRuleId: deviceManagementComplianceScheduledActionForRuleId,
		DeviceManagementComplianceActionItemId:             deviceManagementComplianceActionItemId,
	}
}

// ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID parses 'input' into a DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
func ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(input string) (*DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively parses 'input' case-insensitively into a DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationIDInsensitively(input string) (*DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementCompliancePolicyId, ok = input.Parsed["deviceManagementCompliancePolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementCompliancePolicyId", input)
	}

	if id.DeviceManagementComplianceScheduledActionForRuleId, ok = input.Parsed["deviceManagementComplianceScheduledActionForRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementComplianceScheduledActionForRuleId", input)
	}

	if id.DeviceManagementComplianceActionItemId, ok = input.Parsed["deviceManagementComplianceActionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementComplianceActionItemId", input)
	}

	return nil
}

// ValidateDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID checks that 'input' can be parsed as a Device Management Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func ValidateDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func (id DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) ID() string {
	fmtString := "/deviceManagement/compliancePolicies/%s/scheduledActionsForRule/%s/scheduledActionConfigurations/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementCompliancePolicyId, id.DeviceManagementComplianceScheduledActionForRuleId, id.DeviceManagementComplianceActionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func (id DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("compliancePolicies", "compliancePolicies", "compliancePolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementCompliancePolicyId", "deviceManagementCompliancePolicyId"),
		resourceids.StaticSegment("scheduledActionsForRule", "scheduledActionsForRule", "scheduledActionsForRule"),
		resourceids.UserSpecifiedSegment("deviceManagementComplianceScheduledActionForRuleId", "deviceManagementComplianceScheduledActionForRuleId"),
		resourceids.StaticSegment("scheduledActionConfigurations", "scheduledActionConfigurations", "scheduledActionConfigurations"),
		resourceids.UserSpecifiedSegment("deviceManagementComplianceActionItemId", "deviceManagementComplianceActionItemId"),
	}
}

// String returns a human-readable description of this Device Management Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration ID
func (id DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Compliance Policy: %q", id.DeviceManagementCompliancePolicyId),
		fmt.Sprintf("Device Management Compliance Scheduled Action For Rule: %q", id.DeviceManagementComplianceScheduledActionForRuleId),
		fmt.Sprintf("Device Management Compliance Action Item: %q", id.DeviceManagementComplianceActionItemId),
	}
	return fmt.Sprintf("Device Management Compliance Policy Id Scheduled Actions For Rule Id Scheduled Action Configuration (%s)", strings.Join(components, "\n"))
}
