package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{}

// DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId is a struct representing the Resource ID for a Device Management User Experience Analytics Not Autopilot Ready Device
type DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId struct {
	UserExperienceAnalyticsNotAutopilotReadyDeviceId string
}

// NewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID returns a new DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId struct
func NewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(userExperienceAnalyticsNotAutopilotReadyDeviceId string) DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId {
	return DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{
		UserExperienceAnalyticsNotAutopilotReadyDeviceId: userExperienceAnalyticsNotAutopilotReadyDeviceId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID parses 'input' into a DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId
func ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(input string) (*DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsNotAutopilotReadyDeviceId, ok = input.Parsed["userExperienceAnalyticsNotAutopilotReadyDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsNotAutopilotReadyDeviceId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID checks that 'input' can be parsed as a Device Management User Experience Analytics Not Autopilot Ready Device ID
func ValidateDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Not Autopilot Ready Device ID
func (id DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsNotAutopilotReadyDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Not Autopilot Ready Device ID
func (id DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsNotAutopilotReadyDevice", "userExperienceAnalyticsNotAutopilotReadyDevice", "userExperienceAnalyticsNotAutopilotReadyDevice"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsNotAutopilotReadyDeviceId", "userExperienceAnalyticsNotAutopilotReadyDeviceId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Not Autopilot Ready Device ID
func (id DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Not Autopilot Ready Device: %q", id.UserExperienceAnalyticsNotAutopilotReadyDeviceId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Not Autopilot Ready Device (%s)", strings.Join(components, "\n"))
}
