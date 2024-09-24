package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemVersionId{}

// UserIdDriveIdItemIdListItemVersionId is a struct representing the Resource ID for a User Id Drive Id Item Id List Item Version
type UserIdDriveIdItemIdListItemVersionId struct {
	UserId            string
	DriveId           string
	DriveItemId       string
	ListItemVersionId string
}

// NewUserIdDriveIdItemIdListItemVersionID returns a new UserIdDriveIdItemIdListItemVersionId struct
func NewUserIdDriveIdItemIdListItemVersionID(userId string, driveId string, driveItemId string, listItemVersionId string) UserIdDriveIdItemIdListItemVersionId {
	return UserIdDriveIdItemIdListItemVersionId{
		UserId:            userId,
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseUserIdDriveIdItemIdListItemVersionID parses 'input' into a UserIdDriveIdItemIdListItemVersionId
func ParseUserIdDriveIdItemIdListItemVersionID(input string) (*UserIdDriveIdItemIdListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdListItemVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdListItemVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdListItemVersionIDInsensitively(input string) (*UserIdDriveIdItemIdListItemVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdListItemVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdListItemVersionID checks that 'input' can be parsed as a User Id Drive Id Item Id List Item Version ID
func ValidateUserIdDriveIdItemIdListItemVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdListItemVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id List Item Version ID
func (id UserIdDriveIdItemIdListItemVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/listItem/versions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id List Item Version ID
func (id UserIdDriveIdItemIdListItemVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id List Item Version ID
func (id UserIdDriveIdItemIdListItemVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id List Item Version (%s)", strings.Join(components, "\n"))
}
