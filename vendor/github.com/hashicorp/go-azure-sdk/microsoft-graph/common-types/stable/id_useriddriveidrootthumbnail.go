package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootThumbnailId{}

// UserIdDriveIdRootThumbnailId is a struct representing the Resource ID for a User Id Drive Id Root Thumbnail
type UserIdDriveIdRootThumbnailId struct {
	UserId         string
	DriveId        string
	ThumbnailSetId string
}

// NewUserIdDriveIdRootThumbnailID returns a new UserIdDriveIdRootThumbnailId struct
func NewUserIdDriveIdRootThumbnailID(userId string, driveId string, thumbnailSetId string) UserIdDriveIdRootThumbnailId {
	return UserIdDriveIdRootThumbnailId{
		UserId:         userId,
		DriveId:        driveId,
		ThumbnailSetId: thumbnailSetId,
	}
}

// ParseUserIdDriveIdRootThumbnailID parses 'input' into a UserIdDriveIdRootThumbnailId
func ParseUserIdDriveIdRootThumbnailID(input string) (*UserIdDriveIdRootThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootThumbnailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootThumbnailIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootThumbnailId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootThumbnailIDInsensitively(input string) (*UserIdDriveIdRootThumbnailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootThumbnailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootThumbnailId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootThumbnailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ThumbnailSetId, ok = input.Parsed["thumbnailSetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "thumbnailSetId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootThumbnailID checks that 'input' can be parsed as a User Id Drive Id Root Thumbnail ID
func ValidateUserIdDriveIdRootThumbnailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootThumbnailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Thumbnail ID
func (id UserIdDriveIdRootThumbnailId) ID() string {
	fmtString := "/users/%s/drives/%s/root/thumbnails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ThumbnailSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Thumbnail ID
func (id UserIdDriveIdRootThumbnailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("thumbnails", "thumbnails", "thumbnails"),
		resourceids.UserSpecifiedSegment("thumbnailSetId", "thumbnailSetId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Thumbnail ID
func (id UserIdDriveIdRootThumbnailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Thumbnail Set: %q", id.ThumbnailSetId),
	}
	return fmt.Sprintf("User Id Drive Id Root Thumbnail (%s)", strings.Join(components, "\n"))
}
