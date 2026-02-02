package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdDeviceConfigurationStateId{}

// UserIdManagedDeviceIdDeviceConfigurationStateId is a struct representing the Resource ID for a User Id Managed Device Id Device Configuration State
type UserIdManagedDeviceIdDeviceConfigurationStateId struct {
	UserId                     string
	ManagedDeviceId            string
	DeviceConfigurationStateId string
}

// NewUserIdManagedDeviceIdDeviceConfigurationStateID returns a new UserIdManagedDeviceIdDeviceConfigurationStateId struct
func NewUserIdManagedDeviceIdDeviceConfigurationStateID(userId string, managedDeviceId string, deviceConfigurationStateId string) UserIdManagedDeviceIdDeviceConfigurationStateId {
	return UserIdManagedDeviceIdDeviceConfigurationStateId{
		UserId:                     userId,
		ManagedDeviceId:            managedDeviceId,
		DeviceConfigurationStateId: deviceConfigurationStateId,
	}
}

// ParseUserIdManagedDeviceIdDeviceConfigurationStateID parses 'input' into a UserIdManagedDeviceIdDeviceConfigurationStateId
func ParseUserIdManagedDeviceIdDeviceConfigurationStateID(input string) (*UserIdManagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdDeviceConfigurationStateIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdDeviceConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdDeviceConfigurationStateIDInsensitively(input string) (*UserIdManagedDeviceIdDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdDeviceConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdDeviceConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceConfigurationStateId, ok = input.Parsed["deviceConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdDeviceConfigurationStateID checks that 'input' can be parsed as a User Id Managed Device Id Device Configuration State ID
func ValidateUserIdManagedDeviceIdDeviceConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdDeviceConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Device Configuration State ID
func (id UserIdManagedDeviceIdDeviceConfigurationStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/deviceConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DeviceConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Device Configuration State ID
func (id UserIdManagedDeviceIdDeviceConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceConfigurationStates", "deviceConfigurationStates", "deviceConfigurationStates"),
		resourceids.UserSpecifiedSegment("deviceConfigurationStateId", "deviceConfigurationStateId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Device Configuration State ID
func (id UserIdManagedDeviceIdDeviceConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Configuration State: %q", id.DeviceConfigurationStateId),
	}
	return fmt.Sprintf("User Id Managed Device Id Device Configuration State (%s)", strings.Join(components, "\n"))
}
