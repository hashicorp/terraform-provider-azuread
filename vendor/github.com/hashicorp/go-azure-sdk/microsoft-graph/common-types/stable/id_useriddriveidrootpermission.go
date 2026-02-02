package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootPermissionId{}

// UserIdDriveIdRootPermissionId is a struct representing the Resource ID for a User Id Drive Id Root Permission
type UserIdDriveIdRootPermissionId struct {
	UserId       string
	DriveId      string
	PermissionId string
}

// NewUserIdDriveIdRootPermissionID returns a new UserIdDriveIdRootPermissionId struct
func NewUserIdDriveIdRootPermissionID(userId string, driveId string, permissionId string) UserIdDriveIdRootPermissionId {
	return UserIdDriveIdRootPermissionId{
		UserId:       userId,
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseUserIdDriveIdRootPermissionID parses 'input' into a UserIdDriveIdRootPermissionId
func ParseUserIdDriveIdRootPermissionID(input string) (*UserIdDriveIdRootPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootPermissionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootPermissionIDInsensitively(input string) (*UserIdDriveIdRootPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootPermissionID checks that 'input' can be parsed as a User Id Drive Id Root Permission ID
func ValidateUserIdDriveIdRootPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Permission ID
func (id UserIdDriveIdRootPermissionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/permissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Permission ID
func (id UserIdDriveIdRootPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Permission ID
func (id UserIdDriveIdRootPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("User Id Drive Id Root Permission (%s)", strings.Join(components, "\n"))
}
