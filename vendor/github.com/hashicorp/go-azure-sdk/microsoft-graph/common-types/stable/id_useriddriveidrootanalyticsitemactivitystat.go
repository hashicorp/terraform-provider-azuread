package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdRootAnalyticsItemActivityStatId{}

// UserIdDriveIdRootAnalyticsItemActivityStatId is a struct representing the Resource ID for a User Id Drive Id Root Analytics Item Activity Stat
type UserIdDriveIdRootAnalyticsItemActivityStatId struct {
	UserId             string
	DriveId            string
	ItemActivityStatId string
}

// NewUserIdDriveIdRootAnalyticsItemActivityStatID returns a new UserIdDriveIdRootAnalyticsItemActivityStatId struct
func NewUserIdDriveIdRootAnalyticsItemActivityStatID(userId string, driveId string, itemActivityStatId string) UserIdDriveIdRootAnalyticsItemActivityStatId {
	return UserIdDriveIdRootAnalyticsItemActivityStatId{
		UserId:             userId,
		DriveId:            driveId,
		ItemActivityStatId: itemActivityStatId,
	}
}

// ParseUserIdDriveIdRootAnalyticsItemActivityStatID parses 'input' into a UserIdDriveIdRootAnalyticsItemActivityStatId
func ParseUserIdDriveIdRootAnalyticsItemActivityStatID(input string) (*UserIdDriveIdRootAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdDriveIdRootAnalyticsItemActivityStatIDInsensitively parses 'input' case-insensitively into a UserIdDriveIdRootAnalyticsItemActivityStatId
// note: this method should only be used for API response data and not user input
func ParseUserIdDriveIdRootAnalyticsItemActivityStatIDInsensitively(input string) (*UserIdDriveIdRootAnalyticsItemActivityStatId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdDriveIdRootAnalyticsItemActivityStatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdDriveIdRootAnalyticsItemActivityStatId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdDriveIdRootAnalyticsItemActivityStatId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateUserIdDriveIdRootAnalyticsItemActivityStatID checks that 'input' can be parsed as a User Id Drive Id Root Analytics Item Activity Stat ID
func ValidateUserIdDriveIdRootAnalyticsItemActivityStatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdDriveIdRootAnalyticsItemActivityStatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Drive Id Root Analytics Item Activity Stat ID
func (id UserIdDriveIdRootAnalyticsItemActivityStatId) ID() string {
	fmtString := "/users/%s/drives/%s/root/analytics/itemActivityStats/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DriveId, id.ItemActivityStatId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Drive Id Root Analytics Item Activity Stat ID
func (id UserIdDriveIdRootAnalyticsItemActivityStatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
	}
}

// String returns a human-readable description of this User Id Drive Id Root Analytics Item Activity Stat ID
func (id UserIdDriveIdRootAnalyticsItemActivityStatId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
	}
	return fmt.Sprintf("User Id Drive Id Root Analytics Item Activity Stat (%s)", strings.Join(components, "\n"))
}
