package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageRuleId{}

// UserIdMailFolderIdMessageRuleId is a struct representing the Resource ID for a User Id Mail Folder Id Message Rule
type UserIdMailFolderIdMessageRuleId struct {
	UserId        string
	MailFolderId  string
	MessageRuleId string
}

// NewUserIdMailFolderIdMessageRuleID returns a new UserIdMailFolderIdMessageRuleId struct
func NewUserIdMailFolderIdMessageRuleID(userId string, mailFolderId string, messageRuleId string) UserIdMailFolderIdMessageRuleId {
	return UserIdMailFolderIdMessageRuleId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MessageRuleId: messageRuleId,
	}
}

// ParseUserIdMailFolderIdMessageRuleID parses 'input' into a UserIdMailFolderIdMessageRuleId
func ParseUserIdMailFolderIdMessageRuleID(input string) (*UserIdMailFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdMessageRuleIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdMessageRuleId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdMessageRuleIDInsensitively(input string) (*UserIdMailFolderIdMessageRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdMessageRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MessageRuleId, ok = input.Parsed["messageRuleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "messageRuleId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdMessageRuleID checks that 'input' can be parsed as a User Id Mail Folder Id Message Rule ID
func ValidateUserIdMailFolderIdMessageRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdMessageRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Message Rule ID
func (id UserIdMailFolderIdMessageRuleId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messageRules/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Message Rule ID
func (id UserIdMailFolderIdMessageRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messageRules", "messageRules", "messageRules"),
		resourceids.UserSpecifiedSegment("messageRuleId", "messageRuleId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Message Rule ID
func (id UserIdMailFolderIdMessageRuleId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message Rule: %q", id.MessageRuleId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Message Rule (%s)", strings.Join(components, "\n"))
}
