package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelMessageId{}

// GroupIdTeamPrimaryChannelMessageId is a struct representing the Resource ID for a Group Id Team Primary Channel Message
type GroupIdTeamPrimaryChannelMessageId struct {
	GroupId       string
	ChatMessageId string
}

// NewGroupIdTeamPrimaryChannelMessageID returns a new GroupIdTeamPrimaryChannelMessageId struct
func NewGroupIdTeamPrimaryChannelMessageID(groupId string, chatMessageId string) GroupIdTeamPrimaryChannelMessageId {
	return GroupIdTeamPrimaryChannelMessageId{
		GroupId:       groupId,
		ChatMessageId: chatMessageId,
	}
}

// ParseGroupIdTeamPrimaryChannelMessageID parses 'input' into a GroupIdTeamPrimaryChannelMessageId
func ParseGroupIdTeamPrimaryChannelMessageID(input string) (*GroupIdTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelMessageIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelMessageIDInsensitively(input string) (*GroupIdTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelMessageID checks that 'input' can be parsed as a Group Id Team Primary Channel Message ID
func ValidateGroupIdTeamPrimaryChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Message ID
func (id GroupIdTeamPrimaryChannelMessageId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Message ID
func (id GroupIdTeamPrimaryChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Message ID
func (id GroupIdTeamPrimaryChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Message (%s)", strings.Join(components, "\n"))
}
