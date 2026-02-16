package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMessageIdAttachmentId{}

// MeMessageIdAttachmentId is a struct representing the Resource ID for a Me Message Id Attachment
type MeMessageIdAttachmentId struct {
	MessageId    string
	AttachmentId string
}

// NewMeMessageIdAttachmentID returns a new MeMessageIdAttachmentId struct
func NewMeMessageIdAttachmentID(messageId string, attachmentId string) MeMessageIdAttachmentId {
	return MeMessageIdAttachmentId{
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseMeMessageIdAttachmentID parses 'input' into a MeMessageIdAttachmentId
func ParseMeMessageIdAttachmentID(input string) (*MeMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMessageIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeMessageIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeMessageIdAttachmentIDInsensitively(input string) (*MeMessageIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMessageIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMessageIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMessageIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeMessageIdAttachmentID checks that 'input' can be parsed as a Me Message Id Attachment ID
func ValidateMeMessageIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message Id Attachment ID
func (id MeMessageIdAttachmentId) ID() string {
	fmtString := "/me/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message Id Attachment ID
func (id MeMessageIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentId"),
	}
}

// String returns a human-readable description of this Me Message Id Attachment ID
func (id MeMessageIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Message Id Attachment (%s)", strings.Join(components, "\n"))
}
