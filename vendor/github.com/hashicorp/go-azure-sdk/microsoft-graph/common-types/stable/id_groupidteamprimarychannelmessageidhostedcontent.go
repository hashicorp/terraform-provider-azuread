package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelMessageIdHostedContentId{}

// GroupIdTeamPrimaryChannelMessageIdHostedContentId is a struct representing the Resource ID for a Group Id Team Primary Channel Message Id Hosted Content
type GroupIdTeamPrimaryChannelMessageIdHostedContentId struct {
	GroupId                    string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewGroupIdTeamPrimaryChannelMessageIdHostedContentID returns a new GroupIdTeamPrimaryChannelMessageIdHostedContentId struct
func NewGroupIdTeamPrimaryChannelMessageIdHostedContentID(groupId string, chatMessageId string, chatMessageHostedContentId string) GroupIdTeamPrimaryChannelMessageIdHostedContentId {
	return GroupIdTeamPrimaryChannelMessageIdHostedContentId{
		GroupId:                    groupId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupIdTeamPrimaryChannelMessageIdHostedContentID parses 'input' into a GroupIdTeamPrimaryChannelMessageIdHostedContentId
func ParseGroupIdTeamPrimaryChannelMessageIdHostedContentID(input string) (*GroupIdTeamPrimaryChannelMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelMessageIdHostedContentIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelMessageIdHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelMessageIdHostedContentIDInsensitively(input string) (*GroupIdTeamPrimaryChannelMessageIdHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelMessageIdHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelMessageIdHostedContentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelMessageIdHostedContentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ChatMessageId, ok = input.Parsed["chatMessageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", input)
	}

	if id.ChatMessageHostedContentId, ok = input.Parsed["chatMessageHostedContentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelMessageIdHostedContentID checks that 'input' can be parsed as a Group Id Team Primary Channel Message Id Hosted Content ID
func ValidateGroupIdTeamPrimaryChannelMessageIdHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelMessageIdHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Message Id Hosted Content ID
func (id GroupIdTeamPrimaryChannelMessageIdHostedContentId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Message Id Hosted Content ID
func (id GroupIdTeamPrimaryChannelMessageIdHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageId"),
		resourceids.StaticSegment("hostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Message Id Hosted Content ID
func (id GroupIdTeamPrimaryChannelMessageIdHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Message Id Hosted Content (%s)", strings.Join(components, "\n"))
}
