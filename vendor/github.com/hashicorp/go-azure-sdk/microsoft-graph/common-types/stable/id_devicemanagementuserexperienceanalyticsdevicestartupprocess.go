package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{}

// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId is a struct representing the Resource ID for a Device Management User Experience Analytics Device Startup Process
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId struct {
	UserExperienceAnalyticsDeviceStartupProcessId string
}

// NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID returns a new DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId struct
func NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(userExperienceAnalyticsDeviceStartupProcessId string) DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId {
	return DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{
		UserExperienceAnalyticsDeviceStartupProcessId: userExperienceAnalyticsDeviceStartupProcessId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID parses 'input' into a DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId
func ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(input string) (*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsDeviceStartupProcessId, ok = input.Parsed["userExperienceAnalyticsDeviceStartupProcessId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsDeviceStartupProcessId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID checks that 'input' can be parsed as a Device Management User Experience Analytics Device Startup Process ID
func ValidateDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsDeviceStartupProcessID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Device Startup Process ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsDeviceStartupProcessId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Device Startup Process ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsDeviceStartupProcesses", "userExperienceAnalyticsDeviceStartupProcesses", "userExperienceAnalyticsDeviceStartupProcesses"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsDeviceStartupProcessId", "userExperienceAnalyticsDeviceStartupProcessId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Device Startup Process ID
func (id DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Device Startup Process: %q", id.UserExperienceAnalyticsDeviceStartupProcessId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Device Startup Process (%s)", strings.Join(components, "\n"))
}
