package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdChildFolderIdContactIdExtensionId{}

// UserIdContactFolderIdChildFolderIdContactIdExtensionId is a struct representing the Resource ID for a User Id Contact Folder Id Child Folder Id Contact Id Extension
type UserIdContactFolderIdChildFolderIdContactIdExtensionId struct {
	UserId           string
	ContactFolderId  string
	ContactFolderId1 string
	ContactId        string
	ExtensionId      string
}

// NewUserIdContactFolderIdChildFolderIdContactIdExtensionID returns a new UserIdContactFolderIdChildFolderIdContactIdExtensionId struct
func NewUserIdContactFolderIdChildFolderIdContactIdExtensionID(userId string, contactFolderId string, contactFolderId1 string, contactId string, extensionId string) UserIdContactFolderIdChildFolderIdContactIdExtensionId {
	return UserIdContactFolderIdChildFolderIdContactIdExtensionId{
		UserId:           userId,
		ContactFolderId:  contactFolderId,
		ContactFolderId1: contactFolderId1,
		ContactId:        contactId,
		ExtensionId:      extensionId,
	}
}

// ParseUserIdContactFolderIdChildFolderIdContactIdExtensionID parses 'input' into a UserIdContactFolderIdChildFolderIdContactIdExtensionId
func ParseUserIdContactFolderIdChildFolderIdContactIdExtensionID(input string) (*UserIdContactFolderIdChildFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdChildFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdChildFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactFolderIdChildFolderIdContactIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdContactFolderIdChildFolderIdContactIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactFolderIdChildFolderIdContactIdExtensionIDInsensitively(input string) (*UserIdContactFolderIdChildFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdChildFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdChildFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactFolderIdChildFolderIdContactIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdContactFolderIdChildFolderIdContactIdExtensionID checks that 'input' can be parsed as a User Id Contact Folder Id Child Folder Id Contact Id Extension ID
func ValidateUserIdContactFolderIdChildFolderIdContactIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactFolderIdChildFolderIdContactIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Folder Id Child Folder Id Contact Id Extension ID
func (id UserIdContactFolderIdChildFolderIdContactIdExtensionId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/childFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactFolderId1, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Folder Id Child Folder Id Contact Id Extension ID
func (id UserIdContactFolderIdChildFolderIdContactIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId1", "contactFolderId1"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Contact Folder Id Child Folder Id Contact Id Extension ID
func (id UserIdContactFolderIdChildFolderIdContactIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact Folder Id 1: %q", id.ContactFolderId1),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Contact Folder Id Child Folder Id Contact Id Extension (%s)", strings.Join(components, "\n"))
}
