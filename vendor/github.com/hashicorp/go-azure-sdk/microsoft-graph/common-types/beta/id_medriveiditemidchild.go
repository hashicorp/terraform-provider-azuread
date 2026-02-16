package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdChildId{}

// MeDriveIdItemIdChildId is a struct representing the Resource ID for a Me Drive Id Item Id Child
type MeDriveIdItemIdChildId struct {
	DriveId      string
	DriveItemId  string
	DriveItemId1 string
}

// NewMeDriveIdItemIdChildID returns a new MeDriveIdItemIdChildId struct
func NewMeDriveIdItemIdChildID(driveId string, driveItemId string, driveItemId1 string) MeDriveIdItemIdChildId {
	return MeDriveIdItemIdChildId{
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		DriveItemId1: driveItemId1,
	}
}

// ParseMeDriveIdItemIdChildID parses 'input' into a MeDriveIdItemIdChildId
func ParseMeDriveIdItemIdChildID(input string) (*MeDriveIdItemIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdChildIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdChildId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdChildIDInsensitively(input string) (*MeDriveIdItemIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdChildId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.DriveItemId1, ok = input.Parsed["driveItemId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId1", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdChildID checks that 'input' can be parsed as a Me Drive Id Item Id Child ID
func ValidateMeDriveIdItemIdChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Child ID
func (id MeDriveIdItemIdChildId) ID() string {
	fmtString := "/me/drives/%s/items/%s/children/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.DriveItemId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Child ID
func (id MeDriveIdItemIdChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("driveItemId1", "driveItemId1"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Child ID
func (id MeDriveIdItemIdChildId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Drive Item Id 1: %q", id.DriveItemId1),
	}
	return fmt.Sprintf("Me Drive Id Item Id Child (%s)", strings.Join(components, "\n"))
}
