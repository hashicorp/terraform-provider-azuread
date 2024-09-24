package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemId{}

// UserIdDriveIdItemId is a struct representing the Resource ID for a User Id Drive Id Item
type UserIdDriveIdItemId struct {
	UserId      string
	DriveId     string
	DriveItemId string
}

// NewUserIdDriveIdItemID returns a new UserIdDriveIdItemId struct
func NewUserIdDriveIdItemID(userId string, driveId string, driveItemId string) UserIdDriveIdItemId {
	return UserIdDriveIdItemId{
		UserId:      userId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseUserIdDriveIdItemID parses 'input' into a UserIdDriveIdItemId
func ParseUserIdDriveIdItemID(input string) (*UserIdDriveIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIDInsensitively(input string) (*UserIdDriveIdItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdItemID checks that 'input' can be parsed as a User Id Drive Id Item ID
func ValidateUserIdDriveIdItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item ID
func (id UserIdDriveIdItemId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item ID
func (id UserIdDriveIdItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item ID
func (id UserIdDriveIdItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("User Id Drive Id Item (%s)", strings.Join(components, "\n"))
}
