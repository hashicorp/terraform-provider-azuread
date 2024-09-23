package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingId{}

// UserIdOnlineMeetingId is a struct representing the Resource ID for a User Id Online Meeting
type UserIdOnlineMeetingId struct {
	UserId          string
	OnlineMeetingId string
}

// NewUserIdOnlineMeetingID returns a new UserIdOnlineMeetingId struct
func NewUserIdOnlineMeetingID(userId string, onlineMeetingId string) UserIdOnlineMeetingId {
	return UserIdOnlineMeetingId{
		UserId:          userId,
		OnlineMeetingId: onlineMeetingId,
	}
}

// ParseUserIdOnlineMeetingID parses 'input' into a UserIdOnlineMeetingId
func ParseUserIdOnlineMeetingID(input string) (*UserIdOnlineMeetingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIDInsensitively(input string) (*UserIdOnlineMeetingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingID checks that 'input' can be parsed as a User Id Online Meeting ID
func ValidateUserIdOnlineMeetingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting ID
func (id UserIdOnlineMeetingId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting ID
func (id UserIdOnlineMeetingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting ID
func (id UserIdOnlineMeetingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
	}
	return fmt.Sprintf("User Id Online Meeting (%s)", strings.Join(components, "\n"))
}
