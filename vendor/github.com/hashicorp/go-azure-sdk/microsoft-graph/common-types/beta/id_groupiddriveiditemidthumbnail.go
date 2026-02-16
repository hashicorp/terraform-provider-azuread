package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdThumbnailId{}

// GroupIdDriveIdItemIdThumbnailId is a struct representing the Resource ID for a Group Id Drive Id Item Id Thumbnail
type GroupIdDriveIdItemIdThumbnailId struct {
	GroupId        string
	DriveId        string
	DriveItemId    string
	ThumbnailSetId string
}

// NewGroupIdDriveIdItemIdThumbnailID returns a new GroupIdDriveIdItemIdThumbnailId struct
func NewGroupIdDriveIdItemIdThumbnailID(groupId string, driveId string, driveItemId string, thumbnailSetId string) GroupIdDriveIdItemIdThumbnailId {
	return GroupIdDriveIdItemIdThumbnailId{
		GroupId:        groupId,
		DriveId:        driveId,
		DriveItemId:    driveItemId,
		ThumbnailSetId: thumbnailSetId,
	}
}

// ParseGroupIdDriveIdItemIdThumbnailID parses 'input' into a GroupIdDriveIdItemIdThumbnailId
func ParseGroupIdDriveIdItemIdThumbnailID(input string) (*GroupIdDriveIdItemIdThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdThumbnailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdThumbnailIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdThumbnailId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdThumbnailIDInsensitively(input string) (*GroupIdDriveIdItemIdThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdThumbnailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdThumbnailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

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

// ValidateGroupIdDriveIdItemIdThumbnailID checks that 'input' can be parsed as a Group Id Drive Id Item Id Thumbnail ID
func ValidateGroupIdDriveIdItemIdThumbnailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdThumbnailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Thumbnail ID
func (id GroupIdDriveIdItemIdThumbnailId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/thumbnails/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ThumbnailSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Thumbnail ID
func (id GroupIdDriveIdItemIdThumbnailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("thumbnails", "thumbnails", "thumbnails"),
		resourceids.UserSpecifiedSegment("thumbnailSetId", "thumbnailSetId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Thumbnail ID
func (id GroupIdDriveIdItemIdThumbnailId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Thumbnail Set: %q", id.ThumbnailSetId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Thumbnail (%s)", strings.Join(components, "\n"))
}
