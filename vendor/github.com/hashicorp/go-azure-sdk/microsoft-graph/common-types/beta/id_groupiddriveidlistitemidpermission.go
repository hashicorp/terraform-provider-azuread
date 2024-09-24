package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdListItemIdPermissionId{}

// GroupIdDriveIdListItemIdPermissionId is a struct representing the Resource ID for a Group Id Drive Id List Item Id Permission
type GroupIdDriveIdListItemIdPermissionId struct {
	GroupId      string
	DriveId      string
	ListItemId   string
	PermissionId string
}

// NewGroupIdDriveIdListItemIdPermissionID returns a new GroupIdDriveIdListItemIdPermissionId struct
func NewGroupIdDriveIdListItemIdPermissionID(groupId string, driveId string, listItemId string, permissionId string) GroupIdDriveIdListItemIdPermissionId {
	return GroupIdDriveIdListItemIdPermissionId{
		GroupId:      groupId,
		DriveId:      driveId,
		ListItemId:   listItemId,
		PermissionId: permissionId,
	}
}

// ParseGroupIdDriveIdListItemIdPermissionID parses 'input' into a GroupIdDriveIdListItemIdPermissionId
func ParseGroupIdDriveIdListItemIdPermissionID(input string) (*GroupIdDriveIdListItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdListItemIdPermissionIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdListItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdListItemIdPermissionIDInsensitively(input string) (*GroupIdDriveIdListItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdListItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdListItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdListItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdListItemIdPermissionID checks that 'input' can be parsed as a Group Id Drive Id List Item Id Permission ID
func ValidateGroupIdDriveIdListItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdListItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id List Item Id Permission ID
func (id GroupIdDriveIdListItemIdPermissionId) ID() string {
	fmtString := "/groups/%s/drives/%s/list/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ListItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id List Item Id Permission ID
func (id GroupIdDriveIdListItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id List Item Id Permission ID
func (id GroupIdDriveIdListItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Group Id Drive Id List Item Id Permission (%s)", strings.Join(components, "\n"))
}
