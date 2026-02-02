package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMessageIdReplyIdHostedContentId{}

// UserIdChatIdMessageIdReplyIdHostedContentId is a struct representing the Resource ID for a User Id Chat Id Message Id Reply Id Hosted Content
type UserIdChatIdMessageIdReplyIdHostedContentId struct {
	UserId                     string
	ChatId                     string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewUserIdChatIdMessageIdReplyIdHostedContentID returns a new UserIdChatIdMessageIdReplyIdHostedContentId struct
func NewUserIdChatIdMessageIdReplyIdHostedContentID(userId string, chatId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) UserIdChatIdMessageIdReplyIdHostedContentId {
	return UserIdChatIdMessageIdReplyIdHostedContentId{
		UserId:                     userId,
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserIdChatIdMessageIdReplyIdHostedContentID parses 'input' into a UserIdChatIdMessageIdReplyIdHostedContentId
func ParseUserIdChatIdMessageIdReplyIdHostedContentID(input string) (*UserIdChatIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdMessageIdReplyIdHostedContentIDInsensitively parses 'input' case-insensitively into a UserIdChatIdMessageIdReplyIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdMessageIdReplyIdHostedContentIDInsensitively(input string) (*UserIdChatIdMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdMessageIdReplyIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageId1, ok = input.Parsed["chatMessageId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", input)
	}

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateUserIdChatIdMessageIdReplyIdHostedContentID checks that 'input' can be parsed as a User Id Chat Id Message Id Reply Id Hosted Content ID
func ValidateUserIdChatIdMessageIdReplyIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdMessageIdReplyIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Message Id Reply Id Hosted Content ID
func (id UserIdChatIdMessageIdReplyIdHostedContentId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Message Id Reply Id Hosted Content ID
func (id UserIdChatIdMessageIdReplyIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Message Id Reply Id Hosted Content ID
func (id UserIdChatIdMessageIdReplyIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Id Chat Id Message Id Reply Id Hosted Content (%s)", strings.Join(components, "\n"))
}
