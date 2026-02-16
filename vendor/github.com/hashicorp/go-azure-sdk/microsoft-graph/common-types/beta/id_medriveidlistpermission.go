package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdListPermissionId{}

// MeDriveIdListPermissionId is a struct representing the Resource ID for a Me Drive Id List Permission
type MeDriveIdListPermissionId struct {
	DriveId      string
	PermissionId string
}

// NewMeDriveIdListPermissionID returns a new MeDriveIdListPermissionId struct
func NewMeDriveIdListPermissionID(driveId string, permissionId string) MeDriveIdListPermissionId {
	return MeDriveIdListPermissionId{
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseMeDriveIdListPermissionID parses 'input' into a MeDriveIdListPermissionId
func ParseMeDriveIdListPermissionID(input string) (*MeDriveIdListPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdListPermissionIDInsensitively parses 'input' case-insensitively into a MeDriveIdListPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdListPermissionIDInsensitively(input string) (*MeDriveIdListPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdListPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdListPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdListPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateMeDriveIdListPermissionID checks that 'input' can be parsed as a Me Drive Id List Permission ID
func ValidateMeDriveIdListPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdListPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id List Permission ID
func (id MeDriveIdListPermissionId) ID() string {
	fmtString := "/me/drives/%s/list/permissions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id List Permission ID
func (id MeDriveIdListPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Me Drive Id List Permission ID
func (id MeDriveIdListPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Me Drive Id List Permission (%s)", strings.Join(components, "\n"))
}
