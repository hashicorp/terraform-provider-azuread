package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamOwnerId{}

// GroupIdTeamOwnerId is a struct representing the Resource ID for a Group Id Team Owner
type GroupIdTeamOwnerId struct {
	GroupId string
	UserId  string
}

// NewGroupIdTeamOwnerID returns a new GroupIdTeamOwnerId struct
func NewGroupIdTeamOwnerID(groupId string, userId string) GroupIdTeamOwnerId {
	return GroupIdTeamOwnerId{
		GroupId: groupId,
		UserId:  userId,
	}
}

// ParseGroupIdTeamOwnerID parses 'input' into a GroupIdTeamOwnerId
func ParseGroupIdTeamOwnerID(input string) (*GroupIdTeamOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamOwnerIDInsensitively parses 'input' case-insensitively into a GroupIdTeamOwnerId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamOwnerIDInsensitively(input string) (*GroupIdTeamOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamOwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	return nil
}

// ValidateGroupIdTeamOwnerID checks that 'input' can be parsed as a Group Id Team Owner ID
func ValidateGroupIdTeamOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Owner ID
func (id GroupIdTeamOwnerId) ID() string {
	fmtString := "/groups/%s/team/owners/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.UserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Owner ID
func (id GroupIdTeamOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("owners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
	}
}

// String returns a human-readable description of this Group Id Team Owner ID
func (id GroupIdTeamOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("User: %q", id.UserId),
	}
	return fmt.Sprintf("Group Id Team Owner (%s)", strings.Join(components, "\n"))
}
