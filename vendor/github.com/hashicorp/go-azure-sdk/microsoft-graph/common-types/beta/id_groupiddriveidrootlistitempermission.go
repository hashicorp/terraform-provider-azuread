package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootListItemPermissionId{}

// GroupIdDriveIdRootListItemPermissionId is a struct representing the Resource ID for a Group Id Drive Id Root List Item Permission
type GroupIdDriveIdRootListItemPermissionId struct {
	GroupId      string
	DriveId      string
	PermissionId string
}

// NewGroupIdDriveIdRootListItemPermissionID returns a new GroupIdDriveIdRootListItemPermissionId struct
func NewGroupIdDriveIdRootListItemPermissionID(groupId string, driveId string, permissionId string) GroupIdDriveIdRootListItemPermissionId {
	return GroupIdDriveIdRootListItemPermissionId{
		GroupId:      groupId,
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdDriveIdRootListItemPermissionID parses 'input' into a GroupIdDriveIdRootListItemPermissionId
func ParseGroupIdDriveIdRootListItemPermissionID(input string) (*GroupIdDriveIdRootListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootListItemPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootListItemPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootListItemPermissionIDInsensitively(input string) (*GroupIdDriveIdRootListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootListItemPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootListItemPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdRootListItemPermissionID checks that 'input' can be parsed as a Group Id Drive Id Root List Item Permission ID
func ValidateGroupIdDriveIdRootListItemPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootListItemPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root List Item Permission ID
func (id GroupIdDriveIdRootListItemPermissionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/listItem/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root List Item Permission ID
func (id GroupIdDriveIdRootListItemPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root List Item Permission ID
func (id GroupIdDriveIdRootListItemPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root List Item Permission (%s)", strings.Join(components, "\n"))
}
