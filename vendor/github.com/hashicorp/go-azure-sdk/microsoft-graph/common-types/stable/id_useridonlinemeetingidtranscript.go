package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdTranscriptId{}

// UserIdOnlineMeetingIdTranscriptId is a struct representing the Resource ID for a User Id Online Meeting Id Transcript
type UserIdOnlineMeetingIdTranscriptId struct {
	UserId           string
	OnlineMeetingId  string
	CallTranscriptId string
}

// NewUserIdOnlineMeetingIdTranscriptID returns a new UserIdOnlineMeetingIdTranscriptId struct
func NewUserIdOnlineMeetingIdTranscriptID(userId string, onlineMeetingId string, callTranscriptId string) UserIdOnlineMeetingIdTranscriptId {
	return UserIdOnlineMeetingIdTranscriptId{
		UserId:           userId,
		OnlineMeetingId:  onlineMeetingId,
		CallTranscriptId: callTranscriptId,
	}
}

// ParseUserIdOnlineMeetingIdTranscriptID parses 'input' into a UserIdOnlineMeetingIdTranscriptId
func ParseUserIdOnlineMeetingIdTranscriptID(input string) (*UserIdOnlineMeetingIdTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdTranscriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdTranscriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdTranscriptIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdTranscriptId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdTranscriptIDInsensitively(input string) (*UserIdOnlineMeetingIdTranscriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdTranscriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdTranscriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdTranscriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.CallTranscriptId, ok = input.Parsed["callTranscriptId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "callTranscriptId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingIdTranscriptID checks that 'input' can be parsed as a User Id Online Meeting Id Transcript ID
func ValidateUserIdOnlineMeetingIdTranscriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdTranscriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Transcript ID
func (id UserIdOnlineMeetingIdTranscriptId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/transcripts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.CallTranscriptId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Transcript ID
func (id UserIdOnlineMeetingIdTranscriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("transcripts", "transcripts", "transcripts"),
		resourceids.UserSpecifiedSegment("callTranscriptId", "callTranscriptId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Transcript ID
func (id UserIdOnlineMeetingIdTranscriptId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Transcript: %q", id.CallTranscriptId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Transcript (%s)", strings.Join(components, "\n"))
}
