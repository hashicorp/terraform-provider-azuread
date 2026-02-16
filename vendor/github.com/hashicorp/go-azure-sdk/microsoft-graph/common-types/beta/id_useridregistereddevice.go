package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdRegisteredDeviceId{}

// UserIdRegisteredDeviceId is a struct representing the Resource ID for a User Id Registered Device
type UserIdRegisteredDeviceId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdRegisteredDeviceID returns a new UserIdRegisteredDeviceId struct
func NewUserIdRegisteredDeviceID(userId string, directoryObjectId string) UserIdRegisteredDeviceId {
	return UserIdRegisteredDeviceId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdRegisteredDeviceID parses 'input' into a UserIdRegisteredDeviceId
func ParseUserIdRegisteredDeviceID(input string) (*UserIdRegisteredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdRegisteredDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdRegisteredDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdRegisteredDeviceIDInsensitively parses 'input' case-insensitively into a UserIdRegisteredDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserIdRegisteredDeviceIDInsensitively(input string) (*UserIdRegisteredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdRegisteredDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdRegisteredDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdRegisteredDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdRegisteredDeviceID checks that 'input' can be parsed as a User Id Registered Device ID
func ValidateUserIdRegisteredDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdRegisteredDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Registered Device ID
func (id UserIdRegisteredDeviceId) ID() string {
	fmtString := "/users/%s/registeredDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Registered Device ID
func (id UserIdRegisteredDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("registeredDevices", "registeredDevices", "registeredDevices"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Registered Device ID
func (id UserIdRegisteredDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Registered Device (%s)", strings.Join(components, "\n"))
}
