package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListItemIdPermissionId{}

// MeDriveIdListItemIdPermissionId is a struct representing the Resource ID for a Me Drive Id List Item Id Permission
type MeDriveIdListItemIdPermissionId struct {
	DriveId      string
	ListItemId   string
	PermissionId string
}

// NewMeDriveIdListItemIdPermissionID returns a new MeDriveIdListItemIdPermissionId struct
func NewMeDriveIdListItemIdPermissionID(driveId string, listItemId string, permissionId string) MeDriveIdListItemIdPermissionId {
	return MeDriveIdListItemIdPermissionId{
		DriveId:      driveId,
		ListItemId:   listItemId,
		PermissionId: permissionId,
	}
}

// ParseMeDriveIdListItemIdPermissionID parses 'input' into a MeDriveIdListItemIdPermissionId
func ParseMeDriveIdListItemIdPermissionID(input string) (*MeDriveIdListItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListItemIdPermissionIDInsensitively parses 'input' case-insensitively into a MeDriveIdListItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListItemIdPermissionIDInsensitively(input string) (*MeDriveIdListItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeDriveIdListItemIdPermissionID checks that 'input' can be parsed as a Me Drive Id List Item Id Permission ID
func ValidateMeDriveIdListItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Item Id Permission ID
func (id MeDriveIdListItemIdPermissionId) ID() string {
	fmtString := "/me/drives/%s/list/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ListItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Item Id Permission ID
func (id MeDriveIdListItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Item Id Permission ID
func (id MeDriveIdListItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Me Drive Id List Item Id Permission (%s)", strings.Join(components, "\n"))
}
