package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootListItemPermissionId{}

// MeDriveIdRootListItemPermissionId is a struct representing the Resource ID for a Me Drive Id Root List Item Permission
type MeDriveIdRootListItemPermissionId struct {
	DriveId      string
	PermissionId string
}

// NewMeDriveIdRootListItemPermissionID returns a new MeDriveIdRootListItemPermissionId struct
func NewMeDriveIdRootListItemPermissionID(driveId string, permissionId string) MeDriveIdRootListItemPermissionId {
	return MeDriveIdRootListItemPermissionId{
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseMeDriveIdRootListItemPermissionID parses 'input' into a MeDriveIdRootListItemPermissionId
func ParseMeDriveIdRootListItemPermissionID(input string) (*MeDriveIdRootListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootListItemPermissionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootListItemPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootListItemPermissionIDInsensitively(input string) (*MeDriveIdRootListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootListItemPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootListItemPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootListItemPermissionID checks that 'input' can be parsed as a Me Drive Id Root List Item Permission ID
func ValidateMeDriveIdRootListItemPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootListItemPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root List Item Permission ID
func (id MeDriveIdRootListItemPermissionId) ID() string {
	fmtString := "/me/drives/%s/root/listItem/permissions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root List Item Permission ID
func (id MeDriveIdRootListItemPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root List Item Permission ID
func (id MeDriveIdRootListItemPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Me Drive Id Root List Item Permission (%s)", strings.Join(components, "\n"))
}
