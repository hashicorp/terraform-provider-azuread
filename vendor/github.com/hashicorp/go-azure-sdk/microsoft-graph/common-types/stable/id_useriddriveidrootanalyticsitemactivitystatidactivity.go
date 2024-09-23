package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId{}

// UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId is a struct representing the Resource ID for a User Id Drive Id Root Analytics Item Activity Stat Id Activity
type UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId struct {
	UserId             string
	DriveId            string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID returns a new UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId struct
func NewUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID(userId string, driveId string, itemActivityStatId string, itemActivityId string) UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId {
	return UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId{
		UserId:             userId,
		DriveId:            driveId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID parses 'input' into a UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId
func ParseUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID(input string) (*UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively(input string) (*UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	if id.ItemActivityId, ok = input.Parsed["itemActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", input)
	}

	return nil
}

// ValidateUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID checks that 'input' can be parsed as a User Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func ValidateUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootAnalyticsItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId) ID() string {
	fmtString := "/users/%s/drives/%s/root/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("User Id Drive Id Root Analytics Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
