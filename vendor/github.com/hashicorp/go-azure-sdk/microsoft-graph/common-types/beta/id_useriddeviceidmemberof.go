package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdMemberOfId{}

// UserIdDeviceIdMemberOfId is a struct representing the Resource ID for a User Id Device Id Member Of
type UserIdDeviceIdMemberOfId struct {
	UserId            string
	DeviceId          string
	DirectoryObjectId string
}

// NewUserIdDeviceIdMemberOfID returns a new UserIdDeviceIdMemberOfId struct
func NewUserIdDeviceIdMemberOfID(userId string, deviceId string, directoryObjectId string) UserIdDeviceIdMemberOfId {
	return UserIdDeviceIdMemberOfId{
		UserId:            userId,
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdDeviceIdMemberOfID parses 'input' into a UserIdDeviceIdMemberOfId
func ParseUserIdDeviceIdMemberOfID(input string) (*UserIdDeviceIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdMemberOfIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdMemberOfIDInsensitively(input string) (*UserIdDeviceIdMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdMemberOfId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdDeviceIdMemberOfID checks that 'input' can be parsed as a User Id Device Id Member Of ID
func ValidateUserIdDeviceIdMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Member Of ID
func (id UserIdDeviceIdMemberOfId) ID() string {
	fmtString := "/users/%s/devices/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Member Of ID
func (id UserIdDeviceIdMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("memberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Device Id Member Of ID
func (id UserIdDeviceIdMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Device Id Member Of (%s)", strings.Join(components, "\n"))
}
