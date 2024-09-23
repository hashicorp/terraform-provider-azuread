package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOwnedDeviceId{}

// UserIdOwnedDeviceId is a struct representing the Resource ID for a User Id Owned Device
type UserIdOwnedDeviceId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdOwnedDeviceID returns a new UserIdOwnedDeviceId struct
func NewUserIdOwnedDeviceID(userId string, directoryObjectId string) UserIdOwnedDeviceId {
	return UserIdOwnedDeviceId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdOwnedDeviceID parses 'input' into a UserIdOwnedDeviceId
func ParseUserIdOwnedDeviceID(input string) (*UserIdOwnedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOwnedDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOwnedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOwnedDeviceIDInsensitively parses 'input' case-insensitively into a UserIdOwnedDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserIdOwnedDeviceIDInsensitively(input string) (*UserIdOwnedDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOwnedDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOwnedDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOwnedDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdOwnedDeviceID checks that 'input' can be parsed as a User Id Owned Device ID
func ValidateUserIdOwnedDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOwnedDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Owned Device ID
func (id UserIdOwnedDeviceId) ID() string {
	fmtString := "/users/%s/ownedDevices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Owned Device ID
func (id UserIdOwnedDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("ownedDevices", "ownedDevices", "ownedDevices"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Owned Device ID
func (id UserIdOwnedDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Owned Device (%s)", strings.Join(components, "\n"))
}
