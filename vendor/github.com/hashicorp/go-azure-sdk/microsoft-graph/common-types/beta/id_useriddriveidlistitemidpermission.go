package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListItemIdPermissionId{}

// UserIdDriveIdListItemIdPermissionId is a struct representing the Resource ID for a User Id Drive Id List Item Id Permission
type UserIdDriveIdListItemIdPermissionId struct {
	UserId       string
	DriveId      string
	ListItemId   string
	PermissionId string
}

// NewUserIdDriveIdListItemIdPermissionID returns a new UserIdDriveIdListItemIdPermissionId struct
func NewUserIdDriveIdListItemIdPermissionID(userId string, driveId string, listItemId string, permissionId string) UserIdDriveIdListItemIdPermissionId {
	return UserIdDriveIdListItemIdPermissionId{
		UserId:       userId,
		DriveId:      driveId,
		ListItemId:   listItemId,
		PermissionId: permissionId,
	}
}

// ParseUserIdDriveIdListItemIdPermissionID parses 'input' into a UserIdDriveIdListItemIdPermissionId
func ParseUserIdDriveIdListItemIdPermissionID(input string) (*UserIdDriveIdListItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListItemIdPermissionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListItemIdPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListItemIdPermissionIDInsensitively(input string) (*UserIdDriveIdListItemIdPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListItemIdPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ListItemId, ok = input.Parsed["listItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListItemIdPermissionID checks that 'input' can be parsed as a User Id Drive Id List Item Id Permission ID
func ValidateUserIdDriveIdListItemIdPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListItemIdPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Item Id Permission ID
func (id UserIdDriveIdListItemIdPermissionId) ID() string {
	fmtString := "/users/%s/drives/%s/list/items/%s/permissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ListItemId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Item Id Permission ID
func (id UserIdDriveIdListItemIdPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Item Id Permission ID
func (id UserIdDriveIdListItemIdPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("User Id Drive Id List Item Id Permission (%s)", strings.Join(components, "\n"))
}
