package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdUserConfigurationId{}

// UserIdMailFolderIdUserConfigurationId is a struct representing the Resource ID for a User Id Mail Folder Id User Configuration
type UserIdMailFolderIdUserConfigurationId struct {
	UserId              string
	MailFolderId        string
	UserConfigurationId string
}

// NewUserIdMailFolderIdUserConfigurationID returns a new UserIdMailFolderIdUserConfigurationId struct
func NewUserIdMailFolderIdUserConfigurationID(userId string, mailFolderId string, userConfigurationId string) UserIdMailFolderIdUserConfigurationId {
	return UserIdMailFolderIdUserConfigurationId{
		UserId:              userId,
		MailFolderId:        mailFolderId,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseUserIdMailFolderIdUserConfigurationID parses 'input' into a UserIdMailFolderIdUserConfigurationId
func ParseUserIdMailFolderIdUserConfigurationID(input string) (*UserIdMailFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdMailFolderIdUserConfigurationIDInsensitively parses 'input' case-insensitively into a UserIdMailFolderIdUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseUserIdMailFolderIdUserConfigurationIDInsensitively(input string) (*UserIdMailFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdMailFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdMailFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdMailFolderIdUserConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.UserConfigurationId, ok = input.Parsed["userConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", input)
	}

	return nil
}

// ValidateUserIdMailFolderIdUserConfigurationID checks that 'input' can be parsed as a User Id Mail Folder Id User Configuration ID
func ValidateUserIdMailFolderIdUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdMailFolderIdUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Mail Folder Id User Configuration ID
func (id UserIdMailFolderIdUserConfigurationId) ID() string {
	fmtString := "/users/%s/mailFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MailFolderId, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Mail Folder Id User Configuration ID
func (id UserIdMailFolderIdUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("userConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationId"),
	}
}

// String returns a human-readable description of this User Id Mail Folder Id User Configuration ID
func (id UserIdMailFolderIdUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("User Id Mail Folder Id User Configuration (%s)", strings.Join(components, "\n"))
}
