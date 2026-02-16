package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceId{}

// UserIdManagedDeviceId is a struct representing the Resource ID for a User Id Managed Device
type UserIdManagedDeviceId struct {
	UserId          string
	ManagedDeviceId string
}

// NewUserIdManagedDeviceID returns a new UserIdManagedDeviceId struct
func NewUserIdManagedDeviceID(userId string, managedDeviceId string) UserIdManagedDeviceId {
	return UserIdManagedDeviceId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
	}
}

// ParseUserIdManagedDeviceID parses 'input' into a UserIdManagedDeviceId
func ParseUserIdManagedDeviceID(input string) (*UserIdManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedDeviceIDInsensitively parses 'input' case-insensitively into a UserIdManagedDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedDeviceIDInsensitively(input string) (*UserIdManagedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	return nil
}

// ValidateUserIdManagedDeviceID checks that 'input' can be parsed as a User Id Managed Device ID
func ValidateUserIdManagedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed Device ID
func (id UserIdManagedDeviceId) ID() string {
	fmtString := "/users/%s/managedDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed Device ID
func (id UserIdManagedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
	}
}

// String returns a human-readable description of this User Id Managed Device ID
func (id UserIdManagedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
	}
	return fmt.Sprintf("User Id Managed Device (%s)", strings.Join(components, "\n"))
}
