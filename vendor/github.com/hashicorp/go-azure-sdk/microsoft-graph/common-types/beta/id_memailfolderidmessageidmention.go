package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdMessageIdMentionId{}

// MeMailFolderIdMessageIdMentionId is a struct representing the Resource ID for a Me Mail Folder Id Message Id Mention
type MeMailFolderIdMessageIdMentionId struct {
	MailFolderId string
	MessageId    string
	MentionId    string
}

// NewMeMailFolderIdMessageIdMentionID returns a new MeMailFolderIdMessageIdMentionId struct
func NewMeMailFolderIdMessageIdMentionID(mailFolderId string, messageId string, mentionId string) MeMailFolderIdMessageIdMentionId {
	return MeMailFolderIdMessageIdMentionId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		MentionId:    mentionId,
	}
}

// ParseMeMailFolderIdMessageIdMentionID parses 'input' into a MeMailFolderIdMessageIdMentionId
func ParseMeMailFolderIdMessageIdMentionID(input string) (*MeMailFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdMessageIdMentionIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdMessageIdMentionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdMessageIdMentionIDInsensitively(input string) (*MeMailFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdMessageIdMentionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateMeMailFolderIdMessageIdMentionID checks that 'input' can be parsed as a Me Mail Folder Id Message Id Mention ID
func ValidateMeMailFolderIdMessageIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdMessageIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Message Id Mention ID
func (id MeMailFolderIdMessageIdMentionId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Message Id Mention ID
func (id MeMailFolderIdMessageIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Message Id Mention ID
func (id MeMailFolderIdMessageIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Me Mail Folder Id Message Id Mention (%s)", strings.Join(components, "\n"))
}
