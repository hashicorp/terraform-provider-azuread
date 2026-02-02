package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSettingWindowIdInstanceId{}

// UserIdSettingWindowIdInstanceId is a struct representing the Resource ID for a User Id Setting Window Id Instance
type UserIdSettingWindowIdInstanceId struct {
	UserId                   string
	WindowsSettingId         string
	WindowsSettingInstanceId string
}

// NewUserIdSettingWindowIdInstanceID returns a new UserIdSettingWindowIdInstanceId struct
func NewUserIdSettingWindowIdInstanceID(userId string, windowsSettingId string, windowsSettingInstanceId string) UserIdSettingWindowIdInstanceId {
	return UserIdSettingWindowIdInstanceId{
		UserId:                   userId,
		WindowsSettingId:         windowsSettingId,
		WindowsSettingInstanceId: windowsSettingInstanceId,
	}
}

// ParseUserIdSettingWindowIdInstanceID parses 'input' into a UserIdSettingWindowIdInstanceId
func ParseUserIdSettingWindowIdInstanceID(input string) (*UserIdSettingWindowIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSettingWindowIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSettingWindowIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdSettingWindowIdInstanceIDInsensitively parses 'input' case-insensitively into a UserIdSettingWindowIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdSettingWindowIdInstanceIDInsensitively(input string) (*UserIdSettingWindowIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSettingWindowIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSettingWindowIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdSettingWindowIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.WindowsSettingId, ok = input.Parsed["windowsSettingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsSettingId", input)
	}

	if id.WindowsSettingInstanceId, ok = input.Parsed["windowsSettingInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "windowsSettingInstanceId", input)
	}

	return nil
}

// ValidateUserIdSettingWindowIdInstanceID checks that 'input' can be parsed as a User Id Setting Window Id Instance ID
func ValidateUserIdSettingWindowIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdSettingWindowIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Setting Window Id Instance ID
func (id UserIdSettingWindowIdInstanceId) ID() string {
	fmtString := "/users/%s/settings/windows/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WindowsSettingId, id.WindowsSettingInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Setting Window Id Instance ID
func (id UserIdSettingWindowIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("settings", "settings", "settings"),
		resourceids.StaticSegment("windows", "windows", "windows"),
		resourceids.UserSpecifiedSegment("windowsSettingId", "windowsSettingId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("windowsSettingInstanceId", "windowsSettingInstanceId"),
	}
}

// String returns a human-readable description of this User Id Setting Window Id Instance ID
func (id UserIdSettingWindowIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Windows Setting: %q", id.WindowsSettingId),
		fmt.Sprintf("Windows Setting Instance: %q", id.WindowsSettingInstanceId),
	}
	return fmt.Sprintf("User Id Setting Window Id Instance (%s)", strings.Join(components, "\n"))
}
