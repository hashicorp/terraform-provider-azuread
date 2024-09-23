package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatId{}

// UserIdChatId is a struct representing the Resource ID for a User Id Chat
type UserIdChatId struct {
	UserId string
	ChatId string
}

// NewUserIdChatID returns a new UserIdChatId struct
func NewUserIdChatID(userId string, chatId string) UserIdChatId {
	return UserIdChatId{
		UserId: userId,
		ChatId: chatId,
	}
}

// ParseUserIdChatID parses 'input' into a UserIdChatId
func ParseUserIdChatID(input string) (*UserIdChatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIDInsensitively parses 'input' case-insensitively into a UserIdChatId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIDInsensitively(input string) (*UserIdChatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	return nil
}

// ValidateUserIdChatID checks that 'input' can be parsed as a User Id Chat ID
func ValidateUserIdChatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat ID
func (id UserIdChatId) ID() string {
	fmtString := "/users/%s/chats/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat ID
func (id UserIdChatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
	}
}

// String returns a human-readable description of this User Id Chat ID
func (id UserIdChatId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
	}
	return fmt.Sprintf("User Id Chat (%s)", strings.Join(components, "\n"))
}
