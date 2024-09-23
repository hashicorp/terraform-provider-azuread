package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdItemIdThumbnailId{}

// MeDriveIdItemIdThumbnailId is a struct representing the Resource ID for a Me Drive Id Item Id Thumbnail
type MeDriveIdItemIdThumbnailId struct {
	DriveId        string
	DriveItemId    string
	ThumbnailSetId string
}

// NewMeDriveIdItemIdThumbnailID returns a new MeDriveIdItemIdThumbnailId struct
func NewMeDriveIdItemIdThumbnailID(driveId string, driveItemId string, thumbnailSetId string) MeDriveIdItemIdThumbnailId {
	return MeDriveIdItemIdThumbnailId{
		DriveId:        driveId,
		DriveItemId:    driveItemId,
		ThumbnailSetId: thumbnailSetId,
	}
}

// ParseMeDriveIdItemIdThumbnailID parses 'input' into a MeDriveIdItemIdThumbnailId
func ParseMeDriveIdItemIdThumbnailID(input string) (*MeDriveIdItemIdThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdThumbnailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdItemIdThumbnailIDInsensitively parses 'input' case-insensitively into a MeDriveIdItemIdThumbnailId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdItemIdThumbnailIDInsensitively(input string) (*MeDriveIdItemIdThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdItemIdThumbnailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdItemIdThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdItemIdThumbnailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.DriveItemId, ok = input.Parsed["driveItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId", input)
	}

	if id.ThumbnailSetId, ok = input.Parsed["thumbnailSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "thumbnailSetId", input)
	}

	return nil
}

// ValidateMeDriveIdItemIdThumbnailID checks that 'input' can be parsed as a Me Drive Id Item Id Thumbnail ID
func ValidateMeDriveIdItemIdThumbnailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdItemIdThumbnailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Item Id Thumbnail ID
func (id MeDriveIdItemIdThumbnailId) ID() string {
	fmtString := "/me/drives/%s/items/%s/thumbnails/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.DriveItemId, id.ThumbnailSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Item Id Thumbnail ID
func (id MeDriveIdItemIdThumbnailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("thumbnails", "thumbnails", "thumbnails"),
		resourceids.UserSpecifiedSegment("thumbnailSetId", "thumbnailSetId"),
	}
}

// String returns a human-readable description of this Me Drive Id Item Id Thumbnail ID
func (id MeDriveIdItemIdThumbnailId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Thumbnail Set: %q", id.ThumbnailSetId),
	}
	return fmt.Sprintf("Me Drive Id Item Id Thumbnail (%s)", strings.Join(components, "\n"))
}
