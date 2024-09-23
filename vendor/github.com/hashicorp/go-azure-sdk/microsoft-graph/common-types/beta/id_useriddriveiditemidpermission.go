package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdPermissionId{}

// UserIdDriveIdItemIdPermissionId is a struct representing the Resource ID for a User Id Drive Id Item Id Permission
type UserIdDriveIdItemIdPermissionId struct {
	UserId       string
	DriveId      string
	DriveItemId  string
	PermissionId string
}

// NewUserIdDriveIdItemIdPermissionID returns a new UserIdDriveIdItemIdPermissionId struct
func NewUserIdDriveIdItemIdPermissionID(userId string, driveId string, driveItemId string, permissionId string) UserIdDriveIdItemIdPermissionId {
	return UserIdDriveIdItemIdPermissionId{
		UserId:       userId,
		DriveId:      driveId,
		DriveItemId:  driveItemId,
		PermissionId: permissionId,
	}
}

// ParseUserIdDriveIdItemIdPermissionID parses 'input' into a UserIdDriveIdItemIdPermissionId
func ParseUserIdDriveIdItemIdPermissionID(input string) (*UserIdDriveIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdPermissionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdPermissionIDInsensitively(input string) (*UserIdDriveIdItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdItemIdPermissionID checks that 'input' can be parsed as a User Id Drive Id Item Id Permission ID
func ValidateUserIdDriveIdItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Permission ID
func (id UserIdDriveIdItemIdPermissionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Permission ID
func (id UserIdDriveIdItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Permission ID
func (id UserIdDriveIdItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Permission (%s)", strings.Join(components, "\n"))
}
