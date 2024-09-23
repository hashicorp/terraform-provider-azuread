package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPermissionGrantId{}

// MeJoinedTeamIdPermissionGrantId is a struct representing the Resource ID for a Me Joined Team Id Permission Grant
type MeJoinedTeamIdPermissionGrantId struct {
	TeamId                            string
	ResourceSpecificPermissionGrantId string
}

// NewMeJoinedTeamIdPermissionGrantID returns a new MeJoinedTeamIdPermissionGrantId struct
func NewMeJoinedTeamIdPermissionGrantID(teamId string, resourceSpecificPermissionGrantId string) MeJoinedTeamIdPermissionGrantId {
	return MeJoinedTeamIdPermissionGrantId{
		TeamId:                            teamId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseMeJoinedTeamIdPermissionGrantID parses 'input' into a MeJoinedTeamIdPermissionGrantId
func ParseMeJoinedTeamIdPermissionGrantID(input string) (*MeJoinedTeamIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdPermissionGrantIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdPermissionGrantIDInsensitively(input string) (*MeJoinedTeamIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdPermissionGrantID checks that 'input' can be parsed as a Me Joined Team Id Permission Grant ID
func ValidateMeJoinedTeamIdPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Permission Grant ID
func (id MeJoinedTeamIdPermissionGrantId) ID() string {
	fmtString := "/me/joinedTeams/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Permission Grant ID
func (id MeJoinedTeamIdPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Permission Grant ID
func (id MeJoinedTeamIdPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Me Joined Team Id Permission Grant (%s)", strings.Join(components, "\n"))
}
