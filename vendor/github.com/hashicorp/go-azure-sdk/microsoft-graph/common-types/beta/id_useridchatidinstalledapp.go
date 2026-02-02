package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdInstalledAppId{}

// UserIdChatIdInstalledAppId is a struct representing the Resource ID for a User Id Chat Id Installed App
type UserIdChatIdInstalledAppId struct {
	UserId                 string
	ChatId                 string
	TeamsAppInstallationId string
}

// NewUserIdChatIdInstalledAppID returns a new UserIdChatIdInstalledAppId struct
func NewUserIdChatIdInstalledAppID(userId string, chatId string, teamsAppInstallationId string) UserIdChatIdInstalledAppId {
	return UserIdChatIdInstalledAppId{
		UserId:                 userId,
		ChatId:                 chatId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseUserIdChatIdInstalledAppID parses 'input' into a UserIdChatIdInstalledAppId
func ParseUserIdChatIdInstalledAppID(input string) (*UserIdChatIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdChatIdInstalledAppIDInsensitively parses 'input' case-insensitively into a UserIdChatIdInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseUserIdChatIdInstalledAppIDInsensitively(input string) (*UserIdChatIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdChatIdInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdChatIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdChatIdInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ChatId, ok = input.Parsed["chatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "chatId", input)
	}

	if id.TeamsAppInstallationId, ok = input.Parsed["teamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", input)
	}

	return nil
}

// ValidateUserIdChatIdInstalledAppID checks that 'input' can be parsed as a User Id Chat Id Installed App ID
func ValidateUserIdChatIdInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdChatIdInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Chat Id Installed App ID
func (id UserIdChatIdInstalledAppId) ID() string {
	fmtString := "/users/%s/chats/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Chat Id Installed App ID
func (id UserIdChatIdInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("chats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatId"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationId"),
	}
}

// String returns a human-readable description of this User Id Chat Id Installed App ID
func (id UserIdChatIdInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("User Id Chat Id Installed App (%s)", strings.Join(components, "\n"))
}
