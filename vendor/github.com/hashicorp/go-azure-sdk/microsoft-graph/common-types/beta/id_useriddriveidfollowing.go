package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdFollowingId{}

// UserIdDriveIdFollowingId is a struct representing the Resource ID for a User Id Drive Id Following
type UserIdDriveIdFollowingId struct {
	UserId      string
	DriveId     string
	DriveItemId string
}

// NewUserIdDriveIdFollowingID returns a new UserIdDriveIdFollowingId struct
func NewUserIdDriveIdFollowingID(userId string, driveId string, driveItemId string) UserIdDriveIdFollowingId {
	return UserIdDriveIdFollowingId{
		UserId:      userId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseUserIdDriveIdFollowingID parses 'input' into a UserIdDriveIdFollowingId
func ParseUserIdDriveIdFollowingID(input string) (*UserIdDriveIdFollowingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdFollowingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdFollowingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdFollowingIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdFollowingId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdFollowingIDInsensitively(input string) (*UserIdDriveIdFollowingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdFollowingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdFollowingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdFollowingId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdDriveIdFollowingID checks that 'input' can be parsed as a User Id Drive Id Following ID
func ValidateUserIdDriveIdFollowingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdFollowingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Following ID
func (id UserIdDriveIdFollowingId) ID() string {
	fmtString := "/users/%s/drives/%s/following/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Following ID
func (id UserIdDriveIdFollowingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("following", "following", "following"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Following ID
func (id UserIdDriveIdFollowingId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("User Id Drive Id Following (%s)", strings.Join(components, "\n"))
}
