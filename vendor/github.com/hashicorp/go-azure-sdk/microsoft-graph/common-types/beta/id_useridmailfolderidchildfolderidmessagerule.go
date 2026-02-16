package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdMessageRuleId{}

// UserIdMailFolderIdChildFolderIdMessageRuleId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder Id Message Rule
type UserIdMailFolderIdChildFolderIdMessageRuleId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
	MessageRuleId string
}

// NewUserIdMailFolderIdChildFolderIdMessageRuleID returns a new UserIdMailFolderIdChildFolderIdMessageRuleId struct
func NewUserIdMailFolderIdChildFolderIdMessageRuleID(userId string, mailFolderId string, mailFolderId1 string, messageRuleId string) UserIdMailFolderIdChildFolderIdMessageRuleId {
	return UserIdMailFolderIdChildFolderIdMessageRuleId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
		MessageRuleId: messageRuleId,
	}
}

// ParseUserIdMailFolderIdChildFolderIdMessageRuleID parses 'input' into a UserIdMailFolderIdChildFolderIdMessageRuleId
func ParseUserIdMailFolderIdChildFolderIdMessageRuleID(input string) (*UserIdMailFolderIdChildFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIdMessageRuleIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderIdMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIdMessageRuleIDInsensitively(input string) (*UserIdMailFolderIdChildFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderIdMessageRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdMailFolderIdChildFolderIdMessageRuleID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder Id Message Rule ID
func ValidateUserIdMailFolderIdChildFolderIdMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderIdMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder Id Message Rule ID
func (id UserIdMailFolderIdChildFolderIdMessageRuleId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder Id Message Rule ID
func (id UserIdMailFolderIdChildFolderIdMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("messageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Child Folder Id Message Rule ID
func (id UserIdMailFolderIdChildFolderIdMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder Id Message Rule (%s)", strings.Join(components, "\n"))
}
