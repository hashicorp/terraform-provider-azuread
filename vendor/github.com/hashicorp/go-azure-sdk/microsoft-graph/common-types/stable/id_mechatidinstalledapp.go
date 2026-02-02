package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdInstalledAppId{}

// MeChatIdInstalledAppId is a struct representing the Resource ID for a Me Chat Id Installed App
type MeChatIdInstalledAppId struct {
	ChatId                 string
	TeamsAppInstallationId string
}

// NewMeChatIdInstalledAppID returns a new MeChatIdInstalledAppId struct
func NewMeChatIdInstalledAppID(chatId string, teamsAppInstallationId string) MeChatIdInstalledAppId {
	return MeChatIdInstalledAppId{
		ChatId:                 chatId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseMeChatIdInstalledAppID parses 'input' into a MeChatIdInstalledAppId
func ParseMeChatIdInstalledAppID(input string) (*MeChatIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeChatIdInstalledAppIDInsensitively parses 'input' case-insensitively into a MeChatIdInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseMeChatIdInstalledAppIDInsensitively(input string) (*MeChatIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeChatIdInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeChatIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeChatIdInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.TeamsAppInstallationId, ok = input.Parsed["teamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", input)
	}

	return nil
}

// ValidateMeChatIdInstalledAppID checks that 'input' can be parsed as a Me Chat Id Installed App ID
func ValidateMeChatIdInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatIdInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Id Installed App ID
func (id MeChatIdInstalledAppId) ID() string {
	fmtString := "/me/chats/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Id Installed App ID
func (id MeChatIdInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationId"),
	}
}

// String returns a human-readable description of this Me Chat Id Installed App ID
func (id MeChatIdInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("Me Chat Id Installed App (%s)", strings.Join(components, "\n"))
}
