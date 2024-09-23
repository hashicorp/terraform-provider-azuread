package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdMessageIdExtensionId{}

// MeMailFolderIdChildFolderIdMessageIdExtensionId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id Message Id Extension
type MeMailFolderIdChildFolderIdMessageIdExtensionId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageId     string
	ExtensionId   string
}

// NewMeMailFolderIdChildFolderIdMessageIdExtensionID returns a new MeMailFolderIdChildFolderIdMessageIdExtensionId struct
func NewMeMailFolderIdChildFolderIdMessageIdExtensionID(mailFolderId string, mailFolderId1 string, messageId string, extensionId string) MeMailFolderIdChildFolderIdMessageIdExtensionId {
	return MeMailFolderIdChildFolderIdMessageIdExtensionId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageId:     messageId,
		ExtensionId:   extensionId,
	}
}

// ParseMeMailFolderIdChildFolderIdMessageIdExtensionID parses 'input' into a MeMailFolderIdChildFolderIdMessageIdExtensionId
func ParseMeMailFolderIdChildFolderIdMessageIdExtensionID(input string) (*MeMailFolderIdChildFolderIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdMessageIdExtensionIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdMessageIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdMessageIdExtensionIDInsensitively(input string) (*MeMailFolderIdChildFolderIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdMessageIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderIdMessageIdExtensionID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id Message Id Extension ID
func ValidateMeMailFolderIdChildFolderIdMessageIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdMessageIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id Message Id Extension ID
func (id MeMailFolderIdChildFolderIdMessageIdExtensionId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id Message Id Extension ID
func (id MeMailFolderIdChildFolderIdMessageIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id Message Id Extension ID
func (id MeMailFolderIdChildFolderIdMessageIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id Message Id Extension (%s)", strings.Join(components, "\n"))
}
