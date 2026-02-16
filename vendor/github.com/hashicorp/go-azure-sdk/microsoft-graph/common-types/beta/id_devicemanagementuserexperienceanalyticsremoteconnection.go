package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsRemoteConnectionId{}

// DeviceManagementUserExperienceAnalyticsRemoteConnectionId is a struct representing the Resource ID for a Device Management User Experience Analytics Remote Connection
type DeviceManagementUserExperienceAnalyticsRemoteConnectionId struct {
	UserExperienceAnalyticsRemoteConnectionId string
}

// NewDeviceManagementUserExperienceAnalyticsRemoteConnectionID returns a new DeviceManagementUserExperienceAnalyticsRemoteConnectionId struct
func NewDeviceManagementUserExperienceAnalyticsRemoteConnectionID(userExperienceAnalyticsRemoteConnectionId string) DeviceManagementUserExperienceAnalyticsRemoteConnectionId {
	return DeviceManagementUserExperienceAnalyticsRemoteConnectionId{
		UserExperienceAnalyticsRemoteConnectionId: userExperienceAnalyticsRemoteConnectionId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionID parses 'input' into a DeviceManagementUserExperienceAnalyticsRemoteConnectionId
func ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionID(input string) (*DeviceManagementUserExperienceAnalyticsRemoteConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsRemoteConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsRemoteConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsRemoteConnectionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsRemoteConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsRemoteConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsRemoteConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsRemoteConnectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsRemoteConnectionId, ok = input.Parsed["userExperienceAnalyticsRemoteConnectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsRemoteConnectionId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsRemoteConnectionID checks that 'input' can be parsed as a Device Management User Experience Analytics Remote Connection ID
func ValidateDeviceManagementUserExperienceAnalyticsRemoteConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsRemoteConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Remote Connection ID
func (id DeviceManagementUserExperienceAnalyticsRemoteConnectionId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsRemoteConnection/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsRemoteConnectionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Remote Connection ID
func (id DeviceManagementUserExperienceAnalyticsRemoteConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsRemoteConnection", "userExperienceAnalyticsRemoteConnection", "userExperienceAnalyticsRemoteConnection"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsRemoteConnectionId", "userExperienceAnalyticsRemoteConnectionId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Remote Connection ID
func (id DeviceManagementUserExperienceAnalyticsRemoteConnectionId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Remote Connection: %q", id.UserExperienceAnalyticsRemoteConnectionId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Remote Connection (%s)", strings.Join(components, "\n"))
}
