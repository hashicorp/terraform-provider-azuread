package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListPermissionId{}

// GroupIdDriveIdListPermissionId is a struct representing the Resource ID for a Group Id Drive Id List Permission
type GroupIdDriveIdListPermissionId struct {
	GroupId      string
	DriveId      string
	PermissionId string
}

// NewGroupIdDriveIdListPermissionID returns a new GroupIdDriveIdListPermissionId struct
func NewGroupIdDriveIdListPermissionID(groupId string, driveId string, permissionId string) GroupIdDriveIdListPermissionId {
	return GroupIdDriveIdListPermissionId{
		GroupId:      groupId,
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdDriveIdListPermissionID parses 'input' into a GroupIdDriveIdListPermissionId
func ParseGroupIdDriveIdListPermissionID(input string) (*GroupIdDriveIdListPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListPermissionIDInsensitively(input string) (*GroupIdDriveIdListPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListPermissionID checks that 'input' can be parsed as a Group Id Drive Id List Permission ID
func ValidateGroupIdDriveIdListPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Permission ID
func (id GroupIdDriveIdListPermissionId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Permission ID
func (id GroupIdDriveIdListPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Permission ID
func (id GroupIdDriveIdListPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Permission (%s)", strings.Join(components, "\n"))
}
