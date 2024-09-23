package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId{}

// GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId is a struct representing the Resource ID for a Group Id Team Primary Channel Message Id Reply Id Hosted Content
type GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId struct {
	GroupId                    string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID returns a new GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId struct
func NewGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID(groupId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId {
	return GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId{
		GroupId:                    groupId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID parses 'input' into a GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId
func ParseGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID(input string) (*GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentIDInsensitively(input string) (*GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
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

// ValidateGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID checks that 'input' can be parsed as a Group Id Team Primary Channel Message Id Reply Id Hosted Content ID
func ValidateGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Message Id Reply Id Hosted Content ID
func (id GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Message Id Reply Id Hosted Content ID
func (id GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("replies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Message Id Reply Id Hosted Content ID
func (id GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Message Id Reply Id Hosted Content (%s)", strings.Join(components, "\n"))
}
