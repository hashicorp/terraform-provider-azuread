package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdUserConfigurationId{}

// UserIdMailFolderIdChildFolderIdUserConfigurationId is a struct representing the Resource ID for a User Id Mail Folder Id Child Folder Id User Configuration
type UserIdMailFolderIdChildFolderIdUserConfigurationId struct {
	UserId              string
	MailFolderId        string
	MailFolderId1       string
	UserConfigurationId string
}

// NewUserIdMailFolderIdChildFolderIdUserConfigurationID returns a new UserIdMailFolderIdChildFolderIdUserConfigurationId struct
func NewUserIdMailFolderIdChildFolderIdUserConfigurationID(userId string, mailFolderId string, mailFolderId1 string, userConfigurationId string) UserIdMailFolderIdChildFolderIdUserConfigurationId {
	return UserIdMailFolderIdChildFolderIdUserConfigurationId{
		UserId:              userId,
		MailFolderId:        mailFolderId,
		MailFolderId1:       mailFolderId1,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseUserIdMailFolderIdChildFolderIdUserConfigurationID parses 'input' into a UserIdMailFolderIdChildFolderIdUserConfigurationId
func ParseUserIdMailFolderIdChildFolderIdUserConfigurationID(input string) (*UserIdMailFolderIdChildFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdChildFolderIdUserConfigurationIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdChildFolderIdUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdChildFolderIdUserConfigurationIDInsensitively(input string) (*UserIdMailFolderIdChildFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdChildFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdChildFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdChildFolderIdUserConfigurationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.UserConfigurationId, ok = input.Parsed["userConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdChildFolderIdUserConfigurationID checks that 'input' can be parsed as a User Id Mail Folder Id Child Folder Id User Configuration ID
func ValidateUserIdMailFolderIdChildFolderIdUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdChildFolderIdUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id Child Folder Id User Configuration ID
func (id UserIdMailFolderIdChildFolderIdUserConfigurationId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/childFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.MailFolderId1, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id Child Folder Id User Configuration ID
func (id UserIdMailFolderIdChildFolderIdUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("userConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id Child Folder Id User Configuration ID
func (id UserIdMailFolderIdChildFolderIdUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("User Id Mail Folder Id Child Folder Id User Configuration (%s)", strings.Join(components, "\n"))
}
