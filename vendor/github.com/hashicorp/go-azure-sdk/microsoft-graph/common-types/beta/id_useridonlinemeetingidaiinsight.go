package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdAiInsightId{}

// UserIdOnlineMeetingIdAiInsightId is a struct representing the Resource ID for a User Id Online Meeting Id Ai Insight
type UserIdOnlineMeetingIdAiInsightId struct {
	UserId          string
	OnlineMeetingId string
	CallAiInsightId string
}

// NewUserIdOnlineMeetingIdAiInsightID returns a new UserIdOnlineMeetingIdAiInsightId struct
func NewUserIdOnlineMeetingIdAiInsightID(userId string, onlineMeetingId string, callAiInsightId string) UserIdOnlineMeetingIdAiInsightId {
	return UserIdOnlineMeetingIdAiInsightId{
		UserId:          userId,
		OnlineMeetingId: onlineMeetingId,
		CallAiInsightId: callAiInsightId,
	}
}

// ParseUserIdOnlineMeetingIdAiInsightID parses 'input' into a UserIdOnlineMeetingIdAiInsightId
func ParseUserIdOnlineMeetingIdAiInsightID(input string) (*UserIdOnlineMeetingIdAiInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdAiInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdAiInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdAiInsightIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdAiInsightId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdAiInsightIDInsensitively(input string) (*UserIdOnlineMeetingIdAiInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdAiInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdAiInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdAiInsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.CallAiInsightId, ok = input.Parsed["callAiInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "callAiInsightId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingIdAiInsightID checks that 'input' can be parsed as a User Id Online Meeting Id Ai Insight ID
func ValidateUserIdOnlineMeetingIdAiInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdAiInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Ai Insight ID
func (id UserIdOnlineMeetingIdAiInsightId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/aiInsights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.CallAiInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Ai Insight ID
func (id UserIdOnlineMeetingIdAiInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("aiInsights", "aiInsights", "aiInsights"),
		resourceids.UserSpecifiedSegment("callAiInsightId", "callAiInsightId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Ai Insight ID
func (id UserIdOnlineMeetingIdAiInsightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Call Ai Insight: %q", id.CallAiInsightId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Ai Insight (%s)", strings.Join(components, "\n"))
}
