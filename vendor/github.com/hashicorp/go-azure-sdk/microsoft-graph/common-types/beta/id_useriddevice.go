package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceId{}

// UserIdDeviceId is a struct representing the Resource ID for a User Id Device
type UserIdDeviceId struct {
	UserId   string
	DeviceId string
}

// NewUserIdDeviceID returns a new UserIdDeviceId struct
func NewUserIdDeviceID(userId string, deviceId string) UserIdDeviceId {
	return UserIdDeviceId{
		UserId:   userId,
		DeviceId: deviceId,
	}
}

// ParseUserIdDeviceID parses 'input' into a UserIdDeviceId
func ParseUserIdDeviceID(input string) (*UserIdDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIDInsensitively parses 'input' case-insensitively into a UserIdDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIDInsensitively(input string) (*UserIdDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	return nil
}

// ValidateUserIdDeviceID checks that 'input' can be parsed as a User Id Device ID
func ValidateUserIdDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device ID
func (id UserIdDeviceId) ID() string {
	fmtString := "/users/%s/devices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device ID
func (id UserIdDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
	}
}

// String returns a human-readable description of this User Id Device ID
func (id UserIdDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
	}
	return fmt.Sprintf("User Id Device (%s)", strings.Join(components, "\n"))
}
