package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdPermissionId{}

// GroupIdDriveIdItemIdPermissionId is a struct representing the Resource ID for a Group Id Drive Id Item Id Permission
type GroupIdDriveIdItemIdPermissionId struct {
	GroupId      string
	DriveId      string
	DriveItemId  string
	PermissionId string
}

// NewGroupIdDriveIdItemIdPermissionID returns a new GroupIdDriveIdItemIdPermissionId struct
func NewGroupIdDriveIdItemIdPermissionID(groupId string, driveId string, driveItemId string, permissionId string) GroupIdDriveIdItemIdPermissionId {
	return GroupIdDriveIdItemIdPermissionId{
		GroupId:      groupId,
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdDriveIdItemIdPermissionID parses 'input' into a GroupIdDriveIdItemIdPermissionId
func ParseGroupIdDriveIdItemIdPermissionID(input string) (*GroupIdDriveIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdPermissionIDInsensitively(input string) (*GroupIdDriveIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdPermissionID checks that 'input' can be parsed as a Group Id Drive Id Item Id Permission ID
func ValidateGroupIdDriveIdItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Permission ID
func (id GroupIdDriveIdItemIdPermissionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Permission ID
func (id GroupIdDriveIdItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Permission ID
func (id GroupIdDriveIdItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Permission (%s)", strings.Join(components, "\n"))
}
