package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{}

// DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId is a struct representing the Resource ID for a Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Assignment
type DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId struct {
	DeviceManagementReusablePolicySettingId         string
	DeviceManagementConfigurationPolicyId           string
	DeviceManagementConfigurationPolicyAssignmentId string
}

// NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID returns a new DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId struct
func NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(deviceManagementReusablePolicySettingId string, deviceManagementConfigurationPolicyId string, deviceManagementConfigurationPolicyAssignmentId string) DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId {
	return DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
		DeviceManagementReusablePolicySettingId:         deviceManagementReusablePolicySettingId,
		DeviceManagementConfigurationPolicyId:           deviceManagementConfigurationPolicyId,
		DeviceManagementConfigurationPolicyAssignmentId: deviceManagementConfigurationPolicyAssignmentId,
	}
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID parses 'input' into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentIDInsensitively parses 'input' case-insensitively into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentIDInsensitively(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementReusablePolicySettingId, ok = input.Parsed["deviceManagementReusablePolicySettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementReusablePolicySettingId", input)
	}

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId, ok = input.Parsed["deviceManagementConfigurationPolicyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyAssignmentId", input)
	}

	return nil
}

// ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID checks that 'input' can be parsed as a Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Assignment ID
func ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Assignment ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId) ID() string {
	fmtString := "/deviceManagement/reusablePolicySettings/%s/referencingConfigurationPolicies/%s/assignments/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementReusablePolicySettingId, id.DeviceManagementConfigurationPolicyId, id.DeviceManagementConfigurationPolicyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Assignment ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reusablePolicySettings", "reusablePolicySettings", "reusablePolicySettings"),
		resourceids.UserSpecifiedSegment("deviceManagementReusablePolicySettingId", "deviceManagementReusablePolicySettingId"),
		resourceids.StaticSegment("referencingConfigurationPolicies", "referencingConfigurationPolicies", "referencingConfigurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
		resourceids.StaticSegment("assignments", "assignments", "assignments"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyAssignmentId", "deviceManagementConfigurationPolicyAssignmentId"),
	}
}

// String returns a human-readable description of this Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Assignment ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Reusable Policy Setting: %q", id.DeviceManagementReusablePolicySettingId),
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
		fmt.Sprintf("Device Management Configuration Policy Assignment: %q", id.DeviceManagementConfigurationPolicyAssignmentId),
	}
	return fmt.Sprintf("Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Assignment (%s)", strings.Join(components, "\n"))
}
