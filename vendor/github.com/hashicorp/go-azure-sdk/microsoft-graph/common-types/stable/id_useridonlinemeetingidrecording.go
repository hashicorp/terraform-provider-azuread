package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdRecordingId{}

// UserIdOnlineMeetingIdRecordingId is a struct representing the Resource ID for a User Id Online Meeting Id Recording
type UserIdOnlineMeetingIdRecordingId struct {
	UserId          string
	OnlineMeetingId string
	CallRecordingId string
}

// NewUserIdOnlineMeetingIdRecordingID returns a new UserIdOnlineMeetingIdRecordingId struct
func NewUserIdOnlineMeetingIdRecordingID(userId string, onlineMeetingId string, callRecordingId string) UserIdOnlineMeetingIdRecordingId {
	return UserIdOnlineMeetingIdRecordingId{
		UserId:          userId,
		OnlineMeetingId: onlineMeetingId,
		CallRecordingId: callRecordingId,
	}
}

// ParseUserIdOnlineMeetingIdRecordingID parses 'input' into a UserIdOnlineMeetingIdRecordingId
func ParseUserIdOnlineMeetingIdRecordingID(input string) (*UserIdOnlineMeetingIdRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdRecordingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdRecordingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdRecordingIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdRecordingId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdRecordingIDInsensitively(input string) (*UserIdOnlineMeetingIdRecordingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdRecordingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdRecordingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdRecordingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.CallRecordingId, ok = input.Parsed["callRecordingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "callRecordingId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingIdRecordingID checks that 'input' can be parsed as a User Id Online Meeting Id Recording ID
func ValidateUserIdOnlineMeetingIdRecordingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdRecordingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Recording ID
func (id UserIdOnlineMeetingIdRecordingId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/recordings/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.CallRecordingId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Recording ID
func (id UserIdOnlineMeetingIdRecordingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("recordings", "recordings", "recordings"),
		resourceids.UserSpecifiedSegment("callRecordingId", "callRecordingId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Recording ID
func (id UserIdOnlineMeetingIdRecordingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Recording: %q", id.CallRecordingId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Recording (%s)", strings.Join(components, "\n"))
}
