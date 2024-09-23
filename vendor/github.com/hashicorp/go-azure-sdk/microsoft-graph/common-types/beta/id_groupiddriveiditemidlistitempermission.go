package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdListItemPermissionId{}

// GroupIdDriveIdItemIdListItemPermissionId is a struct representing the Resource ID for a Group Id Drive Id Item Id List Item Permission
type GroupIdDriveIdItemIdListItemPermissionId struct {
	GroupId      string
	DriveId      string
	DriveItemId  string
	PermissionId string
}

// NewGroupIdDriveIdItemIdListItemPermissionID returns a new GroupIdDriveIdItemIdListItemPermissionId struct
func NewGroupIdDriveIdItemIdListItemPermissionID(groupId string, driveId string, driveItemId string, permissionId string) GroupIdDriveIdItemIdListItemPermissionId {
	return GroupIdDriveIdItemIdListItemPermissionId{
		GroupId:      groupId,
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdDriveIdItemIdListItemPermissionID parses 'input' into a GroupIdDriveIdItemIdListItemPermissionId
func ParseGroupIdDriveIdItemIdListItemPermissionID(input string) (*GroupIdDriveIdItemIdListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdListItemPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdListItemPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdListItemPermissionIDInsensitively(input string) (*GroupIdDriveIdItemIdListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdListItemPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdListItemPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateGroupIdDriveIdItemIdListItemPermissionID checks that 'input' can be parsed as a Group Id Drive Id Item Id List Item Permission ID
func ValidateGroupIdDriveIdItemIdListItemPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdListItemPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id List Item Permission ID
func (id GroupIdDriveIdItemIdListItemPermissionId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/listItem/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id List Item Permission ID
func (id GroupIdDriveIdItemIdListItemPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id List Item Permission ID
func (id GroupIdDriveIdItemIdListItemPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id List Item Permission (%s)", strings.Join(components, "\n"))
}
