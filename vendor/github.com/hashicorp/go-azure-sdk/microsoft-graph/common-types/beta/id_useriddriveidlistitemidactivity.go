package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdListItemIdActivityId{}

// UserIdDriveIdListItemIdActivityId is a struct representing the Resource ID for a User Id Drive Id List Item Id Activity
type UserIdDriveIdListItemIdActivityId struct {
	UserId            string
	DriveId           string
	ListItemId        string
	ItemActivityOLDId string
}

// NewUserIdDriveIdListItemIdActivityID returns a new UserIdDriveIdListItemIdActivityId struct
func NewUserIdDriveIdListItemIdActivityID(userId string, driveId string, listItemId string, itemActivityOLDId string) UserIdDriveIdListItemIdActivityId {
	return UserIdDriveIdListItemIdActivityId{
		UserId:            userId,
		DriveId:           driveId,
		ListItemId:        listItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseUserIdDriveIdListItemIdActivityID parses 'input' into a UserIdDriveIdListItemIdActivityId
func ParseUserIdDriveIdListItemIdActivityID(input string) (*UserIdDriveIdListItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdListItemIdActivityIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdListItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdListItemIdActivityIDInsensitively(input string) (*UserIdDriveIdListItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdListItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdListItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdListItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateUserIdDriveIdListItemIdActivityID checks that 'input' can be parsed as a User Id Drive Id List Item Id Activity ID
func ValidateUserIdDriveIdListItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdListItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id List Item Id Activity ID
func (id UserIdDriveIdListItemIdActivityId) ID() string {
	fmtString := "/users/%s/drives/%s/list/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ListItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id List Item Id Activity ID
func (id UserIdDriveIdListItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("list", "list", "list"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("listItemId", "listItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this User Id Drive Id List Item Id Activity ID
func (id UserIdDriveIdListItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("List Item: %q", id.ListItemId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("User Id Drive Id List Item Id Activity (%s)", strings.Join(components, "\n"))
}
