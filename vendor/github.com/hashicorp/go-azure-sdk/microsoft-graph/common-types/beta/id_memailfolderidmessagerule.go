package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdMessageRuleId{}

// MeMailFolderIdMessageRuleId is a struct representing the Resource ID for a Me Mail Folder Id Message Rule
type MeMailFolderIdMessageRuleId struct {
	MailFolderId  string
	MessageRuleId string
}

// NewMeMailFolderIdMessageRuleID returns a new MeMailFolderIdMessageRuleId struct
func NewMeMailFolderIdMessageRuleID(mailFolderId string, messageRuleId string) MeMailFolderIdMessageRuleId {
	return MeMailFolderIdMessageRuleId{
		MailFolderId:  mailFolderId,
		MessageRuleId: messageRuleId,
	}
}

// ParseMeMailFolderIdMessageRuleID parses 'input' into a MeMailFolderIdMessageRuleId
func ParseMeMailFolderIdMessageRuleID(input string) (*MeMailFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdMessageRuleIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdMessageRuleIDInsensitively(input string) (*MeMailFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdMessageRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MessageRuleId, ok = input.Parsed["messageRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", input)
	}

	return nil
}

// ValidateMeMailFolderIdMessageRuleID checks that 'input' can be parsed as a Me Mail Folder Id Message Rule ID
func ValidateMeMailFolderIdMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Message Rule ID
func (id MeMailFolderIdMessageRuleId) ID() string {
	fmtString := "/me/mailFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Message Rule ID
func (id MeMailFolderIdMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Message Rule ID
func (id MeMailFolderIdMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("Me Mail Folder Id Message Rule (%s)", strings.Join(components, "\n"))
}
