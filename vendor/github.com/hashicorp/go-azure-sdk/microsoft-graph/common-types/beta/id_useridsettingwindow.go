package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSettingWindowId{}

// UserIdSettingWindowId is a struct representing the Resource ID for a User Id Setting Window
type UserIdSettingWindowId struct {
	UserId           string
	WindowsSettingId string
}

// NewUserIdSettingWindowID returns a new UserIdSettingWindowId struct
func NewUserIdSettingWindowID(userId string, windowsSettingId string) UserIdSettingWindowId {
	return UserIdSettingWindowId{
		UserId:           userId,
		WindowsSettingId: windowsSettingId,
	}
}

// ParseUserIdSettingWindowID parses 'input' into a UserIdSettingWindowId
func ParseUserIdSettingWindowID(input string) (*UserIdSettingWindowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSettingWindowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSettingWindowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdSettingWindowIDInsensitively parses 'input' case-insensitively into a UserIdSettingWindowId
// note: this method should only be used for API response data and not user input
func ParseUserIdSettingWindowIDInsensitively(input string) (*UserIdSettingWindowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSettingWindowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSettingWindowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdSettingWindowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.WindowsSettingId, ok = input.Parsed["windowsSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsSettingId", input)
	}

	return nil
}

// ValidateUserIdSettingWindowID checks that 'input' can be parsed as a User Id Setting Window ID
func ValidateUserIdSettingWindowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdSettingWindowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Setting Window ID
func (id UserIdSettingWindowId) ID() string {
	fmtString := "/users/%s/settings/windows/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WindowsSettingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Setting Window ID
func (id UserIdSettingWindowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.StaticSegment("windows", "windows", "windows"),
		resourceids.UserSpecifiedSegment("windowsSettingId", "windowsSettingId"),
	}
}

// String returns a human-readable description of this User Id Setting Window ID
func (id UserIdSettingWindowId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Windows Setting: %q", id.WindowsSettingId),
	}
	return fmt.Sprintf("User Id Setting Window (%s)", strings.Join(components, "\n"))
}
