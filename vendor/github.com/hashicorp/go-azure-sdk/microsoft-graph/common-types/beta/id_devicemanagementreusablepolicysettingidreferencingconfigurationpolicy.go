package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{}

// DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId is a struct representing the Resource ID for a Device Management Reusable Policy Setting Id Referencing Configuration Policy
type DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId struct {
	DeviceManagementReusablePolicySettingId string
	DeviceManagementConfigurationPolicyId   string
}

// NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID returns a new DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId struct
func NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(deviceManagementReusablePolicySettingId string, deviceManagementConfigurationPolicyId string) DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId {
	return DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{
		DeviceManagementReusablePolicySettingId: deviceManagementReusablePolicySettingId,
		DeviceManagementConfigurationPolicyId:   deviceManagementConfigurationPolicyId,
	}
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID parses 'input' into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIDInsensitively parses 'input' case-insensitively into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIDInsensitively(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementReusablePolicySettingId, ok = input.Parsed["deviceManagementReusablePolicySettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementReusablePolicySettingId", input)
	}

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	return nil
}

// ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID checks that 'input' can be parsed as a Device Management Reusable Policy Setting Id Referencing Configuration Policy ID
func ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Reusable Policy Setting Id Referencing Configuration Policy ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId) ID() string {
	fmtString := "/deviceManagement/reusablePolicySettings/%s/referencingConfigurationPolicies/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementReusablePolicySettingId, id.DeviceManagementConfigurationPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Reusable Policy Setting Id Referencing Configuration Policy ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reusablePolicySettings", "reusablePolicySettings", "reusablePolicySettings"),
		resourceids.UserSpecifiedSegment("deviceManagementReusablePolicySettingId", "deviceManagementReusablePolicySettingId"),
		resourceids.StaticSegment("referencingConfigurationPolicies", "referencingConfigurationPolicies", "referencingConfigurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
	}
}

// String returns a human-readable description of this Device Management Reusable Policy Setting Id Referencing Configuration Policy ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Reusable Policy Setting: %q", id.DeviceManagementReusablePolicySettingId),
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
	}
	return fmt.Sprintf("Device Management Reusable Policy Setting Id Referencing Configuration Policy (%s)", strings.Join(components, "\n"))
}
