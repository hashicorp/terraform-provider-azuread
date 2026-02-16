package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdChildFolderId{}

// UserIdContactFolderIdChildFolderId is a struct representing the Resource ID for a User Id Contact Folder Id Child Folder
type UserIdContactFolderIdChildFolderId struct {
	UserId           string
	ContactFolderId  string
	ContactFolderId1 string
}

// NewUserIdContactFolderIdChildFolderID returns a new UserIdContactFolderIdChildFolderId struct
func NewUserIdContactFolderIdChildFolderID(userId string, contactFolderId string, contactFolderId1 string) UserIdContactFolderIdChildFolderId {
	return UserIdContactFolderIdChildFolderId{
		UserId:           userId,
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
	}
}

// ParseUserIdContactFolderIdChildFolderID parses 'input' into a UserIdContactFolderIdChildFolderId
func ParseUserIdContactFolderIdChildFolderID(input string) (*UserIdContactFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactFolderIdChildFolderIDInsensitively parses 'input' case-insensitively into a UserIdContactFolderIdChildFolderId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactFolderIdChildFolderIDInsensitively(input string) (*UserIdContactFolderIdChildFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdChildFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdChildFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactFolderIdChildFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	if id.ContactFolderId1, ok = input.Parsed["contactFolderId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId1", input)
	}

	return nil
}

// ValidateUserIdContactFolderIdChildFolderID checks that 'input' can be parsed as a User Id Contact Folder Id Child Folder ID
func ValidateUserIdContactFolderIdChildFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactFolderIdChildFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Folder Id Child Folder ID
func (id UserIdContactFolderIdChildFolderId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/childFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactFolderId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Folder Id Child Folder ID
func (id UserIdContactFolderIdChildFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1"),
	}
}

// String returns a human-readable description of this User Id Contact Folder Id Child Folder ID
func (id UserIdContactFolderIdChildFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
	}
	return fmt.Sprintf("User Id Contact Folder Id Child Folder (%s)", strings.Join(components, "\n"))
}
