package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPermissionGrantId{}

// GroupIdTeamPermissionGrantId is a struct representing the Resource ID for a Group Id Team Permission Grant
type GroupIdTeamPermissionGrantId struct {
	GroupId                           string
	ResourceSpecificPermissionGrantId string
}

// NewGroupIdTeamPermissionGrantID returns a new GroupIdTeamPermissionGrantId struct
func NewGroupIdTeamPermissionGrantID(groupId string, resourceSpecificPermissionGrantId string) GroupIdTeamPermissionGrantId {
	return GroupIdTeamPermissionGrantId{
		GroupId:                           groupId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseGroupIdTeamPermissionGrantID parses 'input' into a GroupIdTeamPermissionGrantId
func ParseGroupIdTeamPermissionGrantID(input string) (*GroupIdTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdTeamPermissionGrantIDInsensitively parses 'input' case-insensitively into a GroupIdTeamPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseGroupIdTeamPermissionGrantIDInsensitively(input string) (*GroupIdTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdTeamPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdTeamPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateGroupIdTeamPermissionGrantID checks that 'input' can be parsed as a Group Id Team Permission Grant ID
func ValidateGroupIdTeamPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdTeamPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Team Permission Grant ID
func (id GroupIdTeamPermissionGrantId) ID() string {
	fmtString := "/groups/%s/team/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Team Permission Grant ID
func (id GroupIdTeamPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("team", "team", "team"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this Group Id Team Permission Grant ID
func (id GroupIdTeamPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Group Id Team Permission Grant (%s)", strings.Join(components, "\n"))
}
