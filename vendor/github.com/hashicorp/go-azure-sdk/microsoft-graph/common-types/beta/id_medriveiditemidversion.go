package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdVersionId{}

// MeDriveIdItemIdVersionId is a struct representing the Resource ID for a Me Drive Id Item Id Version
type MeDriveIdItemIdVersionId struct {
	DriveId            string
	DriveItemId        string
	DriveItemVersionId string
}

// NewMeDriveIdItemIdVersionID returns a new MeDriveIdItemIdVersionId struct
func NewMeDriveIdItemIdVersionID(driveId string, driveItemId string, driveItemVersionId string) MeDriveIdItemIdVersionId {
	return MeDriveIdItemIdVersionId{
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		DriveItemVersionId: driveItemVersionId,
	}
}

// ParseMeDriveIdItemIdVersionID parses 'input' into a MeDriveIdItemIdVersionId
func ParseMeDriveIdItemIdVersionID(input string) (*MeDriveIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdVersionIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdVersionIDInsensitively(input string) (*MeDriveIdItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.DriveItemVersionId, ok = input.Parsed["driveItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemVersionId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdVersionID checks that 'input' can be parsed as a Me Drive Id Item Id Version ID
func ValidateMeDriveIdItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Version ID
func (id MeDriveIdItemIdVersionId) ID() string {
	fmtString := "/me/drives/%s/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.DriveItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Version ID
func (id MeDriveIdItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("driveItemVersionId", "driveItemVersionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Version ID
func (id MeDriveIdItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Drive Item Version: %q", id.DriveItemVersionId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Version (%s)", strings.Join(components, "\n"))
}
