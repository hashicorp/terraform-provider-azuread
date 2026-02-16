package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdMessageId{}

// MeMailFolderIdMessageId is a struct representing the Resource ID for a Me Mail Folder Id Message
type MeMailFolderIdMessageId struct {
	MailFolderId string
	MessageId    string
}

// NewMeMailFolderIdMessageID returns a new MeMailFolderIdMessageId struct
func NewMeMailFolderIdMessageID(mailFolderId string, messageId string) MeMailFolderIdMessageId {
	return MeMailFolderIdMessageId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
	}
}

// ParseMeMailFolderIdMessageID parses 'input' into a MeMailFolderIdMessageId
func ParseMeMailFolderIdMessageID(input string) (*MeMailFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdMessageIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdMessageId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdMessageIDInsensitively(input string) (*MeMailFolderIdMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdMessageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	return nil
}

// ValidateMeMailFolderIdMessageID checks that 'input' can be parsed as a Me Mail Folder Id Message ID
func ValidateMeMailFolderIdMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Message ID
func (id MeMailFolderIdMessageId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Message ID
func (id MeMailFolderIdMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Message ID
func (id MeMailFolderIdMessageId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("Me Mail Folder Id Message (%s)", strings.Join(components, "\n"))
}
