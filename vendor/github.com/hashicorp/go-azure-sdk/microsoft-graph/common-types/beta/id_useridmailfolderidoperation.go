package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdOperationId{}

// UserIdMailFolderIdOperationId is a struct representing the Resource ID for a User Id Mail Folder Id Operation
type UserIdMailFolderIdOperationId struct {
	UserId                string
	MailFolderId          string
	MailFolderOperationId string
}

// NewUserIdMailFolderIdOperationID returns a new UserIdMailFolderIdOperationId struct
func NewUserIdMailFolderIdOperationID(userId string, mailFolderId string, mailFolderOperationId string) UserIdMailFolderIdOperationId {
	return UserIdMailFolderIdOperationId{
		UserId:                userId,
		MailFolderId:          mailFolderId,
		MailFolderOperationId: mailFolderOperationId,
	}
}

// ParseUserIdMailFolderIdOperationID parses 'input' into a UserIdMailFolderIdOperationId
func ParseUserIdMailFolderIdOperationID(input string) (*UserIdMailFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdOperationIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdOperationIDInsensitively(input string) (*UserIdMailFolderIdOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.MailFolderOperationId, ok = input.Parsed["mailFolderOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderOperationId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdOperationID checks that 'input' can be parsed as a User Id Mail Folder Id Operation ID
func ValidateUserIdMailFolderIdOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Operation ID
func (id UserIdMailFolderIdOperationId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Operation ID
func (id UserIdMailFolderIdOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("mailFolderOperationId", "mailFolderOperationId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Operation ID
func (id UserIdMailFolderIdOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Operation: %q", id.MailFolderOperationId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Operation (%s)", strings.Join(components, "\n"))
}
