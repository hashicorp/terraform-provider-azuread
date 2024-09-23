package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdContactFolderIdContactIdExtensionId{}

// UserIdContactFolderIdContactIdExtensionId is a struct representing the Resource ID for a User Id Contact Folder Id Contact Id Extension
type UserIdContactFolderIdContactIdExtensionId struct {
	UserId          string
	ContactFolderId string
	ContactId       string
	ExtensionId     string
}

// NewUserIdContactFolderIdContactIdExtensionID returns a new UserIdContactFolderIdContactIdExtensionId struct
func NewUserIdContactFolderIdContactIdExtensionID(userId string, contactFolderId string, contactId string, extensionId string) UserIdContactFolderIdContactIdExtensionId {
	return UserIdContactFolderIdContactIdExtensionId{
		UserId:          userId,
		ContactFolderId: contactFolderId,
		ContactId:       contactId,
		ExtensionId:     extensionId,
	}
}

// ParseUserIdContactFolderIdContactIdExtensionID parses 'input' into a UserIdContactFolderIdContactIdExtensionId
func ParseUserIdContactFolderIdContactIdExtensionID(input string) (*UserIdContactFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdContactFolderIdContactIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdContactFolderIdContactIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdContactFolderIdContactIdExtensionIDInsensitively(input string) (*UserIdContactFolderIdContactIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdContactFolderIdContactIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdContactFolderIdContactIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdContactFolderIdContactIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdContactFolderIdContactIdExtensionID checks that 'input' can be parsed as a User Id Contact Folder Id Contact Id Extension ID
func ValidateUserIdContactFolderIdContactIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdContactFolderIdContactIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Contact Folder Id Contact Id Extension ID
func (id UserIdContactFolderIdContactIdExtensionId) ID() string {
	fmtString := "/users/%s/contactFolders/%s/contacts/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ContactFolderId, id.ContactId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Contact Folder Id Contact Id Extension ID
func (id UserIdContactFolderIdContactIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("contactFolders", "contactFolders", "contactFolders"),
		resourceids.UserSpecifiedSegment("contactFolderId", "contactFolderId"),
		resourceids.StaticSegment("contacts", "contacts", "contacts"),
		resourceids.UserSpecifiedSegment("contactId", "contactId"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this User Id Contact Folder Id Contact Id Extension ID
func (id UserIdContactFolderIdContactIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Contact Folder: %q", id.ContactFolderId),
		fmt.Sprintf("Contact: %q", id.ContactId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Contact Folder Id Contact Id Extension (%s)", strings.Join(components, "\n"))
}
