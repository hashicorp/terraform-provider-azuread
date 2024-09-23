package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdThumbnailId{}

// UserIdDriveIdItemIdThumbnailId is a struct representing the Resource ID for a User Id Drive Id Item Id Thumbnail
type UserIdDriveIdItemIdThumbnailId struct {
	UserId         string
	DriveId        string
	DriveItemId    string
	ThumbnailSetId string
}

// NewUserIdDriveIdItemIdThumbnailID returns a new UserIdDriveIdItemIdThumbnailId struct
func NewUserIdDriveIdItemIdThumbnailID(userId string, driveId string, driveItemId string, thumbnailSetId string) UserIdDriveIdItemIdThumbnailId {
	return UserIdDriveIdItemIdThumbnailId{
		UserId:         userId,
		DriveId:        driveId,
		DriveItemId:    driveItemId,
		ThumbnailSetId: thumbnailSetId,
	}
}

// ParseUserIdDriveIdItemIdThumbnailID parses 'input' into a UserIdDriveIdItemIdThumbnailId
func ParseUserIdDriveIdItemIdThumbnailID(input string) (*UserIdDriveIdItemIdThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdThumbnailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdThumbnailIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdThumbnailId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdThumbnailIDInsensitively(input string) (*UserIdDriveIdItemIdThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdThumbnailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdThumbnailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
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

// ValidateUserIdDriveIdItemIdThumbnailID checks that 'input' can be parsed as a User Id Drive Id Item Id Thumbnail ID
func ValidateUserIdDriveIdItemIdThumbnailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdThumbnailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Thumbnail ID
func (id UserIdDriveIdItemIdThumbnailId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/thumbnails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.ThumbnailSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Thumbnail ID
func (id UserIdDriveIdItemIdThumbnailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("thumbnails", "thumbnails", "thumbnails"),
		resourceids.UserSpecifiedSegment("thumbnailSetId", "thumbnailSetId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Thumbnail ID
func (id UserIdDriveIdItemIdThumbnailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Thumbnail Set: %q", id.ThumbnailSetId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Thumbnail (%s)", strings.Join(components, "\n"))
}
