package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdActivityId{}

// UserIdDriveIdItemIdActivityId is a struct representing the Resource ID for a User Id Drive Id Item Id Activity
type UserIdDriveIdItemIdActivityId struct {
	UserId            string
	DriveId           string
	DriveItemId       string
	ItemActivityOLDId string
}

// NewUserIdDriveIdItemIdActivityID returns a new UserIdDriveIdItemIdActivityId struct
func NewUserIdDriveIdItemIdActivityID(userId string, driveId string, driveItemId string, itemActivityOLDId string) UserIdDriveIdItemIdActivityId {
	return UserIdDriveIdItemIdActivityId{
		UserId:            userId,
		DriveId:           driveId,
		DriveItemId:       driveItemId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseUserIdDriveIdItemIdActivityID parses 'input' into a UserIdDriveIdItemIdActivityId
func ParseUserIdDriveIdItemIdActivityID(input string) (*UserIdDriveIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdActivityIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdActivityIDInsensitively(input string) (*UserIdDriveIdItemIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdActivityId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdDriveIdItemIdActivityID checks that 'input' can be parsed as a User Id Drive Id Item Id Activity ID
func ValidateUserIdDriveIdItemIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Activity ID
func (id UserIdDriveIdItemIdActivityId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Activity ID
func (id UserIdDriveIdItemIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Activity ID
func (id UserIdDriveIdItemIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity O L D: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Activity (%s)", strings.Join(components, "\n"))
}
