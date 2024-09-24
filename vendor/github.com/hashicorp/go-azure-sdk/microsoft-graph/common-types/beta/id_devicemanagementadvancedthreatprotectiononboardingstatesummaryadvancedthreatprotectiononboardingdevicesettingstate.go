package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{}

// DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId is a struct representing the Resource ID for a Device Management Advanced Threat Protection Onboarding State Summary Advanced Threat Protection Onboarding Device Setting State
type DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId struct {
	AdvancedThreatProtectionOnboardingDeviceSettingStateId string
}

// NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID returns a new DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId struct
func NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(advancedThreatProtectionOnboardingDeviceSettingStateId string) DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId {
	return DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{
		AdvancedThreatProtectionOnboardingDeviceSettingStateId: advancedThreatProtectionOnboardingDeviceSettingStateId,
	}
}

// ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID parses 'input' into a DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId
func ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(input string) (*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateIDInsensitively parses 'input' case-insensitively into a DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateIDInsensitively(input string) (*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AdvancedThreatProtectionOnboardingDeviceSettingStateId, ok = input.Parsed["advancedThreatProtectionOnboardingDeviceSettingStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "advancedThreatProtectionOnboardingDeviceSettingStateId", input)
	}

	return nil
}

// ValidateDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID checks that 'input' can be parsed as a Device Management Advanced Threat Protection Onboarding State Summary Advanced Threat Protection Onboarding Device Setting State ID
func ValidateDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Advanced Threat Protection Onboarding State Summary Advanced Threat Protection Onboarding Device Setting State ID
func (id DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId) ID() string {
	fmtString := "/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/%s"
	return fmt.Sprintf(fmtString, id.AdvancedThreatProtectionOnboardingDeviceSettingStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Advanced Threat Protection Onboarding State Summary Advanced Threat Protection Onboarding Device Setting State ID
func (id DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("advancedThreatProtectionOnboardingStateSummary", "advancedThreatProtectionOnboardingStateSummary", "advancedThreatProtectionOnboardingStateSummary"),
		resourceids.StaticSegment("advancedThreatProtectionOnboardingDeviceSettingStates", "advancedThreatProtectionOnboardingDeviceSettingStates", "advancedThreatProtectionOnboardingDeviceSettingStates"),
		resourceids.UserSpecifiedSegment("advancedThreatProtectionOnboardingDeviceSettingStateId", "advancedThreatProtectionOnboardingDeviceSettingStateId"),
	}
}

// String returns a human-readable description of this Device Management Advanced Threat Protection Onboarding State Summary Advanced Threat Protection Onboarding Device Setting State ID
func (id DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateId) String() string {
	components := []string{
		fmt.Sprintf("Advanced Threat Protection Onboarding Device Setting State: %q", id.AdvancedThreatProtectionOnboardingDeviceSettingStateId),
	}
	return fmt.Sprintf("Device Management Advanced Threat Protection Onboarding State Summary Advanced Threat Protection Onboarding Device Setting State (%s)", strings.Join(components, "\n"))
}
