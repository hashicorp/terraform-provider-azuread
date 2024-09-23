package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdUserConfigurationId{}

// MeMailFolderIdUserConfigurationId is a struct representing the Resource ID for a Me Mail Folder Id User Configuration
type MeMailFolderIdUserConfigurationId struct {
	MailFolderId        string
	UserConfigurationId string
}

// NewMeMailFolderIdUserConfigurationID returns a new MeMailFolderIdUserConfigurationId struct
func NewMeMailFolderIdUserConfigurationID(mailFolderId string, userConfigurationId string) MeMailFolderIdUserConfigurationId {
	return MeMailFolderIdUserConfigurationId{
		MailFolderId:        mailFolderId,
		UserConfigurationId: userConfigurationId,
	}
}

// ParseMeMailFolderIdUserConfigurationID parses 'input' into a MeMailFolderIdUserConfigurationId
func ParseMeMailFolderIdUserConfigurationID(input string) (*MeMailFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeMailFolderIdUserConfigurationIDInsensitively parses 'input' case-insensitively into a MeMailFolderIdUserConfigurationId
// note: this method should only be used for API response data and not user input
func ParseMeMailFolderIdUserConfigurationIDInsensitively(input string) (*MeMailFolderIdUserConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeMailFolderIdUserConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeMailFolderIdUserConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeMailFolderIdUserConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MailFolderId, ok = input.Parsed["mailFolderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mailFolderId", input)
	}

	if id.UserConfigurationId, ok = input.Parsed["userConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userConfigurationId", input)
	}

	return nil
}

// ValidateMeMailFolderIdUserConfigurationID checks that 'input' can be parsed as a Me Mail Folder Id User Configuration ID
func ValidateMeMailFolderIdUserConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMailFolderIdUserConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mail Folder Id User Configuration ID
func (id MeMailFolderIdUserConfigurationId) ID() string {
	fmtString := "/me/mailFolders/%s/userConfigurations/%s"
	return fmt.Sprintf(fmtString, id.MailFolderId, id.UserConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mail Folder Id User Configuration ID
func (id MeMailFolderIdUserConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("mailFolders", "mailFolders", "mailFolders"),
		resourceids.UserSpecifiedSegment("mailFolderId", "mailFolderId"),
		resourceids.StaticSegment("userConfigurations", "userConfigurations", "userConfigurations"),
		resourceids.UserSpecifiedSegment("userConfigurationId", "userConfigurationId"),
	}
}

// String returns a human-readable description of this Me Mail Folder Id User Configuration ID
func (id MeMailFolderIdUserConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Mail Folder: %q", id.MailFolderId),
		fmt.Sprintf("User Configuration: %q", id.UserConfigurationId),
	}
	return fmt.Sprintf("Me Mail Folder Id User Configuration (%s)", strings.Join(components, "\n"))
}
