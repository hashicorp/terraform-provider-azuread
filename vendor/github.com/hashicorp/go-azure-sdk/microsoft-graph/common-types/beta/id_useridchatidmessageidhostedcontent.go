package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMessageIdHostedContentId{}

// UserIdChatIdMessageIdHostedContentId is a struct representing the Resource ID for a User Id Chat Id Message Id Hosted Content
type UserIdChatIdMessageIdHostedContentId struct {
	UserId                     string
	ChatId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewUserIdChatIdMessageIdHostedContentID returns a new UserIdChatIdMessageIdHostedContentId struct
func NewUserIdChatIdMessageIdHostedContentID(userId string, chatId string, chatMessageId string, chatMessageHostedContentId string) UserIdChatIdMessageIdHostedContentId {
	return UserIdChatIdMessageIdHostedContentId{
		UserId:                     userId,
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserIdChatIdMessageIdHostedContentID parses 'input' into a UserIdChatIdMessageIdHostedContentId
func ParseUserIdChatIdMessageIdHostedContentID(input string) (*UserIdChatIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a UserIdChatIdMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdMessageIdHostedContentIDInsensitively(input string) (*UserIdChatIdMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateUserIdChatIdMessageIdHostedContentID checks that 'input' can be parsed as a User Id Chat Id Message Id Hosted Content ID
func ValidateUserIdChatIdMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Message Id Hosted Content ID
func (id UserIdChatIdMessageIdHostedContentId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Message Id Hosted Content ID
func (id UserIdChatIdMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Message Id Hosted Content ID
func (id UserIdChatIdMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Id Chat Id Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
