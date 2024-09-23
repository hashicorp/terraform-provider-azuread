package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdContactId{}

// UserIdContactFolderIdContactId is a struct representing the Resource ID for a User Id Contact Folder Id Contact
type UserIdContactFolderIdContactId struct {
	UserId          string
	ContactFolderId string
	ContactId       string
}

// NewUserIdContactFolderIdContactID returns a new UserIdContactFolderIdContactId struct
func NewUserIdContactFolderIdContactID(userId string, contactFolderId string, contactId string) UserIdContactFolderIdContactId {
	return UserIdContactFolderIdContactId{
		UserId:          userId,
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
	}
}

// ParseUserIdContactFolderIdContactID parses 'input' into a UserIdContactFolderIdContactId
func ParseUserIdContactFolderIdContactID(input string) (*UserIdContactFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactFolderIdContactIDInsensitively parses 'input' case-insensitively into a UserIdContactFolderIdContactId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactFolderIdContactIDInsensitively(input string) (*UserIdContactFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactFolderIdContactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ContactFolderId, ok = input.Parsed["contactFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactFolderId", input)
	}

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	return nil
}

// ValidateUserIdContactFolderIdContactID checks that 'input' can be parsed as a User Id Contact Folder Id Contact ID
func ValidateUserIdContactFolderIdContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactFolderIdContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Folder Id Contact ID
func (id UserIdContactFolderIdContactId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Folder Id Contact ID
func (id UserIdContactFolderIdContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
	}
}

// String returns a human-readable description of this User Id Contact Folder Id Contact ID
func (id UserIdContactFolderIdContactId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("User Id Contact Folder Id Contact (%s)", strings.Join(components, "\n"))
}
