package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationId{}

// GroupIdConversationId is a struct representing the Resource ID for a Group Id Conversation
type GroupIdConversationId struct {
	GroupId        string
	ConversationId string
}

// NewGroupIdConversationID returns a new GroupIdConversationId struct
func NewGroupIdConversationID(groupId string, conversationId string) GroupIdConversationId {
	return GroupIdConversationId{
		GroupId:        groupId,
		ConversationId: conversationId,
	}
}

// ParseGroupIdConversationID parses 'input' into a GroupIdConversationId
func ParseGroupIdConversationID(input string) (*GroupIdConversationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdConversationIDInsensitively parses 'input' case-insensitively into a GroupIdConversationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdConversationIDInsensitively(input string) (*GroupIdConversationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdConversationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdConversationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdConversationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationId, ok = input.Parsed["conversationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationId", input)
	}

	return nil
}

// ValidateGroupIdConversationID checks that 'input' can be parsed as a Group Id Conversation ID
func ValidateGroupIdConversationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdConversationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Conversation ID
func (id GroupIdConversationId) ID() string {
	fmtString := "/groups/%s/conversations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Conversation ID
func (id GroupIdConversationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("conversations", "conversations", "conversations"),
		resourceids.UserSpecifiedSegment("conversationId", "conversationId"),
	}
}

// String returns a human-readable description of this Group Id Conversation ID
func (id GroupIdConversationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation: %q", id.ConversationId),
	}
	return fmt.Sprintf("Group Id Conversation (%s)", strings.Join(components, "\n"))
}
