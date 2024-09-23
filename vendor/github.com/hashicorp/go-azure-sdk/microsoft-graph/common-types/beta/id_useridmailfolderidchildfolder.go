package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderId{}

// UserIdMailFolderIdChildFolderId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder
type UserIdMailFolderIdChildFolderId struct {
	UserId        string
	MailFolderId  string
	MailFolderId1 string
}

// NewUserIdMailFolderIdChildFolderID returns a new UserIdMailFolderIdChildFolderId struct
func NewUserIdMailFolderIdChildFolderID(userId string, mailFolderId string, mailFolderId1 string) UserIdMailFolderIdChildFolderId {
	return UserIdMailFolderIdChildFolderId{
		UserId:        userId,
		MailFolderId:  mailFolderId,
		MailFolderId1: mailFolderId1,
	}
}

// ParseUserIdMailFolderIdChildFolderID parses 'input' into a UserIdMailFolderIdChildFolderId
func ParseUserIdMailFolderIdChildFolderID(input string) (*UserIdMailFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIDInsensitively(input string) (*UserIdMailFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdMailFolderIdChildFolderID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder ID
func ValidateUserIdMailFolderIdChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder ID
func (id UserIdMailFolderIdChildFolderId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder ID
func (id UserIdMailFolderIdChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Child Folder ID
func (id UserIdMailFolderIdChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder (%s)", strings.Join(components, "\n"))
}
