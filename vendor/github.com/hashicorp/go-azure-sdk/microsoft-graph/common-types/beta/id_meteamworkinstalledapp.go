package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTeamworkInstalledAppId{}

// MeTeamworkInstalledAppId is a struct representing the Resource ID for a Me Teamwork Installed App
type MeTeamworkInstalledAppId struct {
	UserScopeTeamsAppInstallationId string
}

// NewMeTeamworkInstalledAppID returns a new MeTeamworkInstalledAppId struct
func NewMeTeamworkInstalledAppID(userScopeTeamsAppInstallationId string) MeTeamworkInstalledAppId {
	return MeTeamworkInstalledAppId{
		UserScopeTeamsAppInstallationId: userScopeTeamsAppInstallationId,
	}
}

// ParseMeTeamworkInstalledAppID parses 'input' into a MeTeamworkInstalledAppId
func ParseMeTeamworkInstalledAppID(input string) (*MeTeamworkInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTeamworkInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTeamworkInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeTeamworkInstalledAppIDInsensitively parses 'input' case-insensitively into a MeTeamworkInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseMeTeamworkInstalledAppIDInsensitively(input string) (*MeTeamworkInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeTeamworkInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeTeamworkInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeTeamworkInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserScopeTeamsAppInstallationId, ok = input.Parsed["userScopeTeamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userScopeTeamsAppInstallationId", input)
	}

	return nil
}

// ValidateMeTeamworkInstalledAppID checks that 'input' can be parsed as a Me Teamwork Installed App ID
func ValidateMeTeamworkInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTeamworkInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Teamwork Installed App ID
func (id MeTeamworkInstalledAppId) ID() string {
	fmtString := "/me/teamwork/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserScopeTeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Teamwork Installed App ID
func (id MeTeamworkInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("teamwork", "teamwork", "teamwork"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("userScopeTeamsAppInstallationId", "userScopeTeamsAppInstallationId"),
	}
}

// String returns a human-readable description of this Me Teamwork Installed App ID
func (id MeTeamworkInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User Scope Teams App Installation: %q", id.UserScopeTeamsAppInstallationId),
	}
	return fmt.Sprintf("Me Teamwork Installed App (%s)", strings.Join(components, "\n"))
}
