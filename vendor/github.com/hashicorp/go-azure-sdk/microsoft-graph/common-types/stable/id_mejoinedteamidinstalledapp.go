package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdInstalledAppId{}

// MeJoinedTeamIdInstalledAppId is a struct representing the Resource ID for a Me Joined Team Id Installed App
type MeJoinedTeamIdInstalledAppId struct {
	TeamId                 string
	TeamsAppInstallationId string
}

// NewMeJoinedTeamIdInstalledAppID returns a new MeJoinedTeamIdInstalledAppId struct
func NewMeJoinedTeamIdInstalledAppID(teamId string, teamsAppInstallationId string) MeJoinedTeamIdInstalledAppId {
	return MeJoinedTeamIdInstalledAppId{
		TeamId:                 teamId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseMeJoinedTeamIdInstalledAppID parses 'input' into a MeJoinedTeamIdInstalledAppId
func ParseMeJoinedTeamIdInstalledAppID(input string) (*MeJoinedTeamIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdInstalledAppIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdInstalledAppIDInsensitively(input string) (*MeJoinedTeamIdInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamsAppInstallationId, ok = input.Parsed["teamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdInstalledAppID checks that 'input' can be parsed as a Me Joined Team Id Installed App ID
func ValidateMeJoinedTeamIdInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Installed App ID
func (id MeJoinedTeamIdInstalledAppId) ID() string {
	fmtString := "/me/joinedTeams/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Installed App ID
func (id MeJoinedTeamIdInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Installed App ID
func (id MeJoinedTeamIdInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("Me Joined Team Id Installed App (%s)", strings.Join(components, "\n"))
}
