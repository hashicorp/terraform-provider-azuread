package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{}

// DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId is a struct representing the Resource ID for a Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting
type DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId struct {
	DeviceManagementReusablePolicySettingId string
	DeviceManagementConfigurationPolicyId   string
	DeviceManagementConfigurationSettingId  string
}

// NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID returns a new DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId struct
func NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(deviceManagementReusablePolicySettingId string, deviceManagementConfigurationPolicyId string, deviceManagementConfigurationSettingId string) DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId {
	return DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{
		DeviceManagementReusablePolicySettingId: deviceManagementReusablePolicySettingId,
		DeviceManagementConfigurationPolicyId:   deviceManagementConfigurationPolicyId,
		DeviceManagementConfigurationSettingId:  deviceManagementConfigurationSettingId,
	}
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID parses 'input' into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIDInsensitively(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementReusablePolicySettingId, ok = input.Parsed["deviceManagementReusablePolicySettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementReusablePolicySettingId", input)
	}

	if id.DeviceManagementConfigurationPolicyId, ok = input.Parsed["deviceManagementConfigurationPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationPolicyId", input)
	}

	if id.DeviceManagementConfigurationSettingId, ok = input.Parsed["deviceManagementConfigurationSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingId", input)
	}

	return nil
}

// ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID checks that 'input' can be parsed as a Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting ID
func ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId) ID() string {
	fmtString := "/deviceManagement/reusablePolicySettings/%s/referencingConfigurationPolicies/%s/settings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementReusablePolicySettingId, id.DeviceManagementConfigurationPolicyId, id.DeviceManagementConfigurationSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reusablePolicySettings", "reusablePolicySettings", "reusablePolicySettings"),
		resourceids.UserSpecifiedSegment("deviceManagementReusablePolicySettingId", "deviceManagementReusablePolicySettingId"),
		resourceids.StaticSegment("referencingConfigurationPolicies", "referencingConfigurationPolicies", "referencingConfigurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingId", "deviceManagementConfigurationSettingId"),
	}
}

// String returns a human-readable description of this Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Reusable Policy Setting: %q", id.DeviceManagementReusablePolicySettingId),
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
		fmt.Sprintf("Device Management Configuration Setting: %q", id.DeviceManagementConfigurationSettingId),
	}
	return fmt.Sprintf("Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting (%s)", strings.Join(components, "\n"))
}
