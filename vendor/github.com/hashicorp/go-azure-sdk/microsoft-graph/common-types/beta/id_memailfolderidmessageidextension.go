package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdMessageIdExtensionId{}

// MeMailFolderIdMessageIdExtensionId is a struct representing the Resource ID for a Me Mail Folder Id Message Id Extension
type MeMailFolderIdMessageIdExtensionId struct {
	MailFolderId string
	MessageId    string
	ExtensionId  string
}

// NewMeMailFolderIdMessageIdExtensionID returns a new MeMailFolderIdMessageIdExtensionId struct
func NewMeMailFolderIdMessageIdExtensionID(mailFolderId string, messageId string, extensionId string) MeMailFolderIdMessageIdExtensionId {
	return MeMailFolderIdMessageIdExtensionId{
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		ExtensionId:  extensionId,
	}
}

// ParseMeMailFolderIdMessageIdExtensionID parses 'input' into a MeMailFolderIdMessageIdExtensionId
func ParseMeMailFolderIdMessageIdExtensionID(input string) (*MeMailFolderIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdMessageIdExtensionIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdMessageIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdMessageIdExtensionIDInsensitively(input string) (*MeMailFolderIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdMessageIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MessageId, ok = input.Parsed["messageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeMailFolderIdMessageIdExtensionID checks that 'input' can be parsed as a Me Mail Folder Id Message Id Extension ID
func ValidateMeMailFolderIdMessageIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdMessageIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Message Id Extension ID
func (id MeMailFolderIdMessageIdExtensionId) ID() string {
	fmtString := "/me/mailFolders/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Message Id Extension ID
func (id MeMailFolderIdMessageIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Message Id Extension ID
func (id MeMailFolderIdMessageIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Mail Folder Id Message Id Extension (%s)", strings.Join(components, "\n"))
}
