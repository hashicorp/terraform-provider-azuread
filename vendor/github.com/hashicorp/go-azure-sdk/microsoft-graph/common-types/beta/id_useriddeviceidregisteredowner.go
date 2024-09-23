package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdRegisteredOwnerId{}

// UserIdDeviceIdRegisteredOwnerId is a struct representing the Resource ID for a User Id Device Id Registered Owner
type UserIdDeviceIdRegisteredOwnerId struct {
	UserId            string
	DeviceId          string
	DirectoryObjectId string
}

// NewUserIdDeviceIdRegisteredOwnerID returns a new UserIdDeviceIdRegisteredOwnerId struct
func NewUserIdDeviceIdRegisteredOwnerID(userId string, deviceId string, directoryObjectId string) UserIdDeviceIdRegisteredOwnerId {
	return UserIdDeviceIdRegisteredOwnerId{
		UserId:            userId,
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdDeviceIdRegisteredOwnerID parses 'input' into a UserIdDeviceIdRegisteredOwnerId
func ParseUserIdDeviceIdRegisteredOwnerID(input string) (*UserIdDeviceIdRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdRegisteredOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdRegisteredOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDeviceIdRegisteredOwnerIDInsensitively parses 'input' case-insensitively into a UserIdDeviceIdRegisteredOwnerId
// note: this method should only be used for API response data and not user input
func ParseUserIdDeviceIdRegisteredOwnerIDInsensitively(input string) (*UserIdDeviceIdRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDeviceIdRegisteredOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDeviceIdRegisteredOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDeviceIdRegisteredOwnerId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDeviceIdRegisteredOwnerID checks that 'input' can be parsed as a User Id Device Id Registered Owner ID
func ValidateUserIdDeviceIdRegisteredOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDeviceIdRegisteredOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Device Id Registered Owner ID
func (id UserIdDeviceIdRegisteredOwnerId) ID() string {
	fmtString := "/users/%s/devices/%s/registeredOwners/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Device Id Registered Owner ID
func (id UserIdDeviceIdRegisteredOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("devices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("registeredOwners", "registeredOwners", "registeredOwners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Device Id Registered Owner ID
func (id UserIdDeviceIdRegisteredOwnerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Device Id Registered Owner (%s)", strings.Join(components, "\n"))
}
