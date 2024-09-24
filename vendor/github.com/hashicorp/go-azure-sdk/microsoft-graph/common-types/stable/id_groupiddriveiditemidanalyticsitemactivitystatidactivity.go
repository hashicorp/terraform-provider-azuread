package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}

// GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId is a struct representing the Resource ID for a Group Id Drive Id Item Id Analytics Item Activity Stat Id Activity
type GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId struct {
	GroupId            string
	DriveId            string
	DriveItemId        string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID returns a new GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId struct
func NewGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(groupId string, driveId string, driveItemId string, itemActivityStatId string, itemActivityId string) GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId {
	return GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{
		GroupId:            groupId,
		DriveId:            driveId,
		DriveItemId:        driveItemId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID parses 'input' into a GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId
func ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(input string) (*GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityIDInsensitively(input string) (*GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
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

	if id.ItemActivityId, ok = input.Parsed["itemActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemActivityId", input)
	}

	return nil
}

// ValidateGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID checks that 'input' can be parsed as a Group Id Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func ValidateGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func (id GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/items/%s/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.DriveItemId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func (id GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("driveItemId", "driveItemId"),
		resourceids.StaticSegment("analytics", "analytics", "analytics"),
		resourceids.StaticSegment("itemActivityStats", "itemActivityStats", "itemActivityStats"),
		resourceids.UserSpecifiedSegment("itemActivityStatId", "itemActivityStatId"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.UserSpecifiedSegment("itemActivityId", "itemActivityId"),
	}
}

// String returns a human-readable description of this Group Id Drive Id Item Id Analytics Item Activity Stat Id Activity ID
func (id GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Drive Item: %q", id.DriveItemId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Group Id Drive Id Item Id Analytics Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
