package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdOperationId{}

// UserIdChatIdOperationId is a struct representing the Resource ID for a User Id Chat Id Operation
type UserIdChatIdOperationId struct {
	UserId                string
	ChatId                string
	TeamsAsyncOperationId string
}

// NewUserIdChatIdOperationID returns a new UserIdChatIdOperationId struct
func NewUserIdChatIdOperationID(userId string, chatId string, teamsAsyncOperationId string) UserIdChatIdOperationId {
	return UserIdChatIdOperationId{
		UserId:                userId,
		ChatId:                chatId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseUserIdChatIdOperationID parses 'input' into a UserIdChatIdOperationId
func ParseUserIdChatIdOperationID(input string) (*UserIdChatIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdOperationIDInsensitively parses 'input' case-insensitively into a UserIdChatIdOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdOperationIDInsensitively(input string) (*UserIdChatIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.TeamsAsyncOperationId, ok = input.Parsed["teamsAsyncOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", input)
	}

	return nil
}

// ValidateUserIdChatIdOperationID checks that 'input' can be parsed as a User Id Chat Id Operation ID
func ValidateUserIdChatIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Operation ID
func (id UserIdChatIdOperationId) ID() string {
	fmtString := "/users/%s/chats/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Operation ID
func (id UserIdChatIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Operation ID
func (id UserIdChatIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("User Id Chat Id Operation (%s)", strings.Join(components, "\n"))
}
