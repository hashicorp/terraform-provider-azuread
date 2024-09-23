package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemId{}

// MeDriveIdItemId is a struct representing the Resource ID for a Me Drive Id Item
type MeDriveIdItemId struct {
	DriveId     string
	DriveItemId string
}

// NewMeDriveIdItemID returns a new MeDriveIdItemId struct
func NewMeDriveIdItemID(driveId string, driveItemId string) MeDriveIdItemId {
	return MeDriveIdItemId{
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseMeDriveIdItemID parses 'input' into a MeDriveIdItemId
func ParseMeDriveIdItemID(input string) (*MeDriveIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIDInsensitively(input string) (*MeDriveIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	return nil
}

// ValidateMeDriveIdItemID checks that 'input' can be parsed as a Me Drive Id Item ID
func ValidateMeDriveIdItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item ID
func (id MeDriveIdItemId) ID() string {
	fmtString := "/me/drives/%s/items/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item ID
func (id MeDriveIdItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item ID
func (id MeDriveIdItemId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("Me Drive Id Item (%s)", strings.Join(components, "\n"))
}
