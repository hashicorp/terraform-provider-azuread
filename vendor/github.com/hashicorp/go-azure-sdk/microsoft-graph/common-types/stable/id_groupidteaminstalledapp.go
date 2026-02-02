package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamInstalledAppId{}

// GroupIdTeamInstalledAppId is a struct representing the Resource ID for a Group Id Team Installed App
type GroupIdTeamInstalledAppId struct {
	GroupId                string
	TeamsAppInstallationId string
}

// NewGroupIdTeamInstalledAppID returns a new GroupIdTeamInstalledAppId struct
func NewGroupIdTeamInstalledAppID(groupId string, teamsAppInstallationId string) GroupIdTeamInstalledAppId {
	return GroupIdTeamInstalledAppId{
		GroupId:                groupId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseGroupIdTeamInstalledAppID parses 'input' into a GroupIdTeamInstalledAppId
func ParseGroupIdTeamInstalledAppID(input string) (*GroupIdTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamInstalledAppIDInsensitively parses 'input' case-insensitively into a GroupIdTeamInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamInstalledAppIDInsensitively(input string) (*GroupIdTeamInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamInstalledAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamInstalledAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TeamsAppInstallationId, ok = input.Parsed["teamsAppInstallationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", input)
	}

	return nil
}

// ValidateGroupIdTeamInstalledAppID checks that 'input' can be parsed as a Group Id Team Installed App ID
func ValidateGroupIdTeamInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Installed App ID
func (id GroupIdTeamInstalledAppId) ID() string {
	fmtString := "/groups/%s/team/installedApps/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Installed App ID
func (id GroupIdTeamInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("installedApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationId"),
	}
}

// String returns a human-readable description of this Group Id Team Installed App ID
func (id GroupIdTeamInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("Group Id Team Installed App (%s)", strings.Join(components, "\n"))
}
