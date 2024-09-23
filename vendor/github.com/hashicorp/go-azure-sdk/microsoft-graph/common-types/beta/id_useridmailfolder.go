package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderId{}

// UserIdMailFolderId is a struct representing the Resource ID for a User Id Mail Folder
type UserIdMailFolderId struct {
	UserId       string
	MailFolderId string
}

// NewUserIdMailFolderID returns a new UserIdMailFolderId struct
func NewUserIdMailFolderID(userId string, mailFolderId string) UserIdMailFolderId {
	return UserIdMailFolderId{
		UserId:       userId,
		MailFolderId: mailFolderId,
	}
}

// ParseUserIdMailFolderID parses 'input' into a UserIdMailFolderId
func ParseUserIdMailFolderID(input string) (*UserIdMailFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIDInsensitively(input string) (*UserIdMailFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	return nil
}

// ValidateUserIdMailFolderID checks that 'input' can be parsed as a User Id Mail Folder ID
func ValidateUserIdMailFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder ID
func (id UserIdMailFolderId) ID() string {
	fmtString := "/users/%s/mailFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder ID
func (id UserIdMailFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder ID
func (id UserIdMailFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
	}
	return fmt.Sprintf("User Id Mail Folder (%s)", strings.Join(components, "\n"))
}
