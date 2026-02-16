package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageIdExtensionId{}

// UserIdMailFolderIdMessageIdExtensionId is a struct representing the Resource ID for a User Id Mail Folder Id Message Id Extension
type UserIdMailFolderIdMessageIdExtensionId struct {
	UserId       string
	MailFolderId string
	MessageId    string
	ExtensionId  string
}

// NewUserIdMailFolderIdMessageIdExtensionID returns a new UserIdMailFolderIdMessageIdExtensionId struct
func NewUserIdMailFolderIdMessageIdExtensionID(userId string, mailFolderId string, messageId string, extensionId string) UserIdMailFolderIdMessageIdExtensionId {
	return UserIdMailFolderIdMessageIdExtensionId{
		UserId:       userId,
		MailFolderId: mailFolderId,
		MessageId:    messageId,
		ExtensionId:  extensionId,
	}
}

// ParseUserIdMailFolderIdMessageIdExtensionID parses 'input' into a UserIdMailFolderIdMessageIdExtensionId
func ParseUserIdMailFolderIdMessageIdExtensionID(input string) (*UserIdMailFolderIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdMessageIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdMessageIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdMessageIdExtensionIDInsensitively(input string) (*UserIdMailFolderIdMessageIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdMessageIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdMessageIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdMessageIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdMailFolderIdMessageIdExtensionID checks that 'input' can be parsed as a User Id Mail Folder Id Message Id Extension ID
func ValidateUserIdMailFolderIdMessageIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdMessageIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Message Id Extension ID
func (id UserIdMailFolderIdMessageIdExtensionId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Message Id Extension ID
func (id UserIdMailFolderIdMessageIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("messages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Message Id Extension ID
func (id UserIdMailFolderIdMessageIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Message Id Extension (%s)", strings.Join(components, "\n"))
}
