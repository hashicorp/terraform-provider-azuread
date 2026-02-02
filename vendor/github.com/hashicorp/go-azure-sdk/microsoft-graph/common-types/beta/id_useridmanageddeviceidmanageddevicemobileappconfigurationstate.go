package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}

// UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId is a struct representing the Resource ID for a User Id Managed Device Id Managed Device Mobile App Configuration State
type UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct {
	UserId                                     string
	ManagedDeviceId                            string
	ManagedDeviceMobileAppConfigurationStateId string
}

// NewUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID returns a new UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId struct
func NewUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(userId string, managedDeviceId string, managedDeviceMobileAppConfigurationStateId string) UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId {
	return UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
		ManagedDeviceMobileAppConfigurationStateId: managedDeviceMobileAppConfigurationStateId,
	}
}

// ParseUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID parses 'input' into a UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
func ParseUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input string) (*UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(input string) (*UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.ManagedDeviceMobileAppConfigurationStateId, ok = input.Parsed["managedDeviceMobileAppConfigurationStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceMobileAppConfigurationStateId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID checks that 'input' can be parsed as a User Id Managed Device Id Managed Device Mobile App Configuration State ID
func ValidateUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Managed Device Mobile App Configuration State ID
func (id UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/managedDeviceMobileAppConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.ManagedDeviceMobileAppConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Managed Device Mobile App Configuration State ID
func (id UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates", "managedDeviceMobileAppConfigurationStates"),
		resourceids.UserSpecifiedSegment("managedDeviceMobileAppConfigurationStateId", "managedDeviceMobileAppConfigurationStateId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Managed Device Mobile App Configuration State ID
func (id UserIdManagedDeviceIdManagedDeviceMobileAppConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Managed Device Mobile App Configuration State: %q", id.ManagedDeviceMobileAppConfigurationStateId),
	}
	return fmt.Sprintf("User Id Managed Device Id Managed Device Mobile App Configuration State (%s)", strings.Join(components, "\n"))
}
