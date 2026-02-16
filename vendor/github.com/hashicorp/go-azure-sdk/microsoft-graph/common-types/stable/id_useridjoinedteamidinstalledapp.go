package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdInstalledAppId{}

// UserIdJoinedTeamIdInstalledAppId is a struct representing the Resource ID for a User Id Joined Team Id Installed App
type UserIdJoinedTeamIdInstalledAppId struct {
	UserId                 string
	TeamId                 string
	TeamsAppInstallationId string
}

// NewUserIdJoinedTeamIdInstalledAppID returns a new UserIdJoinedTeamIdInstalledAppId struct
func NewUserIdJoinedTeamIdInstalledAppID(userId string, teamId string, teamsAppInstallationId string) UserIdJoinedTeamIdInstalledAppId {
	return UserIdJoinedTeamIdInstalledAppId{
		UserId:                 userId,
		TeamId:                 teamId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseUserIdJoinedTeamIdInstalledAppID parses 'input' into a UserIdJoinedTeamIdInstalledAppId
func ParseUserIdJoinedTeamIdInstalledAppID(input string) (*UserIdJoinedTeamIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdJoinedTeamIdInstalledAppIDInsensitively parses 'input' case-insensitively into a UserIdJoinedTeamIdInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseUserIdJoinedTeamIdInstalledAppIDInsensitively(input string) (*UserIdJoinedTeamIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdJoinedTeamIdInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdJoinedTeamIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdJoinedTeamIdInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamsAppInstallationId, ok = input.Parsed["teamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", input)
	}

	return nil
}

// ValidateUserIdJoinedTeamIdInstalledAppID checks that 'input' can be parsed as a User Id Joined Team Id Installed App ID
func ValidateUserIdJoinedTeamIdInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdJoinedTeamIdInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Joined Team Id Installed App ID
func (id UserIdJoinedTeamIdInstalledAppId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Joined Team Id Installed App ID
func (id UserIdJoinedTeamIdInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationId"),
	}
}

// String returns a human-readable description of this User Id Joined Team Id Installed App ID
func (id UserIdJoinedTeamIdInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("User Id Joined Team Id Installed App (%s)", strings.Join(components, "\n"))
}
