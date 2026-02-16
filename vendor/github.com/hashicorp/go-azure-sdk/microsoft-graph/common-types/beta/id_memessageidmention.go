package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMessageIdMentionId{}

// MeMessageIdMentionId is a struct representing the Resource ID for a Me Message Id Mention
type MeMessageIdMentionId struct {
	MessageId string
	MentionId string
}

// NewMeMessageIdMentionID returns a new MeMessageIdMentionId struct
func NewMeMessageIdMentionID(messageId string, mentionId string) MeMessageIdMentionId {
	return MeMessageIdMentionId{
		MessageId: messageId,
		MentionId: mentionId,
	}
}

// ParseMeMessageIdMentionID parses 'input' into a MeMessageIdMentionId
func ParseMeMessageIdMentionID(input string) (*MeMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMessageIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMessageIdMentionIDInsensitively parses 'input' case-insensitively into a MeMessageIdMentionId
// note: this method should only be used for API response data and not user input
func ParseMeMessageIdMentionIDInsensitively(input string) (*MeMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMessageIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMessageIdMentionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateMeMessageIdMentionID checks that 'input' can be parsed as a Me Message Id Mention ID
func ValidateMeMessageIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message Id Mention ID
func (id MeMessageIdMentionId) ID() string {
	fmtString := "/me/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message Id Mention ID
func (id MeMessageIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this Me Message Id Mention ID
func (id MeMessageIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Me Message Id Mention (%s)", strings.Join(components, "\n"))
}
