package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootChildId{}

// MeDriveIdRootChildId is a struct representing the Resource ID for a Me Drive Id Root Child
type MeDriveIdRootChildId struct {
	DriveId     string
	DriveItemId string
}

// NewMeDriveIdRootChildID returns a new MeDriveIdRootChildId struct
func NewMeDriveIdRootChildID(driveId string, driveItemId string) MeDriveIdRootChildId {
	return MeDriveIdRootChildId{
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseMeDriveIdRootChildID parses 'input' into a MeDriveIdRootChildId
func ParseMeDriveIdRootChildID(input string) (*MeDriveIdRootChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootChildIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootChildId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootChildIDInsensitively(input string) (*MeDriveIdRootChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootChildId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	return nil
}

// ValidateMeDriveIdRootChildID checks that 'input' can be parsed as a Me Drive Id Root Child ID
func ValidateMeDriveIdRootChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Child ID
func (id MeDriveIdRootChildId) ID() string {
	fmtString := "/me/drives/%s/root/children/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Child ID
func (id MeDriveIdRootChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Child ID
func (id MeDriveIdRootChildId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Me Drive Id Root Child (%s)", strings.Join(components, "\n"))
}
