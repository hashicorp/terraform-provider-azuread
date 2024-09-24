package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootThumbnailId{}

// MeDriveIdRootThumbnailId is a struct representing the Resource ID for a Me Drive Id Root Thumbnail
type MeDriveIdRootThumbnailId struct {
	DriveId        string
	ThumbnailSetId string
}

// NewMeDriveIdRootThumbnailID returns a new MeDriveIdRootThumbnailId struct
func NewMeDriveIdRootThumbnailID(driveId string, thumbnailSetId string) MeDriveIdRootThumbnailId {
	return MeDriveIdRootThumbnailId{
		DriveId:        driveId,
		ThumbnailSetId: thumbnailSetId,
	}
}

// ParseMeDriveIdRootThumbnailID parses 'input' into a MeDriveIdRootThumbnailId
func ParseMeDriveIdRootThumbnailID(input string) (*MeDriveIdRootThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootThumbnailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootThumbnailIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootThumbnailId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootThumbnailIDInsensitively(input string) (*MeDriveIdRootThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootThumbnailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootThumbnailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ThumbnailSetId, ok = input.Parsed["thumbnailSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "thumbnailSetId", input)
	}

	return nil
}

// ValidateMeDriveIdRootThumbnailID checks that 'input' can be parsed as a Me Drive Id Root Thumbnail ID
func ValidateMeDriveIdRootThumbnailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootThumbnailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Thumbnail ID
func (id MeDriveIdRootThumbnailId) ID() string {
	fmtString := "/me/drives/%s/root/thumbnails/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.ThumbnailSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Thumbnail ID
func (id MeDriveIdRootThumbnailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("thumbnails", "thumbnails", "thumbnails"),
		resourceids.UserSpecifiedSegment("thumbnailSetId", "thumbnailSetId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Thumbnail ID
func (id MeDriveIdRootThumbnailId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Thumbnail Set: %q", id.ThumbnailSetId),
	}
	return fmt.Sprintf("Me Drive Id Root Thumbnail (%s)", strings.Join(components, "\n"))
}
