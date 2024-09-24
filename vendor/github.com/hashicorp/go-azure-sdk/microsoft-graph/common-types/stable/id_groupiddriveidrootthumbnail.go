package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootThumbnailId{}

// GroupIdDriveIdRootThumbnailId is a struct representing the Resource ID for a Group Id Drive Id Root Thumbnail
type GroupIdDriveIdRootThumbnailId struct {
	GroupId        string
	DriveId        string
	ThumbnailSetId string
}

// NewGroupIdDriveIdRootThumbnailID returns a new GroupIdDriveIdRootThumbnailId struct
func NewGroupIdDriveIdRootThumbnailID(groupId string, driveId string, thumbnailSetId string) GroupIdDriveIdRootThumbnailId {
	return GroupIdDriveIdRootThumbnailId{
		GroupId:        groupId,
		DriveId:        driveId,
		ThumbnailSetId: thumbnailSetId,
	}
}

// ParseGroupIdDriveIdRootThumbnailID parses 'input' into a GroupIdDriveIdRootThumbnailId
func ParseGroupIdDriveIdRootThumbnailID(input string) (*GroupIdDriveIdRootThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootThumbnailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootThumbnailIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootThumbnailId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootThumbnailIDInsensitively(input string) (*GroupIdDriveIdRootThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootThumbnailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootThumbnailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ThumbnailSetId, ok = input.Parsed["thumbnailSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "thumbnailSetId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdRootThumbnailID checks that 'input' can be parsed as a Group Id Drive Id Root Thumbnail ID
func ValidateGroupIdDriveIdRootThumbnailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootThumbnailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Thumbnail ID
func (id GroupIdDriveIdRootThumbnailId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/thumbnails/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ThumbnailSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Thumbnail ID
func (id GroupIdDriveIdRootThumbnailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("thumbnails", "thumbnails", "thumbnails"),
		resourceids.UserSpecifiedSegment("thumbnailSetId", "thumbnailSetId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Root Thumbnail ID
func (id GroupIdDriveIdRootThumbnailId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Thumbnail Set: %q", id.ThumbnailSetId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Thumbnail (%s)", strings.Join(components, "\n"))
}
