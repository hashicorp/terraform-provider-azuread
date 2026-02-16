package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{}

// DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId is a struct representing the Resource ID for a Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting Id Setting Definition
type DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId struct {
	DeviceManagementReusablePolicySettingId          string
	DeviceManagementConfigurationPolicyId            string
	DeviceManagementConfigurationSettingId           string
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID returns a new DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId struct
func NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(deviceManagementReusablePolicySettingId string, deviceManagementConfigurationPolicyId string, deviceManagementConfigurationSettingId string, deviceManagementConfigurationSettingDefinitionId string) DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId {
	return DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{
		DeviceManagementReusablePolicySettingId:          deviceManagementReusablePolicySettingId,
		DeviceManagementConfigurationPolicyId:            deviceManagementConfigurationPolicyId,
		DeviceManagementConfigurationSettingId:           deviceManagementConfigurationSettingId,
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID parses 'input' into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively(input string) (*DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID checks that 'input' can be parsed as a Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting Id Setting Definition ID
func ValidateDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting Id Setting Definition ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId) ID() string {
	fmtString := "/deviceManagement/reusablePolicySettings/%s/referencingConfigurationPolicies/%s/settings/%s/settingDefinitions/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementReusablePolicySettingId, id.DeviceManagementConfigurationPolicyId, id.DeviceManagementConfigurationSettingId, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting Id Setting Definition ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reusablePolicySettings", "reusablePolicySettings", "reusablePolicySettings"),
		resourceids.UserSpecifiedSegment("deviceManagementReusablePolicySettingId", "deviceManagementReusablePolicySettingId"),
		resourceids.StaticSegment("referencingConfigurationPolicies", "referencingConfigurationPolicies", "referencingConfigurationPolicies"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingId", "deviceManagementConfigurationSettingId"),
		resourceids.StaticSegment("settingDefinitions", "settingDefinitions", "settingDefinitions"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting Id Setting Definition ID
func (id DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Reusable Policy Setting: %q", id.DeviceManagementReusablePolicySettingId),
		fmt.Sprintf("Device Management Configuration Policy: %q", id.DeviceManagementConfigurationPolicyId),
		fmt.Sprintf("Device Management Configuration Setting: %q", id.DeviceManagementConfigurationSettingId),
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Reusable Policy Setting Id Referencing Configuration Policy Id Setting Id Setting Definition (%s)", strings.Join(components, "\n"))
}
