package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdChildFolderIdContactId{}

// UserIdContactFolderIdChildFolderIdContactId is a struct representing the Resource ID for a User Id Contact Folder Id Child Folder Id Contact
type UserIdContactFolderIdChildFolderIdContactId struct {
	UserId           string
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
}

// NewUserIdContactFolderIdChildFolderIdContactID returns a new UserIdContactFolderIdChildFolderIdContactId struct
func NewUserIdContactFolderIdChildFolderIdContactID(userId string, contactFolderId string, contactFolderId1 string, contactId string) UserIdContactFolderIdChildFolderIdContactId {
	return UserIdContactFolderIdChildFolderIdContactId{
		UserId:           userId,
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
	}
}

// ParseUserIdContactFolderIdChildFolderIdContactID parses 'input' into a UserIdContactFolderIdChildFolderIdContactId
func ParseUserIdContactFolderIdChildFolderIdContactID(input string) (*UserIdContactFolderIdChildFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdChildFolderIdContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdChildFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactFolderIdChildFolderIdContactIDInsensitively parses 'input' case-insensitively into a UserIdContactFolderIdChildFolderIdContactId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactFolderIdChildFolderIdContactIDInsensitively(input string) (*UserIdContactFolderIdChildFolderIdContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdChildFolderIdContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdChildFolderIdContactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactFolderIdChildFolderIdContactId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContactId, ok = input.Parsed["contactId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contactId", input)
	}

	return nil
}

// ValidateUserIdContactFolderIdChildFolderIdContactID checks that 'input' can be parsed as a User Id Contact Folder Id Child Folder Id Contact ID
func ValidateUserIdContactFolderIdChildFolderIdContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactFolderIdChildFolderIdContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Folder Id Child Folder Id Contact ID
func (id UserIdContactFolderIdChildFolderIdContactId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/childFolders/%s/contacts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactFolderId1, id.ContactId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Folder Id Child Folder Id Contact ID
func (id UserIdContactFolderIdChildFolderIdContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
	}
}

// String returns a human-readable description of this User Id Contact Folder Id Child Folder Id Contact ID
func (id UserIdContactFolderIdChildFolderIdContactId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
	}
	return fmt.Sprintf("User Id Contact Folder Id Child Folder Id Contact (%s)", strings.Join(components, "\n"))
}
