package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{}

// DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId is a struct representing the Resource ID for a Device Management User Experience Analytics Battery Health Device Runtime History
type DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId struct {
	UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId string
}

// NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID returns a new DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId struct
func NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId string) DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId {
	return DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{
		UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId: userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID parses 'input' into a DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId, ok = input.Parsed["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID checks that 'input' can be parsed as a Device Management User Experience Analytics Battery Health Device Runtime History ID
func ValidateDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Battery Health Device Runtime History ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Battery Health Device Runtime History ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId", "userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Battery Health Device Runtime History ID
func (id DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Battery Health Device Runtime History: %q", id.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Battery Health Device Runtime History (%s)", strings.Join(components, "\n"))
}
