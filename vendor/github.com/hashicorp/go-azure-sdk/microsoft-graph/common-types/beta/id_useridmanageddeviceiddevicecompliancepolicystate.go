package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdDeviceCompliancePolicyStateId{}

// UserIdManagedDeviceIdDeviceCompliancePolicyStateId is a struct representing the Resource ID for a User Id Managed Device Id Device Compliance Policy State
type UserIdManagedDeviceIdDeviceCompliancePolicyStateId struct {
	UserId                        string
	ManagedDeviceId               string
	DeviceCompliancePolicyStateId string
}

// NewUserIdManagedDeviceIdDeviceCompliancePolicyStateID returns a new UserIdManagedDeviceIdDeviceCompliancePolicyStateId struct
func NewUserIdManagedDeviceIdDeviceCompliancePolicyStateID(userId string, managedDeviceId string, deviceCompliancePolicyStateId string) UserIdManagedDeviceIdDeviceCompliancePolicyStateId {
	return UserIdManagedDeviceIdDeviceCompliancePolicyStateId{
		UserId:                        userId,
		ManagedDeviceId:               managedDeviceId,
		DeviceCompliancePolicyStateId: deviceCompliancePolicyStateId,
	}
}

// ParseUserIdManagedDeviceIdDeviceCompliancePolicyStateID parses 'input' into a UserIdManagedDeviceIdDeviceCompliancePolicyStateId
func ParseUserIdManagedDeviceIdDeviceCompliancePolicyStateID(input string) (*UserIdManagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceIdDeviceCompliancePolicyStateId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively(input string) (*UserIdManagedDeviceIdDeviceCompliancePolicyStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceIdDeviceCompliancePolicyStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceIdDeviceCompliancePolicyStateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceIdDeviceCompliancePolicyStateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DeviceCompliancePolicyStateId, ok = input.Parsed["deviceCompliancePolicyStateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceCompliancePolicyStateId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceIdDeviceCompliancePolicyStateID checks that 'input' can be parsed as a User Id Managed Device Id Device Compliance Policy State ID
func ValidateUserIdManagedDeviceIdDeviceCompliancePolicyStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceIdDeviceCompliancePolicyStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device Id Device Compliance Policy State ID
func (id UserIdManagedDeviceIdDeviceCompliancePolicyStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/deviceCompliancePolicyStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DeviceCompliancePolicyStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device Id Device Compliance Policy State ID
func (id UserIdManagedDeviceIdDeviceCompliancePolicyStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("deviceCompliancePolicyStates", "deviceCompliancePolicyStates", "deviceCompliancePolicyStates"),
		resourceids.UserSpecifiedSegment("deviceCompliancePolicyStateId", "deviceCompliancePolicyStateId"),
	}
}

// String returns a human-readable description of this User Id Managed Device Id Device Compliance Policy State ID
func (id UserIdManagedDeviceIdDeviceCompliancePolicyStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Compliance Policy State: %q", id.DeviceCompliancePolicyStateId),
	}
	return fmt.Sprintf("User Id Managed Device Id Device Compliance Policy State (%s)", strings.Join(components, "\n"))
}
