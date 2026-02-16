package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListPermissionId{}

// UserIdDriveIdListPermissionId is a struct representing the Resource ID for a User Id Drive Id List Permission
type UserIdDriveIdListPermissionId struct {
	UserId       string
	DriveId      string
	PermissionId string
}

// NewUserIdDriveIdListPermissionID returns a new UserIdDriveIdListPermissionId struct
func NewUserIdDriveIdListPermissionID(userId string, driveId string, permissionId string) UserIdDriveIdListPermissionId {
	return UserIdDriveIdListPermissionId{
		UserId:       userId,
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseUserIdDriveIdListPermissionID parses 'input' into a UserIdDriveIdListPermissionId
func ParseUserIdDriveIdListPermissionID(input string) (*UserIdDriveIdListPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListPermissionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListPermissionIDInsensitively(input string) (*UserIdDriveIdListPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdListPermissionID checks that 'input' can be parsed as a User Id Drive Id List Permission ID
func ValidateUserIdDriveIdListPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Permission ID
func (id UserIdDriveIdListPermissionId) ID() string {
	fmtString := "/users/%s/drives/%s/list/permissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Permission ID
func (id UserIdDriveIdListPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Permission ID
func (id UserIdDriveIdListPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("User Id Drive Id List Permission (%s)", strings.Join(components, "\n"))
}
