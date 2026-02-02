package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelTabId{}

// MeJoinedTeamIdPrimaryChannelTabId is a struct representing the Resource ID for a Me Joined Team Id Primary Channel Tab
type MeJoinedTeamIdPrimaryChannelTabId struct {
	TeamId     string
	TeamsTabId string
}

// NewMeJoinedTeamIdPrimaryChannelTabID returns a new MeJoinedTeamIdPrimaryChannelTabId struct
func NewMeJoinedTeamIdPrimaryChannelTabID(teamId string, teamsTabId string) MeJoinedTeamIdPrimaryChannelTabId {
	return MeJoinedTeamIdPrimaryChannelTabId{
		TeamId:     teamId,
		TeamsTabId: teamsTabId,
	}
}

// ParseMeJoinedTeamIdPrimaryChannelTabID parses 'input' into a MeJoinedTeamIdPrimaryChannelTabId
func ParseMeJoinedTeamIdPrimaryChannelTabID(input string) (*MeJoinedTeamIdPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPrimaryChannelTabIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPrimaryChannelTabId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPrimaryChannelTabIDInsensitively(input string) (*MeJoinedTeamIdPrimaryChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPrimaryChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPrimaryChannelTabId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPrimaryChannelTabId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.TeamsTabId, ok = input.Parsed["teamsTabId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPrimaryChannelTabID checks that 'input' can be parsed as a Me Joined Team Id Primary Channel Tab ID
func ValidateMeJoinedTeamIdPrimaryChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPrimaryChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Primary Channel Tab ID
func (id MeJoinedTeamIdPrimaryChannelTabId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/tabs/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Primary Channel Tab ID
func (id MeJoinedTeamIdPrimaryChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("tabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Primary Channel Tab ID
func (id MeJoinedTeamIdPrimaryChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Me Joined Team Id Primary Channel Tab (%s)", strings.Join(components, "\n"))
}
