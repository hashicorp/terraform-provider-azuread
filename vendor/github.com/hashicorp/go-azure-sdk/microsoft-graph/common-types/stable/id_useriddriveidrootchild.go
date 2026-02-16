package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootChildId{}

// UserIdDriveIdRootChildId is a struct representing the Resource ID for a User Id Drive Id Root Child
type UserIdDriveIdRootChildId struct {
	UserId      string
	DriveId     string
	DriveItemId string
}

// NewUserIdDriveIdRootChildID returns a new UserIdDriveIdRootChildId struct
func NewUserIdDriveIdRootChildID(userId string, driveId string, driveItemId string) UserIdDriveIdRootChildId {
	return UserIdDriveIdRootChildId{
		UserId:      userId,
		DriveId:     driveId,
		DriveItemId: driveItemId,
	}
}

// ParseUserIdDriveIdRootChildID parses 'input' into a UserIdDriveIdRootChildId
func ParseUserIdDriveIdRootChildID(input string) (*UserIdDriveIdRootChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootChildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootChildIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootChildId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootChildIDInsensitively(input string) (*UserIdDriveIdRootChildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootChildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootChildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootChildId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdRootChildID checks that 'input' can be parsed as a User Id Drive Id Root Child ID
func ValidateUserIdDriveIdRootChildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootChildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Child ID
func (id UserIdDriveIdRootChildId) ID() string {
	fmtString := "/users/%s/drives/%s/root/children/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Child ID
func (id UserIdDriveIdRootChildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Child ID
func (id UserIdDriveIdRootChildId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
	}
	return fmt.Sprintf("User Id Drive Id Root Child (%s)", strings.Join(components, "\n"))
}
