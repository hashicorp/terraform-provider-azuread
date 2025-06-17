package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdOperationId{}

// UserIdMailFolderIdChildFolderIdOperationId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder Id Operation
type UserIdMailFolderIdChildFolderIdOperationId struct {
	UserId                string
	MailFolderId          string
	MailFolderId1         string
	MailFolderOperationId string
}

// NewUserIdMailFolderIdChildFolderIdOperationID returns a new UserIdMailFolderIdChildFolderIdOperationId struct
func NewUserIdMailFolderIdChildFolderIdOperationID(userId string, mailFolderId string, mailFolderId1 string, mailFolderOperationId string) UserIdMailFolderIdChildFolderIdOperationId {
	return UserIdMailFolderIdChildFolderIdOperationId{
		UserId:                userId,
		MailFolderId:          mailFolderId,
		MailFolderId1:         mailFolderId1,
		MailFolderOperationId: mailFolderOperationId,
	}
}

// ParseUserIdMailFolderIdChildFolderIdOperationID parses 'input' into a UserIdMailFolderIdChildFolderIdOperationId
func ParseUserIdMailFolderIdChildFolderIdOperationID(input string) (*UserIdMailFolderIdChildFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIdOperationIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderIdOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIdOperationIDInsensitively(input string) (*UserIdMailFolderIdChildFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderIdOperationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.MailFolderOperationId, ok = input.Parsed["mailFolderOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderOperationId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdChildFolderIdOperationID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder Id Operation ID
func ValidateUserIdMailFolderIdChildFolderIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder Id Operation ID
func (id UserIdMailFolderIdChildFolderIdOperationId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.MailFolderOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder Id Operation ID
func (id UserIdMailFolderIdChildFolderIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("mailFolderOperationId", "mailFolderOperationId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Child Folder Id Operation ID
func (id UserIdMailFolderIdChildFolderIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("Mail Folder Operation: %q", id.MailFolderOperationId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder Id Operation (%s)", strings.Join(components, "\n"))
}
