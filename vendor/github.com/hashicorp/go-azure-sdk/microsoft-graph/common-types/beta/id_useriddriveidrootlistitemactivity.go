package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootListItemActivityId{}

// UserIdDriveIdRootListItemActivityId is a struct representing the Resource ID for a User Id Drive Id Root List Item Activity
type UserIdDriveIdRootListItemActivityId struct {
	UserId            string
	DriveId           string
	ItemActivityOLDId string
}

// NewUserIdDriveIdRootListItemActivityID returns a new UserIdDriveIdRootListItemActivityId struct
func NewUserIdDriveIdRootListItemActivityID(userId string, driveId string, itemActivityOLDId string) UserIdDriveIdRootListItemActivityId {
	return UserIdDriveIdRootListItemActivityId{
		UserId:            userId,
		DriveId:           driveId,
		ItemActivityOLDId: itemActivityOLDId,
	}
}

// ParseUserIdDriveIdRootListItemActivityID parses 'input' into a UserIdDriveIdRootListItemActivityId
func ParseUserIdDriveIdRootListItemActivityID(input string) (*UserIdDriveIdRootListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootListItemActivityIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootListItemActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootListItemActivityIDInsensitively(input string) (*UserIdDriveIdRootListItemActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootListItemActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootListItemActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootListItemActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityOLDId, ok = input.Parsed["itemActivityOLDId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityOLDId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootListItemActivityID checks that 'input' can be parsed as a User Id Drive Id Root List Item Activity ID
func ValidateUserIdDriveIdRootListItemActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootListItemActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root List Item Activity ID
func (id UserIdDriveIdRootListItemActivityId) ID() string {
	fmtString := "/users/%s/drives/%s/root/listItem/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ItemActivityOLDId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root List Item Activity ID
func (id UserIdDriveIdRootListItemActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("listItem", "listItem", "listItem"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityOLDId", "itemActivityOLDId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root List Item Activity ID
func (id UserIdDriveIdRootListItemActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity OLD: %q", id.ItemActivityOLDId),
	}
	return fmt.Sprintf("User Id Drive Id Root List Item Activity (%s)", strings.Join(components, "\n"))
}
