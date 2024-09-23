package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId{}

// GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId is a struct representing the Resource ID for a Group Id Drive Id Root Analytics Item Activity Stat Id Activity
type GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId struct {
	GroupId            string
	DriveId            string
	ItemActivityStatId string
	ItemActivityId     string
}

// NewGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID returns a new GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId struct
func NewGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID(groupId string, driveId string, itemActivityStatId string, itemActivityId string) GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId {
	return GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId{
		GroupId:            groupId,
		DriveId:            driveId,
		ItemActivityStatId: itemActivityStatId,
		ItemActivityId:     itemActivityId,
	}
}

// ParseGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID parses 'input' into a GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId
func ParseGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID(input string) (*GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively parses 'input' case-insensitively into a GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId
// note: this method should only be used for API response data and not user input
func ParseGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityIDInsensitively(input string) (*GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
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

// ValidateGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID checks that 'input' can be parsed as a Group Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func ValidateGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdDriveIdRootAnalyticsItemActivityStatIdActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId) ID() string {
	fmtString := "/groups/%s/drives/%s/root/analytics/itemActivityStats/%s/activities/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DriveId, id.ItemActivityStatId, id.ItemActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
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

// String returns a human-readable description of this Group Id Drive Id Root Analytics Item Activity Stat Id Activity ID
func (id GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Item Activity Stat: %q", id.ItemActivityStatId),
		fmt.Sprintf("Item Activity: %q", id.ItemActivityId),
	}
	return fmt.Sprintf("Group Id Drive Id Root Analytics Item Activity Stat Id Activity (%s)", strings.Join(components, "\n"))
}
