package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdChildId{}

// UserIdDriveIdItemIdChildId is a struct representing the Resource ID for a User Id Drive Id Item Id Child
type UserIdDriveIdItemIdChildId struct {
	UserId       string
	DriveId      string
	DriveItemId  string
	DriveItemId1 string
}

// NewUserIdDriveIdItemIdChildID returns a new UserIdDriveIdItemIdChildId struct
func NewUserIdDriveIdItemIdChildID(userId string, driveId string, driveItemId string, driveItemId1 string) UserIdDriveIdItemIdChildId {
	return UserIdDriveIdItemIdChildId{
		UserId:       userId,
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		DriveItemId1: driveItemId1,
	}
}

// ParseUserIdDriveIdItemIdChildID parses 'input' into a UserIdDriveIdItemIdChildId
func ParseUserIdDriveIdItemIdChildID(input string) (*UserIdDriveIdItemIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdChildIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdChildId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdChildIDInsensitively(input string) (*UserIdDriveIdItemIdChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdChildId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DriveItemId1, ok = input.Parsed["driveItemId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveItemId1", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdChildID checks that 'input' can be parsed as a User Id Drive Id Item Id Child ID
func ValidateUserIdDriveIdItemIdChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Child ID
func (id UserIdDriveIdItemIdChildId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/children/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.DriveItemId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Child ID
func (id UserIdDriveIdItemIdChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("driveItemId1", "driveItemId1"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Child ID
func (id UserIdDriveIdItemIdChildId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Drive Item Id 1: %q", id.DriveItemId1),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Child (%s)", strings.Join(components, "\n"))
}
