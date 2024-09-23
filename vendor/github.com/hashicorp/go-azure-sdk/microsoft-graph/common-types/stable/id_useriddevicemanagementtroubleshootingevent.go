package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceManagementTroubleshootingEventId{}

// UserIdDeviceManagementTroubleshootingEventId is a struct representing the Resource ID for a User Id Device Management Troubleshooting Event
type UserIdDeviceManagementTroubleshootingEventId struct {
	UserId                                 string
	DeviceManagementTroubleshootingEventId string
}

// NewUserIdDeviceManagementTroubleshootingEventID returns a new UserIdDeviceManagementTroubleshootingEventId struct
func NewUserIdDeviceManagementTroubleshootingEventID(userId string, deviceManagementTroubleshootingEventId string) UserIdDeviceManagementTroubleshootingEventId {
	return UserIdDeviceManagementTroubleshootingEventId{
		UserId:                                 userId,
		DeviceManagementTroubleshootingEventId: deviceManagementTroubleshootingEventId,
	}
}

// ParseUserIdDeviceManagementTroubleshootingEventID parses 'input' into a UserIdDeviceManagementTroubleshootingEventId
func ParseUserIdDeviceManagementTroubleshootingEventID(input string) (*UserIdDeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceManagementTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceManagementTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a UserIdDeviceManagementTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceManagementTroubleshootingEventIDInsensitively(input string) (*UserIdDeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceManagementTroubleshootingEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceManagementTroubleshootingEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceManagementTroubleshootingEventId, ok = input.Parsed["deviceManagementTroubleshootingEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTroubleshootingEventId", input)
	}

	return nil
}

// ValidateUserIdDeviceManagementTroubleshootingEventID checks that 'input' can be parsed as a User Id Device Management Troubleshooting Event ID
func ValidateUserIdDeviceManagementTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceManagementTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Management Troubleshooting Event ID
func (id UserIdDeviceManagementTroubleshootingEventId) ID() string {
	fmtString := "/users/%s/deviceManagementTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceManagementTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Management Troubleshooting Event ID
func (id UserIdDeviceManagementTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("deviceManagementTroubleshootingEvents", "deviceManagementTroubleshootingEvents", "deviceManagementTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("deviceManagementTroubleshootingEventId", "deviceManagementTroubleshootingEventId"),
	}
}

// String returns a human-readable description of this User Id Device Management Troubleshooting Event ID
func (id UserIdDeviceManagementTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device Management Troubleshooting Event: %q", id.DeviceManagementTroubleshootingEventId),
	}
	return fmt.Sprintf("User Id Device Management Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
