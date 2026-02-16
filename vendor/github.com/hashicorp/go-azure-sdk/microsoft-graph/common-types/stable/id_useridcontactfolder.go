package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderId{}

// UserIdContactFolderId is a struct representing the Resource ID for a User Id Contact Folder
type UserIdContactFolderId struct {
	UserId          string
	ContactFolderId string
}

// NewUserIdContactFolderID returns a new UserIdContactFolderId struct
func NewUserIdContactFolderID(userId string, contactFolderId string) UserIdContactFolderId {
	return UserIdContactFolderId{
		UserId:          userId,
		ContactFolderId: contactFolderId,
	}
}

// ParseUserIdContactFolderID parses 'input' into a UserIdContactFolderId
func ParseUserIdContactFolderID(input string) (*UserIdContactFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactFolderIDInsensitively parses 'input' case-insensitively into a UserIdContactFolderId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactFolderIDInsensitively(input string) (*UserIdContactFolderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactFolderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	return nil
}

// ValidateUserIdContactFolderID checks that 'input' can be parsed as a User Id Contact Folder ID
func ValidateUserIdContactFolderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactFolderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Folder ID
func (id UserIdContactFolderId) ID() string {
	fmtString := "/users/%s/contactFolders/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Folder ID
func (id UserIdContactFolderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
	}
}

// String returns a human-readable description of this User Id Contact Folder ID
func (id UserIdContactFolderId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
	}
	return fmt.Sprintf("User Id Contact Folder (%s)", strings.Join(components, "\n"))
}
