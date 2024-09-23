package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamTagIdMemberId{}

// GroupIdTeamTagIdMemberId is a struct representing the Resource ID for a Group Id Team Tag Id Member
type GroupIdTeamTagIdMemberId struct {
	GroupId             string
	TeamworkTagId       string
	TeamworkTagMemberId string
}

// NewGroupIdTeamTagIdMemberID returns a new GroupIdTeamTagIdMemberId struct
func NewGroupIdTeamTagIdMemberID(groupId string, teamworkTagId string, teamworkTagMemberId string) GroupIdTeamTagIdMemberId {
	return GroupIdTeamTagIdMemberId{
		GroupId:             groupId,
		TeamworkTagId:       teamworkTagId,
		TeamworkTagMemberId: teamworkTagMemberId,
	}
}

// ParseGroupIdTeamTagIdMemberID parses 'input' into a GroupIdTeamTagIdMemberId
func ParseGroupIdTeamTagIdMemberID(input string) (*GroupIdTeamTagIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamTagIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamTagIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamTagIdMemberIDInsensitively parses 'input' case-insensitively into a GroupIdTeamTagIdMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamTagIdMemberIDInsensitively(input string) (*GroupIdTeamTagIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamTagIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamTagIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamTagIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TeamworkTagId, ok = input.Parsed["teamworkTagId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagId", input)
	}

	if id.TeamworkTagMemberId, ok = input.Parsed["teamworkTagMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamworkTagMemberId", input)
	}

	return nil
}

// ValidateGroupIdTeamTagIdMemberID checks that 'input' can be parsed as a Group Id Team Tag Id Member ID
func ValidateGroupIdTeamTagIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamTagIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Tag Id Member ID
func (id GroupIdTeamTagIdMemberId) ID() string {
	fmtString := "/groups/%s/team/tags/%s/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamworkTagId, id.TeamworkTagMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Tag Id Member ID
func (id GroupIdTeamTagIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("tags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("teamworkTagId", "teamworkTagId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("teamworkTagMemberId", "teamworkTagMemberId"),
	}
}

// String returns a human-readable description of this Group Id Team Tag Id Member ID
func (id GroupIdTeamTagIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teamwork Tag: %q", id.TeamworkTagId),
		fmt.Sprintf("Teamwork Tag Member: %q", id.TeamworkTagMemberId),
	}
	return fmt.Sprintf("Group Id Team Tag Id Member (%s)", strings.Join(components, "\n"))
}
