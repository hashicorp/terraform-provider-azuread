package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadId{}

// GroupIdThreadId is a struct representing the Resource ID for a Group Id Thread
type GroupIdThreadId struct {
	GroupId              string
	ConversationThreadId string
}

// NewGroupIdThreadID returns a new GroupIdThreadId struct
func NewGroupIdThreadID(groupId string, conversationThreadId string) GroupIdThreadId {
	return GroupIdThreadId{
		GroupId:              groupId,
		ConversationThreadId: conversationThreadId,
	}
}

// ParseGroupIdThreadID parses 'input' into a GroupIdThreadId
func ParseGroupIdThreadID(input string) (*GroupIdThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdThreadIDInsensitively parses 'input' case-insensitively into a GroupIdThreadId
// note: this method should only be used for API response data and not user input
func ParseGroupIdThreadIDInsensitively(input string) (*GroupIdThreadId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdThreadId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdThreadId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdThreadId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ConversationThreadId, ok = input.Parsed["conversationThreadId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationThreadId", input)
	}

	return nil
}

// ValidateGroupIdThreadID checks that 'input' can be parsed as a Group Id Thread ID
func ValidateGroupIdThreadID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdThreadID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Thread ID
func (id GroupIdThreadId) ID() string {
	fmtString := "/groups/%s/threads/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationThreadId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Thread ID
func (id GroupIdThreadId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("threads", "threads", "threads"),
		resourceids.UserSpecifiedSegment("conversationThreadId", "conversationThreadId"),
	}
}

// String returns a human-readable description of this Group Id Thread ID
func (id GroupIdThreadId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Thread: %q", id.ConversationThreadId),
	}
	return fmt.Sprintf("Group Id Thread (%s)", strings.Join(components, "\n"))
}
