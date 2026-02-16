package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdPermissionId{}

// MeDriveIdItemIdPermissionId is a struct representing the Resource ID for a Me Drive Id Item Id Permission
type MeDriveIdItemIdPermissionId struct {
	DriveId      string
	DriveItemId  string
	PermissionId string
}

// NewMeDriveIdItemIdPermissionID returns a new MeDriveIdItemIdPermissionId struct
func NewMeDriveIdItemIdPermissionID(driveId string, driveItemId string, permissionId string) MeDriveIdItemIdPermissionId {
	return MeDriveIdItemIdPermissionId{
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		PermissionId: permissionId,
	}
}

// ParseMeDriveIdItemIdPermissionID parses 'input' into a MeDriveIdItemIdPermissionId
func ParseMeDriveIdItemIdPermissionID(input string) (*MeDriveIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdPermissionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdPermissionIDInsensitively(input string) (*MeDriveIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeDriveIdItemIdPermissionID checks that 'input' can be parsed as a Me Drive Id Item Id Permission ID
func ValidateMeDriveIdItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Permission ID
func (id MeDriveIdItemIdPermissionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Permission ID
func (id MeDriveIdItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Permission ID
func (id MeDriveIdItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Permission (%s)", strings.Join(components, "\n"))
}
