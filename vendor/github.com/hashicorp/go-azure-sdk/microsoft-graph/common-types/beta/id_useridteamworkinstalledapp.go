package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTeamworkInstalledAppId{}

// UserIdTeamworkInstalledAppId is a struct representing the Resource ID for a User Id Teamwork Installed App
type UserIdTeamworkInstalledAppId struct {
	UserId                          string
	UserScopeTeamsAppInstallationId string
}

// NewUserIdTeamworkInstalledAppID returns a new UserIdTeamworkInstalledAppId struct
func NewUserIdTeamworkInstalledAppID(userId string, userScopeTeamsAppInstallationId string) UserIdTeamworkInstalledAppId {
	return UserIdTeamworkInstalledAppId{
		UserId:                          userId,
		UserScopeTeamsAppInstallationId: userScopeTeamsAppInstallationId,
	}
}

// ParseUserIdTeamworkInstalledAppID parses 'input' into a UserIdTeamworkInstalledAppId
func ParseUserIdTeamworkInstalledAppID(input string) (*UserIdTeamworkInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTeamworkInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTeamworkInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdTeamworkInstalledAppIDInsensitively parses 'input' case-insensitively into a UserIdTeamworkInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseUserIdTeamworkInstalledAppIDInsensitively(input string) (*UserIdTeamworkInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdTeamworkInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdTeamworkInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdTeamworkInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.UserScopeTeamsAppInstallationId, ok = input.Parsed["userScopeTeamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userScopeTeamsAppInstallationId", input)
	}

	return nil
}

// ValidateUserIdTeamworkInstalledAppID checks that 'input' can be parsed as a User Id Teamwork Installed App ID
func ValidateUserIdTeamworkInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdTeamworkInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Teamwork Installed App ID
func (id UserIdTeamworkInstalledAppId) ID() string {
	fmtString := "/users/%s/teamwork/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserScopeTeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Teamwork Installed App ID
func (id UserIdTeamworkInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("teamwork", "teamwork", "teamwork"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("userScopeTeamsAppInstallationId", "userScopeTeamsAppInstallationId"),
	}
}

// String returns a human-readable description of this User Id Teamwork Installed App ID
func (id UserIdTeamworkInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Scope Teams App Installation: %q", id.UserScopeTeamsAppInstallationId),
	}
	return fmt.Sprintf("User Id Teamwork Installed App (%s)", strings.Join(components, "\n"))
}
