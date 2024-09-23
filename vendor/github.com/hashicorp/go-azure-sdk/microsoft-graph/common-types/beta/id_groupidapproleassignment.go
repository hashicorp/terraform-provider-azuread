package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdAppRoleAssignmentId{}

// GroupIdAppRoleAssignmentId is a struct representing the Resource ID for a Group Id App Role Assignment
type GroupIdAppRoleAssignmentId struct {
	GroupId             string
	AppRoleAssignmentId string
}

// NewGroupIdAppRoleAssignmentID returns a new GroupIdAppRoleAssignmentId struct
func NewGroupIdAppRoleAssignmentID(groupId string, appRoleAssignmentId string) GroupIdAppRoleAssignmentId {
	return GroupIdAppRoleAssignmentId{
		GroupId:             groupId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseGroupIdAppRoleAssignmentID parses 'input' into a GroupIdAppRoleAssignmentId
func ParseGroupIdAppRoleAssignmentID(input string) (*GroupIdAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a GroupIdAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseGroupIdAppRoleAssignmentIDInsensitively(input string) (*GroupIdAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdAppRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdAppRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.AppRoleAssignmentId, ok = input.Parsed["appRoleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", input)
	}

	return nil
}

// ValidateGroupIdAppRoleAssignmentID checks that 'input' can be parsed as a Group Id App Role Assignment ID
func ValidateGroupIdAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id App Role Assignment ID
func (id GroupIdAppRoleAssignmentId) ID() string {
	fmtString := "/groups/%s/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id App Role Assignment ID
func (id GroupIdAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("appRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentId"),
	}
}

// String returns a human-readable description of this Group Id App Role Assignment ID
func (id GroupIdAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Group Id App Role Assignment (%s)", strings.Join(components, "\n"))
}
