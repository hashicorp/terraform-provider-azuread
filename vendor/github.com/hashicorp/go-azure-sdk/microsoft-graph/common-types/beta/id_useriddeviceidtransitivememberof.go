package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdTransitiveMemberOfId{}

// UserIdDeviceIdTransitiveMemberOfId is a struct representing the Resource ID for a User Id Device Id Transitive Member Of
type UserIdDeviceIdTransitiveMemberOfId struct {
	UserId            string
	DeviceId          string
	DirectoryObjectId string
}

// NewUserIdDeviceIdTransitiveMemberOfID returns a new UserIdDeviceIdTransitiveMemberOfId struct
func NewUserIdDeviceIdTransitiveMemberOfID(userId string, deviceId string, directoryObjectId string) UserIdDeviceIdTransitiveMemberOfId {
	return UserIdDeviceIdTransitiveMemberOfId{
		UserId:            userId,
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdDeviceIdTransitiveMemberOfID parses 'input' into a UserIdDeviceIdTransitiveMemberOfId
func ParseUserIdDeviceIdTransitiveMemberOfID(input string) (*UserIdDeviceIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdTransitiveMemberOfIDInsensitively(input string) (*UserIdDeviceIdTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdTransitiveMemberOfId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdTransitiveMemberOfId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDeviceIdTransitiveMemberOfID checks that 'input' can be parsed as a User Id Device Id Transitive Member Of ID
func ValidateUserIdDeviceIdTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Transitive Member Of ID
func (id UserIdDeviceIdTransitiveMemberOfId) ID() string {
	fmtString := "/users/%s/devices/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Transitive Member Of ID
func (id UserIdDeviceIdTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("transitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Device Id Transitive Member Of ID
func (id UserIdDeviceIdTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Device Id Transitive Member Of (%s)", strings.Join(components, "\n"))
}
