package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdMessageIdMentionId{}

// MeMailFolderIdChildFolderIdMessageIdMentionId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id Message Id Mention
type MeMailFolderIdChildFolderIdMessageIdMentionId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	MentionId     string
}

// NewMeMailFolderIdChildFolderIdMessageIdMentionID returns a new MeMailFolderIdChildFolderIdMessageIdMentionId struct
func NewMeMailFolderIdChildFolderIdMessageIdMentionID(mailFolderId string, mailFolderId1 string, messageId string, mentionId string) MeMailFolderIdChildFolderIdMessageIdMentionId {
	return MeMailFolderIdChildFolderIdMessageIdMentionId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		MentionId:     mentionId,
	}
}

// ParseMeMailFolderIdChildFolderIdMessageIdMentionID parses 'input' into a MeMailFolderIdChildFolderIdMessageIdMentionId
func ParseMeMailFolderIdChildFolderIdMessageIdMentionID(input string) (*MeMailFolderIdChildFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdMessageIdMentionIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdMessageIdMentionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdMessageIdMentionIDInsensitively(input string) (*MeMailFolderIdChildFolderIdMessageIdMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageIdMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageIdMentionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdMessageIdMentionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderId1, ok = input.Parsed["mailFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.MentionId, ok = input.Parsed["mentionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mentionId", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderIdMessageIdMentionID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id Message Id Mention ID
func ValidateMeMailFolderIdChildFolderIdMessageIdMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdMessageIdMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id Message Id Mention ID
func (id MeMailFolderIdChildFolderIdMessageIdMentionId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id Message Id Mention ID
func (id MeMailFolderIdChildFolderIdMessageIdMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("mentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id Message Id Mention ID
func (id MeMailFolderIdChildFolderIdMessageIdMentionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id Message Id Mention (%s)", strings.Join(components, "\n"))
}
