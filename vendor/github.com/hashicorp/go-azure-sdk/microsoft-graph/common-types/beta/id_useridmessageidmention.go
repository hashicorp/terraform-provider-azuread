package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMessageIdMentionId{}

// UserIdMessageIdMentionId is a struct representing the Resource ID for a User Id Message Id Mention
type UserIdMessageIdMentionId struct {
	UserId    string
	MessageId string
	MentionId string
}

// NewUserIdMessageIdMentionID returns a new UserIdMessageIdMentionId struct
func NewUserIdMessageIdMentionID(userId string, messageId string, mentionId string) UserIdMessageIdMentionId {
	return UserIdMessageIdMentionId{
		UserId:    userId,
		MessageId: messageId,
		MentionId: mentionId,
	}
}

// ParseUserIdMessageIdMentionID parses 'input' into a UserIdMessageIdMentionId
func ParseUserIdMessageIdMentionID(input string) (*UserIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMessageIdMentionIDInsensitively parses 'input' case-insensitively into a UserIdMessageIdMentionId
// note: this method should only be used for API response data and not user input
func ParseUserIdMessageIdMentionIDInsensitively(input string) (*UserIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMessageIdMentionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateUserIdMessageIdMentionID checks that 'input' can be parsed as a User Id Message Id Mention ID
func ValidateUserIdMessageIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMessageIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Message Id Mention ID
func (id UserIdMessageIdMentionId) ID() string {
	fmtString := "/users/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Message Id Mention ID
func (id UserIdMessageIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this User Id Message Id Mention ID
func (id UserIdMessageIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("User Id Message Id Mention (%s)", strings.Join(components, "\n"))
}
