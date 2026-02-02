package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdCommandId{}

// UserIdDeviceIdCommandId is a struct representing the Resource ID for a User Id Device Id Command
type UserIdDeviceIdCommandId struct {
	UserId    string
	DeviceId  string
	CommandId string
}

// NewUserIdDeviceIdCommandID returns a new UserIdDeviceIdCommandId struct
func NewUserIdDeviceIdCommandID(userId string, deviceId string, commandId string) UserIdDeviceIdCommandId {
	return UserIdDeviceIdCommandId{
		UserId:    userId,
		DeviceId:  deviceId,
		CommandId: commandId,
	}
}

// ParseUserIdDeviceIdCommandID parses 'input' into a UserIdDeviceIdCommandId
func ParseUserIdDeviceIdCommandID(input string) (*UserIdDeviceIdCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdCommandIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdCommandId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdCommandIDInsensitively(input string) (*UserIdDeviceIdCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdCommandId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.CommandId, ok = input.Parsed["commandId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "commandId", input)
	}

	return nil
}

// ValidateUserIdDeviceIdCommandID checks that 'input' can be parsed as a User Id Device Id Command ID
func ValidateUserIdDeviceIdCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Command ID
func (id UserIdDeviceIdCommandId) ID() string {
	fmtString := "/users/%s/devices/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.CommandId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Command ID
func (id UserIdDeviceIdCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("commands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandId", "commandId"),
	}
}

// String returns a human-readable description of this User Id Device Id Command ID
func (id UserIdDeviceIdCommandId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Command: %q", id.CommandId),
	}
	return fmt.Sprintf("User Id Device Id Command (%s)", strings.Join(components, "\n"))
}
