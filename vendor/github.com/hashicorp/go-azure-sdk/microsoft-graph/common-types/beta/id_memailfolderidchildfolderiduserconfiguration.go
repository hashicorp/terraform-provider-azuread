package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdChildFolderIdUserConfigurationId{}

// MeMailFolderIdChildFolderIdUserConfigurationId is a struct representing the Resource ID for a Me Mail Folder Id Child Folder Id User Configuration
type MeMailFolderIdChildFolderIdUserConfigurationId struct {
	MailFolderId        string
	MailFolderId1       string
	UserConfigurationId string
}

// NewMeMailFolderIdChildFolderIdUserConfigurationID returns a new MeMailFolderIdChildFolderIdUserConfigurationId struct
func NewMeMailFolderIdChildFolderIdUserConfigurationID(mailFolderId string, mailFolderId1 string, userConfigurationId string) MeMailFolderIdChildFolderIdUserConfigurationId {
	return MeMailFolderIdChildFolderIdUserConfigurationId{
		MailFolderId:        mailFolderId,
		MailFolderId1:       mailFolderId1,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseMeMailFolderIdChildFolderIdUserConfigurationID parses 'input' into a MeMailFolderIdChildFolderIdUserConfigurationId
func ParseMeMailFolderIdChildFolderIdUserConfigurationID(input string) (*MeMailFolderIdChildFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdChildFolderIdUserConfigurationIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdChildFolderIdUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdChildFolderIdUserConfigurationIDInsensitively(input string) (*MeMailFolderIdChildFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdChildFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdChildFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdChildFolderIdUserConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeMailFolderIdChildFolderIdUserConfigurationID checks that 'input' can be parsed as a Me Mail Folder Id Child Folder Id User Configuration ID
func ValidateMeMailFolderIdChildFolderIdUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdChildFolderIdUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id Child Folder Id User Configuration ID
func (id MeMailFolderIdChildFolderIdUserConfigurationId) ID() string {
	fmtString := "/me/mailFolders/%s/childFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.MailFolderId1, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id Child Folder Id User Configuration ID
func (id MeMailFolderIdChildFolderIdUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("childFolders", "childFolders", "childFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId1", "mailFolderId1"),
		resourceids.StaticSegment("userConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id Child Folder Id User Configuration ID
func (id MeMailFolderIdChildFolderIdUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("Mail Folder Id 1: %q", id.MailFolderId1),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("Me Mail Folder Id Child Folder Id User Configuration (%s)", strings.Join(components, "\n"))
}
