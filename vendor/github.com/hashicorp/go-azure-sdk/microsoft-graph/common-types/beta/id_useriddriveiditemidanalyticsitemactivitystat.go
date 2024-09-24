package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdAnalyticsItemActivityStatId{}

// UserIdDriveIdItemIdAnalyticsItemActivityStatId is a struct representing the Resource ID for a User Id Drive Id Item Id Analytics Item Activity Stat
type UserIdDriveIdItemIdAnalyticsItemActivityStatId struct {
	UserId             string
	DriveId            string
	DriveItemId        string
	ItemActivityStatId string
}

// NewUserIdDriveIdItemIdAnalyticsItemActivityStatID returns a new UserIdDriveIdItemIdAnalyticsItemActivityStatId struct
func NewUserIdDriveIdItemIdAnalyticsItemActivityStatID(userId string, driveId string, driveItemId string, itemActivityStatId string) UserIdDriveIdItemIdAnalyticsItemActivityStatId {
	return UserIdDriveIdItemIdAnalyticsItemActivityStatId{
		UserId:             userId,
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseUserIdDriveIdItemIdAnalyticsItemActivityStatID parses 'input' into a UserIdDriveIdItemIdAnalyticsItemActivityStatId
func ParseUserIdDriveIdItemIdAnalyticsItemActivityStatID(input string) (*UserIdDriveIdItemIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdItemIdAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdItemIdAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdItemIdAnalyticsItemActivityStatIDInsensitively(input string) (*UserIdDriveIdItemIdAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdItemIdAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdItemIdAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdItemIdAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ItemActivityStatId, ok = input.Parsed["itemActivityStatId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityStatId", input)
	}

	return nil
}

// ValidateUserIdDriveIdItemIdAnalyticsItemActivityStatID checks that 'input' can be parsed as a User Id Drive Id Item Id Analytics Item Activity Stat ID
func ValidateUserIdDriveIdItemIdAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdItemIdAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Item Id Analytics Item Activity Stat ID
func (id UserIdDriveIdItemIdAnalyticsItemActivityStatId) ID() string {
	fmtString := "/users/%s/drives/%s/items/%s/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.DriveItemId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Item Id Analytics Item Activity Stat ID
func (id UserIdDriveIdItemIdAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Item Id Analytics Item Activity Stat ID
func (id UserIdDriveIdItemIdAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("User Id Drive Id Item Id Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
