package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamOperationId{}

// GroupIdTeamOperationId is a struct representing the Resource ID for a Group Id Team Operation
type GroupIdTeamOperationId struct {
	GroupId               string
	TeamsAsyncOperationId string
}

// NewGroupIdTeamOperationID returns a new GroupIdTeamOperationId struct
func NewGroupIdTeamOperationID(groupId string, teamsAsyncOperationId string) GroupIdTeamOperationId {
	return GroupIdTeamOperationId{
		GroupId:               groupId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseGroupIdTeamOperationID parses 'input' into a GroupIdTeamOperationId
func ParseGroupIdTeamOperationID(input string) (*GroupIdTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamOperationIDInsensitively parses 'input' case-insensitively into a GroupIdTeamOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamOperationIDInsensitively(input string) (*GroupIdTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.TeamsAsyncOperationId, ok = input.Parsed["teamsAsyncOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", input)
	}

	return nil
}

// ValidateGroupIdTeamOperationID checks that 'input' can be parsed as a Group Id Team Operation ID
func ValidateGroupIdTeamOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Operation ID
func (id GroupIdTeamOperationId) ID() string {
	fmtString := "/groups/%s/team/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Operation ID
func (id GroupIdTeamOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationId"),
	}
}

// String returns a human-readable description of this Group Id Team Operation ID
func (id GroupIdTeamOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("Group Id Team Operation (%s)", strings.Join(components, "\n"))
}
