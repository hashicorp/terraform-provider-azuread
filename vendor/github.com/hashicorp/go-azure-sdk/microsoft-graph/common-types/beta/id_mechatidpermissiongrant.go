package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdPermissionGrantId{}

// MeChatIdPermissionGrantId is a struct representing the Resource ID for a Me Chat Id Permission Grant
type MeChatIdPermissionGrantId struct {
	ChatId                            string
	ResourceSpecificPermissionGrantId string
}

// NewMeChatIdPermissionGrantID returns a new MeChatIdPermissionGrantId struct
func NewMeChatIdPermissionGrantID(chatId string, resourceSpecificPermissionGrantId string) MeChatIdPermissionGrantId {
	return MeChatIdPermissionGrantId{
		ChatId:                            chatId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseMeChatIdPermissionGrantID parses 'input' into a MeChatIdPermissionGrantId
func ParseMeChatIdPermissionGrantID(input string) (*MeChatIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdPermissionGrantIDInsensitively parses 'input' case-insensitively into a MeChatIdPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdPermissionGrantIDInsensitively(input string) (*MeChatIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateMeChatIdPermissionGrantID checks that 'input' can be parsed as a Me Chat Id Permission Grant ID
func ValidateMeChatIdPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Permission Grant ID
func (id MeChatIdPermissionGrantId) ID() string {
	fmtString := "/me/chats/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Permission Grant ID
func (id MeChatIdPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this Me Chat Id Permission Grant ID
func (id MeChatIdPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Me Chat Id Permission Grant (%s)", strings.Join(components, "\n"))
}
