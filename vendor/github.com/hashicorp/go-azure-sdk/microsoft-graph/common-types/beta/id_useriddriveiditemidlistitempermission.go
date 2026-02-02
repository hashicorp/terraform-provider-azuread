package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemPermissionId{}

// UserIdDriveIdItemIdListItemPermissionId is a struct representing the Resource ID for a User Id Drive Id Item Id List Item Permission
type UserIdDriveIdItemIdListItemPermissionId struct {
	UserId       string
	DriveId      string
	DriveItemId  string
	PermissionId string
}

// NewUserIdDriveIdItemIdListItemPermissionID returns a new UserIdDriveIdItemIdListItemPermissionId struct
func NewUserIdDriveIdItemIdListItemPermissionID(userId string, driveId string, driveItemId string, permissionId string) UserIdDriveIdItemIdListItemPermissionId {
	return UserIdDriveIdItemIdListItemPermissionId{
		UserId:       userId,
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		PermissionId: permissionId,
	}
}

// ParseUserIdDriveIdItemIdListItemPermissionID parses 'input' into a UserIdDriveIdItemIdListItemPermissionId
func ParseUserIdDriveIdItemIdListItemPermissionID(input string) (*UserIdDriveIdItemIdListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdListItemPermissionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdListItemPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdListItemPermissionIDInsensitively(input string) (*UserIdDriveIdItemIdListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdListItemPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdListItemPermissionID checks that 'input' can be parsed as a User Id Drive Id Item Id List Item Permission ID
func ValidateUserIdDriveIdItemIdListItemPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdListItemPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id List Item Permission ID
func (id UserIdDriveIdItemIdListItemPermissionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/listItem/permissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id List Item Permission ID
func (id UserIdDriveIdItemIdListItemPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id List Item Permission ID
func (id UserIdDriveIdItemIdListItemPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id List Item Permission (%s)", strings.Join(components, "\n"))
}
