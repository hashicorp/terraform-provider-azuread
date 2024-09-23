package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdTabId{}

// UserIdChatIdTabId is a struct representing the Resource ID for a User Id Chat Id Tab
type UserIdChatIdTabId struct {
	UserId     string
	ChatId     string
	TeamsTabId string
}

// NewUserIdChatIdTabID returns a new UserIdChatIdTabId struct
func NewUserIdChatIdTabID(userId string, chatId string, teamsTabId string) UserIdChatIdTabId {
	return UserIdChatIdTabId{
		UserId:     userId,
		ChatId:     chatId,
		TeamsTabId: teamsTabId,
	}
}

// ParseUserIdChatIdTabID parses 'input' into a UserIdChatIdTabId
func ParseUserIdChatIdTabID(input string) (*UserIdChatIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdTabIDInsensitively parses 'input' case-insensitively into a UserIdChatIdTabId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdTabIDInsensitively(input string) (*UserIdChatIdTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateUserIdChatIdTabID checks that 'input' can be parsed as a User Id Chat Id Tab ID
func ValidateUserIdChatIdTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Tab ID
func (id UserIdChatIdTabId) ID() string {
	fmtString := "/users/%s/chats/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Tab ID
func (id UserIdChatIdTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Tab ID
func (id UserIdChatIdTabId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("User Id Chat Id Tab (%s)", strings.Join(components, "\n"))
}
