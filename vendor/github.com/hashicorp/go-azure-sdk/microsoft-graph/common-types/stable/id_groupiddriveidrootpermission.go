package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootPermissionId{}

// GroupIdDriveIdRootPermissionId is a struct representing the Resource ID for a Group Id Drive Id Root Permission
type GroupIdDriveIdRootPermissionId struct {
	GroupId      string
	DriveId      string
	PermissionId string
}

// NewGroupIdDriveIdRootPermissionID returns a new GroupIdDriveIdRootPermissionId struct
func NewGroupIdDriveIdRootPermissionID(groupId string, driveId string, permissionId string) GroupIdDriveIdRootPermissionId {
	return GroupIdDriveIdRootPermissionId{
		GroupId:      groupId,
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdDriveIdRootPermissionID parses 'input' into a GroupIdDriveIdRootPermissionId
func ParseGroupIdDriveIdRootPermissionID(input string) (*GroupIdDriveIdRootPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootPermissionIDInsensitively(input string) (*GroupIdDriveIdRootPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdRootPermissionID checks that 'input' can be parsed as a Group Id Drive Id Root Permission ID
func ValidateGroupIdDriveIdRootPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Permission ID
func (id GroupIdDriveIdRootPermissionId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Permission ID
func (id GroupIdDriveIdRootPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Permission ID
func (id GroupIdDriveIdRootPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Permission (%s)", strings.Join(components, "\n"))
}
