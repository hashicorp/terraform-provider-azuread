package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelSharedWithTeamId{}

// GroupIdTeamPrimaryChannelSharedWithTeamId is a struct representing the Resource ID for a Group Id Team Primary Channel Shared With Team
type GroupIdTeamPrimaryChannelSharedWithTeamId struct {
	GroupId                     string
	SharedWithChannelTeamInfoId string
}

// NewGroupIdTeamPrimaryChannelSharedWithTeamID returns a new GroupIdTeamPrimaryChannelSharedWithTeamId struct
func NewGroupIdTeamPrimaryChannelSharedWithTeamID(groupId string, sharedWithChannelTeamInfoId string) GroupIdTeamPrimaryChannelSharedWithTeamId {
	return GroupIdTeamPrimaryChannelSharedWithTeamId{
		GroupId:                     groupId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseGroupIdTeamPrimaryChannelSharedWithTeamID parses 'input' into a GroupIdTeamPrimaryChannelSharedWithTeamId
func ParseGroupIdTeamPrimaryChannelSharedWithTeamID(input string) (*GroupIdTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPrimaryChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPrimaryChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPrimaryChannelSharedWithTeamIDInsensitively(input string) (*GroupIdTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPrimaryChannelSharedWithTeamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPrimaryChannelSharedWithTeamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	return nil
}

// ValidateGroupIdTeamPrimaryChannelSharedWithTeamID checks that 'input' can be parsed as a Group Id Team Primary Channel Shared With Team ID
func ValidateGroupIdTeamPrimaryChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPrimaryChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Primary Channel Shared With Team ID
func (id GroupIdTeamPrimaryChannelSharedWithTeamId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Primary Channel Shared With Team ID
func (id GroupIdTeamPrimaryChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("primaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
	}
}

// String returns a human-readable description of this Group Id Team Primary Channel Shared With Team ID
func (id GroupIdTeamPrimaryChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Group Id Team Primary Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
