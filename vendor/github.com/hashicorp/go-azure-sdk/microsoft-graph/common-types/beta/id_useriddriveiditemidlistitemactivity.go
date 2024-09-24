package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemActivityId{}

// UserIdDriveIdItemIdListItemActivityId is a struct representing the Resource ID for a User Id Drive Id Item Id List Item Activity
type UserIdDriveIdItemIdListItemActivityId struct {
	UserId            string
	DriveId           string
	DriveItemId       string
	ItemActivityOLDId string
}

// NewUserIdDriveIdItemIdListItemActivityID returns a new UserIdDriveIdItemIdListItemActivityId struct
func NewUserIdDriveIdItemIdListItemActivityID(userId string, driveId string, driveItemId string, itemActivityOLDId string) UserIdDriveIdItemIdListItemActivityId {
	return UserIdDriveIdItemIdListItemActivityId{
		UserId:            userId,
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseUserIdDriveIdItemIdListItemActivityID parses 'input' into a UserIdDriveIdItemIdListItemActivityId
func ParseUserIdDriveIdItemIdListItemActivityID(input string) (*UserIdDriveIdItemIdListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdListItemActivityIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdListItemActivityIDInsensitively(input string) (*UserIdDriveIdItemIdListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdListItemActivityId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdListItemActivityID checks that 'input' can be parsed as a User Id Drive Id Item Id List Item Activity ID
func ValidateUserIdDriveIdItemIdListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id List Item Activity ID
func (id UserIdDriveIdItemIdListItemActivityId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/listItem/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id List Item Activity ID
func (id UserIdDriveIdItemIdListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id List Item Activity ID
func (id UserIdDriveIdItemIdListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id List Item Activity (%s)", strings.Join(components, "\n"))
}
