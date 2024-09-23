package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdListItemPermissionId{}

// MeDriveIdItemIdListItemPermissionId is a struct representing the Resource ID for a Me Drive Id Item Id List Item Permission
type MeDriveIdItemIdListItemPermissionId struct {
	DriveId      string
	DriveItemId  string
	PermissionId string
}

// NewMeDriveIdItemIdListItemPermissionID returns a new MeDriveIdItemIdListItemPermissionId struct
func NewMeDriveIdItemIdListItemPermissionID(driveId string, driveItemId string, permissionId string) MeDriveIdItemIdListItemPermissionId {
	return MeDriveIdItemIdListItemPermissionId{
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		PermissionId: permissionId,
	}
}

// ParseMeDriveIdItemIdListItemPermissionID parses 'input' into a MeDriveIdItemIdListItemPermissionId
func ParseMeDriveIdItemIdListItemPermissionID(input string) (*MeDriveIdItemIdListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdListItemPermissionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdListItemPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdListItemPermissionIDInsensitively(input string) (*MeDriveIdItemIdListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdListItemPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdListItemPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeDriveIdItemIdListItemPermissionID checks that 'input' can be parsed as a Me Drive Id Item Id List Item Permission ID
func ValidateMeDriveIdItemIdListItemPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdListItemPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id List Item Permission ID
func (id MeDriveIdItemIdListItemPermissionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/listItem/permissions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id List Item Permission ID
func (id MeDriveIdItemIdListItemPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id List Item Permission ID
func (id MeDriveIdItemIdListItemPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id List Item Permission (%s)", strings.Join(components, "\n"))
}
