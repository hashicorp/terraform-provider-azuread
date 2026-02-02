package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListItemId{}

// UserIdDriveIdListItemId is a struct representing the Resource ID for a User Id Drive Id List Item
type UserIdDriveIdListItemId struct {
	UserId     string
	DriveId    string
	ListItemId string
}

// NewUserIdDriveIdListItemID returns a new UserIdDriveIdListItemId struct
func NewUserIdDriveIdListItemID(userId string, driveId string, listItemId string) UserIdDriveIdListItemId {
	return UserIdDriveIdListItemId{
		UserId:     userId,
		DriveId:    driveId,
		ListItemId: listItemId,
	}
}

// ParseUserIdDriveIdListItemID parses 'input' into a UserIdDriveIdListItemId
func ParseUserIdDriveIdListItemID(input string) (*UserIdDriveIdListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListItemIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListItemId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListItemIDInsensitively(input string) (*UserIdDriveIdListItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListItemId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdDriveIdListItemID checks that 'input' can be parsed as a User Id Drive Id List Item ID
func ValidateUserIdDriveIdListItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Item ID
func (id UserIdDriveIdListItemId) ID() string {
	fmtString := "/users/%s/drives/%s/list/items/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ListItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Item ID
func (id UserIdDriveIdListItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Item ID
func (id UserIdDriveIdListItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
	}
	return fmt.Sprintf("User Id Drive Id List Item (%s)", strings.Join(components, "\n"))
}
