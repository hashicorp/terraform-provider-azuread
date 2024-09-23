package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdExtensionId{}

// UserIdDeviceIdExtensionId is a struct representing the Resource ID for a User Id Device Id Extension
type UserIdDeviceIdExtensionId struct {
	UserId      string
	DeviceId    string
	ExtensionId string
}

// NewUserIdDeviceIdExtensionID returns a new UserIdDeviceIdExtensionId struct
func NewUserIdDeviceIdExtensionID(userId string, deviceId string, extensionId string) UserIdDeviceIdExtensionId {
	return UserIdDeviceIdExtensionId{
		UserId:      userId,
		DeviceId:    deviceId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdDeviceIdExtensionID parses 'input' into a UserIdDeviceIdExtensionId
func ParseUserIdDeviceIdExtensionID(input string) (*UserIdDeviceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdExtensionIDInsensitively(input string) (*UserIdDeviceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdDeviceIdExtensionID checks that 'input' can be parsed as a User Id Device Id Extension ID
func ValidateUserIdDeviceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Extension ID
func (id UserIdDeviceIdExtensionId) ID() string {
	fmtString := "/users/%s/devices/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Extension ID
func (id UserIdDeviceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Device Id Extension ID
func (id UserIdDeviceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Device Id Extension (%s)", strings.Join(components, "\n"))
}
