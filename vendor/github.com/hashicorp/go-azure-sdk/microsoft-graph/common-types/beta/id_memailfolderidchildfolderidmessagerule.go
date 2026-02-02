package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdMessageRuleId{}

// MeMailFolderIdChildFolderIdMessageRuleId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id Message Rule
type MeMailFolderIdChildFolderIdMessageRuleId struct {
	MailFolderId  string
	MailFolderId1 string
	MessageRuleId string
}

// NewMeMailFolderIdChildFolderIdMessageRuleID returns a new MeMailFolderIdChildFolderIdMessageRuleId struct
func NewMeMailFolderIdChildFolderIdMessageRuleID(mailFolderId string, mailFolderId1 string, messageRuleId string) MeMailFolderIdChildFolderIdMessageRuleId {
	return MeMailFolderIdChildFolderIdMessageRuleId{
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageRuleId: messageRuleId,
	}
}

// ParseMeMailFolderIdChildFolderIdMessageRuleID parses 'input' into a MeMailFolderIdChildFolderIdMessageRuleId
func ParseMeMailFolderIdChildFolderIdMessageRuleID(input string) (*MeMailFolderIdChildFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdMessageRuleIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdMessageRuleIDInsensitively(input string) (*MeMailFolderIdChildFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdMessageRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderId1, ok = input.Parsed["mailFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId1", input)
	}

	if id.MessageRuleId, ok = input.Parsed["messageRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", input)
	}

	return nil
}

// ValidateMeMailFolderIdChildFolderIdMessageRuleID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id Message Rule ID
func ValidateMeMailFolderIdChildFolderIdMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id Message Rule ID
func (id MeMailFolderIdChildFolderIdMessageRuleId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id Message Rule ID
func (id MeMailFolderIdChildFolderIdMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id Message Rule ID
func (id MeMailFolderIdChildFolderIdMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id Message Rule (%s)", strings.Join(components, "\n"))
}
