package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootListItemPermissionId{}

// UserIdDriveIdRootListItemPermissionId is a struct representing the Resource ID for a User Id Drive Id Root List Item Permission
type UserIdDriveIdRootListItemPermissionId struct {
	UserId       string
	DriveId      string
	PermissionId string
}

// NewUserIdDriveIdRootListItemPermissionID returns a new UserIdDriveIdRootListItemPermissionId struct
func NewUserIdDriveIdRootListItemPermissionID(userId string, driveId string, permissionId string) UserIdDriveIdRootListItemPermissionId {
	return UserIdDriveIdRootListItemPermissionId{
		UserId:       userId,
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseUserIdDriveIdRootListItemPermissionID parses 'input' into a UserIdDriveIdRootListItemPermissionId
func ParseUserIdDriveIdRootListItemPermissionID(input string) (*UserIdDriveIdRootListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootListItemPermissionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootListItemPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootListItemPermissionIDInsensitively(input string) (*UserIdDriveIdRootListItemPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootListItemPermissionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdRootListItemPermissionID checks that 'input' can be parsed as a User Id Drive Id Root List Item Permission ID
func ValidateUserIdDriveIdRootListItemPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootListItemPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root List Item Permission ID
func (id UserIdDriveIdRootListItemPermissionId) ID() string {
	fmtString := "/users/%s/drives/%s/root/listItem/permissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root List Item Permission ID
func (id UserIdDriveIdRootListItemPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root List Item Permission ID
func (id UserIdDriveIdRootListItemPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("User Id Drive Id Root List Item Permission (%s)", strings.Join(components, "\n"))
}
