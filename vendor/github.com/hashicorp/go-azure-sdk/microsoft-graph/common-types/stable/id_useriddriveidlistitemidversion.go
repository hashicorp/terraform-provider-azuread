package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListItemIdVersionId{}

// UserIdDriveIdListItemIdVersionId is a struct representing the Resource ID for a User Id Drive Id List Item Id Version
type UserIdDriveIdListItemIdVersionId struct {
	UserId            string
	DriveId           string
	ListItemId        string
	ListItemVersionId string
}

// NewUserIdDriveIdListItemIdVersionID returns a new UserIdDriveIdListItemIdVersionId struct
func NewUserIdDriveIdListItemIdVersionID(userId string, driveId string, listItemId string, listItemVersionId string) UserIdDriveIdListItemIdVersionId {
	return UserIdDriveIdListItemIdVersionId{
		UserId:            userId,
		DriveId:           driveId,
		ListItemId:        listItemId,
		ListItemVersionId: listItemVersionId,
	}
}

// ParseUserIdDriveIdListItemIdVersionID parses 'input' into a UserIdDriveIdListItemIdVersionId
func ParseUserIdDriveIdListItemIdVersionID(input string) (*UserIdDriveIdListItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListItemIdVersionIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListItemIdVersionId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListItemIdVersionIDInsensitively(input string) (*UserIdDriveIdListItemIdVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListItemIdVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ListItemVersionId, ok = input.Parsed["listItemVersionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "listItemVersionId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListItemIdVersionID checks that 'input' can be parsed as a User Id Drive Id List Item Id Version ID
func ValidateUserIdDriveIdListItemIdVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListItemIdVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Item Id Version ID
func (id UserIdDriveIdListItemIdVersionId) ID() string {
	fmtString := "/users/%s/drives/%s/list/items/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ListItemId, id.ListItemVersionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Item Id Version ID
func (id UserIdDriveIdListItemIdVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("versions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("listItemVersionId", "listItemVersionId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Item Id Version ID
func (id UserIdDriveIdListItemIdVersionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("List Item Version: %q", id.ListItemVersionId),
	}
	return fmt.Sprintf("User Id Drive Id List Item Id Version (%s)", strings.Join(components, "\n"))
}
