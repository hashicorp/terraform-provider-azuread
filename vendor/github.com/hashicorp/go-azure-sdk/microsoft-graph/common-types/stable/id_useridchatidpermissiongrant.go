package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdPermissionGrantId{}

// UserIdChatIdPermissionGrantId is a struct representing the Resource ID for a User Id Chat Id Permission Grant
type UserIdChatIdPermissionGrantId struct {
	UserId                            string
	ChatId                            string
	ResourceSpecificPermissionGrantId string
}

// NewUserIdChatIdPermissionGrantID returns a new UserIdChatIdPermissionGrantId struct
func NewUserIdChatIdPermissionGrantID(userId string, chatId string, resourceSpecificPermissionGrantId string) UserIdChatIdPermissionGrantId {
	return UserIdChatIdPermissionGrantId{
		UserId:                            userId,
		ChatId:                            chatId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseUserIdChatIdPermissionGrantID parses 'input' into a UserIdChatIdPermissionGrantId
func ParseUserIdChatIdPermissionGrantID(input string) (*UserIdChatIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdPermissionGrantIDInsensitively parses 'input' case-insensitively into a UserIdChatIdPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdPermissionGrantIDInsensitively(input string) (*UserIdChatIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateUserIdChatIdPermissionGrantID checks that 'input' can be parsed as a User Id Chat Id Permission Grant ID
func ValidateUserIdChatIdPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Permission Grant ID
func (id UserIdChatIdPermissionGrantId) ID() string {
	fmtString := "/users/%s/chats/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Permission Grant ID
func (id UserIdChatIdPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Permission Grant ID
func (id UserIdChatIdPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("User Id Chat Id Permission Grant (%s)", strings.Join(components, "\n"))
}
