package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdUsageRightId{}

// UserIdDeviceIdUsageRightId is a struct representing the Resource ID for a User Id Device Id Usage Right
type UserIdDeviceIdUsageRightId struct {
	UserId       string
	DeviceId     string
	UsageRightId string
}

// NewUserIdDeviceIdUsageRightID returns a new UserIdDeviceIdUsageRightId struct
func NewUserIdDeviceIdUsageRightID(userId string, deviceId string, usageRightId string) UserIdDeviceIdUsageRightId {
	return UserIdDeviceIdUsageRightId{
		UserId:       userId,
		DeviceId:     deviceId,
		UsageRightId: usageRightId,
	}
}

// ParseUserIdDeviceIdUsageRightID parses 'input' into a UserIdDeviceIdUsageRightId
func ParseUserIdDeviceIdUsageRightID(input string) (*UserIdDeviceIdUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdUsageRightIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdUsageRightId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdUsageRightIDInsensitively(input string) (*UserIdDeviceIdUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdUsageRightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdUsageRightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.UsageRightId, ok = input.Parsed["usageRightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", input)
	}

	return nil
}

// ValidateUserIdDeviceIdUsageRightID checks that 'input' can be parsed as a User Id Device Id Usage Right ID
func ValidateUserIdDeviceIdUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Usage Right ID
func (id UserIdDeviceIdUsageRightId) ID() string {
	fmtString := "/users/%s/devices/%s/usageRights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Usage Right ID
func (id UserIdDeviceIdUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("usageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightId"),
	}
}

// String returns a human-readable description of this User Id Device Id Usage Right ID
func (id UserIdDeviceIdUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("User Id Device Id Usage Right (%s)", strings.Join(components, "\n"))
}
